package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

// Function to perform the following actions:
// (1) Translate StorageSummary resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func StorageSummaryModel(d *schema.ResourceData) *models.StorageSummary {
	allocatedMB, _ := d.Get("allocated_m_b").(float64)
	totalMB, _ := d.Get("total_m_b").(float64)
	usedMB, _ := d.Get("used_m_b").(float64)
	return &models.StorageSummary{
		AllocatedMB: allocatedMB,
		TotalMB:     totalMB,
		UsedMB:      usedMB,
	}
}

func StorageSummaryModelFromMap(m map[string]interface{}) *models.StorageSummary {
	allocatedMB := m["allocated_m_b"].(float64)
	totalMB := m["total_m_b"].(float64)
	usedMB := m["used_m_b"].(float64)
	return &models.StorageSummary{
		AllocatedMB: allocatedMB,
		TotalMB:     totalMB,
		UsedMB:      usedMB,
	}
}

// Update the underlying StorageSummary resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetStorageSummaryResourceData(d *schema.ResourceData, m *models.StorageSummary) {
	d.Set("allocated_m_b", m.AllocatedMB)
	d.Set("total_m_b", m.TotalMB)
	d.Set("used_m_b", m.UsedMB)
}

// Iterate through and update the StorageSummary resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetStorageSummarySubResourceData(m []*models.StorageSummary) (d []*map[string]interface{}) {
	for _, StorageSummaryModel := range m {
		if StorageSummaryModel != nil {
			properties := make(map[string]interface{})
			properties["allocated_m_b"] = StorageSummaryModel.AllocatedMB
			properties["total_m_b"] = StorageSummaryModel.TotalMB
			properties["used_m_b"] = StorageSummaryModel.UsedMB
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the StorageSummary resource defined in the Terraform configuration
func StorageSummarySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"allocated_m_b": {
			Description: `Total reserved for running applications + temp. images etc`,
			Type:        schema.TypeFloat,
			Optional:    true,
		},

		"total_m_b": {
			Description: `Total Storage for the device in MBs`,
			Type:        schema.TypeFloat,
			Optional:    true,
		},

		"used_m_b": {
			Description: `How much is used within the allocated total storage`,
			Type:        schema.TypeFloat,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the StorageSummary resource
func GetStorageSummaryPropertyFields() (t []string) {
	return []string{
		"allocated_m_b",
		"total_m_b",
		"used_m_b",
	}
}
