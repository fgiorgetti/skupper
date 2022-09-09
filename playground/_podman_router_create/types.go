package main

import "github.com/skupperproject/skupper/api/types"

const (
	ContainerNetworkName = "skupper"
)

var (
	SkupperContainerVolumes = []string{"skupper-local-server", "router-config", "skupper-site-server", "skupper-router-certs"}
)

type SiteInitializer interface {
	Prepare() (*SkupperSite, error)
	Initialize(site *SkupperSite) (*SkupperSite, error)
	PostInitialize(site *SkupperSite) (*SkupperSite, error)
}

type HostIngressResolver interface {
	List() ([]HostIngressInfo, error)
	Realise(host HostIngressInfo) error
}

type SkupperSite struct {
	Config *types.SiteConfig
	Spec   *types.RouterSpec
}

type HostIngressInfo struct {
	Host       string
	Ports      map[int]int
	Port       int
	Target     string
	TargetPort int
}
