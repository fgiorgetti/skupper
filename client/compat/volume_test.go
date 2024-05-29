package compat

import (
	"encoding/json"
	"testing"

	"github.com/skupperproject/skupper/client/container"
	"gotest.tools/assert"
)

func TestVolumeCreate(t *testing.T) {
	cli, err := NewCompatClient("/tmp/docker.sock", "")
	assert.Assert(t, err)
	volume, err := cli.VolumeCreate(&container.Volume{
		Name: "batatinha",
		Labels: map[string]string{
			"skupper.io/type": "volume",
		},
	})
	volumeJson, _ := json.MarshalIndent(volume, "", "  ")
	t.Logf("volumeJson: %v", string(volumeJson))
}

func TestVolumeInspect(t *testing.T) {
	cli, err := NewCompatClient("/tmp/docker.sock", "")
	assert.Assert(t, err)
	volume, err := cli.VolumeInspect("batatinha")
	assert.Assert(t, err)
	volumeJson, _ := json.MarshalIndent(volume, "", "  ")
	t.Logf("volumeJson: %v", string(volumeJson))
}

func TestVolumeList(t *testing.T) {
	cli, err := NewCompatClient("/tmp/docker.sock", "")
	assert.Assert(t, err)
	volumeList, err := cli.VolumeList()
	assert.Assert(t, err)
	assert.Assert(t, len(volumeList) > 0)
	for _, volume := range volumeList {
		t.Logf("volume: %v", volume)
	}
}

func TestVolumeRemove(t *testing.T) {
	cli, err := NewCompatClient("/tmp/docker.sock", "")
	assert.Assert(t, err)
	err = cli.VolumeRemove("manualvolume")
	assert.Assert(t, err)
}
