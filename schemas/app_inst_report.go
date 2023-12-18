package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AppInstReportModel(d *schema.ResourceData) *models.AppInstReport {
	var appInstSummaryReport *models.AppInstSummaryReport // AppInstSummaryReport
	appInstSummaryReportInterface, appInstSummaryReportIsSet := d.GetOk("app_inst_summary_report")
	if appInstSummaryReportIsSet && appInstSummaryReportInterface != nil {
		appInstSummaryReportMap := appInstSummaryReportInterface.([]interface{})
		if len(appInstSummaryReportMap) > 0 {
			appInstSummaryReport = AppInstSummaryReportModelFromMap(appInstSummaryReportMap[0].(map[string]interface{}))
		}
	}
	entpID, _ := d.Get("entp_id").(string)
	error, _ := d.Get("error").(string)
	return &models.AppInstReport{
		AppInstSummaryReport: appInstSummaryReport,
		EntpID:               entpID,
		Error:                error,
	}
}

func AppInstReportModelFromMap(m map[string]interface{}) *models.AppInstReport {
	var appInstSummaryReport *models.AppInstSummaryReport // AppInstSummaryReport
	appInstSummaryReportInterface, appInstSummaryReportIsSet := m["app_inst_summary_report"]
	if appInstSummaryReportIsSet && appInstSummaryReportInterface != nil {
		appInstSummaryReportMap := appInstSummaryReportInterface.([]interface{})
		if len(appInstSummaryReportMap) > 0 {
			appInstSummaryReport = AppInstSummaryReportModelFromMap(appInstSummaryReportMap[0].(map[string]interface{}))
		}
	}
	//
	entpID := m["entp_id"].(string)
	error := m["error"].(string)
	return &models.AppInstReport{
		AppInstSummaryReport: appInstSummaryReport,
		EntpID:               entpID,
		Error:                error,
	}
}

func SetAppInstReportResourceData(d *schema.ResourceData, m *models.AppInstReport) {
	d.Set("app_inst_summary_report", SetAppInstSummaryReportSubResourceData([]*models.AppInstSummaryReport{m.AppInstSummaryReport}))
	d.Set("entp_id", m.EntpID)
	d.Set("error", m.Error)
}

func SetAppInstReportSubResourceData(m []*models.AppInstReport) (d []*map[string]interface{}) {
	for _, AppInstReportModel := range m {
		if AppInstReportModel != nil {
			properties := make(map[string]interface{})
			properties["app_inst_summary_report"] = SetAppInstSummaryReportSubResourceData([]*models.AppInstSummaryReport{AppInstReportModel.AppInstSummaryReport})
			properties["entp_id"] = AppInstReportModel.EntpID
			properties["error"] = AppInstReportModel.Error
			d = append(d, &properties)
		}
	}
	return
}

func AppInstReportSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"app_inst_summary_report": {
			Description: `Enterprise appInst report`,
			Type:        schema.TypeList, //GoType: AppInstSummaryReport
			Elem: &schema.Resource{
				Schema: AppInstSummaryReportSchema(),
			},
			Optional: true,
		},

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
	}
}

func GetAppInstReportPropertyFields() (t []string) {
	return []string{
		"app_inst_summary_report",
		"entp_id",
		"error",
	}
}
