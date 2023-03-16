package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate ConfigFormat resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ConfigFormatModel(d *schema.ResourceData) *models.ConfigFormat {
	configFormat, _ := d.Get("config_format").(models.ConfigFormat)
	return &configFormat
}

func ConfigFormatModelFromMap(m map[string]interface{}) *models.ConfigFormat {
	configFormat := m["config_format"].(models.ConfigFormat)
	return &configFormat
}

// Update the underlying ConfigFormat resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetConfigFormatResourceData(d *schema.ResourceData, m *models.ConfigFormat) {
}

// Iterate through and update the ConfigFormat resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetConfigFormatSubResourceData(m []*models.ConfigFormat) (d []*map[string]interface{}) {
	for _, ConfigFormatModel := range m {
		if ConfigFormatModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the ConfigFormat resource defined in the Terraform configuration
func ConfigFormatSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the ConfigFormat resource
func GetConfigFormatPropertyFields() (t []string) {
	return []string{}
}
