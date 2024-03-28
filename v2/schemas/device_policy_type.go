package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

// Function to perform the following actions:
// (1) Translate DevicePolicyType resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func DevicePolicyTypeModel(d *schema.ResourceData) *models.DevicePolicyType {
	devicePolicyType, _ := d.Get("device_policy_type").(models.DevicePolicyType)
	return &devicePolicyType
}

func DevicePolicyTypeModelFromMap(m map[string]interface{}) *models.DevicePolicyType {
	devicePolicyType := m["device_policy_type"].(models.DevicePolicyType)
	return &devicePolicyType
}

// Update the underlying DevicePolicyType resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDevicePolicyTypeResourceData(d *schema.ResourceData, m *models.DevicePolicyType) {
}

// Iterate through and update the DevicePolicyType resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDevicePolicyTypeSubResourceData(m []*models.DevicePolicyType) (d []*map[string]interface{}) {
	for _, DevicePolicyTypeModel := range m {
		if DevicePolicyTypeModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the DevicePolicyType resource defined in the Terraform configuration
func DevicePolicyTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the DevicePolicyType resource
func GetDevicePolicyTypePropertyFields() (t []string) {
	return []string{}
}
