package provider

import "github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/features"

func expandFeatures(input []interface{}) features.UserFeatures {
	// TODO: in 2.0 when Required this can become:
	//val := input[0].(map[string]interface{})

	var val map[string]interface{}
	if len(input) > 0 {
		val = input[0].(map[string]interface{})
	}

	features := features.UserFeatures{
		// NOTE: ensure all nested objects are fully populated
		VirtualMachine: &features.VirtualMachineFeatures{
			DeleteOSDiskOnDeletion: true,
		},
	}

	if raw, ok := val["virtual_machines"]; ok {
		items := raw.([]interface{})
		if len(items) > 0 {
			virtualMachinesRaw := items[0].(map[string]interface{})
			if v, ok := virtualMachinesRaw["delete_os_disk_on_deletion"]; ok {
				features.VirtualMachine.DeleteOSDiskOnDeletion = v.(bool)
			}
		}
	}

	return features
}
