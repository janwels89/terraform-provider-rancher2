package rancher2

import (
	norman "github.com/rancher/norman/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	machineConfigV2PveKind       = "PveConfig"
	machineConfigV2PveAPIVersion = "rke-machine-config.cattle.io/v1"
	machineConfigV2PveAPIType    = "rke-machine-config.cattle.io.pveconfig"
)

type machineConfigV2Pve struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Connection
	PveURL         string `json:"url,omitempty" yaml:"url,omitempty"`
	PveTokenID     string `json:"tokenId,omitempty" yaml:"tokenId,omitempty"`
	PveTokenSecret string `json:"tokenSecret,omitempty" yaml:"tokenSecret,omitempty"`
	PveInsecureTLS bool   `json:"insecureTls,omitempty" yaml:"insecureTls,omitempty"`

	// Placement and identity
	PveNode             string `json:"node,omitempty" yaml:"node,omitempty"`
	PveAllowedNodes     string `json:"allowedNodes,omitempty" yaml:"allowedNodes,omitempty"`
	PveTemplateVMID     string `json:"templateVmid,omitempty" yaml:"templateVmid,omitempty"`
	PveTemplateTag      string `json:"templateTag,omitempty" yaml:"templateTag,omitempty"`
	PveTemplateTagMatch string `json:"templateTagMatch,omitempty" yaml:"templateTagMatch,omitempty"`
	PveLinkedClone      bool   `json:"linkedClone,omitempty" yaml:"linkedClone,omitempty"`
	PveCloneStorage     string `json:"cloneStorage,omitempty" yaml:"cloneStorage,omitempty"`
	PveCloneFormat      string `json:"cloneFormat,omitempty" yaml:"cloneFormat,omitempty"`
	PveVMID             string `json:"vmid,omitempty" yaml:"vmid,omitempty"`
	PveVMIDRange        string `json:"vmidRange,omitempty" yaml:"vmidRange,omitempty"`
	PveTags             string `json:"tags,omitempty" yaml:"tags,omitempty"`
	PveDescription      string `json:"description,omitempty" yaml:"description,omitempty"`
	PveVMNamePrefix     string `json:"vmNamePrefix,omitempty" yaml:"vmNamePrefix,omitempty"`
	PveOnboot           bool   `json:"onboot,omitempty" yaml:"onboot,omitempty"`
	PveHA               bool   `json:"ha,omitempty" yaml:"ha,omitempty"`
	PveHAGroup          string `json:"haGroup,omitempty" yaml:"haGroup,omitempty"`

	// Sizing
	PveCores   string `json:"cores,omitempty" yaml:"cores,omitempty"`
	PveSockets string `json:"sockets,omitempty" yaml:"sockets,omitempty"`
	PveMemory  string `json:"memory,omitempty" yaml:"memory,omitempty"`

	// Disks
	PveBootDiskSize     string   `json:"bootDiskSize,omitempty" yaml:"bootDiskSize,omitempty"`
	PveBootDiskDevice   string   `json:"bootDiskDevice,omitempty" yaml:"bootDiskDevice,omitempty"`
	PveBackup           string   `json:"backup,omitempty" yaml:"backup,omitempty"`
	PveDataDisk         []string `json:"dataDisk,omitempty" yaml:"dataDisk,omitempty"`
	PveDiskSetupTimeout string   `json:"diskSetupTimeout,omitempty" yaml:"diskSetupTimeout,omitempty"`

	// Networking
	PveNetDevice        string `json:"netDevice,omitempty" yaml:"netDevice,omitempty"`
	PveNetBridge        string `json:"netBridge,omitempty" yaml:"netBridge,omitempty"`
	PveNetModel         string `json:"netModel,omitempty" yaml:"netModel,omitempty"`
	PveNetVlanTag       string `json:"netVlanTag,omitempty" yaml:"netVlanTag,omitempty"`
	PveNetMtu           string `json:"netMtu,omitempty" yaml:"netMtu,omitempty"`
	PveNetFirewall      string `json:"netFirewall,omitempty" yaml:"netFirewall,omitempty"`
	PveNetIface         string `json:"netIface,omitempty" yaml:"netIface,omitempty"`
	PveAgentTimeout     string `json:"agentTimeout,omitempty" yaml:"agentTimeout,omitempty"`
	PveCloudinitTimeout string `json:"cloudinitTimeout,omitempty" yaml:"cloudinitTimeout,omitempty"`
	PveProvisionDelay   string `json:"provisionDelay,omitempty" yaml:"provisionDelay,omitempty"`

	// Cloud-init and access
	PveIPMode       string `json:"ipMode,omitempty" yaml:"ipMode,omitempty"`
	PveIPStart      string `json:"ipStart,omitempty" yaml:"ipStart,omitempty"`
	PveIPEnd        string `json:"ipEnd,omitempty" yaml:"ipEnd,omitempty"`
	PveIPPrefix     string `json:"ipPrefix,omitempty" yaml:"ipPrefix,omitempty"`
	PveGateway      string `json:"gateway,omitempty" yaml:"gateway,omitempty"`
	PveNameservers  string `json:"nameservers,omitempty" yaml:"nameservers,omitempty"`
	PveSearchdomain string `json:"searchdomain,omitempty" yaml:"searchdomain,omitempty"`
	PveCiuser       string `json:"ciuser,omitempty" yaml:"ciuser,omitempty"`
	PveSSHKeys      string `json:"sshkeys,omitempty" yaml:"sshkeys,omitempty"`
	PveSSHUser      string `json:"sshUser,omitempty" yaml:"sshUser,omitempty"`
	PveSSHPort      string `json:"sshPort,omitempty" yaml:"sshPort,omitempty"`
	PveCicustom     string `json:"cicustom,omitempty" yaml:"cicustom,omitempty"`

	// Additional PVE options
	PveExtraConfig []string `json:"extraConfig,omitempty" yaml:"extraConfig,omitempty"`

	// Debugging
	PveKeepOnFailure bool `json:"keepOnFailure,omitempty" yaml:"keepOnFailure,omitempty"`
}

type MachineConfigV2Pve struct {
	norman.Resource
	machineConfigV2Pve
}

// Flatteners

func flattenMachineConfigV2Pve(in *MachineConfigV2Pve) []interface{} {
	if in == nil {
		return nil
	}
	obj := make(map[string]interface{})

	if len(in.PveURL) > 0 {
		obj["pve_url"] = in.PveURL
	}
	if len(in.PveTokenID) > 0 {
		obj["pve_token_id"] = in.PveTokenID
	}
	if len(in.PveTokenSecret) > 0 {
		obj["pve_token_secret"] = in.PveTokenSecret
	}
	obj["pve_insecure_tls"] = in.PveInsecureTLS

	if len(in.PveNode) > 0 {
		obj["pve_node"] = in.PveNode
	}
	if len(in.PveAllowedNodes) > 0 {
		obj["pve_allowed_nodes"] = in.PveAllowedNodes
	}
	if len(in.PveTemplateVMID) > 0 {
		obj["pve_template_vmid"] = in.PveTemplateVMID
	}
	if len(in.PveTemplateTag) > 0 {
		obj["pve_template_tag"] = in.PveTemplateTag
	}
	if len(in.PveTemplateTagMatch) > 0 {
		obj["pve_template_tag_match"] = in.PveTemplateTagMatch
	}
	obj["pve_linked_clone"] = in.PveLinkedClone
	if len(in.PveCloneStorage) > 0 {
		obj["pve_clone_storage"] = in.PveCloneStorage
	}
	if len(in.PveCloneFormat) > 0 {
		obj["pve_clone_format"] = in.PveCloneFormat
	}
	if len(in.PveVMID) > 0 {
		obj["pve_vmid"] = in.PveVMID
	}
	if len(in.PveVMIDRange) > 0 {
		obj["pve_vmid_range"] = in.PveVMIDRange
	}
	if len(in.PveTags) > 0 {
		obj["pve_tags"] = in.PveTags
	}
	if len(in.PveDescription) > 0 {
		obj["pve_description"] = in.PveDescription
	}
	if len(in.PveVMNamePrefix) > 0 {
		obj["pve_vm_name_prefix"] = in.PveVMNamePrefix
	}
	obj["pve_onboot"] = in.PveOnboot
	obj["pve_ha"] = in.PveHA
	if len(in.PveHAGroup) > 0 {
		obj["pve_ha_group"] = in.PveHAGroup
	}

	if len(in.PveCores) > 0 {
		obj["pve_cores"] = in.PveCores
	}
	if len(in.PveSockets) > 0 {
		obj["pve_sockets"] = in.PveSockets
	}
	if len(in.PveMemory) > 0 {
		obj["pve_memory"] = in.PveMemory
	}

	if len(in.PveBootDiskSize) > 0 {
		obj["pve_boot_disk_size"] = in.PveBootDiskSize
	}
	if len(in.PveBootDiskDevice) > 0 {
		obj["pve_boot_disk_device"] = in.PveBootDiskDevice
	}
	if len(in.PveBackup) > 0 {
		obj["pve_backup"] = in.PveBackup
	}
	if len(in.PveDataDisk) > 0 {
		obj["pve_data_disk"] = toArrayInterface(in.PveDataDisk)
	}
	if len(in.PveDiskSetupTimeout) > 0 {
		obj["pve_disk_setup_timeout"] = in.PveDiskSetupTimeout
	}

	if len(in.PveNetDevice) > 0 {
		obj["pve_net_device"] = in.PveNetDevice
	}
	if len(in.PveNetBridge) > 0 {
		obj["pve_net_bridge"] = in.PveNetBridge
	}
	if len(in.PveNetModel) > 0 {
		obj["pve_net_model"] = in.PveNetModel
	}
	if len(in.PveNetVlanTag) > 0 {
		obj["pve_net_vlan_tag"] = in.PveNetVlanTag
	}
	if len(in.PveNetMtu) > 0 {
		obj["pve_net_mtu"] = in.PveNetMtu
	}
	if len(in.PveNetFirewall) > 0 {
		obj["pve_net_firewall"] = in.PveNetFirewall
	}
	if len(in.PveNetIface) > 0 {
		obj["pve_net_iface"] = in.PveNetIface
	}
	if len(in.PveAgentTimeout) > 0 {
		obj["pve_agent_timeout"] = in.PveAgentTimeout
	}
	if len(in.PveCloudinitTimeout) > 0 {
		obj["pve_cloudinit_timeout"] = in.PveCloudinitTimeout
	}
	if len(in.PveProvisionDelay) > 0 {
		obj["pve_provision_delay"] = in.PveProvisionDelay
	}

	if len(in.PveIPMode) > 0 {
		obj["pve_ip_mode"] = in.PveIPMode
	}
	if len(in.PveIPStart) > 0 {
		obj["pve_ip_start"] = in.PveIPStart
	}
	if len(in.PveIPEnd) > 0 {
		obj["pve_ip_end"] = in.PveIPEnd
	}
	if len(in.PveIPPrefix) > 0 {
		obj["pve_ip_prefix"] = in.PveIPPrefix
	}
	if len(in.PveGateway) > 0 {
		obj["pve_gateway"] = in.PveGateway
	}
	if len(in.PveNameservers) > 0 {
		obj["pve_nameservers"] = in.PveNameservers
	}
	if len(in.PveSearchdomain) > 0 {
		obj["pve_searchdomain"] = in.PveSearchdomain
	}
	if len(in.PveCiuser) > 0 {
		obj["pve_ciuser"] = in.PveCiuser
	}
	if len(in.PveSSHKeys) > 0 {
		obj["pve_sshkeys"] = in.PveSSHKeys
	}
	if len(in.PveSSHUser) > 0 {
		obj["pve_ssh_user"] = in.PveSSHUser
	}
	if len(in.PveSSHPort) > 0 {
		obj["pve_ssh_port"] = in.PveSSHPort
	}
	if len(in.PveCicustom) > 0 {
		obj["pve_cicustom"] = in.PveCicustom
	}

	if len(in.PveExtraConfig) > 0 {
		obj["pve_extra_config"] = toArrayInterface(in.PveExtraConfig)
	}

	obj["pve_keep_on_failure"] = in.PveKeepOnFailure

	return []interface{}{obj}
}

// Expanders

func expandMachineConfigV2Pve(p []interface{}, source *MachineConfigV2) *MachineConfigV2Pve {
	if len(p) == 0 || p[0] == nil {
		return nil
	}
	obj := &MachineConfigV2Pve{}
	if len(source.ID) > 0 {
		obj.ID = source.ID
	}
	in := p[0].(map[string]interface{})

	obj.TypeMeta.Kind = machineConfigV2PveKind
	obj.TypeMeta.APIVersion = machineConfigV2PveAPIVersion
	source.TypeMeta = obj.TypeMeta
	obj.ObjectMeta = source.ObjectMeta

	if v, ok := in["pve_url"].(string); ok && len(v) > 0 {
		obj.PveURL = v
	}
	if v, ok := in["pve_token_id"].(string); ok && len(v) > 0 {
		obj.PveTokenID = v
	}
	if v, ok := in["pve_token_secret"].(string); ok && len(v) > 0 {
		obj.PveTokenSecret = v
	}
	if v, ok := in["pve_insecure_tls"].(bool); ok {
		obj.PveInsecureTLS = v
	}

	if v, ok := in["pve_node"].(string); ok && len(v) > 0 {
		obj.PveNode = v
	}
	if v, ok := in["pve_allowed_nodes"].(string); ok && len(v) > 0 {
		obj.PveAllowedNodes = v
	}
	if v, ok := in["pve_template_vmid"].(string); ok && len(v) > 0 {
		obj.PveTemplateVMID = v
	}
	if v, ok := in["pve_template_tag"].(string); ok && len(v) > 0 {
		obj.PveTemplateTag = v
	}
	if v, ok := in["pve_template_tag_match"].(string); ok && len(v) > 0 {
		obj.PveTemplateTagMatch = v
	}
	if v, ok := in["pve_linked_clone"].(bool); ok {
		obj.PveLinkedClone = v
	}
	if v, ok := in["pve_clone_storage"].(string); ok && len(v) > 0 {
		obj.PveCloneStorage = v
	}
	if v, ok := in["pve_clone_format"].(string); ok && len(v) > 0 {
		obj.PveCloneFormat = v
	}
	if v, ok := in["pve_vmid"].(string); ok && len(v) > 0 {
		obj.PveVMID = v
	}
	if v, ok := in["pve_vmid_range"].(string); ok && len(v) > 0 {
		obj.PveVMIDRange = v
	}
	if v, ok := in["pve_tags"].(string); ok && len(v) > 0 {
		obj.PveTags = v
	}
	if v, ok := in["pve_description"].(string); ok && len(v) > 0 {
		obj.PveDescription = v
	}
	if v, ok := in["pve_vm_name_prefix"].(string); ok && len(v) > 0 {
		obj.PveVMNamePrefix = v
	}
	if v, ok := in["pve_onboot"].(bool); ok {
		obj.PveOnboot = v
	}
	if v, ok := in["pve_ha"].(bool); ok {
		obj.PveHA = v
	}
	if v, ok := in["pve_ha_group"].(string); ok && len(v) > 0 {
		obj.PveHAGroup = v
	}

	if v, ok := in["pve_cores"].(string); ok && len(v) > 0 {
		obj.PveCores = v
	}
	if v, ok := in["pve_sockets"].(string); ok && len(v) > 0 {
		obj.PveSockets = v
	}
	if v, ok := in["pve_memory"].(string); ok && len(v) > 0 {
		obj.PveMemory = v
	}

	if v, ok := in["pve_boot_disk_size"].(string); ok && len(v) > 0 {
		obj.PveBootDiskSize = v
	}
	if v, ok := in["pve_boot_disk_device"].(string); ok && len(v) > 0 {
		obj.PveBootDiskDevice = v
	}
	if v, ok := in["pve_backup"].(string); ok && len(v) > 0 {
		obj.PveBackup = v
	}
	if v, ok := in["pve_data_disk"].([]interface{}); ok && len(v) > 0 {
		obj.PveDataDisk = toArrayString(v)
	}
	if v, ok := in["pve_disk_setup_timeout"].(string); ok && len(v) > 0 {
		obj.PveDiskSetupTimeout = v
	}

	if v, ok := in["pve_net_device"].(string); ok && len(v) > 0 {
		obj.PveNetDevice = v
	}
	if v, ok := in["pve_net_bridge"].(string); ok && len(v) > 0 {
		obj.PveNetBridge = v
	}
	if v, ok := in["pve_net_model"].(string); ok && len(v) > 0 {
		obj.PveNetModel = v
	}
	if v, ok := in["pve_net_vlan_tag"].(string); ok && len(v) > 0 {
		obj.PveNetVlanTag = v
	}
	if v, ok := in["pve_net_mtu"].(string); ok && len(v) > 0 {
		obj.PveNetMtu = v
	}
	if v, ok := in["pve_net_firewall"].(string); ok && len(v) > 0 {
		obj.PveNetFirewall = v
	}
	if v, ok := in["pve_net_iface"].(string); ok && len(v) > 0 {
		obj.PveNetIface = v
	}
	if v, ok := in["pve_agent_timeout"].(string); ok && len(v) > 0 {
		obj.PveAgentTimeout = v
	}
	if v, ok := in["pve_cloudinit_timeout"].(string); ok && len(v) > 0 {
		obj.PveCloudinitTimeout = v
	}
	if v, ok := in["pve_provision_delay"].(string); ok && len(v) > 0 {
		obj.PveProvisionDelay = v
	}

	if v, ok := in["pve_ip_mode"].(string); ok && len(v) > 0 {
		obj.PveIPMode = v
	}
	if v, ok := in["pve_ip_start"].(string); ok && len(v) > 0 {
		obj.PveIPStart = v
	}
	if v, ok := in["pve_ip_end"].(string); ok && len(v) > 0 {
		obj.PveIPEnd = v
	}
	if v, ok := in["pve_ip_prefix"].(string); ok && len(v) > 0 {
		obj.PveIPPrefix = v
	}
	if v, ok := in["pve_gateway"].(string); ok && len(v) > 0 {
		obj.PveGateway = v
	}
	if v, ok := in["pve_nameservers"].(string); ok && len(v) > 0 {
		obj.PveNameservers = v
	}
	if v, ok := in["pve_searchdomain"].(string); ok && len(v) > 0 {
		obj.PveSearchdomain = v
	}
	if v, ok := in["pve_ciuser"].(string); ok && len(v) > 0 {
		obj.PveCiuser = v
	}
	if v, ok := in["pve_sshkeys"].(string); ok && len(v) > 0 {
		obj.PveSSHKeys = v
	}
	if v, ok := in["pve_ssh_user"].(string); ok && len(v) > 0 {
		obj.PveSSHUser = v
	}
	if v, ok := in["pve_ssh_port"].(string); ok && len(v) > 0 {
		obj.PveSSHPort = v
	}
	if v, ok := in["pve_cicustom"].(string); ok && len(v) > 0 {
		obj.PveCicustom = v
	}

	if v, ok := in["pve_extra_config"].([]interface{}); ok && len(v) > 0 {
		obj.PveExtraConfig = toArrayString(v)
	}

	if v, ok := in["pve_keep_on_failure"].(bool); ok {
		obj.PveKeepOnFailure = v
	}

	return obj
}
