package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func NetInstListModel(d *schema.ResourceData) *models.NetInstList {
	var cfgList []*models.NetInstConfig // []*NetInstConfig
	cfgListInterface, cfgListIsSet := d.GetOk("cfg_list")
	if cfgListIsSet {
		var items []interface{}
		if listItems, isList := cfgListInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = cfgListInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := NetworkInstanceModelFromMap(v.(map[string]interface{}))
			cfgList = append(cfgList, m)
		}
	}
	var list []*models.NetInstShortConfig // []*NetInstShortConfig
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
			m := NetInstShortConfigModelFromMap(v.(map[string]interface{}))
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
	var summaryByAddressType *models.Summary // Summary
	summaryByAddressTypeInterface, summaryByAddressTypeIsSet := d.GetOk("summary_by_address_type")
	if summaryByAddressTypeIsSet && summaryByAddressTypeInterface != nil {
		summaryByAddressTypeMap := summaryByAddressTypeInterface.([]interface{})
		if len(summaryByAddressTypeMap) > 0 {
			summaryByAddressType = SummaryModelFromMap(summaryByAddressTypeMap[0].(map[string]interface{}))
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
	return &models.NetInstList{
		CfgList:              cfgList,
		List:                 list,
		Next:                 next,
		SummaryByAddressType: summaryByAddressType,
		SummaryByKind:        summaryByKind,
	}
}

func NetInstListModelFromMap(m map[string]interface{}) *models.NetInstList {
	var cfgList []*models.NetInstConfig // []*NetInstConfig
	cfgListInterface, cfgListIsSet := m["cfg_list"]
	if cfgListIsSet {
		var items []interface{}
		if listItems, isList := cfgListInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = cfgListInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := NetworkInstanceModelFromMap(v.(map[string]interface{}))
			cfgList = append(cfgList, m)
		}
	}
	var list []*models.NetInstShortConfig // []*NetInstShortConfig
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
			m := NetInstShortConfigModelFromMap(v.(map[string]interface{}))
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
	var summaryByAddressType *models.Summary // Summary
	summaryByAddressTypeInterface, summaryByAddressTypeIsSet := m["summary_by_address_type"]
	if summaryByAddressTypeIsSet && summaryByAddressTypeInterface != nil {
		summaryByAddressTypeMap := summaryByAddressTypeInterface.([]interface{})
		if len(summaryByAddressTypeMap) > 0 {
			summaryByAddressType = SummaryModelFromMap(summaryByAddressTypeMap[0].(map[string]interface{}))
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
	return &models.NetInstList{
		CfgList:              cfgList,
		List:                 list,
		Next:                 next,
		SummaryByAddressType: summaryByAddressType,
		SummaryByKind:        summaryByKind,
	}
}

// Update the underlying NetInstList resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetNetInstListResourceData(d *schema.ResourceData, m *models.NetInstList) {
	d.Set("cfg_list", SetNetworkInstanceSubResourceData(m.CfgList))
	d.Set("list", SetNetInstShortConfigSubResourceData(m.List))
	d.Set("next", SetCursorSubResourceData([]*models.Cursor{m.Next}))
	d.Set("summary_by_address_type", SetSummarySubResourceData([]*models.Summary{m.SummaryByAddressType}))
	d.Set("summary_by_kind", SetSummarySubResourceData([]*models.Summary{m.SummaryByKind}))
}

// Iterate through and update the NetInstList resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetNetInstListSubResourceData(m []*models.NetInstList) (d []*map[string]interface{}) {
	for _, NetInstListModel := range m {
		if NetInstListModel != nil {
			properties := make(map[string]interface{})
			properties["cfg_list"] = SetNetworkInstanceSubResourceData(NetInstListModel.CfgList)
			properties["list"] = SetNetInstShortConfigSubResourceData(NetInstListModel.List)
			properties["next"] = SetCursorSubResourceData([]*models.Cursor{NetInstListModel.Next})
			properties["summary_by_address_type"] = SetSummarySubResourceData([]*models.Summary{NetInstListModel.SummaryByAddressType})
			properties["summary_by_kind"] = SetSummarySubResourceData([]*models.Summary{NetInstListModel.SummaryByKind})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the NetInstList resource defined in the Terraform configuration
func NetInstListSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cfg_list": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []*NetInstConfig
			Elem: &schema.Resource{
				Schema: NetworkInstance(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"list": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []*NetInstShortConfig
			Elem: &schema.Resource{
				Schema: NetInstShortConfigSchema(),
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

		"summary_by_address_type": {
			Description: `Summary information about netinstance list records by addressing type.`,
			Type:        schema.TypeList, //GoType: Summary
			Elem: &schema.Resource{
				Schema: SummarySchema(),
			},
			Optional: true,
		},

		"summary_by_kind": {
			Description: `Summary information about netinstance list records by network instance kind.`,
			Type:        schema.TypeList, //GoType: Summary
			Elem: &schema.Resource{
				Schema: SummarySchema(),
			},
			Optional: true,
		},
	}
}

// Retrieve property field names for updating the NetInstList resource
func GetNetInstListPropertyFields() (t []string) {
	return []string{
		"cfg_list",
		"list",
		"next",
		"summary_by_address_type",
		"summary_by_kind",
	}
}
