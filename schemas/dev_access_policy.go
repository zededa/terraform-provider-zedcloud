package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func DevAccessPolicyModel(d *schema.ResourceData) *models.DevAccessPolicy {
	allowDev, _ := d.Get("allow_dev").(bool)
	return &models.DevAccessPolicy{
		AllowDev: allowDev,
	}
}

func DevAccessPolicyModelFromMap(m map[string]interface{}) *models.DevAccessPolicy {
	allowDev := m["allow_dev"].(bool)
	return &models.DevAccessPolicy{
		AllowDev: allowDev,
	}
}

func SetDevAccessPolicyResourceData(d *schema.ResourceData, m *models.DevAccessPolicy) {
	d.Set("allow_dev", m.AllowDev)
}

func SetDevAccessPolicySubResourceData(m []*models.DevAccessPolicy) (d []*map[string]interface{}) {
	for _, DevAccessPolicyModel := range m {
		if DevAccessPolicyModel != nil {
			properties := make(map[string]interface{})
			properties["allow_dev"] = DevAccessPolicyModel.AllowDev
			d = append(d, &properties)
		}
	}
	return
}

func DevAccessPolicy() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"allow_dev": {
			Description: ``,
			Type:        schema.TypeBool,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the DevAccessPolicy resource
func GetDevAccessPolicyPropertyFields() (t []string) {
	return []string{
		"allow_dev",
	}
}
