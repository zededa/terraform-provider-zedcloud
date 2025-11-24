package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ZksInstanceStatusListRespModel(d *schema.ResourceData) *models.ZksInstanceStatusListResp {
	var list []*models.ZksInstanceStatus // []*ZksInstanceStatus
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
			m := ZksInstanceStatusModelFromMap(v.(map[string]interface{}))
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
	var result *models.ZsrvResponse // ZsrvResponse
	resultInterface, resultIsSet := d.GetOk("result")
	if resultIsSet && resultInterface != nil {
		resultMap := resultInterface.([]interface{})
		if len(resultMap) > 0 {
			result = ZsrvResponseModelFromMap(resultMap[0].(map[string]interface{}))
		}
	}
	var summaries *models.ZksInstanceStatusSummaries // ZksInstanceStatusSummaries
	summariesInterface, summariesIsSet := d.GetOk("summaries")
	if summariesIsSet && summariesInterface != nil {
		summariesMap := summariesInterface.([]interface{})
		if len(summariesMap) > 0 {
			summaries = ZksInstanceStatusSummariesModelFromMap(summariesMap[0].(map[string]interface{}))
		}
	}
	totalCountInt, _ := d.Get("total_count").(int)
	totalCount := int64(totalCountInt)
	return &models.ZksInstanceStatusListResp{
		List:       list,
		Next:       next,
		Result:     result,
		Summaries:  summaries,
		TotalCount: totalCount,
	}
}

func ZksInstanceStatusListRespModelFromMap(m map[string]interface{}) *models.ZksInstanceStatusListResp {
	var list []*models.ZksInstanceStatus // []*ZksInstanceStatus
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
			m := ZksInstanceStatusModelFromMap(v.(map[string]interface{}))
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
	var result *models.ZsrvResponse // ZsrvResponse
	resultInterface, resultIsSet := m["result"]
	if resultIsSet && resultInterface != nil {
		resultMap := resultInterface.([]interface{})
		if len(resultMap) > 0 {
			result = ZsrvResponseModelFromMap(resultMap[0].(map[string]interface{}))
		}
	}
	//
	var summaries *models.ZksInstanceStatusSummaries // ZksInstanceStatusSummaries
	summariesInterface, summariesIsSet := m["summaries"]
	if summariesIsSet && summariesInterface != nil {
		summariesMap := summariesInterface.([]interface{})
		if len(summariesMap) > 0 {
			summaries = ZksInstanceStatusSummariesModelFromMap(summariesMap[0].(map[string]interface{}))
		}
	}
	//
	totalCount := int64(m["total_count"].(int)) // int64
	return &models.ZksInstanceStatusListResp{
		List:       list,
		Next:       next,
		Result:     result,
		Summaries:  summaries,
		TotalCount: totalCount,
	}
}

func SetZksInstanceStatusListRespResourceData(d *schema.ResourceData, m *models.ZksInstanceStatusListResp) {
	d.Set("list", SetZksInstanceStatusSubResourceData(m.List))
	d.Set("next", SetCursorSubResourceData([]*models.Cursor{m.Next}))
	d.Set("result", SetZsrvResponseSubResourceData([]*models.ZsrvResponse{m.Result}))
	d.Set("summaries", SetZksInstanceStatusSummariesSubResourceData([]*models.ZksInstanceStatusSummaries{m.Summaries}))
	d.Set("total_count", m.TotalCount)
}

func SetZksInstanceStatusListRespSubResourceData(m []*models.ZksInstanceStatusListResp) (d []*map[string]interface{}) {
	for _, ZksInstanceStatusListRespModel := range m {
		if ZksInstanceStatusListRespModel != nil {
			properties := make(map[string]interface{})
			properties["list"] = SetZksInstanceStatusSubResourceData(ZksInstanceStatusListRespModel.List)
			properties["next"] = SetCursorSubResourceData([]*models.Cursor{ZksInstanceStatusListRespModel.Next})
			properties["result"] = SetZsrvResponseSubResourceData([]*models.ZsrvResponse{ZksInstanceStatusListRespModel.Result})
			properties["summaries"] = SetZksInstanceStatusSummariesSubResourceData([]*models.ZksInstanceStatusSummaries{ZksInstanceStatusListRespModel.Summaries})
			properties["total_count"] = ZksInstanceStatusListRespModel.TotalCount
			d = append(d, &properties)
		}
	}
	return
}

func ZksInstanceStatusListRespSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"list": {
			Description: `Detailed statuses of the zks instances`,
			Type:        schema.TypeList, //GoType: []*ZksInstanceStatus
			Elem: &schema.Resource{
				Schema: ZksInstanceStatusSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"next": {
			Description: `Cursor for pagination`,
			Type:        schema.TypeList, //GoType: Cursor
			Elem: &schema.Resource{
				Schema: CursorSchema(),
			},
			Optional: true,
		},

		"result": {
			Description: `Response result`,
			Type:        schema.TypeList, //GoType: ZsrvResponse
			Elem: &schema.Resource{
				Schema: ZsrvResponseSchema(),
			},
			Optional: true,
		},

		"summaries": {
			Description: `Summaries of the ZKS instance statuses`,
			Type:        schema.TypeList, //GoType: ZksInstanceStatusSummaries
			Elem: &schema.Resource{
				Schema: ZksInstanceStatusSummariesSchema(),
			},
			Optional: true,
		},

		"total_count": {
			Description: `Total count of zks instances matching the filter`,
			Type:        schema.TypeInt,
			Optional:    true,
		},
	}
}

func GetZksInstanceStatusListRespPropertyFields() (t []string) {
	return []string{
		"list",
		"next",
		"result",
		"summaries",
		"total_count",
	}
}
