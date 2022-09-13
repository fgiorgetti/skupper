package v2

import "github.com/skupperproject/skupper/api/types"

type SkupperDeployment interface {
	Name() string
	GetComponents() []SkupperComponent
}

type SkupperDeploymentRouter struct {
	Components []SkupperComponent
}

func (s *SkupperDeploymentRouter) Name() string {
	return types.TransportDeploymentName
}

func (s *SkupperDeploymentRouter) GetComponents() []SkupperComponent {
	return s.Components
}
