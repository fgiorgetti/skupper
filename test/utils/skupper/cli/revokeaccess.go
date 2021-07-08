package cli

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/skupperproject/skupper/api/types"
	"github.com/skupperproject/skupper/test/utils/base"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"
)

var (
	expectUpdatedSecrets = []string{types.SiteCaSecret, types.SiteServerSecret, types.ClaimsServerSecret}
	updatedSecrets       = []string{}
)

const (
	timeout     = time.Minute
	deleteLabel = "skupper.io/type=token-claim-record"
)

// RevokeAccessTester allows running and validating `skupper revoke-access`.
type RevokeAccessTester struct {
	ExpectClaimRecordsDeleted bool
	secretInformer            cache.SharedIndexInformer
	claimRecordsDeleted       bool
}

func (d *RevokeAccessTester) Command(cluster *base.ClusterContext) []string {
	args := SkupperCommonOptions(cluster)
	args = append(args, "revoke-access")
	return args
}

func (d *RevokeAccessTester) Run(cluster *base.ClusterContext) (stdout string, stderr string, err error) {

	//
	// Creating informers to monitor the following secrets (before revoke-access is issued):
	// - Removal of those labeled as 'skupper.io/type=token-claim-record'
	// - Updates to: skupper-site-ca secret, skupper-site-server, skupper-claims-server
	//
	stopCh := make(chan struct{})
	defer close(stopCh)
	doneCh := d.initializeInformer(cluster, stopCh)

	// Execute revoke-access command
	stdout, stderr, err = RunSkupperCli(d.Command(cluster))
	if err != nil {
		return
	}

	//
	// output is currently empty so we must validate if secrets have been recycled
	//
	log.Printf("Validating 'skupper revoke-access'")
	if stdout != "" {
		err = fmt.Errorf("expected an empty output - found: %s", stdout)
		return
	}

	//
	// Waiting for secret updates to complete or timeout
	//
	log.Printf("validating secrets deleted and updated")
	timeoutCh := time.After(timeout)
	select {
	case <-doneCh:
	case <-timeoutCh:
		err = fmt.Errorf("timed out waiting on secrets to be deleted or updated")
	}

	return
}

func (d *RevokeAccessTester) initializeInformer(cluster *base.ClusterContext, stop <-chan struct{}) chan struct{} {
	done := make(chan struct{})
	validateDone := func() {
		if (!d.ExpectClaimRecordsDeleted || d.claimRecordsDeleted) && reflect.DeepEqual(expectUpdatedSecrets, updatedSecrets) {
			close(done)
		}
	}
	factory := informers.NewSharedInformerFactory(cluster.VanClient.KubeClient, 0)
	d.secretInformer = factory.Core().V1().Secrets().Informer()
	d.secretInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		UpdateFunc: func(oldObj, newObj interface{}) {
			oldSecret := oldObj.(v1.Secret)
			newSecret := newObj.(v1.Secret)
			if !reflect.DeepEqual(oldSecret.Data, newSecret.Data) {
				updatedSecrets = append(updatedSecrets, newSecret.Name)
			}
			validateDone()
		}, DeleteFunc: func(obj interface{}) {
			svc := obj.(v1.Service)
			if svc.ObjectMeta.Labels != nil {
				if _, ok := svc.ObjectMeta.Labels[deleteLabel]; ok {
					d.claimRecordsDeleted = true
					validateDone()
				}
			}
		},
	})
	go d.secretInformer.Run(stop)

	return done
}
