package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

func CustomConfigVariableGroupModel(d *schema.ResourceData) *models.CustomConfigVariableGroup {
	var condition *models.VariableGroupCondition // VariableGroupCondition
	conditionInterface, conditionIsSet := d.GetOk("condition")
	if conditionIsSet && conditionInterface != nil {
		conditionMap := conditionInterface.([]interface{})
		if len(conditionMap) > 0 {
			condition = VariableGroupConditionModelFromMap(conditionMap[0].(map[string]interface{}))
		}
	}
	name, _ := d.Get("name").(string)
	required, _ := d.Get("required").(bool)
	var variables []*models.VariableGroupVariable // []*VariableGroupVariable
	variablesInterface, variablesIsSet := d.GetOk("variables")
	if variablesIsSet {
		var items []interface{}
		if listItems, isList := variablesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = variablesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := VariableGroupVariableModelFromMap(v.(map[string]interface{}))
			variables = append(variables, m)
		}
	}
	return &models.CustomConfigVariableGroup{
		Condition: condition,
		Name:      name,
		Required:  required,
		Variables: variables,
	}
}

func CustomConfigVariableGroupModelFromMap(m map[string]interface{}) *models.CustomConfigVariableGroup {
	var condition *models.VariableGroupCondition // VariableGroupCondition
	conditionInterface, conditionIsSet := m["condition"]
	if conditionIsSet && conditionInterface != nil {
		conditionMap := conditionInterface.([]interface{})
		if len(conditionMap) > 0 {
			condition = VariableGroupConditionModelFromMap(conditionMap[0].(map[string]interface{}))
		}
	}
	//
	name := m["name"].(string)
	required := m["required"].(bool)
	var variables []*models.VariableGroupVariable // []*VariableGroupVariable
	variablesInterface, variablesIsSet := m["variables"]
	if variablesIsSet {
		var items []interface{}
		if listItems, isList := variablesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = variablesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := VariableGroupVariableModelFromMap(v.(map[string]interface{}))
			variables = append(variables, m)
		}
	}
	return &models.CustomConfigVariableGroup{
		Condition: condition,
		Name:      name,
		Required:  required,
		Variables: variables,
	}
}

func SetCustomConfigVariableGroupResourceData(d *schema.ResourceData, m *models.CustomConfigVariableGroup) {
	d.Set("condition", SetVariableGroupConditionSubResourceData([]*models.VariableGroupCondition{m.Condition}))
	d.Set("name", m.Name)
	d.Set("required", m.Required)
	d.Set("variables", SetVariableGroupVariableSubResourceData(m.Variables))
}

func SetCustomConfigVariableGroupSubResourceData(m []*models.CustomConfigVariableGroup) (d []*map[string]interface{}) {
	for _, CustomConfigVariableGroupModel := range m {
		if CustomConfigVariableGroupModel != nil {
			properties := make(map[string]interface{})
			properties["condition"] = SetVariableGroupConditionSubResourceData([]*models.VariableGroupCondition{CustomConfigVariableGroupModel.Condition})
			properties["name"] = CustomConfigVariableGroupModel.Name
			properties["required"] = CustomConfigVariableGroupModel.Required
			properties["variables"] = SetVariableGroupVariableSubResourceData(CustomConfigVariableGroupModel.Variables)
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the CustomConfigVariableGroup resource defined in the Terraform configuration
func VariableGroup() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"condition": {
			Description: `Condition to apply the variable group. (Optional. Default: None)`,
			Type:        schema.TypeList, //GoType: VariableGroupCondition
			Elem: &schema.Resource{
				Schema: VariableGroupConditionSchema(),
			},
			Optional: true,
		},

		"name": {
			Description: `Name of the Variable Group(Required)`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"required": {
			Description: `Indicates if the variable group is required to be specified for the App Instance. (Optional. Default:False)`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"variables": {
			Description: `List of variables(Required)`,
			Type:        schema.TypeList, //GoType: []*VariableGroupVariable
			Elem: &schema.Resource{
				Schema: VariableGroupVariable(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
	}
}

// Retrieve property field names for updating the CustomConfigVariableGroup resource
func GetCustomConfigVariableGroupPropertyFields() (t []string) {
	return []string{
		"condition",
		"name",
		"required",
		"variables",
	}
}
