package site

type securedAccessMapping struct {
	RouterAccessName string
	Group            string
}

type networkAccessMapping struct {
	NetworkAccessName string
	Group             string
}

func newSecuredAccessMapping(routerAccess, group string) securedAccessMapping {
	return securedAccessMapping{
		RouterAccessName: routerAccess,
		Group:            group,
	}
}

func newNetworkAccessMapping(networkAccess, group string) networkAccessMapping {
	return networkAccessMapping{
		NetworkAccessName: networkAccess,
		Group:             group,
	}
}

type securedAccessMap map[string]securedAccessMapping
type networkAccessMap map[string]networkAccessMapping
