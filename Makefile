VERSION := $(shell git describe --tags --dirty=-`git rev-parse --short HEAD`-modified --always)
SERVICE_CONTROLLER_IMAGE ?= quay.io/fgiorgetti/service-controller:multiarch
SITE_CONTROLLER_IMAGE ?= quay.io/fgiorgetti/site-controller:multiarch
CONFIG_SYNC_IMAGE ?= quay.io/fgiorgetti/config-sync:multiarch
FLOW_COLLECTOR_IMAGE ?= quay.io/fgiorgetti/flow-collector:multiarch
TEST_IMAGE ?= quay.io/fgiorgetti/skupper-tests:multiarch
TEST_BINARIES_FOLDER := ${PWD}/test/integration/bin
DOCKER := docker
LDFLAGS := -X github.com/skupperproject/skupper/pkg/version.Version=${VERSION}
GOOS ?= linux
GOARCH ?= amd64

all: build-cmd build-get build-config-sync build-controllers build-tests

build-tests:
	mkdir -p ${TEST_BINARIES_FOLDER}
	GOOS=${GOOS} GOARCH=${GOARCH} go test -c -tags=job -v ./test/integration/examples/tcp_echo/job -o ${TEST_BINARIES_FOLDER}/tcp_echo_test
	GOOS=${GOOS} GOARCH=${GOARCH} go test -c -tags=job -v ./test/integration/examples/http/job -o ${TEST_BINARIES_FOLDER}/http_test
	GOOS=${GOOS} GOARCH=${GOARCH} go test -c -tags=job -v ./test/integration/examples/bookinfo/job -o ${TEST_BINARIES_FOLDER}/bookinfo_test
	GOOS=${GOOS} GOARCH=${GOARCH} go test -c -tags=job -v ./test/integration/examples/mongodb/job -o ${TEST_BINARIES_FOLDER}/mongo_test
	GOOS=${GOOS} GOARCH=${GOARCH} go test -c -tags=job -v ./test/integration/examples/custom/hipstershop/job -o ${TEST_BINARIES_FOLDER}/grpcclient_test
	GOOS=${GOOS} GOARCH=${GOARCH} go test -c -tags=job -v ./test/integration/examples/tls_t/job -o ${TEST_BINARIES_FOLDER}/tls_test

build-cmd:
	GOOS=${GOOS} GOARCH=${GOARCH} go build -ldflags="${LDFLAGS}"  -o skupper ./cmd/skupper

build-get:
	GOOS=${GOOS} GOARCH=${GOARCH} go build -ldflags="${LDFLAGS}"  -o get ./cmd/get

build-service-controller:
	GOOS=${GOOS} GOARCH=${GOARCH} go build -ldflags="${LDFLAGS}"  -o service-controller cmd/service-controller/main.go cmd/service-controller/controller.go cmd/service-controller/ports.go cmd/service-controller/definition_monitor.go cmd/service-controller/console_server.go cmd/service-controller/site_query.go cmd/service-controller/ip_lookup.go cmd/service-controller/claim_verifier.go cmd/service-controller/token_handler.go cmd/service-controller/secret_controller.go cmd/service-controller/claim_handler.go cmd/service-controller/tokens.go cmd/service-controller/links.go cmd/service-controller/services.go cmd/service-controller/policies.go cmd/service-controller/policy_controller.go cmd/service-controller/revoke_access.go  cmd/service-controller/nodes.go

build-site-controller:
	GOOS=${GOOS} GOARCH=${GOARCH} go build -ldflags="${LDFLAGS}"  -o site-controller cmd/site-controller/main.go cmd/site-controller/controller.go

build-flow-collector:
	GOOS=${GOOS} GOARCH=${GOARCH} go build -ldflags="${LDFLAGS}"  -o flow-collector cmd/flow-collector/main.go cmd/flow-collector/controller.go cmd/flow-collector/handlers.go

build-config-sync:
	GOOS=${GOOS} GOARCH=${GOARCH} go build -ldflags="${LDFLAGS}"  -o config-sync cmd/config-sync/main.go cmd/config-sync/config_sync.go

build-controllers: build-site-controller build-service-controller build-flow-collector

docker-build-test-image:
	${DOCKER} buildx build --no-cache --push --platform linux/s390x,linux/amd64 -t ${TEST_IMAGE} -f Dockerfile.ci-test .

docker-build: docker-build-test-image
	${DOCKER} buildx build --no-cache --push --platform linux/s390x,linux/amd64 -t ${SERVICE_CONTROLLER_IMAGE} -f Dockerfile.service-controller .
	${DOCKER} buildx build --no-cache --push --platform linux/s390x,linux/amd64 -t ${SITE_CONTROLLER_IMAGE} -f Dockerfile.site-controller .
	${DOCKER} buildx build --no-cache --push --platform linux/s390x,linux/amd64 -t ${CONFIG_SYNC_IMAGE} -f Dockerfile.config-sync .
	${DOCKER} buildx build --no-cache --push --platform linux/s390x,linux/amd64 -t ${FLOW_COLLECTOR_IMAGE} -f Dockerfile.flow-collector .

docker-build-fc:
	${DOCKER} buildx build --no-cache --push --platform linux/s390x,linux/amd64 -t ${FLOW_COLLECTOR_IMAGE} -f Dockerfile.flow-collector .

format:
	go fmt ./...

generate-client:
	./scripts/update-codegen.sh
	./scripts/libpod-generate.sh

force-generate-client:
	FORCE=true ./scripts/update-codegen.sh
	FORCE=true ./scripts/libpod-generate.sh

client-mock-test:
	go test -v -count=1 ./client

client-cluster-test:
	go test -v -count=1 ./client -use-cluster

vet:
	go vet ./...

cmd-test:
	go test -v -count=1 ./cmd/...

pkg-test:
	go test -v -count=1 ./pkg/...

.PHONY: test
test:
	go test -v -count=1 ./pkg/... ./cmd/... ./client/...

clean:
	rm -rf skupper service-controller site-controller release get config-sync ${TEST_BINARIES_FOLDER}

package: release/windows.zip release/darwin.zip release/linux.tgz release/s390x.tgz release/arm64.tgz

release/linux.tgz: release/linux/skupper
	tar -czf release/linux.tgz -C release/linux/ skupper

release/linux/skupper: cmd/skupper/skupper.go
	GOOS=linux GOARCH=amd64 go build -ldflags="${LDFLAGS}" -o release/linux/skupper ./cmd/skupper

release/windows/skupper: cmd/skupper/skupper.go
	GOOS=windows GOARCH=amd64 go build -ldflags="${LDFLAGS}" -o release/windows/skupper ./cmd/skupper

release/windows.zip: release/windows/skupper
	zip -j release/windows.zip release/windows/skupper

release/darwin/skupper: cmd/skupper/skupper.go
	GOOS=darwin GOARCH=amd64 go build -ldflags="${LDFLAGS}" -o release/darwin/skupper ./cmd/skupper

release/darwin.zip: release/darwin/skupper
	zip -j release/darwin.zip release/darwin/skupper

release/s390x/skupper: cmd/skupper/skupper.go
	GOOS=linux GOARCH=s390x go build -ldflags="${LDFLAGS}" -o release/s390x/skupper ./cmd/skupper

release/s390x.tgz: release/s390x/skupper
	tar -czf release/s390x.tgz release/s390x/skupper

release/arm64/skupper: cmd/skupper/skupper.go
	GOOS=linux GOARCH=arm64 go build -ldflags="${LDFLAGS}" -o release/arm64/skupper ./cmd/skupper

release/arm64.tgz: release/arm64/skupper
	tar -czf release/arm64.tgz release/arm64/skupper
