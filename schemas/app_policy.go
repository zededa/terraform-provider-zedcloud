package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

func AppPolicyModel(d *schema.ResourceData) *models.AppPolicy {
	var apps []*models.AppConfig // []*AppConfig
	appsInterface, appsIsSet := d.GetOk("apps")
	if appsIsSet {
		var items []interface{}
		if listItems, isList := appsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = appsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := AppConfigModelFromMap(v.(map[string]interface{}))
			apps = append(apps, m)
		}
	}
	return &models.AppPolicy{
		Apps: apps,
	}
}

func AppPolicyModelFromMap(m map[string]interface{}) *models.AppPolicy {
	var apps []*models.AppConfig // []*AppConfig
	appsInterface, appsIsSet := m["apps"]
	if appsIsSet {
		var items []interface{}
		if listItems, isList := appsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = appsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := AppConfigModelFromMap(v.(map[string]interface{}))
			apps = append(apps, m)
		}
	}
	return &models.AppPolicy{
		Apps: apps,
	}
}

func SetAppPolicyResourceData(d *schema.ResourceData, m *models.AppPolicy) {
	d.Set("apps", SetAppConfigSubResourceData(m.Apps))
}

func SetAppPolicySubResourceData(m []*models.AppPolicy) (d []*map[string]interface{}) {
	for _, AppPolicyModel := range m {
		if AppPolicyModel != nil {
			properties := make(map[string]interface{})
			properties["apps"] = SetAppConfigSubResourceData(AppPolicyModel.Apps)
			d = append(d, &properties)
		}
	}
	return
}

func AppPolicy() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"apps": {
			Description: `list of app details that will be provisioned on all the devices of the project to which this policy is attached`,
			Type:        schema.TypeList, //GoType: []*AppConfig
			Elem: &schema.Resource{
				Schema: AppConfig(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Required: true,
		},
	}
}

func GetAppPolicyPropertyFields() (t []string) {
	return []string{
		"apps",
	}
}
