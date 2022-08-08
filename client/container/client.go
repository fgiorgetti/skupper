package container

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/skupperproject/skupper/client/generated/libpod/models"
)

type Client interface {
	Version() (*Version, error)
	ContainerList() ([]Container, error)
	ContainerInspect(id string) (*Container, error)
	ContainerCreate(container *Container) error
	ContainerRemove(id string) error
	ContainerExec(id string, command []string) (string, string, error)
	ContainerLogs(id string) (string, error)
	ContainerStart(id string) error
	ContainerStop(id string) error
	ContainerRestart(id string) error
	ImageList()
	ImageInspect()
	ImagePull()
	NetworkList()
	NetworkInspect()
	NetworkCreate()
	NetworkRemove()
	NetworkConnect()
	NetworkDisconnect()
}

type VersionInfo struct {
	Version    string
	APIVersion string
}

type Version struct {
	Client VersionInfo
	Server VersionInfo
}

type Container struct {
	ID            string
	Name          string
	Pod           string
	Image         string
	Env           map[string]string
	Labels        map[string]string
	Annotations   map[string]string
	Networks      map[string]NetworkInfo
	Mounts        []Volume
	Ports         []Port
	EntryPoint    []string
	Command       []string
	RestartPolicy string
	RestartCount  int
	Running       bool
	CreatedAt     string
	StartedAt     string
	ExitedAt      string
	ExitCode      int
}

func (c *Container) FromEnv(env []string) {
	for _, e := range env {
		if !strings.Contains(e, "=") {
			continue
		}
		es := strings.SplitN(e, "=", 2)
		c.Env[es[0]] = es[1]
	}
}

func (c *Container) EnvSlice() []string {
	es := []string{}
	for k, v := range c.Env {
		es = append(es, fmt.Sprintf("%s=%s", k, v))
	}
	return es
}

func (c *Container) NetworkNames() []string {
	var networks []string
	for name, _ := range c.Networks {
		networks = append(networks, name)
	}
	return networks
}

func (c *Container) VolumesToMounts() []*models.Mount {
	var mounts []*models.Mount
	for _, v := range c.Mounts {
		m := &models.Mount{
			ReadOnly:    !v.RW,
			Source:      v.Source,
			Target:      v.Destination,
			Destination: v.Destination,
			Type:        "bind",
			Options:     []string{"Z"},
		}
		mounts = append(mounts, m)
	}
	return mounts
}

func (c *Container) ToPortmappings() []*models.PortMapping {
	var mapping []*models.PortMapping
	for _, port := range c.Ports {
		target, _ := strconv.Atoi(port.Target)
		host, _ := strconv.Atoi(port.Host)

		mapping = append(mapping, &models.PortMapping{
			ContainerPort: uint16(target),
			HostIP:        port.HostIP,
			HostPort:      uint16(host),
			Protocol:      port.Protocol,
		})
	}
	return mapping
}

func (c *Container) ToSpecGenerator() *models.SpecGenerator {
	spec := &models.SpecGenerator{
		Annotations:   c.Annotations,
		CNINetworks:   c.NetworkNames(),
		Command:       c.Command,
		Entrypoint:    c.EntryPoint,
		Env:           c.Env,
		Image:         c.Image,
		Labels:        c.Labels,
		Mounts:        c.VolumesToMounts(),
		Name:          c.Name,
		Pod:           c.Pod,
		PortMappings:  c.ToPortmappings(),
		RestartPolicy: c.RestartPolicy,
	}

	// Network info
	spec.Aliases = map[string][]string{}
	for networkName, network := range c.Networks {
		spec.Aliases[networkName] = network.Aliases
	}
	return spec
}

type Volume struct {
	Source      string
	Destination string
	Mode        string
	RW          bool
}

type Port struct {
	Host     string
	HostIP   string
	Target   string
	Protocol string
}

type NetworkInfo struct {
	ID          string
	IPAddress   string
	IPPrefixLen int
	MacAddress  string
	Gateway     string
	Aliases     []string
}
