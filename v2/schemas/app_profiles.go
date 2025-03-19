package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func AppProfilesModel(d *schema.ResourceData) *models.AppProfiles {
	var list []*models.AppProfileConfigSummary // []*AppProfileConfigSummary
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
			m := AppProfileConfigSummaryModelFromMap(v.(map[string]interface{}))
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
	return &models.AppProfiles{
		List:           list,
		Next:           next,
		SummaryByState: summaryByState,
	}
}

func AppProfilesModelFromMap(m map[string]interface{}) *models.AppProfiles {
	var list []*models.AppProfileConfigSummary // []*AppProfileConfigSummary
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
			m := AppProfileConfigSummaryModelFromMap(v.(map[string]interface{}))
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
	return &models.AppProfiles{
		List:           list,
		Next:           next,
		SummaryByState: summaryByState,
	}
}

func SetAppProfilesResourceData(d *schema.ResourceData, m *models.AppProfiles) {
	d.Set("list", SetAppProfileConfigSummarySubResourceData(m.List))
	d.Set("next", SetCursorSubResourceData([]*models.Cursor{m.Next}))
	d.Set("summary_by_state", SetSummarySubResourceData([]*models.Summary{m.SummaryByState}))
}

func SetAppProfilesSubResourceData(m []*models.AppProfiles) (d []*map[string]interface{}) {
	for _, AppProfilesModel := range m {
		if AppProfilesModel != nil {
			properties := make(map[string]interface{})
			properties["list"] = SetAppProfileConfigSummarySubResourceData(AppProfilesModel.List)
			properties["next"] = SetCursorSubResourceData([]*models.Cursor{AppProfilesModel.Next})
			properties["summary_by_state"] = SetSummarySubResourceData([]*models.Summary{AppProfilesModel.SummaryByState})
			d = append(d, &properties)
		}
	}
	return
}

func AppProfilesSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"list": {
			Description: `App profile summary list response`,
			Type:        schema.TypeList, //GoType: []*AppProfileConfigSummary
			Elem: &schema.Resource{
				Schema: AppProfileConfigSummarySchema(),
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
			Description: `app profile by summary`,
			Type:        schema.TypeList, //GoType: Summary
			Elem: &schema.Resource{
				Schema: SummarySchema(),
			},
			Optional: true,
		},
	}
}

func GetAppProfilesPropertyFields() (t []string) {
	return []string{
		"list",
		"next",
		"summary_by_state",
	}
}
