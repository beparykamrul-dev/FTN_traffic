output "managed_vm_ids" {
  description = "VM IDs managed by this OpenTofu configuration."
  value       = { for name, vm in proxmox_virtual_environment_vm.ftn_grid : name => vm.vm_id }
}

output "managed_vm_nodes" {
  description = "Node placement for managed VMs."
  value       = { for name, vm in proxmox_virtual_environment_vm.ftn_grid : name => vm.node_name }
}
