package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func AppProfileInfoModel(d *schema.ResourceData) *models.AppProfileInfo {
	versionInt, _ := d.Get("version").(int)
	version := int64(versionInt)
	appProfileID, _ := d.Get("app_profile_id").(string)
	return &models.AppProfileInfo{
		Version:      version,
		AppProfileID: appProfileID,
	}
}

func AppProfileInfoModelFromMap(m map[string]interface{}) *models.AppProfileInfo {
	appProfileID := m["app_profile_id"].(string)
	version := int64(m["version"].(int)) // int64
	return &models.AppProfileInfo{
		AppProfileID: appProfileID,
		Version:      version,
	}
}

func SetAppProfileInfoResourceData(d *schema.ResourceData, m *models.AppProfileInfo) {
	d.Set("app_profile_id", m.AppProfileID)
	d.Set("version", m.Version)
}

func SetAppProfileInfoSubResourceData(m []*models.AppProfileInfo) (d []*map[string]interface{}) {
	for _, AppProfileInfoModel := range m {
		if AppProfileInfoModel != nil {
			properties := make(map[string]interface{})
			properties["app_profile_id"] = AppProfileInfoModel.AppProfileID
			properties["version"] = AppProfileInfoModel.Version
			d = append(d, &properties)
		}
	}
	return
}

func AppProfileInfoSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"app_profile_id": {
			Description: `Unique ID of the app profile.`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"version": {
			Description: `Version of the app profile`,
			Type:        schema.TypeInt,
			Required:    true,
		},
	}
}

func GetAppProfileInfoPropertyFields() (t []string) {
	return []string{
		"app_profile_id",
		"version",
	}
}
