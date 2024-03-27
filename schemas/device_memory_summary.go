package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

// Function to perform the following actions:
// (1) Translate DeviceMemorySummary resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func DeviceMemorySummaryModel(d *schema.ResourceData) *models.DeviceMemorySummary {
	allocatedAppsMB, _ := d.Get("allocated_apps_m_b").(float64)
	allocatedEveMB, _ := d.Get("allocated_eve_m_b").(float64)
	availableForNewApps, _ := d.Get("available_for_new_apps").(float64)
	deviceMemoryMB, _ := d.Get("device_memory_m_b").(float64)
	usedEveMB, _ := d.Get("used_eve_m_b").(float64)
	return &models.DeviceMemorySummary{
		AllocatedAppsMB:     allocatedAppsMB,
		AllocatedEveMB:      allocatedEveMB,
		AvailableForNewApps: availableForNewApps,
		DeviceMemoryMB:      deviceMemoryMB,
		UsedEveMB:           usedEveMB,
	}
}

func DeviceMemorySummaryModelFromMap(m map[string]interface{}) *models.DeviceMemorySummary {
	allocatedAppsMB := m["allocated_apps_m_b"].(float64)
	allocatedEveMB := m["allocated_eve_m_b"].(float64)
	availableForNewApps := m["available_for_new_apps"].(float64)
	deviceMemoryMB := m["device_memory_m_b"].(float64)
	usedEveMB := m["used_eve_m_b"].(float64)
	return &models.DeviceMemorySummary{
		AllocatedAppsMB:     allocatedAppsMB,
		AllocatedEveMB:      allocatedEveMB,
		AvailableForNewApps: availableForNewApps,
		DeviceMemoryMB:      deviceMemoryMB,
		UsedEveMB:           usedEveMB,
	}
}

// Update the underlying DeviceMemorySummary resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDeviceMemorySummaryResourceData(d *schema.ResourceData, m *models.DeviceMemorySummary) {
	d.Set("allocated_apps_m_b", m.AllocatedAppsMB)
	d.Set("allocated_eve_m_b", m.AllocatedEveMB)
	d.Set("available_for_new_apps", m.AvailableForNewApps)
	d.Set("device_memory_m_b", m.DeviceMemoryMB)
	d.Set("used_eve_m_b", m.UsedEveMB)
}

// Iterate through and update the DeviceMemorySummary resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDeviceMemorySummarySubResourceData(m []*models.DeviceMemorySummary) (d []*map[string]interface{}) {
	for _, DeviceMemorySummaryModel := range m {
		if DeviceMemorySummaryModel != nil {
			properties := make(map[string]interface{})
			properties["allocated_apps_m_b"] = DeviceMemorySummaryModel.AllocatedAppsMB
			properties["allocated_eve_m_b"] = DeviceMemorySummaryModel.AllocatedEveMB
			properties["available_for_new_apps"] = DeviceMemorySummaryModel.AvailableForNewApps
			properties["device_memory_m_b"] = DeviceMemorySummaryModel.DeviceMemoryMB
			properties["used_eve_m_b"] = DeviceMemorySummaryModel.UsedEveMB
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the DeviceMemorySummary resource defined in the Terraform configuration
func DeviceMemorySummarySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"allocated_apps_m_b": {
			Description: `Total memory allocated to app instances.`,
			Type:        schema.TypeFloat,
			Optional:    true,
		},

		"allocated_eve_m_b": {
			Description: `Total memory allocated to EVE on the device.`,
			Type:        schema.TypeFloat,
			Optional:    true,
		},

		"available_for_new_apps": {
			Description: `Memory Available for new app-instances on the device.`,
			Type:        schema.TypeFloat,
			Optional:    true,
		},

		"device_memory_m_b": {
			Description: `Total memory on the device.`,
			Type:        schema.TypeFloat,
			Optional:    true,
		},

		"used_eve_m_b": {
			Description: `Memory currently used by EVE on the device.`,
			Type:        schema.TypeFloat,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the DeviceMemorySummary resource
func GetDeviceMemorySummaryPropertyFields() (t []string) {
	return []string{
		"allocated_apps_m_b",
		"allocated_eve_m_b",
		"available_for_new_apps",
		"device_memory_m_b",
		"used_eve_m_b",
	}
}
