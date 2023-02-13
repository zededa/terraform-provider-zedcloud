package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func TransDetailsModel(d *schema.ResourceData) *models.TransDetails {
	var cause *models.TransCause // TransCause
	causeInterface, causeIsSet := d.GetOk("cause")
	if causeIsSet {
		causeModel := causeInterface.(string)
		cause = models.NewTransCause(models.TransCause(causeModel))
	}
	var scope *models.ActionScope // ActionScope
	scopeInterface, scopeIsSet := d.GetOk("scope")
	if scopeIsSet && scopeInterface != nil {
		scopeMap := scopeInterface.([]interface{})
		if len(scopeMap) > 0 {
			scope = ActionScopeModelFromMap(scopeMap[0].(map[string]interface{}))
		}
	}
	severity, _ := d.Get("severity").(string)
	return &models.TransDetails{
		Cause:    cause,
		Scope:    scope,
		Severity: severity,
	}
}

func TransDetailsModelFromMap(m map[string]interface{}) *models.TransDetails {
	var cause *models.TransCause // TransCause
	causeInterface, causeIsSet := m["cause"]
	if causeIsSet {
		causeModel := causeInterface.(string)
		cause = models.NewTransCause(models.TransCause(causeModel))
	}
	var scope *models.ActionScope // ActionScope
	scopeInterface, scopeIsSet := m["scope"]
	if scopeIsSet && scopeInterface != nil {
		scopeMap := scopeInterface.([]interface{})
		if len(scopeMap) > 0 {
			scope = ActionScopeModelFromMap(scopeMap[0].(map[string]interface{}))
		}
	}
	//
	severity := m["severity"].(string)
	return &models.TransDetails{
		Cause:    cause,
		Scope:    scope,
		Severity: severity,
	}
}

// Update the underlying TransDetails resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetTransDetailsResourceData(d *schema.ResourceData, m *models.TransDetails) {
	d.Set("cause", m.Cause)
	d.Set("scope", SetActionScopeSubResourceData([]*models.ActionScope{m.Scope}))
	d.Set("severity", m.Severity)
}

// Iterate through and update the TransDetails resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetTransDetailsSubResourceData(m []*models.TransDetails) (d []*map[string]interface{}) {
	for _, TransDetailsModel := range m {
		if TransDetailsModel != nil {
			properties := make(map[string]interface{})
			properties["cause"] = TransDetailsModel.Cause
			properties["scope"] = SetActionScopeSubResourceData([]*models.ActionScope{TransDetailsModel.Scope})
			properties["severity"] = TransDetailsModel.Severity
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the TransDetails resource defined in the Terraform configuration
func TransDetailsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cause": {
			Description: `Cause of the transition action`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"scope": {
			Description: `Scope of the action (e.g. project, cluster, instance)`,
			Type:        schema.TypeList, //GoType: ActionScope
			Elem: &schema.Resource{
				Schema: ActionScopeSchema(),
			},
			Optional: true,
		},

		"severity": {
			Description: `Urgency of the recommended action (ERROR - must take action immediately, WARN - must take action but not immediately, NOTICE - may take action`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the TransDetails resource
func GetTransDetailsPropertyFields() (t []string) {
	return []string{
		"cause",
		"scope",
		"severity",
	}
}
