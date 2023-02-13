package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AppInstancesModel(d *schema.ResourceData) *models.AppInstances {
	var list []*models.AppInstance // []*AppInstance
	listInterface, listIsSet := d.GetOk("list")
	if listIsSet {
		var items []interface{}
		if listItems, isList := listInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = listInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := ApplicationInstanceModelFromMap(v.(map[string]interface{}))
			list = append(list, m)
		}
	}
	var next *models.Cursor // Cursor
	nextInterface, nextIsSet := d.GetOk("next")
	if nextIsSet && nextInterface != nil {
		nextMap := nextInterface.([]interface{})
		if len(nextMap) > 0 {
			next = CursorModelFromMap(nextMap[0].(map[string]interface{}))
		}
	}
	var summaryByState *models.Summary // Summary
	summaryByStateInterface, summaryByStateIsSet := d.GetOk("summary_by_state")
	if summaryByStateIsSet && summaryByStateInterface != nil {
		summaryByStateMap := summaryByStateInterface.([]interface{})
		if len(summaryByStateMap) > 0 {
			summaryByState = SummaryModelFromMap(summaryByStateMap[0].(map[string]interface{}))
		}
	}
	return &models.AppInstances{
		List:           list,
		Next:           next,
		SummaryByState: summaryByState,
	}
}

func AppInstancesModelFromMap(m map[string]interface{}) *models.AppInstances {
	var list []*models.AppInstance // []*AppInstance
	listInterface, listIsSet := m["list"]
	if listIsSet {
		var items []interface{}
		if listItems, isList := listInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = listInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := ApplicationInstanceModelFromMap(v.(map[string]interface{}))
			list = append(list, m)
		}
	}
	var next *models.Cursor // Cursor
	nextInterface, nextIsSet := m["next"]
	if nextIsSet && nextInterface != nil {
		nextMap := nextInterface.([]interface{})
		if len(nextMap) > 0 {
			next = CursorModelFromMap(nextMap[0].(map[string]interface{}))
		}
	}
	//
	var summaryByState *models.Summary // Summary
	summaryByStateInterface, summaryByStateIsSet := m["summary_by_state"]
	if summaryByStateIsSet && summaryByStateInterface != nil {
		summaryByStateMap := summaryByStateInterface.([]interface{})
		if len(summaryByStateMap) > 0 {
			summaryByState = SummaryModelFromMap(summaryByStateMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.AppInstances{
		List:           list,
		Next:           next,
		SummaryByState: summaryByState,
	}
}

// Update the underlying AppInstances resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetAppInstancesResourceData(d *schema.ResourceData, m *models.AppInstances) {
	d.Set("list", SetApplicationInstanceSubResourceData(m.List))
	d.Set("next", SetCursorSubResourceData([]*models.Cursor{m.Next}))
	d.Set("summary_by_state", SetSummarySubResourceData([]*models.Summary{m.SummaryByState}))
}

// Iterate through and update the AppInstances resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetAppInstancesSubResourceData(m []*models.AppInstances) (d []*map[string]interface{}) {
	for _, AppInstancesModel := range m {
		if AppInstancesModel != nil {
			properties := make(map[string]interface{})
			properties["list"] = SetApplicationInstanceSubResourceData(AppInstancesModel.List)
			properties["next"] = SetCursorSubResourceData([]*models.Cursor{AppInstancesModel.Next})
			properties["summary_by_state"] = SetSummarySubResourceData([]*models.Summary{AppInstancesModel.SummaryByState})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the AppInstances resource defined in the Terraform configuration
func AppInstancesSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"list": {
			Description: `app insatance list response`,
			Type:        schema.TypeList, //GoType: []*AppInstance
			Elem: &schema.Resource{
				Schema: ApplicationInstance(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"next": {
			Description: `cursor next`,
			Type:        schema.TypeList, //GoType: Cursor
			Elem: &schema.Resource{
				Schema: CursorSchema(),
			},
			Optional: true,
		},

		"summary_by_state": {
			Description: `app instance by summary`,
			Type:        schema.TypeList, //GoType: Summary
			Elem: &schema.Resource{
				Schema: SummarySchema(),
			},
			Optional: true,
		},
	}
}

// Retrieve property field names for updating the AppInstances resource
func GetAppInstancesPropertyFields() (t []string) {
	return []string{
		"list",
		"next",
		"summary_by_state",
	}
}
