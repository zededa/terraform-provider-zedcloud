package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ResourceUsageModel(d *schema.ResourceData) *models.ResourceUsage {
	current, _ := d.Get("current").(float64)
	total, _ := d.Get("total").(float64)
	return &models.ResourceUsage{
		Current: current,
		Total:   total,
	}
}

func ResourceUsageModelFromMap(m map[string]interface{}) *models.ResourceUsage {
	current := m["current"].(float64)
	total := m["total"].(float64)
	return &models.ResourceUsage{
		Current: current,
		Total:   total,
	}
}

func SetResourceUsageResourceData(d *schema.ResourceData, m *models.ResourceUsage) {
	d.Set("current", m.Current)
	d.Set("total", m.Total)
}

func SetResourceUsageSubResourceData(m []*models.ResourceUsage) (d []*map[string]interface{}) {
	for _, ResourceUsageModel := range m {
		if ResourceUsageModel != nil {
			properties := make(map[string]interface{})
			properties["current"] = ResourceUsageModel.Current
			properties["total"] = ResourceUsageModel.Total
			d = append(d, &properties)
		}
	}
	return
}

func ResourceUsageSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"current": {
			Description: `Current resource usage`,
			Type:        schema.TypeFloat,
			Optional:    true,
		},

		"total": {
			Description: `Total resource usage`,
			Type:        schema.TypeFloat,
			Optional:    true,
		},
	}
}

func GetResourceUsagePropertyFields() (t []string) {
	return []string{
		"current",
		"total",
	}
}
