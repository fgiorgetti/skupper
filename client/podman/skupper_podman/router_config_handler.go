package skupper_podman

import (
	"fmt"

	"github.com/skupperproject/skupper/api/types"
	"github.com/skupperproject/skupper/api/types/v2"
	"github.com/skupperproject/skupper/client/container"
	"github.com/skupperproject/skupper/client/generated/libpod/client/volumes"
	"github.com/skupperproject/skupper/client/podman"
)

type RouterConfigHandlerPodman struct {
	cli *podman.PodmanRestClient
}

func NewRouterConfigHandlerPodman(cli *podman.PodmanRestClient) *RouterConfigHandlerPodman {
	return &RouterConfigHandlerPodman{
		cli: cli,
	}
}

func (r *RouterConfigHandlerPodman) GetRouterConfig() (*v2.RouterConfig, error) {
	var configVolume *container.Volume
	configVolume, err := r.cli.VolumeInspect(types.TransportConfigMapName)
	if err != nil {
		return nil, fmt.Errorf("error retrieving volume %s - %v", types.TransportConfigMapName, err)
	}
	routerConfigStr, err := configVolume.ReadFile(types.TransportConfigFile)
	if err != nil {
		return nil, fmt.Errorf("error reading config file %s from volume %s - %v",
			types.TransportConfigFile, types.TransportConfigMapName, err)
	}
	routerConfig, err := v2.UnmarshalRouterConfig(routerConfigStr)
	if err != nil {
		return nil, fmt.Errorf("error parsing config file %s from volume %s - %v",
			types.TransportConfigFile, types.TransportConfigMapName, err)
	}
	return &routerConfig, nil
}

func (r *RouterConfigHandlerPodman) SaveRouterConfig(config *v2.RouterConfig) error {
	var configVolume *container.Volume
	configVolume, err := r.cli.VolumeInspect(types.TransportConfigMapName)
	if err != nil {
		if _, notFound := err.(*volumes.VolumeInspectLibpodNotFound); !notFound {
			return fmt.Errorf("error retrieving volume %s - %v", types.TransportConfigMapName, err)
		}
		// try to create volume since not found given
		if configVolume, err = r.cli.VolumeCreate(&container.Volume{Name: types.TransportConfigMapName}); err != nil {
			return fmt.Errorf("error creating volume %s - %v", types.TransportConfigMapName, err)
		}
	}
	routerConfig, err := v2.MarshalRouterConfig(*config)
	if err != nil {
		return fmt.Errorf("error serializing router config - %v", err)
	}
	_, err = configVolume.CreateFile(types.TransportConfigFile, []byte(routerConfig), true)
	if err != nil {
		return fmt.Errorf("error creating router config - %v", err)
	}
	return nil
}
