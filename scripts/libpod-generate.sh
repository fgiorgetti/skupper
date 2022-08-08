LIBPOD_SPEC='https://storage.googleapis.com/libpod-master-releases/swagger-v3.4.7.yaml'

# Generating libpod clients
./scripts/swagger-generate.sh client/generated/libpod ${LIBPOD_SPEC}

# Model has an issue: https://github.com/containers/podman/issues/13092
sed '/Target string `json:"Target,omitempty"`/a Destination string `json:"Destination,omitempty"`' client/generated/libpod/models/mount.go

