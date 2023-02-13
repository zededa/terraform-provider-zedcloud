package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate VariableGroupConditionOperator resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func VariableGroupConditionOperatorModel(d *schema.ResourceData) *models.VariableGroupConditionOperator {
	variableGroupConditionOperator, _ := d.Get("variable_group_condition_operator").(models.VariableGroupConditionOperator)
	return &variableGroupConditionOperator
}

func VariableGroupConditionOperatorModelFromMap(m map[string]interface{}) *models.VariableGroupConditionOperator {
	variableGroupConditionOperator := m["variable_group_condition_operator"].(models.VariableGroupConditionOperator)
	return &variableGroupConditionOperator
}

// Update the underlying VariableGroupConditionOperator resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetVariableGroupConditionOperatorResourceData(d *schema.ResourceData, m *models.VariableGroupConditionOperator) {
}

// Iterate through and update the VariableGroupConditionOperator resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetVariableGroupConditionOperatorSubResourceData(m []*models.VariableGroupConditionOperator) (d []*map[string]interface{}) {
	for _, VariableGroupConditionOperatorModel := range m {
		if VariableGroupConditionOperatorModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the VariableGroupConditionOperator resource defined in the Terraform configuration
func VariableGroupConditionOperatorSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the VariableGroupConditionOperator resource
func GetVariableGroupConditionOperatorPropertyFields() (t []string) {
	return []string{}
}
