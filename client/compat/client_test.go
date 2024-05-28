package compat

import (
	"testing"

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
