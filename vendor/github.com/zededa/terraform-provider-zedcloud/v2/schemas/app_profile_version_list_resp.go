package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func AppProfileVersionListRespModel(d *schema.ResourceData) *models.AppProfileVersionListResp {
	var next *models.Cursor // Cursor
	nextInterface, nextIsSet := d.GetOk("next")
	if nextIsSet && nextInterface != nil {
		nextMap := nextInterface.([]interface{})
		if len(nextMap) > 0 {
			next = CursorModelFromMap(nextMap[0].(map[string]interface{}))
		}
	}
	var versionList []*models.AppProfileVersion // []*AppProfileVersion
	versionListInterface, versionListIsSet := d.GetOk("version_list")
	if versionListIsSet {
		var items []interface{}
		if listItems, isList := versionListInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = versionListInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := AppProfileVersionModelFromMap(v.(map[string]interface{}))
			versionList = append(versionList, m)
		}
	}
	return &models.AppProfileVersionListResp{
		Next:        next,
		VersionList: versionList,
	}
}

func AppProfileVersionListRespModelFromMap(m map[string]interface{}) *models.AppProfileVersionListResp {
	var next *models.Cursor // Cursor
	nextInterface, nextIsSet := m["next"]
	if nextIsSet && nextInterface != nil {
		nextMap := nextInterface.([]interface{})
		if len(nextMap) > 0 {
			next = CursorModelFromMap(nextMap[0].(map[string]interface{}))
		}
	}
	//
	var versionList []*models.AppProfileVersion // []*AppProfileVersion
	versionListInterface, versionListIsSet := m["version_list"]
	if versionListIsSet {
		var items []interface{}
		if listItems, isList := versionListInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = versionListInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := AppProfileVersionModelFromMap(v.(map[string]interface{}))
			versionList = append(versionList, m)
		}
	}
	return &models.AppProfileVersionListResp{
		Next:        next,
		VersionList: versionList,
	}
}

func SetAppProfileVersionListRespResourceData(d *schema.ResourceData, m *models.AppProfileVersionListResp) {
	d.Set("next", SetCursorSubResourceData([]*models.Cursor{m.Next}))
	d.Set("version_list", SetAppProfileVersionSubResourceData(m.VersionList))
}

func SetAppProfileVersionListRespSubResourceData(m []*models.AppProfileVersionListResp) (d []*map[string]interface{}) {
	for _, AppProfileVersionListRespModel := range m {
		if AppProfileVersionListRespModel != nil {
			properties := make(map[string]interface{})
			properties["next"] = SetCursorSubResourceData([]*models.Cursor{AppProfileVersionListRespModel.Next})
			properties["version_list"] = SetAppProfileVersionSubResourceData(AppProfileVersionListRespModel.VersionList)
			d = append(d, &properties)
		}
	}
	return
}

func AppProfileVersionListRespSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"next": {
			Description: `cursor next`,
			Type:        schema.TypeList, //GoType: Cursor
			Elem: &schema.Resource{
				Schema: CursorSchema(),
			},
			Optional: true,
		},

		"version_list": {
			Description: `App profile version list response`,
			Type:        schema.TypeList, //GoType: []*AppProfileVersion
			Elem: &schema.Resource{
				Schema: AppProfileVersionSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
	}
}

func GetAppProfileVersionListRespPropertyFields() (t []string) {
	return []string{
		"next",
		"version_list",
	}
}
