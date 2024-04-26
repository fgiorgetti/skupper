package apis

import "github.com/skupperproject/skupper/pkg/apis/skupper/v1alpha1"

type StaticSiteStateRenderer interface {
	Render(state SiteState) error
}

type SiteState struct {
	Site         v1alpha1.Site
	Listeners    []v1alpha1.Listener
	Connectors   []v1alpha1.Connector
	LinkAccesses []v1alpha1.LinkAccess
	Grants       []v1alpha1.Grant
}

type SiteStateLoader interface {
	Load() (*SiteState, error)
}
