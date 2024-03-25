package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

// Function to perform the following actions:
// (1) Translate VariableVariableFormat resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func VariableVariableFormatModel(d *schema.ResourceData) *models.VariableVariableFormat {
	variableVariableFormat, _ := d.Get("variable_variable_format").(models.VariableVariableFormat)
	return &variableVariableFormat
}

func VariableVariableFormatModelFromMap(m map[string]interface{}) *models.VariableVariableFormat {
	variableVariableFormat := m["variable_variable_format"].(models.VariableVariableFormat)
	return &variableVariableFormat
}

// Update the underlying VariableVariableFormat resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetVariableVariableFormatResourceData(d *schema.ResourceData, m *models.VariableVariableFormat) {
}

// Iterate through and update the VariableVariableFormat resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetVariableVariableFormatSubResourceData(m []*models.VariableVariableFormat) (d []*map[string]interface{}) {
	for _, VariableVariableFormatModel := range m {
		if VariableVariableFormatModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the VariableVariableFormat resource defined in the Terraform configuration
func VariableVariableFormatSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the VariableVariableFormat resource
func GetVariableVariableFormatPropertyFields() (t []string) {
	return []string{}
}
