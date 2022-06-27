# Things to keep in mind

* Try to keep model generic and platform-agnostic as much as possible
* CLI has to deal with multiple local sites
* Try to generalize what seems to be k8s specific, like:
  * service target
  * secrets and configmaps (use a resourceRef)
* Avoid defining credentials in the model and try to use a reference to them
* Service sync on non-k8s sites
* How to do a gateway forward on new model
* How a k8s service will be handled on non-k8s sites

# Model validation matrix

## Platforms 

1. Kubernetes
2. Systemd
3. Docker
4. Podman

## Scenarios

1. Platform exclusive mesh (two sites mesh per platform)
2. Hybrid mesh (Four sites mesh with one site on each platform)
3. Define one service per site (frontend or backend)

## Methods

1. Demonstrate CLI using `skupper init`
2. Model should:
   1. Create tokens
   2. Link sites
   3. Expose a service
3. How to achieve the same using CLI

# Procedure and results

## Kubernetes mesh

1. Demonstrate CLI using `skupper init`

The init command can consume the yaml from a local file (creating the configmap)
or the configmap can be created previously, then skupper init will detect its
presence and consume the `skupper.yaml` key, if present.

```bash
$ skupper init --config=<local.yaml>

# or

$ kubectl create configmap skupper-site --from-file=skupper.yaml=<local.yaml>
$ skupper init
```

## Systemd mesh

## Docker mesh

## Podman mesh

## Hybrid mesh
