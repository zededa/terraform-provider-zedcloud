package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate AppInstPolicy resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func AppInstPolicyModel(d *schema.ResourceData) *models.AppInstPolicy {
	var appInstConfig *models.AppInstConfig // AppInstConfig
	appInstConfigInterface, appInstConfigIsSet := d.GetOk("app_inst_config")
	if appInstConfigIsSet {
		appInstConfigMap := appInstConfigInterface.([]interface{})[0].(map[string]interface{})
		appInstConfig = AppInstConfigModelFromMap(appInstConfigMap)
	}
	var metaData *models.PolicyCommon // PolicyCommon
	metaDataInterface, metaDataIsSet := d.GetOk("meta_data")
	if metaDataIsSet {
		metaDataMap := metaDataInterface.([]interface{})[0].(map[string]interface{})
		metaData = PolicyCommonModelFromMap(metaDataMap)
	}
	return &models.AppInstPolicy{
		AppInstConfig: appInstConfig,
		MetaData:      metaData,
	}
}

func AppInstPolicyModelFromMap(m map[string]interface{}) *models.AppInstPolicy {
	var appInstConfig *models.AppInstConfig // AppInstConfig
	appInstConfigInterface, appInstConfigIsSet := m["app_inst_config"]
	if appInstConfigIsSet {
		appInstConfigMap := appInstConfigInterface.([]interface{})[0].(map[string]interface{})
		appInstConfig = AppInstConfigModelFromMap(appInstConfigMap)
	}
	//
	var metaData *models.PolicyCommon // PolicyCommon
	metaDataInterface, metaDataIsSet := m["meta_data"]
	if metaDataIsSet {
		metaDataMap := metaDataInterface.([]interface{})[0].(map[string]interface{})
		metaData = PolicyCommonModelFromMap(metaDataMap)
	}
	//
	return &models.AppInstPolicy{
		AppInstConfig: appInstConfig,
		MetaData:      metaData,
	}
}

// Update the underlying AppInstPolicy resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetAppInstPolicyResourceData(d *schema.ResourceData, m *models.AppInstPolicy) {
	d.Set("app_inst_config", SetAppInstConfigSubResourceData([]*models.AppInstConfig{m.AppInstConfig}))
	d.Set("meta_data", SetPolicyCommonSubResourceData([]*models.PolicyCommon{m.MetaData}))
}

// Iterate through and update the AppInstPolicy resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
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

// Schema mapping representing the AppInstPolicy resource defined in the Terraform configuration
func AppInstPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"app_inst_config": {
			Description: `app instance config for automated deployment`,
			Type:        schema.TypeList, //GoType: AppInstConfig
			Elem: &schema.Resource{
				Schema: AppInstConfigSchema(),
			},
			Optional: true,
		},

		"meta_data": {
			Description: `all the required metadata for a policy like id, name, different types of tags`,
			Type:        schema.TypeList, //GoType: PolicyCommon
			Elem: &schema.Resource{
				Schema: PolicyCommonSchema(),
			},
			Optional: true,
		},
	}
}

// Retrieve property field names for updating the AppInstPolicy resource
func GetAppInstPolicyPropertyFields() (t []string) {
	return []string{
		"app_inst_config",
		"meta_data",
	}
}
