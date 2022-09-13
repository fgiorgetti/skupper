package skupper_podman

import (
	"fmt"

	"github.com/skupperproject/skupper/api/types"
	v2 "github.com/skupperproject/skupper/api/types/v2"
	"github.com/skupperproject/skupper/client/container"
	"github.com/skupperproject/skupper/client/podman"
	"github.com/skupperproject/skupper/pkg/utils"
)

type SitePodman struct {
	*v2.SiteCommon
	IngressBindHost            string
	IngressBindInterRouterPort int
	IngressBindEdgePort        int
	ContainerNetworks          []string
	PodmanEndpoint             string
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
	s.ConfigurePodmanDeployments(podmanSite)

	if err := s.canCreate(podmanSite); err != nil {
		return nil, err
	}
	return podmanSite, nil
}

func (s *SitePodmanHandler) ConfigurePodmanDeployments(site *SitePodman) {
	// CAs
	cas := []types.CertAuthority{}
	cas = append(cas, types.CertAuthority{Name: types.LocalCaSecret})
	if !site.IsEdge() {
		cas = append(cas, types.CertAuthority{Name: types.SiteCaSecret})
	}
	cas = append(cas, types.CertAuthority{Name: types.ServiceCaSecret})
	site.CertAuthorities = cas

	// Certificates
	credentials := []types.Credential{}
	credentials = append(credentials, types.Credential{
		CA:          types.LocalCaSecret,
		Name:        types.LocalServerSecret,
		Subject:     types.LocalTransportServiceName,
		Hosts:       []string{types.LocalTransportServiceName},
		ConnectJson: false,
		Post:        false,
	})
	credentials = append(credentials, types.Credential{
		CA:          types.LocalCaSecret,
		Name:        types.LocalClientSecret,
		Subject:     types.LocalTransportServiceName,
		Hosts:       []string{},
		ConnectJson: true,
		Post:        false,
	})

	credentials = append(credentials, types.Credential{
		CA:          types.ServiceCaSecret,
		Name:        types.ServiceClientSecret,
		Hosts:       []string{},
		ConnectJson: false,
		Post:        false,
		Simple:      true,
	})

	if !site.IsEdge() {
		credentials = append(credentials, types.Credential{
			CA:          types.SiteCaSecret,
			Name:        types.SiteServerSecret,
			Subject:     types.TransportServiceName,
			Hosts:       []string{types.TransportServiceName, site.IngressBindHost},
			ConnectJson: false,
		})
	}
	site.Credentials = credentials

	// Router Deployment
	volumeMounts := map[string]string{
		types.LocalServerSecret:      "/etc/skupper-router-certs/skupper-amqps/",
		types.TransportConfigMapName: "/etc/skupper-router/config/",
		"skupper-router-certs":       "/etc/skupper-router-certs",
	}
	if !site.IsEdge() {
		volumeMounts[types.SiteServerSecret] = "/etc/skupper-router-certs/skupper-internal/"
	}
	routerDepl := &SkupperDeploymentRouterPodman{
		SkupperDeploymentRouter: &v2.SkupperDeploymentRouter{
			Components: []v2.SkupperComponent{
				&v2.Router{
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
		SkupperDeploymentPodman: &SkupperDeploymentPodman{
			Aliases:      []string{types.TransportServiceName, types.LocalTransportServiceName},
			VolumeMounts: volumeMounts,
		},
	}
	site.Deployments = append(site.Deployments, routerDepl)
}

func (s *SitePodmanHandler) Create(site v2.Site) error {
	return fmt.Errorf("not implemented")
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
		container, err := cli.ContainerInspect(skupperDepl.Name())
		if err == nil && container != nil {
			return fmt.Errorf("%s container already defined", skupperDepl.Name())
		}
	}

	// Validating skupper networks available
	for _, networkName := range site.ContainerNetworks {
		net, err := cli.NetworkInspect(networkName)
		if err == nil && net != nil {
			return fmt.Errorf("network %s already exists", networkName)
		}
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
	testNetwork := site.ContainerNetworks[0]
	createdNetwork, err := cli.NetworkCreate(&container.Network{
		Name:     testNetwork,
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
	}(cli, testNetwork)
	if !createdNetwork.DNS {
		return fmt.Errorf("network %s cannot resolve names - podman plugins must be installed", testNetwork)
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
