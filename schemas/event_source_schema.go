package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate EventSource resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func EventSourceModel(d *schema.ResourceData) *models.EventSource {
	eventSource := d.Get("event_source").(models.EventSource)
	return &eventSource
}

func EventSourceModelFromMap(m map[string]interface{}) *models.EventSource {
	eventSource := m["event_source"].(models.EventSource)
	return &eventSource
}

// Update the underlying EventSource resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetEventSourceResourceData(d *schema.ResourceData, m *models.EventSource) {
}

// Iterate throught and update the EventSource resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetEventSourceSubResourceData(m []*models.EventSource) (d []*map[string]interface{}) {
	for _, EventSourceModel := range m {
		if EventSourceModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the EventSource resource defined in the Terraform configuration
func EventSourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the EventSource resource
func GetEventSourcePropertyFields() (t []string) {
	return []string{}
}
