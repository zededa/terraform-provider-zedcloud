package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func SnapshotOperationTypeModel(d *schema.ResourceData) *models.SnapshotOperationType {
	snapshotOperationType, _ := d.Get("snapshot_operation_type").(models.SnapshotOperationType)
	return &snapshotOperationType
}

func SnapshotOperationTypeModelFromMap(m map[string]interface{}) *models.SnapshotOperationType {
	snapshotOperationType := m["snapshot_operation_type"].(models.SnapshotOperationType)
	return &snapshotOperationType
}

func SetSnapshotOperationTypeResourceData(d *schema.ResourceData, m *models.SnapshotOperationType) {
}

func SetSnapshotOperationTypeSubResourceData(m []*models.SnapshotOperationType) (d []*map[string]interface{}) {
	for _, SnapshotOperationTypeModel := range m {
		if SnapshotOperationTypeModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func SnapshotOperationTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func GetSnapshotOperationTypePropertyFields() (t []string) {
	return []string{}
}
