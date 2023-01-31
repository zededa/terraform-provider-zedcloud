package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AppInstMemorySummaryModel(d *schema.ResourceData) *models.AppInstMemorySummary {
	allocatedMB, _ := d.Get("allocated_m_b").(float64)
	availableMB, _ := d.Get("available_m_b").(float64)
	usedMB, _ := d.Get("used_m_b").(float64)
	return &models.AppInstMemorySummary{
		AllocatedMB: allocatedMB,
		AvailableMB: availableMB,
		UsedMB:      usedMB,
	}
}

func AppInstMemorySummaryModelFromMap(m map[string]interface{}) *models.AppInstMemorySummary {
	allocatedMB := m["allocated_m_b"].(float64)
	availableMB := m["available_m_b"].(float64)
	usedMB := m["used_m_b"].(float64)
	return &models.AppInstMemorySummary{
		AllocatedMB: allocatedMB,
		AvailableMB: availableMB,
		UsedMB:      usedMB,
	}
}

// Update the underlying AppInstMemorySummary resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetAppInstMemorySummaryResourceData(d *schema.ResourceData, m *models.AppInstMemorySummary) {
	d.Set("allocated_m_b", m.AllocatedMB)
	d.Set("available_m_b", m.AvailableMB)
	d.Set("used_m_b", m.UsedMB)
}

// Iterate through and update the AppInstMemorySummary resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetAppInstMemorySummarySubResourceData(m []*models.AppInstMemorySummary) (d []*map[string]interface{}) {
	for _, AppInstMemorySummaryModel := range m {
		if AppInstMemorySummaryModel != nil {
			properties := make(map[string]interface{})
			properties["allocated_m_b"] = AppInstMemorySummaryModel.AllocatedMB
			properties["available_m_b"] = AppInstMemorySummaryModel.AvailableMB
			properties["used_m_b"] = AppInstMemorySummaryModel.UsedMB
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the AppInstMemorySummary resource defined in the Terraform configuration
func AppInstMemorySummarySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"allocated_m_b": {
			Description: `Total memory allocated for the App instance.`,
			Type:        schema.TypeFloat,
			Optional:    true,
		},

		"available_m_b": {
			Description: `Available / Free memory in the App`,
			Type:        schema.TypeFloat,
			Optional:    true,
		},

		"used_m_b": {
			Description: `Total memory used by the App within allocated memory`,
			Type:        schema.TypeFloat,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the AppInstMemorySummary resource
func GetAppInstMemorySummaryPropertyFields() (t []string) {
	return []string{
		"allocated_m_b",
		"available_m_b",
		"used_m_b",
	}
}
