package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ModuleSummaryModel(d *schema.ResourceData) *models.ModuleSummary {
	var moduleType *models.ModuleType // ModuleType
	moduleTypeInterface, moduleTypeIsSet := d.GetOk("module_type")
	if moduleTypeIsSet {
		moduleTypeModel := moduleTypeInterface.(string)
		moduleType = models.NewModuleType(models.ModuleType(moduleTypeModel))
	}
	return &models.ModuleSummary{
		ModuleType: moduleType,
	}
}

func ModuleSummaryModelFromMap(m map[string]interface{}) *models.ModuleSummary {
	var moduleType *models.ModuleType // ModuleType
	moduleTypeInterface, moduleTypeIsSet := m["module_type"]
	if moduleTypeIsSet {
		moduleTypeModel := moduleTypeInterface.(string)
		moduleType = models.NewModuleType(models.ModuleType(moduleTypeModel))
	}
	return &models.ModuleSummary{
		ModuleType: moduleType,
	}
}

// Update the underlying ModuleSummary resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetModuleSummaryResourceData(d *schema.ResourceData, m *models.ModuleSummary) {
	d.Set("module_type", m.ModuleType)
}

// Iterate through and update the ModuleSummary resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetModuleSummarySubResourceData(m []*models.ModuleSummary) (d []*map[string]interface{}) {
	for _, ModuleSummaryModel := range m {
		if ModuleSummaryModel != nil {
			properties := make(map[string]interface{})
			properties["module_type"] = ModuleSummaryModel.ModuleType
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the ModuleSummary resource defined in the Terraform configuration
func ModuleSummarySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"module_type": {
			Description: `Type of modules`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the ModuleSummary resource
func GetModuleSummaryPropertyFields() (t []string) {
	return []string{
		"module_type",
	}
}
