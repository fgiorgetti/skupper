package main

import (
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
	fmt.Println(cli)

	var siteHandler v2.SiteHandler
	siteHandler = skupper_podman.NewSitePodmanHandler(cli)

	// Deleting site
	err = siteHandler.Delete()
	if err != nil {
		log.Fatal(err)
	}

}
