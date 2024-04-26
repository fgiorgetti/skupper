package podman

import (
	"github.com/skupperproject/skupper/pkg/apis/skupper/v1alpha1"
	"github.com/skupperproject/skupper/pkg/non_kube/apis"
)

type StaticLibpodRenderer struct {
}

func (l *StaticLibpodRenderer) Render(state apis.SiteState) error {
	//TODO implement me
	panic("implement me")
}

type LibpodSiteState struct {
	site         v1alpha1.Site
	listeners    []v1alpha1.Listener
	connectors   []v1alpha1.Connector
	linkAccesses []v1alpha1.LinkAccess
}

func (l *LibpodSiteState) Site() v1alpha1.Site {
	return l.site
}

func (l *LibpodSiteState) Listeners() []v1alpha1.Listener {
	return l.listeners
}

func (l *LibpodSiteState) Connectors() []v1alpha1.Connector {
	return l.connectors
}

func (l *LibpodSiteState) LinkAccesses() []v1alpha1.LinkAccess {
	return l.linkAccesses
}
