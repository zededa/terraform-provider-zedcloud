package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate TagStatusListMsg resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func TagStatusListMsgModel(d *schema.ResourceData) *models.TagStatusListMsg {
	list := d.Get("list").([]*models.TagStatusMsg) // []*TagStatusMsg
	var next *models.Cursor                        // Cursor
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
	return &models.TagStatusListMsg{
		List:           list,
		Next:           next,
		SummaryByState: summaryByState,
		SummaryByType:  summaryByType,
	}
}

func TagStatusListMsgModelFromMap(m map[string]interface{}) *models.TagStatusListMsg {
	list := m["list"].([]*models.TagStatusMsg) // []*TagStatusMsg
	var next *models.Cursor                    // Cursor
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
	return &models.TagStatusListMsg{
		List:           list,
		Next:           next,
		SummaryByState: summaryByState,
		SummaryByType:  summaryByType,
	}
}

// Update the underlying TagStatusListMsg resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetTagStatusListMsgResourceData(d *schema.ResourceData, m *models.TagStatusListMsg) {
	d.Set("list", SetTagStatusMsgSubResourceData(m.List))
	d.Set("next", SetCursorSubResourceData([]*models.Cursor{m.Next}))
	d.Set("summary_by_state", SetSummarySubResourceData([]*models.Summary{m.SummaryByState}))
	d.Set("summary_by_type", SetSummarySubResourceData([]*models.Summary{m.SummaryByType}))
}

// Iterate throught and update the TagStatusListMsg resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetTagStatusListMsgSubResourceData(m []*models.TagStatusListMsg) (d []*map[string]interface{}) {
	for _, TagStatusListMsgModel := range m {
		if TagStatusListMsgModel != nil {
			properties := make(map[string]interface{})
			properties["list"] = SetTagStatusMsgSubResourceData(TagStatusListMsgModel.List)
			properties["next"] = SetCursorSubResourceData([]*models.Cursor{TagStatusListMsgModel.Next})
			properties["summary_by_state"] = SetSummarySubResourceData([]*models.Summary{TagStatusListMsgModel.SummaryByState})
			properties["summary_by_type"] = SetSummarySubResourceData([]*models.Summary{TagStatusListMsgModel.SummaryByType})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the TagStatusListMsg resource defined in the Terraform configuration
func TagStatusListMsgSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"list": {
			Description: `List of filtered resource group records`,
			Type:        schema.TypeList, //GoType: []*TagStatusMsg
			Elem: &schema.Resource{
				Schema: TagStatusMsgSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"next": {
			Description: `Responded page details of filtered records`,
			// We assume it's an enum type
			Type:     schema.TypeString,
			Optional: true,
		},

		"summary_by_state": {
			Description: `Summary of filtered resource group records`,
			// We assume it's an enum type
			Type:     schema.TypeString,
			Optional: true,
		},

		"summary_by_type": {
			Description: `Summary of filtered resource group records`,
			// We assume it's an enum type
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

// Retrieve property field names for updating the TagStatusListMsg resource
func GetTagStatusListMsgPropertyFields() (t []string) {
	return []string{
		"list",
		"next",
		"summary_by_state",
		"summary_by_type",
	}
}
