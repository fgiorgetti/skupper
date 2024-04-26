package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/skupperproject/skupper/pkg/version"
)

type Mode string

const (
	dataDir        string = "/home/fgiorget/git/skupper-v2/cmd/controller-podman-v2/crs" //TODO use "/data"
	modeController Mode   = "CONTROLLER"
	modeBootstrap  Mode   = "BOOTSTRAP"
)

func main() {
	// if -version used, report and exit
	isVersion := flag.Bool("version", false, "Report the version of the Skupper Controller")
	flag.Parse()
	if *isVersion {
		fmt.Println(version.Version)
		os.Exit(0)
	}

	mode := os.Getenv("MODE")
	switch Mode(mode) {
	case modeBootstrap:
		bootstrap()
	default:
		log.Println("FORCING bootstrap")
		bootstrap()
		// TODO enable once implemented
		//controller()
	}
}

func bootstrap() {
	log.Printf("Skupper podman V2 bootstrap")
	log.Printf("Version: %s", version.Version)

}

func controller() {
	log.Printf("Skupper podman V2 controller")
	log.Printf("Version: %s", version.Version)
	panic("not implemented")
}
