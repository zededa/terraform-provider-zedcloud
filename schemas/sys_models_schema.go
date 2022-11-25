package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate SysModels resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func SysModelsModel(d *schema.ResourceData) *models.SysModels {
	list := d.Get("list").([]*models.SysModel) // []*SysModel
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
	return &models.SysModels{
		List:  list,
		Next:  next,
		Terse: terse,
	}
}

func SysModelsModelFromMap(m map[string]interface{}) *models.SysModels {
	list := m["list"].([]*models.SysModel) // []*SysModel
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
	return &models.SysModels{
		List:  list,
		Next:  next,
		Terse: terse,
	}
}

// Update the underlying SysModels resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetSysModelsResourceData(d *schema.ResourceData, m *models.SysModels) {
	d.Set("list", SetSysModelSubResourceData(m.List))
	d.Set("next", SetCursorSubResourceData([]*models.Cursor{m.Next}))
	d.Set("terse", SetSummarySubResourceData([]*models.Summary{m.Terse}))
}

// Iterate throught and update the SysModels resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetSysModelsSubResourceData(m []*models.SysModels) (d []*map[string]interface{}) {
	for _, SysModelsModel := range m {
		if SysModelsModel != nil {
			properties := make(map[string]interface{})
			properties["list"] = SetSysModelSubResourceData(SysModelsModel.List)
			properties["next"] = SetCursorSubResourceData([]*models.Cursor{SysModelsModel.Next})
			properties["terse"] = SetSummarySubResourceData([]*models.Summary{SysModelsModel.Terse})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the SysModels resource defined in the Terraform configuration
func SysModelsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"list": {
			Description: `List of filtered Sys Models`,
			Type:        schema.TypeList, //GoType: []*SysModel
			Elem: &schema.Resource{
				Schema: SysModelSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
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

// Retrieve property field names for updating the SysModels resource
func GetSysModelsPropertyFields() (t []string) {
	return []string{
		"list",
		"next",
		"terse",
	}
}
