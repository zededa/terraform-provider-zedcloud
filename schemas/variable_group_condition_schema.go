package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate VariableGroupCondition resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func VariableGroupConditionModel(d *schema.ResourceData) *models.VariableGroupCondition {
	name, _ := d.Get("name").(string)
	var operator *models.VariableGroupConditionOperator // VariableGroupConditionOperator
	operatorInterface, operatorIsSet := d.GetOk("operator")
	if operatorIsSet {
		operatorModel := operatorInterface.(string)
		operator = models.NewVariableGroupConditionOperator(models.VariableGroupConditionOperator(operatorModel))
	}
	value, _ := d.Get("value").(string)
	return &models.VariableGroupCondition{
		Name:     name,
		Operator: operator,
		Value:    value,
	}
}

func VariableGroupConditionModelFromMap(m map[string]interface{}) *models.VariableGroupCondition {
	name := m["name"].(string)
	operator := m["operator"].(*models.VariableGroupConditionOperator) // VariableGroupConditionOperator
	value := m["value"].(string)
	return &models.VariableGroupCondition{
		Name:     name,
		Operator: operator,
		Value:    value,
	}
}

// Update the underlying VariableGroupCondition resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetVariableGroupConditionResourceData(d *schema.ResourceData, m *models.VariableGroupCondition) {
	d.Set("name", m.Name)
	d.Set("operator", m.Operator)
	d.Set("value", m.Value)
}

// Iterate through and update the VariableGroupCondition resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetVariableGroupConditionSubResourceData(m []*models.VariableGroupCondition) (d []*map[string]interface{}) {
	for _, VariableGroupConditionModel := range m {
		if VariableGroupConditionModel != nil {
			properties := make(map[string]interface{})
			properties["name"] = VariableGroupConditionModel.Name
			properties["operator"] = VariableGroupConditionModel.Operator
			properties["value"] = VariableGroupConditionModel.Value
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the VariableGroupCondition resource defined in the Terraform configuration
func VariableGroupConditionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"operator": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"value": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the VariableGroupCondition resource
func GetVariableGroupConditionPropertyFields() (t []string) {
	return []string{
		"name",
		"operator",
		"value",
	}
}
