package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate OpaqueConfigType resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func OpaqueConfigTypeModel(d *schema.ResourceData) *models.OpaqueConfigType {
	opaqueConfigType, _ := d.Get("opaque_config_type").(models.OpaqueConfigType)
	return &opaqueConfigType
}

func OpaqueConfigTypeModelFromMap(m map[string]interface{}) *models.OpaqueConfigType {
	opaqueConfigType := m["opaque_config_type"].(models.OpaqueConfigType)
	return &opaqueConfigType
}

// Update the underlying OpaqueConfigType resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetOpaqueConfigTypeResourceData(d *schema.ResourceData, m *models.OpaqueConfigType) {
}

// Iterate through and update the OpaqueConfigType resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetOpaqueConfigTypeSubResourceData(m []*models.OpaqueConfigType) (d []*map[string]interface{}) {
	for _, OpaqueConfigTypeModel := range m {
		if OpaqueConfigTypeModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the OpaqueConfigType resource defined in the Terraform configuration
func OpaqueConfigTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the OpaqueConfigType resource
func GetOpaqueConfigTypePropertyFields() (t []string) {
	return []string{}
}
