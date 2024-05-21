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
podman rm -f ${site}-skupper-router
rm -rf ${sites_path}/${site}/
service="skupper-site-${site}.service"
${systemctl} stop ${service}
${systemctl} disable ${service}
rm -f ${service_path}/${service}
${systemctl} daemon-reload
${systemctl} reset-failed
