package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func DhcpIPRangeModel(d *schema.ResourceData) *models.DhcpIPRange {
	end, _ := d.Get("end").(string)
	start, _ := d.Get("start").(string)
	return &models.DhcpIPRange{
		End:   end,
		Start: start,
	}
}

func DhcpIPRangeModelFromMap(m map[string]interface{}) *models.DhcpIPRange {
	end := m["end"].(string)
	start := m["start"].(string)
	return &models.DhcpIPRange{
		End:   end,
		Start: start,
	}
}

func SetDhcpIPRangeResourceData(d *schema.ResourceData, m *models.DhcpIPRange) {
	d.Set("end", m.End)
	d.Set("start", m.Start)
}

func SetDhcpIPRangeSubResourceData(m []*models.DhcpIPRange) (d []*map[string]interface{}) {
	for _, DhcpIPRangeModel := range m {
		if DhcpIPRangeModel != nil {
			properties := make(map[string]interface{})
			properties["end"] = DhcpIPRangeModel.End
			properties["start"] = DhcpIPRangeModel.Start
			d = append(d, &properties)
		}
	}
	return
}

func DhcpIPRange() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"end": {
			Description: `ending IP`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"start": {
			Description: `starting IP`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func DhcpIPRangePropertyFields() (t []string) {
	return []string{
		"end",
		"start",
	}
}
