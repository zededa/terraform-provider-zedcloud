package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AuthTypeModel(d *schema.ResourceData) *models.AuthType {
	authType, _ := d.Get("auth_type").(models.AuthType)
	return &authType
}

func AuthTypeModelFromMap(m map[string]interface{}) *models.AuthType {
	authType := m["auth_type"].(models.AuthType)
	return &authType
}

func SetAuthTypeResourceData(d *schema.ResourceData, m *models.AuthType) {
}

func SetAuthTypeSubResourceData(m []*models.AuthType) (d []*map[string]interface{}) {
	for _, AuthTypeModel := range m {
		if AuthTypeModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func AuthTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func GetAuthTypePropertyFields() (t []string) {
	return []string{}
}
