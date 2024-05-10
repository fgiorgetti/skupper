package podman

import (
	"context"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/skupperproject/skupper/api/types"
	"github.com/skupperproject/skupper/client/container"
	"github.com/skupperproject/skupper/client/podman"
	"github.com/skupperproject/skupper/pkg/images"
	"github.com/skupperproject/skupper/pkg/non_kube/apis"
	"github.com/skupperproject/skupper/pkg/non_kube/common"
)

type SiteStateRenderer struct {
	loadedSiteState apis.SiteState
	siteState       apis.SiteState
	configRenderer  *common.FileSystemConfigurationRenderer
	containers      map[string]container.Container
	cli             *podman.PodmanRestClient
}

func (s *SiteStateRenderer) Render(loadedSiteState apis.SiteState) error {
	var err error
	var validator apis.SiteStateValidator = &common.SiteStateValidator{}
	// TODO enhance site state validator (too basic yet)
	err = validator.Validate(loadedSiteState)
	if err != nil {
		return err
	}
	s.loadedSiteState = loadedSiteState
	// TODO define how to get podman socket endpoint from Site CR
	s.cli, err = podman.NewPodmanClient("", "")
	if err != nil {
		return fmt.Errorf("failed to create podman client: %v", err)
	}
	// active (runtime) SiteState
	s.siteState = common.CopySiteState(s.loadedSiteState)
	err = common.PrepareCertificatesAndLinkAccess(&s.siteState)
	if err != nil {
		return fmt.Errorf("failed to prepare podman site: %w", err)
	}
	// rendering non-kube configuration files and certificates
	s.configRenderer = &common.FileSystemConfigurationRenderer{
		Force: true, // TODO discuss how this should be handled?
	}
	err = s.configRenderer.Render(s.siteState)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Serializing loaded and runtime site states
	if err = apis.MarshalSiteState(s.loadedSiteState, path.Join(s.configRenderer.OutputPath, common.LoadedSiteStatePath)); err != nil {
		return err
	}
	if err = apis.MarshalSiteState(s.siteState, path.Join(s.configRenderer.OutputPath, common.RuntimeSiteStatePath)); err != nil {
		return err
	}

	// TODO Render containers
	// TODO Controller might need some more thinking still
	//      - collector is being separated
	//      - claim api needs to be added
	//      - an API for interacting with CRs (OpenAPI/Rest)

	// TODO How to get timeout setting from Site CR
	ctx, cn := context.WithTimeout(context.Background(), time.Minute*10)
	defer cn()
	if err = s.prepareContainers(); err != nil {
		return err
	}
	if err = s.pullImages(ctx); err != nil {
		return err
	}
	if err = s.createContainers(); err != nil {
		return err
	}
	if err = s.startContainers(); err != nil {
		return err
	}
	return err
}

func (s *SiteStateRenderer) prepareContainers() error {
	siteConfigPath, err := common.GetHostSiteHome(s.siteState.Site)
	if err != nil {
		return err
	}
	s.containers = make(map[string]container.Container)
	s.containers[types.RouterComponent] = container.Container{
		Name:  fmt.Sprintf("%s-skupper-router", s.siteState.Site.Name),
		Image: images.GetRouterImageName(),
		Env: map[string]string{
			"APPLICATION_NAME":      "skupper-router",
			"QDROUTERD_CONF":        "/etc/skupper-router/config/" + types.TransportConfigFile,
			"QDROUTERD_CONF_TYPE":   "json",
			"SKUPPER_SITE_ID":       s.configRenderer.RouterConfig.GetSiteMetadata().Id,
			"SSL_PROFILE_BASE_PATH": "/etc/skupper-router",
		},
		Labels: map[string]string{
			types.ComponentAnnotation: types.TransportDeploymentName,
			types.SiteId:              s.configRenderer.RouterConfig.GetSiteMetadata().Id,
		},
		FileMounts: []container.FileMount{
			{
				Source:      path.Join(siteConfigPath, "config/router"),
				Destination: "/etc/skupper-router/config",
				Options:     []string{"z"},
			},
			{
				Source:      path.Join(siteConfigPath, "certificates"),
				Destination: "/etc/skupper-router/certificates",
				Options:     []string{"z"},
			},
		},
		RestartPolicy: "always",
		// TODO handle resource utilization with podman sites
	}
	return nil
}

func (s *SiteStateRenderer) pullImages(ctx context.Context) error {
	var err error
	for component, skupperContainer := range s.containers {
		err = s.cli.ImagePull(ctx, skupperContainer.Image)
		if err != nil {
			return fmt.Errorf("failed to pull %s image %s: %w", component, skupperContainer.Image, err)
		}
	}
	return nil
}

func (s *SiteStateRenderer) cleanupContainers(err error) {
	if err == nil {
		return
	}
	for _, createdContainer := range s.containers {
		_ = s.cli.ContainerStop(createdContainer.Name)
		_ = s.cli.ContainerRemove(createdContainer.Name)
	}
}

func (s *SiteStateRenderer) createContainers() error {
	var err error
	defer s.cleanupContainers(err)
	// validate if containers already exist before creating anything
	for component, skupperContainer := range s.containers {
		existingContainer, err := s.cli.ContainerInspect(skupperContainer.Name)
		if err == nil && existingContainer != nil {
			return fmt.Errorf("container %s already exists (component: %s)", skupperContainer.Name, component)
		}
	}
	for component, skupperContainer := range s.containers {
		err = s.cli.ContainerCreate(&skupperContainer)
		if err != nil {
			return fmt.Errorf("failed to create %q container (%s): %w", component, skupperContainer.Name, err)
		}
	}
	return nil
}

func (s *SiteStateRenderer) startContainers() error {
	var err error
	defer s.cleanupContainers(err)
	for component, skupperContainer := range s.containers {
		err = s.cli.ContainerStart(skupperContainer.Name)
		if err != nil {
			return fmt.Errorf("failed to start %s container %q: %w", component, skupperContainer.Name, err)
		}
	}
	return nil
}
