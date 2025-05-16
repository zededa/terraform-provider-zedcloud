package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ZserviceAssetGroupListResponseModel(d *schema.ResourceData) *models.ZserviceAssetGroupListResponse {
	var assetGroups []*models.ZserviceAssetGroupReadRO // []*ZserviceAssetGroupReadRO
	assetGroupsInterface, assetGroupsIsSet := d.GetOk("asset_groups")
	if assetGroupsIsSet {
		var items []interface{}
		if listItems, isList := assetGroupsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = assetGroupsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := ZserviceAssetGroupReadROModelFromMap(v.(map[string]interface{}))
			assetGroups = append(assetGroups, m)
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
	totalCountInt, _ := d.Get("total_count").(int)
	totalCount := int64(totalCountInt)
	return &models.ZserviceAssetGroupListResponse{
		AssetGroups: assetGroups,
		Next:        next,
		TotalCount:  totalCount,
	}
}

func ZserviceAssetGroupListResponseModelFromMap(m map[string]interface{}) *models.ZserviceAssetGroupListResponse {
	var assetGroups []*models.ZserviceAssetGroupReadRO // []*ZserviceAssetGroupReadRO
	assetGroupsInterface, assetGroupsIsSet := m["asset_groups"]
	if assetGroupsIsSet {
		var items []interface{}
		if listItems, isList := assetGroupsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = assetGroupsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := ZserviceAssetGroupReadROModelFromMap(v.(map[string]interface{}))
			assetGroups = append(assetGroups, m)
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
	totalCount := int64(m["total_count"].(int)) // int64
	return &models.ZserviceAssetGroupListResponse{
		AssetGroups: assetGroups,
		Next:        next,
		TotalCount:  totalCount,
	}
}

func SetZserviceAssetGroupListResponseResourceData(d *schema.ResourceData, m *models.ZserviceAssetGroupListResponse) {
	d.Set("asset_groups", SetZserviceAssetGroupReadROSubResourceData(m.AssetGroups))
	d.Set("next", SetCursorSubResourceData([]*models.Cursor{m.Next}))
	d.Set("total_count", m.TotalCount)
}

func SetZserviceAssetGroupListResponseSubResourceData(m []*models.ZserviceAssetGroupListResponse) (d []*map[string]interface{}) {
	for _, ZserviceAssetGroupListResponseModel := range m {
		if ZserviceAssetGroupListResponseModel != nil {
			properties := make(map[string]interface{})
			properties["asset_groups"] = SetZserviceAssetGroupReadROSubResourceData(ZserviceAssetGroupListResponseModel.AssetGroups)
			properties["next"] = SetCursorSubResourceData([]*models.Cursor{ZserviceAssetGroupListResponseModel.Next})
			properties["total_count"] = ZserviceAssetGroupListResponseModel.TotalCount
			d = append(d, &properties)
		}
	}
	return
}

func ZserviceAssetGroupListResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"asset_groups": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []*ZserviceAssetGroupReadRO
			Elem: &schema.Resource{
				Schema: ZserviceAssetGroupReadROSchema(),
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

		"total_count": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},
	}
}

func GetZserviceAssetGroupListResponsePropertyFields() (t []string) {
	return []string{
		"asset_groups",
		"next",
		"total_count",
	}
}
