// +build integration cli gateway

package gateway

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"testing"
	"time"

	"github.com/skupperproject/skupper/test/integration/tcp_echo"
	"github.com/skupperproject/skupper/test/utils/base"
	"github.com/skupperproject/skupper/test/utils/k8s"
	"github.com/skupperproject/skupper/test/utils/skupper/cli"
	"github.com/skupperproject/skupper/test/utils/skupper/cli/gateway"
	"github.com/skupperproject/skupper/test/utils/skupper/cli/service"
	"gotest.tools/assert"
	v1 "k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//
// TestLocalGateway uses localhost to run a TCP Echo server
// bound to a dynamic port and expose it through a local
// gateway into the skupper network, against two connected
// clusters. It also forwards local requests to a cluster
// port reaching out tcp-echo-cluster service using a
// dynamic port
//
func TestLocalGateway(t *testing.T) {

	pub, _ := testRunner.GetPublicContext(1)

	var generatedGwName string

	setupScenario := []cli.TestScenario{
		cli.TestScenario{
			Name: "local-gateway-setup",
			Tasks: []cli.SkupperTask{
				{Ctx: pub, Commands: []cli.SkupperCommandTester{
					// skupper gateway init
					&gateway.InitTester{
						GeneratedName: &generatedGwName,
					},
					// skupper service create
					&service.CreateTester{
						Name:    "tcp-echo-host",
						Port:    9090,
						Mapping: "tcp",
					},
					// skupper gateway bind
					&gateway.BindTester{
						Address:  "tcp-echo-host",
						Host:     "0.0.0.0",
						Port:     strconv.Itoa(localTcpEchoPort),
						Protocol: "tcp",
						Name:     generatedGwName,
					},
					// skupper gateway forward
					&gateway.ForwardTester{
						Address: "tcp-echo-cluster",
						Port:    strconv.Itoa(forwardTcpEchoPort),
						Name:    generatedGwName,
					},
				}},
			},
		},
	}

	tearDownScenario := []cli.TestScenario{
		cli.TestScenario{
			Name: "local-gateway-teardown",
			Tasks: []cli.SkupperTask{
				{Ctx: pub, Commands: []cli.SkupperCommandTester{
					// skupper gateway unforward
					&gateway.UnforwardTester{
						Address:  "tcp-echo-cluster",
						Protocol: "tcp",
						Name:     generatedGwName,
					},
					// skupper gateway unbind
					&gateway.UnbindTester{
						Address:  "tcp-echo-host",
						Name:     generatedGwName,
						Protocol: "tcp",
					},
					// skupper service delete
					&service.DeleteTester{
						Name: "tcp-echo-host",
					},
					// skupper gateway delete
					&gateway.DeleteTester{},
				}},
			},
		},
	}

	// Running the setup scenarios
	cli.RunScenarios(t, setupScenario)

	// Testing service communication across gateway and cluster services
	testServices(t)

	// Running the teardown scenarios
	cli.RunScenarios(t, tearDownScenario)

}

//
// TestDowloadedGateway prepares a gateway for download, then
// it installs the prepared gateway in the localhost which is
// used to run a TCP Echo server bound to a dynamic port and expose
// it through a local gateway into the skupper network, against
// two connected clusters. It also forwards local requests to a
// cluster port reaching out tcp-echo-cluster service using a
// dynamic port
//
func TestDownloadedGateway(t *testing.T) {
	pub, _ := testRunner.GetPublicContext(1)
	gwName := "prepared-gateway"

	setupScenario := []cli.TestScenario{
		cli.TestScenario{
			Name: "prepared-gateway-setup",
			Tasks: []cli.SkupperTask{
				{Ctx: pub, Commands: []cli.SkupperCommandTester{
					// skupper gateway init
					&gateway.InitTester{
						Name:         gwName,
						DownloadOnly: true,
					},
					// skupper service create
					&service.CreateTester{
						Name:    "tcp-echo-host",
						Port:    9090,
						Mapping: "tcp",
					},
					// skupper gateway bind
					&gateway.BindTester{
						Address:         "tcp-echo-host",
						Host:            "0.0.0.0",
						Port:            strconv.Itoa(localTcpEchoPort),
						Protocol:        "tcp",
						Name:            gwName,
						IsGatewayActive: true,
					},
					// skupper gateway forward
					&gateway.ForwardTester{
						Address:         "tcp-echo-cluster",
						Port:            strconv.Itoa(forwardTcpEchoPort),
						Name:            gwName,
						IsGatewayActive: true,
					},
					// skupper gateway download
					&gateway.DownloadTester{
						OutputPath: "/tmp",
						Name:       gwName,
					},
				}},
			},
		},
	}

	tearDownScenario := []cli.TestScenario{
		cli.TestScenario{
			Name: "prepared-gateway-teardown",
			Tasks: []cli.SkupperTask{
				{Ctx: pub, Commands: []cli.SkupperCommandTester{
					// skupper gateway delete
					&gateway.DeleteTester{
						Name: gwName,
					},
				}},
			},
		},
	}

	// Running the setup scenarios
	cli.RunScenarios(t, setupScenario)
	if !t.Failed() {
		// Installing the gateway tarball
		installed := t.Run("prepared-gateway-install", func(t *testing.T) {
			assert.Assert(t, installGateway("/tmp/"+gwName+".tar.gz"))
		})

		// Testing service communication across gateway and cluster services
		if installed {
			testServices(t)
		}
	} else {
		t.Logf("skipping further tests due to previous failures...")
	}

	// Running the teardown scenarios
	cli.RunScenarios(t, tearDownScenario)
}

func testServices(t *testing.T) {
	pub, _ := testRunner.GetPublicContext(1)
	prv, _ := testRunner.GetPrivateContext(1)

	runClusterJob := func(cluster *base.ClusterContext, name, address string) error {
		job := k8s.NewJob(name, cluster.Namespace, k8s.JobOpts{
			Image:        k8s.GetTestImage(),
			BackoffLimit: 3,
			Restart:      v1.RestartPolicyOnFailure,
			Env:          map[string]string{"ADDRESS": address},
			Command:      []string{"/app/tcp_echo_test"},
		})
		_, err := cluster.VanClient.KubeClient.BatchV1().Jobs(cluster.Namespace).Create(job)
		if err != nil {
			return err
		}
		_, err = k8s.WaitForJob(cluster.Namespace, cluster.VanClient.KubeClient, name, time.Minute)
		if err != nil {
			cluster.KubectlExec("logs job/" + name)
			return err
		}
		cluster.VanClient.KubeClient.BatchV1().Jobs(cluster.Namespace).Delete(job.Name, &v12.DeleteOptions{})
		return nil
	}

	t.Run("tcp-echo-host", func(t *testing.T) {
		t.Run("client-host", func(t *testing.T) {
			assert.Assert(t, tcp_echo.SendReceive("0.0.0.0:"+strconv.Itoa(localTcpEchoPort)))
		})
		t.Run("client-cluster-public", func(t *testing.T) {
			assert.Assert(t, runClusterJob(pub, "tcp-echo-host", "tcp-echo-host:9090"))
		})
		t.Run("client-cluster-private", func(t *testing.T) {
			assert.Assert(t, runClusterJob(prv, "tcp-echo-host", "tcp-echo-host:9090"))
		})
	})

	t.Run("tcp-echo-cluster", func(t *testing.T) {
		t.Run("client-host", func(t *testing.T) {
			assert.Assert(t, tcp_echo.SendReceive("0.0.0.0:"+strconv.Itoa(forwardTcpEchoPort)))
		})
		t.Run("client-cluster-public", func(t *testing.T) {
			assert.Assert(t, runClusterJob(pub, "tcp-echo-cluster", "tcp-echo-cluster:9090"))
		})
		t.Run("client-cluster-private", func(t *testing.T) {
			assert.Assert(t, runClusterJob(prv, "tcp-echo-cluster", "tcp-echo-cluster:9090"))
		})
	})
}

func installGateway(tarball string) error {
	log.Printf("installing the gateway from tarball %s", tarball)
	gzStream, err := os.Open(tarball)
	if err != nil {
		return err
	}
	tarStream, err := gzip.NewReader(gzStream)
	if err != nil {
		return err
	}
	tarReader := tar.NewReader(tarStream)

	// uncompress all files under temp directory
	dir, err := ioutil.TempDir("", "gateway")
	if err != nil {
		return err
	}

	for {
		h, err := tarReader.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		switch h.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(dir+"/"+h.Name, 0755); err != nil {
				return err
			}
		case tar.TypeReg:
			fileName := dir + "/" + h.Name
			if err := os.MkdirAll(filepath.Dir(fileName), 0755); err != nil {
				return err
			}
			f, err := os.Create(fileName)
			if err != nil {
				return err
			}
			if _, err = io.Copy(f, tarReader); err != nil {
				return err
			}
			if err = f.Close(); err != nil {
				return err
			}
		}
	}

	// running the launch.sh script
	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)
	launch := exec.Command("bash", dir+"/launch.sh")
	launch.Dir = dir
	launch.Stdout = stdout
	launch.Stderr = stderr

	if err = launch.Run(); err != nil {
		log.Println("error executing launch.sh script:")
		log.Printf("stdout:\n%s", stdout.String())
		log.Printf("stderr:\n%s", stderr.String())
		return err
	}

	return nil
}
