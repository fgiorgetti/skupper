package skupper_podman

import (
	"fmt"
	"strconv"

	"github.com/skupperproject/skupper/api/types"
	v2 "github.com/skupperproject/skupper/api/types/v2"
	"github.com/skupperproject/skupper/client/container"
	"github.com/skupperproject/skupper/client/podman"
)

type SkupperDeploymentPodman struct {
	*v2.SkupperDeploymentCommon
	Name         string
	Aliases      []string
	VolumeMounts map[string]string
	Networks     []string
}

func (s *SkupperDeploymentPodman) GetName() string {
	return s.Name
}

type SkupperDeploymentHandlerPodman struct {
	cli *podman.PodmanRestClient
}

func NewSkupperDeploymentHandlerPodman(cli *podman.PodmanRestClient) *SkupperDeploymentHandlerPodman {
	return &SkupperDeploymentHandlerPodman{
		cli: cli,
	}
}

func (s *SkupperDeploymentHandlerPodman) Deploy(deployment v2.SkupperDeployment) error {
	var err error
	var cleanupContainers []string

	defer func() {
		if err != nil {
			for _, containerName := range cleanupContainers {
				_ = s.cli.ContainerStop(containerName)
				_ = s.cli.ContainerRemove(containerName)
			}
		}
	}()

	podmanDeployment := deployment.(*SkupperDeploymentPodman)
	for _, component := range deployment.GetComponents() {

		// Pulling image first
		err = s.cli.ImagePull(component.Image())
		if err != nil {
			return err
		}

		// Setting network aliases
		networkMap := map[string]container.ContainerNetworkInfo{}
		for _, network := range podmanDeployment.Networks {
			networkMap[network] = container.ContainerNetworkInfo{
				Aliases: podmanDeployment.Aliases,
			}
		}

		// Defining the mounted volumes
		mounts := []container.Volume{}
		for volumeName, destDir := range podmanDeployment.VolumeMounts {
			var volume *container.Volume
			volume, err = s.cli.VolumeInspect(volumeName)
			if err != nil {
				err = fmt.Errorf("error reading volume %s - %v", volumeName, err)
				return err
			}
			volume.Destination = destDir
			mounts = append(mounts, *volume)
		}

		// Ports
		ports := []container.Port{}
		for _, siteIngress := range component.GetSiteIngresses() {
			ports = append(ports, container.Port{
				Host:     strconv.Itoa(siteIngress.GetPort()),
				HostIP:   siteIngress.GetHost(),
				Target:   strconv.Itoa(siteIngress.GetTarget().GetPort()),
				Protocol: "tcp",
			})
		}

		// Defining the container
		labels := component.GetLabels()
		labels["application"] = types.TransportDeploymentName
		c := &container.Container{
			Name:          component.Name(),
			Image:         component.Image(),
			Env:           component.GetEnv(),
			Labels:        labels,
			Networks:      networkMap,
			Mounts:        mounts,
			Ports:         ports,
			RestartPolicy: "always",
		}

		err = s.cli.ContainerCreate(c)
		if err != nil {
			return fmt.Errorf("error creating skupper component: %s - %v", c.Name, err)
		}
		cleanupContainers = append(cleanupContainers, c.Name)

		err = s.cli.ContainerStart(c.Name)
		if err != nil {
			return fmt.Errorf("error starting skupper component: %s - %v", c.Name, err)
		}
	}

	return nil
}

func (s *SkupperDeploymentHandlerPodman) Undeploy(name string) error {
	containers, err := s.cli.ContainerList()
	if err != nil {
		return fmt.Errorf("error listing containers - %w", err)
	}

	stopContainers := []string{}
	for _, c := range containers {
		if appName, ok := c.Labels["application"]; ok && appName == name {
			stopContainers = append(stopContainers, c.Name)
		}
	}

	if len(stopContainers) == 0 {
		return nil
	}

	for _, c := range stopContainers {
		if err = s.cli.ContainerStop(c); err != nil {
			return err
		}
		if err = s.cli.ContainerRemove(c); err != nil {
			return err
		}
	}
	return nil
}
