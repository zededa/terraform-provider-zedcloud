package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func DeviceEntitlementModel(d *schema.ResourceData) *models.DeviceEntitlement {
	limitInt, _ := d.Get("limit").(int)
	limit := int32(limitInt)
	var size *models.DeviceSize // DeviceSize
	sizeInterface, sizeIsSet := d.GetOk("size")
	if sizeIsSet {
		sizeModel := sizeInterface.(string)
		size = models.NewDeviceSize(models.DeviceSize(sizeModel))
	}
	usageInt, _ := d.Get("usage").(int)
	usage := int64(usageInt)
	return &models.DeviceEntitlement{
		Limit: limit,
		Size:  size,
		Usage: usage,
	}
}

func DeviceEntitlementModelFromMap(m map[string]interface{}) *models.DeviceEntitlement {
	limit := int32(m["limit"].(int)) // int32
	var size *models.DeviceSize      // DeviceSize
	sizeInterface, sizeIsSet := m["size"]
	if sizeIsSet {
		sizeModel := sizeInterface.(string)
		size = models.NewDeviceSize(models.DeviceSize(sizeModel))
	}
	usage := int64(m["usage"].(int)) // int64
	return &models.DeviceEntitlement{
		Limit: limit,
		Size:  size,
		Usage: usage,
	}
}

func SetDeviceEntitlementResourceData(d *schema.ResourceData, m *models.DeviceEntitlement) {
	d.Set("limit", m.Limit)
	d.Set("size", m.Size)
	d.Set("usage", m.Usage)
}

func SetDeviceEntitlementSubResourceData(m []*models.DeviceEntitlement) (d []*map[string]interface{}) {
	for _, DeviceEntitlementModel := range m {
		if DeviceEntitlementModel != nil {
			properties := make(map[string]interface{})
			properties["limit"] = DeviceEntitlementModel.Limit
			properties["size"] = DeviceEntitlementModel.Size
			properties["usage"] = DeviceEntitlementModel.Usage
			d = append(d, &properties)
		}
	}
	return
}

func DeviceEntitlementSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"limit": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"size": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"usage": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},
	}
}

func GetDeviceEntitlementPropertyFields() (t []string) {
	return []string{
		"limit",
		"size",
		"usage",
	}
}
