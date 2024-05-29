package compat

import (
	"fmt"
	"testing"

	"github.com/skupperproject/skupper/client/generated/libpod/client/volumes_compat"
	"gotest.tools/assert"
)

func TestCompatRestClient(t *testing.T) {
	//cli, err := NewCompatClient("/run/user/1000/podman/podman.sock", "")
	cli, err := NewCompatClient("/run/docker.sock", "")
	if err != nil {
		t.Fatal(err)
	}
	version, err := cli.Version()
	assert.Assert(t, err)
	assert.Assert(t, version.Server.Version != "")
}

func TestToAPIError(t *testing.T) {
	notFoundErr := volumes_compat.NewVolumeDeleteNotFound()
	assert.Assert(t, notFoundErr != nil)
	notFoundErr.Payload = new(volumes_compat.VolumeDeleteNotFoundBody)
	notFoundErr.Payload.Message = "Sample error message"
	notFoundErr.Payload.Because = "Because it has to fail"
	notFoundErr.Payload.ResponseCode = 404
	// validating result only and both result and error
	for _, e := range []interface{}{notFoundErr, fmt.Errorf("unused error")} {
		apiErr := ToAPIError(e)
		assert.Assert(t, apiErr != nil)
		assert.Equal(t, apiErr.Message, notFoundErr.Payload.Message)
		assert.Equal(t, apiErr.Because, notFoundErr.Payload.Because)
		assert.Equal(t, apiErr.ResponseCode, notFoundErr.Payload.ResponseCode)
	}

	// validating none
	assert.Assert(t, ToAPIError(nil) == nil)
}
