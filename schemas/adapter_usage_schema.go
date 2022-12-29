package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate AdapterUsage resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func AdapterUsageModel(d *schema.ResourceData) *models.AdapterUsage {
	adapterUsage, _ := d.Get("adapter_usage").(models.AdapterUsage)
	return &adapterUsage
}

func AdapterUsageModelFromMap(m map[string]interface{}) *models.AdapterUsage {
	adapterUsage := m["adapter_usage"].(models.AdapterUsage)
	return &adapterUsage
}

// Update the underlying AdapterUsage resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetAdapterUsageResourceData(d *schema.ResourceData, m *models.AdapterUsage) {
}

// Iterate through and update the AdapterUsage resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetAdapterUsageSubResourceData(m []*models.AdapterUsage) (d []*map[string]interface{}) {
	for _, AdapterUsageModel := range m {
		if AdapterUsageModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the AdapterUsage resource defined in the Terraform configuration
func AdapterUsageSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the AdapterUsage resource
func GetAdapterUsagePropertyFields() (t []string) {
	return []string{}
}
