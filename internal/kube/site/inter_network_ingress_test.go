package site

import (
	"testing"

	"github.com/skupperproject/skupper/internal/qdr"
	skupperv2alpha1 "github.com/skupperproject/skupper/pkg/apis/skupper/v2alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func linkFixture(name string) *NetworkLink {
	return &NetworkLink{
		Name:      name,
		NetworkId: "my-van",
		Link: &skupperv2alpha1.NetworkLink{
			ObjectMeta: v1.ObjectMeta{Name: name, Namespace: "test"},
			Spec: skupperv2alpha1.NetworkLinkSpec{
				Hostname:       "remote.example.com",
				Port:           55555,
				TlsCredentials: name + "-creds",
			},
		},
	}
}

func accessFixture(name string) *skupperv2alpha1.NetworkAccess {
	return &skupperv2alpha1.NetworkAccess{
		ObjectMeta: v1.ObjectMeta{Name: name, Namespace: "test"},
		Spec: skupperv2alpha1.NetworkAccessSpec{
			BindHost:       "0.0.0.0",
			TlsCredentials: name + "-creds",
			Port:           55556,
		},
	}
}

func ingressFixture(name, routingKey, linkName, accessName string) *skupperv2alpha1.InterNetworkIngress {
	return &skupperv2alpha1.InterNetworkIngress{
		ObjectMeta: v1.ObjectMeta{Name: name, Namespace: "test"},
		Spec: skupperv2alpha1.InterNetworkIngressSpec{
			RoutingKey:    routingKey,
			NetworkLink:   linkName,
			NetworkAccess: accessName,
		},
	}
}

func TestInterNetworkIngress_Update(t *testing.T) {
	link := linkFixture("link1")
	access := accessFixture("na1")
	ingressLink := ingressFixture("ing1", "mykey", "link1", "")
	ingressAccess := ingressFixture("ing1", "mykey", "", "na1")

	tests := []struct {
		name    string
		initial *InterNetworkIngress
		ingress *skupperv2alpha1.InterNetworkIngress
		link    *NetworkLink
		access  *skupperv2alpha1.NetworkAccess
		want    bool
	}{
		{
			name:    "nil ingress returns true",
			initial: NewInterNetworkIngress(ingressLink, link, access),
			ingress: nil,
			link:    nil,
			access:  nil,
			want:    true,
		},
		{
			name:    "routing key change returns true",
			initial: NewInterNetworkIngress(ingressLink, link, nil),
			ingress: ingressFixture("ing1", "otherkey", "link1", ""),
			link:    link,
			access:  nil,
			want:    true,
		},
		{
			name:    "new network link reference returns true",
			initial: NewInterNetworkIngress(ingressFixture("ing1", "mykey", "linkother", ""), link, nil),
			ingress: ingressLink,
			link:    link,
			access:  nil,
			want:    true,
		},
		{
			name:    "new network access reference returns true",
			initial: NewInterNetworkIngress(ingressFixture("ing1", "mykey", "", "naother"), nil, access),
			ingress: ingressAccess,
			link:    nil,
			access:  access,
			want:    true,
		},
		{
			name:    "removing network link reference returns true",
			initial: NewInterNetworkIngress(ingressLink, link, nil),
			ingress: ingressFixture("ing1", "mykey", "", "na1"),
			link:    nil,
			access:  access,
			want:    true,
		},
		{
			name:    "removing network access reference returns true",
			initial: NewInterNetworkIngress(ingressLink, nil, access),
			ingress: ingressFixture("ing1", "mykey", "link1", ""),
			link:    link,
			access:  nil,
			want:    true,
		},
		{
			name:    "identical state returns false",
			initial: NewInterNetworkIngress(ingressLink, link, access),
			ingress: ingressLink,
			link:    link,
			access:  access,
			want:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.initial.Update(tt.ingress, tt.link, tt.access); got != tt.want {
				t.Errorf("InterNetworkIngress.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterNetworkIngress_DesiredAutoLinks(t *testing.T) {
	link := linkFixture("link1")
	access := accessFixture("na1")
	ingress := ingressFixture("ing1", "mykey", "link1", "na1")

	tests := []struct {
		name    string
		subject *InterNetworkIngress
		wantLen int
	}{
		{
			name:    "nil ingress returns empty slice",
			subject: &InterNetworkIngress{Name: "ing1"},
			wantLen: 0,
		},
		{
			name:    "ingress with link only returns one autolink",
			subject: NewInterNetworkIngress(ingress, link, nil),
			wantLen: 1,
		},
		{
			name:    "ingress with access only returns one autolink",
			subject: NewInterNetworkIngress(ingress, nil, access),
			wantLen: 1,
		},
		{
			name:    "ingress with both link and access returns two autolinks",
			subject: NewInterNetworkIngress(ingress, link, access),
			wantLen: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.subject.DesiredAutoLinks()
			if len(got) != tt.wantLen {
				t.Errorf("DesiredAutoLinks() returned %d autolinks, want %d", len(got), tt.wantLen)
			}
			if tt.wantLen >= 1 && tt.subject.Link != nil {
				if got[0].Connection != tt.subject.Link.connectorName() {
					t.Errorf("expected link autolink connection to be %q, got %q", tt.subject.Link.connectorName(), got[0].Connection)
				}
			}
		})
	}
}

func TestInterNetworkIngress_Apply(t *testing.T) {
	link := linkFixture("link1")
	access := accessFixture("na1")
	ingress := ingressFixture("ing1", "mykey", "link1", "na1")

	tests := []struct {
		name         string
		subject      *InterNetworkIngress
		want         bool
		wantLinkAL   bool
		wantAccessAL bool
	}{
		{
			name: "nil ingress removes both autolinks",
			subject: func() *InterNetworkIngress {
				i := NewInterNetworkIngress(ingress, link, access)
				i.Ingress = nil
				return i
			}(),
			want:         false,
			wantLinkAL:   false,
			wantAccessAL: false,
		},
		{
			name:         "ingress with link adds link autolink",
			subject:      NewInterNetworkIngress(ingress, link, nil),
			want:         true,
			wantLinkAL:   true,
			wantAccessAL: false,
		},
		{
			name:         "ingress with access adds access autolink",
			subject:      NewInterNetworkIngress(ingress, nil, access),
			want:         true,
			wantLinkAL:   false,
			wantAccessAL: true,
		},
		{
			name:         "ingress with both adds two autolinks",
			subject:      NewInterNetworkIngress(ingress, link, access),
			want:         true,
			wantLinkAL:   true,
			wantAccessAL: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := qdr.InitialConfig("r", "s", "v", false, 3)
			if got := tt.subject.Apply(&config); got != tt.want {
				t.Errorf("InterNetworkIngress.Apply() = %v, want %v", got, tt.want)
			}
			linkALName := tt.subject.AutoLinkNameForNetworkLink()
			accessALName := tt.subject.AutoLinkNameForNetworkAccess()
			if _, ok := config.AutoLinks[linkALName]; ok != tt.wantLinkAL {
				t.Errorf("autolink %q present = %v, want %v", linkALName, ok, tt.wantLinkAL)
			}
			if _, ok := config.AutoLinks[accessALName]; ok != tt.wantAccessAL {
				t.Errorf("autolink %q present = %v, want %v", accessALName, ok, tt.wantAccessAL)
			}
			if tt.want {
				got2 := tt.subject.Apply(&config)
				if got2 {
					t.Errorf("expected idempotent second Apply to return false")
				}
			}
		})
	}
}
