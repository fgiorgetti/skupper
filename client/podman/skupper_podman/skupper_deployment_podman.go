package skupper_podman

import (
	v2 "github.com/skupperproject/skupper/api/types/v2"
)

type SkupperDeploymentPodman struct {
	Aliases      []string
	VolumeMounts map[string]string
}

type SkupperDeploymentRouterPodman struct {
	*v2.SkupperDeploymentRouter
	*SkupperDeploymentPodman
}
