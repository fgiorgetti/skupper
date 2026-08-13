package site

import (
	"testing"

	"github.com/skupperproject/skupper/internal/qdr"
	skupperv2alpha1 "github.com/skupperproject/skupper/pkg/apis/skupper/v2alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestNetworkLink_Update(t *testing.T) {
	type fields struct {
		networkId string
		link      *skupperv2alpha1.NetworkLink
	}
	type args struct {
		networkId string
		desired   *skupperv2alpha1.NetworkLink
	}
	tests := []struct {
		name    string
		current fields
		new     args
		expect  bool
	}{
		{
			name: "networkId change returns true",
			current: fields{
				networkId: "old-van",
				link:      nil,
			},
			new: args{
				networkId: "new-van",
				desired:   nil,
			},
			expect: true,
		},
		{
			name: "add link returns true",
			current: fields{
				networkId: "my-van",
				link:      nil,
			},
			new: args{
				networkId: "my-van",
				desired: &skupperv2alpha1.NetworkLink{
					ObjectMeta: v1.ObjectMeta{Name: "link1", Namespace: "test"},
					Spec: skupperv2alpha1.NetworkLinkSpec{
						Hostname: "remote.example.com",
						Port:     55555,
					},
				},
			},
			expect: true,
		},
		{
			name: "nil link clears m.Link and returns true",
			current: fields{
				networkId: "my-van",
				link: &skupperv2alpha1.NetworkLink{
					ObjectMeta: v1.ObjectMeta{Name: "link1", Namespace: "test"},
					Spec: skupperv2alpha1.NetworkLinkSpec{
						Hostname: "remote.example.com",
						Port:     55555,
					},
				},
			},
			new: args{
				networkId: "my-van",
				desired:   nil,
			},
			expect: true,
		},
		{
			name: "spec change returns true",
			current: fields{
				networkId: "my-van",
				link: &skupperv2alpha1.NetworkLink{
					ObjectMeta: v1.ObjectMeta{Name: "link1", Namespace: "test"},
					Spec: skupperv2alpha1.NetworkLinkSpec{
						Hostname: "old.example.com",
						Port:     55555,
					},
				},
			},
			new: args{
				networkId: "my-van",
				desired: &skupperv2alpha1.NetworkLink{
					ObjectMeta: v1.ObjectMeta{Name: "link1", Namespace: "test"},
					Spec: skupperv2alpha1.NetworkLinkSpec{
						Hostname: "new.example.com",
						Port:     55555,
					},
				},
			},
			expect: true,
		},
		{
			name: "identical spec returns false",
			current: fields{
				networkId: "my-van",
				link: &skupperv2alpha1.NetworkLink{
					ObjectMeta: v1.ObjectMeta{Name: "link1", Namespace: "test"},
					Spec: skupperv2alpha1.NetworkLinkSpec{
						Hostname: "remote.example.com",
						Port:     55555,
					},
				},
			},
			new: args{
				networkId: "my-van",
				desired: &skupperv2alpha1.NetworkLink{
					ObjectMeta: v1.ObjectMeta{Name: "link1", Namespace: "test"},
					Spec: skupperv2alpha1.NetworkLinkSpec{
						Hostname: "remote.example.com",
						Port:     55555,
					},
				},
			},
			expect: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &NetworkLink{
				Name:      "link1",
				Link:      tt.current.link,
				NetworkId: tt.current.networkId,
			}
			if got := m.Update(tt.new.networkId, tt.new.desired); got != tt.expect {
				t.Errorf("NetworkLink.Update() = %v, want %v", got, tt.expect)
			}
		})
	}
}

func TestNetworkLink_Apply(t *testing.T) {
	type fields struct {
		name      string
		networkId string
		link      *skupperv2alpha1.NetworkLink
	}
	tests := []struct {
		name    string
		fields  fields
		preHook func(*qdr.RouterConfig)
		expect  bool
		present bool
	}{
		{
			name: "nil link deletes connector autolink and ssl profile",
			fields: fields{
				name:      "link1",
				networkId: "my-van",
				link:      nil,
			},
			expect:  false,
			present: false,
		},
		{
			name: "empty networkId deletes same as nil link",
			fields: fields{
				name:      "link1",
				networkId: "",
				link: &skupperv2alpha1.NetworkLink{
					ObjectMeta: v1.ObjectMeta{Name: "link1", Namespace: "test"},
					Spec: skupperv2alpha1.NetworkLinkSpec{
						Hostname:       "remote.example.com",
						Port:           55555,
						TlsCredentials: "link1-creds",
					},
				},
			},
			expect:  false,
			present: false,
		},
		{
			name: "nil link deletes existing connector autolink and ssl profile",
			fields: fields{
				name:      "link1",
				networkId: "my-van",
				link:      nil,
			},
			preHook: func(config *qdr.RouterConfig) {
				n := &NetworkLink{
					Name:      "link1",
					NetworkId: "my-van",
					Link: &skupperv2alpha1.NetworkLink{
						ObjectMeta: v1.ObjectMeta{Name: "link1", Namespace: "test"},
						Spec: skupperv2alpha1.NetworkLinkSpec{
							Hostname:       "remote.example.com",
							Port:           55555,
							TlsCredentials: "link1-creds",
						},
					},
				}
				n.Apply(config)
			},
			expect:  true,
			present: false,
		},
		{
			name: "empty networkId deletes existing resources same as nil link",
			fields: fields{
				name:      "link1",
				networkId: "",
				link: &skupperv2alpha1.NetworkLink{
					ObjectMeta: v1.ObjectMeta{Name: "link1", Namespace: "test"},
					Spec: skupperv2alpha1.NetworkLinkSpec{
						Hostname:       "remote.example.com",
						Port:           55555,
						TlsCredentials: "link1-creds",
					},
				},
			},
			preHook: func(config *qdr.RouterConfig) {
				n := &NetworkLink{
					Name:      "link1",
					NetworkId: "my-van",
					Link: &skupperv2alpha1.NetworkLink{
						ObjectMeta: v1.ObjectMeta{Name: "link1", Namespace: "test"},
						Spec: skupperv2alpha1.NetworkLinkSpec{
							Hostname:       "remote.example.com",
							Port:           55555,
							TlsCredentials: "link1-creds",
						},
					},
				}
				n.Apply(config)
			},
			expect:  true,
			present: false,
		},
		{
			name: "valid link with networkId adds ssl profile connector and autolink",
			fields: fields{
				name:      "link1",
				networkId: "my-van",
				link: &skupperv2alpha1.NetworkLink{
					ObjectMeta: v1.ObjectMeta{Name: "link1", Namespace: "test"},
					Spec: skupperv2alpha1.NetworkLinkSpec{
						Hostname:       "remote.example.com",
						Port:           55555,
						TlsCredentials: "link1-creds",
					},
				},
			},
			expect:  true,
			present: true,
		},
		{
			name: "custom TlsCredentials sets ssl profile name from credentials",
			fields: fields{
				name:      "link1",
				networkId: "my-van",
				link: &skupperv2alpha1.NetworkLink{
					ObjectMeta: v1.ObjectMeta{Name: "link1", Namespace: "test"},
					Spec: skupperv2alpha1.NetworkLinkSpec{
						Hostname:       "remote.example.com",
						Port:           55555,
						TlsCredentials: "custom-creds",
					},
				},
			},
			expect:  true,
			present: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := qdr.InitialConfig("r", "s", "v", false, 3)
			if tt.preHook != nil {
				tt.preHook(&config)
			}
			m := &NetworkLink{
				Name:      tt.fields.name,
				Link:      tt.fields.link,
				NetworkId: tt.fields.networkId,
			}
			if got := m.Apply(&config); got != tt.expect {
				t.Errorf("NetworkLink.Apply() = %v, want %v", got, tt.expect)
			}
			if tt.expect {
				connectorName := tt.fields.name + "-networklink-connector"
				if _, ok := config.Connectors[connectorName]; ok != tt.present {
					t.Errorf("expected connector %q to be present", connectorName)
				}
				autoLinkName := tt.fields.name + "-connector-autoLink"
				if _, ok := config.AutoLinks[autoLinkName]; ok != tt.present {
					t.Errorf("expected autolink %q to be present", autoLinkName)
				}
				var expectedProfileName string
				if tt.fields.link != nil && tt.fields.link.Spec.TlsCredentials != "" {
					expectedProfileName = tt.fields.link.Spec.TlsCredentials + "-profile"
				} else {
					expectedProfileName = tt.fields.name + "-profile"
				}
				if _, ok := config.SslProfiles[expectedProfileName]; ok != tt.present {
					t.Errorf("expected ssl profile %q to be present", expectedProfileName)
				}

				got2 := m.Apply(&config)
				if got2 {
					t.Errorf("expected idempotent second Apply to return false")
				}
			}
		})
	}
}
