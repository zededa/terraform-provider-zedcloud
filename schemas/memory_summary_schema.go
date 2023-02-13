package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate MemorySummary resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func MemorySummaryModel(d *schema.ResourceData) *models.MemorySummary {
	allocatedMB, _ := d.Get("allocated_m_b").(float64)
	totalMB, _ := d.Get("total_m_b").(float64)
	usedMB, _ := d.Get("used_m_b").(float64)
	return &models.MemorySummary{
		AllocatedMB: allocatedMB,
		TotalMB:     totalMB,
		UsedMB:      usedMB,
	}
}

func MemorySummaryModelFromMap(m map[string]interface{}) *models.MemorySummary {
	allocatedMB := m["allocated_m_b"].(float64)
	totalMB := m["total_m_b"].(float64)
	usedMB := m["used_m_b"].(float64)
	return &models.MemorySummary{
		AllocatedMB: allocatedMB,
		TotalMB:     totalMB,
		UsedMB:      usedMB,
	}
}

// Update the underlying MemorySummary resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetMemorySummaryResourceData(d *schema.ResourceData, m *models.MemorySummary) {
	d.Set("allocated_m_b", m.AllocatedMB)
	d.Set("total_m_b", m.TotalMB)
	d.Set("used_m_b", m.UsedMB)
}

// Iterate through and update the MemorySummary resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetMemorySummarySubResourceData(m []*models.MemorySummary) (d []*map[string]interface{}) {
	for _, MemorySummaryModel := range m {
		if MemorySummaryModel != nil {
			properties := make(map[string]interface{})
			properties["allocated_m_b"] = MemorySummaryModel.AllocatedMB
			properties["total_m_b"] = MemorySummaryModel.TotalMB
			properties["used_m_b"] = MemorySummaryModel.UsedMB
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the MemorySummary resource defined in the Terraform configuration
func MemorySummarySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"allocated_m_b": {
			Description: `Total allocated memory in MBs`,
			Type:        schema.TypeFloat,
			Optional:    true,
		},

		"total_m_b": {
			Description: `Total memory for the device in MBs`,
			Type:        schema.TypeFloat,
			Optional:    true,
		},

		"used_m_b": {
			Description: `Total memory used by the device in MBs within allocated memory`,
			Type:        schema.TypeFloat,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the MemorySummary resource
func GetMemorySummaryPropertyFields() (t []string) {
	return []string{
		"allocated_m_b",
		"total_m_b",
		"used_m_b",
	}
}
