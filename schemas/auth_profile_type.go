package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AuthProfileTypeModel(d *schema.ResourceData) *models.AuthProfileType {
	authProfileType, _ := d.Get("auth_profile_type").(models.AuthProfileType)
	return &authProfileType
}

func AuthProfileTypeModelFromMap(m map[string]interface{}) *models.AuthProfileType {
	authProfileType := m["auth_profile_type"].(models.AuthProfileType)
	return &authProfileType
}

func SetAuthProfileTypeResourceData(d *schema.ResourceData, m *models.AuthProfileType) {
}

func SetAuthProfileTypeSubResourceData(m []*models.AuthProfileType) (d []*map[string]interface{}) {
	for _, AuthProfileTypeModel := range m {
		if AuthProfileTypeModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func AuthProfileTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func GetAuthProfileTypePropertyFields() (t []string) {
	return []string{}
}
