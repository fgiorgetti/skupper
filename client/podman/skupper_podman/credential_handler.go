package skupper_podman

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/skupperproject/skupper/api/types"
	"github.com/skupperproject/skupper/client/container"
	"github.com/skupperproject/skupper/client/generated/libpod/client/volumes"
	"github.com/skupperproject/skupper/client/podman"
	"github.com/skupperproject/skupper/pkg/certs"
	"github.com/skupperproject/skupper/pkg/kube"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PodmanCredentialHandler struct {
	cli *podman.PodmanRestClient
}

func NewPodmanCredentialHandler(cli *podman.PodmanRestClient) *PodmanCredentialHandler {
	return &PodmanCredentialHandler{
		cli: cli,
	}
}

func (p *PodmanCredentialHandler) LoadVolumeAsSecret(vol *container.Volume) (*corev1.Secret, error) {
	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      vol.Name,
			Namespace: Namespace,
		},
		Data: map[string][]byte{},
		Type: "kubernetes.io/tls",
	}

	if metadataLabel, ok := vol.Labels[types.InternalMetadataQualifier]; ok {
		err := json.Unmarshal([]byte(metadataLabel), &secret.ObjectMeta)
		if err != nil {
			return nil, fmt.Errorf("error loading secret metadata from volume - %v", err)
		}
	}

	files, err := vol.ListFiles()
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		data, err := ioutil.ReadFile(file.Name())
		if err != nil {
			return nil, fmt.Errorf("error reading file %s for secret %s - %v", file.Name(), vol.Name, err)
		}
		secret.Data[file.Name()] = data
	}

	return secret, nil
}

func (p *PodmanCredentialHandler) SaveSecretAsVolume(secret *corev1.Secret) (*container.Volume, error) {
	vol, err := p.cli.VolumeInspect(secret.Name)

	if err != nil {
		if _, notFound := err.(*volumes.VolumeInspectLibpodNotFound); !notFound {
			return nil, err
		}
		// creating new volume
		metadataStr, err := json.Marshal(secret.ObjectMeta)
		if err != nil {
			return nil, fmt.Errorf("error marshalling secret info for %s - %v", secret.Name, err)
		}
		vol = &container.Volume{
			Name: secret.Name,
			Labels: map[string]string{
				types.InternalMetadataQualifier: string(metadataStr),
			},
		}
		vol, err = p.cli.VolumeCreate(vol)
		if err != nil {
			return nil, fmt.Errorf("error creating volume %s - %v", secret.Name, err)
		}
	}
	_, err = vol.CreateDataFiles(secret.Data, true)
	return nil, err
}

func (p *PodmanCredentialHandler) NewCertAuthority(ca types.CertAuthority) (*corev1.Secret, error) {
	caSecret, err := p.GetSecret(ca.Name)
	if err == nil {
		return caSecret, nil
	}
	if _, notFound := err.(*volumes.VolumeInspectLibpodNotFound); !notFound {
		return nil, fmt.Errorf("Failed to check CA %s : %w", ca.Name, err)
	}
	newCA := certs.GenerateCASecret(ca.Name, ca.Name)
	_, err = p.SaveSecretAsVolume(&newCA)
	return &newCA, err
}

func (p *PodmanCredentialHandler) NewCredential(cred types.Credential) (*corev1.Secret, error) {
	var caSecret *corev1.Secret
	var err error
	if cred.CA != "" {
		caSecret, err = p.GetSecret(cred.CA)
		if err != nil {
			return nil, fmt.Errorf("error loading CA secret %s - %v", cred.CA, err)
		}
	}
	secret := kube.PrepareNewSecret(cred, caSecret, types.TransportDeploymentName)
	_, err = p.SaveSecretAsVolume(&secret)
	return &secret, err
}

func (p *PodmanCredentialHandler) GetSecret(name string) (*corev1.Secret, error) {
	vol, err := p.cli.VolumeInspect(name)
	if err != nil {
		return nil, err
	}
	return p.LoadVolumeAsSecret(vol)
}
