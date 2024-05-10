package apis

import (
	"fmt"
	"os"
	"path"

	"github.com/skupperproject/skupper/pkg/apis/skupper/v1alpha1"
	"gopkg.in/yaml.v3"
)

type StaticSiteStateRenderer interface {
	Render(state SiteState) error
}

type SiteState struct {
	Site            v1alpha1.Site
	Listeners       map[string]v1alpha1.Listener
	Connectors      map[string]v1alpha1.Connector
	LinkAccesses    map[string]v1alpha1.LinkAccess
	Grants          map[string]v1alpha1.Grant
	Links           map[string]v1alpha1.Link
	Claims          map[string]v1alpha1.Claim
	Certificates    map[string]v1alpha1.Certificate
	SecuredAccesses map[string]v1alpha1.SecuredAccess
}

func (s *SiteState) IsInterior() bool {
	// TODO Site.Spec.Settings is not working
	// TODO Define how router mode will be defined
	return s.Site.Spec.Settings == nil || s.Site.Spec.Settings["mode"] != "edge"
}

func marshal(outputDirectory, resourceType, resourceName string, resource interface{}) error {
	var err error
	writeDirectory := path.Join(outputDirectory, resourceType)
	err = os.MkdirAll(writeDirectory, 0755)
	if err != nil {
		return fmt.Errorf("error creating directory %s: %w", writeDirectory, err)
	}
	resourceYaml, err := yaml.Marshal(resource)
	if err != nil {
		return fmt.Errorf("error marshalling resource %s: %w", resourceName, err)
	}
	fileName := path.Join(writeDirectory, fmt.Sprintf("%s.yaml", resourceName))
	err = os.WriteFile(fileName, resourceYaml, 0644)
	if err != nil {
		return fmt.Errorf("error writing resource %s: %w", fileName, err)
	}
	return nil
}

func marshalMap[V any](outputDirectory, resourceType string, resourceMap map[string]V) error {
	var err error
	for resourceName, resource := range resourceMap {
		if err = marshal(outputDirectory, resourceType, resourceName, resource); err != nil {
			return err
		}
	}
	return nil
}

func MarshalSiteState(siteState SiteState, outputDirectory string) error {
	var err error
	if err = marshal(outputDirectory, "site", siteState.Site.Name, siteState.Site); err != nil {
		return err
	}
	if err = marshalMap(outputDirectory, "listeners", siteState.Listeners); err != nil {
		return err
	}
	if err = marshalMap(outputDirectory, "connectors", siteState.Connectors); err != nil {
		return err
	}
	if err = marshalMap(outputDirectory, "linkAccesses", siteState.LinkAccesses); err != nil {
		return err
	}
	if err = marshalMap(outputDirectory, "links", siteState.Links); err != nil {
		return err
	}
	if err = marshalMap(outputDirectory, "grants", siteState.Grants); err != nil {
		return err
	}
	if err = marshalMap(outputDirectory, "claims", siteState.Claims); err != nil {
		return err
	}
	if err = marshalMap(outputDirectory, "certificates", siteState.Certificates); err != nil {
		return err
	}
	if err = marshalMap(outputDirectory, "securedAccesses", siteState.SecuredAccesses); err != nil {
		return err
	}
	return nil
}

type SiteStateLoader interface {
	Load() (*SiteState, error)
}

type SiteStateValidator interface {
	Validate(site SiteState) error
}
