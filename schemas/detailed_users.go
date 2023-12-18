package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func DetailedUsersModel(d *schema.ResourceData) *models.DetailedUsers {
	var list []*models.DetailedUser // []*DetailedUser
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
			m := DetailedUserModelFromMap(v.(map[string]interface{}))
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
	var summaryByRoleDistribution *models.Summary // Summary
	summaryByRoleDistributionInterface, summaryByRoleDistributionIsSet := d.GetOk("summary_by_role_distribution")
	if summaryByRoleDistributionIsSet && summaryByRoleDistributionInterface != nil {
		summaryByRoleDistributionMap := summaryByRoleDistributionInterface.([]interface{})
		if len(summaryByRoleDistributionMap) > 0 {
			summaryByRoleDistribution = SummaryModelFromMap(summaryByRoleDistributionMap[0].(map[string]interface{}))
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
	var summaryByUserActivity *models.Summary // Summary
	summaryByUserActivityInterface, summaryByUserActivityIsSet := d.GetOk("summary_by_user_activity")
	if summaryByUserActivityIsSet && summaryByUserActivityInterface != nil {
		summaryByUserActivityMap := summaryByUserActivityInterface.([]interface{})
		if len(summaryByUserActivityMap) > 0 {
			summaryByUserActivity = SummaryModelFromMap(summaryByUserActivityMap[0].(map[string]interface{}))
		}
	}
	return &models.DetailedUsers{
		List:                      list,
		Next:                      next,
		SummaryByRoleDistribution: summaryByRoleDistribution,
		SummaryByState:            summaryByState,
		SummaryByUserActivity:     summaryByUserActivity,
	}
}

func DetailedUsersModelFromMap(m map[string]interface{}) *models.DetailedUsers {
	var list []*models.DetailedUser // []*DetailedUser
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
			m := DetailedUserModelFromMap(v.(map[string]interface{}))
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
	var summaryByRoleDistribution *models.Summary // Summary
	summaryByRoleDistributionInterface, summaryByRoleDistributionIsSet := m["summary_by_role_distribution"]
	if summaryByRoleDistributionIsSet && summaryByRoleDistributionInterface != nil {
		summaryByRoleDistributionMap := summaryByRoleDistributionInterface.([]interface{})
		if len(summaryByRoleDistributionMap) > 0 {
			summaryByRoleDistribution = SummaryModelFromMap(summaryByRoleDistributionMap[0].(map[string]interface{}))
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
	var summaryByUserActivity *models.Summary // Summary
	summaryByUserActivityInterface, summaryByUserActivityIsSet := m["summary_by_user_activity"]
	if summaryByUserActivityIsSet && summaryByUserActivityInterface != nil {
		summaryByUserActivityMap := summaryByUserActivityInterface.([]interface{})
		if len(summaryByUserActivityMap) > 0 {
			summaryByUserActivity = SummaryModelFromMap(summaryByUserActivityMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.DetailedUsers{
		List:                      list,
		Next:                      next,
		SummaryByRoleDistribution: summaryByRoleDistribution,
		SummaryByState:            summaryByState,
		SummaryByUserActivity:     summaryByUserActivity,
	}
}

func SetDetailedUsersResourceData(d *schema.ResourceData, m *models.DetailedUsers) {
	d.Set("list", SetDetailedUserSubResourceData(m.List))
	d.Set("next", SetCursorSubResourceData([]*models.Cursor{m.Next}))
	d.Set("summary_by_role_distribution", SetSummarySubResourceData([]*models.Summary{m.SummaryByRoleDistribution}))
	d.Set("summary_by_state", SetSummarySubResourceData([]*models.Summary{m.SummaryByState}))
	d.Set("summary_by_user_activity", SetSummarySubResourceData([]*models.Summary{m.SummaryByUserActivity}))
}

func SetDetailedUsersSubResourceData(m []*models.DetailedUsers) (d []*map[string]interface{}) {
	for _, DetailedUsersModel := range m {
		if DetailedUsersModel != nil {
			properties := make(map[string]interface{})
			properties["list"] = SetDetailedUserSubResourceData(DetailedUsersModel.List)
			properties["next"] = SetCursorSubResourceData([]*models.Cursor{DetailedUsersModel.Next})
			properties["summary_by_role_distribution"] = SetSummarySubResourceData([]*models.Summary{DetailedUsersModel.SummaryByRoleDistribution})
			properties["summary_by_state"] = SetSummarySubResourceData([]*models.Summary{DetailedUsersModel.SummaryByState})
			properties["summary_by_user_activity"] = SetSummarySubResourceData([]*models.Summary{DetailedUsersModel.SummaryByUserActivity})
			d = append(d, &properties)
		}
	}
	return
}

func DetailedUsersSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"list": {
			Description: `List of users`,
			Type:        schema.TypeList, //GoType: []*DetailedUser
			Elem: &schema.Resource{
				Schema: DetailedUserSchema(),
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

		"summary_by_role_distribution": {
			Description: `User distribution by role`,
			Type:        schema.TypeList, //GoType: Summary
			Elem: &schema.Resource{
				Schema: SummarySchema(),
			},
			Optional: true,
		},

		"summary_by_state": {
			Description: `Summary of filtered users`,
			Type:        schema.TypeList, //GoType: Summary
			Elem: &schema.Resource{
				Schema: SummarySchema(),
			},
			Optional: true,
		},

		"summary_by_user_activity": {
			Description: `Count of online/offline users`,
			Type:        schema.TypeList, //GoType: Summary
			Elem: &schema.Resource{
				Schema: SummarySchema(),
			},
			Optional: true,
		},
	}
}

func GetDetailedUsersPropertyFields() (t []string) {
	return []string{
		"list",
		"next",
		"summary_by_role_distribution",
		"summary_by_state",
		"summary_by_user_activity",
	}
}
