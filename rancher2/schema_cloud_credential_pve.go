package rancher2

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func cloudCredentialPveFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"pve_url": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Proxmox VE URL (e.g. 'https://<PROXMOX VE ADDRESS>:8006')",
		},
		"pve_token_id": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Proxmox VE API Token ID (including username and realm, e.g. 'root@pam!rancher')",
		},
		"pve_token_secret": {
			Type:        schema.TypeString,
			Required:    true,
			Sensitive:   true,
			Description: "Proxmox VE API Token secret",
		},
		"pve_insecure_tls": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "Disables Proxmox VE TLS certificate verification",
		},
		"pve_ca_cert": {
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			Description: "PEM CA certificate content (not a path) to trust for the Proxmox VE API",
		},
		"pve_pool": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Proxmox VE resource pool new VMs are created into. Must match the API token's ACL scope",
		},
		"pve_skip_permission_check": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "Skips the token-privilege probe run before provisioning",
		},
	}
}
