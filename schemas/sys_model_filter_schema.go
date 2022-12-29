package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate SysModelFilter resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func SysModelFilterModel(d *schema.ResourceData) *models.SysModelFilter {
	brandID, _ := d.Get("brand_id").(string)
	brandIds, _ := d.Get("brand_ids").([]string)
	namePattern, _ := d.Get("name_pattern").(string)
	originTypeModel, _ := d.Get("origin_type").(models.Origin) // Origin
	originType := &originTypeModel
	if !ok {
		originType = nil
	}
	return &models.SysModelFilter{
		BrandID:     brandID,
		BrandIds:    brandIds,
		NamePattern: namePattern,
		OriginType:  originType,
	}
}

func SysModelFilterModelFromMap(m map[string]interface{}) *models.SysModelFilter {
	brandID := m["brand_id"].(string)
	brandIds := m["brand_ids"].([]string)
	namePattern := m["name_pattern"].(string)
	originType := m["origin_type"].(*models.Origin) // Origin
	return &models.SysModelFilter{
		BrandID:     brandID,
		BrandIds:    brandIds,
		NamePattern: namePattern,
		OriginType:  originType,
	}
}

// Update the underlying SysModelFilter resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetSysModelFilterResourceData(d *schema.ResourceData, m *models.SysModelFilter) {
	d.Set("brand_id", m.BrandID)
	d.Set("brand_ids", m.BrandIds)
	d.Set("name_pattern", m.NamePattern)
	d.Set("origin_type", m.OriginType)
}

// Iterate through and update the SysModelFilter resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetSysModelFilterSubResourceData(m []*models.SysModelFilter) (d []*map[string]interface{}) {
	for _, SysModelFilterModel := range m {
		if SysModelFilterModel != nil {
			properties := make(map[string]interface{})
			properties["brand_id"] = SysModelFilterModel.BrandID
			properties["brand_ids"] = SysModelFilterModel.BrandIds
			properties["name_pattern"] = SysModelFilterModel.NamePattern
			properties["origin_type"] = SysModelFilterModel.OriginType
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the SysModelFilter resource defined in the Terraform configuration
func SysModelFilterSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"brand_id": {
			Description: `System defined universally unique Id of the brand.`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"brand_ids": {
			Description: `System defined universally unique Ids of the brand.`,
			Type:        schema.TypeList, //GoType: []string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"name_pattern": {
			Description: `Model name pattern to be matched.`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"origin_type": {
			Description: `origin of object`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the SysModelFilter resource
func GetSysModelFilterPropertyFields() (t []string) {
	return []string{
		"brand_id",
		"brand_ids",
		"name_pattern",
		"origin_type",
	}
}
