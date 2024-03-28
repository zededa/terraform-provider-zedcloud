package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func UserTemplateModel(d *schema.ResourceData) *models.UserDataTemplate {
	var customConfig *models.CustomConfig // CustomConfig
	customConfigInterface, customConfigIsSet := d.GetOk("custom_config")
	if customConfigIsSet {
		customConfigMap := customConfigInterface.([]interface{})[0].(map[string]interface{})
		customConfig = CustomConfigModelFromMap(customConfigMap)
	}
	return &models.UserDataTemplate{
		CustomConfig: customConfig,
	}
}

func UserTemplateModelFromMap(m map[string]interface{}) *models.UserDataTemplate {
	var customConfig *models.CustomConfig // CustomConfig
	customConfigInterface, customConfigIsSet := m["custom_config"]
	if customConfigIsSet {
		customConfigMap := customConfigInterface.([]interface{})[0].(map[string]interface{})
		customConfig = CustomConfigModelFromMap(customConfigMap)
	}
	//
	return &models.UserDataTemplate{
		CustomConfig: customConfig,
	}
}

func SetUserTemplateResourceData(d *schema.ResourceData, m *models.UserDataTemplate) {
	d.Set("custom_config", SetCustomConfigSubResourceData([]*models.CustomConfig{m.CustomConfig}))
}

func SetUserTemplateSubResourceData(m []*models.UserDataTemplate) (d []*map[string]interface{}) {
	for _, UserDataTemplateModel := range m {
		if UserDataTemplateModel != nil {
			properties := make(map[string]interface{})
			properties["custom_config"] = SetCustomConfigSubResourceData([]*models.CustomConfig{UserDataTemplateModel.CustomConfig})
			d = append(d, &properties)
		}
	}
	return
}

func UserTemplate() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"custom_config": {
			Description: ``,
			Type:        schema.TypeList, //GoType: CustomConfig
			Elem: &schema.Resource{
				Schema: CustomConfig(),
			},
			Optional: true,
		},
	}
}

func GetUserTemplatePropertyFields() (t []string) {
	return []string{
		"custom_config",
	}
}
