package rancher2

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func machineConfigV2PveFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		// Connection - mandatory if rancher2_cloud_credential.pve_credential_config is not used
		"pve_url": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Proxmox VE URL (e.g. 'https://<PROXMOX VE ADDRESS>:8006')",
		},
		"pve_token_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Proxmox VE API Token ID (including username and realm, e.g. 'root@pam!rancher')",
		},
		"pve_token_secret": {
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			Description: "Proxmox VE API Token secret",
		},
		"pve_insecure_tls": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "Disables Proxmox VE TLS certificate verification",
		},

		// Placement and identity
		"pve_node": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Proxmox VE node to place the VM on. Empty lets the driver pick the online node with the most free memory. Conflicts with pve_allowed_nodes",
		},
		"pve_allowed_nodes": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Comma-separated Proxmox VE node names the driver may place the VM on (e.g. 'pve1,pve2'). Empty considers every online node. Conflicts with pve_node",
		},
		"pve_template_vmid": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "VMID of the Proxmox VE template to clone. Conflicts with pve_template_tag",
		},
		"pve_template_tag": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Select the template by Proxmox VE tag instead of VMID (e.g. 'rancher-node'). Exactly one template must match. Conflicts with pve_template_vmid",
		},
		"pve_template_tag_match": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "subset",
			Description: "How pve_template_tag is matched: 'subset' (template carries at least the given tags) or 'exact' (its tags are exactly the given ones)",
		},
		"pve_linked_clone": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "Clones the template as a linked clone (thin overlay) instead of a full, independent copy. Requires snapshot-capable storage",
		},
		"pve_clone_storage": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Proxmox VE storage ID the clone's disks are created on (e.g. 'ceph-rbd'). Full clones only. Empty uses the template's own storage",
		},
		"pve_clone_format": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Disk format for the clone: 'raw', 'qcow2' or 'vmdk'. Full clones only, and only on file-based storage",
		},
		"pve_vmid": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Explicit VMID for the created VM. '0' or empty auto-assigns. Conflicts with pve_vmid_range",
		},
		"pve_vmid_range": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Allocate the VMID from this inclusive range (e.g. '200-299'). Empty lets Proxmox pick the next free ID cluster-wide",
		},
		"pve_tags": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Comma-separated list of tags to assign to the VM (e.g. 'foo,bar'). Informational only",
		},
		"pve_description": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "VM Notes text. Empty writes a default line naming the machine and its template",
		},
		"pve_vm_name_prefix": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Prefix for the Proxmox VE VM name, rendered as '<prefix>-<machine name>'. Letters, digits and inner hyphens only",
		},
		"pve_onboot": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "Starts the VM automatically when the Proxmox VE host boots",
		},
		"pve_ha": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "Registers the VM as a cluster HA resource. Multi-node clusters only; requires Sys.Console on the API token",
		},
		"pve_ha_group": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "HA group constraining which nodes the resource may run on. Requires pve_ha",
		},

		// Sizing
		"pve_cores": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "2",
			Description: "CPU cores per socket",
		},
		"pve_sockets": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "1",
			Description: "CPU sockets",
		},
		"pve_memory": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "2048",
			Description: "RAM in MiB",
		},

		// Disks
		"pve_boot_disk_size": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "0",
			Description: "Grows the cloned boot disk to this size in GB. '0' keeps the template's own size. Proxmox VE can only grow a disk, never shrink it",
		},
		"pve_boot_disk_device": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "scsi0",
			Description: "Proxmox VE config key of the boot disk grown by pve_boot_disk_size (e.g. 'scsi0', 'virtio0', 'sata0')",
		},
		"pve_backup": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "'true'/'false' to include the boot disk in Proxmox VE backups. Empty keeps whatever the template set",
		},
		"pve_data_disk": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Description: "Data disk to attach; repeatable. Comma-separated key=value pairs, e.g. " +
				"'size=100,storage=local-lvm,fs=ext4,mount=/var/lib/longhorn'. Required keys: size, storage. " +
				"Optional keys: fs (ext4, xfs or none; default ext4), mount (required unless fs=none), label, device, discard, iothread, backup",
		},
		"pve_disk_setup_timeout": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "300",
			Description: "Seconds to wait for SSH plus formatting and mounting of the data disks",
		},

		// Networking
		"pve_net_device": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "net0",
			Description: "Proxmox VE network config device (net0..net31) whose MAC pins down IP discovery, and which the pve_net_* settings are written to",
		},
		"pve_net_bridge": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Proxmox VE bridge to attach the NIC to (e.g. 'vmbr1'). Empty leaves the template's network untouched; setting it rewrites pve_net_device",
		},
		"pve_net_model": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "virtio",
			Description: "Emulated NIC model. Requires pve_net_bridge",
		},
		"pve_net_vlan_tag": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "0",
			Description: "802.1Q VLAN tag ('0' = untagged). Requires pve_net_bridge",
		},
		"pve_net_mtu": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "0",
			Description: "NIC MTU ('0' = Proxmox VE default). Requires pve_net_bridge",
		},
		"pve_net_firewall": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "'true'/'false' to toggle the Proxmox VE firewall on the NIC. Empty keeps the Proxmox VE default. Requires pve_net_bridge",
		},
		"pve_net_iface": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Restricts IP discovery to this guest interface name. Rarely needed, MAC matching already pins it",
		},
		"pve_agent_timeout": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "300",
			Description: "Seconds to wait for the QEMU guest agent to report an IP",
		},
		"pve_cloudinit_timeout": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "300",
			Description: "Seconds to wait for cloud-init to finish inside the guest before handing the machine to Rancher. '0' skips the wait",
		},
		"pve_provision_delay": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "30",
			Description: "Seconds to wait after the VM is up before handing it to Rancher for provisioning",
		},

		// Cloud-init and access
		"pve_ip_mode": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "dhcp",
			Description: "'dhcp' or 'static'. Static derives each machine's address from its VMID, so it requires pve_vmid_range",
		},
		"pve_ip_start": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Static only. First address of the pool, e.g. '192.168.15.150'",
		},
		"pve_ip_end": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Static only. Last address of the pool. Caps how many machines the pool can hold",
		},
		"pve_ip_prefix": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Static only. Subnet prefix length the machines get, e.g. '24'. This is the netmask, not the pool size",
		},
		"pve_gateway": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Static only. Default gateway. May sit outside the pool, but must be inside the subnet pve_ip_prefix describes",
		},
		"pve_nameservers": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "DNS servers, space- or comma-separated. Applies in both IP modes; empty keeps the DHCP-supplied resolver",
		},
		"pve_searchdomain": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "DNS search domain, e.g. 'cluster.lan'. Applies in both IP modes",
		},
		"pve_ciuser": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Cloud-init user to create and install the keys for. Empty derives it from pve_ssh_user",
		},
		"pve_sshkeys": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Extra OpenSSH public keys, one per line. The machine's own generated key is always injected as well",
		},
		"pve_ssh_user": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "root",
			Description: "Account the driver and Rancher log in as. Must exist in the guest image (e.g. 'debian', 'rancher'). The default 'root' works with neither documented template",
		},
		"pve_ssh_port": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "22",
			Description: "Port to use when connecting to the machine via SSH",
		},
		"pve_cicustom": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Proxmox VE cicustom value naming cloud-init snippets, e.g. 'vendor=local:snippets/rancher.yaml'",
		},

		// Additional PVE options
		"pve_extra_config": {
			Type:        schema.TypeList,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "Raw Proxmox VE VM config key=value; repeatable, e.g. 'cpu=host'. Forwarded as-is, not validated",
		},

		// Debugging
		"pve_keep_on_failure": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "Leaves the cloned VM in place when Create fails, instead of rolling it back. Standalone debugging only, leaks VMs when used through Rancher",
		},
	}
}
