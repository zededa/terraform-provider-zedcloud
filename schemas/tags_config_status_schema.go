package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate TagsConfigStatus resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func TagsConfigStatusModel(d *schema.ResourceData) *models.TagsConfigStatus {
	list := d.Get("list").([]*models.TagConfigStatus) // []*TagConfigStatus
	var next *models.Cursor                           // Cursor
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
	totalCount := int32(d.Get("total_count").(int))
	return &models.TagsConfigStatus{
		List:           list,
		Next:           next,
		SummaryByState: summaryByState,
		SummaryByType:  summaryByType,
		TotalCount:     totalCount,
	}
}

func TagsConfigStatusModelFromMap(m map[string]interface{}) *models.TagsConfigStatus {
	list := m["list"].([]*models.TagConfigStatus) // []*TagConfigStatus
	var next *models.Cursor                       // Cursor
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
	totalCount := int32(m["total_count"].(int)) // int32 false false false
	return &models.TagsConfigStatus{
		List:           list,
		Next:           next,
		SummaryByState: summaryByState,
		SummaryByType:  summaryByType,
		TotalCount:     totalCount,
	}
}

// Update the underlying TagsConfigStatus resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetTagsConfigStatusResourceData(d *schema.ResourceData, m *models.TagsConfigStatus) {
	d.Set("list", SetTagConfigStatusSubResourceData(m.List))
	d.Set("next", SetCursorSubResourceData([]*models.Cursor{m.Next}))
	d.Set("summary_by_state", SetSummarySubResourceData([]*models.Summary{m.SummaryByState}))
	d.Set("summary_by_type", SetSummarySubResourceData([]*models.Summary{m.SummaryByType}))
	d.Set("total_count", m.TotalCount)
}

// Iterate throught and update the TagsConfigStatus resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetTagsConfigStatusSubResourceData(m []*models.TagsConfigStatus) (d []*map[string]interface{}) {
	for _, TagsConfigStatusModel := range m {
		if TagsConfigStatusModel != nil {
			properties := make(map[string]interface{})
			properties["list"] = SetTagConfigStatusSubResourceData(TagsConfigStatusModel.List)
			properties["next"] = SetCursorSubResourceData([]*models.Cursor{TagsConfigStatusModel.Next})
			properties["summary_by_state"] = SetSummarySubResourceData([]*models.Summary{TagsConfigStatusModel.SummaryByState})
			properties["summary_by_type"] = SetSummarySubResourceData([]*models.Summary{TagsConfigStatusModel.SummaryByType})
			properties["total_count"] = TagsConfigStatusModel.TotalCount
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the TagsConfigStatus resource defined in the Terraform configuration
func TagsConfigStatusSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"list": {
			Description: `List of filtered resource group records`,
			Type:        schema.TypeList, //GoType: []*TagConfigStatus
			Elem: &schema.Resource{
				Schema: TagConfigStatusSchema(),
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

		"total_count": {
			Description: `total count of projects`,
			Type:        schema.TypeInt,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the TagsConfigStatus resource
func GetTagsConfigStatusPropertyFields() (t []string) {
	return []string{
		"list",
		"next",
		"summary_by_state",
		"summary_by_type",
		"total_count",
	}
}
