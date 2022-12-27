package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate NetworkPolicy resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func NetworkPolicyModel(d *schema.ResourceData) *models.NetworkPolicy {
	netInstanceConfig := d.Get("net_instance_config").([]*models.NetInstConfig) // []*NetInstConfig
	return &models.NetworkPolicy{
		NetInstanceConfig: netInstanceConfig,
	}
}

func NetworkPolicyModelFromMap(m map[string]interface{}) *models.NetworkPolicy {
	netInstanceConfig := m["net_instance_config"].([]*models.NetInstConfig) // []*NetInstConfig
	return &models.NetworkPolicy{
		NetInstanceConfig: netInstanceConfig,
	}
}

// Update the underlying NetworkPolicy resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetNetworkPolicyResourceData(d *schema.ResourceData, m *models.NetworkPolicy) {
	d.Set("net_instance_config", SetNetInstConfigSubResourceData(m.NetInstanceConfig))
}

// Iterate throught and update the NetworkPolicy resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetNetworkPolicySubResourceData(m []*models.NetworkPolicy) (d []*map[string]interface{}) {
	for _, NetworkPolicyModel := range m {
		if NetworkPolicyModel != nil {
			properties := make(map[string]interface{})
			properties["net_instance_config"] = SetNetInstConfigSubResourceData(NetworkPolicyModel.NetInstanceConfig)
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the NetworkPolicy resource defined in the Terraform configuration
func NetworkPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"net_instance_config": {
			Description: `list of network details that will be created on all the devices of the project to which this policy is attached`,
			Type:        schema.TypeList, //GoType: []*NetInstConfig
			Elem: &schema.Resource{
				Schema: NetInstConfigSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Required: true,
		},
	}
}

// Retrieve property field names for updating the NetworkPolicy resource
func GetNetworkPolicyPropertyFields() (t []string) {
	return []string{
		"net_instance_config",
	}
}
