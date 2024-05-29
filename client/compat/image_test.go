package compat

import (
	"context"
	"encoding/json"
	"testing"

	"gotest.tools/assert"
)

func TestImageList(t *testing.T) {
	ccli, err := NewCompatClient("", "")
	//ccli, err := NewCompatClient("/run/docker.sock", "")
	assert.Assert(t, err)
	images, err := ccli.ImageList()
	assert.Assert(t, err)
	for _, image := range images {
		if image.Repository == "quay.io/skupper/skupper-router:main" {
			imageJson, _ := json.MarshalIndent(image, "", "  ")
			t.Log(string(imageJson))
		}
	}
}

func TestImageInspect(t *testing.T) {
	ccli, err := NewCompatClient("", "")
	//ccli, err := NewCompatClient("/run/docker.sock", "")
	assert.Assert(t, err)
	img, err := ccli.ImageInspect("quay.io/skupper/skupper-router:main")
	assert.Assert(t, err)
	imageJson, _ := json.MarshalIndent(img, "", "  ")
	t.Log(string(imageJson))
}

func TestImagePull(t *testing.T) {
	//ccli, err := NewCompatClient("", "")
	ccli, err := NewCompatClient("/tmp/docker.sock", "")
	assert.Assert(t, err)
	//os.Setenv("REGISTRY_AUTH_FILE", "/home/fgiorget/.docker/config.json")
	err = ccli.ImagePull(context.Background(), "quay.io/fgiorgetti/registry-private")
	//err = ccli.ImagePull(context.Background(), "quay.io/skupper/skupper-router:main")
	//err = ccli.ImagePull(context.Background(), "quay.io/skupper/skupper-router@sha256:4643bdf98d8e551eb7292718c16428ff467c5657b0d98bdbe7b559f516591c64")
	assert.Assert(t, err)
}
