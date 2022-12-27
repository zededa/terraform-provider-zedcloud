package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate NetworkInstPolicy resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func NetworkInstPolicyModel(d *schema.ResourceData) *models.NetworkInstPolicy {
	var metaData *models.PolicyCommon // PolicyCommon
	metaDataInterface, metaDataIsSet := d.GetOk("meta_data")
	if metaDataIsSet {
		metaDataMap := metaDataInterface.([]interface{})[0].(map[string]interface{})
		metaData = PolicyCommonModelFromMap(metaDataMap)
	}
	var netInstConfig *models.NetworkInstConfig // NetworkInstConfig
	netInstConfigInterface, netInstConfigIsSet := d.GetOk("net_inst_config")
	if netInstConfigIsSet {
		netInstConfigMap := netInstConfigInterface.([]interface{})[0].(map[string]interface{})
		netInstConfig = NetworkInstConfigModelFromMap(netInstConfigMap)
	}
	return &models.NetworkInstPolicy{
		MetaData:      metaData,
		NetInstConfig: netInstConfig,
	}
}

func NetworkInstPolicyModelFromMap(m map[string]interface{}) *models.NetworkInstPolicy {
	var metaData *models.PolicyCommon // PolicyCommon
	metaDataInterface, metaDataIsSet := m["meta_data"]
	if metaDataIsSet {
		metaDataMap := metaDataInterface.([]interface{})[0].(map[string]interface{})
		metaData = PolicyCommonModelFromMap(metaDataMap)
	}
	//
	var netInstConfig *models.NetworkInstConfig // NetworkInstConfig
	netInstConfigInterface, netInstConfigIsSet := m["net_inst_config"]
	if netInstConfigIsSet {
		netInstConfigMap := netInstConfigInterface.([]interface{})[0].(map[string]interface{})
		netInstConfig = NetworkInstConfigModelFromMap(netInstConfigMap)
	}
	//
	return &models.NetworkInstPolicy{
		MetaData:      metaData,
		NetInstConfig: netInstConfig,
	}
}

// Update the underlying NetworkInstPolicy resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetNetworkInstPolicyResourceData(d *schema.ResourceData, m *models.NetworkInstPolicy) {
	d.Set("meta_data", SetPolicyCommonSubResourceData([]*models.PolicyCommon{m.MetaData}))
	d.Set("net_inst_config", SetNetworkInstConfigSubResourceData([]*models.NetworkInstConfig{m.NetInstConfig}))
}

// Iterate throught and update the NetworkInstPolicy resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetNetworkInstPolicySubResourceData(m []*models.NetworkInstPolicy) (d []*map[string]interface{}) {
	for _, NetworkInstPolicyModel := range m {
		if NetworkInstPolicyModel != nil {
			properties := make(map[string]interface{})
			properties["meta_data"] = SetPolicyCommonSubResourceData([]*models.PolicyCommon{NetworkInstPolicyModel.MetaData})
			properties["net_inst_config"] = SetNetworkInstConfigSubResourceData([]*models.NetworkInstConfig{NetworkInstPolicyModel.NetInstConfig})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the NetworkInstPolicy resource defined in the Terraform configuration
func NetworkInstPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"meta_data": {
			Description: `all the required metadata for a policy like id, name, different types of tags`,
			Type:        schema.TypeList, //GoType: PolicyCommon
			Elem: &schema.Resource{
				Schema: PolicyCommonSchema(),
			},
			Optional: true,
		},

		"net_inst_config": {
			Description: `network instance config details`,
			Type:        schema.TypeList, //GoType: NetworkInstConfig
			Elem: &schema.Resource{
				Schema: NetworkInstConfigSchema(),
			},
			Optional: true,
		},
	}
}

// Retrieve property field names for updating the NetworkInstPolicy resource
func GetNetworkInstPolicyPropertyFields() (t []string) {
	return []string{
		"meta_data",
		"net_inst_config",
	}
}
