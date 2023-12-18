package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func PluginSummaryReportModel(d *schema.ResourceData) *models.PluginSummaryReport {
	totalPluginsInt, _ := d.Get("total_plugins").(int)
	totalPlugins := int64(totalPluginsInt)
	return &models.PluginSummaryReport{
		TotalPlugins: totalPlugins,
	}
}

func PluginSummaryReportModelFromMap(m map[string]interface{}) *models.PluginSummaryReport {
	totalPlugins := int64(m["total_plugins"].(int)) // int64
	return &models.PluginSummaryReport{
		TotalPlugins: totalPlugins,
	}
}

func SetPluginSummaryReportResourceData(d *schema.ResourceData, m *models.PluginSummaryReport) {
	d.Set("total_plugins", m.TotalPlugins)
}

func SetPluginSummaryReportSubResourceData(m []*models.PluginSummaryReport) (d []*map[string]interface{}) {
	for _, PluginSummaryReportModel := range m {
		if PluginSummaryReportModel != nil {
			properties := make(map[string]interface{})
			properties["total_plugins"] = PluginSummaryReportModel.TotalPlugins
			d = append(d, &properties)
		}
	}
	return
}

func PluginSummaryReportSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"total_plugins": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},
	}
}

func GetPluginSummaryReportPropertyFields() (t []string) {
	return []string{
		"total_plugins",
	}
}
