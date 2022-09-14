package v2

import "github.com/skupperproject/skupper/api/types"

type SkupperComponent interface {
	Name() string
	Image() string
	GetEnv() map[string]string
	GetLabels() map[string]string
	GetSiteIngresses() []SiteIngress
}

type Router struct {
	Env           map[string]string
	Labels        map[string]string
	SiteIngresses []SiteIngress
}

func (r *Router) Name() string {
	return types.TransportDeploymentName
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

func (r *Router) GetLabels() map[string]string {
	if r.Labels == nil {
		r.Labels = map[string]string{}
	}
	return r.Labels
}

func (r *Router) GetSiteIngresses() []SiteIngress {
	return r.SiteIngresses
}
