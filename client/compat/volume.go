package compat

import (
	"github.com/skupperproject/skupper/client/container"
	"github.com/skupperproject/skupper/client/generated/libpod/models"
)

func (c *CompatClient) VolumeCreate(volume *container.Volume) (*container.Volume, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CompatClient) VolumeInspect(id string) (*container.Volume, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CompatClient) VolumeRemove(id string) error {
	//TODO implement me
	panic("implement me")
}

func (c *CompatClient) VolumeList() ([]*container.Volume, error) {
	//TODO implement me
	panic("implement me")
}

func FilesToMounts(c *container.Container) []*models.Volume {
	var mounts []*models.Volume
	for _, fm := range c.FileMounts {
		m := &models.Volume{
			Mountpoint: &fm.Destination,
			Name:       &fm.Source,
			//Options:    fm.Options,
		}
		/*			Type:        "bind",
					Source:      fm.Source,
					Destination: fm.Destination,
					Options:     fm.Options,
		*/
		mounts = append(mounts, m)
	}
	return mounts
}

func VolumesToNamedVolumes(c *container.Container) []*models.NamedVolume {
	var namedVolumes []*models.NamedVolume
	for _, v := range c.Mounts {
		m := &models.NamedVolume{
			Dest:    v.Destination,
			Name:    v.Name,
			Options: []string{"z", "U"}, // shared between containers
		}
		namedVolumes = append(namedVolumes, m)
	}
	return namedVolumes
}
