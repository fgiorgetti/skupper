package v2

import "github.com/skupperproject/skupper/api/types"

type SkupperComponent interface {
	Name() string
	Image() string
	GetEnv() map[string]string
	GetSiteIngresses() []SiteIngress
}

type Router struct {
	Env           map[string]string
	SiteIngresses []SiteIngress
}

func (r *Router) Name() string {
	return types.TransportComponentName
}

func (r *Router) Image() string {
	return GetRouterImageName()
}

func (r *Router) GetEnv() map[string]string {
	if r.Env == nil {
		r.Env = map[string]string{}
	}
	return r.Env
}

func (r *Router) GetSiteIngresses() []SiteIngress {
	return r.SiteIngresses
}
