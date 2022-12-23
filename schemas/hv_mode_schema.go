package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate HvMode resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func HvModeModel(d *schema.ResourceData) *models.HvMode {
	hvMode := d.Get("hv_mode").(models.HvMode)
	return &hvMode
}

func HvModeModelFromMap(m map[string]interface{}) *models.HvMode {
	hvMode := m["hv_mode"].(models.HvMode)
	return &hvMode
}

// Update the underlying HvMode resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetHvModeResourceData(d *schema.ResourceData, m *models.HvMode) {
}

// Iterate throught and update the HvMode resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetHvModeSubResourceData(m []*models.HvMode) (d []*map[string]interface{}) {
	for _, HvModeModel := range m {
		if HvModeModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the HvMode resource defined in the Terraform configuration
func HvModeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the HvMode resource
func GetHvModePropertyFields() (t []string) {
	return []string{}
}
