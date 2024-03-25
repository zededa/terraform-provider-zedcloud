package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

// Function to perform the following actions:
// (1) Translate ModuleType resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ModuleTypeModel(d *schema.ResourceData) *models.ModuleType {
	moduleType, _ := d.Get("module_type").(models.ModuleType)
	return &moduleType
}

func ModuleTypeModelFromMap(m map[string]interface{}) *models.ModuleType {
	moduleType := m["module_type"].(models.ModuleType)
	return &moduleType
}

// Update the underlying ModuleType resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetModuleTypeResourceData(d *schema.ResourceData, m *models.ModuleType) {
}

// Iterate through and update the ModuleType resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetModuleTypeSubResourceData(m []*models.ModuleType) (d []*map[string]interface{}) {
	for _, ModuleTypeModel := range m {
		if ModuleTypeModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the ModuleType resource defined in the Terraform configuration
func ModuleTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the ModuleType resource
func GetModuleTypePropertyFields() (t []string) {
	return []string{}
}
