# 11. Multi-VAN resources

Date: 2026-07-28

## Status

Proposed

## Context

New CRDs have been added to support multi-van management through VMS.
The RouterAccess CRD has also been modified to allow dynamically allocated ports to be persisted into the status field.

These new CRDs must be in place for a given Site of a VAN to be managed through VMS.

## Decision

### New CRDs

1. Network

Defines the networkId to be set on the skupper-router configuration.
It is used as a singleton and only reconciled when named: **network**.

A Site must be in READY state for it to be processed.

We have also considered adding the networkId field as part of the Site CRD spec, or reading it from settings map, but we ended up with a new CRD as the Site resource is usually created and owned by the Site owner, while the Network resource can be provided by VMS (VAN management system). Keeping it separate prevents potential collisions with GitOps managed sites/vans.

2. NetworkAccess

Provides inter-network access to other VANs.
It is handled similarly to the RouterAccess resource, but instead of having spec.roles it has just the port field (which can be dynamically allocated).

The reason why it is a new CRD instead of a new role into the existing RouterAccess, is just for RBAC flexibility.

This resource is only processed when the Network resource is READY.

3. NetworkLink

Counter-part of the Link resource for inter-network links between VANs.
The reason for it to be a new CRD is again just for RBAC flexibility.

This resource is only processed when the Network resource is READY.

4. InterNetworkIngress

Allows a routingKey, available in the local VAN, to be exposed to another VAN through a NetworkLink or a NetworkAccess.

It is only processed when the respective NetworkAccess or NetworkLink resource is READY.

### Update to RouterAccess

The `.status.roles[]` field has been introduced so that the controller can persist the allocated port, in case a `.spec.roles[]` entry is defined without a port (set to 0).

When a role is defined without a port, the controller should allocate a port and store it into the `.status.roles[].port`.

### Make them optional

The new CRDs as well as the change to the RouterAccess CRD must be handled as optional by the Controller.
Therefore, if the new CRDs as well as the modified version of the RouterAccess are not defined, the respective functionality should be considered as disabled and not cause problems to upgraded or new sites.

## Consequences

* Controller will have the ability to handle the new CRDs as well as the updated RouterAccess CRD optionally. If they are not in place, new resources and dynamic port allocation through RouterAccess will be ignored.

* New deployment YAML files (cluster and namespace scoped), including the multi-van CRDs should be produced as part of the release process.

* A new helm-chart to install only the CRDs should be provided. Along with it, a flag to optionally install the multi-van resources.
