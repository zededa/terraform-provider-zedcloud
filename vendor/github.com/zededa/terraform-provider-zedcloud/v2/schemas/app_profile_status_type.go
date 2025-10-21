package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func AppProfileStatusTypeModel(d *schema.ResourceData) *models.AppProfileStatusType {
	appProfileStatusType, _ := d.Get("app_profile_status_type").(models.AppProfileStatusType)
	return &appProfileStatusType
}

func AppProfileStatusTypeModelFromMap(m map[string]interface{}) *models.AppProfileStatusType {
	appProfileStatusType := m["app_profile_status_type"].(models.AppProfileStatusType)
	return &appProfileStatusType
}

func SetAppProfileStatusTypeResourceData(d *schema.ResourceData, m *models.AppProfileStatusType) {
}

func SetAppProfileStatusTypeSubResourceData(m []*models.AppProfileStatusType) (d []*map[string]interface{}) {
	for _, AppProfileStatusTypeModel := range m {
		if AppProfileStatusTypeModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func AppProfileStatusTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func GetAppProfileStatusTypePropertyFields() (t []string) {
	return []string{}
}
