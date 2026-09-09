# FTN Traffic open-source runtime boundary

FTN_traffic is the local control/orchestration layer. Open-source networking projects remain replaceable dataplane backends rather than runtime dependencies on GitHub.

## Coverage

- Routing: FRRouting, BIRD, GoBGP
- Linux dataplane: iproute2, tc, nftables, eBPF/XDP
- Edge proxy/cache: HAProxy, Envoy, Apache Traffic Server, Varnish
- DNS: Unbound, CoreDNS, PowerDNS Authoritative, dnsdist
- Flow telemetry: GoFlow2, pmacct, nfdump
- Observability: Prometheus, OpenTelemetry, Loki
- Encrypted transport: WireGuard

The controller selects a backend by capability and health. It does not assume that public routing data grants authorization. Production mutations remain approval-gated and credentials stay outside Git.

FRR is suitable for BGP/OSPF/IS-IS/BFD and Linux routing integration; BIRD is also suitable for large routing and IXP route-server deployments. nftables provides Linux packet filtering/NAT and flowtable facilities. These are integration targets, not a claim that every backend is installed on the current FTN host.
