package client

import (
	"context"
	"fmt"
	"github.com/skupperproject/skupper/api/types"
	"github.com/skupperproject/skupper/pkg/kube"
	"github.com/skupperproject/skupper/pkg/qdr"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/util/retry"
)

func (cli *VanClient) ServiceInterfaceCreate(ctx context.Context, service *types.ServiceInterface) error {
	policy := NewPolicyValidatorAPI(cli)
	res, err := policy.Service(service.Address)
	if err != nil {
		return fmt.Errorf("Unable to validate policies: %s", err)
	}
	if !res.Allowed {
		return fmt.Errorf("Service creation for: %s is not allowed", service.Address)
	}
	owner, err := getRootObject(cli)
	if err == nil {
		err = validateServiceInterface(service)
		if err != nil {
			return err
		}

		if len(service.TlsCredentials) > 0 {

			configmap, err := cli.KubeClient.CoreV1().ConfigMaps(cli.Namespace).Get(types.TransportConfigMapName, metav1.GetOptions{})

			if err != nil {
				return err
			}

			serviceCredential := types.Credential{
				CA:          types.ServiceCaSecret,
				Name:        service.TlsCredentials,
				Subject:     service.Address,
				Hosts:       []string{service.Address},
				ConnectJson: false,
				Post:        false,
			}

			ownerReference := metav1.OwnerReference{
				APIVersion: "v1",
				Kind:       "ConfigMap",
				Name:       configmap.Name,
				UID:        configmap.UID,
			}
			serviceSecret, err := kube.NewSecret(serviceCredential, &ownerReference, cli.Namespace, cli.KubeClient)
			if err != nil {
				return err
			}

			err = qdr.AddSslProfile(serviceSecret.Name, cli.Namespace, cli.KubeClient)
			if err != nil {
				return err
			}

			err = retry.RetryOnConflict(retry.DefaultRetry, func() error {

				err = kube.AppendSecretAndUpdateDeployment(
					serviceSecret.Name,
					"/etc/qpid-dispatch-certs/",
					types.TransportDeploymentName,
					cli.Namespace,
					cli.KubeClient,
					true)
				if err != nil {
					return err
				}
				return nil
			})

			if err != nil {
				return err
			}
		}

		return updateServiceInterface(service, false, owner, cli)
	} else if errors.IsNotFound(err) {
		return fmt.Errorf("Skupper not initialised in %s", cli.Namespace)
	} else {
		return err
	}
}
