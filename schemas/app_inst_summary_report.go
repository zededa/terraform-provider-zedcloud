package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AppInstSummaryReportModel(d *schema.ResourceData) *models.AppInstSummaryReport {
	totalAppInstsInt, _ := d.Get("total_app_insts").(int)
	totalAppInsts := int64(totalAppInstsInt)
	return &models.AppInstSummaryReport{
		TotalAppInsts: totalAppInsts,
	}
}

func AppInstSummaryReportModelFromMap(m map[string]interface{}) *models.AppInstSummaryReport {
	totalAppInsts := int64(m["total_app_insts"].(int)) // int64
	return &models.AppInstSummaryReport{
		TotalAppInsts: totalAppInsts,
	}
}

func SetAppInstSummaryReportResourceData(d *schema.ResourceData, m *models.AppInstSummaryReport) {
	d.Set("total_app_insts", m.TotalAppInsts)
}

func SetAppInstSummaryReportSubResourceData(m []*models.AppInstSummaryReport) (d []*map[string]interface{}) {
	for _, AppInstSummaryReportModel := range m {
		if AppInstSummaryReportModel != nil {
			properties := make(map[string]interface{})
			properties["total_app_insts"] = AppInstSummaryReportModel.TotalAppInsts
			d = append(d, &properties)
		}
	}
	return
}

func AppInstSummaryReportSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"total_app_insts": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},
	}
}

func GetAppInstSummaryReportPropertyFields() (t []string) {
	return []string{
		"total_app_insts",
	}
}
