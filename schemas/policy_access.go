package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func PolicyAccessModel(d *schema.ResourceData) *models.PolicyAccess {
	policyAccess, _ := d.Get("policy_access").(models.PolicyAccess)
	return &policyAccess
}

func PolicyAccessModelFromMap(m map[string]interface{}) *models.PolicyAccess {
	policyAccess := m["policy_access"].(models.PolicyAccess)
	return &policyAccess
}

func SetPolicyAccessResourceData(d *schema.ResourceData, m *models.PolicyAccess) {
}

func SetPolicyAccessSubResourceData(m []*models.PolicyAccess) (d []*map[string]interface{}) {
	for _, PolicyAccessModel := range m {
		if PolicyAccessModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func PolicyAccessSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func GetPolicyAccessPropertyFields() (t []string) {
	return []string{}
}
