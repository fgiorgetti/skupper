set -Ceu

site=$1
sites_path="${HOME}/.local/share/skupper/sites"
service_path="${HOME}/.config/systemd/user"
systemctl="systemctl --user"
if [[ ${UID} -eq 0 ]]; then
    sites_path="/usr/local/share/skupper/sites"
    service_path="/etc/systemd/system"
    systemctl="systemctl"
fi
if [[ ! -d "${sites_path}/${site}" ]]; then
    echo "Site does not exist"
    exit 0
fi
platform_file="${sites_path}/${site}/runtime/state/platform.yaml"
SKUPPER_PLATFORM=$(grep '^platform: ' "${platform_file}" | sed -e 's/.*: //g')
if [[ "${SKUPPER_PLATFORM}" != "systemd" ]]; then
    ${SKUPPER_PLATFORM} rm -f ${site}-skupper-router
fi
rm -rf ${sites_path}/${site}/
service="skupper-site-${site}.service"
${systemctl} stop ${service}
${systemctl} disable ${service}
rm -f ${service_path}/${service}
${systemctl} daemon-reload
${systemctl} reset-failed
