set -o errexit
set -o nounset
set -o pipefail

# Target is a local directory that will be created
TARGET=${1}
# Source can be a local or remote (http) openapi definition
SOURCE=${2}

# Validate if running inside a go module
[[ ! -f go.mod ]] && echo "Run it inside a go project" && exit 1

# Creating the target location
[[ ! -d ${TARGET} ]] && mkdir -p ${TARGET}

# Generating the client code
OWNUSR=`id -u`
OWNGRP=`id -g`
docker run -u ${OWNUSR}:${OWNGRP} --rm -it -v `pwd`:/go/app:Z -w /go/app quay.io/goswagger/swagger generate client -t "${TARGET}" -f "${SOURCE}"
