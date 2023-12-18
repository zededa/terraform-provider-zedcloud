package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func DeviceDistributionModel(d *schema.ResourceData) *models.DeviceDistribution {
	deviceCountInt, _ := d.Get("device_count").(int)
	deviceCount := int64(deviceCountInt)
	projectName, _ := d.Get("project_name").(string)
	return &models.DeviceDistribution{
		DeviceCount: deviceCount,
		ProjectName: projectName,
	}
}

func DeviceDistributionModelFromMap(m map[string]interface{}) *models.DeviceDistribution {
	deviceCount := int64(m["device_count"].(int)) // int64
	projectName := m["project_name"].(string)
	return &models.DeviceDistribution{
		DeviceCount: deviceCount,
		ProjectName: projectName,
	}
}

func SetDeviceDistributionResourceData(d *schema.ResourceData, m *models.DeviceDistribution) {
	d.Set("device_count", m.DeviceCount)
	d.Set("project_name", m.ProjectName)
}

func SetDeviceDistributionSubResourceData(m []*models.DeviceDistribution) (d []*map[string]interface{}) {
	for _, DeviceDistributionModel := range m {
		if DeviceDistributionModel != nil {
			properties := make(map[string]interface{})
			properties["device_count"] = DeviceDistributionModel.DeviceCount
			properties["project_name"] = DeviceDistributionModel.ProjectName
			d = append(d, &properties)
		}
	}
	return
}

func DeviceDistributionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"device_count": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"project_name": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetDeviceDistributionPropertyFields() (t []string) {
	return []string{
		"device_count",
		"project_name",
	}
}
