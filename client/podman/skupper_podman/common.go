package skupper_podman

import "os"

var (
	Namespace               = os.Getenv("USER")
	SkupperContainerVolumes = []string{"skupper-local-server", "router-config", "skupper-site-server", "skupper-router-certs"}
)
