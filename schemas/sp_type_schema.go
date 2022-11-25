package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate SpType resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func SpTypeModel(d *schema.ResourceData) *models.SpType {
	spType := d.Get("sp_type").(models.SpType)
	return &spType
}

func SpTypeModelFromMap(m map[string]interface{}) *models.SpType {
	spType := m["sp_type"].(models.SpType)
	return &spType
}

// Update the underlying SpType resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetSpTypeResourceData(d *schema.ResourceData, m *models.SpType) {
}

// Iterate throught and update the SpType resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetSpTypeSubResourceData(m []*models.SpType) (d []*map[string]interface{}) {
	for _, SpTypeModel := range m {
		if SpTypeModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the SpType resource defined in the Terraform configuration
func SpTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the SpType resource
func GetSpTypePropertyFields() (t []string) {
	return []string{}
}
