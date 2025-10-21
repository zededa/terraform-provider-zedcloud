package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func TargetAssetGroupModel(d *schema.ResourceData) *models.TargetAssetGroup {
	groupID, _ := d.Get("group_id").(string)
	return &models.TargetAssetGroup{
		GroupID: groupID,
	}
}

func TargetAssetGroupModelFromMap(m map[string]interface{}) *models.TargetAssetGroup {
	groupID := m["group_id"].(string)
	return &models.TargetAssetGroup{
		GroupID: groupID,
	}
}

func SetTargetAssetGroupResourceData(d *schema.ResourceData, m *models.TargetAssetGroup) {
	d.Set("group_id", m.GroupID)
}

func SetTargetAssetGroupSubResourceData(m []*models.TargetAssetGroup) (d []*map[string]interface{}) {
	for _, TargetAssetGroupModel := range m {
		if TargetAssetGroupModel != nil {
			properties := make(map[string]interface{})
			properties["group_id"] = TargetAssetGroupModel.GroupID
			d = append(d, &properties)
		}
	}
	return
}

func TargetAssetGroupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"group_id": {
			Description: `unique Id of the asset group.`,
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

func GetTargetAssetGroupPropertyFields() (t []string) {
	return []string{
		"group_id",
	}
}
