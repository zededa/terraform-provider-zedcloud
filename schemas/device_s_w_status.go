package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

// Function to perform the following actions:
// (1) Translate DeviceSWStatus resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func DeviceSWStatusModel(d *schema.ResourceData) *models.DeviceSWStatus {
	deviceSWStatus, _ := d.Get("device_s_w_status").(models.DeviceSWStatus)
	return &deviceSWStatus
}

func DeviceSWStatusModelFromMap(m map[string]interface{}) *models.DeviceSWStatus {
	deviceSWStatus := m["device_s_w_status"].(models.DeviceSWStatus)
	return &deviceSWStatus
}

// Update the underlying DeviceSWStatus resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDeviceSWStatusResourceData(d *schema.ResourceData, m *models.DeviceSWStatus) {
}

// Iterate through and update the DeviceSWStatus resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDeviceSWStatusSubResourceData(m []*models.DeviceSWStatus) (d []*map[string]interface{}) {
	for _, DeviceSWStatusModel := range m {
		if DeviceSWStatusModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the DeviceSWStatus resource defined in the Terraform configuration
func DeviceSWStatusSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the DeviceSWStatus resource
func GetDeviceSWStatusPropertyFields() (t []string) {
	return []string{}
}
