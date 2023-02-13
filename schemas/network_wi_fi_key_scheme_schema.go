package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func NetworkWiFiKeySchemeModel(d *schema.ResourceData) *models.NetworkWiFiKeyScheme {
	networkWiFiKeyScheme, _ := d.Get("network_wi_fi_key_scheme").(models.NetworkWiFiKeyScheme)
	return &networkWiFiKeyScheme
}

func NetworkWiFiKeySchemeModelFromMap(m map[string]interface{}) *models.NetworkWiFiKeyScheme {
	networkWiFiKeyScheme := m["network_wi_fi_key_scheme"].(models.NetworkWiFiKeyScheme)
	return &networkWiFiKeyScheme
}

// Update the underlying NetworkWiFiKeyScheme resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetNetworkWiFiKeySchemeResourceData(d *schema.ResourceData, m *models.NetworkWiFiKeyScheme) {
}

// Iterate through and update the NetworkWiFiKeyScheme resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetNetworkWiFiKeySchemeSubResourceData(m []*models.NetworkWiFiKeyScheme) (d []*map[string]interface{}) {
	for _, NetworkWiFiKeySchemeModel := range m {
		if NetworkWiFiKeySchemeModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the NetworkWiFiKeyScheme resource defined in the Terraform configuration
func NetworkWiFiKeySchemeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the NetworkWiFiKeyScheme resource
func GetNetworkWiFiKeySchemePropertyFields() (t []string) {
	return []string{}
}
