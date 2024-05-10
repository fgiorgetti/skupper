package compat

import (
	"strings"

	"github.com/skupperproject/skupper/client/container"
	system "github.com/skupperproject/skupper/client/generated/libpod/client/system_compat"
)

func (p *CompatRestClient) Version() (*container.Version, error) {
	systemCli := system.New(p.RestClient, formats)
	params := system.NewSystemVersionParams()
	info, err := systemCli.SystemVersion(params)
	if err != nil {
		return nil, err
	}
	v := &container.Version{}
	if info.Payload != nil {
		v.Server = container.VersionInfo{
			Version:    info.Payload.Version,
			APIVersion: info.Payload.APIVersion,
		}
		v.Arch = info.Payload.Arch
		v.Kernel = info.Payload.KernelVersion
		v.OS = info.Payload.Os
		for _, cmp := range info.Payload.Components {
			if strings.Contains(strings.ToLower(cmp.Name), "podman") {
				v.Engine = "podman"
			} else {
				v.Engine = "docker"
			}
		}
	}

	return v, nil
}
