package v2

import (
	"fmt"

	"github.com/skupperproject/skupper/api/types"
	"github.com/skupperproject/skupper/pkg/utils"
)

type Site interface {
	GetName() string
	GetId() string
	GetMode() string
	GetPlatform() string
	GetCertAuthorities() []types.CertAuthority
	SetCertAuthorities(cas []types.CertAuthority)
	GetCredentials() []types.Credential
	SetCredentials(credentials []types.Credential)
	GetIngressClasses() []string
	GetDeployments() []SkupperDeployment
	SetDeployments(deployments []SkupperDeployment)
}

type SiteHandler interface {
	Create(s Site) error
	Get() (Site, error)
	Delete() error
	Update() error
}

// SiteCommon base implementation of the Site interface
type SiteCommon struct {
	Name            string
	Id              string
	Mode            string
	Platform        string
	CertAuthorities []types.CertAuthority
	Credentials     []types.Credential
	Deployments     []SkupperDeployment
}

func (s *SiteCommon) GetCertAuthorities() []types.CertAuthority {
	return s.CertAuthorities
}

func (s *SiteCommon) SetCertAuthorities(cas []types.CertAuthority) {
	s.CertAuthorities = cas
}

func (s *SiteCommon) GetCredentials() []types.Credential {
	if s.Credentials == nil {
		s.Credentials = []types.Credential{}
	}
	return s.Credentials
}

func (s *SiteCommon) SetCredentials(credentials []types.Credential) {
	s.Credentials = credentials
}

func (s *SiteCommon) GetDeployments() []SkupperDeployment {
	return s.Deployments
}

func (s *SiteCommon) SetDeployments(deployments []SkupperDeployment) {
	s.Deployments = deployments
}

func (s *SiteCommon) GetName() string {
	return s.Name
}

func (s *SiteCommon) GetId() string {
	return s.Id
}

func (s *SiteCommon) GetMode() string {
	return s.Mode
}

func (s *SiteCommon) IsEdge() bool {
	return s.Mode == ModeEdge
}

func (s *SiteCommon) ValidateMinimumRequirements() error {
	reqMsg := func(field string) error {
		return fmt.Errorf("%s cannot be empty", field)
	}
	if s.Name == "" {
		return reqMsg("name")
	}
	if s.Platform == "" {
		return reqMsg("platform")
	}
	if s.Mode == "" {
		return reqMsg("mode")
	}
	if s.Id == "" {
		s.Id = utils.RandomId(10)
	}
	return nil
}
