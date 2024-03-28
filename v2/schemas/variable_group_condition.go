package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

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
	var operator *models.VariableGroupConditionOperator // VariableGroupConditionOperator
	operatorInterface, operatorIsSet := m["operator"]
	if operatorIsSet {
		operatorModel := operatorInterface.(string)
		operator = models.NewVariableGroupConditionOperator(models.VariableGroupConditionOperator(operatorModel))
	}
	value := m["value"].(string)
	return &models.VariableGroupCondition{
		Name:     name,
		Operator: operator,
		Value:    value,
	}
}

func SetVariableGroupConditionResourceData(d *schema.ResourceData, m *models.VariableGroupCondition) {
	d.Set("name", m.Name)
	d.Set("operator", m.Operator)
	d.Set("value", m.Value)
}

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
