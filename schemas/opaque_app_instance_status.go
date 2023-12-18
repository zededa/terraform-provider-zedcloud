package schemas

import (
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/zededa/terraform-provider/models"
)

func OpaqueAppInstanceStatusModel(d *schema.ResourceData) *models.OpaqueAppInstanceStatus {
	deviceID, _ := d.Get("device_id").(string)
	id, _ := d.Get("id").(string)
	lastUpdateTime, _ := d.Get("last_update_time").(strfmt.DateTime)
	name, _ := d.Get("name").(string)
	opaqueAppInstanceStatus, _ := d.Get("opaque_app_instance_status").(strfmt.Base64)
	return &models.OpaqueAppInstanceStatus{
		DeviceID:                deviceID,
		ID:                      id,
		LastUpdateTime:          lastUpdateTime,
		Name:                    name,
		OpaqueAppInstanceStatus: opaqueAppInstanceStatus,
	}
}

func OpaqueAppInstanceStatusModelFromMap(m map[string]interface{}) *models.OpaqueAppInstanceStatus {
	deviceID := m["device_id"].(string)
	id := m["id"].(string)
	lastUpdateTime := m["last_update_time"].(strfmt.DateTime)
	name := m["name"].(string)
	opaqueAppInstanceStatus := m["opaque_app_instance_status"].(strfmt.Base64)
	return &models.OpaqueAppInstanceStatus{
		DeviceID:                deviceID,
		ID:                      id,
		LastUpdateTime:          lastUpdateTime,
		Name:                    name,
		OpaqueAppInstanceStatus: opaqueAppInstanceStatus,
	}
}

func SetOpaqueAppInstanceStatusResourceData(d *schema.ResourceData, m *models.OpaqueAppInstanceStatus) {
	d.Set("device_id", m.DeviceID)
	d.Set("id", m.ID)
	d.Set("last_update_time", m.LastUpdateTime)
	d.Set("name", m.Name)
	d.Set("opaque_app_instance_status", m.OpaqueAppInstanceStatus.String())
}

func SetOpaqueAppInstanceStatusSubResourceData(m []*models.OpaqueAppInstanceStatus) (d []*map[string]interface{}) {
	for _, OpaqueAppInstanceStatusModel := range m {
		if OpaqueAppInstanceStatusModel != nil {
			properties := make(map[string]interface{})
			properties["device_id"] = OpaqueAppInstanceStatusModel.DeviceID
			properties["id"] = OpaqueAppInstanceStatusModel.ID
			properties["last_update_time"] = OpaqueAppInstanceStatusModel.LastUpdateTime.String()
			properties["name"] = OpaqueAppInstanceStatusModel.Name
			properties["opaque_app_instance_status"] = OpaqueAppInstanceStatusModel.OpaqueAppInstanceStatus.String()
			d = append(d, &properties)
		}
	}
	return
}

func OpaqueAppInstanceStatusSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"device_id": {
			Description: `User defined name of the device, unique across the enterprise.`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"id": {
			Description: `System defined universally unique Id of the app instance`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"last_update_time": {
			Description:  `last update time`,
			Type:         schema.TypeString,
			ValidateFunc: validation.IsRFC3339Time,
			Optional:     true,
		},

		"name": {
			Description: `App instance name. Unique across the application instance`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"opaque_app_instance_status": {
			Description: `opaque app instance status info`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetOpaqueAppInstanceStatusPropertyFields() (t []string) {
	return []string{
		"device_id",
		"id",
		"last_update_time",
		"name",
		"opaque_app_instance_status",
	}
}
