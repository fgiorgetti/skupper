package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/skupperproject/skupper/pkg/non_kube/apis"
	"github.com/skupperproject/skupper/pkg/non_kube/common"
	"github.com/skupperproject/skupper/pkg/non_kube/podman"
	"github.com/skupperproject/skupper/pkg/version"
)

func main() {
	// if -version used, report and exit
	isVersion := flag.Bool("version", false, "Report the version of the Skupper Controller")
	flag.Parse()
	if *isVersion {
		fmt.Println(version.Version)
		os.Exit(0)
	}

	//
	// NOTE FOR CONTAINERS
	// When running bootstrap process through a container
	// the /input path must be mapped to a directory containing a site
	// definition based on CR files.
	// It also expects the /output path to be mapped to the
	// Host's XDG_DATA_HOME/skupper or $HOME/.local/share/skupper
	//

	log.Printf("Skupper V2 - nonkube bootstrap")
	log.Printf("Version: %s", version.Version)

	var inputPath string
	var outputPath string
	inputPath = flag.Arg(0)
	if common.IsRunningInContainer() {
		//
		inputPath = "/input"
		outputPath = "/output"
		for _, directory := range []string{inputPath, outputPath} {
			stat, err := os.Stat(directory)
			if err != nil {
				fmt.Printf("Failed to stat %s: %s\n", directory, err)
				os.Exit(1)
			}
			if !stat.IsDir() {
				fmt.Printf("%s is not a directory\n", directory)
				os.Exit(1)
			}
		}
	}
	// TODO defined standard places for input path?
	if inputPath == "" {
		fmt.Printf("No input path specified\n")
		os.Exit(1)
	}

	siteState, err := bootstrap(inputPath)
	if err != nil {
		fmt.Println("Failed to bootstrap:", err)
		os.Exit(1)
	}
	fmt.Printf("Site %q has been rendered\n", siteState.Site.Name)
}

func bootstrap(inputPath string) (*apis.SiteState, error) {
	var siteStateLoader apis.SiteStateLoader
	siteStateLoader = &common.FileSystemSiteStateLoader{
		Path: inputPath,
	}
	siteState, err := siteStateLoader.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load site state: %v", err)
	}
	var siteStateRenderer = &podman.SiteStateRenderer{}
	err = siteStateRenderer.Render(*siteState)
	if err != nil {
		return nil, fmt.Errorf("failed to render site state: %v", err)
	}
	return siteState, nil
}
