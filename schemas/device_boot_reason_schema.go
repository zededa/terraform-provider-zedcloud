package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate DeviceBootReason resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func DeviceBootReasonModel(d *schema.ResourceData) *models.DeviceBootReason {
	deviceBootReason, _ := d.Get("device_boot_reason").(models.DeviceBootReason)
	return &deviceBootReason
}

func DeviceBootReasonModelFromMap(m map[string]interface{}) *models.DeviceBootReason {
	deviceBootReason := m["device_boot_reason"].(models.DeviceBootReason)
	return &deviceBootReason
}

// Update the underlying DeviceBootReason resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDeviceBootReasonResourceData(d *schema.ResourceData, m *models.DeviceBootReason) {
}

// Iterate through and update the DeviceBootReason resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDeviceBootReasonSubResourceData(m []*models.DeviceBootReason) (d []*map[string]interface{}) {
	for _, DeviceBootReasonModel := range m {
		if DeviceBootReasonModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the DeviceBootReason resource defined in the Terraform configuration
func DeviceBootReasonSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the DeviceBootReason resource
func GetDeviceBootReasonPropertyFields() (t []string) {
	return []string{}
}
