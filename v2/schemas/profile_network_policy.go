package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ProfileNetworkPolicyModel(d *schema.ResourceData) *models.ProfileNetworkPolicy {
	var metaData *models.ProfilePolicyCommon // ProfilePolicyCommon
	metaDataInterface, metaDataIsSet := d.GetOk("meta_data")
	if metaDataIsSet && metaDataInterface != nil {
		metaDataMap := metaDataInterface.([]interface{})
		if len(metaDataMap) > 0 {
			metaData = ProfilePolicyCommonModelFromMap(metaDataMap[0].(map[string]interface{}))
		}
	}
	var networkConfig *models.ProfileNetworkConfig // ProfileNetworkConfig
	networkConfigInterface, networkConfigIsSet := d.GetOk("network_config")
	if networkConfigIsSet && networkConfigInterface != nil {
		networkConfigMap := networkConfigInterface.([]interface{})
		if len(networkConfigMap) > 0 {
			networkConfig = ProfileNetworkConfigModelFromMap(networkConfigMap[0].(map[string]interface{}))
		}
	}
	return &models.ProfileNetworkPolicy{
		MetaData:      metaData,
		NetworkConfig: networkConfig,
	}
}

func ProfileNetworkPolicyModelFromMap(m map[string]interface{}) *models.ProfileNetworkPolicy {
	var metaData *models.ProfilePolicyCommon // ProfilePolicyCommon
	metaDataInterface, metaDataIsSet := m["meta_data"]
	if metaDataIsSet && metaDataInterface != nil {
		metaDataMap := metaDataInterface.([]interface{})
		if len(metaDataMap) > 0 {
			metaData = ProfilePolicyCommonModelFromMap(metaDataMap[0].(map[string]interface{}))
		}
	}
	//
	var networkConfig *models.ProfileNetworkConfig // ProfileNetworkConfig
	networkConfigInterface, networkConfigIsSet := m["network_config"]
	if networkConfigIsSet && networkConfigInterface != nil {
		networkConfigMap := networkConfigInterface.([]interface{})
		if len(networkConfigMap) > 0 {
			networkConfig = ProfileNetworkConfigModelFromMap(networkConfigMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.ProfileNetworkPolicy{
		MetaData:      metaData,
		NetworkConfig: networkConfig,
	}
}

func SetProfileNetworkPolicyResourceData(d *schema.ResourceData, m *models.ProfileNetworkPolicy) {
	d.Set("meta_data", SetProfilePolicyCommonSubResourceData([]*models.ProfilePolicyCommon{m.MetaData}))
	d.Set("network_config", SetProfileNetworkConfigSubResourceData([]*models.ProfileNetworkConfig{m.NetworkConfig}))
}

func SetProfileNetworkPolicySubResourceData(m []*models.ProfileNetworkPolicy) (d []*map[string]interface{}) {
	for _, ProfileNetworkPolicyModel := range m {
		if ProfileNetworkPolicyModel != nil {
			properties := make(map[string]interface{})
			properties["meta_data"] = SetProfilePolicyCommonSubResourceData([]*models.ProfilePolicyCommon{ProfileNetworkPolicyModel.MetaData})
			properties["network_config"] = SetProfileNetworkConfigSubResourceData([]*models.ProfileNetworkConfig{ProfileNetworkPolicyModel.NetworkConfig})
			d = append(d, &properties)
		}
	}
	return
}

func ProfileNetworkPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"meta_data": {
			Description: `all the required metadata for a policy like id, name, different types of tags`,
			Type:        schema.TypeList, //GoType: ProfilePolicyCommon
			Elem: &schema.Resource{
				Schema: ProfilePolicyCommonSchema(),
			},
			Optional: true,
		},

		"network_config": {
			Description: `network instance config details`,
			Type:        schema.TypeList, //GoType: ProfileNetworkConfig
			Elem: &schema.Resource{
				Schema: ProfileNetworkConfigSchema(),
			},
			Optional: true,
		},
	}
}

func GetProfileNetworkPolicyPropertyFields() (t []string) {
	return []string{
		"meta_data",
		"network_config",
	}
}
