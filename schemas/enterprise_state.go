package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func EnterpriseStateModel(d *schema.ResourceData) *models.EnterpriseState {
	enterpriseState, _ := d.Get("enterprise_state").(models.EnterpriseState)
	return &enterpriseState
}

func EnterpriseStateModelFromMap(m map[string]interface{}) *models.EnterpriseState {
	enterpriseState := m["enterprise_state"].(models.EnterpriseState)
	return &enterpriseState
}

func SetEnterpriseStateResourceData(d *schema.ResourceData, m *models.EnterpriseState) {
}

func SetEnterpriseStateSubResourceData(m []*models.EnterpriseState) (d []*map[string]interface{}) {
	for _, EnterpriseStateModel := range m {
		if EnterpriseStateModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func EnterpriseStateSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func GetEnterpriseStatePropertyFields() (t []string) {
	return []string{}
}
