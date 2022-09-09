package podman

import (
	"github.com/skupperproject/skupper/client/container"
	"github.com/skupperproject/skupper/client/generated/libpod/client/volumes"
	"github.com/skupperproject/skupper/client/generated/libpod/models"
)

func (p *PodmanRestClient) VolumeCreate(volume *container.Volume) (*container.Volume, error) {
	cli := volumes.New(p.RestClient, formats)
	params := volumes.NewVolumeCreateLibpodParams()
	params.Create = ToVolumeCreateOptions(volume)
	created, err := cli.VolumeCreateLibpod(params)
	if err != nil {
		return nil, err
	}
	return FromCreatedToVolume(created), nil
}

func ToVolumeCreateOptions(v *container.Volume) *models.VolumeCreateOptions {
	nv := &models.VolumeCreateOptions{
		Name: v.Name,
	}
	return nv
}

func FromCreatedToVolume(created *volumes.VolumeCreateLibpodCreated) *container.Volume {
	v := &container.Volume{
		Name:   created.Payload.Name,
		Source: created.Payload.Mountpoint,
		Labels: created.Payload.Labels,
	}
	return v
}

func (p *PodmanRestClient) VolumeRemove(id string) error {
	cli := volumes.New(p.RestClient, formats)
	params := volumes.NewVolumeDeleteLibpodParams()
	params.Name = id
	_, err := cli.VolumeDeleteLibpod(params)
	if err != nil {
		return err
	}
	return nil
}

func VolumesToMounts(c *container.Container) []*models.Mount {
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

func (p *PodmanRestClient) VolumeInspect(id string) (*container.Volume, error) {
	cli := volumes.New(p.RestClient, formats)
	params := volumes.NewVolumeInspectLibpodParams()
	params.Name = id
	vi, err := cli.VolumeInspectLibpod(params)
	if err != nil {
		return nil, err
	}
	v := FromInspectToVolume(vi)
	return v, err
}

func FromInspectToVolume(vi *volumes.VolumeInspectLibpodOK) *container.Volume {
	return &container.Volume{
		Name:   vi.Payload.Name,
		Source: vi.Payload.Mountpoint,
		Labels: vi.Payload.Labels,
	}
}
