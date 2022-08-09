package podman

import (
	"fmt"

	"github.com/skupperproject/skupper/client/container"
	"github.com/skupperproject/skupper/client/generated/libpod/client/networks"
	"github.com/skupperproject/skupper/client/generated/libpod/models"
)

func (p *PodmanRestClient) NetworkList() ([]*container.Network, error) {
	cli := networks.New(p.RestClient, formats)
	params := networks.NewNetworkListLibpodParams()
	res, err := cli.NetworkListLibpod(params)
	if err != nil {
		return nil, fmt.Errorf("error listing networks: %v", err)
	}
	return ToNetworkInfoList(res.Payload), nil
}

func ToNetworkInfoList(networks []*models.Network) []*container.Network {
	var nets []*container.Network
	for _, net := range networks {
		nets = append(nets, ToNetworkInfo(net))
	}
	return nets
}

func ToNetworkInfo(network *models.Network) *container.Network {
	var ss []*container.Subnet

	for _, s := range network.Subnets {
		ss = append(ss, &container.Subnet{
			Subnet:  s.Subnet,
			Gateway: s.Gateway,
		})
	}

	n := &container.Network{
		ID:        network.ID,
		Name:      network.Name,
		Subnets:   ss,
		Driver:    network.Driver,
		DNS:       network.DNSEnabled,
		Internal:  network.Internal,
		Labels:    network.Labels,
		Options:   network.Options,
		CreatedAt: network.Created.String(),
	}

	return n
}

/*
	NetworkInspect()
	NetworkCreate()
	NetworkRemove()
	NetworkConnect()
	NetworkDisconnect()

*/
