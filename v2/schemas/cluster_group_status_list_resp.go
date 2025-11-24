package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ClusterGroupStatusListRespModel(d *schema.ResourceData) *models.ClusterGroupStatusListResp {
	var list []*models.ClusterGroupStatus // []*ClusterGroupStatus
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
			m := ClusterGroupStatusModelFromMap(v.(map[string]interface{}))
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
	totalCountInt, _ := d.Get("total_count").(int)
	totalCount := int64(totalCountInt)
	return &models.ClusterGroupStatusListResp{
		List:       list,
		Next:       next,
		Result:     result,
		TotalCount: totalCount,
	}
}

func ClusterGroupStatusListRespModelFromMap(m map[string]interface{}) *models.ClusterGroupStatusListResp {
	var list []*models.ClusterGroupStatus // []*ClusterGroupStatus
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
			m := ClusterGroupStatusModelFromMap(v.(map[string]interface{}))
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
	totalCount := int64(m["total_count"].(int)) // int64
	return &models.ClusterGroupStatusListResp{
		List:       list,
		Next:       next,
		Result:     result,
		TotalCount: totalCount,
	}
}

func SetClusterGroupStatusListRespResourceData(d *schema.ResourceData, m *models.ClusterGroupStatusListResp) {
	d.Set("list", SetClusterGroupStatusSubResourceData(m.List))
	d.Set("next", SetCursorSubResourceData([]*models.Cursor{m.Next}))
	d.Set("result", SetZsrvResponseSubResourceData([]*models.ZsrvResponse{m.Result}))
	d.Set("total_count", m.TotalCount)
}

func SetClusterGroupStatusListRespSubResourceData(m []*models.ClusterGroupStatusListResp) (d []*map[string]interface{}) {
	for _, ClusterGroupStatusListRespModel := range m {
		if ClusterGroupStatusListRespModel != nil {
			properties := make(map[string]interface{})
			properties["list"] = SetClusterGroupStatusSubResourceData(ClusterGroupStatusListRespModel.List)
			properties["next"] = SetCursorSubResourceData([]*models.Cursor{ClusterGroupStatusListRespModel.Next})
			properties["result"] = SetZsrvResponseSubResourceData([]*models.ZsrvResponse{ClusterGroupStatusListRespModel.Result})
			properties["total_count"] = ClusterGroupStatusListRespModel.TotalCount
			d = append(d, &properties)
		}
	}
	return
}

func ClusterGroupStatusListRespSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"list": {
			Description: `Detailed statuses of the cluster groups`,
			Type:        schema.TypeList, //GoType: []*ClusterGroupStatus
			Elem: &schema.Resource{
				Schema: ClusterGroupStatusSchema(),
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

		"total_count": {
			Description: `Total number of cluster groups satisfying the filter criteria`,
			Type:        schema.TypeInt,
			Optional:    true,
		},
	}
}

func GetClusterGroupStatusListRespPropertyFields() (t []string) {
	return []string{
		"list",
		"next",
		"result",
		"total_count",
	}
}
