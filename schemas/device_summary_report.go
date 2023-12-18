package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func DeviceSummaryReportModel(d *schema.ResourceData) *models.DeviceSummaryReport {
	var summaryByState *models.Summary // Summary
	summaryByStateInterface, summaryByStateIsSet := d.GetOk("summary_by_state")
	if summaryByStateIsSet && summaryByStateInterface != nil {
		summaryByStateMap := summaryByStateInterface.([]interface{})
		if len(summaryByStateMap) > 0 {
			summaryByState = SummaryModelFromMap(summaryByStateMap[0].(map[string]interface{}))
		}
	}
	totalDevicesInt, _ := d.Get("total_devices").(int)
	totalDevices := int64(totalDevicesInt)
	return &models.DeviceSummaryReport{
		SummaryByState: summaryByState,
		TotalDevices:   totalDevices,
	}
}

func DeviceSummaryReportModelFromMap(m map[string]interface{}) *models.DeviceSummaryReport {
	var summaryByState *models.Summary // Summary
	summaryByStateInterface, summaryByStateIsSet := m["summary_by_state"]
	if summaryByStateIsSet && summaryByStateInterface != nil {
		summaryByStateMap := summaryByStateInterface.([]interface{})
		if len(summaryByStateMap) > 0 {
			summaryByState = SummaryModelFromMap(summaryByStateMap[0].(map[string]interface{}))
		}
	}
	//
	totalDevices := int64(m["total_devices"].(int)) // int64
	return &models.DeviceSummaryReport{
		SummaryByState: summaryByState,
		TotalDevices:   totalDevices,
	}
}

func SetDeviceSummaryReportResourceData(d *schema.ResourceData, m *models.DeviceSummaryReport) {
	d.Set("summary_by_state", SetSummarySubResourceData([]*models.Summary{m.SummaryByState}))
	d.Set("total_devices", m.TotalDevices)
}

func SetDeviceSummaryReportSubResourceData(m []*models.DeviceSummaryReport) (d []*map[string]interface{}) {
	for _, DeviceSummaryReportModel := range m {
		if DeviceSummaryReportModel != nil {
			properties := make(map[string]interface{})
			properties["summary_by_state"] = SetSummarySubResourceData([]*models.Summary{DeviceSummaryReportModel.SummaryByState})
			properties["total_devices"] = DeviceSummaryReportModel.TotalDevices
			d = append(d, &properties)
		}
	}
	return
}

func DeviceSummaryReportSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"summary_by_state": {
			Description: ``,
			Type:        schema.TypeList, //GoType: Summary
			Elem: &schema.Resource{
				Schema: SummarySchema(),
			},
			Optional: true,
		},

		"total_devices": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},
	}
}

func GetDeviceSummaryReportPropertyFields() (t []string) {
	return []string{
		"summary_by_state",
		"total_devices",
	}
}
