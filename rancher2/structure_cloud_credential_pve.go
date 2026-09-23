package rancher2

// Flatteners

func flattenCloudCredentialPve(in map[string]interface{}, p []interface{}) []interface{} {
	var obj map[string]interface{}
	if len(p) == 0 || p[0] == nil {
		obj = make(map[string]interface{})
	} else {
		obj = p[0].(map[string]interface{})
	}

	if in == nil {
		return []interface{}{}
	}

	if v, ok := in["url"].(string); ok && len(v) > 0 {
		obj["pve_url"] = v
	}
	if v, ok := in["tokenId"].(string); ok && len(v) > 0 {
		obj["pve_token_id"] = v
	}
	if v, ok := in["tokenSecret"].(string); ok && len(v) > 0 {
		obj["pve_token_secret"] = v
	}
	if v, ok := in["insecureTls"].(bool); ok {
		obj["pve_insecure_tls"] = v
	}
	if v, ok := in["caCert"].(string); ok && len(v) > 0 {
		obj["pve_ca_cert"] = v
	}
	if v, ok := in["pool"].(string); ok && len(v) > 0 {
		obj["pve_pool"] = v
	}
	if v, ok := in["skipPermissionCheck"].(bool); ok {
		obj["pve_skip_permission_check"] = v
	}

	return []interface{}{obj}
}

// Expanders

func expandCloudCredentialPve(p []interface{}) map[string]interface{} {
	obj := make(map[string]interface{})
	if len(p) == 0 || p[0] == nil {
		return obj
	}
	in := p[0].(map[string]interface{})

	if v, ok := in["pve_url"].(string); ok && len(v) > 0 {
		obj["url"] = v
	}
	if v, ok := in["pve_token_id"].(string); ok && len(v) > 0 {
		obj["tokenId"] = v
	}
	if v, ok := in["pve_token_secret"].(string); ok && len(v) > 0 {
		obj["tokenSecret"] = v
	}
	if v, ok := in["pve_insecure_tls"].(bool); ok {
		obj["insecureTls"] = v
	}
	if v, ok := in["pve_ca_cert"].(string); ok && len(v) > 0 {
		obj["caCert"] = v
	}
	if v, ok := in["pve_pool"].(string); ok && len(v) > 0 {
		obj["pool"] = v
	}
	if v, ok := in["pve_skip_permission_check"].(bool); ok {
		obj["skipPermissionCheck"] = v
	}

	return obj
}
