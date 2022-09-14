package v2

type SkupperDeployment interface {
	GetName() string
	GetComponents() []SkupperComponent
}

type SkupperDeploymentHandler interface {
	Deploy(deployment SkupperDeployment) error
	Undeploy(name string) error
}

type SkupperDeploymentCommon struct {
	Components []SkupperComponent
}

func (s *SkupperDeploymentCommon) GetComponents() []SkupperComponent {
	return s.Components
}
