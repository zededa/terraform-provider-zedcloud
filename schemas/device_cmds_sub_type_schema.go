package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate DeviceCmdsSubType resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func DeviceCmdsSubTypeModel(d *schema.ResourceData) *models.DeviceCmdsSubType {
	deviceCmdsSubType, _ := d.Get("device_cmds_sub_type").(models.DeviceCmdsSubType)
	return &deviceCmdsSubType
}

func DeviceCmdsSubTypeModelFromMap(m map[string]interface{}) *models.DeviceCmdsSubType {
	deviceCmdsSubType := m["device_cmds_sub_type"].(models.DeviceCmdsSubType)
	return &deviceCmdsSubType
}

// Update the underlying DeviceCmdsSubType resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDeviceCmdsSubTypeResourceData(d *schema.ResourceData, m *models.DeviceCmdsSubType) {
}

// Iterate through and update the DeviceCmdsSubType resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDeviceCmdsSubTypeSubResourceData(m []*models.DeviceCmdsSubType) (d []*map[string]interface{}) {
	for _, DeviceCmdsSubTypeModel := range m {
		if DeviceCmdsSubTypeModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the DeviceCmdsSubType resource defined in the Terraform configuration
func DeviceCmdsSubTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the DeviceCmdsSubType resource
func GetDeviceCmdsSubTypePropertyFields() (t []string) {
	return []string{}
}
