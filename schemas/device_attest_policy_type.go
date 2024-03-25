package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

// Function to perform the following actions:
// (1) Translate DeviceAttestPolicyType resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func DeviceAttestPolicyTypeModel(d *schema.ResourceData) *models.DeviceAttestPolicyType {
	deviceAttestPolicyType, _ := d.Get("device_attest_policy_type").(models.DeviceAttestPolicyType)
	return &deviceAttestPolicyType
}

func DeviceAttestPolicyTypeModelFromMap(m map[string]interface{}) *models.DeviceAttestPolicyType {
	deviceAttestPolicyType := m["device_attest_policy_type"].(models.DeviceAttestPolicyType)
	return &deviceAttestPolicyType
}

// Update the underlying DeviceAttestPolicyType resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDeviceAttestPolicyTypeResourceData(d *schema.ResourceData, m *models.DeviceAttestPolicyType) {
}

// Iterate through and update the DeviceAttestPolicyType resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDeviceAttestPolicyTypeSubResourceData(m []*models.DeviceAttestPolicyType) (d []*map[string]interface{}) {
	for _, DeviceAttestPolicyTypeModel := range m {
		if DeviceAttestPolicyTypeModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the DeviceAttestPolicyType resource defined in the Terraform configuration
func DeviceAttestPolicyTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the DeviceAttestPolicyType resource
func GetDeviceAttestPolicyTypePropertyFields() (t []string) {
	return []string{}
}
