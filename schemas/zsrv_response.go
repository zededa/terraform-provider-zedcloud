package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ZsrvResponseModel(d *schema.ResourceData) *models.ZsrvResponse {
	endTime, _ := d.Get("end_time").(string)
	var error []*models.ZsrvError // []*ZsrvError
	errorInterface, errorIsSet := d.GetOk("error")
	if errorIsSet {
		var items []interface{}
		if listItems, isList := errorInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = errorInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := ZsrvErrorModelFromMap(v.(map[string]interface{}))
			error = append(error, m)
		}
	}
	hTTPStatusCodeInt, _ := d.Get("http_status_code").(int)
	hTTPStatusCode := int32(hTTPStatusCodeInt)
	hTTPStatusMsg, _ := d.Get("http_status_msg").(string)
	jobID, _ := d.Get("job_id").(string)
	objectID, _ := d.Get("object_id").(string)
	objectKind, _ := d.Get("object_kind").(string)
	objectName, _ := d.Get("object_name").(string)
	objectRevision, _ := d.Get("object_revision").(string)
	var objectType *models.ObjectType // ObjectType
	objectTypeInterface, objectTypeIsSet := d.GetOk("object_type")
	if objectTypeIsSet {
		objectTypeModel := objectTypeInterface.(string)
		objectType = models.NewObjectType(models.ObjectType(objectTypeModel))
	}
	var operationStatus *models.ZcOpsStatus // ZcOpsStatus
	operationStatusInterface, operationStatusIsSet := d.GetOk("operation_status")
	if operationStatusIsSet {
		operationStatusModel := operationStatusInterface.(string)
		operationStatus = models.NewZcOpsStatus(models.ZcOpsStatus(operationStatusModel))
	}
	operationTime, _ := d.Get("operation_time").(string)
	var operationType *models.ZcOpsType // ZcOpsType
	operationTypeInterface, operationTypeIsSet := d.GetOk("operation_type")
	if operationTypeIsSet {
		operationTypeModel := operationTypeInterface.(string)
		operationType = models.NewZcOpsType(models.ZcOpsType(operationTypeModel))
	}
	startTime, _ := d.Get("start_time").(string)
	user, _ := d.Get("user").(string)
	return &models.ZsrvResponse{
		EndTime:         endTime,
		Error:           error,
		HTTPStatusCode:  hTTPStatusCode,
		HTTPStatusMsg:   hTTPStatusMsg,
		JobID:           jobID,
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
	var error []*models.ZsrvError // []*ZsrvError
	errorInterface, errorIsSet := m["error"]
	if errorIsSet {
		var items []interface{}
		if listItems, isList := errorInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = errorInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := ZsrvErrorModelFromMap(v.(map[string]interface{}))
			error = append(error, m)
		}
	}
	hTTPStatusCode := int32(m["http_status_code"].(int)) // int32
	hTTPStatusMsg := m["http_status_msg"].(string)
	jobID := m["job_id"].(string)
	objectID := m["object_id"].(string)
	objectKind := m["object_kind"].(string)
	objectName := m["object_name"].(string)
	objectRevision := m["object_revision"].(string)
	var objectType *models.ObjectType // ObjectType
	objectTypeInterface, objectTypeIsSet := m["object_type"]
	if objectTypeIsSet {
		objectTypeModel := objectTypeInterface.(string)
		objectType = models.NewObjectType(models.ObjectType(objectTypeModel))
	}
	var operationStatus *models.ZcOpsStatus // ZcOpsStatus
	operationStatusInterface, operationStatusIsSet := m["operation_status"]
	if operationStatusIsSet {
		operationStatusModel := operationStatusInterface.(string)
		operationStatus = models.NewZcOpsStatus(models.ZcOpsStatus(operationStatusModel))
	}
	operationTime := m["operation_time"].(string)
	var operationType *models.ZcOpsType // ZcOpsType
	operationTypeInterface, operationTypeIsSet := m["operation_type"]
	if operationTypeIsSet {
		operationTypeModel := operationTypeInterface.(string)
		operationType = models.NewZcOpsType(models.ZcOpsType(operationTypeModel))
	}
	startTime := m["start_time"].(string)
	user := m["user"].(string)
	return &models.ZsrvResponse{
		EndTime:         endTime,
		Error:           error,
		HTTPStatusCode:  hTTPStatusCode,
		HTTPStatusMsg:   hTTPStatusMsg,
		JobID:           jobID,
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
	d.Set("job_id", m.JobID)
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

// Iterate through and update the ZsrvResponse resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetZsrvResponseSubResourceData(m []*models.ZsrvResponse) (d []*map[string]interface{}) {
	for _, ZsrvResponseModel := range m {
		if ZsrvResponseModel != nil {
			properties := make(map[string]interface{})
			properties["end_time"] = ZsrvResponseModel.EndTime
			properties["error"] = SetZsrvErrorSubResourceData(ZsrvResponseModel.Error)
			properties["http_status_code"] = ZsrvResponseModel.HTTPStatusCode
			properties["http_status_msg"] = ZsrvResponseModel.HTTPStatusMsg
			properties["job_id"] = ZsrvResponseModel.JobID
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
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
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

		"job_id": {
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
			Type:        schema.TypeString,
			Optional:    true,
		},

		"operation_status": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"operation_time": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"operation_type": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
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
		"job_id",
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
