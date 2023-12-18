package schemas

import (
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/zededa/terraform-provider/models"
)

func LastLoggedinUserInfoModel(d *schema.ResourceData) *models.LastLoggedinUserInfo {
	lastLoginTime, _ := d.Get("last_login_time").(strfmt.DateTime)
	name, _ := d.Get("name").(string)
	return &models.LastLoggedinUserInfo{
		LastLoginTime: lastLoginTime,
		Name:          name,
	}
}

func LastLoggedinUserInfoModelFromMap(m map[string]interface{}) *models.LastLoggedinUserInfo {
	lastLoginTime := m["last_login_time"].(strfmt.DateTime)
	name := m["name"].(string)
	return &models.LastLoggedinUserInfo{
		LastLoginTime: lastLoginTime,
		Name:          name,
	}
}

func SetLastLoggedinUserInfoResourceData(d *schema.ResourceData, m *models.LastLoggedinUserInfo) {
	d.Set("last_login_time", m.LastLoginTime)
	d.Set("name", m.Name)
}

func SetLastLoggedinUserInfoSubResourceData(m []*models.LastLoggedinUserInfo) (d []*map[string]interface{}) {
	for _, LastLoggedinUserInfoModel := range m {
		if LastLoggedinUserInfoModel != nil {
			properties := make(map[string]interface{})
			properties["last_login_time"] = LastLoggedinUserInfoModel.LastLoginTime.String()
			properties["name"] = LastLoggedinUserInfoModel.Name
			d = append(d, &properties)
		}
	}
	return
}

func LastLoggedinUserInfoSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"last_login_time": {
			Description:  ``,
			Type:         schema.TypeString,
			ValidateFunc: validation.IsRFC3339Time,
			Optional:     true,
		},

		"name": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetLastLoggedinUserInfoPropertyFields() (t []string) {
	return []string{
		"last_login_time",
		"name",
	}
}
