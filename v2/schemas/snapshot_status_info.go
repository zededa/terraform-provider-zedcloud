package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func SnapshotStatusInfoModel(d *schema.ResourceData) *models.SnapshotStatusInfo {
	message, _ := d.Get("message").(string)
	readyToUse, _ := d.Get("ready_to_use").(bool)
	state, _ := d.Get("state").(string)
	return &models.SnapshotStatusInfo{
		Message:    message,
		ReadyToUse: readyToUse,
		State:      state,
	}
}

func SnapshotStatusInfoModelFromMap(m map[string]interface{}) *models.SnapshotStatusInfo {
	message := m["message"].(string)
	readyToUse := m["ready_to_use"].(bool)
	state := m["state"].(string)
	return &models.SnapshotStatusInfo{
		Message:    message,
		ReadyToUse: readyToUse,
		State:      state,
	}
}

func SetSnapshotStatusInfoResourceData(d *schema.ResourceData, m *models.SnapshotStatusInfo) {
	d.Set("message", m.Message)
	d.Set("ready_to_use", m.ReadyToUse)
	d.Set("state", m.State)
}

func SetSnapshotStatusInfoSubResourceData(m []*models.SnapshotStatusInfo) (d []*map[string]interface{}) {
	for _, SnapshotStatusInfoModel := range m {
		if SnapshotStatusInfoModel != nil {
			properties := make(map[string]interface{})
			properties["message"] = SnapshotStatusInfoModel.Message
			properties["ready_to_use"] = SnapshotStatusInfoModel.ReadyToUse
			properties["state"] = SnapshotStatusInfoModel.State
			d = append(d, &properties)
		}
	}
	return
}

func SnapshotStatusInfoSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"message": {
			Description: `Status message with additional details`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"ready_to_use": {
			Description: `Whether the snapshot is ready to use for restore`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"state": {
			Description: `Current state of the snapshot (e.g., ready, creating, error)`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetSnapshotStatusInfoPropertyFields() (t []string) {
	return []string{
		"message",
		"ready_to_use",
		"state",
	}
}
