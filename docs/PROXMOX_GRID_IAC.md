# FTN Proxmox Grid / IaC Architecture

This document carries forward the requirements recorded in `FTN up.txt` and turns them into the FTN_traffic architecture and source tree.

## 1. Network and VLAN layer

- Proxmox bridges such as `vmbr0`/`vmbr1` remain deployment inputs, not hard-coded assumptions.
- VLAN-aware bridges and dedicated management/storage/cluster networks should be selected per site.
- BGP/routing remains a separate authorized network plane; FTN_traffic only integrates telemetry and provider-neutral contracts.

## 2. Storage layer

- Local ZFS/LVM-thin/directory storage can be managed per node.
- Ceph is the distributed-storage option for a real multi-node shared datastore.
- OSDs, MONs and MGRs are cluster resources and must be provisioned only after validating disk layout, network separation, quorum and recovery capacity.
- Automated storage selection is a desired-state function, not an unconditional claim that every disk should be consumed.

## 3. Compute and HA

- Proxmox clustering provides centralized node membership and VM/CT management.
- Live migration can move a VM between compatible nodes.
- HA can restart workloads on another healthy node after a node failure when the cluster has adequate quorum/resources.
- A Proxmox cluster does **not** transparently combine RAM from multiple physical servers into one VM. Memory is consumed by the node running the VM; resource pooling is scheduling/capacity pooling, not a single shared RAM bank.
- A production HA cluster normally needs reliable quorum; use an odd number of voting members or a properly designed QDevice where appropriate.

## 4. Automation / IaC

The repository now retains all IaC families requested in `FTN up.txt`:

- OpenTofu/Terraform-compatible Proxmox provisioning: `infra/opentofu/`
- Ansible configuration/orchestration: `infra/ansible/`
- Packer golden-image pipeline: `infra/packer/`
- Pulumi integration boundary: `infra/pulumi/`

Only one IaC engine should own a given resource set at a time to avoid state contention.

## 5. Monitoring

The intended monitoring layer includes Prometheus/Node Exporter and Grafana, with FTN_traffic health and traffic telemetry feeding the wider NOC. Existing repository monitoring/metrics endpoints remain authoritative for application telemetry.

## 6. Alternative control planes retained from the source requirements

The source requirements also mentioned Proxmox Datacenter Manager, Cockpit, Portainer, Apache CloudStack and OpenNebula. They are treated as optional control/orchestration integrations rather than mandatory replacements for Proxmox VE.

The FTN control plane should prefer a single authoritative lifecycle path for Proxmox resources and use read-only integrations where another platform would otherwise create conflicting state.

## 7. Security and privacy

- No passwords, API tokens, private keys or customer records in Git.
- Proxmox API tokens are sensitive and must be injected by the deployment environment.
- TLS verification stays enabled in production.
- Destructive infrastructure mutations must remain approval-gated in the FTN control plane.
- Management interfaces should be reachable only through the intended private/admin network or an explicitly authorized access gateway.

## 8. Source preservation

`FTN up.txt` remains in the repository as the original requirement archive. This document is the normalized implementation contract; it does not delete the original material.
