package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func SnapshotOperationStatusModel(d *schema.ResourceData) *models.SnapshotOperationStatus {
	snapshotOperationStatus, _ := d.Get("snapshot_operation_status").(models.SnapshotOperationStatus)
	return &snapshotOperationStatus
}

func SnapshotOperationStatusModelFromMap(m map[string]interface{}) *models.SnapshotOperationStatus {
	snapshotOperationStatus := m["snapshot_operation_status"].(models.SnapshotOperationStatus)
	return &snapshotOperationStatus
}

func SetSnapshotOperationStatusResourceData(d *schema.ResourceData, m *models.SnapshotOperationStatus) {
}

func SetSnapshotOperationStatusSubResourceData(m []*models.SnapshotOperationStatus) (d []*map[string]interface{}) {
	for _, SnapshotOperationStatusModel := range m {
		if SnapshotOperationStatusModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func SnapshotOperationStatusSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func GetSnapshotOperationStatusPropertyFields() (t []string) {
	return []string{}
}
