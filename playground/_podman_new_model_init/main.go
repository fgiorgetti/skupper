package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/skupperproject/skupper/api/types"
	v2 "github.com/skupperproject/skupper/api/types/v2"
	"github.com/skupperproject/skupper/client/podman"
	"github.com/skupperproject/skupper/client/podman/skupper_podman"
)

func main() {
	cli, err := podman.NewPodmanClient("", "")
	if err != nil {
		log.Fatal(err)
	}
	siteHandler := skupper_podman.NewSitePodmanHandler(cli)
	site := &skupper_podman.SitePodman{
		SiteCommon: &v2.SiteCommon{
			Name:     os.Getenv("USER"),
			Mode:     "interior",
			Platform: types.PlatformPodman,
		},
		IngressBindHost:            "192.168.122.1",
		IngressBindInterRouterPort: int(types.InterRouterListenerPort),
		IngressBindEdgePort:        int(types.EdgeListenerPort),
		ContainerNetworks:          []string{"skupper"},
	}
	_, err = siteHandler.Prepare(site)

	out, _ := json.MarshalIndent(site, "", "    ")
	fmt.Println(string(out))
}
