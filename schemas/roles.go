package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func RolesModel(d *schema.ResourceData) *models.Roles {
	var list []*models.Role // []*Role
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
			m := RoleModelFromMap(v.(map[string]interface{}))
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
	return &models.Roles{
		List:           list,
		Next:           next,
		SummaryByState: summaryByState,
	}
}

func RolesModelFromMap(m map[string]interface{}) *models.Roles {
	var list []*models.Role // []*Role
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
			m := RoleModelFromMap(v.(map[string]interface{}))
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
	return &models.Roles{
		List:           list,
		Next:           next,
		SummaryByState: summaryByState,
	}
}

func SetRolesResourceData(d *schema.ResourceData, m *models.Roles) {
	d.Set("list", SetRoleSubResourceData(m.List))
	d.Set("next", SetCursorSubResourceData([]*models.Cursor{m.Next}))
	d.Set("summary_by_state", SetSummarySubResourceData([]*models.Summary{m.SummaryByState}))
}

func SetRolesSubResourceData(m []*models.Roles) (d []*map[string]interface{}) {
	for _, RolesModel := range m {
		if RolesModel != nil {
			properties := make(map[string]interface{})
			properties["list"] = SetRoleSubResourceData(RolesModel.List)
			properties["next"] = SetCursorSubResourceData([]*models.Cursor{RolesModel.Next})
			properties["summary_by_state"] = SetSummarySubResourceData([]*models.Summary{RolesModel.SummaryByState})
			d = append(d, &properties)
		}
	}
	return
}

func RolesSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"list": {
			Description: `List of roles`,
			Type:        schema.TypeList, //GoType: []*Role
			Elem: &schema.Resource{
				Schema: RoleSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"next": {
			Description: `Page details of the filtered records`,
			Type:        schema.TypeList, //GoType: Cursor
			Elem: &schema.Resource{
				Schema: CursorSchema(),
			},
			Optional: true,
		},

		"summary_by_state": {
			Description: `Summary of filtered roles`,
			Type:        schema.TypeList, //GoType: Summary
			Elem: &schema.Resource{
				Schema: SummarySchema(),
			},
			Optional: true,
		},
	}
}

func GetRolesPropertyFields() (t []string) {
	return []string{
		"list",
		"next",
		"summary_by_state",
	}
}
