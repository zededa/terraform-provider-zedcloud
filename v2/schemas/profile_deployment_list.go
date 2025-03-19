package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ProfileDeploymentListModel(d *schema.ResourceData) *models.ProfileDeploymentList {
	var list []*models.ProfileDeployment // []*ProfileDeployment
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
			m := ProfileDeploymentModelFromMap(v.(map[string]interface{}))
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
	return &models.ProfileDeploymentList{
		List: list,
		Next: next,
	}
}

func ProfileDeploymentListModelFromMap(m map[string]interface{}) *models.ProfileDeploymentList {
	var list []*models.ProfileDeployment // []*ProfileDeployment
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
			m := ProfileDeploymentModelFromMap(v.(map[string]interface{}))
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
	return &models.ProfileDeploymentList{
		List: list,
		Next: next,
	}
}

func SetProfileDeploymentListResourceData(d *schema.ResourceData, m *models.ProfileDeploymentList) {
	d.Set("list", SetProfileDeploymentSubResourceData(m.List))
	d.Set("next", SetCursorSubResourceData([]*models.Cursor{m.Next}))
}

func SetProfileDeploymentListSubResourceData(m []*models.ProfileDeploymentList) (d []*map[string]interface{}) {
	for _, ProfileDeploymentListModel := range m {
		if ProfileDeploymentListModel != nil {
			properties := make(map[string]interface{})
			properties["list"] = SetProfileDeploymentSubResourceData(ProfileDeploymentListModel.List)
			properties["next"] = SetCursorSubResourceData([]*models.Cursor{ProfileDeploymentListModel.Next})
			d = append(d, &properties)
		}
	}
	return
}

func ProfileDeploymentListSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"list": {
			Description: `deployment summary list`,
			Type:        schema.TypeList, //GoType: []*ProfileDeployment
			Elem: &schema.Resource{
				Schema: ProfileDeploymentSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"next": {
			Description: `cursor next`,
			Type:        schema.TypeList, //GoType: Cursor
			Elem: &schema.Resource{
				Schema: CursorSchema(),
			},
			Optional: true,
		},
	}
}

func GetProfileDeploymentListPropertyFields() (t []string) {
	return []string{
		"list",
		"next",
	}
}
