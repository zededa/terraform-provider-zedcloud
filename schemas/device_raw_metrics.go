package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

// Function to perform the following actions:
// (1) Translate DeviceRawMetrics resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func DeviceRawMetricsModel(d *schema.ResourceData) *models.DeviceRawMetrics {
	id, _ := d.Get("id").(string)
	name, _ := d.Get("name").(string)
	rawMetrics, _ := d.Get("raw_metrics").(string)
	return &models.DeviceRawMetrics{
		ID:         id,
		Name:       name,
		RawMetrics: rawMetrics,
	}
}

func DeviceRawMetricsModelFromMap(m map[string]interface{}) *models.DeviceRawMetrics {
	id := m["id"].(string)
	name := m["name"].(string)
	rawMetrics := m["raw_metrics"].(string)
	return &models.DeviceRawMetrics{
		ID:         id,
		Name:       name,
		RawMetrics: rawMetrics,
	}
}

// Update the underlying DeviceRawMetrics resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDeviceRawMetricsResourceData(d *schema.ResourceData, m *models.DeviceRawMetrics) {
	d.Set("id", m.ID)
	d.Set("name", m.Name)
	d.Set("raw_metrics", m.RawMetrics)
}

// Iterate through and update the DeviceRawMetrics resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDeviceRawMetricsSubResourceData(m []*models.DeviceRawMetrics) (d []*map[string]interface{}) {
	for _, DeviceRawMetricsModel := range m {
		if DeviceRawMetricsModel != nil {
			properties := make(map[string]interface{})
			properties["id"] = DeviceRawMetricsModel.ID
			properties["name"] = DeviceRawMetricsModel.Name
			properties["raw_metrics"] = DeviceRawMetricsModel.RawMetrics
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the DeviceRawMetrics resource defined in the Terraform configuration
func DeviceRawMetricsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Description: ``,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"name": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"raw_metrics": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the DeviceRawMetrics resource
func GetDeviceRawMetricsPropertyFields() (t []string) {
	return []string{
		"id",
		"name",
		"raw_metrics",
	}
}
