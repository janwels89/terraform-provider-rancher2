package rancher2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testMachineConfigV2PveConf = &MachineConfigV2Pve{
		machineConfigV2Pve: machineConfigV2Pve{
			PveURL:              "https://pve.example.com:8006",
			PveTokenID:          "root@pam!rancher",
			PveTokenSecret:      "secret-uuid",
			PveInsecureTLS:      false,
			PveNode:             "pve1",
			PveAllowedNodes:     "pve1,pve2",
			PveTemplateVMID:     "100",
			PveTemplateTag:      "rancher-node",
			PveTemplateTagMatch: "subset",
			PveLinkedClone:      true,
			PveCloneStorage:     "ceph-rbd",
			PveCloneFormat:      "qcow2",
			PveVMID:             "0",
			PveVMIDRange:        "200-299",
			PveTags:             "foo,bar",
			PveDescription:      "managed by rancher2",
			PveVMNamePrefix:     "k8s",
			PveOnboot:           true,
			PveHA:               true,
			PveHAGroup:          "ha-group",
			PveCores:            "4",
			PveSockets:          "2",
			PveMemory:           "4096",
			PveBootDiskSize:     "40",
			PveBootDiskDevice:   "scsi0",
			PveBackup:           "true",
			PveDataDisk:         []string{"size=100,storage=local-lvm,fs=ext4,mount=/var/lib/longhorn"},
			PveDiskSetupTimeout: "300",
			PveNetDevice:        "net0",
			PveNetBridge:        "vmbr1",
			PveNetModel:         "virtio",
			PveNetVlanTag:       "100",
			PveNetMtu:           "9000",
			PveNetFirewall:      "true",
			PveNetIface:         "eth0",
			PveAgentTimeout:     "300",
			PveCloudinitTimeout: "300",
			PveProvisionDelay:   "30",
			PveIPMode:           "static",
			PveIPStart:          "192.168.15.150",
			PveIPEnd:            "192.168.15.159",
			PveIPPrefix:         "24",
			PveGateway:          "192.168.15.1",
			PveNameservers:      "1.1.1.1",
			PveSearchdomain:     "cluster.lan",
			PveCiuser:           "rancher",
			PveSSHKeys:          "ssh-ed25519 AAAA...",
			PveSSHUser:          "rancher",
			PveSSHPort:          "22",
			PveCicustom:         "vendor=local:snippets/rancher.yaml",
			PveExtraConfig:      []string{"cpu=host"},
			PveKeepOnFailure:    false,
		},
	}
	testMachineConfigV2PveInterface = []interface{}{
		map[string]interface{}{
			"pve_url":                "https://pve.example.com:8006",
			"pve_token_id":           "root@pam!rancher",
			"pve_token_secret":       "secret-uuid",
			"pve_insecure_tls":       false,
			"pve_node":               "pve1",
			"pve_allowed_nodes":      "pve1,pve2",
			"pve_template_vmid":      "100",
			"pve_template_tag":       "rancher-node",
			"pve_template_tag_match": "subset",
			"pve_linked_clone":       true,
			"pve_clone_storage":      "ceph-rbd",
			"pve_clone_format":       "qcow2",
			"pve_vmid":               "0",
			"pve_vmid_range":         "200-299",
			"pve_tags":               "foo,bar",
			"pve_description":        "managed by rancher2",
			"pve_vm_name_prefix":     "k8s",
			"pve_onboot":             true,
			"pve_ha":                 true,
			"pve_ha_group":           "ha-group",
			"pve_cores":              "4",
			"pve_sockets":            "2",
			"pve_memory":             "4096",
			"pve_boot_disk_size":     "40",
			"pve_boot_disk_device":   "scsi0",
			"pve_backup":             "true",
			"pve_data_disk":          []interface{}{"size=100,storage=local-lvm,fs=ext4,mount=/var/lib/longhorn"},
			"pve_disk_setup_timeout": "300",
			"pve_net_device":         "net0",
			"pve_net_bridge":         "vmbr1",
			"pve_net_model":          "virtio",
			"pve_net_vlan_tag":       "100",
			"pve_net_mtu":            "9000",
			"pve_net_firewall":       "true",
			"pve_net_iface":          "eth0",
			"pve_agent_timeout":      "300",
			"pve_cloudinit_timeout":  "300",
			"pve_provision_delay":    "30",
			"pve_ip_mode":            "static",
			"pve_ip_start":           "192.168.15.150",
			"pve_ip_end":             "192.168.15.159",
			"pve_ip_prefix":          "24",
			"pve_gateway":            "192.168.15.1",
			"pve_nameservers":        "1.1.1.1",
			"pve_searchdomain":       "cluster.lan",
			"pve_ciuser":             "rancher",
			"pve_sshkeys":            "ssh-ed25519 AAAA...",
			"pve_ssh_user":           "rancher",
			"pve_ssh_port":           "22",
			"pve_cicustom":           "vendor=local:snippets/rancher.yaml",
			"pve_extra_config":       []interface{}{"cpu=host"},
			"pve_keep_on_failure":    false,
		},
	}
)

func TestFlattenMachineConfigV2Pve(t *testing.T) {
	result := flattenMachineConfigV2Pve(testMachineConfigV2PveConf)
	assert.Equal(t, testMachineConfigV2PveInterface, result)
}

func TestExpandMachineConfigV2Pve(t *testing.T) {
	source := &MachineConfigV2{}
	result := expandMachineConfigV2Pve(testMachineConfigV2PveInterface, source)
	assert.Equal(t, machineConfigV2PveKind, result.TypeMeta.Kind)
	assert.Equal(t, machineConfigV2PveAPIVersion, result.TypeMeta.APIVersion)
	assert.Equal(t, machineConfigV2PveKind, source.TypeMeta.Kind)
	assert.Equal(t, machineConfigV2PveAPIVersion, source.TypeMeta.APIVersion)
	assert.Equal(t, testMachineConfigV2PveConf.PveURL, result.PveURL)
	assert.Equal(t, testMachineConfigV2PveConf.PveTokenID, result.PveTokenID)
	assert.Equal(t, testMachineConfigV2PveConf.PveTokenSecret, result.PveTokenSecret)
	assert.Equal(t, testMachineConfigV2PveConf.PveInsecureTLS, result.PveInsecureTLS)
	assert.Equal(t, testMachineConfigV2PveConf.PveNode, result.PveNode)
	assert.Equal(t, testMachineConfigV2PveConf.PveAllowedNodes, result.PveAllowedNodes)
	assert.Equal(t, testMachineConfigV2PveConf.PveTemplateVMID, result.PveTemplateVMID)
	assert.Equal(t, testMachineConfigV2PveConf.PveTemplateTag, result.PveTemplateTag)
	assert.Equal(t, testMachineConfigV2PveConf.PveTemplateTagMatch, result.PveTemplateTagMatch)
	assert.Equal(t, testMachineConfigV2PveConf.PveLinkedClone, result.PveLinkedClone)
	assert.Equal(t, testMachineConfigV2PveConf.PveCloneStorage, result.PveCloneStorage)
	assert.Equal(t, testMachineConfigV2PveConf.PveCloneFormat, result.PveCloneFormat)
	assert.Equal(t, testMachineConfigV2PveConf.PveVMID, result.PveVMID)
	assert.Equal(t, testMachineConfigV2PveConf.PveVMIDRange, result.PveVMIDRange)
	assert.Equal(t, testMachineConfigV2PveConf.PveTags, result.PveTags)
	assert.Equal(t, testMachineConfigV2PveConf.PveDescription, result.PveDescription)
	assert.Equal(t, testMachineConfigV2PveConf.PveVMNamePrefix, result.PveVMNamePrefix)
	assert.Equal(t, testMachineConfigV2PveConf.PveOnboot, result.PveOnboot)
	assert.Equal(t, testMachineConfigV2PveConf.PveHA, result.PveHA)
	assert.Equal(t, testMachineConfigV2PveConf.PveHAGroup, result.PveHAGroup)
	assert.Equal(t, testMachineConfigV2PveConf.PveCores, result.PveCores)
	assert.Equal(t, testMachineConfigV2PveConf.PveSockets, result.PveSockets)
	assert.Equal(t, testMachineConfigV2PveConf.PveMemory, result.PveMemory)
	assert.Equal(t, testMachineConfigV2PveConf.PveBootDiskSize, result.PveBootDiskSize)
	assert.Equal(t, testMachineConfigV2PveConf.PveBootDiskDevice, result.PveBootDiskDevice)
	assert.Equal(t, testMachineConfigV2PveConf.PveBackup, result.PveBackup)
	assert.Equal(t, testMachineConfigV2PveConf.PveDataDisk, result.PveDataDisk)
	assert.Equal(t, testMachineConfigV2PveConf.PveDiskSetupTimeout, result.PveDiskSetupTimeout)
	assert.Equal(t, testMachineConfigV2PveConf.PveNetDevice, result.PveNetDevice)
	assert.Equal(t, testMachineConfigV2PveConf.PveNetBridge, result.PveNetBridge)
	assert.Equal(t, testMachineConfigV2PveConf.PveNetModel, result.PveNetModel)
	assert.Equal(t, testMachineConfigV2PveConf.PveNetVlanTag, result.PveNetVlanTag)
	assert.Equal(t, testMachineConfigV2PveConf.PveNetMtu, result.PveNetMtu)
	assert.Equal(t, testMachineConfigV2PveConf.PveNetFirewall, result.PveNetFirewall)
	assert.Equal(t, testMachineConfigV2PveConf.PveNetIface, result.PveNetIface)
	assert.Equal(t, testMachineConfigV2PveConf.PveAgentTimeout, result.PveAgentTimeout)
	assert.Equal(t, testMachineConfigV2PveConf.PveCloudinitTimeout, result.PveCloudinitTimeout)
	assert.Equal(t, testMachineConfigV2PveConf.PveProvisionDelay, result.PveProvisionDelay)
	assert.Equal(t, testMachineConfigV2PveConf.PveIPMode, result.PveIPMode)
	assert.Equal(t, testMachineConfigV2PveConf.PveIPStart, result.PveIPStart)
	assert.Equal(t, testMachineConfigV2PveConf.PveIPEnd, result.PveIPEnd)
	assert.Equal(t, testMachineConfigV2PveConf.PveIPPrefix, result.PveIPPrefix)
	assert.Equal(t, testMachineConfigV2PveConf.PveGateway, result.PveGateway)
	assert.Equal(t, testMachineConfigV2PveConf.PveNameservers, result.PveNameservers)
	assert.Equal(t, testMachineConfigV2PveConf.PveSearchdomain, result.PveSearchdomain)
	assert.Equal(t, testMachineConfigV2PveConf.PveCiuser, result.PveCiuser)
	assert.Equal(t, testMachineConfigV2PveConf.PveSSHKeys, result.PveSSHKeys)
	assert.Equal(t, testMachineConfigV2PveConf.PveSSHUser, result.PveSSHUser)
	assert.Equal(t, testMachineConfigV2PveConf.PveSSHPort, result.PveSSHPort)
	assert.Equal(t, testMachineConfigV2PveConf.PveCicustom, result.PveCicustom)
	assert.Equal(t, testMachineConfigV2PveConf.PveExtraConfig, result.PveExtraConfig)
	assert.Equal(t, testMachineConfigV2PveConf.PveKeepOnFailure, result.PveKeepOnFailure)
}
