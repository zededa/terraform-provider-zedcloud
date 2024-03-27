package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

func NetworkInstanceKindModel(d *schema.ResourceData) *models.NetworkInstanceKind {
	networkInstanceKind, _ := d.Get("network_instance_kind").(models.NetworkInstanceKind)
	return &networkInstanceKind
}

func NetworkInstanceKindModelFromMap(m map[string]interface{}) *models.NetworkInstanceKind {
	networkInstanceKind := m["network_instance_kind"].(models.NetworkInstanceKind)
	return &networkInstanceKind
}

// Update the underlying NetworkInstanceKind resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetNetworkInstanceKindResourceData(d *schema.ResourceData, m *models.NetworkInstanceKind) {
}

// Iterate through and update the NetworkInstanceKind resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetNetworkInstanceKindSubResourceData(m []*models.NetworkInstanceKind) (d []*map[string]interface{}) {
	for _, NetworkInstanceKindModel := range m {
		if NetworkInstanceKindModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the NetworkInstanceKind resource defined in the Terraform configuration
func NetworkInstanceKindSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the NetworkInstanceKind resource
func GetNetworkInstanceKindPropertyFields() (t []string) {
	return []string{}
}
