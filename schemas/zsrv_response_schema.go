package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate ZsrvResponse resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ZsrvResponseModel(d *schema.ResourceData) *models.ZsrvResponse {
	endTime := d.Get("end_time").(string)
	error := d.Get("error").([]*models.ZsrvError) // []*ZsrvError
	hTTPStatusCode := int32(d.Get("http_status_code").(int))
	hTTPStatusMsg := d.Get("http_status_msg").(string)
	objectID := d.Get("object_id").(string)
	objectKind := d.Get("object_kind").(string)
	objectName := d.Get("object_name").(string)
	objectRevision := d.Get("object_revision").(string)
	objectType := d.Get("object_type").(*models.ObjectType)            // ObjectType
	operationStatus := d.Get("operation_status").(*models.ZcOpsStatus) // ZcOpsStatus
	operationTime := d.Get("operation_time").(string)
	operationType := d.Get("operation_type").(*models.ZcOpsType) // ZcOpsType
	startTime := d.Get("start_time").(string)
	user := d.Get("user").(string)
	return &models.ZsrvResponse{
		EndTime:         endTime,
		Error:           error,
		HTTPStatusCode:  hTTPStatusCode,
		HTTPStatusMsg:   hTTPStatusMsg,
		ObjectID:        objectID,
		ObjectKind:      objectKind,
		ObjectName:      objectName,
		ObjectRevision:  objectRevision,
		ObjectType:      objectType,
		OperationStatus: operationStatus,
		OperationTime:   operationTime,
		OperationType:   operationType,
		StartTime:       startTime,
		User:            user,
	}
}

func ZsrvResponseModelFromMap(m map[string]interface{}) *models.ZsrvResponse {
	endTime := m["end_time"].(string)
	error := m["error"].([]*models.ZsrvError)            // []*ZsrvError
	hTTPStatusCode := int32(m["http_status_code"].(int)) // int32 false false false
	hTTPStatusMsg := m["http_status_msg"].(string)
	objectID := m["object_id"].(string)
	objectKind := m["object_kind"].(string)
	objectName := m["object_name"].(string)
	objectRevision := m["object_revision"].(string)
	objectType := m["object_type"].(*models.ObjectType)            // ObjectType
	operationStatus := m["operation_status"].(*models.ZcOpsStatus) // ZcOpsStatus
	operationTime := m["operation_time"].(string)
	operationType := m["operation_type"].(*models.ZcOpsType) // ZcOpsType
	startTime := m["start_time"].(string)
	user := m["user"].(string)
	return &models.ZsrvResponse{
		EndTime:         endTime,
		Error:           error,
		HTTPStatusCode:  hTTPStatusCode,
		HTTPStatusMsg:   hTTPStatusMsg,
		ObjectID:        objectID,
		ObjectKind:      objectKind,
		ObjectName:      objectName,
		ObjectRevision:  objectRevision,
		ObjectType:      objectType,
		OperationStatus: operationStatus,
		OperationTime:   operationTime,
		OperationType:   operationType,
		StartTime:       startTime,
		User:            user,
	}
}

// Update the underlying ZsrvResponse resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetZsrvResponseResourceData(d *schema.ResourceData, m *models.ZsrvResponse) {
	d.Set("end_time", m.EndTime)
	d.Set("error", SetZsrvErrorSubResourceData(m.Error))
	d.Set("http_status_code", m.HTTPStatusCode)
	d.Set("http_status_msg", m.HTTPStatusMsg)
	d.Set("object_id", m.ObjectID)
	d.Set("object_kind", m.ObjectKind)
	d.Set("object_name", m.ObjectName)
	d.Set("object_revision", m.ObjectRevision)
	d.Set("object_type", m.ObjectType)
	d.Set("operation_status", m.OperationStatus)
	d.Set("operation_time", m.OperationTime)
	d.Set("operation_type", m.OperationType)
	d.Set("start_time", m.StartTime)
	d.Set("user", m.User)
}

// Iterate throught and update the ZsrvResponse resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetZsrvResponseSubResourceData(m []*models.ZsrvResponse) (d []*map[string]interface{}) {
	for _, ZsrvResponseModel := range m {
		if ZsrvResponseModel != nil {
			properties := make(map[string]interface{})
			properties["end_time"] = ZsrvResponseModel.EndTime
			properties["error"] = SetZsrvErrorSubResourceData(ZsrvResponseModel.Error)
			properties["http_status_code"] = ZsrvResponseModel.HTTPStatusCode
			properties["http_status_msg"] = ZsrvResponseModel.HTTPStatusMsg
			properties["object_id"] = ZsrvResponseModel.ObjectID
			properties["object_kind"] = ZsrvResponseModel.ObjectKind
			properties["object_name"] = ZsrvResponseModel.ObjectName
			properties["object_revision"] = ZsrvResponseModel.ObjectRevision
			properties["object_type"] = ZsrvResponseModel.ObjectType
			properties["operation_status"] = ZsrvResponseModel.OperationStatus
			properties["operation_time"] = ZsrvResponseModel.OperationTime
			properties["operation_type"] = ZsrvResponseModel.OperationType
			properties["start_time"] = ZsrvResponseModel.StartTime
			properties["user"] = ZsrvResponseModel.User
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the ZsrvResponse resource defined in the Terraform configuration
func ZsrvResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"end_time": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"error": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []*ZsrvError
			Elem: &schema.Resource{
				Schema: ZsrvErrorSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"http_status_code": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"http_status_msg": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"object_id": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"object_kind": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"object_name": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"object_revision": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"object_type": {
			Description: ``,
			Type:        schema.TypeList, //GoType: ObjectType
			Elem: &schema.Resource{
				Schema: ObjectTypeSchema(),
			},
			Optional: true,
		},

		"operation_status": {
			Description: ``,
			Type:        schema.TypeList, //GoType: ZcOpsStatus
			Elem: &schema.Resource{
				Schema: ZcOpsStatusSchema(),
			},
			Optional: true,
		},

		"operation_time": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"operation_type": {
			Description: ``,
			Type:        schema.TypeList, //GoType: ZcOpsType
			Elem: &schema.Resource{
				Schema: ZcOpsTypeSchema(),
			},
			Optional: true,
		},

		"start_time": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"user": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the ZsrvResponse resource
func GetZsrvResponsePropertyFields() (t []string) {
	return []string{
		"end_time",
		"error",
		"http_status_code",
		"http_status_msg",
		"object_id",
		"object_kind",
		"object_name",
		"object_revision",
		"object_type",
		"operation_status",
		"operation_time",
		"operation_type",
		"start_time",
		"user",
	}
}
