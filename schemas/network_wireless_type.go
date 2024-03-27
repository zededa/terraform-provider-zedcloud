package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

func NetworkWirelessTypeModel(d *schema.ResourceData) *models.Type {
	networkWirelessType, _ := d.Get("network_wireless_type").(models.Type)
	return &networkWirelessType
}

func NetworkWirelessTypeModelFromMap(m map[string]interface{}) *models.Type {
	networkWirelessType := m["network_wireless_type"].(models.Type)
	return &networkWirelessType
}

// Update the underlying NetworkWirelessType resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetNetworkWirelessTypeResourceData(d *schema.ResourceData, m *models.Type) {
}

// Iterate through and update the NetworkWirelessType resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetNetworkWirelessTypeSubResourceData(m []*models.Type) (d []*map[string]interface{}) {
	for _, NetworkWirelessTypeModel := range m {
		if NetworkWirelessTypeModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the NetworkWirelessType resource defined in the Terraform configuration
func NetworkWirelessTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the NetworkWirelessType resource
func GetNetworkWirelessTypePropertyFields() (t []string) {
	return []string{}
}
