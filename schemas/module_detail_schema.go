package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate ModuleDetail resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ModuleDetailModel(d *schema.ResourceData) *models.ModuleDetail {
	environment := d.Get("environment").(map[string]string) // map[string]string
	moduleType := d.Get("module_type").(*models.ModuleType) // ModuleType
	routes := d.Get("routes").(map[string]string)           // map[string]string
	twinDetail := d.Get("twin_detail").(string)
	return &models.ModuleDetail{
		Environment: environment,
		ModuleType:  moduleType,
		Routes:      routes,
		TwinDetail:  twinDetail,
	}
}

func ModuleDetailModelFromMap(m map[string]interface{}) *models.ModuleDetail {
	environment := m["environment"].(map[string]string)
	moduleType := m["module_type"].(*models.ModuleType) // ModuleType
	routes := m["routes"].(map[string]string)
	twinDetail := m["twin_detail"].(string)
	return &models.ModuleDetail{
		Environment: environment,
		ModuleType:  moduleType,
		Routes:      routes,
		TwinDetail:  twinDetail,
	}
}

// Update the underlying ModuleDetail resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetModuleDetailResourceData(d *schema.ResourceData, m *models.ModuleDetail) {
	d.Set("environment", m.Environment)
	d.Set("module_type", m.ModuleType)
	d.Set("routes", m.Routes)
	d.Set("twin_detail", m.TwinDetail)
}

// Iterate throught and update the ModuleDetail resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetModuleDetailSubResourceData(m []*models.ModuleDetail) (d []*map[string]interface{}) {
	for _, ModuleDetailModel := range m {
		if ModuleDetailModel != nil {
			properties := make(map[string]interface{})
			properties["environment"] = ModuleDetailModel.Environment
			properties["module_type"] = ModuleDetailModel.ModuleType
			properties["routes"] = ModuleDetailModel.Routes
			properties["twin_detail"] = ModuleDetailModel.TwinDetail
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the ModuleDetail resource defined in the Terraform configuration
func ModuleDetailSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"environment": {
			Description: `Extra information to module to make configuration easier`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"module_type": {
			Description: `Type of modules`,
			// We assume it's an enum type
			Type:     schema.TypeString,
			Optional: true,
		},

		"routes": {
			Description: `Send messages between modules or send messages from modules to iot hub`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"twin_detail": {
			Description: `Base64 encoded module twin details, desired properties of the module will be updated to reflect these values`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the ModuleDetail resource
func GetModuleDetailPropertyFields() (t []string) {
	return []string{
		"environment",
		"module_type",
		"routes",
		"twin_detail",
	}
}
