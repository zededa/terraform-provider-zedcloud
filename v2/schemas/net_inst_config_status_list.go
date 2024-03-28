package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func NetInstConfigStatusListModel(d *schema.ResourceData) *models.NetInstConfigStatusList {
	var list []*models.NetInstConfigStatus // []*NetInstConfigStatus
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
			m := NetInstConfigStatusModelFromMap(v.(map[string]interface{}))
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
	return &models.NetInstConfigStatusList{
		List:                 list,
		Next:                 next,
		SummaryByAddressType: summaryByAddressType,
		SummaryByKind:        summaryByKind,
	}
}

func NetInstConfigStatusListModelFromMap(m map[string]interface{}) *models.NetInstConfigStatusList {
	var list []*models.NetInstConfigStatus // []*NetInstConfigStatus
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
			m := NetInstConfigStatusModelFromMap(v.(map[string]interface{}))
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
	return &models.NetInstConfigStatusList{
		List:                 list,
		Next:                 next,
		SummaryByAddressType: summaryByAddressType,
		SummaryByKind:        summaryByKind,
	}
}

// Update the underlying NetInstConfigStatusList resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetNetInstConfigStatusListResourceData(d *schema.ResourceData, m *models.NetInstConfigStatusList) {
	d.Set("list", SetNetInstConfigStatusSubResourceData(m.List))
	d.Set("next", SetCursorSubResourceData([]*models.Cursor{m.Next}))
	d.Set("summary_by_address_type", SetSummarySubResourceData([]*models.Summary{m.SummaryByAddressType}))
	d.Set("summary_by_kind", SetSummarySubResourceData([]*models.Summary{m.SummaryByKind}))
}

// Iterate through and update the NetInstConfigStatusList resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetNetInstConfigStatusListSubResourceData(m []*models.NetInstConfigStatusList) (d []*map[string]interface{}) {
	for _, NetInstConfigStatusListModel := range m {
		if NetInstConfigStatusListModel != nil {
			properties := make(map[string]interface{})
			properties["list"] = SetNetInstConfigStatusSubResourceData(NetInstConfigStatusListModel.List)
			properties["next"] = SetCursorSubResourceData([]*models.Cursor{NetInstConfigStatusListModel.Next})
			properties["summary_by_address_type"] = SetSummarySubResourceData([]*models.Summary{NetInstConfigStatusListModel.SummaryByAddressType})
			properties["summary_by_kind"] = SetSummarySubResourceData([]*models.Summary{NetInstConfigStatusListModel.SummaryByKind})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the NetInstConfigStatusList resource defined in the Terraform configuration
func NetInstConfigStatusListSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"list": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []*NetInstConfigStatus
			Elem: &schema.Resource{
				Schema: NetInstConfigStatusSchema(),
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
			Description: `Summary information about netinstance config status list records by addressing type.`,
			Type:        schema.TypeList, //GoType: Summary
			Elem: &schema.Resource{
				Schema: SummarySchema(),
			},
			Optional: true,
		},

		"summary_by_kind": {
			Description: `Summary information about netinstance config status list records by network instance kind.`,
			Type:        schema.TypeList, //GoType: Summary
			Elem: &schema.Resource{
				Schema: SummarySchema(),
			},
			Optional: true,
		},
	}
}

// Retrieve property field names for updating the NetInstConfigStatusList resource
func GetNetInstConfigStatusListPropertyFields() (t []string) {
	return []string{
		"list",
		"next",
		"summary_by_address_type",
		"summary_by_kind",
	}
}
