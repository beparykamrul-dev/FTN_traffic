# FTN Global Advanced DDNS

FTN DDNS is designed as a provider-neutral control layer for authorized zones. It supports IPv4/IPv6, multi-provider synchronization, health-based failover, authenticated agents, webhooks and router/DHCP event sources.

## Flow

`Client/Router/Agent -> Authenticated DDNS API -> Approval/Policy -> DDNS Manager -> Authorized Provider(s) -> DNS mesh`

## Capabilities

- A and AAAA dynamic records
- Dual-stack hostnames
- Multi-provider targets
- Health-based failover and stale-record protection
- Configurable TTL bounds
- Provider-neutral adapters for PowerDNS, Technitium, Cloudflare DNS, DNSPod/Tencent Cloud, Akamai, Route53 and others
- Prometheus metrics and audit integration
- Credential references only; secrets never belong in Git

## Safety model

DDNS updates are state-changing operations. The manager rejects mutations without an explicit approval identifier. Provider adapters must additionally enforce active provider authorization before touching external DNS systems.

No mechanism in this layer attempts to obtain or modify third-party zones without authorization.
