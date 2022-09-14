package skupper_podman

import (
	"fmt"

	"github.com/skupperproject/skupper/api/types"
	v2 "github.com/skupperproject/skupper/api/types/v2"
	"github.com/skupperproject/skupper/client"
	"github.com/skupperproject/skupper/client/common"
	"github.com/skupperproject/skupper/client/container"
	"github.com/skupperproject/skupper/client/podman"
	"github.com/skupperproject/skupper/pkg/utils"
)

type SitePodman struct {
	*v2.SiteCommon
	IngressBindHost            string
	IngressBindInterRouterPort int
	IngressBindEdgePort        int
	ContainerNetwork           string
	PodmanEndpoint             string
	prepared                   bool
}

func (s *SitePodman) GetPlatform() string {
	return "podman"
}

func (s *SitePodman) GetIngressClasses() []string {
	return []string{"host"}
}

type SitePodmanHandler struct {
	cli *podman.PodmanRestClient
}

func NewSitePodmanHandler(cli *podman.PodmanRestClient) *SitePodmanHandler {
	return &SitePodmanHandler{
		cli: cli,
	}
}

func (s *SitePodmanHandler) Prepare(site v2.Site) (v2.Site, error) {
	podmanSite, ok := site.(*SitePodman)
	if !ok {
		return nil, fmt.Errorf("not a valid podman site definition")
	}

	// Validating basic info
	if err := podmanSite.ValidateMinimumRequirements(); err != nil {
		return nil, err
	}

	// Preparing site
	common.ConfigureSiteCredentials(podmanSite, podmanSite.IngressBindHost)
	s.ConfigurePodmanDeployments(podmanSite)

	if err := s.canCreate(podmanSite); err != nil {
		return nil, err
	}

	podmanSite.prepared = true
	return podmanSite, nil
}

func (s *SitePodmanHandler) ConfigurePodmanDeployments(site *SitePodman) {
	// Router Deployment
	volumeMounts := map[string]string{
		types.LocalServerSecret:      "/etc/skupper-router-certs/skupper-amqps/",
		types.TransportConfigMapName: "/etc/skupper-router/config/",
		"skupper-router-certs":       "/etc/skupper-router-certs",
	}
	if !site.IsEdge() {
		volumeMounts[types.SiteServerSecret] = "/etc/skupper-router-certs/skupper-internal/"
	}
	routerDepl := &SkupperDeploymentPodman{
		Name: types.TransportDeploymentName,
		SkupperDeploymentCommon: &v2.SkupperDeploymentCommon{
			Components: []v2.SkupperComponent{
				&v2.Router{
					// TODO ADD Labels
					Labels: map[string]string{},
					Env: map[string]string{
						"APPLICATION_NAME":    "skupper-router",
						"QDROUTERD_CONF":      "/etc/skupper-router/config/" + types.TransportConfigFile,
						"QDROUTERD_CONF_TYPE": "json",
						"SKUPPER_SITE_ID":     site.Id,
					},
					SiteIngresses: []v2.SiteIngress{
						&SiteIngressPodmanHost{
							SiteIngressCommon: &v2.SiteIngressCommon{
								Name: types.InterRouterIngressPrefix,
								Host: site.IngressBindHost,
								Port: site.IngressBindInterRouterPort,
								Target: &v2.PortCommon{
									Name: types.InterRouterIngressPrefix,
									Port: int(types.InterRouterListenerPort),
								},
							},
						},
						&SiteIngressPodmanHost{
							SiteIngressCommon: &v2.SiteIngressCommon{
								Name: types.EdgeIngressPrefix,
								Host: site.IngressBindHost,
								Port: site.IngressBindEdgePort,
								Target: &v2.PortCommon{
									Name: types.EdgeIngressPrefix,
									Port: int(types.EdgeListenerPort),
								},
							},
						},
					},
				},
			},
		},
		Aliases:      []string{types.TransportServiceName, types.LocalTransportServiceName},
		VolumeMounts: volumeMounts,
		Networks:     []string{site.ContainerNetwork},
	}
	site.Deployments = append(site.Deployments, routerDepl)
}

func (s *SitePodmanHandler) Create(site v2.Site) error {
	var err error
	var cleanupFns []func()

	podmanSite := site.(*SitePodman)
	if !podmanSite.prepared {
		var preparedSite v2.Site
		preparedSite, err = s.Prepare(podmanSite)
		if err != nil {
			return err
		}
		podmanSite = preparedSite.(*SitePodman)
	}

	// cleanup on error
	defer func() {
		if err != nil {
			for _, fn := range cleanupFns {
				fn()
			}
		}
	}()

	// Create network
	err = s.createNetwork(podmanSite)
	if err != nil {
		return err
	}
	cleanupFns = append(cleanupFns, func() {
		_ = s.cli.NetworkRemove(podmanSite.ContainerNetwork)
	})

	// Create cert authorities and credentials
	var credHandler types.CredentialHandler
	credHandler = NewPodmanCredentialHandler(s.cli)

	// - creating cert authorities
	cleanupFns = append(cleanupFns, func() {
		for _, ca := range podmanSite.GetCertAuthorities() {
			_ = credHandler.DeleteCertAuthority(ca.Name)
		}
	})
	for _, ca := range podmanSite.GetCertAuthorities() {
		if _, err = credHandler.NewCertAuthority(ca); err != nil {
			return err
		}
	}

	// - create credentials
	cleanupFns = append(cleanupFns, func() {
		for _, cred := range podmanSite.GetCredentials() {
			_ = credHandler.DeleteCredential(cred.Name)
		}
	})
	for _, cred := range podmanSite.GetCredentials() {
		if _, err = credHandler.NewCredential(cred); err != nil {
			return err
		}
	}

	// Create initial transport config file
	// TODO add log and debug options
	initialRouterConfig := v2.InitialConfigSkupperRouter(podmanSite.GetName(), podmanSite.GetId(), client.Version, podmanSite.IsEdge(), 3, types.RouterOptions{})
	var routerConfigHandler v2.RouterConfigHandler
	routerConfigHandler = NewRouterConfigHandlerPodman(s.cli)
	err = routerConfigHandler.SaveRouterConfig(&initialRouterConfig)
	cleanupFns = append(cleanupFns, func() {
		_ = routerConfigHandler.RemoveRouterConfig()
	})
	if err != nil {
		return err
	}

	// Verify volumes not yet created and create them
	for _, volumeName := range SkupperContainerVolumes {
		var vol *container.Volume
		vol, err = s.cli.VolumeInspect(volumeName)
		if vol == nil && err != nil {
			_, err = s.cli.VolumeCreate(&container.Volume{Name: volumeName})
			if err != nil {
				return err
			}
			cleanupFns = append(cleanupFns, func() {
				_ = s.cli.VolumeRemove(volumeName)
			})
		}
	}

	// Deploy container(s)
	deployHandler := NewSkupperDeploymentHandlerPodman(s.cli)
	for _, depl := range podmanSite.GetDeployments() {
		err = deployHandler.Deploy(depl)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *SitePodmanHandler) Get() (v2.Site, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *SitePodmanHandler) Delete(site v2.Site) error {
	return fmt.Errorf("not implemented")
}

func (s *SitePodmanHandler) Update(site v2.Site) error {
	return fmt.Errorf("not implemented")
}

func (s *SitePodmanHandler) canCreate(site *SitePodman) error {

	// Validating podman endpoint
	if s.cli == nil {
		cli, err := podman.NewPodmanClient(site.PodmanEndpoint, "")
		if err != nil {
			// TODO try to start podman's user service instance?
			return fmt.Errorf("unable to communicate with podman service through %s - %v", site.PodmanEndpoint, err)
		}
		s.cli = cli
	}

	// Validate podman version
	cli := s.cli
	version, err := cli.Version()
	if err != nil {
		return fmt.Errorf("error validating podman version - %v", err)
	}
	apiVersion := utils.ParseVersion(version.Server.APIVersion)
	if apiVersion.Major < 4 {
		return fmt.Errorf("podman version must be 4.0.0 or greater, found: %s", version.Server.APIVersion)
	}

	// TODO improve on container and network already exists
	// Validating any of the required deployment exists
	for _, skupperDepl := range site.Deployments {
		container, err := cli.ContainerInspect(skupperDepl.GetName())
		if err == nil && container != nil {
			return fmt.Errorf("%s container already defined", skupperDepl.GetName())
		}
	}

	// Validating skupper networks available
	net, err := cli.NetworkInspect(site.ContainerNetwork)
	if err == nil && net != nil {
		return fmt.Errorf("network %s already exists", site.ContainerNetwork)
	}

	// Validating bind ports
	for _, skupperDepl := range site.GetDeployments() {
		for _, skupperComp := range skupperDepl.GetComponents() {
			for _, ingress := range skupperComp.GetSiteIngresses() {
				if utils.TcpPortInUse(ingress.GetHost(), ingress.GetPort()) {
					return fmt.Errorf("ingress port already bound %s:%d", ingress.GetHost(), ingress.GetPort())
				}

			}
		}
	}

	// Validate network ability to resolve names
	createdNetwork, err := cli.NetworkCreate(&container.Network{
		Name:     site.ContainerNetwork,
		DNS:      true,
		Internal: false,
	})
	if err != nil {
		return fmt.Errorf("error validating network creation - %v", err)
	}
	defer func(cli *podman.PodmanRestClient, id string) {
		err := cli.NetworkRemove(id)
		if err != nil {
			fmt.Printf("ERROR removing network %s - %v\n", id, err)
		}
	}(cli, site.ContainerNetwork)
	if !createdNetwork.DNS {
		return fmt.Errorf("network %s cannot resolve names - podman plugins must be installed", site.ContainerNetwork)
	}

	// Validating existing volumes
	for _, v := range SkupperContainerVolumes {
		_, err := cli.VolumeInspect(v)
		if err == nil {
			return fmt.Errorf("required volume already exists %s", v)
		}
	}

	return nil
}

func (s *SitePodmanHandler) createNetwork(site *SitePodman) error {
	_, err := s.cli.NetworkCreate(&container.Network{
		Name:     site.ContainerNetwork,
		DNS:      true,
		Internal: false,
	})
	if err != nil {
		return fmt.Errorf("error creating network %s - %v", site.ContainerNetwork, err)
	}
	return nil
}
