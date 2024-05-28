package compat

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/skupperproject/skupper/client/container"
	"gotest.tools/assert"
)

func TestNetworkList(t *testing.T) {
	cli, err := NewCompatClient("/tmp/docker.sock", "")
	assert.Assert(t, err)
	networks, err := cli.NetworkList()
	assert.Assert(t, err)
	assert.Assert(t, len(networks) > 0)
	for _, network := range networks {
		networkJson, _ := json.MarshalIndent(network, "", "  ")
		fmt.Println(string(networkJson))
	}
}

func TestNetworkCreate(t *testing.T) {
	cli, err := NewCompatClient("/tmp/docker.sock", "")
	assert.Assert(t, err)

	n := &container.Network{
		Name: "my-network",
		Subnets: []*container.Subnet{
			{
				Subnet:  "10.1.2.0/24",
				Gateway: "10.1.2.1",
			},
		},
		DNS: true,
		Labels: map[string]string{
			"foo": "bar",
		},
	}
	_ = cli.NetworkRemove(n.Name)
	nc, err := cli.NetworkCreate(n)
	assert.Assert(t, err)
	assert.Assert(t, nc != nil)
	//ncJson, _ := json.MarshalIndent(nc, "", "  ")
	//fmt.Println(string(ncJson))
}

func TestNetworkConnect(t *testing.T) {
	cli, err := NewCompatClient("/tmp/docker.sock", "")
	assert.Assert(t, err)
	_ = cli.NetworkDisconnect("my-network", "my-container")
	err = cli.NetworkConnect("my-network", "my-container", "alias1", "alias2")
	assert.Assert(t, err)
	_ = cli.NetworkDisconnect("my-network", "my-container")
	assert.Assert(t, err)

}
