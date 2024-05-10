set -Ceu

IMAGE="quay.io/fgiorgetti/bootstrap"
PODMAN_ENDPOINT_DEFAULT="unix://${XDG_RUNTIME_DIR:-/run/user/${UID}}/podman/podman.sock"
if [[ ${UID} = "0" ]]; then
    PODMAN_ENDPOINT_DEFAULT="unix://${XDG_RUNTIME_DIR:-/run}/podman/podman.sock"
fi
PODMAN_ENDPOINT="${PODMAN_ENDPOINT:-${PODMAN_ENDPOINT_DEFAULT}}"
OUTPUT_PATH="${XDG_DATA_HOME:-${HOME}/.local/share}/skupper"

exit_error() {
    echo $*
    exit 1
}

is_sock_endpoint() {
    [[ "${PODMAN_ENDPOINT}" =~ ^(\/|unix:\/\/) ]] && return 0
    return 1
}

main() {
    # Must be mounted into the container
    MOUNTS=()
    ENV_VARS=()
    
    # Mounts
    if is_sock_endpoint; then
        MOUNTS+=(-v ${PODMAN_ENDPOINT/unix:\/\//}:/podman.sock:z)
    fi
    MOUNTS+=(-v ${INPUT_PATH}:/input:z)
    MOUNTS+=(-v ${OUTPUT_PATH}:/output:z)
    
    # Env vars
    if is_sock_endpoint; then
        ENV_VARS+=(-e PODMAN_ENDPOINT="/podman.sock")
    else
        ENV_VARS+=(-e PODMAN_ENDPOINT="${PODMAN_ENDPOINT}")
    fi
    
    # Running the bootstrap
    podman pull ${IMAGE}
    podman run --rm --name skupper-bootstrap \
        --security-opt label=disable -u ${UID} --userns=keep-id \
        --name skupper-podman-bootstrap \
        ${MOUNTS[@]} \
        ${ENV_VARS[@]} \
        ${IMAGE}
}

main $@
