package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate DeviceAttestationPolicy resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func DeviceAttestationPolicyModel(d *schema.ResourceData) *models.DeviceAttestationPolicy {
	typeVar := d.Get("type").(*models.DeviceAttestPolicyType) // DeviceAttestPolicyType
	return &models.DeviceAttestationPolicy{
		Type: typeVar,
	}
}

func DeviceAttestationPolicyModelFromMap(m map[string]interface{}) *models.DeviceAttestationPolicy {
	typeVar := m["type"].(*models.DeviceAttestPolicyType) // DeviceAttestPolicyType
	return &models.DeviceAttestationPolicy{
		Type: typeVar,
	}
}

// Update the underlying DeviceAttestationPolicy resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDeviceAttestationPolicyResourceData(d *schema.ResourceData, m *models.DeviceAttestationPolicy) {
	d.Set("type", m.Type)
}

// Iterate throught and update the DeviceAttestationPolicy resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDeviceAttestationPolicySubResourceData(m []*models.DeviceAttestationPolicy) (d []*map[string]interface{}) {
	for _, DeviceAttestationPolicyModel := range m {
		if DeviceAttestationPolicyModel != nil {
			properties := make(map[string]interface{})
			properties["type"] = DeviceAttestationPolicyModel.Type
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the DeviceAttestationPolicy resource defined in the Terraform configuration
func DeviceAttestationPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"type": {
			Description: `Attestation policy type`,
			// We assume it's an enum type
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

// Retrieve property field names for updating the DeviceAttestationPolicy resource
func GetDeviceAttestationPolicyPropertyFields() (t []string) {
	return []string{
		"type",
	}
}
