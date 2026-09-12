package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/go-openapi/strfmt"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func SnapshotOperationResponseModel(d *schema.ResourceData) *models.SnapshotOperationResponse {
	clusterID, _ := d.Get("cluster_id").(string)
	createdAt, _ := d.Get("created_at").(strfmt.DateTime)
	message, _ := d.Get("message").(string)
	operationID, _ := d.Get("operation_id").(string)
	var operationType *models.SnapshotOperationType // SnapshotOperationType
	operationTypeInterface, operationTypeIsSet := d.GetOk("operation_type")
	if operationTypeIsSet {
		operationTypeModel := operationTypeInterface.(string)
		operationType = models.NewSnapshotOperationType(models.SnapshotOperationType(operationTypeModel))
	}
	phase, _ := d.Get("phase").(string)
	snapshotID, _ := d.Get("snapshot_id").(string)
	snapshotName, _ := d.Get("snapshot_name").(string)
	var status *models.SnapshotOperationStatus // SnapshotOperationStatus
	statusInterface, statusIsSet := d.GetOk("status")
	if statusIsSet {
		statusModel := statusInterface.(string)
		status = models.NewSnapshotOperationStatus(models.SnapshotOperationStatus(statusModel))
	}
	updatedAt, _ := d.Get("updated_at").(strfmt.DateTime)
	return &models.SnapshotOperationResponse{
		ClusterID:     clusterID,
		CreatedAt:     createdAt,
		Message:       message,
		OperationID:   operationID,
		OperationType: operationType,
		Phase:         phase,
		SnapshotID:    snapshotID,
		SnapshotName:  snapshotName,
		Status:        status,
		UpdatedAt:     updatedAt,
	}
}

func SnapshotOperationResponseModelFromMap(m map[string]interface{}) *models.SnapshotOperationResponse {
	clusterID := m["cluster_id"].(string)
	createdAt := m["created_at"].(strfmt.DateTime)
	message := m["message"].(string)
	operationID := m["operation_id"].(string)
	var operationType *models.SnapshotOperationType // SnapshotOperationType
	operationTypeInterface, operationTypeIsSet := m["operation_type"]
	if operationTypeIsSet {
		operationTypeModel := operationTypeInterface.(string)
		operationType = models.NewSnapshotOperationType(models.SnapshotOperationType(operationTypeModel))
	}
	phase := m["phase"].(string)
	snapshotID := m["snapshot_id"].(string)
	snapshotName := m["snapshot_name"].(string)
	var status *models.SnapshotOperationStatus // SnapshotOperationStatus
	statusInterface, statusIsSet := m["status"]
	if statusIsSet {
		statusModel := statusInterface.(string)
		status = models.NewSnapshotOperationStatus(models.SnapshotOperationStatus(statusModel))
	}
	updatedAt := m["updated_at"].(strfmt.DateTime)
	return &models.SnapshotOperationResponse{
		ClusterID:     clusterID,
		CreatedAt:     createdAt,
		Message:       message,
		OperationID:   operationID,
		OperationType: operationType,
		Phase:         phase,
		SnapshotID:    snapshotID,
		SnapshotName:  snapshotName,
		Status:        status,
		UpdatedAt:     updatedAt,
	}
}

func SetSnapshotOperationResponseResourceData(d *schema.ResourceData, m *models.SnapshotOperationResponse) {
	d.Set("cluster_id", m.ClusterID)
	d.Set("created_at", m.CreatedAt)
	d.Set("message", m.Message)
	d.Set("operation_id", m.OperationID)
	d.Set("operation_type", m.OperationType)
	d.Set("phase", m.Phase)
	d.Set("snapshot_id", m.SnapshotID)
	d.Set("snapshot_name", m.SnapshotName)
	d.Set("status", m.Status)
	d.Set("updated_at", m.UpdatedAt)
}

func SetSnapshotOperationResponseSubResourceData(m []*models.SnapshotOperationResponse) (d []*map[string]interface{}) {
	for _, SnapshotOperationResponseModel := range m {
		if SnapshotOperationResponseModel != nil {
			properties := make(map[string]interface{})
			properties["cluster_id"] = SnapshotOperationResponseModel.ClusterID
			properties["created_at"] = SnapshotOperationResponseModel.CreatedAt.String()
			properties["message"] = SnapshotOperationResponseModel.Message
			properties["operation_id"] = SnapshotOperationResponseModel.OperationID
			properties["operation_type"] = SnapshotOperationResponseModel.OperationType
			properties["phase"] = SnapshotOperationResponseModel.Phase
			properties["snapshot_id"] = SnapshotOperationResponseModel.SnapshotID
			properties["snapshot_name"] = SnapshotOperationResponseModel.SnapshotName
			properties["status"] = SnapshotOperationResponseModel.Status
			properties["updated_at"] = SnapshotOperationResponseModel.UpdatedAt.String()
			d = append(d, &properties)
		}
	}
	return
}

func SnapshotOperationResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cluster_id": {
			Description: `Cluster ID associated with the operation`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"created_at": {
			Description:  `When the operation was created`,
			Type:         schema.TypeString,
			ValidateFunc: validation.IsRFC3339Time,
			Optional:     true,
		},

		"message": {
			Description: `Status message`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"operation_id": {
			Description: `Unique identifier for the async operation`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"operation_type": {
			Description: `Type of the operation`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"phase": {
			Description: `Current phase of the operation`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"snapshot_id": {
			Description: `Snapshot id generated for the operation`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"snapshot_name": {
			Description: `Snapshot Name`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"status": {
			Description: `Current status of the operation`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"updated_at": {
			Description:  `When the operation was last updated`,
			Type:         schema.TypeString,
			ValidateFunc: validation.IsRFC3339Time,
			Optional:     true,
		},
	}
}

func GetSnapshotOperationResponsePropertyFields() (t []string) {
	return []string{
		"cluster_id",
		"created_at",
		"message",
		"operation_id",
		"operation_type",
		"phase",
		"snapshot_id",
		"snapshot_name",
		"status",
		"updated_at",
	}
}
