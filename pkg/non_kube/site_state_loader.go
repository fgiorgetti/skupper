package non_kube

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"github.com/skupperproject/skupper/pkg/apis/skupper/v1alpha1"
	"github.com/skupperproject/skupper/pkg/non_kube/apis"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer/yaml"
	yamlutil "k8s.io/apimachinery/pkg/util/yaml"
)

type FileSystemSiteStateLoader struct {
	Path string
}

func (f *FileSystemSiteStateLoader) Load() (*apis.SiteState, error) {
	var siteState = &apis.SiteState{}
	yamlFileNames, err := f.readAllFiles(f.Path)
	if err != nil {
		return nil, err
	}
	// Reading all yaml files found
	for _, yamlFileName := range yamlFileNames {
		yamlFile, err := os.Open(yamlFileName)
		if err != nil {
			return nil, err
		}
		yamlDecoder := yamlutil.NewYAMLOrJSONDecoder(bufio.NewReader(yamlFile), 1024)
		// read eventual multiple-document yaml
		for {
			var rawObj runtime.RawExtension
			err = yamlDecoder.Decode(&rawObj)
			if err != nil {
				if err != io.EOF {
					return nil, fmt.Errorf("error decoding yaml: %s", err)
				}
				break
			}
			// Decoded object from rawObject, with gvk (Group Version Kind)
			obj, gvk, err := yaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme).Decode(rawObj.Raw, nil, nil)
			if err != nil {
				return nil, err
			}
			// We only care about our v1alpha1 types
			if v1alpha1.SchemeGroupVersion == gvk.GroupVersion() {
				switch gvk.Kind {
				case "Site":
					runtime.DefaultUnstructuredConverter.FromUnstructured(obj.(runtime.Unstructured).UnstructuredContent(), &siteState.Site)
				case "Listener":
					var listener v1alpha1.Listener
					runtime.DefaultUnstructuredConverter.FromUnstructured(obj.(runtime.Unstructured).UnstructuredContent(), &listener)
					siteState.Listeners = append(siteState.Listeners, listener)
				case "Connector":
					var connector v1alpha1.Connector
					runtime.DefaultUnstructuredConverter.FromUnstructured(obj.(runtime.Unstructured).UnstructuredContent(), &connector)
					siteState.Connectors = append(siteState.Connectors, connector)
				case "LinkAccess":
					var linkAccess v1alpha1.LinkAccess
					runtime.DefaultUnstructuredConverter.FromUnstructured(obj.(runtime.Unstructured).UnstructuredContent(), &linkAccess)
					siteState.LinkAccesses = append(siteState.LinkAccesses, linkAccess)
				case "Grant":
					var grant v1alpha1.Grant
					runtime.DefaultUnstructuredConverter.FromUnstructured(obj.(runtime.Unstructured).UnstructuredContent(), &grant)
					siteState.Grants = append(siteState.Grants, grant)
				case "Link":
					var link v1alpha1.Link
					runtime.DefaultUnstructuredConverter.FromUnstructured(obj.(runtime.Unstructured).UnstructuredContent(), &link)
					siteState.Links = append(siteState.Links, link)
				case "Claim":
					var claim v1alpha1.Claim
					runtime.DefaultUnstructuredConverter.FromUnstructured(obj.(runtime.Unstructured).UnstructuredContent(), &claim)
					siteState.Claims = append(siteState.Claims, claim)
				case "Certificate":
					var certificate v1alpha1.Certificate
					runtime.DefaultUnstructuredConverter.FromUnstructured(obj.(runtime.Unstructured).UnstructuredContent(), &certificate)
					siteState.Certificates = append(siteState.Certificates, certificate)
				case "SecuredAccess":
					var securedAccess v1alpha1.SecuredAccess
					runtime.DefaultUnstructuredConverter.FromUnstructured(obj.(runtime.Unstructured).UnstructuredContent(), &securedAccess)
					siteState.SecuredAccesses = append(siteState.SecuredAccesses, securedAccess)
				}
			}
		}
	}
	return siteState, nil
}

func (f *FileSystemSiteStateLoader) readAllFiles(inputDir string) ([]string, error) {
	dir, err := os.Open(inputDir)
	if err != nil {
		return nil, err
	}
	dirInfo, err := dir.Stat()
	if err != nil {
		return nil, err
	}
	if !dirInfo.IsDir() {
		return nil, fmt.Errorf("%s is not a directory", inputDir)
	}
	files, err := dir.ReadDir(0)
	if err != nil {
		return nil, err
	}
	var fileNames []string
	for _, file := range files {
		if file.IsDir() {
			recursiveFiles, err := f.readAllFiles(path.Join(inputDir, file.Name()))
			if err != nil {
				return nil, err
			}
			fileNames = append(fileNames, recursiveFiles...)
		} else {
			if strings.HasSuffix(file.Name(), ".yaml") || strings.HasSuffix(file.Name(), ".yml") {
				fileNames = append(fileNames, path.Join(inputDir, file.Name()))
			}
		}
	}
	return fileNames, nil
}
