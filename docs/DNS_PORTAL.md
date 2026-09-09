# FTN DNS Portal & Global DNS Mesh

FTN DNS is designed as a control-plane-managed DNS platform, not as a single resolver.

## Engine roles

- PowerDNS Enterprise: authoritative/enterprise DNS integration.
- Technitium DNS: authoritative/recursive DNS option.
- CoreDNS: cloud-native/service-discovery DNS.
- Unbound: recursive resolver using root hints.
- dnsdist: DNS load-balancing, routing and protection layer.
- GoDNS: FTN DNS automation/API integration.
- Hickory DNS: Rust-based DNS stack option.
- Numa DNS: provider/engine integration slot.
- miekg/dns: Go DNS protocol library for FTN services.
- SmartDNS: recursive forwarding/cache option.

## Provider integrations

DNSPod/Tencent Cloud DNS, Cloudflare DNS, Akamai DNS, DuckDNS, Porkbun and Caddy DNS are represented as provider adapters. Let's Encrypt DNS-01 through DNSPod is represented as a certificate automation path.

Provider credentials are references only. Secrets must be supplied through the deployment secret store and never committed to Git.

## Portal

The portal should expose:

1. DNS engines and node health.
2. Authoritative zones and records.
3. Recursive resolver status/cache statistics.
4. dnsdist pools and routing policy.
5. Anycast node health and advertised prefixes.
6. Provider authorization status.
7. DNSSEC/key status.
8. DoH/DoT endpoints.
9. Query/error/latency metrics.
10. Audit and approval queue.

## Global DNS mesh

`Client -> Anycast IP -> nearest healthy FTN DNS node -> dnsdist -> authoritative/recursive engine -> provider/upstream`

The mesh supports active/standby and health-based failover. Anycast/BGP changes remain approval-gated.

## Control plane

ASP.NET Core may serve the administrative portal/API while the Go backend performs DNS orchestration and protocol-level operations. PostgreSQL stores configuration/state/audit metadata and PgBouncer provides connection pooling.

## Safety and authorization

FTN can manage zones and providers that FTN owns or is authorized to administer. The platform does not bypass provider authorization or obtain private third-party DNS data without permission.
