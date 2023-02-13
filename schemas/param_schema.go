package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate Param resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ParamModel(d *schema.ResourceData) *models.Param {
	name, _ := d.Get("name").(string)
	value, _ := d.Get("value").(string)
	return &models.Param{
		Name:  name,
		Value: value,
	}
}

func ParamModelFromMap(m map[string]interface{}) *models.Param {
	name := m["name"].(string)
	value := m["value"].(string)
	return &models.Param{
		Name:  name,
		Value: value,
	}
}

// Update the underlying Param resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetParamResourceData(d *schema.ResourceData, m *models.Param) {
	d.Set("name", m.Name)
	d.Set("value", m.Value)
}

// Iterate through and update the Param resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetParamSubResourceData(m []*models.Param) (d []*map[string]interface{}) {
	for _, ParamModel := range m {
		if ParamModel != nil {
			properties := make(map[string]interface{})
			properties["name"] = ParamModel.Name
			properties["value"] = ParamModel.Value
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the Param resource defined in the Terraform configuration
func ParamSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Description: `Name of the Parameter (Required)`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"value": {
			Description: `Value of the parameter (Required)`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the Param resource
func GetParamPropertyFields() (t []string) {
	return []string{
		"name",
		"value",
	}
}
