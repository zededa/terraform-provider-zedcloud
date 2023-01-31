package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func VolInstListModel(d *schema.ResourceData) *models.VolInstList {
	var cfgList []*models.VolumeInstConfig // []*VolInstConfig
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
			m := VolumeInstanceModelFromMap(v.(map[string]interface{}))
			cfgList = append(cfgList, m)
		}
	}
	var list []*models.VolInstShortConfig // []*VolInstShortConfig
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
			m := VolInstShortConfigModelFromMap(v.(map[string]interface{}))
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
	var summaryByType *models.Summary // Summary
	summaryByTypeInterface, summaryByTypeIsSet := d.GetOk("summary_by_type")
	if summaryByTypeIsSet && summaryByTypeInterface != nil {
		summaryByTypeMap := summaryByTypeInterface.([]interface{})
		if len(summaryByTypeMap) > 0 {
			summaryByType = SummaryModelFromMap(summaryByTypeMap[0].(map[string]interface{}))
		}
	}
	return &models.VolInstList{
		CfgList:       cfgList,
		List:          list,
		Next:          next,
		SummaryByType: summaryByType,
	}
}

func VolInstListModelFromMap(m map[string]interface{}) *models.VolInstList {
	var cfgList []*models.VolumeInstConfig // []*VolInstConfig
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
			m := VolumeInstanceModelFromMap(v.(map[string]interface{}))
			cfgList = append(cfgList, m)
		}
	}
	var list []*models.VolInstShortConfig // []*VolInstShortConfig
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
			m := VolInstShortConfigModelFromMap(v.(map[string]interface{}))
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
	var summaryByType *models.Summary // Summary
	summaryByTypeInterface, summaryByTypeIsSet := m["summary_by_type"]
	if summaryByTypeIsSet && summaryByTypeInterface != nil {
		summaryByTypeMap := summaryByTypeInterface.([]interface{})
		if len(summaryByTypeMap) > 0 {
			summaryByType = SummaryModelFromMap(summaryByTypeMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.VolInstList{
		CfgList:       cfgList,
		List:          list,
		Next:          next,
		SummaryByType: summaryByType,
	}
}

// Update the underlying VolInstList resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetVolInstListResourceData(d *schema.ResourceData, m *models.VolInstList) {
	d.Set("cfg_list", SetVolumeInstanceSubResourceData(m.CfgList))
	d.Set("list", SetVolInstShortConfigSubResourceData(m.List))
	d.Set("next", SetCursorSubResourceData([]*models.Cursor{m.Next}))
	d.Set("summary_by_type", SetSummarySubResourceData([]*models.Summary{m.SummaryByType}))
}

// Iterate through and update the VolInstList resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetVolInstListSubResourceData(m []*models.VolInstList) (d []*map[string]interface{}) {
	for _, VolInstListModel := range m {
		if VolInstListModel != nil {
			properties := make(map[string]interface{})
			properties["cfg_list"] = SetVolumeInstanceSubResourceData(VolInstListModel.CfgList)
			properties["list"] = SetVolInstShortConfigSubResourceData(VolInstListModel.List)
			properties["next"] = SetCursorSubResourceData([]*models.Cursor{VolInstListModel.Next})
			properties["summary_by_type"] = SetSummarySubResourceData([]*models.Summary{VolInstListModel.SummaryByType})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the VolInstList resource defined in the Terraform configuration
func VolInstListSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cfg_list": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []*VolInstConfig
			Elem: &schema.Resource{
				Schema: VolumeInstance(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"list": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []*VolInstShortConfig
			Elem: &schema.Resource{
				Schema: VolInstShortConfigSchema(),
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

		"summary_by_type": {
			Description: ``,
			Type:        schema.TypeList, //GoType: Summary
			Elem: &schema.Resource{
				Schema: SummarySchema(),
			},
			Optional: true,
		},
	}
}

// Retrieve property field names for updating the VolInstList resource
func GetVolInstListPropertyFields() (t []string) {
	return []string{
		"cfg_list",
		"list",
		"next",
		"summary_by_type",
	}
}
