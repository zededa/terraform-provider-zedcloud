package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

func NetworkDHCPTypeModel(d *schema.ResourceData) *models.NetworkDHCPType {
	networkDHCPType, _ := d.Get("network_d_h_c_p_type").(models.NetworkDHCPType)
	return &networkDHCPType
}

func NetworkDHCPTypeModelFromMap(m map[string]interface{}) *models.NetworkDHCPType {
	networkDHCPType := m["network_d_h_c_p_type"].(models.NetworkDHCPType)
	return &networkDHCPType
}

// Update the underlying NetworkDHCPType resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetNetworkDHCPTypeResourceData(d *schema.ResourceData, m *models.NetworkDHCPType) {
}

// Iterate through and update the NetworkDHCPType resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetNetworkDHCPTypeSubResourceData(m []*models.NetworkDHCPType) (d []*map[string]interface{}) {
	for _, NetworkDHCPTypeModel := range m {
		if NetworkDHCPTypeModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the NetworkDHCPType resource defined in the Terraform configuration
func NetworkDHCPTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the NetworkDHCPType resource
func GetNetworkDHCPTypePropertyFields() (t []string) {
	return []string{}
}
