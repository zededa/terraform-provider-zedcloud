package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func PolicyScopeModel(d *schema.ResourceData) *models.PolicyScope {
	policyScope, _ := d.Get("policy_scope").(models.PolicyScope)
	return &policyScope
}

func PolicyScopeModelFromMap(m map[string]interface{}) *models.PolicyScope {
	policyScope := m["policy_scope"].(models.PolicyScope)
	return &policyScope
}

func SetPolicyScopeResourceData(d *schema.ResourceData, m *models.PolicyScope) {
}

func SetPolicyScopeSubResourceData(m []*models.PolicyScope) (d []*map[string]interface{}) {
	for _, PolicyScopeModel := range m {
		if PolicyScopeModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func PolicyScopeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func GetPolicyScopePropertyFields() (t []string) {
	return []string{}
}
