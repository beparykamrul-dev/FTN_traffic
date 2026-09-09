variable "proxmox_endpoint" {
  type        = string
  description = "Proxmox API endpoint. Supply through an ignored tfvars file or environment-backed workflow."
}

variable "proxmox_api_token" {
  type        = string
  sensitive   = true
  description = "Proxmox API token. Never commit this value."
}

variable "proxmox_insecure" {
  type        = bool
  default     = false
  description = "Allow invalid TLS certificates. Keep false in production."
}

variable "proxmox_ssh_username" {
  type    = string
  default = "root"
}

variable "nodes" {
  description = "Explicit VM definitions keyed by FTN node name."
  type = map(object({
    node_name    = string
    name         = string
    vm_id        = number
    cpu_cores    = number
    memory_mb    = number
    disk_gb      = number
    datastore_id = string
    bridge       = string
    qemu_agent   = bool
  }))
  default = {}
}
