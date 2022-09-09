package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/skupperproject/skupper/api/types"
	"github.com/skupperproject/skupper/client"
	"github.com/skupperproject/skupper/client/container"
	"github.com/skupperproject/skupper/client/generated/libpod/client/volumes"
	"github.com/skupperproject/skupper/client/podman"
	"github.com/skupperproject/skupper/pkg/certs"
	"github.com/skupperproject/skupper/pkg/kube"
	"github.com/skupperproject/skupper/pkg/qdr"
	"github.com/skupperproject/skupper/pkg/utils"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PodmanSiteInitializer struct {
	spec   types.SiteConfigSpec
	params PodmanInitParams
	cli    *podman.PodmanRestClient
}

type PodmanCredentialHandler struct {
	cli *podman.PodmanRestClient
}

func NewPodmanCredentialHandler(cli *podman.PodmanRestClient) *PodmanCredentialHandler {
	return &PodmanCredentialHandler{
		cli: cli,
	}
}

func (p *PodmanCredentialHandler) LoadVolumeAsSecret(vol *container.Volume) (*corev1.Secret, error) {
	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      vol.Name,
			Namespace: Namespace,
		},
		Data: map[string][]byte{},
		Type: "kubernetes.io/tls",
	}

	if metadataLabel, ok := vol.Labels[types.InternalMetadataQualifier]; ok {
		err := json.Unmarshal([]byte(metadataLabel), &secret.ObjectMeta)
		if err != nil {
			return nil, fmt.Errorf("error loading secret metadata from volume - %v", err)
		}
	}

	files, err := vol.ListFiles()
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		data, err := ioutil.ReadFile(file.Name())
		if err != nil {
			return nil, fmt.Errorf("error reading file %s for secret %s - %v", file.Name(), vol.Name, err)
		}
		secret.Data[file.Name()] = data
	}

	return secret, nil
}

func (p *PodmanCredentialHandler) SaveSecretAsVolume(secret *corev1.Secret) (*container.Volume, error) {
	vol, err := p.cli.VolumeInspect(secret.Name)

	if err != nil {
		if _, notFound := err.(*volumes.VolumeInspectLibpodNotFound); !notFound {
			return nil, err
		}
		// creating new volume
		metadataStr, err := json.Marshal(secret.ObjectMeta)
		if err != nil {
			return nil, fmt.Errorf("error marshalling secret info for %s - %v", secret.Name, err)
		}
		vol = &container.Volume{
			Name: secret.Name,
			Labels: map[string]string{
				types.InternalMetadataQualifier: string(metadataStr),
			},
		}
		vol, err = p.cli.VolumeCreate(vol)
		if err != nil {
			return nil, fmt.Errorf("error creating volume %s - %v", secret.Name, err)
		}
	}
	_, err = vol.CreateDataFiles(secret.Data, true)
	return nil, err
}

func (p *PodmanCredentialHandler) NewCertAuthority(ca types.CertAuthority) (*corev1.Secret, error) {
	caSecret, err := p.GetSecret(ca.Name)
	if err == nil {
		return caSecret, nil
	}
	if _, notFound := err.(*volumes.VolumeInspectLibpodNotFound); !notFound {
		return nil, fmt.Errorf("Failed to check CA %s : %w", ca.Name, err)
	}
	newCA := certs.GenerateCASecret(ca.Name, ca.Name)
	_, err = p.SaveSecretAsVolume(&newCA)
	return &newCA, err
}

func (p *PodmanCredentialHandler) NewCredential(cred types.Credential) (*corev1.Secret, error) {
	var caSecret *corev1.Secret
	var err error
	if cred.CA != "" {
		caSecret, err = p.GetSecret(cred.CA)
		if err != nil {
			return nil, fmt.Errorf("error loading CA secret %s - %v", cred.CA, err)
		}
	}
	secret := kube.PrepareNewSecret(cred, caSecret, types.TransportDeploymentName)
	_, err = p.SaveSecretAsVolume(&secret)
	return &secret, err
}

func (p *PodmanCredentialHandler) GetSecret(name string) (*corev1.Secret, error) {
	vol, err := p.cli.VolumeInspect(name)
	if err != nil {
		return nil, err
	}
	return p.LoadVolumeAsSecret(vol)
}

type SkupperSitePodman struct {
	Containers    []container.Container
	VolumeSecrets map[string][]string
}

func NewSkupperSitePodman(skupper *SkupperSite) (*SkupperSitePodman, error) {

	return nil, nil
}

func NewPodmanSiteInitializer(spec types.SiteConfigSpec, params PodmanInitParams) (*PodmanSiteInitializer, error) {
	si := &PodmanSiteInitializer{
		spec:   spec,
		params: params,
	}
	return si, nil
}

func (p *PodmanSiteInitializer) Prepare() (*SkupperSite, error) {
	site := &SkupperSite{
		Config: &types.SiteConfig{
			Spec: p.spec,
		},
		Spec: &types.RouterSpec{},
	}

	// Validating parameters and environment
	err := p.validate()
	if err != nil {
		return nil, err
	}

	// Preparing the router deployment
	siteId := client.GetSiteId(*site.Config)

	// Preparing the router spec
	site.Spec = client.GetRouterSpecFromOpts(site.Config.Spec, siteId, Namespace)
	// Renaming router-config volume as skupper-internal (so that the podman volume name matches)
	for idx, volume := range site.Spec.Transport.Volumes {
		if volume.Name == "router-config" {
			site.Spec.Transport.Volumes[idx].Name = types.TransportConfigMapName
		}
	}
	for idx, vm := range site.Spec.Transport.VolumeMounts[0] {
		if vm.Name == "router-config" {
			site.Spec.Transport.VolumeMounts[0][idx].Name = types.TransportConfigMapName
		}
	}

	return site, nil
}

func (p *PodmanSiteInitializer) Initialize(site *SkupperSite) (*SkupperSite, error) {
	var err error
	var cleanupFns []func()

	// cleanup on error
	defer func() {
		if err != nil {
			for _, fn := range cleanupFns {
				fn()
			}
		}
	}()

	// Create network: skupper
	_, err = p.cli.NetworkCreate(&container.Network{
		Name:     ContainerNetworkName,
		DNS:      true,
		Internal: false,
	})
	if err != nil {
		return nil, fmt.Errorf("error creating network %s - %v", ContainerNetworkName, err)
	}
	cleanupFns = append(cleanupFns, func() {
		_ = p.cli.NetworkRemove(ContainerNetworkName)
	})

	// Create volumes: [skupper-local-server, router-config, skupper-site-server, skupper-router-certs]
	for _, volume := range site.Spec.Transport.Volumes {
		var vc *container.Volume
		vc, err = p.cli.VolumeCreate(&container.Volume{Name: volume.Name})
		if err != nil {
			return nil, fmt.Errorf("error creating volume %s - %v", volume, err)
		}
		cleanupFns = append(cleanupFns, func() {
			_ = p.cli.VolumeRemove(vc.Name)
		})
	}

	// Save transport config file
	var transportConfig *container.Volume
	transportConfig, err = p.cli.VolumeInspect(types.TransportConfigMapName)
	if err != nil {
		return nil, fmt.Errorf("error creating %s volume - %v", types.TransportConfigMapName, err)
	}
	cleanupFns = append(cleanupFns, func() {
		_ = p.cli.VolumeRemove(transportConfig.Name)
	})
	transportConfigMap := qdr.AsConfigMapData(site.Spec.RouterConfig)
	_, err = transportConfig.CreateFiles(transportConfigMap, false)
	if err != nil {
		return nil, err
	}

	// Credentials
	handler := NewPodmanCredentialHandler(p.cli)

	// Create cert authorities
	for _, ca := range site.Spec.CertAuthoritys {
		_, err = handler.NewCertAuthority(ca)
		if err != nil {
			return nil, fmt.Errorf("error creating certificate authority %s - %v", ca.Name, err)
		}
		cleanupFns = append(cleanupFns, func() {
			_ = p.cli.VolumeRemove(ca.Name)
		})
	}

	// Create credentials
	for _, cred := range site.Spec.TransportCredentials {
		handler.NewCredential(cred)
	}

	// Create container: router (current container name for skupper-router deployment)
	/*
		mount volumes:
		- skupper-local-server -> /etc/skupper-router-certs/skupper-amqps/
		- router-config -> /etc/skupper-router/config/
		- skupper-site-server -> /etc/skupper-router-certs/skupper-internal/
		- skupper-router-certs -> /etc/skupper-router-certs
	*/

	// Create startup script

	// Create systemd user service

	return site, err
}

func (p *PodmanSiteInitializer) PostInitialize(site *SkupperSite) (*SkupperSite, error) {
	panic("implement me")
}

func (p *PodmanSiteInitializer) validate() error {
	// Validating podman endpoint
	cli, err := podman.NewPodmanClient(p.params.PodmanEndpoint, "")
	if err != nil {
		// TODO try to start podman's user service instance?
		return fmt.Errorf("unable to communicate with podman service through %s - %v", p.params.PodmanEndpoint, err)
	}
	p.cli = cli

	// Validate podman version
	version, err := cli.Version()
	if err != nil {
		return fmt.Errorf("error validating podman version - %v", err)
	}
	apiVersion := utils.ParseVersion(version.Server.APIVersion)
	if apiVersion.Major < 4 {
		return fmt.Errorf("podman version must be 4.0.0 or greater, found: %s", version.Server.APIVersion)
	}

	// TODO improve on container and network already exists
	// Validating router container already defined
	routerContainer, err := cli.ContainerInspect(types.TransportDeploymentName)
	if err == nil && routerContainer != nil {
		return fmt.Errorf("%s container already defined", types.TransportDeploymentName)
	}

	// Validating skupper network available
	net, err := cli.NetworkInspect(ContainerNetworkName)
	if err == nil && net != nil {
		return fmt.Errorf("network %s already exists", ContainerNetworkName)
	}

	// Validating bind ports
	for _, port := range []int{p.params.IngressBindInterRouterPort, p.params.IngressBindEdgePort} {
		if TcpPortInUse(p.params.IngressBindHost, port) {
			return fmt.Errorf("ingress port already bound %s:%d", p.params.IngressBindHost, port)
		}
	}

	// Validate network ability to resolve names
	createdNetwork, err := cli.NetworkCreate(&container.Network{
		Name:     ContainerNetworkName,
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
	}(cli, ContainerNetworkName)
	if !createdNetwork.DNS {
		return fmt.Errorf("network %s cannot resolve names - podman plugins must be installed", ContainerNetworkName)
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

type PodmanInitParams struct {
	PodmanEndpoint             string
	IngressBindHost            string
	IngressBindInterRouterPort int
	IngressBindEdgePort        int
	ContainerNetworks          []string
}
