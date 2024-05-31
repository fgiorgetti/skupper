set -Ceu

IMAGE="quay.io/fgiorgetti/bootstrap"
INPUT_PATH="${1:-}"
OUTPUT_PATH="${XDG_DATA_HOME:-${HOME}/.local/share}/skupper"
SERVICE_DIR="${XDG_CONFIG_HOME:-${HOME}/.config}/systemd/user"
LOG_FILE="$(mktemp /tmp/skupper-bootstrap.XXXXX.log)"

exit_error() {
    echo "$*"
    exit 1
}

is_sock_endpoint() {
    [[ "${CONTAINER_ENDPOINT}" =~ ^(\/|unix:\/\/) ]] && return 0
    return 1
}

is_container_platform() {
    [[ "${SKUPPER_PLATFORM}" != "systemd" ]] && return 0
    return 1
}

container_env() {
    export SKUPPER_PLATFORM=${SKUPPER_PLATFORM:-podman}
    export CONTAINER_ENGINE=${CONTAINER_ENGINE:-podman}
    # Determining if the respective binaries are available
    case "${SKUPPER_PLATFORM}" in
        systemd)
            command -v skrouterd > /dev/null || exit_error "SystemD platform cannot be used: skrouterd not found"
            ;;
        docker)
            command -v docker > /dev/null || exit_error "Docker platform cannot be used: docker not found"
            export CONTAINER_ENGINE=docker
            ;;
        *)
            command -v podman > /dev/null || exit_error "Podman platform cannot be used: podman not found"
            export CONTAINER_ENGINE=podman
            ;;
    esac
    export CONTAINER_ENDPOINT_DEFAULT="unix://${XDG_RUNTIME_DIR:-/run/user/${UID}}/podman/podman.sock"
    export RUNAS="${UID}"
    export USERNS="keep-id"
    if [[ "${CONTAINER_ENGINE}" = "docker" ]]; then
        export CONTAINER_ENDPOINT_DEFAULT="unix:///run/docker.sock"
        export USERNS="host"
        export DOCKERGID=$(getent group docker | cut -d: -f3)
        export RUNAS="${UID}:${DOCKERGID}"
    fi
    if [[ ${UID} -eq 0 ]]; then
        if [[ "${CONTAINER_ENGINE}" = "podman" ]]; then
            export CONTAINER_ENDPOINT_DEFAULT="unix:///run/podman/podman.sock"
        fi
        export USERNS=""
        export OUTPUT_PATH="/usr/local/share/skupper"
        mkdir -p "${OUTPUT_PATH}"
        export SERVICE_DIR="/etc/systemd/system"
    fi
    export CONTAINER_ENDPOINT="${CONTAINER_ENDPOINT:-${CONTAINER_ENDPOINT_DEFAULT}}"
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
        if [[ ! -d "${SERVICE_DIR}" ]]; then
            echo "Unable to define path to SystemD service"
            return
        fi
        mv "${service_file}" "${SERVICE_DIR}"
        systemctl --user enable --now "${service_name}"
        systemctl --user daemon-reload
    fi
}

main() {
    if [[ -z "${INPUT_PATH}" ]] || [[ ! -d "${INPUT_PATH}" ]]; then
        exit_error "Use: bootstrap.sh <local path to CRs>"
    fi

    # Parse Skupper Platform and Container Engine settings
    container_env

    # Must be mounted into the container
    MOUNTS=()
    ENV_VARS=()
    
    # Mounts
    if is_sock_endpoint && is_container_platform; then
        MOUNTS+=(-v "${CONTAINER_ENDPOINT/unix:\/\//}":/${CONTAINER_ENGINE}.sock:z)
    fi
    MOUNTS+=(-v "${INPUT_PATH}":/input:z)
    MOUNTS+=(-v "${OUTPUT_PATH}":/output:z)
    
    # Env vars
    if is_container_platform; then
        if is_sock_endpoint; then
            ENV_VARS+=(-e CONTAINER_ENDPOINT="/${CONTAINER_ENGINE}.sock")
        else
            ENV_VARS+=(-e CONTAINER_ENDPOINT="${CONTAINER_ENDPOINT}")
        fi
    fi
    ENV_VARS+=(-e OUTPUT_PATH="${OUTPUT_PATH}")
    ENV_VARS+=(-e SKUPPER_PLATFORM="${SKUPPER_PLATFORM}")

    # Running the bootstrap
    ${CONTAINER_ENGINE} pull ${IMAGE}
    ${CONTAINER_ENGINE} run --rm --name skupper-bootstrap \
        --network host --security-opt label=disable -u "${RUNAS}" --userns="${USERNS}" \
        ${MOUNTS[@]} \
        ${ENV_VARS[@]} \
        ${IMAGE} 2>&1 | tee "${LOG_FILE}"

    if [[ $? -eq 0 ]]; then
        create_service
    fi
}

main $@
