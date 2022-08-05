package container

import (
	"fmt"
	"strings"
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
	ID           string
	Name         string
	Image        string
	Env          map[string]string
	Labels       map[string]string
	Networks     map[string]NetworkInfo
	Mounts       []Volume
	Ports        []Port
	EntryPoint   []string
	Command      []string
	RestartCount int
	Running      bool
	CreatedAt    string
	StartedAt    string
	ExitedAt     string
	ExitCode     int
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
