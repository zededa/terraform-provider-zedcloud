package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate SysBrands resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func SysBrandsModel(d *schema.ResourceData) *models.SysBrands {
	list := d.Get("list").([]*models.SysBrand) // []*SysBrand
	var next *models.Cursor                    // Cursor
	nextInterface, nextIsSet := d.GetOk("next")
	if nextIsSet {
		nextMap := nextInterface.([]interface{})[0].(map[string]interface{})
		next = CursorModelFromMap(nextMap)
	}
	var terse *models.Summary // Summary
	terseInterface, terseIsSet := d.GetOk("terse")
	if terseIsSet {
		terseMap := terseInterface.([]interface{})[0].(map[string]interface{})
		terse = SummaryModelFromMap(terseMap)
	}
	return &models.SysBrands{
		List:  list,
		Next:  next,
		Terse: terse,
	}
}

func SysBrandsModelFromMap(m map[string]interface{}) *models.SysBrands {
	list := m["list"].([]*models.SysBrand) // []*SysBrand
	var next *models.Cursor                // Cursor
	nextInterface, nextIsSet := m["next"]
	if nextIsSet {
		nextMap := nextInterface.([]interface{})[0].(map[string]interface{})
		next = CursorModelFromMap(nextMap)
	}
	//
	var terse *models.Summary // Summary
	terseInterface, terseIsSet := m["terse"]
	if terseIsSet {
		terseMap := terseInterface.([]interface{})[0].(map[string]interface{})
		terse = SummaryModelFromMap(terseMap)
	}
	//
	return &models.SysBrands{
		List:  list,
		Next:  next,
		Terse: terse,
	}
}

// Update the underlying SysBrands resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetSysBrandsResourceData(d *schema.ResourceData, m *models.SysBrands) {
	d.Set("list", SetSysBrandSubResourceData(m.List))
	d.Set("next", SetCursorSubResourceData([]*models.Cursor{m.Next}))
	d.Set("terse", SetSummarySubResourceData([]*models.Summary{m.Terse}))
}

// Iterate throught and update the SysBrands resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetSysBrandsSubResourceData(m []*models.SysBrands) (d []*map[string]interface{}) {
	for _, SysBrandsModel := range m {
		if SysBrandsModel != nil {
			properties := make(map[string]interface{})
			properties["list"] = SetSysBrandSubResourceData(SysBrandsModel.List)
			properties["next"] = SetCursorSubResourceData([]*models.Cursor{SysBrandsModel.Next})
			properties["terse"] = SetSummarySubResourceData([]*models.Summary{SysBrandsModel.Terse})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the SysBrands resource defined in the Terraform configuration
func SysBrandsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"list": {
			Description: `List of filtered Sys Models`,
			Type:        schema.TypeList, //GoType: []*SysBrand
			Elem: &schema.Resource{
				Schema: SysBrandSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"next": {
			Description: `Responded page details of filtered records`,
			Type:        schema.TypeList, //GoType: Cursor
			Elem: &schema.Resource{
				Schema: CursorSchema(),
			},
			Optional: true,
		},

		"terse": {
			Description: `Summary of filtered model records`,
			Type:        schema.TypeList, //GoType: Summary
			Elem: &schema.Resource{
				Schema: SummarySchema(),
			},
			Optional: true,
		},
	}
}

// Retrieve property field names for updating the SysBrands resource
func GetSysBrandsPropertyFields() (t []string) {
	return []string{
		"list",
		"next",
		"terse",
	}
}
