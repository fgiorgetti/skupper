package gateway

import (
	"context"
	"fmt"
	"time"

	"github.com/skupperproject/skupper/api/types"
	"github.com/skupperproject/skupper/pkg/utils"
	"github.com/skupperproject/skupper/test/utils/base"
	"github.com/skupperproject/skupper/test/utils/skupper/cli"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// InitTester runs `skupper gateway init` and asserts that
// the gateway is defined accordingly
type InitTester struct {
	Name          string
	DownloadOnly  bool
	GeneratedName *string
}

func (i *InitTester) Command(cluster *base.ClusterContext) []string {
	args := cli.SkupperCommonOptions(cluster)
	args = append(args, "gateway", "init")

	if i.Name != "" {
		args = append(args, "--name", i.Name)
	}
	if i.DownloadOnly {
		args = append(args, "--downloadonly")
	}

	return args
}

func (i *InitTester) Run(cluster *base.ClusterContext) (stdout string, stderr string, err error) {
	//
	// Retrieve existing list of gateways
	//
	ctx := context.Background()
	existingGateways, err := cluster.VanClient.GatewayList(ctx)
	if err != nil {
		return
	}

	// Execute the gateway init command
	stdout, stderr, err = cli.RunSkupperCli(i.Command(cluster))
	if err != nil {
		return
	}

	//
	// Retrieve updated list of gateways
	//
	var currentGateways []*types.GatewayInspectResponse
	err = utils.Retry(time.Second, 5, func() (bool, error) {
		currentGateways, err = cluster.VanClient.GatewayList(ctx)
		if err != nil {
			return false, err
		}
		if len(currentGateways) > len(existingGateways) {
			for _, gw := range currentGateways {
				if gw.GatewayName != "" {
					return true, nil
				}
			}
		}
		return false, nil
	})
	if err != nil {
		return
	}

	// If i.Name is empty we need to discover the generated gateway name
	gatewayName := i.Name
	if gatewayName == "" {
		if len(currentGateways) == 1 {
			gatewayName = currentGateways[0].GatewayName
		} else if len(currentGateways) > len(existingGateways) {
			for _, gw := range currentGateways {
				found := false
				for _, existingGw := range existingGateways {
					if existingGw.GatewayName == gw.GatewayName {
						found = true
					}
				}
				if !found {
					gatewayName = gw.GatewayName
					break
				}
			}
			if gatewayName == "" {
				err = fmt.Errorf("unable to discover gateway name")
				return
			}
		} else {
			err = fmt.Errorf("could not find a new gateway")
			return
		}
	} else {
		found := false
		for _, existingGw := range currentGateways {
			if existingGw.GatewayName == gatewayName {
				found = true
				break
			}
		}
		if !found {
			err = fmt.Errorf("gateway %s not found", gatewayName)
			return
		}
	}

	//
	// Retrieve ConfigMap with skupper.io/type: gateway-definition (label)
	//
	cm, err := cluster.VanClient.KubeClient.CoreV1().ConfigMaps(cluster.Namespace).Get("skupper-gateway-"+gatewayName, v1.GetOptions{})
	if err != nil {
		return
	}

	//
	// Retrieve Secret (token) with same ConfigMap name
	//
	_, err = cluster.VanClient.KubeClient.CoreV1().Secrets(cluster.Namespace).Get(cm.Name, v1.GetOptions{})
	if err != nil {
		return
	}

	// Validating systemd user service created
	expectAvailable := !i.DownloadOnly
	available := SystemdUnitAvailable(gatewayName)
	if available != expectAvailable {
		err = fmt.Errorf("systemd unit %s.service availability issue - available: %v - expected: %v", gatewayName, available, expectAvailable)
		return
	}

	// Validating systemd user service enabled
	enabled := SystemdUnitEnabled(gatewayName)
	if enabled != expectAvailable {
		err = fmt.Errorf("systemd unit %s.service availability issue - enabled: %v - expected: %v", gatewayName, enabled, expectAvailable)
		return
	}

	// Setting Generated Name
	if i.GeneratedName != nil && i.Name == "" {
		*i.GeneratedName = gatewayName
	}

	return
}
