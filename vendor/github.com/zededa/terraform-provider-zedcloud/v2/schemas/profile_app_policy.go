package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ProfileAppPolicyModel(d *schema.ResourceData) *models.ProfileAppPolicy {
	var appConfig *models.ProfileAppConfig // ProfileAppConfig
	appConfigInterface, appConfigIsSet := d.GetOk("app_config")
	if appConfigIsSet && appConfigInterface != nil {
		appConfigMap := appConfigInterface.([]interface{})
		if len(appConfigMap) > 0 {
			appConfig = ProfileAppConfigModelFromMap(appConfigMap[0].(map[string]interface{}))
		}
	}
	var metaData *models.ProfilePolicyCommon // ProfilePolicyCommon
	metaDataInterface, metaDataIsSet := d.GetOk("meta_data")
	if metaDataIsSet && metaDataInterface != nil {
		metaDataMap := metaDataInterface.([]interface{})
		if len(metaDataMap) > 0 {
			metaData = ProfilePolicyCommonModelFromMap(metaDataMap[0].(map[string]interface{}))
		}
	}
	return &models.ProfileAppPolicy{
		AppConfig: appConfig,
		MetaData:  metaData,
	}
}

func ProfileAppPolicyModelFromMap(m map[string]interface{}) *models.ProfileAppPolicy {
	var appConfig *models.ProfileAppConfig // ProfileAppConfig
	appConfigInterface, appConfigIsSet := m["app_config"]
	if appConfigIsSet && appConfigInterface != nil {
		appConfigMap := appConfigInterface.([]interface{})
		if len(appConfigMap) > 0 {
			appConfig = ProfileAppConfigModelFromMap(appConfigMap[0].(map[string]interface{}))
		}
	}
	//
	var metaData *models.ProfilePolicyCommon // ProfilePolicyCommon
	metaDataInterface, metaDataIsSet := m["meta_data"]
	if metaDataIsSet && metaDataInterface != nil {
		metaDataMap := metaDataInterface.([]interface{})
		if len(metaDataMap) > 0 {
			metaData = ProfilePolicyCommonModelFromMap(metaDataMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.ProfileAppPolicy{
		AppConfig: appConfig,
		MetaData:  metaData,
	}
}

func SetProfileAppPolicyResourceData(d *schema.ResourceData, m *models.ProfileAppPolicy) {
	d.Set("app_config", SetProfileAppConfigSubResourceData([]*models.ProfileAppConfig{m.AppConfig}))
	d.Set("meta_data", SetProfilePolicyCommonSubResourceData([]*models.ProfilePolicyCommon{m.MetaData}))
}

func SetProfileAppPolicySubResourceData(m []*models.ProfileAppPolicy) (d []*map[string]interface{}) {
	for _, ProfileAppPolicyModel := range m {
		if ProfileAppPolicyModel != nil {
			properties := make(map[string]interface{})
			properties["app_config"] = SetProfileAppConfigSubResourceData([]*models.ProfileAppConfig{ProfileAppPolicyModel.AppConfig})
			properties["meta_data"] = SetProfilePolicyCommonSubResourceData([]*models.ProfilePolicyCommon{ProfileAppPolicyModel.MetaData})
			d = append(d, &properties)
		}
	}
	return
}

func ProfileAppPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"app_config": {
			Description: `app config for app profile`,
			Type:        schema.TypeList, //GoType: ProfileAppConfig
			Elem: &schema.Resource{
				Schema: ProfileAppConfigSchema(),
			},
			Optional: true,
		},

		"meta_data": {
			Description: `all the required metadata for a policy like id, name, different types of tags`,
			Type:        schema.TypeList, //GoType: ProfilePolicyCommon
			Elem: &schema.Resource{
				Schema: ProfilePolicyCommonSchema(),
			},
			Optional: true,
		},
	}
}

func GetProfileAppPolicyPropertyFields() (t []string) {
	return []string{
		"app_config",
		"meta_data",
	}
}
