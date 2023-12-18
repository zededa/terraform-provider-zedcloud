package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func RoleStateModel(d *schema.ResourceData) *models.RoleState {
	roleState, _ := d.Get("role_state").(models.RoleState)
	return &roleState
}

func RoleStateModelFromMap(m map[string]interface{}) *models.RoleState {
	roleState := m["role_state"].(models.RoleState)
	return &roleState
}

func SetRoleStateResourceData(d *schema.ResourceData, m *models.RoleState) {
}

func SetRoleStateSubResourceData(m []*models.RoleState) (d []*map[string]interface{}) {
	for _, RoleStateModel := range m {
		if RoleStateModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func RoleStateSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func GetRoleStatePropertyFields() (t []string) {
	return []string{}
}
