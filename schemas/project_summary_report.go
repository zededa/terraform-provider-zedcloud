package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func ProjectSummaryReportModel(d *schema.ResourceData) *models.ProjectSummaryReport {
	var deviceDistribution []*models.DeviceDistribution // []*DeviceDistribution
	deviceDistributionInterface, deviceDistributionIsSet := d.GetOk("device_distribution")
	if deviceDistributionIsSet {
		var items []interface{}
		if listItems, isList := deviceDistributionInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = deviceDistributionInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := DeviceDistributionModelFromMap(v.(map[string]interface{}))
			deviceDistribution = append(deviceDistribution, m)
		}
	}
	totalInt, _ := d.Get("total").(int)
	total := int64(totalInt)
	return &models.ProjectSummaryReport{
		DeviceDistribution: deviceDistribution,
		Total:              total,
	}
}

func ProjectSummaryReportModelFromMap(m map[string]interface{}) *models.ProjectSummaryReport {
	var deviceDistribution []*models.DeviceDistribution // []*DeviceDistribution
	deviceDistributionInterface, deviceDistributionIsSet := m["device_distribution"]
	if deviceDistributionIsSet {
		var items []interface{}
		if listItems, isList := deviceDistributionInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = deviceDistributionInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := DeviceDistributionModelFromMap(v.(map[string]interface{}))
			deviceDistribution = append(deviceDistribution, m)
		}
	}
	total := int64(m["total"].(int)) // int64
	return &models.ProjectSummaryReport{
		DeviceDistribution: deviceDistribution,
		Total:              total,
	}
}

func SetProjectSummaryReportResourceData(d *schema.ResourceData, m *models.ProjectSummaryReport) {
	d.Set("device_distribution", SetDeviceDistributionSubResourceData(m.DeviceDistribution))
	d.Set("total", m.Total)
}

func SetProjectSummaryReportSubResourceData(m []*models.ProjectSummaryReport) (d []*map[string]interface{}) {
	for _, ProjectSummaryReportModel := range m {
		if ProjectSummaryReportModel != nil {
			properties := make(map[string]interface{})
			properties["device_distribution"] = SetDeviceDistributionSubResourceData(ProjectSummaryReportModel.DeviceDistribution)
			properties["total"] = ProjectSummaryReportModel.Total
			d = append(d, &properties)
		}
	}
	return
}

func ProjectSummaryReportSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"device_distribution": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []*DeviceDistribution
			Elem: &schema.Resource{
				Schema: DeviceDistributionSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"total": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},
	}
}

func GetProjectSummaryReportPropertyFields() (t []string) {
	return []string{
		"device_distribution",
		"total",
	}
}
