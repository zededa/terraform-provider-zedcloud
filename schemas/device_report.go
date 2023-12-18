package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func DeviceReportModel(d *schema.ResourceData) *models.DeviceReport {
	var deviceSummaryReport *models.DeviceSummaryReport // DeviceSummaryReport
	deviceSummaryReportInterface, deviceSummaryReportIsSet := d.GetOk("device_summary_report")
	if deviceSummaryReportIsSet && deviceSummaryReportInterface != nil {
		deviceSummaryReportMap := deviceSummaryReportInterface.([]interface{})
		if len(deviceSummaryReportMap) > 0 {
			deviceSummaryReport = DeviceSummaryReportModelFromMap(deviceSummaryReportMap[0].(map[string]interface{}))
		}
	}
	entpID, _ := d.Get("entp_id").(string)
	error, _ := d.Get("error").(string)
	return &models.DeviceReport{
		DeviceSummaryReport: deviceSummaryReport,
		EntpID:              entpID,
		Error:               error,
	}
}

func DeviceReportModelFromMap(m map[string]interface{}) *models.DeviceReport {
	var deviceSummaryReport *models.DeviceSummaryReport // DeviceSummaryReport
	deviceSummaryReportInterface, deviceSummaryReportIsSet := m["device_summary_report"]
	if deviceSummaryReportIsSet && deviceSummaryReportInterface != nil {
		deviceSummaryReportMap := deviceSummaryReportInterface.([]interface{})
		if len(deviceSummaryReportMap) > 0 {
			deviceSummaryReport = DeviceSummaryReportModelFromMap(deviceSummaryReportMap[0].(map[string]interface{}))
		}
	}
	//
	entpID := m["entp_id"].(string)
	error := m["error"].(string)
	return &models.DeviceReport{
		DeviceSummaryReport: deviceSummaryReport,
		EntpID:              entpID,
		Error:               error,
	}
}

func SetDeviceReportResourceData(d *schema.ResourceData, m *models.DeviceReport) {
	d.Set("device_summary_report", SetDeviceSummaryReportSubResourceData([]*models.DeviceSummaryReport{m.DeviceSummaryReport}))
	d.Set("entp_id", m.EntpID)
	d.Set("error", m.Error)
}

func SetDeviceReportSubResourceData(m []*models.DeviceReport) (d []*map[string]interface{}) {
	for _, DeviceReportModel := range m {
		if DeviceReportModel != nil {
			properties := make(map[string]interface{})
			properties["device_summary_report"] = SetDeviceSummaryReportSubResourceData([]*models.DeviceSummaryReport{DeviceReportModel.DeviceSummaryReport})
			properties["entp_id"] = DeviceReportModel.EntpID
			properties["error"] = DeviceReportModel.Error
			d = append(d, &properties)
		}
	}
	return
}

func DeviceReportSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"device_summary_report": {
			Description: `Enterprise device report`,
			Type:        schema.TypeList, //GoType: DeviceSummaryReport
			Elem: &schema.Resource{
				Schema: DeviceSummaryReportSchema(),
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

func GetDeviceReportPropertyFields() (t []string) {
	return []string{
		"device_summary_report",
		"entp_id",
		"error",
	}
}
