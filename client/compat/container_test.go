package compat

import (
	"testing"

	"github.com/skupperproject/skupper/client/container"
	"gotest.tools/assert"
)

func TestContainerCreate(t *testing.T) {
	//ccli, err := NewCompatClient("", "")
	ccli, err := NewCompatClient("/tmp/docker.sock", "")
	assert.Assert(t, err)

	c := &container.Container{
		Name:  "my-container",
		Image: "quay.io/skupper/skupper-router:main",
		Env: map[string]string{
			"env1": "val1",
		},
		Annotations: map[string]string{
			"io.podman.annotations.label": "disable",
		},
		Labels: map[string]string{
			"foo": "bar",
		},
		Networks: map[string]container.ContainerNetworkInfo{
			"my-network": {
				Aliases: []string{"c1", "c2"},
			},
		},
		Mounts: []container.Volume{
			{
				Name:        "vol1",
				Destination: "/opt/vol1",
				Mode:        "z",
			},
		},
		FileMounts: []container.FileMount{
			{
				Source:      "/run/user/1000/podman/podman.sock",
				Destination: "/tmp/podman.sock",
				Options:     []string{"z"},
			},
		},
		Ports: []container.Port{
			{
				Host:     "8888",
				HostIP:   "192.168.124.1",
				Target:   "8080",
				Protocol: "tcp",
			},
		},
		EntryPoint:     []string{"tail", "-f", "/dev/null"},
		Command:        []string{"tail", "-f", "/dev/null"},
		RestartPolicy:  "always",
		MaxCpus:        1,
		MaxMemoryBytes: 4096 * 1024 * 1024,
	}
	_ = ccli.ContainerRemove(c.Name)
	err = ccli.ContainerCreate(c)
	assert.Assert(t, err)
	//cli := containers_compat.New(ccli.RestClient, formats)
	//params := containers_compat.NewContainerCreateParams()
	//
	//params.Name = stringP("my-container")
	//params.Body = &models.CreateContainerConfig{
	//	Cmd: []string{"tail", "-f", "/dev/null"},
	//	Env: []string{"PODMAN_ENDPOINT=/tmp/podman.sock"},
	//	ExposedPorts: models.PortSet{
	//		"8080/tcp": interface{}(nil),
	//	},
	//	//Healthcheck:      nil,
	//	HostConfig: &models.HostConfig{
	//		Binds:    nil,
	//		CPUCount: 1,
	//		//CPUPercent:           0,
	//		CPUPeriod: 100000,
	//		CPUQuota:  int64(1 * 100000),
	//		//CPURealtimePeriod:    0,
	//		//CPURealtimeRuntime:   0,
	//		//CPUShares:            0,
	//		//CapAdd:               nil,
	//		//CapDrop:              nil,
	//		//Cgroup:               "",
	//		//CgroupParent:         "",
	//		//CgroupnsMode:         "",
	//		//ConsoleSize:          nil,
	//		//ContainerIDFile:      "",
	//		//CpusetCpus:           "",
	//		//CpusetMems:           "",
	//		//DNS:                  nil,
	//		//DNSOptions:           nil,
	//		//DNSSearch:            nil,
	//		//DeviceCgroupRules:    nil,
	//		//DeviceRequests:       nil,
	//		//Devices:              nil,
	//		//ExtraHosts:           nil,
	//		//GroupAdd:             nil,
	//		//IOMaximumBandwidth:   0,
	//		//IOMaximumIOps:        0,
	//		//Init:                 false,
	//		//IpcMode:              "",
	//		//Isolation:            "",
	//		//KernelMemory:         0,
	//		//KernelMemoryTCP:      0,
	//		//Links:                nil,
	//		//LogConfig:            nil,
	//		//MaskedPaths:          nil,
	//		Memory: 4096 * 1024 * 1024,
	//		//MemoryReservation:    0,
	//		//MemorySwap:           0,
	//		//MemorySwappiness:     0,
	//		Mounts: make([]*models.Mount, 0),
	//		//NanoCPUs:             0,
	//		NetworkMode: "bridge",
	//		//OomKillDisable:       false,
	//		OomScoreAdj: 100, //default value provided by podman (avoids warning)
	//		//PidMode:              "",
	//		//PidsLimit:            0,
	//		PortBindings: models.PortMap{
	//			"8080/tcp": []models.PortBinding{
	//				{
	//					HostIP:   "192.168.124.1",
	//					HostPort: "8888",
	//				},
	//			},
	//		},
	//		//Privileged:           false,
	//		//PublishAllPorts:      false,
	//		//ReadonlyPaths:        nil,
	//		//ReadonlyRootfs:       false,
	//		//RestartPolicy:        nil,
	//		//Runtime:              "",
	//		SecurityOpt: []string{"label=disable"},
	//		//ShmSize:              0,
	//		//StorageOpt:           nil,
	//		//Sysctls:              nil,
	//		//Tmpfs:                nil,
	//		//UTSMode:              "",
	//		//Ulimits:              nil,
	//		UsernsMode: "keep-id",
	//		//VolumeDriver:         "",
	//		//VolumesFrom:          nil,
	//	},
	//	//Hostname:         "",
	//	Image: images.GetRouterImageName(),
	//	Labels: map[string]string{
	//		"label1": "value1",
	//	},
	//	//MacAddress:       "",
	//	Name:            *params.Name,
	//	NetworkDisabled: false,
	//	NetworkingConfig: &models.NetworkingConfig{
	//		map[string]models.EndpointSettings{
	//			"pitoca": models.EndpointSettings{
	//				NetworkID: "pitoca",
	//			},
	//		},
	//	},
	//	//OnBuild:     nil,
	//	//OpenStdin:   false,
	//	//Shell:       nil,
	//	//StdinOnce:   false,
	//	//StopSignal:  "",
	//	//StopTimeout: 0,
	//	//Tty:         false,
	//	//UnsetEnv:    nil,
	//	//UnsetEnvAll: false,
	//	//Volumes: nil,
	//	//WorkingDir: "",
	//}
	//params.Body.User = fmt.Sprintf("%d:%d", os.Getuid(), os.Getgid())
	//
	//if ccli.engine == "docker" {
	//	params.Body.User = "0:0"
	//}
	//params.Body.HostConfig.Mounts = append(params.Body.HostConfig.Mounts, &models.Mount{
	//	Target:  "/tmp/podman.sock",
	//	Options: []string{"z"},
	//	Source:  "/var/run/user/1000/podman/podman.sock",
	//	Type:    "bind",
	//})
	//params.Body.HostConfig.Mounts = append(params.Body.HostConfig.Mounts, &models.Mount{
	//	Target:  "/tmp/docker.sock",
	//	Options: []string{"z"},
	//	Source:  "/run/docker.sock",
	//	Type:    "bind",
	//})
	//params.Body.HostConfig.Binds = append(params.Body.HostConfig.Binds,
	//	fmt.Sprintf("%s:%s:z", "my-volume", "/opt/my-volume"))
	//_, _ = cli.ContainerDelete(containers_compat.NewContainerDeleteParams().WithName(*params.Name).WithForce(boolTrue()))
	//_, err = cli.ContainerCreate(params)
	//assert.Assert(t, err)
}

func TestContainerExec(t *testing.T) {
	//ccli, err := NewCompatClient("", "")
	ccli, err := NewCompatClient("/tmp/docker.sock", "")
	assert.Assert(t, err)
	out, err := ccli.ContainerExec("my-container", []string{"ls", "-l", "/"})
	assert.Assert(t, err)
	assert.Assert(t, len(out) > 0)
	t.Logf(out)
}

func TestContainerLogs(t *testing.T) {
	ccli, err := NewCompatClient("", "")
	//ccli, err := NewCompatClient("/run/docker.sock", "")
	assert.Assert(t, err)
	out, err := ccli.ContainerLogs("skupper-router")
	assert.Assert(t, err)
	assert.Assert(t, len(out) > 0)
	t.Log(out)
}
