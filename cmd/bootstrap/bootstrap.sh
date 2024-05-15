set -Ceu

IMAGE="quay.io/fgiorgetti/bootstrap"
PODMAN_ENDPOINT_DEFAULT="unix://${XDG_RUNTIME_DIR:-/run/user/${UID}}/podman/podman.sock"
if [[ ${UID} = "0" ]]; then
    PODMAN_ENDPOINT_DEFAULT="unix://${XDG_RUNTIME_DIR:-/run}/podman/podman.sock"
fi
PODMAN_ENDPOINT="${PODMAN_ENDPOINT:-${PODMAN_ENDPOINT_DEFAULT}}"
INPUT_PATH="${1:-}"
OUTPUT_PATH="${XDG_DATA_HOME:-${HOME}/.local/share}/skupper"
CONFIG_HOME="${XDG_CONFIG_HOME:-${HOME}/.config}"
LOG_FILE="$(mktemp /tmp/skupper-bootstrap.XXXXX.log)"

exit_error() {
    echo "$*"
    exit 1
}

is_sock_endpoint() {
    [[ "${PODMAN_ENDPOINT}" =~ ^(\/|unix:\/\/) ]] && return 0
    return 1
}

create_service() {
    # generated service file
    site_name="$(grep -E 'Site ".*" has been created' "${LOG_FILE}" | awk -F'"' '{print $2}')"
    if [[ -z "${site_name}" ]]; then
        echo "Unable to create SystemD service (site name could not be identified)"
        return
    fi
    service_name="skupper-site-${site_name}.service"
    service_file="${OUTPUT_PATH}/sites/${site_name}/runtime/scripts/${service_name}"
    if [[ ! -f ${service_file} ]]; then
        echo "SystemD service has not been defined"
        return
    fi

    # Moving it to the appropriate location
    if [[ ${UID} -eq 0 ]]; then
      mv "${service_file}" /etc/systemd/system/
      systemctl enable --now "${service_name}"
      systemctl daemon-reload
    else
      service_dir="${CONFIG_HOME}/systemd/user/"
      [[ ! -d "${service_dir}" ]] && mkdir -p "${service_dir}"
      mv "${service_file}" "${service_dir}"
      systemctl --user enable --now "${service_name}"
      systemctl --user daemon-reload
    fi
}

main() {
    if [[ -z "${INPUT_PATH}" ]] || [[ ! -d "${INPUT_PATH}" ]]; then
      exit_error "Use: bootstrap.sh <local path to CRs>"
    fi
    # Must be mounted into the container
    MOUNTS=()
    ENV_VARS=()
    
    # Mounts
    if is_sock_endpoint; then
        MOUNTS+=(-v "${PODMAN_ENDPOINT/unix:\/\//}":/podman.sock:z)
    fi
    MOUNTS+=(-v "${INPUT_PATH}":/input:z)
    MOUNTS+=(-v "${OUTPUT_PATH}":/output:z)
    
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
        ${IMAGE} 2>&1 | tee "${LOG_FILE}"

    if [[ $? -eq 0 ]]; then
      create_service
    fi
}

main $@
