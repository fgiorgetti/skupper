#!/bin/bash

# Check if the script is executed with one argument
if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <controller-version>"
    exit 1
fi

APP_VERSION="$1"

# Set chart name and directories
CHART_NAME="skupper-crds"
DEST_DIR="./charts/$CHART_NAME"
CRD_SOURCE_DIR="./config/crd/bases"
MULTI_VAN_SOURCE_DIR="./config/crd/multi-van"
TEMPLATES_DIR="$DEST_DIR/templates"
STABLE_TEMPLATE="$TEMPLATES_DIR/crds.yaml"
MULTI_VAN_TEMPLATE="$TEMPLATES_DIR/multi-van-crds.yaml"

if [ ! -d "$CRD_SOURCE_DIR" ]; then
    echo "Source directory '$CRD_SOURCE_DIR' does not exist. Exiting."
    exit 1
fi

if [ ! -d "$MULTI_VAN_SOURCE_DIR" ]; then
    echo "Source directory '$MULTI_VAN_SOURCE_DIR' does not exist. Exiting."
    exit 1
fi

mkdir -p "$TEMPLATES_DIR"
rm -f "$TEMPLATES_DIR"/*yaml

# Generate templates/crds.yaml — all stable CRDs, conditional on .Values.base
{
    printf '{{- if .Values.base }}\n'
    first=true
    for f in "$CRD_SOURCE_DIR"/*_crd.yaml; do
        if [ "$first" = true ]; then
            first=false
        else
            printf -- '---\n'
        fi
        cat "$f"
    done
    printf '{{- end }}\n'
} > "$STABLE_TEMPLATE"

# Generate templates/multi-van-crds.yaml — multi-van CRDs, conditional on .Values.multiVan
{
    printf '{{- if .Values.multiVan }}\n'
    first=true
    for f in "$MULTI_VAN_SOURCE_DIR"/*_crd.yaml; do
        if [ "$first" = true ]; then
            first=false
        else
            printf -- '---\n'
        fi
        cat "$f"
    done
    printf '{{- end }}\n'
} > "$MULTI_VAN_TEMPLATE"

# Substitute appVersion in Chart.yaml
sed -i "s/^appVersion:.*/appVersion: $APP_VERSION/" "$DEST_DIR/Chart.yaml"

echo "Helm chart '$CHART_NAME' generated successfully with appVersion=$APP_VERSION."
