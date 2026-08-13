package site

import (
	"reflect"
	"testing"

	"github.com/skupperproject/skupper/internal/qdr"
	skupperv2alpha1 "github.com/skupperproject/skupper/pkg/apis/skupper/v2alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestNetworkAccessConfig_Apply(t *testing.T) {
	id := "router-1"
	siteId := "site-1"
	version := "v2.0"
	notEdge := false
	helloAge := 10

	type networkAccessConfig struct {
		listeners   map[string]qdr.Listener
		profilePath string
	}
	type configEntities struct {
		networkId   string
		listeners   map[string]qdr.Listener
		sslProfiles map[string]qdr.SslProfile
		autoLinks   map[string]qdr.AutoLink
	}
	tests := []struct {
		name         string
		config       networkAccessConfig
		routerConfig configEntities
		expect       bool
	}{
		{
			name: "no listeners",
			config: networkAccessConfig{
				listeners:   map[string]qdr.Listener{},
				profilePath: "",
			},
			routerConfig: configEntities{
				networkId:   "my-network",
				listeners:   map[string]qdr.Listener{},
				sslProfiles: map[string]qdr.SslProfile{},
			},
			expect: false,
		},
		{
			name: "add inter-network listener with autolink when networkId is set",
			config: networkAccessConfig{
				listeners: map[string]qdr.Listener{
					"na1-inter-network": qdr.Listener{
						Name:             "na1-inter-network",
						Role:             qdr.RoleInterNetwork,
						Host:             "0.0.0.0",
						Port:             55555,
						SslProfile:       "my-creds",
						SaslMechanisms:   "EXTERNAL",
						AuthenticatePeer: true,
					},
				},
				profilePath: "/etc/skupper",
			},
			routerConfig: configEntities{
				networkId:   "my-network",
				listeners:   map[string]qdr.Listener{},
				sslProfiles: map[string]qdr.SslProfile{},
			},
			expect: true,
		},
		{
			name: "same listener and autolink already present is idempotent",
			config: networkAccessConfig{
				listeners: map[string]qdr.Listener{
					"na1-inter-network": qdr.Listener{
						Name:             "na1-inter-network",
						Role:             qdr.RoleInterNetwork,
						Host:             "0.0.0.0",
						Port:             55555,
						SslProfile:       "my-creds",
						SaslMechanisms:   "EXTERNAL",
						AuthenticatePeer: true,
					},
				},
				profilePath: "/etc/skupper",
			},
			routerConfig: configEntities{
				networkId: "my-network",
				listeners: map[string]qdr.Listener{
					"na1-inter-network": qdr.Listener{
						Name:             "na1-inter-network",
						Role:             qdr.RoleInterNetwork,
						Host:             "0.0.0.0",
						Port:             55555,
						SslProfile:       "my-creds",
						SaslMechanisms:   "EXTERNAL",
						AuthenticatePeer: true,
					},
				},
				sslProfiles: map[string]qdr.SslProfile{
					"my-creds": qdr.SslProfile{
						Name:           "my-creds",
						CertFile:       "/etc/skupper/my-creds/tls.crt",
						PrivateKeyFile: "/etc/skupper/my-creds/tls.key",
						CaCertFile:     "/etc/skupper/my-creds/ca.crt",
					},
				},
				autoLinks: map[string]qdr.AutoLink{
					"na1-inter-network-listener-autoLink": qdr.AutoLink{
						Name:            "na1-inter-network-listener-autoLink",
						ExternalAddress: "_xtopo/my-network",
						Direction:       qdr.DirectionIn,
						Connection:      "na1-inter-network",
					},
				},
			},
			expect: false,
		},
		{
			name: "remove listener also removes its autolink",
			config: networkAccessConfig{
				listeners:   map[string]qdr.Listener{},
				profilePath: "",
			},
			routerConfig: configEntities{
				networkId: "my-network",
				listeners: map[string]qdr.Listener{
					"na1-inter-network": qdr.Listener{
						Name:             "na1-inter-network",
						Role:             qdr.RoleInterNetwork,
						Host:             "0.0.0.0",
						Port:             55555,
						SslProfile:       "my-creds",
						SaslMechanisms:   "EXTERNAL",
						AuthenticatePeer: true,
					},
				},
				sslProfiles: map[string]qdr.SslProfile{
					"my-creds": qdr.SslProfile{
						Name:           "my-creds",
						CertFile:       "/etc/skupper/my-creds/tls.crt",
						PrivateKeyFile: "/etc/skupper/my-creds/tls.key",
						CaCertFile:     "/etc/skupper/my-creds/ca.crt",
					},
				},
			},
			expect: true,
		},
		{
			name: "empty networkId removes autolinks but keeps listeners",
			config: networkAccessConfig{
				listeners: map[string]qdr.Listener{
					"na1-inter-network": qdr.Listener{
						Name:             "na1-inter-network",
						Role:             qdr.RoleInterNetwork,
						Host:             "0.0.0.0",
						Port:             55555,
						SslProfile:       "my-creds",
						SaslMechanisms:   "EXTERNAL",
						AuthenticatePeer: true,
					},
				},
				profilePath: "/etc/skupper",
			},
			routerConfig: configEntities{
				networkId: "",
				listeners: map[string]qdr.Listener{
					"na1-inter-network": qdr.Listener{
						Name:             "na1-inter-network",
						Role:             qdr.RoleInterNetwork,
						Host:             "0.0.0.0",
						Port:             55555,
						SslProfile:       "my-creds",
						SaslMechanisms:   "EXTERNAL",
						AuthenticatePeer: true,
					},
				},
				sslProfiles: map[string]qdr.SslProfile{
					"my-creds": qdr.SslProfile{
						Name:           "my-creds",
						CertFile:       "/etc/skupper/my-creds/tls.crt",
						PrivateKeyFile: "/etc/skupper/my-creds/tls.key",
						CaCertFile:     "/etc/skupper/my-creds/ca.crt",
					},
				},
				autoLinks: map[string]qdr.AutoLink{
					"na1-inter-network-listener-autoLink": qdr.AutoLink{
						Name:            "na1-inter-network-listener-autoLink",
						ExternalAddress: "_xtopo/my-network",
						Direction:       qdr.DirectionIn,
						Connection:      "na1-inter-network",
					},
				},
			},
			expect: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &NetworkAccessConfig{}
			g.listeners = tt.config.listeners
			g.profilePath = tt.config.profilePath

			argsConfig := qdr.InitialConfig(id, siteId, version, notEdge, helloAge)
			argsConfig.Network.NetworkId = tt.routerConfig.networkId
			argsConfig.Listeners = tt.routerConfig.listeners
			argsConfig.SslProfiles = tt.routerConfig.sslProfiles
			for _, al := range tt.routerConfig.autoLinks {
				argsConfig.AddAutoLink(al)
			}

			if got := g.Apply(&argsConfig); got != tt.expect {
				t.Errorf("NetworkAccessConfig.Apply() = %v, want %v (subtest: %s)", got, tt.expect, tt.name)
			}
		})
	}
}

func TestNetworkAccessMap_DesiredConfig(t *testing.T) {
	type args struct {
		targetGroups []string
		profilePath  string
	}
	tests := []struct {
		name string
		m    NetworkAccessMap
		args args
		want NetworkAccessConfig
	}{
		{
			name: "empty map",
			m:    NetworkAccessMap{},
			args: args{
				targetGroups: []string{},
				profilePath:  "",
			},
			want: NetworkAccessConfig{
				listeners:   map[string]qdr.Listener{},
				profilePath: "",
			},
		},
		{
			name: "single entry produces inter-network listener",
			m: NetworkAccessMap{
				"na1": &skupperv2alpha1.NetworkAccess{
					ObjectMeta: v1.ObjectMeta{
						Name:      "na1",
						Namespace: "test",
					},
					Spec: skupperv2alpha1.NetworkAccessSpec{
						BindHost:       "0.0.0.0",
						TlsCredentials: "my-creds",
						Port:           55555,
					},
				},
			},
			args: args{
				targetGroups: []string{},
				profilePath:  "/etc/skupper",
			},
			want: NetworkAccessConfig{
				listeners: map[string]qdr.Listener{
					"na1-inter-network": qdr.Listener{
						Name:             "na1-inter-network",
						Role:             qdr.RoleInterNetwork,
						Host:             "0.0.0.0",
						Port:             55555,
						SslProfile:       "my-creds",
						SaslMechanisms:   "EXTERNAL",
						AuthenticatePeer: true,
					},
				},
				profilePath: "/etc/skupper",
			},
		},
		{
			name: "multiple entries produce multiple listeners",
			m: NetworkAccessMap{
				"na1": &skupperv2alpha1.NetworkAccess{
					ObjectMeta: v1.ObjectMeta{
						Name:      "na1",
						Namespace: "test",
					},
					Spec: skupperv2alpha1.NetworkAccessSpec{
						BindHost:       "0.0.0.0",
						TlsCredentials: "creds-1",
						Port:           55555,
					},
				},
				"na2": &skupperv2alpha1.NetworkAccess{
					ObjectMeta: v1.ObjectMeta{
						Name:      "na2",
						Namespace: "test",
					},
					Spec: skupperv2alpha1.NetworkAccessSpec{
						BindHost:       "0.0.0.0",
						TlsCredentials: "creds-2",
						Port:           55556,
					},
				},
			},
			args: args{
				targetGroups: []string{},
				profilePath:  "/etc/skupper",
			},
			want: NetworkAccessConfig{
				listeners: map[string]qdr.Listener{
					"na1-inter-network": qdr.Listener{
						Name:             "na1-inter-network",
						Role:             qdr.RoleInterNetwork,
						Host:             "0.0.0.0",
						Port:             55555,
						SslProfile:       "creds-1",
						SaslMechanisms:   "EXTERNAL",
						AuthenticatePeer: true,
					},
					"na2-inter-network": qdr.Listener{
						Name:             "na2-inter-network",
						Role:             qdr.RoleInterNetwork,
						Host:             "0.0.0.0",
						Port:             55556,
						SslProfile:       "creds-2",
						SaslMechanisms:   "EXTERNAL",
						AuthenticatePeer: true,
					},
				},
				profilePath: "/etc/skupper",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.DesiredConfig(tt.args.targetGroups, tt.args.profilePath); !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("NetworkAccessMap.DesiredConfig() = %v, want %v", *got, tt.want)
			}
		})
	}
}

func TestNetworkAccessMap_DesiredConfigWithAvailableCredentials(t *testing.T) {
	na := &skupperv2alpha1.NetworkAccess{
		ObjectMeta: v1.ObjectMeta{Name: "na1", Namespace: "ns"},
		Spec: skupperv2alpha1.NetworkAccessSpec{
			TlsCredentials: "missing-secret",
			BindHost:       "0.0.0.0",
			Port:           55555,
		},
	}
	naNoTLS := &skupperv2alpha1.NetworkAccess{
		ObjectMeta: v1.ObjectMeta{Name: "na2", Namespace: "ns"},
		Spec: skupperv2alpha1.NetworkAccessSpec{
			BindHost: "0.0.0.0",
			Port:     55556,
		},
	}
	m := NetworkAccessMap{"na1": na, "na2": naNoTLS}

	allow := func(name string) bool { return name != "missing-secret" }
	got := m.DesiredConfigWithAvailableCredentials([]string{}, "/certs", allow)
	if _, ok := got.listeners["na1-inter-network"]; ok {
		t.Errorf("expected na1 to be excluded when TLS secret is missing")
	}
	if _, ok := got.listeners["na2-inter-network"]; !ok {
		t.Errorf("expected na2 to be included when TLS is not required")
	}

	gotAll := m.DesiredConfigWithAvailableCredentials([]string{}, "/certs", func(string) bool { return true })
	if _, ok := gotAll.listeners["na1-inter-network"]; !ok {
		t.Errorf("expected na1 to be included when TLS secret is present")
	}
	if _, ok := gotAll.listeners["na2-inter-network"]; !ok {
		t.Errorf("expected na2 to be included when TLS secret is present")
	}
}
