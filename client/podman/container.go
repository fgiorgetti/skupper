package podman

import (
	"fmt"
	"strings"

	"github.com/skupperproject/skupper/client/container"
	"github.com/skupperproject/skupper/client/generated/libpod/client/containers"
	"github.com/skupperproject/skupper/client/generated/libpod/models"
)

/*
	TODO
	ContainerExec(id string, command []string) (string, string, error)
	ContainerLogs(id string) (string, error)
*/

func (p *PodmanRestClient) ContainerList() ([]*container.Container, error) {
	cli := containers.New(p.RestClient, formats)
	params := containers.NewContainerListLibpodParams()
	params.All = boolTrue()
	list, err := cli.ContainerListLibpod(params)
	if err != nil {
		return nil, fmt.Errorf("error listing containers: %v", err)
	}
	cts := []*container.Container{}
	for _, c := range list.Payload {
		if c == nil {
			continue
		}
		cts = append(cts, FromListContainer(*c))
	}
	return cts, nil
}

func (p *PodmanRestClient) ContainerInspect(id string) (*container.Container, error) {
	cli := containers.New(p.RestClient, formats)
	param := containers.NewContainerInspectLibpodParams()
	param.Name = id
	res, err := cli.ContainerInspectLibpod(param)
	if err != nil {
		return nil, fmt.Errorf("error inspecting container '%s': %v", id, err)
	}
	return FromInspectContainer(*res.Payload), nil
}

func (p *PodmanRestClient) ContainerCreate(container *container.Container) error {
	cli := containers.New(p.RestClient, formats)
	params := containers.NewContainerCreateLibpodParams()
	params.Create = container.ToSpecGenerator()
	_, err := cli.ContainerCreateLibpod(params)
	if err != nil {
		return fmt.Errorf("error creating container %s: %v", container.Name, err)
	}
	return nil
}

func (p *PodmanRestClient) ContainerRemove(name string) error {
	cli := containers.New(p.RestClient, formats)
	params := containers.NewContainerDeleteLibpodParams()
	params.Name = name
	params.Force = boolTrue()
	_, _, err := cli.ContainerDeleteLibpod(params)
	if err != nil {
		return fmt.Errorf("error deleting container %s: %v", name, err)
	}
	return nil
}

func (p *PodmanRestClient) ContainerStart(name string) error {
	cli := containers.New(p.RestClient, formats)
	params := containers.NewContainerStartLibpodParams()
	params.Name = name
	_, err := cli.ContainerStartLibpod(params)
	if err != nil {
		return fmt.Errorf("error starting container %s: %v", name, err)
	}
	return nil
}

func (p *PodmanRestClient) ContainerStop(name string) error {
	cli := containers.New(p.RestClient, formats)
	params := containers.NewContainerStopLibpodParams()
	params.Name = name
	_, err := cli.ContainerStopLibpod(params)
	if err != nil {
		return fmt.Errorf("error stopping container %s: %v", name, err)
	}
	return nil
}

func (p *PodmanRestClient) ContainerRestart(name string) error {
	cli := containers.New(p.RestClient, formats)
	params := containers.NewContainerRestartLibpodParams()
	params.Name = name
	_, err := cli.ContainerRestartLibpod(params)
	if err != nil {
		return fmt.Errorf("error restarting container %s: %v", name, err)
	}
	return nil
}

func FromListContainer(c models.ListContainer) *container.Container {
	ct := &container.Container{
		ID:        c.ID,
		Name:      c.Names[0],
		Pod:       c.Pod,
		Image:     c.Image,
		Labels:    c.Labels,
		Command:   c.Command,
		Running:   !c.Exited,
		CreatedAt: fmt.Sprint(c.CreatedAt),
		StartedAt: fmt.Sprint(c.CreatedAt),
		ExitedAt:  fmt.Sprint(c.ExitedAt),
		ExitCode:  int(c.ExitCode),
	}
	ct.Networks = map[string]container.NetworkInfo{}
	ct.Env = map[string]string{}

	// base network info
	for _, n := range c.Networks {
		network := container.NetworkInfo{ID: n}
		ct.Networks[n] = network
	}
	// base mount info
	for _, m := range c.Mounts {
		v := container.Volume{Destination: m}
		ct.Mounts = append(ct.Mounts, v)
	}
	// port mapping
	for _, port := range c.Ports {
		ct.Ports = append(ct.Ports, container.Port{
			Host:     fmt.Sprint(port.HostPort),
			HostIP:   port.HostIP,
			Target:   fmt.Sprint(port.ContainerPort),
			Protocol: port.Protocol,
		})
	}
	return ct
}

func FromInspectContainer(c containers.ContainerInspectLibpodOKBody) *container.Container {
	ct := &container.Container{
		ID:           c.ID,
		Name:         c.Name,
		RestartCount: int(c.RestartCount),
		CreatedAt:    c.Created.String(),
		Pod:          c.Pod,
	}
	ct.Networks = map[string]container.NetworkInfo{}
	ct.Labels = map[string]string{}
	ct.Annotations = map[string]string{}
	ct.Env = map[string]string{}

	// Volume mounts
	for _, m := range c.Mounts {
		volume := container.Volume{
			Source:      m.Source,
			Destination: m.Destination,
			Mode:        m.Mode,
			RW:          m.RW,
		}
		ct.Mounts = append(ct.Mounts, volume)
	}

	// Container config
	if c.Config != nil {
		config := c.Config
		ct.Image = config.Image
		ct.FromEnv(config.Env)
		ct.Labels = config.Labels
		ct.Annotations = config.Annotations
		if config.Entrypoint != "" {
			ct.EntryPoint = []string{config.Entrypoint}
		}
		ct.Command = config.Cmd
	}

	// HostConfig
	if c.HostConfig != nil {
		hostConfig := c.HostConfig
		if hostConfig.RestartPolicy != nil {
			ct.RestartPolicy = hostConfig.RestartPolicy.Name
		}
	}

	// Network info
	if c.NetworkSettings != nil {
		// Addressing info
		for k, v := range c.NetworkSettings.Networks {
			netInfo := container.NetworkInfo{
				ID:          v.NetworkID,
				IPAddress:   v.IPAddress,
				IPPrefixLen: int(v.IPPrefixLen),
				MacAddress:  v.MacAddress,
				Gateway:     v.Gateway,
				Aliases:     v.Aliases,
			}
			ct.Networks[k] = netInfo
		}

		// Port mapping
		for portProto, ports := range c.NetworkSettings.Ports {
			portProtoS := strings.Split(portProto, "/")
			protocol := "tcp"
			if len(portProtoS) > 1 {
				protocol = portProtoS[1]
			}
			targetPort := portProtoS[0]
			for _, portInfo := range ports {
				p := container.Port{
					Host:     portInfo.HostPort,
					HostIP:   portInfo.HostIP,
					Target:   targetPort,
					Protocol: protocol,
				}
				ct.Ports = append(ct.Ports, p)
			}
		}
	}

	// State info
	if c.State != nil {
		ct.Running = c.State.Running
		ct.StartedAt = c.State.StartedAt.String()
		ct.ExitedAt = c.State.FinishedAt.String()
		ct.ExitCode = int(c.State.ExitCode)
	}

	return ct
}
