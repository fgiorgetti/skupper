package site

import (
	"testing"

	"github.com/skupperproject/skupper/internal/qdr"
	skupperv2alpha1 "github.com/skupperproject/skupper/pkg/apis/skupper/v2alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestNetwork_Update(t *testing.T) {
	type fields struct {
		networkId string
		image     string
	}
	type args struct {
		network *skupperv2alpha1.Network
	}
	tests := []struct {
		name   string
		from   fields
		to     args
		expect bool
	}{
		{
			name: "nil network clears networkId",
			from: fields{
				networkId: "my-van",
			},
			to: args{
				network: nil,
			},
			expect: true,
		},
		{
			name: "nil network when already empty",
			from: fields{
				networkId: "",
			},
			to: args{
				network: nil,
			},
			expect: false,
		},
		{
			name: "new network sets networkId and image",
			from: fields{
				networkId: "",
			},
			to: args{
				network: &skupperv2alpha1.Network{
					ObjectMeta: v1.ObjectMeta{Name: "network", Namespace: "test"},
					Spec: skupperv2alpha1.NetworkSpec{
						NetworkId: "my-van",
						Image:     "my-image",
					},
				},
			},
			expect: true,
		},
		{
			name: "same network spec returns false",
			from: fields{
				networkId: "my-van",
				image:     "my-image",
			},
			to: args{
				network: &skupperv2alpha1.Network{
					ObjectMeta: v1.ObjectMeta{Name: "network", Namespace: "test"},
					Spec: skupperv2alpha1.NetworkSpec{
						NetworkId: "my-van",
						Image:     "my-image",
					},
				},
			},
			expect: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Network{
				NetworkId: tt.from.networkId,
				Image:     tt.from.image,
			}
			if got := m.Update("test", "network", tt.to.network); got != tt.expect {
				t.Errorf("Network.Update() = %v, want %v", got, tt.expect)
			}
		})
	}
}

func TestNetwork_Apply(t *testing.T) {
	tests := []struct {
		name         string
		newNetworkId string
		config       qdr.RouterConfig
		expect       bool
	}{
		{
			name:         "networkId unchanged returns false",
			newNetworkId: "my-van",
			config: func() qdr.RouterConfig {
				rc := qdr.InitialConfig("r", "s", "v", false, 3)
				rc.Network.NetworkId = "my-van"
				return rc
			}(),
			expect: false,
		},
		{
			name:         "new networkId returns true",
			newNetworkId: "my-van",
			config:       qdr.InitialConfig("r", "s", "v", false, 3),
			expect:       true,
		},
		{
			name:         "networkId updated returns true",
			newNetworkId: "my-van-2",
			config: func() qdr.RouterConfig {
				rc := qdr.InitialConfig("r", "s", "v", false, 3)
				rc.Network.NetworkId = "my-van"
				return rc
			}(),
			expect: true,
		},
		{
			name:         "clearing networkId removes inter-network connectors and their autolinks",
			newNetworkId: "",
			config: func() qdr.RouterConfig {
				rc := qdr.InitialConfig("r", "s", "v", false, 3)
				rc.Network.NetworkId = "my-van"
				rc.AddConnector(qdr.Connector{
					Name: "inet-connector",
					Role: qdr.RoleInterNetwork,
					Host: "remote.example.com",
					Port: "55555",
				})
				rc.AddConnector(qdr.Connector{
					Name: "normal-connector",
					Role: qdr.RoleNormal,
					Host: "peer.example.com",
					Port: "55671",
				})
				rc.AddAutoLink(qdr.AutoLink{
					Name:       "inet-connector-autoLink",
					Connection: "inet-connector",
				})
				rc.AddAutoLink(qdr.AutoLink{
					Name:       "other-autoLink",
					Connection: "normal-connector",
				})
				return rc
			}(),
			expect: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Network{
				NetworkId: tt.newNetworkId,
			}
			config := tt.config
			if got := m.Apply(&config); got != tt.expect {
				t.Errorf("Network.Apply() = %v, want %v", got, tt.expect)
			}
			if tt.name == "clearing networkId removes inter-network connectors and their autolinks" {
				if _, ok := config.Connectors["inet-connector"]; ok {
					t.Errorf("expected inter-network connector to be removed")
				}
				if _, ok := config.Connectors["normal-connector"]; !ok {
					t.Errorf("expected normal connector to remain")
				}
				if _, ok := config.AutoLinks["inet-connector-autoLink"]; ok {
					t.Errorf("expected autolink for inter-network connector to be removed")
				}
				if _, ok := config.AutoLinks["other-autoLink"]; !ok {
					t.Errorf("expected autolink for normal connector to remain")
				}
			}
		})
	}
}
