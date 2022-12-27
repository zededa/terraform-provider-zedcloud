package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate Tags resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func TagsModel(d *schema.ResourceData) *models.Tags {
	list := d.Get("list").([]*models.Tag) // []*Tag
	var next *models.Cursor               // Cursor
	nextInterface, nextIsSet := d.GetOk("next")
	if nextIsSet {
		nextMap := nextInterface.([]interface{})[0].(map[string]interface{})
		next = CursorModelFromMap(nextMap)
	}
	var summaryByState *models.Summary // Summary
	summaryByStateInterface, summaryByStateIsSet := d.GetOk("summary_by_state")
	if summaryByStateIsSet {
		summaryByStateMap := summaryByStateInterface.([]interface{})[0].(map[string]interface{})
		summaryByState = SummaryModelFromMap(summaryByStateMap)
	}
	var summaryByType *models.Summary // Summary
	summaryByTypeInterface, summaryByTypeIsSet := d.GetOk("summary_by_type")
	if summaryByTypeIsSet {
		summaryByTypeMap := summaryByTypeInterface.([]interface{})[0].(map[string]interface{})
		summaryByType = SummaryModelFromMap(summaryByTypeMap)
	}
	return &models.Tags{
		List:           list,
		Next:           next,
		SummaryByState: summaryByState,
		SummaryByType:  summaryByType,
	}
}

func TagsModelFromMap(m map[string]interface{}) *models.Tags {
	list := m["list"].([]*models.Tag) // []*Tag
	var next *models.Cursor           // Cursor
	nextInterface, nextIsSet := m["next"]
	if nextIsSet {
		nextMap := nextInterface.([]interface{})[0].(map[string]interface{})
		next = CursorModelFromMap(nextMap)
	}
	//
	var summaryByState *models.Summary // Summary
	summaryByStateInterface, summaryByStateIsSet := m["summary_by_state"]
	if summaryByStateIsSet {
		summaryByStateMap := summaryByStateInterface.([]interface{})[0].(map[string]interface{})
		summaryByState = SummaryModelFromMap(summaryByStateMap)
	}
	//
	var summaryByType *models.Summary // Summary
	summaryByTypeInterface, summaryByTypeIsSet := m["summary_by_type"]
	if summaryByTypeIsSet {
		summaryByTypeMap := summaryByTypeInterface.([]interface{})[0].(map[string]interface{})
		summaryByType = SummaryModelFromMap(summaryByTypeMap)
	}
	//
	return &models.Tags{
		List:           list,
		Next:           next,
		SummaryByState: summaryByState,
		SummaryByType:  summaryByType,
	}
}

// Update the underlying Tags resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetTagsResourceData(d *schema.ResourceData, m *models.Tags) {
	d.Set("list", SetTagSubResourceData(m.List))
	d.Set("next", SetCursorSubResourceData([]*models.Cursor{m.Next}))
	d.Set("summary_by_state", SetSummarySubResourceData([]*models.Summary{m.SummaryByState}))
	d.Set("summary_by_type", SetSummarySubResourceData([]*models.Summary{m.SummaryByType}))
}

// Iterate throught and update the Tags resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetTagsSubResourceData(m []*models.Tags) (d []*map[string]interface{}) {
	for _, TagsModel := range m {
		if TagsModel != nil {
			properties := make(map[string]interface{})
			properties["list"] = SetTagSubResourceData(TagsModel.List)
			properties["next"] = SetCursorSubResourceData([]*models.Cursor{TagsModel.Next})
			properties["summary_by_state"] = SetSummarySubResourceData([]*models.Summary{TagsModel.SummaryByState})
			properties["summary_by_type"] = SetSummarySubResourceData([]*models.Summary{TagsModel.SummaryByType})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the Tags resource defined in the Terraform configuration
func TagsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"list": {
			Description: `List of filtered resource group records`,
			Type:        schema.TypeList, //GoType: []*Tag
			Elem: &schema.Resource{
				Schema: TagSchema(),
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

		"summary_by_state": {
			Description: `Summary of filtered resource group records`,
			Type:        schema.TypeList, //GoType: Summary
			Elem: &schema.Resource{
				Schema: SummarySchema(),
			},
			Optional: true,
		},

		"summary_by_type": {
			Description: `Summary of filtered resource group records`,
			Type:        schema.TypeList, //GoType: Summary
			Elem: &schema.Resource{
				Schema: SummarySchema(),
			},
			Optional: true,
		},
	}
}

// Retrieve property field names for updating the Tags resource
func GetTagsPropertyFields() (t []string) {
	return []string{
		"list",
		"next",
		"summary_by_state",
		"summary_by_type",
	}
}
