package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AppAccessPolicyModel(d *schema.ResourceData) *models.AppAccessPolicy {
	allowApp, _ := d.Get("allow_app").(bool)
	return &models.AppAccessPolicy{
		AllowApp: allowApp,
	}
}

func AppAccessPolicyModelFromMap(m map[string]interface{}) *models.AppAccessPolicy {
	allowApp := m["allow_app"].(bool)
	return &models.AppAccessPolicy{
		AllowApp: allowApp,
	}
}

func SetAppAccessPolicyResourceData(d *schema.ResourceData, m *models.AppAccessPolicy) {
	d.Set("allow_app", m.AllowApp)
}

func SetAppAccessPolicySubResourceData(m []*models.AppAccessPolicy) (d []*map[string]interface{}) {
	for _, AppAccessPolicyModel := range m {
		if AppAccessPolicyModel != nil {
			properties := make(map[string]interface{})
			properties["allow_app"] = AppAccessPolicyModel.AllowApp
			d = append(d, &properties)
		}
	}
	return
}

func AppAccessPolicy() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"allow_app": {
			Description: ``,
			Type:        schema.TypeBool,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the AppAccessPolicy resource
func GetAppAccessPolicyPropertyFields() (t []string) {
	return []string{
		"allow_app",
	}
}
