package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func UserUsagePerEnterpriseListModel(d *schema.ResourceData) *models.UserUsagePerEnterpriseList {
	var userUsagePerEntp []*models.UserUsagePerEnterprise // []*UserUsagePerEnterprise
	userUsagePerEntpInterface, userUsagePerEntpIsSet := d.GetOk("user_usage_per_entp")
	if userUsagePerEntpIsSet {
		var items []interface{}
		if listItems, isList := userUsagePerEntpInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = userUsagePerEntpInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := UserUsagePerEnterpriseModelFromMap(v.(map[string]interface{}))
			userUsagePerEntp = append(userUsagePerEntp, m)
		}
	}
	return &models.UserUsagePerEnterpriseList{
		UserUsagePerEntp: userUsagePerEntp,
	}
}

func UserUsagePerEnterpriseListModelFromMap(m map[string]interface{}) *models.UserUsagePerEnterpriseList {
	var userUsagePerEntp []*models.UserUsagePerEnterprise // []*UserUsagePerEnterprise
	userUsagePerEntpInterface, userUsagePerEntpIsSet := m["user_usage_per_entp"]
	if userUsagePerEntpIsSet {
		var items []interface{}
		if listItems, isList := userUsagePerEntpInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = userUsagePerEntpInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := UserUsagePerEnterpriseModelFromMap(v.(map[string]interface{}))
			userUsagePerEntp = append(userUsagePerEntp, m)
		}
	}
	return &models.UserUsagePerEnterpriseList{
		UserUsagePerEntp: userUsagePerEntp,
	}
}

func SetUserUsagePerEnterpriseListResourceData(d *schema.ResourceData, m *models.UserUsagePerEnterpriseList) {
	d.Set("user_usage_per_entp", SetUserUsagePerEnterpriseSubResourceData(m.UserUsagePerEntp))
}

func SetUserUsagePerEnterpriseListSubResourceData(m []*models.UserUsagePerEnterpriseList) (d []*map[string]interface{}) {
	for _, UserUsagePerEnterpriseListModel := range m {
		if UserUsagePerEnterpriseListModel != nil {
			properties := make(map[string]interface{})
			properties["user_usage_per_entp"] = SetUserUsagePerEnterpriseSubResourceData(UserUsagePerEnterpriseListModel.UserUsagePerEntp)
			d = append(d, &properties)
		}
	}
	return
}

func UserUsagePerEnterpriseListSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"user_usage_per_entp": {
			Description: `user usage per enterprise list`,
			Type:        schema.TypeList, //GoType: []*UserUsagePerEnterprise
			Elem: &schema.Resource{
				Schema: UserUsagePerEnterpriseSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
	}
}

func GetUserUsagePerEnterpriseListPropertyFields() (t []string) {
	return []string{
		"user_usage_per_entp",
	}
}
