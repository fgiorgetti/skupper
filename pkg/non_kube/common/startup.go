package common

import (
	"bytes"
	_ "embed"
	"os"
	"path"
	"text/template"

	"github.com/skupperproject/skupper/pkg/apis/skupper/v1alpha1"
)

var (
	//go:embed startsh-podman.template
	StartScriptPodmanTemplate string

	//go:embed stopsh-podman.template
	StopScriptPodmanTemplate string
)

type startupScripts struct {
	StartScript string
	StopScript  string
	Site        v1alpha1.Site
	SiteId      string
	path        string
}

func GetStartupScripts(site v1alpha1.Site, siteId string) (*startupScripts, error) {
	scripts := &startupScripts{
		StartScript: StartScriptPodmanTemplate,
		StopScript:  StopScriptPodmanTemplate,
		Site:        site,
		SiteId:      siteId,
	}
	siteHome, err := GetHostSiteHome(site)
	if err != nil {
		return nil, err
	}
	scripts.path = path.Join(siteHome, RuntimeScriptsPath)
	if IsRunningInContainer() {
		scripts.path = path.Join(GetDefaultOutputPath(site.Name), RuntimeScriptsPath)
	}
	return scripts, nil
}

func (s *startupScripts) Create() error {
	var startBuf bytes.Buffer
	var stopBuf bytes.Buffer

	startTemplate := template.Must(template.New("start").Parse(s.StartScript))
	startTemplate.Execute(&startBuf, s)
	startFileName := path.Join(s.path, s.GetStartFileName())
	err := os.WriteFile(startFileName, startBuf.Bytes(), 0755)
	if err != nil {
		return err
	}
	stopTemplate := template.Must(template.New("stop").Parse(s.StopScript))
	stopTemplate.Execute(&stopBuf, s)
	stopFileName := path.Join(s.path, s.GetStopFileName())
	err = os.WriteFile(stopFileName, stopBuf.Bytes(), 0755)
	if err != nil {
		return err
	}
	return nil
}

func (s *startupScripts) Remove() {
	startFileName := path.Join(s.path, s.GetStartFileName())
	stopFileName := path.Join(s.path, s.GetStopFileName())
	_ = os.Remove(startFileName)
	_ = os.Remove(stopFileName)
}

func (s *startupScripts) GetPath() string {
	return s.path
}

func (s *startupScripts) GetStartFileName() string {
	return "start.sh"
}

func (s *startupScripts) GetStopFileName() string {
	return "stop.sh"
}
