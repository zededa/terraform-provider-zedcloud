package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func PluginReportModel(d *schema.ResourceData) *models.PluginReport {
	entpID, _ := d.Get("entp_id").(string)
	error, _ := d.Get("error").(string)
	var pluginSummaryReport *models.PluginSummaryReport // PluginSummaryReport
	pluginSummaryReportInterface, pluginSummaryReportIsSet := d.GetOk("plugin_summary_report")
	if pluginSummaryReportIsSet && pluginSummaryReportInterface != nil {
		pluginSummaryReportMap := pluginSummaryReportInterface.([]interface{})
		if len(pluginSummaryReportMap) > 0 {
			pluginSummaryReport = PluginSummaryReportModelFromMap(pluginSummaryReportMap[0].(map[string]interface{}))
		}
	}
	return &models.PluginReport{
		EntpID:              entpID,
		Error:               error,
		PluginSummaryReport: pluginSummaryReport,
	}
}

func PluginReportModelFromMap(m map[string]interface{}) *models.PluginReport {
	entpID := m["entp_id"].(string)
	error := m["error"].(string)
	var pluginSummaryReport *models.PluginSummaryReport // PluginSummaryReport
	pluginSummaryReportInterface, pluginSummaryReportIsSet := m["plugin_summary_report"]
	if pluginSummaryReportIsSet && pluginSummaryReportInterface != nil {
		pluginSummaryReportMap := pluginSummaryReportInterface.([]interface{})
		if len(pluginSummaryReportMap) > 0 {
			pluginSummaryReport = PluginSummaryReportModelFromMap(pluginSummaryReportMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.PluginReport{
		EntpID:              entpID,
		Error:               error,
		PluginSummaryReport: pluginSummaryReport,
	}
}

func SetPluginReportResourceData(d *schema.ResourceData, m *models.PluginReport) {
	d.Set("entp_id", m.EntpID)
	d.Set("error", m.Error)
	d.Set("plugin_summary_report", SetPluginSummaryReportSubResourceData([]*models.PluginSummaryReport{m.PluginSummaryReport}))
}

func SetPluginReportSubResourceData(m []*models.PluginReport) (d []*map[string]interface{}) {
	for _, PluginReportModel := range m {
		if PluginReportModel != nil {
			properties := make(map[string]interface{})
			properties["entp_id"] = PluginReportModel.EntpID
			properties["error"] = PluginReportModel.Error
			properties["plugin_summary_report"] = SetPluginSummaryReportSubResourceData([]*models.PluginSummaryReport{PluginReportModel.PluginSummaryReport})
			d = append(d, &properties)
		}
	}
	return
}

func PluginReportSchema() map[string]*schema.Schema {
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

		"plugin_summary_report": {
			Description: `Enterprise plugin report`,
			Type:        schema.TypeList, //GoType: PluginSummaryReport
			Elem: &schema.Resource{
				Schema: PluginSummaryReportSchema(),
			},
			Optional: true,
		},
	}
}

func GetPluginReportPropertyFields() (t []string) {
	return []string{
		"entp_id",
		"error",
		"plugin_summary_report",
	}
}
