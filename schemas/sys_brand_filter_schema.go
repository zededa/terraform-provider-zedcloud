package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate SysBrandFilter resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func SysBrandFilterModel(d *schema.ResourceData) *models.SysBrandFilter {
	namePattern := d.Get("name_pattern").(string)
	originType := d.Get("origin_type").(*models.Origin) // Origin
	return &models.SysBrandFilter{
		NamePattern: namePattern,
		OriginType:  originType,
	}
}

func SysBrandFilterModelFromMap(m map[string]interface{}) *models.SysBrandFilter {
	namePattern := m["name_pattern"].(string)
	originType := m["origin_type"].(*models.Origin) // Origin
	return &models.SysBrandFilter{
		NamePattern: namePattern,
		OriginType:  originType,
	}
}

// Update the underlying SysBrandFilter resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetSysBrandFilterResourceData(d *schema.ResourceData, m *models.SysBrandFilter) {
	d.Set("name_pattern", m.NamePattern)
	d.Set("origin_type", m.OriginType)
}

// Iterate throught and update the SysBrandFilter resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetSysBrandFilterSubResourceData(m []*models.SysBrandFilter) (d []*map[string]interface{}) {
	for _, SysBrandFilterModel := range m {
		if SysBrandFilterModel != nil {
			properties := make(map[string]interface{})
			properties["name_pattern"] = SysBrandFilterModel.NamePattern
			properties["origin_type"] = SysBrandFilterModel.OriginType
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the SysBrandFilter resource defined in the Terraform configuration
func SysBrandFilterSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name_pattern": {
			Description: `Brand name pattern to be matched.`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"origin_type": {
			Description: `origin of object`,
			// We assume it's an enum type
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

// Retrieve property field names for updating the SysBrandFilter resource
func GetSysBrandFilterPropertyFields() (t []string) {
	return []string{
		"name_pattern",
		"origin_type",
	}
}
