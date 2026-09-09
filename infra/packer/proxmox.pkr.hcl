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

source "proxmox-iso" "ftn_debian" {
  proxmox_url              = var.proxmox_endpoint
  username                 = "packer@pve!image"
  token                    = var.proxmox_api_token
  insecure_skip_tls_verify = false

  node                 = var.node
  vm_id                = 9200
  vm_name              = "ftn-debian-golden"
  cores                = 2
  memory               = 4096
  scsi_controller      = "virtio-scsi-pci"
  qemu_agent           = true
  os                   = "l26"

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
    iso_url      = "https://cdimage.debian.org/debian-cd/current/amd64/iso-cd/debian-13.0.0-amd64-netinst.iso"
    iso_checksum = "none"
  }

  ssh_username = "ftn"
  ssh_password = "CHANGE_OUTSIDE_GIT"
  ssh_timeout  = "20m"

  cloud_init = true
}

build {
  sources = ["source.proxmox-iso.ftn_debian"]
}
