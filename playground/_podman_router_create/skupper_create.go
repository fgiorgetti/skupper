package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/skupperproject/skupper/api/types"
	"github.com/skupperproject/skupper/client"
)

var (
	Namespace = os.Getenv("USER")
)

func CreateSkupperSite() {
	// init params
	params := PodmanInitParams{
		IngressBindHost:            "0.0.0.0",
		IngressBindInterRouterPort: 55671,
		IngressBindEdgePort:        45671,
		ContainerNetworks:          []string{"skupper-network"},
		PodmanEndpoint:             "",
	}

	routerLogging, _ := client.ParseRouterLogConfig("info")

	ingressHost := "192.168.15.10"
	spec := types.SiteConfigSpec{
		SkupperName: Namespace,
		RouterMode:  "interior",
		Ingress:     "host",
		IngressHost: ingressHost,
		Labels:      map[string]string{"application": types.TransportDeploymentName, types.ComponentAnnotation: types.TransportComponentName},
		Router: types.RouterOptions{
			Logging:     routerLogging,
			DebugMode:   "gdb",
			IngressHost: ingressHost,
		},
		Controller: types.ControllerOptions{},
		ConfigSync: types.ConfigSyncOptions{},
		Platform:   types.PlatformPodman,
	}

	// TODO inspect for existing site-config first

	initializer, err := NewPodmanSiteInitializer(spec, params)
	if err != nil {
		log.Fatal(err)
	}

	site, err := initializer.Prepare()
	if err != nil {
		log.Fatal(err)
	}

	site, err = initializer.Initialize(site)
	log.Println("error =", err)
	log.Println("site  =", site)
	siteData, _ := json.MarshalIndent(site, "", "    ")
	fmt.Println(string(siteData))
	if err != nil {
		log.Fatal(err)
	}
}
