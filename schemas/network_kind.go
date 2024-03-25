package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func NetworkKindModel(d *schema.ResourceData) *models.NetworkKind {
	networkKind, _ := d.Get("network_kind").(models.NetworkKind)
	return &networkKind
}

func NetworkKindModelFromMap(m map[string]interface{}) *models.NetworkKind {
	networkKind := m["network_kind"].(models.NetworkKind)
	return &networkKind
}

// Update the underlying NetworkKind resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetNetworkKindResourceData(d *schema.ResourceData, m *models.NetworkKind) {
}

// Iterate through and update the NetworkKind resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetNetworkKindSubResourceData(m []*models.NetworkKind) (d []*map[string]interface{}) {
	for _, NetworkKindModel := range m {
		if NetworkKindModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the NetworkKind resource defined in the Terraform configuration
func NetworkKindSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the NetworkKind resource
func GetNetworkKindPropertyFields() (t []string) {
	return []string{}
}
