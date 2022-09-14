package skupper_podman

import "os"

var (
	Namespace               = os.Getenv("USER")
	SkupperContainerVolumes = []string{"skupper-local-server", "skupper-internal", "skupper-site-server", "skupper-router-certs"}
)
