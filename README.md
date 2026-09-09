# FTN Traffic

FTN Traffic is the traffic/edge integration plane for Family Time Network.

## Scope

- GitHub-origin public asset distribution
- CDN/edge integration contracts
- Traffic telemetry and health signals
- Transit/IX/BGP traffic kept as a separate network plane
- Provider-neutral adapters
- Approval-first privileged changes
- Proxmox grid/IaC integration
- Secrets stay outside Git

## Architecture

```text
GitHub / FTN Origin -> CDN/Edge -> Global Users

Transit / IX / BGP -> FTN Core -> FTN POPs -> FTN Clients

Proxmox Cluster -> Ceph/Storage -> VM/CT Grid -> FTN Services
                    |
              OpenTofu / Ansible / Packer / Pulumi
```

These planes are integrated for observability and policy, but CDN delivery does not substitute for Internet transit.

## Proxmox Grid / IaC

The original requirements archive is retained in `FTN up.txt`. The normalized implementation contract is in `docs/PROXMOX_GRID_IAC.md`.

- `infra/opentofu/` — Proxmox desired-state provisioning
- `infra/ansible/` — node configuration and observability baseline
- `infra/packer/` — golden-image pipeline
- `infra/pulumi/` — optional infrastructure-as-software integration boundary

Proxmox clusters pool management and storage capacity, but physical RAM is not transparently combined into one VM across nodes. Ceph provides distributed storage; HA/live migration provide workload mobility.

## Safety

Do not commit provider credentials, API tokens, private keys, customer data, or routing secrets. Destructive infrastructure mutations must remain approval-gated.
