package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate DeviceSWSubStatus resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func DeviceSWSubStatusModel(d *schema.ResourceData) *models.DeviceSWSubStatus {
	deviceSWSubStatus := d.Get("device_s_w_sub_status").(models.DeviceSWSubStatus)
	return &deviceSWSubStatus
}

func DeviceSWSubStatusModelFromMap(m map[string]interface{}) *models.DeviceSWSubStatus {
	deviceSWSubStatus := m["device_s_w_sub_status"].(models.DeviceSWSubStatus)
	return &deviceSWSubStatus
}

// Update the underlying DeviceSWSubStatus resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDeviceSWSubStatusResourceData(d *schema.ResourceData, m *models.DeviceSWSubStatus) {
}

// Iterate throught and update the DeviceSWSubStatus resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDeviceSWSubStatusSubResourceData(m []*models.DeviceSWSubStatus) (d []*map[string]interface{}) {
	for _, DeviceSWSubStatusModel := range m {
		if DeviceSWSubStatusModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the DeviceSWSubStatus resource defined in the Terraform configuration
func DeviceSWSubStatusSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the DeviceSWSubStatus resource
func GetDeviceSWSubStatusPropertyFields() (t []string) {
	return []string{}
}
