# RKE2 Cluster on Proxmox VE

This example provisions an RKE2 cluster in Rancher using Proxmox VE (PVE) as the infrastructure provider via the `pve` node driver.

It assumes you already have:

- A Rancher instance with the Proxmox VE node driver installed and active
- A PVE cloud credential already created in Rancher (URL, API token and resource pool all live there)
- A Proxmox VE VM template accessible from Rancher (cloud-init capable)

## Prerequisites

- OpenTofu >= 1.5.0 (or Terraform >= 1.5.0)
- Rancher2 provider >= 5.0.0
- PVE node driver: [Lore09/pve-rancher-driver](https://github.com/Lore09/pve-rancher-driver)

## Variables

| Name | Description | Required |
|------|-------------|----------|
| `rancher_url` | Rancher API URL | yes |
| `rancher_credentials` | Object with `access_key` and `secret_key` | yes |
| `rancher_insecure` | Skip TLS verification (self-signed certs) | no (default: `false`) |
| `cluster_name` | Name for the new cluster in Rancher | yes |
| `kubernetes_version` | RKE2 version (e.g. `v1.34.3+rke2r1`) | yes |
| `pve_cloud_credential_name` | Name of the existing PVE cloud credential in Rancher | yes |
| `pve_node` | Single Proxmox VE node name to pin VMs to, e.g. `pve1`. Mutually exclusive with `pve_allowed_nodes` | no (default: empty) |
| `pve_allowed_nodes` | Comma-separated Proxmox VE node names the driver may place VMs on (host selection), e.g. `pve1,pve2`. Mutually exclusive with `pve_node` | no (default: empty = any online node) |
| `pve_net_device` | Network config device whose MAC pins down IP discovery (e.g. `net0`) | no (default: `net0`) |
| `pve_ssh_user` | Account the driver and Rancher log in as; must exist in the template image | yes |
| `pve_linked_clone` | Clone as a linked clone instead of a full clone | no (default: `false`) |
| `pve_tags` | Comma-separated VM tags | no |
| `pve_server_template_vmid` | Proxmox VM template VMID for server nodes | yes |
| `pve_server_sockets` | CPU sockets for server nodes | no (default: `1`) |
| `pve_server_cores` | CPU cores for server nodes | no (default: `2`) |
| `pve_server_memory` | Memory in MiB for server nodes | no (default: `4096`) |
| `server_quantity` | Number of server (control-plane + etcd) nodes | no (default: `1`) |
| `pve_worker_template_vmid` | Proxmox VM template VMID for worker nodes | yes |
| `pve_worker_sockets` | CPU sockets for worker nodes | no (default: `1`) |
| `pve_worker_cores` | CPU cores for worker nodes | no (default: `2`) |
| `pve_worker_memory` | Memory in MiB for worker nodes | no (default: `4096`) |
| `worker_quantity` | Number of worker nodes | no (default: `1`) |

### Node selection

This example demonstrates both node-selection knobs; set at most one of
`pve_node` (pin to a single host) or `pve_allowed_nodes` (restrict to a
subset). Leave both empty to let the driver place VMs on any online node.

### terraform.tfvars example

```hcl
rancher_url               = "https://rancher.example.com"
rancher_insecure          = false
cluster_name              = "pve-cluster"
kubernetes_version        = "v1.34.3+rke2r1"
pve_cloud_credential_name = "my-pve-credential"
pve_allowed_nodes         = "pve1,pve2"
pve_net_device            = "net0"
pve_ssh_user              = "debian"
pve_linked_clone          = false
pve_server_template_vmid  = "100"
pve_server_sockets        = "1"
pve_server_cores          = "2"
pve_server_memory         = "4096"
server_quantity           = 3
pve_worker_template_vmid  = "100"
pve_worker_sockets        = "1"
pve_worker_cores          = "2"
pve_worker_memory         = "4096"
worker_quantity           = 2
```


