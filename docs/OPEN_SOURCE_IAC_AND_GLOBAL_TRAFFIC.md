# FTN Open-Source IaC + Selective Global Traffic

FTN uses infrastructure-as-code and traffic-control layers as separate concerns.

## IaC roles

- OpenTofu: primary infrastructure state owner where selected.
- Terraform: compatible alternate; never manages the same state/resource set concurrently with OpenTofu.
- Ansible: host configuration, service deployment, hardening, and runtime orchestration.
- Proxmox: virtualization/cluster layer.
- OpenNebula: optional private-cloud compute/orchestration layer for workloads that fit its cloud model.
- Packer: immutable image preparation.
- Pulumi: optional alternate IaC engine; one engine owns each resource set.

## Observability

Local Prometheus remains the operational source. Grafana Cloud remote_write is an optional external observability sink. Its endpoint/token are runtime secrets, never repository configuration. A remote-write failure must not mutate or disable the dataplane.

## Selective global traffic

FTN does not attempt to obtain third-party/private traffic merely because a route, ASN, DNS record, public endpoint, or open-source implementation is visible. The selector can choose only an authorized source:

1. FTN-owned/direct origin
2. public CDN/distribution
3. contracted CDN
4. authorized partner edge
5. authorized IX peering
6. contracted Internet transit

Selection is capability + authorization + health + latency + packet loss + capacity + locality (+ optional cost). Public routing data is telemetry, not permission.

## Migration

Traffic migration follows `preflight -> authorization -> health gate -> gradual shift -> observe -> commit`, with rollback on failure. ECMP/multipath and BGP policy can provide path redundancy, while DNS/edge migration can move application delivery between authorized origins and POPs.

## Security boundary

No customer payload interception, no provider credential reuse, no route hijacking, no secrets in Git, and no privileged route/edge/DNS/IaC mutation without authorization and approval.
