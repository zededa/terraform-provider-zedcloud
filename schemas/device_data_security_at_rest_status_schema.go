package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate DeviceDataSecurityAtRestStatus resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func DeviceDataSecurityAtRestStatusModel(d *schema.ResourceData) *models.DeviceDataSecurityAtRestStatus {
	deviceDataSecurityAtRestStatus, _ := d.Get("device_data_security_at_rest_status").(models.DeviceDataSecurityAtRestStatus)
	return &deviceDataSecurityAtRestStatus
}

func DeviceDataSecurityAtRestStatusModelFromMap(m map[string]interface{}) *models.DeviceDataSecurityAtRestStatus {
	deviceDataSecurityAtRestStatus := m["device_data_security_at_rest_status"].(models.DeviceDataSecurityAtRestStatus)
	return &deviceDataSecurityAtRestStatus
}

// Update the underlying DeviceDataSecurityAtRestStatus resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDeviceDataSecurityAtRestStatusResourceData(d *schema.ResourceData, m *models.DeviceDataSecurityAtRestStatus) {
}

// Iterate through and update the DeviceDataSecurityAtRestStatus resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDeviceDataSecurityAtRestStatusSubResourceData(m []*models.DeviceDataSecurityAtRestStatus) (d []*map[string]interface{}) {
	for _, DeviceDataSecurityAtRestStatusModel := range m {
		if DeviceDataSecurityAtRestStatusModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the DeviceDataSecurityAtRestStatus resource defined in the Terraform configuration
func DeviceDataSecurityAtRestStatusSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the DeviceDataSecurityAtRestStatus resource
func GetDeviceDataSecurityAtRestStatusPropertyFields() (t []string) {
	return []string{}
}
