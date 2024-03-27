package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

func NetworkPolicyModel(d *schema.ResourceData) *models.NetworkPolicy {
	var netInstanceConfig []*models.NetworkInstance // []*NetworkInstance
	netInstanceConfigInterface, netInstanceConfigIsSet := d.GetOk("net_instance_config")
	if netInstanceConfigIsSet {
		var items []interface{}
		if listItems, isList := netInstanceConfigInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = netInstanceConfigInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := NetworkInstanceModelFromMap(v.(map[string]interface{}))
			netInstanceConfig = append(netInstanceConfig, m)
		}
	}
	return &models.NetworkPolicy{
		NetInstanceConfig: netInstanceConfig,
	}
}

func NetworkPolicyModelFromMap(m map[string]interface{}) *models.NetworkPolicy {
	var netInstanceConfig []*models.NetworkInstance // []*NetworkInstance
	netInstanceConfigInterface, netInstanceConfigIsSet := m["net_instance_config"]
	if netInstanceConfigIsSet {
		var items []interface{}
		if listItems, isList := netInstanceConfigInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = netInstanceConfigInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := NetworkInstanceModelFromMap(v.(map[string]interface{}))
			netInstanceConfig = append(netInstanceConfig, m)
		}
	}
	return &models.NetworkPolicy{
		NetInstanceConfig: netInstanceConfig,
	}
}

func SetNetworkPolicyResourceData(d *schema.ResourceData, m *models.NetworkPolicy) {
	d.Set("net_instance_config", SetNetworkInstanceSubResourceData(m.NetInstanceConfig))
}

func SetNetworkPolicySubResourceData(m []*models.NetworkPolicy) (d []*map[string]interface{}) {
	for _, NetworkPolicyModel := range m {
		if NetworkPolicyModel != nil {
			properties := make(map[string]interface{})
			properties["net_instance_config"] = SetNetworkInstanceSubResourceData(NetworkPolicyModel.NetInstanceConfig)
			d = append(d, &properties)
		}
	}
	return
}

func NetworkPolicy() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"net_instance_config": {
			Description: `list of network details that will be created on all the devices of the project to which this policy is attached`,
			Type:        schema.TypeList, //GoType: []*NetworkInstance
			Elem: &schema.Resource{
				Schema: NetworkInstance(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Required: true,
		},
	}
}

func GetNetworkPolicyPropertyFields() (t []string) {
	return []string{
		"net_instance_config",
	}
}
