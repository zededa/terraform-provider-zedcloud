package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func NetworkInstPolicyModel(d *schema.ResourceData) *models.NetworkInstPolicy {
	var metaData *models.PolicyCommon // PolicyCommon
	metaDataInterface, metaDataIsSet := d.GetOk("meta_data")
	if metaDataIsSet && metaDataInterface != nil {
		metaDataMap := metaDataInterface.([]interface{})
		if len(metaDataMap) > 0 {
			metaData = PolicyCommonModelFromMap(metaDataMap[0].(map[string]interface{}))
		}
	}
	var netInstConfig *models.NetworkInstance // NetworkInstConfig
	netInstConfigInterface, netInstConfigIsSet := d.GetOk("net_inst_config")
	if netInstConfigIsSet && netInstConfigInterface != nil {
		netInstConfigMap := netInstConfigInterface.([]interface{})
		if len(netInstConfigMap) > 0 {
			netInstConfig = NetworkInstanceModelFromMap(netInstConfigMap[0].(map[string]interface{}))
		}
	}
	return &models.NetworkInstPolicy{
		MetaData:      metaData,
		NetInstConfig: netInstConfig,
	}
}

func NetworkInstPolicyModelFromMap(m map[string]interface{}) *models.NetworkInstPolicy {
	var metaData *models.PolicyCommon // PolicyCommon
	metaDataInterface, metaDataIsSet := m["meta_data"]
	if metaDataIsSet && metaDataInterface != nil {
		metaDataMap := metaDataInterface.([]interface{})
		if len(metaDataMap) > 0 {
			metaData = PolicyCommonModelFromMap(metaDataMap[0].(map[string]interface{}))
		}
	}
	//
	var netInstConfig *models.NetworkInstance // NetworkInstConfig
	netInstConfigInterface, netInstConfigIsSet := m["net_inst_config"]
	if netInstConfigIsSet && netInstConfigInterface != nil {
		netInstConfigMap := netInstConfigInterface.([]interface{})
		if len(netInstConfigMap) > 0 {
			netInstConfig = NetworkInstanceModelFromMap(netInstConfigMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.NetworkInstPolicy{
		MetaData:      metaData,
		NetInstConfig: netInstConfig,
	}
}

func SetNetworkInstPolicyResourceData(d *schema.ResourceData, m *models.NetworkInstPolicy) {
	d.Set("meta_data", SetPolicyCommonSubResourceData([]*models.PolicyCommon{m.MetaData}))
	d.Set("net_inst_config", SetNetworkInstanceSubResourceData([]*models.NetworkInstance{m.NetInstConfig}))
}

func SetNetworkInstPolicySubResourceData(m []*models.NetworkInstPolicy) (d []*map[string]interface{}) {
	for _, NetworkInstPolicyModel := range m {
		if NetworkInstPolicyModel != nil {
			properties := make(map[string]interface{})
			properties["meta_data"] = SetPolicyCommonSubResourceData([]*models.PolicyCommon{NetworkInstPolicyModel.MetaData})
			properties["net_inst_config"] = SetNetworkInstanceSubResourceData([]*models.NetworkInstance{NetworkInstPolicyModel.NetInstConfig})
			d = append(d, &properties)
		}
	}
	return
}

func NetworkInstPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"meta_data": {
			Description: `all the required metadata for a policy like id, name, different types of tags`,
			Type:        schema.TypeList, //GoType: PolicyCommon
			Elem: &schema.Resource{
				Schema: PolicyCommon(),
			},
			Optional: true,
		},

		"net_inst_config": {
			Description: `network instance config details`,
			Type:        schema.TypeList, //GoType: NetworkInstConfig
			Elem: &schema.Resource{
				Schema: NetworkInstance(),
			},
			Optional: true,
		},
	}
}

func GetNetworkInstPolicyPropertyFields() (t []string) {
	return []string{
		"meta_data",
		"net_inst_config",
	}
}
