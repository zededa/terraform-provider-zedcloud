package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func UserSummaryReportModel(d *schema.ResourceData) *models.UserSummaryReport {
	activeUsersCountInt, _ := d.Get("active_users_count").(int)
	activeUsersCount := int64(activeUsersCountInt)
	var lastLoggedinUserInfo *models.LastLoggedinUserInfo // LastLoggedinUserInfo
	lastLoggedinUserInfoInterface, lastLoggedinUserInfoIsSet := d.GetOk("last_loggedin_user_info")
	if lastLoggedinUserInfoIsSet && lastLoggedinUserInfoInterface != nil {
		lastLoggedinUserInfoMap := lastLoggedinUserInfoInterface.([]interface{})
		if len(lastLoggedinUserInfoMap) > 0 {
			lastLoggedinUserInfo = LastLoggedinUserInfoModelFromMap(lastLoggedinUserInfoMap[0].(map[string]interface{}))
		}
	}
	usersInt, _ := d.Get("users").(int)
	users := int64(usersInt)
	return &models.UserSummaryReport{
		ActiveUsersCount:     activeUsersCount,
		LastLoggedinUserInfo: lastLoggedinUserInfo,
		Users:                users,
	}
}

func UserSummaryReportModelFromMap(m map[string]interface{}) *models.UserSummaryReport {
	activeUsersCount := int64(m["active_users_count"].(int)) // int64
	var lastLoggedinUserInfo *models.LastLoggedinUserInfo    // LastLoggedinUserInfo
	lastLoggedinUserInfoInterface, lastLoggedinUserInfoIsSet := m["last_loggedin_user_info"]
	if lastLoggedinUserInfoIsSet && lastLoggedinUserInfoInterface != nil {
		lastLoggedinUserInfoMap := lastLoggedinUserInfoInterface.([]interface{})
		if len(lastLoggedinUserInfoMap) > 0 {
			lastLoggedinUserInfo = LastLoggedinUserInfoModelFromMap(lastLoggedinUserInfoMap[0].(map[string]interface{}))
		}
	}
	//
	users := int64(m["users"].(int)) // int64
	return &models.UserSummaryReport{
		ActiveUsersCount:     activeUsersCount,
		LastLoggedinUserInfo: lastLoggedinUserInfo,
		Users:                users,
	}
}

func SetUserSummaryReportResourceData(d *schema.ResourceData, m *models.UserSummaryReport) {
	d.Set("active_users_count", m.ActiveUsersCount)
	d.Set("last_loggedin_user_info", SetLastLoggedinUserInfoSubResourceData([]*models.LastLoggedinUserInfo{m.LastLoggedinUserInfo}))
	d.Set("users", m.Users)
}

func SetUserSummaryReportSubResourceData(m []*models.UserSummaryReport) (d []*map[string]interface{}) {
	for _, UserSummaryReportModel := range m {
		if UserSummaryReportModel != nil {
			properties := make(map[string]interface{})
			properties["active_users_count"] = UserSummaryReportModel.ActiveUsersCount
			properties["last_loggedin_user_info"] = SetLastLoggedinUserInfoSubResourceData([]*models.LastLoggedinUserInfo{UserSummaryReportModel.LastLoggedinUserInfo})
			properties["users"] = UserSummaryReportModel.Users
			d = append(d, &properties)
		}
	}
	return
}

func UserSummaryReportSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"active_users_count": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"last_loggedin_user_info": {
			Description: ``,
			Type:        schema.TypeList, //GoType: LastLoggedinUserInfo
			Elem: &schema.Resource{
				Schema: LastLoggedinUserInfoSchema(),
			},
			Optional: true,
		},

		"users": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},
	}
}

func GetUserSummaryReportPropertyFields() (t []string) {
	return []string{
		"active_users_count",
		"last_loggedin_user_info",
		"users",
	}
}
