package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func AppInstPolicyModel(d *schema.ResourceData) *models.AppInstPolicy {
	var appInstConfig *models.AppInstConfig // AppInstConfig
	appInstConfigInterface, appInstConfigIsSet := d.GetOk("app_inst_config")
	if appInstConfigIsSet && appInstConfigInterface != nil {
		appInstConfigMap := appInstConfigInterface.([]interface{})
		if len(appInstConfigMap) > 0 {
			appInstConfig = AppInstConfigModelFromMap(appInstConfigMap[0].(map[string]interface{}))
		}
	}
	var metaData *models.PolicyCommon // PolicyCommon
	metaDataInterface, metaDataIsSet := d.GetOk("meta_data")
	if metaDataIsSet && metaDataInterface != nil {
		metaDataMap := metaDataInterface.([]interface{})
		if len(metaDataMap) > 0 {
			metaData = PolicyCommonModelFromMap(metaDataMap[0].(map[string]interface{}))
		}
	}
	return &models.AppInstPolicy{
		AppInstConfig: appInstConfig,
		MetaData:      metaData,
	}
}

func AppInstPolicyModelFromMap(m map[string]interface{}) *models.AppInstPolicy {
	var appInstConfig *models.AppInstConfig // AppInstConfig
	appInstConfigInterface, appInstConfigIsSet := m["app_inst_config"]
	if appInstConfigIsSet && appInstConfigInterface != nil {
		appInstConfigMap := appInstConfigInterface.([]interface{})
		if len(appInstConfigMap) > 0 {
			appInstConfig = AppInstConfigModelFromMap(appInstConfigMap[0].(map[string]interface{}))
		}
	}
	//
	var metaData *models.PolicyCommon // PolicyCommon
	metaDataInterface, metaDataIsSet := m["meta_data"]
	if metaDataIsSet && metaDataInterface != nil {
		metaDataMap := metaDataInterface.([]interface{})
		if len(metaDataMap) > 0 {
			metaData = PolicyCommonModelFromMap(metaDataMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.AppInstPolicy{
		AppInstConfig: appInstConfig,
		MetaData:      metaData,
	}
}

func SetAppInstPolicyResourceData(d *schema.ResourceData, m *models.AppInstPolicy) {
	d.Set("app_inst_config", SetAppInstConfigSubResourceData([]*models.AppInstConfig{m.AppInstConfig}))
	d.Set("meta_data", SetPolicyCommonSubResourceData([]*models.PolicyCommon{m.MetaData}))
}

func SetAppInstPolicySubResourceData(m []*models.AppInstPolicy) (d []*map[string]interface{}) {
	for _, AppInstPolicyModel := range m {
		if AppInstPolicyModel != nil {
			properties := make(map[string]interface{})
			properties["app_inst_config"] = SetAppInstConfigSubResourceData([]*models.AppInstConfig{AppInstPolicyModel.AppInstConfig})
			properties["meta_data"] = SetPolicyCommonSubResourceData([]*models.PolicyCommon{AppInstPolicyModel.MetaData})
			d = append(d, &properties)
		}
	}
	return
}

func AppInstPolicy() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"app_inst_config": {
			Description: `app instance config for automated deployment`,
			Type:        schema.TypeList, //GoType: AppInstConfig
			Elem: &schema.Resource{
				Schema: AppInstConfig(),
			},
			Optional: true,
		},

		"meta_data": {
			Description: `all the required metadata for a policy like id, name, different types of tags`,
			Type:        schema.TypeList, //GoType: PolicyCommon
			Elem: &schema.Resource{
				Schema: PolicyCommon(),
			},
			Optional: true,
		},
	}
}

func GetAppInstPolicyPropertyFields() (t []string) {
	return []string{
		"app_inst_config",
		"meta_data",
	}
}
