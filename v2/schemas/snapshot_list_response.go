package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func SnapshotListResponseModel(d *schema.ResourceData) *models.SnapshotListResponse {
	var snapshots []*models.SnapshotGetResponse // []*SnapshotGetResponse
	snapshotsInterface, snapshotsIsSet := d.GetOk("snapshots")
	if snapshotsIsSet {
		var items []interface{}
		if listItems, isList := snapshotsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = snapshotsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := SnapshotGetResponseModelFromMap(v.(map[string]interface{}))
			snapshots = append(snapshots, m)
		}
	}
	totalCountInt, _ := d.Get("total_count").(int)
	totalCount := int32(totalCountInt)
	return &models.SnapshotListResponse{
		Snapshots:  snapshots,
		TotalCount: totalCount,
	}
}

func SnapshotListResponseModelFromMap(m map[string]interface{}) *models.SnapshotListResponse {
	var snapshots []*models.SnapshotGetResponse // []*SnapshotGetResponse
	snapshotsInterface, snapshotsIsSet := m["snapshots"]
	if snapshotsIsSet {
		var items []interface{}
		if listItems, isList := snapshotsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = snapshotsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := SnapshotGetResponseModelFromMap(v.(map[string]interface{}))
			snapshots = append(snapshots, m)
		}
	}
	totalCount := int32(m["total_count"].(int)) // int32
	return &models.SnapshotListResponse{
		Snapshots:  snapshots,
		TotalCount: totalCount,
	}
}

func SetSnapshotListResponseResourceData(d *schema.ResourceData, m *models.SnapshotListResponse) {
	d.Set("snapshots", SetSnapshotGetResponseSubResourceData(m.Snapshots))
	d.Set("total_count", m.TotalCount)
}

func SetSnapshotListResponseSubResourceData(m []*models.SnapshotListResponse) (d []*map[string]interface{}) {
	for _, SnapshotListResponseModel := range m {
		if SnapshotListResponseModel != nil {
			properties := make(map[string]interface{})
			properties["snapshots"] = SetSnapshotGetResponseSubResourceData(SnapshotListResponseModel.Snapshots)
			properties["total_count"] = SnapshotListResponseModel.TotalCount
			d = append(d, &properties)
		}
	}
	return
}

func SnapshotListResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"snapshots": {
			Description: `List of snapshot entries`,
			Type:        schema.TypeList, //GoType: []*SnapshotGetResponse
			Elem: &schema.Resource{
				Schema: SnapshotGetResponseSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"total_count": {
			Description: `Total count of snapshots`,
			Type:        schema.TypeInt,
			Optional:    true,
		},
	}
}

func GetSnapshotListResponsePropertyFields() (t []string) {
	return []string{
		"snapshots",
		"total_count",
	}
}
