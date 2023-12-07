package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func PatchStatusFilterModel(d *schema.ResourceData) *models.PatchStatusFilter {
	deviceName, _ := d.Get("device_name").(string)
	return &models.PatchStatusFilter{
		DeviceName: deviceName,
	}
}

func PatchStatusFilterModelFromMap(m map[string]interface{}) *models.PatchStatusFilter {
	deviceName := m["device_name"].(string)
	return &models.PatchStatusFilter{
		DeviceName: deviceName,
	}
}

func SetPatchStatusFilterResourceData(d *schema.ResourceData, m *models.PatchStatusFilter) {
	d.Set("device_name", m.DeviceName)
}

func SetPatchStatusFilterSubResourceData(m []*models.PatchStatusFilter) (d []*map[string]interface{}) {
	for _, PatchStatusFilterModel := range m {
		if PatchStatusFilterModel != nil {
			properties := make(map[string]interface{})
			properties["device_name"] = PatchStatusFilterModel.DeviceName
			d = append(d, &properties)
		}
	}
	return
}

func PatchStatusFilterSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"device_name": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetPatchStatusFilterPropertyFields() (t []string) {
	return []string{
		"device_name",
	}
}
