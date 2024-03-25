package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

// Function to perform the following actions:
// (1) Translate DeviceLoad resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func DeviceLoadModel(d *schema.ResourceData) *models.DeviceLoad {
	deviceLoad, _ := d.Get("device_load").(models.DeviceLoad)
	return &deviceLoad
}

func DeviceLoadModelFromMap(m map[string]interface{}) *models.DeviceLoad {
	deviceLoad := m["device_load"].(models.DeviceLoad)
	return &deviceLoad
}

// Update the underlying DeviceLoad resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDeviceLoadResourceData(d *schema.ResourceData, m *models.DeviceLoad) {
}

// Iterate through and update the DeviceLoad resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDeviceLoadSubResourceData(m []*models.DeviceLoad) (d []*map[string]interface{}) {
	for _, DeviceLoadModel := range m {
		if DeviceLoadModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the DeviceLoad resource defined in the Terraform configuration
func DeviceLoadSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the DeviceLoad resource
func GetDeviceLoadPropertyFields() (t []string) {
	return []string{}
}
