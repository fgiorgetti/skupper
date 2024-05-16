set -Ceu

site=$1
podman rm -f ${site}-skupper-router
rm -rf ~/.local/share/skupper/sites/${site}
service="skupper-site-${site}.service"
systemctl --user stop ${service}
systemctl --user disable ${service}
rm -f ~/.config/systemd/user/${service}
systemctl --user daemon-reload
systemctl --user reset-failed
