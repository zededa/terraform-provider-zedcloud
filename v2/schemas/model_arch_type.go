package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

// Function to perform the following actions:
// (1) Translate ModelArchType resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ModelArchTypeModel(d *schema.ResourceData) *models.ModelArchType {
	modelArchType, _ := d.Get("model_arch_type").(models.ModelArchType)
	return &modelArchType
}

func ModelArchTypeModelFromMap(m map[string]interface{}) *models.ModelArchType {
	modelArchType := m["model_arch_type"].(models.ModelArchType)
	return &modelArchType
}

// Update the underlying ModelArchType resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetModelArchTypeResourceData(d *schema.ResourceData, m *models.ModelArchType) {
}

// Iterate through and update the ModelArchType resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetModelArchTypeSubResourceData(m []*models.ModelArchType) (d []*map[string]interface{}) {
	for _, ModelArchTypeModel := range m {
		if ModelArchTypeModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the ModelArchType resource defined in the Terraform configuration
func ModelArchTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the ModelArchType resource
func GetModelArchTypePropertyFields() (t []string) {
	return []string{}
}
