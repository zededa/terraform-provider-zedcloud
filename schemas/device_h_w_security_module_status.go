package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

// Function to perform the following actions:
// (1) Translate DeviceHWSecurityModuleStatus resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func DeviceHWSecurityModuleStatusModel(d *schema.ResourceData) *models.DeviceHWSecurityModuleStatus {
	deviceHWSecurityModuleStatus, _ := d.Get("device_h_w_security_module_status").(models.DeviceHWSecurityModuleStatus)
	return &deviceHWSecurityModuleStatus
}

func DeviceHWSecurityModuleStatusModelFromMap(m map[string]interface{}) *models.DeviceHWSecurityModuleStatus {
	deviceHWSecurityModuleStatus := m["device_h_w_security_module_status"].(models.DeviceHWSecurityModuleStatus)
	return &deviceHWSecurityModuleStatus
}

// Update the underlying DeviceHWSecurityModuleStatus resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDeviceHWSecurityModuleStatusResourceData(d *schema.ResourceData, m *models.DeviceHWSecurityModuleStatus) {
}

// Iterate through and update the DeviceHWSecurityModuleStatus resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDeviceHWSecurityModuleStatusSubResourceData(m []*models.DeviceHWSecurityModuleStatus) (d []*map[string]interface{}) {
	for _, DeviceHWSecurityModuleStatusModel := range m {
		if DeviceHWSecurityModuleStatusModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the DeviceHWSecurityModuleStatus resource defined in the Terraform configuration
func DeviceHWSecurityModuleStatusSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the DeviceHWSecurityModuleStatus resource
func GetDeviceHWSecurityModuleStatusPropertyFields() (t []string) {
	return []string{}
}
