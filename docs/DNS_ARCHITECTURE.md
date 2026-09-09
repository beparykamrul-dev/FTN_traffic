# FTN DNS Architecture

FTN DNS is designed as a two-plane DNS service:

1. **Recursive resolver plane** — Unbound performs validating, caching recursion using the DNS root hints. This gives FTN an independent recursive resolution service without pretending to be an Internet DNS root server.
2. **Authoritative plane** — an authoritative DNS cluster serves FTN-owned/delegated zones. Public zones are managed from the FTN control plane; private/customer records stay outside Git.

## Root-like behavior

FTN can operate like a root-hints-based recursive resolver: queries are resolved through the DNS hierarchy starting from root-server hints, with DNSSEC validation and caching. FTN is **not** an ICANN/IANA root server; becoming an actual root-server operator is a separate governance and operational role.

## Edge / POP model

```text
Client
  |
  v
FTN DNS VIP / Anycast-ready address
  |
  +--> Authoritative cluster (FTN-owned zones)
  |
  +--> Recursive cluster (Unbound + root hints + DNSSEC)
  |
  +--> Cache / health / policy telemetry
```

Each POP can run a local resolver pair. Anycast/BGP advertisement can be added later through the existing routing plane and remains approval-gated.

## Grafana Cloud

Prometheus scrapes FTN traffic and DNS metrics. Optional `remote_write` sends metrics to Grafana Cloud. The repository stores only environment-variable references; Grafana Cloud credentials must be supplied through the runtime secret store/environment.

Required runtime variables when remote write is enabled:

- `GRAFANA_CLOUD_PROMETHEUS_URL`
- `GRAFANA_CLOUD_USER`
- `GRAFANA_CLOUD_API_KEY`

## Safety / policy defaults

- Public recursive DNS is disabled by policy; explicitly configure authorized customer/FTN networks.
- DNS mutations require approval.
- Credentials and customer DNS records are not committed to Git.
- DNS telemetry contains metrics, not provider credentials.
