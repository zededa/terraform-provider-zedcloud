package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func UserRoleModel(d *schema.ResourceData) *models.UserRole {
	userRole, _ := d.Get("user_role").(models.UserRole)
	return &userRole
}

func UserRoleModelFromMap(m map[string]interface{}) *models.UserRole {
	userRole := m["user_role"].(models.UserRole)
	return &userRole
}

func SetUserRoleResourceData(d *schema.ResourceData, m *models.UserRole) {
}

func SetUserRoleSubResourceData(m []*models.UserRole) (d []*map[string]interface{}) {
	for _, UserRoleModel := range m {
		if UserRoleModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func UserRoleSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func GetUserRolePropertyFields() (t []string) {
	return []string{}
}
