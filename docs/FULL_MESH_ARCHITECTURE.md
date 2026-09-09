# FTN Full-Mesh Architecture

## Purpose

FTN uses an authorized full-mesh network plane to connect POPs, core routers, edge nodes, DNS nodes, and other approved infrastructure directly where operationally appropriate.

Full mesh is a transport and routing topology; it does not grant access to third-party networks or private content.

## Planes

1. **Underlay** — IPv4/IPv6 transport, MTU/path-MTU checks, reachability and encrypted node-to-node transport where required.
2. **Overlay** — BGP/ECMP/multipath routing between authenticated FTN nodes.
3. **Edge** — CDN/cache/origin traffic can use healthy mesh paths.
4. **DNS** — authoritative and recursive DNS nodes can use mesh reachability and health-based failover.
5. **Certificate** — certificate distribution can use authenticated mesh channels; private keys remain in the runtime secret/HSM boundary.
6. **Observability** — Prometheus and flow telemetry expose node, link, route and traffic health.

## Routing behavior

- eBGP/iBGP are supported through provider-neutral adapters.
- RPKI validation is required for public-prefix policy.
- BFD and graceful restart provide fast failure detection/recovery.
- ECMP and weighted path selection support parallel links.
- Route mutations are approval-gated and fail closed when authorization is missing.

## Traffic migration

A migration can move traffic between mesh paths, POPs, origins or authorized providers using:

`preflight -> health gate -> weighted/canary shift -> observe -> commit`

If a health gate fails, the controller can return traffic to the last known-good desired state. Every mutation uses an idempotency key and produces an audit record.

## BDIX independence

The full mesh reduces dependence on any single domestic peering path by allowing FTN-owned POPs and authorized peers/transit paths to remain interconnected through multiple routes.

An FTN-operated peering fabric can additionally provide an IX-like service for authorized participants using route-server/BGP, RPKI/IRR policy, BFD, max-prefix controls and participant isolation. This is a separate peering service, not a mechanism for bypassing another exchange or provider.

## Security invariants

- No secrets or private keys in Git.
- No customer traffic interception or payload capture.
- Third-party private access is authorization-gated.
- Administrative APIs remain private/admin-only.
- Destructive routing, DNS, edge and infrastructure mutations require approval.
- Unauthorized route changes are denied.
