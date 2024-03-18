IMAGE="quay.io/fgiorgetti/bootstrap"

PODMAN_ENDPOINT_DEFAULT="unix://${XDG_RUNTIME_DIR:-/run/user/${UID}}/podman/podman.sock"
PODMAN_ENDPOINT="${PODMAN_ENDPOINT:-${PODMAN_ENDPOINT_DEFAULT}}"
VOLUME_PATH="$(podman info --format=json | jq -r .store.volumePath)"

function exit_error() {
    echo $*
    exit 1
}

function is_sock_endpoint() {
    [[ "${PODMAN_ENDPOINT}" =~ ^(\/|unix:\/\/) ]] && return 0
    return 1
}

[[ ! -d ${VOLUME_PATH} ]] && exit_error "Unable to determine podman volume path"

# Must be mounted into the container
VOLUME_MOUNTS=()
if is_sock_endpoint; then
    VOLUME_MOUNTS+=(-v ${PODMAN_ENDPOINT/unix:\/\//}:/tmp/podman.sock:z)
fi
VOLUME_MOUNTS+=(-v ${VOLUME_PATH}:/opt/podman/volumes:z)

# Env vars
ENV_VARS=()
ENV_VARS+=(-e PODMAN_ENDPOINT="/tmp/podman.sock")
ENV_VARS+=(-e HOST_PODMAN_ENDPOINT="${PODMAN_ENDPOINT}")
ENV_VARS+=(-e HOST_USER="${USER}")
ENV_VARS+=(-e HOST_HOSTNAME="${HOSTNAME}")

# Running the bootstrap
podman pull ${IMAGE}
podman run --rm -it --name skupper-bootstrap \
    --security-opt label=disable -u ${UID} --userns=keep-id \
    --name skupper-podman-bootstrap \
    ${VOLUME_MOUNTS[@]} \
    ${ENV_VARS[@]} \
    ${IMAGE}
