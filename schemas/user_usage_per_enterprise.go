package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func UserUsagePerEnterpriseModel(d *schema.ResourceData) *models.UserUsagePerEnterprise {
	userUsage, _ := d.Get("user_usage").(string)
	return &models.UserUsagePerEnterprise{
		UserUsage: userUsage,
	}
}

func UserUsagePerEnterpriseModelFromMap(m map[string]interface{}) *models.UserUsagePerEnterprise {
	userUsage := m["user_usage"].(string)
	return &models.UserUsagePerEnterprise{
		UserUsage: userUsage,
	}
}

func SetUserUsagePerEnterpriseResourceData(d *schema.ResourceData, m *models.UserUsagePerEnterprise) {
	d.Set("entp_id", m.EntpID)
	d.Set("user_usage", m.UserUsage)
}

func SetUserUsagePerEnterpriseSubResourceData(m []*models.UserUsagePerEnterprise) (d []*map[string]interface{}) {
	for _, UserUsagePerEnterpriseModel := range m {
		if UserUsagePerEnterpriseModel != nil {
			properties := make(map[string]interface{})
			properties["entp_id"] = UserUsagePerEnterpriseModel.EntpID
			properties["user_usage"] = UserUsagePerEnterpriseModel.UserUsage
			d = append(d, &properties)
		}
	}
	return
}

func UserUsagePerEnterpriseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"entp_id": {
			Description: `Enterprise id`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"user_usage": {
			Description: `User usage for that enterprise`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetUserUsagePerEnterprisePropertyFields() (t []string) {
	return []string{
		"user_usage",
	}
}
