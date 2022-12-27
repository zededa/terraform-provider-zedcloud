package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate CustomConfigVariableGroup resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func CustomConfigVariableGroupModel(d *schema.ResourceData) *models.CustomConfigVariableGroup {
	var condition *models.VariableGroupCondition // VariableGroupCondition
	conditionInterface, conditionIsSet := d.GetOk("condition")
	if conditionIsSet {
		conditionMap := conditionInterface.([]interface{})[0].(map[string]interface{})
		condition = VariableGroupConditionModelFromMap(conditionMap)
	}
	name := d.Get("name").(string)
	required := d.Get("required").(bool)
	variables := d.Get("variables").([]*models.VariableGroupVariable) // []*VariableGroupVariable
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
	if conditionIsSet {
		conditionMap := conditionInterface.([]interface{})[0].(map[string]interface{})
		condition = VariableGroupConditionModelFromMap(conditionMap)
	}
	//
	name := m["name"].(string)
	required := m["required"].(bool)
	variables := m["variables"].([]*models.VariableGroupVariable) // []*VariableGroupVariable
	return &models.CustomConfigVariableGroup{
		Condition: condition,
		Name:      name,
		Required:  required,
		Variables: variables,
	}
}

// Update the underlying CustomConfigVariableGroup resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetCustomConfigVariableGroupResourceData(d *schema.ResourceData, m *models.CustomConfigVariableGroup) {
	d.Set("condition", SetVariableGroupConditionSubResourceData([]*models.VariableGroupCondition{m.Condition}))
	d.Set("name", m.Name)
	d.Set("required", m.Required)
	d.Set("variables", SetVariableGroupVariableSubResourceData(m.Variables))
}

// Iterate throught and update the CustomConfigVariableGroup resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
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
func CustomConfigVariableGroupSchema() map[string]*schema.Schema {
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
				Schema: VariableGroupVariableSchema(),
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
