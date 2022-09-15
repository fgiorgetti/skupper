package main

import (
	"encoding/json"
	"fmt"
	"log"

	v2 "github.com/skupperproject/skupper/api/types/v2"
	"github.com/skupperproject/skupper/client/podman"
	"github.com/skupperproject/skupper/client/podman/skupper_podman"
)

func main() {
	cli, err := podman.NewPodmanClient("", "")
	if err != nil {
		log.Fatal(err)
	}

	var siteHandler v2.SiteHandler
	siteHandler = skupper_podman.NewSitePodmanHandler(cli)

	site, err := siteHandler.Get()
	if err != nil {
		log.Fatal(err)
	}

	out, _ := json.MarshalIndent(site, "", "    ")
	fmt.Println(string(out))
}
