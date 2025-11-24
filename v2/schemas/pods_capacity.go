package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func PodsCapacityModel(d *schema.ResourceData) *models.PodsCapacity {
	totalInt, _ := d.Get("total").(int)
	total := int32(totalInt)
	usedInt, _ := d.Get("used").(int)
	used := int32(usedInt)
	return &models.PodsCapacity{
		Total: total,
		Used:  used,
	}
}

func PodsCapacityModelFromMap(m map[string]interface{}) *models.PodsCapacity {
	total := int32(m["total"].(int)) // int32
	used := int32(m["used"].(int))   // int32
	return &models.PodsCapacity{
		Total: total,
		Used:  used,
	}
}

func SetPodsCapacityResourceData(d *schema.ResourceData, m *models.PodsCapacity) {
	d.Set("total", m.Total)
	d.Set("used", m.Used)
}

func SetPodsCapacitySubResourceData(m []*models.PodsCapacity) (d []*map[string]interface{}) {
	for _, PodsCapacityModel := range m {
		if PodsCapacityModel != nil {
			properties := make(map[string]interface{})
			properties["total"] = PodsCapacityModel.Total
			properties["used"] = PodsCapacityModel.Used
			d = append(d, &properties)
		}
	}
	return
}

func PodsCapacitySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"total": {
			Description: `Total pods capacity`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"used": {
			Description: `Used pods capacity`,
			Type:        schema.TypeInt,
			Optional:    true,
		},
	}
}

func GetPodsCapacityPropertyFields() (t []string) {
	return []string{
		"total",
		"used",
	}
}
