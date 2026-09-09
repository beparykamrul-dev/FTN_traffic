# FTN Traffic local dataplane

This batch closes the local orchestration layer around the existing open-source integration contracts.

## Runtime coverage

- Routing: FRRouting, BIRD, GoBGP
- Linux dataplane: iproute2, tc, nftables, eBPF/XDP
- Edge: HAProxy, Envoy, Apache Traffic Server, Varnish
- DNS: Unbound, CoreDNS, PowerDNS Authoritative, dnsdist
- Flow: GoFlow2, pmacct, nfdump
- Observability: Prometheus, OpenTelemetry, Loki
- Transport: WireGuard

The runtime registry is local and backend-neutral. GitHub is not required for service operation.

Route intents are validated for authorization, address family, prefix and next-hop before a local router adapter is called. Capacity headroom and path health provide selection gates. Downstream ISP profiles require authorization, ASN separation, address-family service and a max-prefix limit.

Actual BGP/edge/firewall mutations remain behind the existing approval and authorization boundary. This repository does not claim that an open-source package is installed or that FTN has external routing capacity merely because an adapter exists.
