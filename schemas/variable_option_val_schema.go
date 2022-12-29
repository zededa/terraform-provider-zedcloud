package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate VariableOptionVal resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func VariableOptionValModel(d *schema.ResourceData) *models.VariableOptionVal {
	label, _ := d.Get("label").(string)
	value, _ := d.Get("value").(string)
	return &models.VariableOptionVal{
		Label: label,
		Value: value,
	}
}

func VariableOptionValModelFromMap(m map[string]interface{}) *models.VariableOptionVal {
	label := m["label"].(string)
	value := m["value"].(string)
	return &models.VariableOptionVal{
		Label: label,
		Value: value,
	}
}

// Update the underlying VariableOptionVal resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetVariableOptionValResourceData(d *schema.ResourceData, m *models.VariableOptionVal) {
	d.Set("label", m.Label)
	d.Set("value", m.Value)
}

// Iterate through and update the VariableOptionVal resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetVariableOptionValSubResourceData(m []*models.VariableOptionVal) (d []*map[string]interface{}) {
	for _, VariableOptionValModel := range m {
		if VariableOptionValModel != nil {
			properties := make(map[string]interface{})
			properties["label"] = VariableOptionValModel.Label
			properties["value"] = VariableOptionValModel.Value
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the VariableOptionVal resource defined in the Terraform configuration
func VariableOptionValSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"label": {
			Description: `Display label of the key in User-Agent`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"value": {
			Description: `Value of the key to be used`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the VariableOptionVal resource
func GetVariableOptionValPropertyFields() (t []string) {
	return []string{
		"label",
		"value",
	}
}
