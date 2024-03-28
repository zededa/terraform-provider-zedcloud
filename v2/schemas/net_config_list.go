package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func NetConfigListModel(d *schema.ResourceData) *models.NetConfigList {
	var list []*models.Network // []*NetConfig
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
			m := NetworkModelFromMap(v.(map[string]interface{}))
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
	var summary *models.Summary // Summary
	summaryInterface, summaryIsSet := d.GetOk("summary")
	if summaryIsSet && summaryInterface != nil {
		summaryMap := summaryInterface.([]interface{})
		if len(summaryMap) > 0 {
			summary = SummaryModelFromMap(summaryMap[0].(map[string]interface{}))
		}
	}
	var summaryByDist *models.Summary // Summary
	summaryByDistInterface, summaryByDistIsSet := d.GetOk("summary_by_dist")
	if summaryByDistIsSet && summaryByDistInterface != nil {
		summaryByDistMap := summaryByDistInterface.([]interface{})
		if len(summaryByDistMap) > 0 {
			summaryByDist = SummaryModelFromMap(summaryByDistMap[0].(map[string]interface{}))
		}
	}
	var summaryByKind *models.Summary // Summary
	summaryByKindInterface, summaryByKindIsSet := d.GetOk("summary_by_kind")
	if summaryByKindIsSet && summaryByKindInterface != nil {
		summaryByKindMap := summaryByKindInterface.([]interface{})
		if len(summaryByKindMap) > 0 {
			summaryByKind = SummaryModelFromMap(summaryByKindMap[0].(map[string]interface{}))
		}
	}
	var summaryByProxy *models.Summary // Summary
	summaryByProxyInterface, summaryByProxyIsSet := d.GetOk("summary_by_proxy")
	if summaryByProxyIsSet && summaryByProxyInterface != nil {
		summaryByProxyMap := summaryByProxyInterface.([]interface{})
		if len(summaryByProxyMap) > 0 {
			summaryByProxy = SummaryModelFromMap(summaryByProxyMap[0].(map[string]interface{}))
		}
	}
	return &models.NetConfigList{
		List:           list,
		Next:           next,
		Summary:        summary,
		SummaryByDist:  summaryByDist,
		SummaryByKind:  summaryByKind,
		SummaryByProxy: summaryByProxy,
	}
}

func NetConfigListModelFromMap(m map[string]interface{}) *models.NetConfigList {
	var list []*models.Network // []*NetConfig
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
			m := NetworkModelFromMap(v.(map[string]interface{}))
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
	var summary *models.Summary // Summary
	summaryInterface, summaryIsSet := m["summary"]
	if summaryIsSet && summaryInterface != nil {
		summaryMap := summaryInterface.([]interface{})
		if len(summaryMap) > 0 {
			summary = SummaryModelFromMap(summaryMap[0].(map[string]interface{}))
		}
	}
	//
	var summaryByDist *models.Summary // Summary
	summaryByDistInterface, summaryByDistIsSet := m["summary_by_dist"]
	if summaryByDistIsSet && summaryByDistInterface != nil {
		summaryByDistMap := summaryByDistInterface.([]interface{})
		if len(summaryByDistMap) > 0 {
			summaryByDist = SummaryModelFromMap(summaryByDistMap[0].(map[string]interface{}))
		}
	}
	//
	var summaryByKind *models.Summary // Summary
	summaryByKindInterface, summaryByKindIsSet := m["summary_by_kind"]
	if summaryByKindIsSet && summaryByKindInterface != nil {
		summaryByKindMap := summaryByKindInterface.([]interface{})
		if len(summaryByKindMap) > 0 {
			summaryByKind = SummaryModelFromMap(summaryByKindMap[0].(map[string]interface{}))
		}
	}
	//
	var summaryByProxy *models.Summary // Summary
	summaryByProxyInterface, summaryByProxyIsSet := m["summary_by_proxy"]
	if summaryByProxyIsSet && summaryByProxyInterface != nil {
		summaryByProxyMap := summaryByProxyInterface.([]interface{})
		if len(summaryByProxyMap) > 0 {
			summaryByProxy = SummaryModelFromMap(summaryByProxyMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.NetConfigList{
		List:           list,
		Next:           next,
		Summary:        summary,
		SummaryByDist:  summaryByDist,
		SummaryByKind:  summaryByKind,
		SummaryByProxy: summaryByProxy,
	}
}

// Update the underlying NetConfigList resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetNetConfigListResourceData(d *schema.ResourceData, m *models.NetConfigList) {
	d.Set("list", SetNetworkSubResourceData(m.List))
	d.Set("next", SetCursorSubResourceData([]*models.Cursor{m.Next}))
	d.Set("summary", SetSummarySubResourceData([]*models.Summary{m.Summary}))
	d.Set("summary_by_dist", SetSummarySubResourceData([]*models.Summary{m.SummaryByDist}))
	d.Set("summary_by_kind", SetSummarySubResourceData([]*models.Summary{m.SummaryByKind}))
	d.Set("summary_by_proxy", SetSummarySubResourceData([]*models.Summary{m.SummaryByProxy}))
}

// Iterate through and update the NetConfigList resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetNetConfigListSubResourceData(m []*models.NetConfigList) (d []*map[string]interface{}) {
	for _, NetConfigListModel := range m {
		if NetConfigListModel != nil {
			properties := make(map[string]interface{})
			properties["list"] = SetNetworkSubResourceData(NetConfigListModel.List)
			properties["next"] = SetCursorSubResourceData([]*models.Cursor{NetConfigListModel.Next})
			properties["summary"] = SetSummarySubResourceData([]*models.Summary{NetConfigListModel.Summary})
			properties["summary_by_dist"] = SetSummarySubResourceData([]*models.Summary{NetConfigListModel.SummaryByDist})
			properties["summary_by_kind"] = SetSummarySubResourceData([]*models.Summary{NetConfigListModel.SummaryByKind})
			properties["summary_by_proxy"] = SetSummarySubResourceData([]*models.Summary{NetConfigListModel.SummaryByProxy})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the NetConfigList resource defined in the Terraform configuration
func NetConfigListSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"list": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []*NetConfig
			Elem: &schema.Resource{
				Schema: Network(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"next": {
			Description: ``,
			Type:        schema.TypeList, //GoType: Cursor
			Elem: &schema.Resource{
				Schema: CursorSchema(),
			},
			Optional: true,
		},

		"summary": {
			Description: ``,
			Type:        schema.TypeList, //GoType: Summary
			Elem: &schema.Resource{
				Schema: SummarySchema(),
			},
			Optional: true,
		},

		"summary_by_dist": {
			Description: ``,
			Type:        schema.TypeList, //GoType: Summary
			Elem: &schema.Resource{
				Schema: SummarySchema(),
			},
			Optional: true,
		},

		"summary_by_kind": {
			Description: ``,
			Type:        schema.TypeList, //GoType: Summary
			Elem: &schema.Resource{
				Schema: SummarySchema(),
			},
			Optional: true,
		},

		"summary_by_proxy": {
			Description: ``,
			Type:        schema.TypeList, //GoType: Summary
			Elem: &schema.Resource{
				Schema: SummarySchema(),
			},
			Optional: true,
		},
	}
}

// Retrieve property field names for updating the NetConfigList resource
func GetNetConfigListPropertyFields() (t []string) {
	return []string{
		"list",
		"next",
		"summary",
		"summary_by_dist",
		"summary_by_kind",
		"summary_by_proxy",
	}
}
