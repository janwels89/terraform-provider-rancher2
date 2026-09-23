variable "rancher_credentials" {
  description = "Rancher API access and secret key."
  type = object({
    access_key = string
    secret_key = string
  })
  nullable  = false
  sensitive = true
}

variable "rancher_url" {
  description = "Rancher API URL (e.g. https://rancher.example.com)."
  type        = string
}

variable "cluster_name" {
  description = "Name for the new RKE2 cluster in Rancher."
  type        = string
}

variable "kubernetes_version" {
  description = "RKE2 Kubernetes version (e.g. v1.34.3+rke2r1)."
  type        = string
}

variable "rancher_insecure" {
  description = "Skip TLS verification for Rancher (use when CA cert is not in system trust store)."
  type        = bool
  default     = false
}

# Cloud credential name (must already exist in Rancher)
variable "pve_cloud_credential_name" {
  description = "Name of the Proxmox VE cloud credential already configured in Rancher. The resource pool, URL and API token all live on this credential (see rancher2_cloud_credential.pve_credential_config)."
  type        = string
}

# Proxmox VE — machine-level settings only; URL, credentials and resource pool come from the cloud credential
# Node selection: pve_node pins to one host, pve_allowed_nodes restricts to a subset. Set at most one; leave both empty to consider every online node.
variable "pve_node" {
  description = "Single Proxmox VE node name to place VMs on (e.g. 'pve1'). Empty lets the driver pick. Mutually exclusive with pve_allowed_nodes."
  type        = string
  default     = ""
}

variable "pve_allowed_nodes" {
  description = "Comma-separated Proxmox VE node names the driver may place VMs on (e.g. 'pve1,pve2'). Empty considers every online node. Mutually exclusive with pve_node."
  type        = string
  default     = ""
}

variable "pve_net_device" {
  description = "Proxmox VE network config device whose MAC pins down IP discovery (e.g. net0)."
  type        = string
  default     = "net0"
}

variable "pve_ssh_user" {
  description = "Account the driver and Rancher log in as. Must exist in the guest image (e.g. debian, rancher)."
  type        = string
}

variable "pve_linked_clone" {
  description = "Clones the template as a linked clone instead of a full, independent copy. Requires snapshot-capable storage."
  type        = bool
  default     = false
}

variable "pve_tags" {
  description = "Comma-separated list of tags to assign to the VMs."
  type        = string
  default     = ""
}

# Server (control-plane + etcd) pool
variable "pve_server_template_vmid" {
  description = "VMID of the Proxmox VE VM template to use for server nodes."
  type        = string
}

variable "pve_server_sockets" {
  description = "Number of CPU sockets for server nodes."
  type        = string
  default     = "1"
}

variable "pve_server_cores" {
  description = "Number of CPU cores for server nodes."
  type        = string
  default     = "2"
}

variable "pve_server_memory" {
  description = "Memory in MiB for server nodes."
  type        = string
  default     = "4096"
}

variable "server_quantity" {
  description = "Number of server (control-plane + etcd) nodes."
  type        = number
  default     = 1
}

# Worker pool
variable "pve_worker_template_vmid" {
  description = "VMID of the Proxmox VE VM template to use for worker nodes."
  type        = string
}

variable "pve_worker_sockets" {
  description = "Number of CPU sockets for worker nodes."
  type        = string
  default     = "1"
}

variable "pve_worker_cores" {
  description = "Number of CPU cores for worker nodes."
  type        = string
  default     = "2"
}

variable "pve_worker_memory" {
  description = "Memory in MiB for worker nodes."
  type        = string
  default     = "4096"
}

variable "worker_quantity" {
  description = "Number of worker nodes."
  type        = number
  default     = 1
}
