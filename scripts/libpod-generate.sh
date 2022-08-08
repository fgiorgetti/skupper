LIBPOD_SPEC='https://storage.googleapis.com/libpod-master-releases/swagger-v4.0.0.yaml'

# Generating libpod clients
./scripts/swagger-generate.sh client/generated/libpod ${LIBPOD_SPEC}

# Model has an issue: https://github.com/containers/podman/issues/13092
sed -i '/Target string `json:"Target,omitempty"`/a \\tDestination string `json:"Destination,omitempty"`' client/generated/libpod/models/mount.go
sed -i '/Target string `json:"Target,omitempty"`/a \\tOptions []string `json:"Options,omitempty"`' client/generated/libpod/models/mount.go

