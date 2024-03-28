package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

// Function to perform the following actions:
// (1) Translate VariableFileEncoding resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func VariableFileEncodingModel(d *schema.ResourceData) *models.VariableFileEncoding {
	variableFileEncoding, _ := d.Get("variable_file_encoding").(models.VariableFileEncoding)
	return &variableFileEncoding
}

func VariableFileEncodingModelFromMap(m map[string]interface{}) *models.VariableFileEncoding {
	variableFileEncoding := m["variable_file_encoding"].(models.VariableFileEncoding)
	return &variableFileEncoding
}

// Update the underlying VariableFileEncoding resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetVariableFileEncodingResourceData(d *schema.ResourceData, m *models.VariableFileEncoding) {
}

// Iterate through and update the VariableFileEncoding resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetVariableFileEncodingSubResourceData(m []*models.VariableFileEncoding) (d []*map[string]interface{}) {
	for _, VariableFileEncodingModel := range m {
		if VariableFileEncodingModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the VariableFileEncoding resource defined in the Terraform configuration
func VariableFileEncodingSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the VariableFileEncoding resource
func GetVariableFileEncodingPropertyFields() (t []string) {
	return []string{}
}
