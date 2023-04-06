package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func PermissionModel(d *schema.ResourceData) *models.Permission {
	permission, _ := d.Get("permission").(models.Permission)
	return &permission
}

func PermissionModelFromMap(m map[string]interface{}) *models.Permission {
	permission := m["permission"].(models.Permission)
	return &permission
}

func SetPermissionResourceData(d *schema.ResourceData, m *models.Permission) {
}

func SetPermissionSubResourceData(m []*models.Permission) (d []*map[string]interface{}) {
	for _, PermissionModel := range m {
		if PermissionModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func Permission() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func GetPermissionPropertyFields() (t []string) {
	return []string{}
}
