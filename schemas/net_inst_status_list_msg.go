package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func NetInstStatusListMsgModel(d *schema.ResourceData) *models.NetInstStatusListMsg {
	var list []*models.NetInstStatusSummaryMsg // []*NetInstStatusSummaryMsg
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
			m := NetInstStatusSummaryMsgModelFromMap(v.(map[string]interface{}))
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
	return &models.NetInstStatusListMsg{
		List:                 list,
		Next:                 next,
		SummaryByAddressType: summaryByAddressType,
		SummaryByKind:        summaryByKind,
	}
}

func NetInstStatusListMsgModelFromMap(m map[string]interface{}) *models.NetInstStatusListMsg {
	var list []*models.NetInstStatusSummaryMsg // []*NetInstStatusSummaryMsg
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
			m := NetInstStatusSummaryMsgModelFromMap(v.(map[string]interface{}))
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
	return &models.NetInstStatusListMsg{
		List:                 list,
		Next:                 next,
		SummaryByAddressType: summaryByAddressType,
		SummaryByKind:        summaryByKind,
	}
}

// Update the underlying NetInstStatusListMsg resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetNetInstStatusListMsgResourceData(d *schema.ResourceData, m *models.NetInstStatusListMsg) {
	d.Set("list", SetNetInstStatusSummaryMsgSubResourceData(m.List))
	d.Set("next", SetCursorSubResourceData([]*models.Cursor{m.Next}))
	d.Set("summary_by_address_type", SetSummarySubResourceData([]*models.Summary{m.SummaryByAddressType}))
	d.Set("summary_by_kind", SetSummarySubResourceData([]*models.Summary{m.SummaryByKind}))
}

// Iterate through and update the NetInstStatusListMsg resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetNetInstStatusListMsgSubResourceData(m []*models.NetInstStatusListMsg) (d []*map[string]interface{}) {
	for _, NetInstStatusListMsgModel := range m {
		if NetInstStatusListMsgModel != nil {
			properties := make(map[string]interface{})
			properties["list"] = SetNetInstStatusSummaryMsgSubResourceData(NetInstStatusListMsgModel.List)
			properties["next"] = SetCursorSubResourceData([]*models.Cursor{NetInstStatusListMsgModel.Next})
			properties["summary_by_address_type"] = SetSummarySubResourceData([]*models.Summary{NetInstStatusListMsgModel.SummaryByAddressType})
			properties["summary_by_kind"] = SetSummarySubResourceData([]*models.Summary{NetInstStatusListMsgModel.SummaryByKind})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the NetInstStatusListMsg resource defined in the Terraform configuration
func NetInstStatusListMsgSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"list": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []*NetInstStatusSummaryMsg
			Elem: &schema.Resource{
				Schema: NetInstStatusSummaryMsgSchema(),
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
			Description: `Summary information about netinstance status list records by addressing type.`,
			Type:        schema.TypeList, //GoType: Summary
			Elem: &schema.Resource{
				Schema: SummarySchema(),
			},
			Optional: true,
		},

		"summary_by_kind": {
			Description: `Summary information about netinstance status list records by network instance kind.`,
			Type:        schema.TypeList, //GoType: Summary
			Elem: &schema.Resource{
				Schema: SummarySchema(),
			},
			Optional: true,
		},
	}
}

// Retrieve property field names for updating the NetInstStatusListMsg resource
func GetNetInstStatusListMsgPropertyFields() (t []string) {
	return []string{
		"list",
		"next",
		"summary_by_address_type",
		"summary_by_kind",
	}
}
