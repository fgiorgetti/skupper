package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/skupperproject/skupper/api/types"
	clientpodman "github.com/skupperproject/skupper/client/podman"
	"github.com/skupperproject/skupper/pkg/domain"
	podman "github.com/skupperproject/skupper/pkg/domain/podman"
	"github.com/skupperproject/skupper/pkg/qdr"
)

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	//
	// The goal here is to exercise the bootstrap of a podman
	// site (and eventually a systemd site as well), in order
	// to understand what are all the needs and eventual challenges
	// we may face in the real implementation, which is possibly
	// going to start after we are done with the design of the CRDs.
	//
	// All relevant challenges and points that might need some more
	// thinking will be documented at the notes.md file.
	//
	log.Println("Skupper site bootstrap playground")
	podmanCli, err := clientpodman.NewPodmanClient("", "")
	logFatal(err)

	siteHandler := podman.NewSitePodmanHandlerFromCli(podmanCli)
	exSite, _ := siteHandler.Get()
	if exSite != nil {
		log.Fatal("Skupper is already initialized")
	}
	log.Println("Creating a skupper site using podman")
	defaultName := fmt.Sprintf("%s-%s-%s",
		strings.ToLower(os.Getenv("HOST_HOSTNAME")),
		strings.ToLower(os.Getenv("HOST_USER")),
		uuid.NewString()[:5])

	site := &podman.Site{
		SiteCommon: &domain.SiteCommon{
			Name:     defaultName,
			Mode:     string(qdr.ModeInterior),
			Platform: types.PlatformPodman,
		},
		PodmanEndpoint:      os.Getenv("HOST_" + clientpodman.ENV_PODMAN_ENDPOINT),
		EnableFlowCollector: true,
		EnableConsole:       true,
	}
	ctx, cn := context.WithTimeout(context.Background(), time.Minute*2)
	defer cn()

	err = siteHandler.Create(ctx, site)
	logFatal(err)
}
