# skupper-crds

Helm chart for installing and upgrading [Skupper](https://skupper.io) CRDs.

Unlike the main `skupper` chart — where CRDs are placed in `crds/` and
therefore only installed on the initial `helm install` — this chart places all
CRDs in `templates/` so that `helm upgrade` applies the latest schemas on every
run.

## Prerequisites

- Kubernetes 1.25+
- Helm 3

## Using the chart

The chart exposes two independent flags that control which CRDs are installed:

| Value | Default | Description |
|---|---|---|
| `base` | `true` | Install the stable Skupper CRDs. |
| `multiVan` | `false` | Install the 4 multi-VAN CRDs (Network, NetworkLink, InterNetworkIngress, NetworkAccess). |

### Base CRDs

The stable Skupper CRDs are installed by default:

```bash
helm install skupper-crds oci://quay.io/skupper/helm/skupper-crds
```

To install only the multi-VAN CRDs and skip the base ones (e.g. base CRDs are
already present in the cluster from a previous install):

```bash
helm install skupper-crds oci://quay.io/skupper/helm/skupper-crds \
    --set base=false \
    --set multiVan=true
```

### Multi-VAN CRDs

The chart includes 4 additional CRDs that enable multi-VAN (Virtual Application
Network) support: `Network`, `NetworkLink`, `InterNetworkIngress`, and
`NetworkAccess`. These are **disabled by default** and must be opted in
explicitly:

```bash
helm install skupper-crds oci://quay.io/skupper/helm/skupper-crds \
    --set multiVan=true
```

To install both base and multi-VAN CRDs together using an override `values.yaml`:

```yaml
base: true
multiVan: true
```

> **Note:** Once CRDs are installed, removing them (by running `helm upgrade`
> with a flag set to `false`) will **delete** the CRDs from the cluster, if they
> were installed via this helm chart.
> If the CRDs were installed using a different mechanism, Helm will not delete
> those CRDs during an upgrade or uninstall. If you want to delete them, you have
> to do it manually, with `kubectl delete crd`.

## Upgrading the chart

Because all CRDs are in `templates/`, running `helm upgrade` is sufficient to
apply any schema changes:

```bash
helm upgrade skupper-crds oci://quay.io/skupper/helm/skupper-crds
```

To also upgrade the multi-VAN CRDs:

```bash
helm upgrade skupper-crds oci://quay.io/skupper/helm/skupper-crds \
    --set multiVan=true
```

## Development

The chart has two parts:

- **Static** — `Chart.yaml`, `values.yaml`, `README.md`, and
  `templates/NOTES.txt` are committed to the repository.
- **Generated** — `templates/crds.yaml` (stable CRDs, gated by `base`) and
  `templates/multi-van-crds.yaml` (multi-van CRDs, gated by `multiVan`) are
  produced by the generator script from the source files under `config/crd/`
  and are gitignored. Both template files are applied on every `helm upgrade`
  when their respective flag is `true`.

To regenerate the chart before packaging or linting, run:

```bash
make generate-skupper-crds-helm-chart
```

This generates `templates/crds.yaml` from `config/crd/bases/` and
`templates/multi-van-crds.yaml` from `config/crd/multi-van/`.

To package the chart as a `.tgz`:

```bash
make pack-skupper-crds-helm-chart
```
