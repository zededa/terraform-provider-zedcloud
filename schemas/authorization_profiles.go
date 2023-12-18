package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AuthorizationProfilesModel(d *schema.ResourceData) *models.AuthorizationProfiles {
	var list []*models.AuthorizationProfile // []*AuthorizationProfile
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
			m := AuthorizationProfileModelFromMap(v.(map[string]interface{}))
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
	return &models.AuthorizationProfiles{
		List:           list,
		Next:           next,
		SummaryByState: summaryByState,
	}
}

func AuthorizationProfilesModelFromMap(m map[string]interface{}) *models.AuthorizationProfiles {
	var list []*models.AuthorizationProfile // []*AuthorizationProfile
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
			m := AuthorizationProfileModelFromMap(v.(map[string]interface{}))
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
	return &models.AuthorizationProfiles{
		List:           list,
		Next:           next,
		SummaryByState: summaryByState,
	}
}

func SetAuthorizationProfilesResourceData(d *schema.ResourceData, m *models.AuthorizationProfiles) {
	d.Set("list", SetAuthorizationProfileSubResourceData(m.List))
	d.Set("next", SetCursorSubResourceData([]*models.Cursor{m.Next}))
	d.Set("summary_by_state", SetSummarySubResourceData([]*models.Summary{m.SummaryByState}))
}

func SetAuthorizationProfilesSubResourceData(m []*models.AuthorizationProfiles) (d []*map[string]interface{}) {
	for _, AuthorizationProfilesModel := range m {
		if AuthorizationProfilesModel != nil {
			properties := make(map[string]interface{})
			properties["list"] = SetAuthorizationProfileSubResourceData(AuthorizationProfilesModel.List)
			properties["next"] = SetCursorSubResourceData([]*models.Cursor{AuthorizationProfilesModel.Next})
			properties["summary_by_state"] = SetSummarySubResourceData([]*models.Summary{AuthorizationProfilesModel.SummaryByState})
			d = append(d, &properties)
		}
	}
	return
}

func AuthorizationProfilesSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"list": {
			Description: `List of profiles`,
			Type:        schema.TypeList, //GoType: []*AuthorizationProfile
			Elem: &schema.Resource{
				Schema: AuthorizationProfileSchema(),
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
			Description: `Summary of filtered profiles`,
			Type:        schema.TypeList, //GoType: Summary
			Elem: &schema.Resource{
				Schema: SummarySchema(),
			},
			Optional: true,
		},
	}
}

func GetAuthorizationProfilesPropertyFields() (t []string) {
	return []string{
		"list",
		"next",
		"summary_by_state",
	}
}
