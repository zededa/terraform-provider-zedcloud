package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func DeviceSizeModel(d *schema.ResourceData) *models.DeviceSize {
	deviceSize, _ := d.Get("device_size").(models.DeviceSize)
	return &deviceSize
}

func DeviceSizeModelFromMap(m map[string]interface{}) *models.DeviceSize {
	deviceSize := m["device_size"].(models.DeviceSize)
	return &deviceSize
}

func SetDeviceSizeResourceData(d *schema.ResourceData, m *models.DeviceSize) {
}

func SetDeviceSizeSubResourceData(m []*models.DeviceSize) (d []*map[string]interface{}) {
	for _, DeviceSizeModel := range m {
		if DeviceSizeModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func DeviceSizeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func GetDeviceSizePropertyFields() (t []string) {
	return []string{}
}
