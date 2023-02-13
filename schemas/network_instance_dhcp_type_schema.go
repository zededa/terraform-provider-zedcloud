package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func NetworkInstanceDhcpTypeModel(d *schema.ResourceData) *models.NetworkInstanceDhcpType {
	networkInstanceDhcpType, _ := d.Get("network_instance_dhcp_type").(models.NetworkInstanceDhcpType)
	return &networkInstanceDhcpType
}

func NetworkInstanceDhcpTypeModelFromMap(m map[string]interface{}) *models.NetworkInstanceDhcpType {
	networkInstanceDhcpType := m["network_instance_dhcp_type"].(models.NetworkInstanceDhcpType)
	return &networkInstanceDhcpType
}

// Update the underlying NetworkInstanceDhcpType resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetNetworkInstanceDhcpTypeResourceData(d *schema.ResourceData, m *models.NetworkInstanceDhcpType) {
}

// Iterate through and update the NetworkInstanceDhcpType resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetNetworkInstanceDhcpTypeSubResourceData(m []*models.NetworkInstanceDhcpType) (d []*map[string]interface{}) {
	for _, NetworkInstanceDhcpTypeModel := range m {
		if NetworkInstanceDhcpTypeModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the NetworkInstanceDhcpType resource defined in the Terraform configuration
func NetworkInstanceDhcpTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the NetworkInstanceDhcpType resource
func GetNetworkInstanceDhcpTypePropertyFields() (t []string) {
	return []string{}
}
