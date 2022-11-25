package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate DhcpIPRange resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func DhcpIPRangeModel(d *schema.ResourceData) *models.DhcpIPRange {
	end := d.Get("end").(string)
	start := d.Get("start").(string)
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

// Update the underlying DhcpIPRange resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDhcpIPRangeResourceData(d *schema.ResourceData, m *models.DhcpIPRange) {
	d.Set("end", m.End)
	d.Set("start", m.Start)
}

// Iterate throught and update the DhcpIPRange resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
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

// Schema mapping representing the DhcpIPRange resource defined in the Terraform configuration
func DhcpIPRangeSchema() map[string]*schema.Schema {
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

// Retrieve property field names for updating the DhcpIPRange resource
func GetDhcpIPRangePropertyFields() (t []string) {
	return []string{
		"end",
		"start",
	}
}
