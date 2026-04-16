package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func SnapshotStateModel(d *schema.ResourceData) *models.SnapshotState {
	snapshotState, _ := d.Get("snapshot_state").(models.SnapshotState)
	return &snapshotState
}

func SnapshotStateModelFromMap(m map[string]interface{}) *models.SnapshotState {
	snapshotState := m["snapshot_state"].(models.SnapshotState)
	return &snapshotState
}

func SetSnapshotStateResourceData(d *schema.ResourceData, m *models.SnapshotState) {
}

func SetSnapshotStateSubResourceData(m []*models.SnapshotState) (d []*map[string]interface{}) {
	for _, SnapshotStateModel := range m {
		if SnapshotStateModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func SnapshotStateSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func GetSnapshotStatePropertyFields() (t []string) {
	return []string{}
}
