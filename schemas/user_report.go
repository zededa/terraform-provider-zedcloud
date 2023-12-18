package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func UserReportModel(d *schema.ResourceData) *models.UserReport {
	entpID, _ := d.Get("entp_id").(string)
	error, _ := d.Get("error").(string)
	var userSummaryReport *models.UserSummaryReport // UserSummaryReport
	userSummaryReportInterface, userSummaryReportIsSet := d.GetOk("user_summary_report")
	if userSummaryReportIsSet && userSummaryReportInterface != nil {
		userSummaryReportMap := userSummaryReportInterface.([]interface{})
		if len(userSummaryReportMap) > 0 {
			userSummaryReport = UserSummaryReportModelFromMap(userSummaryReportMap[0].(map[string]interface{}))
		}
	}
	return &models.UserReport{
		EntpID:            entpID,
		Error:             error,
		UserSummaryReport: userSummaryReport,
	}
}

func UserReportModelFromMap(m map[string]interface{}) *models.UserReport {
	entpID := m["entp_id"].(string)
	error := m["error"].(string)
	var userSummaryReport *models.UserSummaryReport // UserSummaryReport
	userSummaryReportInterface, userSummaryReportIsSet := m["user_summary_report"]
	if userSummaryReportIsSet && userSummaryReportInterface != nil {
		userSummaryReportMap := userSummaryReportInterface.([]interface{})
		if len(userSummaryReportMap) > 0 {
			userSummaryReport = UserSummaryReportModelFromMap(userSummaryReportMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.UserReport{
		EntpID:            entpID,
		Error:             error,
		UserSummaryReport: userSummaryReport,
	}
}

func SetUserReportResourceData(d *schema.ResourceData, m *models.UserReport) {
	d.Set("entp_id", m.EntpID)
	d.Set("error", m.Error)
	d.Set("user_summary_report", SetUserSummaryReportSubResourceData([]*models.UserSummaryReport{m.UserSummaryReport}))
}

func SetUserReportSubResourceData(m []*models.UserReport) (d []*map[string]interface{}) {
	for _, UserReportModel := range m {
		if UserReportModel != nil {
			properties := make(map[string]interface{})
			properties["entp_id"] = UserReportModel.EntpID
			properties["error"] = UserReportModel.Error
			properties["user_summary_report"] = SetUserSummaryReportSubResourceData([]*models.UserSummaryReport{UserReportModel.UserSummaryReport})
			d = append(d, &properties)
		}
	}
	return
}

func UserReportSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"entp_id": {
			Description: `Enterprise id for which we want to get summary report for all objects`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"error": {
			Description: `Error while fetching report for enterprise, if any`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"user_summary_report": {
			Description: `Enterprise user report`,
			Type:        schema.TypeList, //GoType: UserSummaryReport
			Elem: &schema.Resource{
				Schema: UserSummaryReportSchema(),
			},
			Optional: true,
		},
	}
}

func GetUserReportPropertyFields() (t []string) {
	return []string{
		"entp_id",
		"error",
		"user_summary_report",
	}
}
