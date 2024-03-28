package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func NetworkProxyProtoModel(d *schema.ResourceData) *models.Proto {
	networkProxyProto, _ := d.Get("network_proxy_proto").(models.Proto)
	return &networkProxyProto
}

func NetworkProxyProtoModelFromMap(m map[string]interface{}) *models.Proto {
	networkProxyProto := m["network_proxy_proto"].(models.Proto)
	return &networkProxyProto
}

// Update the underlying NetworkProxyProto resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetNetworkProxyProtoResourceData(d *schema.ResourceData, m *models.Proto) {
}

// Iterate through and update the NetworkProxyProto resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetNetworkProxyProtoSubResourceData(m []*models.Proto) (d []*map[string]interface{}) {
	for _, NetworkProxyProtoModel := range m {
		if NetworkProxyProtoModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the NetworkProxyProto resource defined in the Terraform configuration
func NetworkProxyProtoSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the NetworkProxyProto resource
func GetNetworkProxyProtoPropertyFields() (t []string) {
	return []string{}
}
