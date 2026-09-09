# Pulumi integration

Pulumi is retained as an optional infrastructure-as-software layer for FTN. It must consume the same provider-neutral inventory and secret references as OpenTofu/Ansible; it must not become a second source of truth for the same Proxmox resources.

## Rules

- Keep credentials outside Git.
- Use an API token with the minimum required Proxmox permissions.
- Keep destructive operations behind explicit approval in the FTN control plane.
- Do not manage the same VM/storage/network resource concurrently from Pulumi and OpenTofu.
- Prefer the FTN inventory as the authoritative desired-state input.

## Recommended role

Pulumi may orchestrate application-level dependencies or workflows that are awkward in declarative IaC. Proxmox infrastructure state remains owned by one selected IaC engine per resource set.
