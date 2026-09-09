# FTN Global Advanced Traffic / Control / Observability Stack

This is the normalized engineering contract for the advanced capabilities requested for FTN_traffic. It separates **traffic acquisition**, **traffic delivery**, **routing**, **control**, and **observability** so that a CDN/cache is never mistaken for Internet transit.

## Data plane
- L3/L4: Linux networking, VLAN/VRF, IPv4/IPv6, ECMP, MTU/path-MTU monitoring.
- Routing: BGP/EBGP/IBGP, route-policy, communities, RPKI validation, graceful restart, BFD integration boundary.
- Edge: Apache Traffic Server-compatible HTTP caching, TLS termination, cache keys, purge, origin shielding, health-based failover.
- Flow: NetFlow/IPFIX/sFlow collector boundary; byte/packet/ASN/POP dimensions.
- Packet telemetry: eBPF/XDP integration boundary for high-rate counters and policy telemetry; no unauthorized interception.

## Control plane
- Provider registry and authorization state.
- Provider-neutral edge and routing adapters.
- DNS authoritative/recursive separation, DNSSEC, DoH/DoT, anycast-ready topology.
- DDNS provider synchronization with explicit approval verification.
- Infrastructure lifecycle through a single authoritative IaC owner per resource set.
- Idempotency, audit events, dry-run/planning, health gates and fail-closed behavior for privileged mutations.

## Observability
- Prometheus metrics and alert rules.
- Grafana-compatible dashboards.
- DNS latency/error/availability telemetry.
- BGP session, prefix, RPKI and route-change telemetry.
- POP/edge cache hit ratio, origin latency, bandwidth and error-rate telemetry.
- Flow top-N, ASN/provider/POP traffic accounting.
- SLO-oriented availability and latency measurements.

## Traffic sources and acquisition
Public CDN/distribution paths may be used for content FTN owns or is authorized to distribute. Internet transit and private third-party traffic require an actual commercial/peering/partner authorization. Public routing/measurement data can inform decisions but does not grant traffic access.

Supported integration classes include CDN contracts, partner edges, IX peering, Internet transit, public CDN endpoints, direct origins, open distribution and public telemetry.

## Security invariants
- Credentials, tokens, private keys and customer private records stay outside Git.
- TLS verification remains enabled.
- Customer traffic interception is disabled by default.
- Third-party private traffic access is authorization-gated.
- Route and edge mutations require approval.
- Management APIs are private/admin-plane only.
- Logs must not contain provider secrets or authorization tokens.

## Production boundary
This document records architecture and integration contracts. Provider-specific SDK clients, BGP daemons, packet collectors, and deployment credentials are separate runtime components and must be enabled only when the corresponding infrastructure and authorization exist.
