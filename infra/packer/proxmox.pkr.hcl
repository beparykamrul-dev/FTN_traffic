packer {
  required_plugins {
    proxmox = {
      version = ">= 1.0.0"
      source  = "github.com/hashicorp/proxmox"
    }
  }
}

variable "proxmox_endpoint" {
  type = string
}

variable "proxmox_api_token" {
  type      = string
  sensitive = true
}

variable "node" {
  type    = string
  default = "pve"
}

variable "iso_url" {
  type        = string
  description = "Pinned installer URL selected by the deployment owner."
}

variable "iso_checksum" {
  type        = string
  description = "Checksum for the pinned installer ISO."
}

variable "ssh_password" {
  type      = string
  sensitive = true
}

source "proxmox-iso" "ftn_linux" {
  proxmox_url              = var.proxmox_endpoint
  username                 = "packer@pve!image"
  token                    = var.proxmox_api_token
  insecure_skip_tls_verify = false

  node            = var.node
  vm_id           = 9200
  vm_name         = "ftn-linux-golden"
  cores           = 2
  memory          = 4096
  scsi_controller = "virtio-scsi-pci"
  qemu_agent      = true
  os              = "l26"

  network_adapters {
    bridge = "vmbr0"
    model  = "virtio"
  }

  disks {
    disk_size         = "20G"
    storage_pool      = "local-lvm"
    storage_pool_type = "lvmthin"
    type              = "scsi"
  }

  boot_iso {
    iso_url      = var.iso_url
    iso_checksum = var.iso_checksum
  }

  ssh_username = "ftn"
  ssh_password = var.ssh_password
  ssh_timeout  = "20m"

  cloud_init = true
}

build {
  sources = ["source.proxmox-iso.ftn_linux"]
}
