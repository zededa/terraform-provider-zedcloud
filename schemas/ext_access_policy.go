package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

func ExtAccessPolicyModel(d *schema.ResourceData) *models.ExtAccessPolicy {
	allowExt, _ := d.Get("allow_ext").(bool)
	return &models.ExtAccessPolicy{
		AllowExt: allowExt,
	}
}

func ExtAccessPolicyModelFromMap(m map[string]interface{}) *models.ExtAccessPolicy {
	allowExt := m["allow_ext"].(bool)
	return &models.ExtAccessPolicy{
		AllowExt: allowExt,
	}
}

func SetExtAccessPolicyResourceData(d *schema.ResourceData, m *models.ExtAccessPolicy) {
	d.Set("allow_ext", m.AllowExt)
}

func SetExtAccessPolicySubResourceData(m []*models.ExtAccessPolicy) (d []*map[string]interface{}) {
	for _, ExtAccessPolicyModel := range m {
		if ExtAccessPolicyModel != nil {
			properties := make(map[string]interface{})
			properties["allow_ext"] = ExtAccessPolicyModel.AllowExt
			d = append(d, &properties)
		}
	}
	return
}

func ExtAccessPolicy() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"allow_ext": {
			Description: ``,
			Type:        schema.TypeBool,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the ExtAccessPolicy resource
func GetExtAccessPolicyPropertyFields() (t []string) {
	return []string{
		"allow_ext",
	}
}
