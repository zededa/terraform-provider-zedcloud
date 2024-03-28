package schemas

import (
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func VolInstStatusSummaryMsgModel(d *schema.ResourceData) *models.VolInstStatusSummaryMsg {
	createTime, _ := d.Get("create_time").(strfmt.DateTime)
	deviceID, _ := d.Get("device_id").(string)
	deviceName, _ := d.Get("device_name").(string)
	var deviceState *models.SWState // SWState
	deviceStateInterface, deviceStateIsSet := d.GetOk("device_state")
	if deviceStateIsSet {
		deviceStateModel := deviceStateInterface.(string)
		deviceState = models.NewSWState(models.SWState(deviceStateModel))
	}
	id, _ := d.Get("id").(string)
	name, _ := d.Get("name").(string)
	progressPercentageInt, _ := d.Get("progress_percentage").(int)
	progressPercentage := int64(progressPercentageInt)
	projectID, _ := d.Get("project_id").(string)
	projectName, _ := d.Get("project_name").(string)
	var runState *models.RunState // RunState
	runStateInterface, runStateIsSet := d.GetOk("run_state")
	if runStateIsSet {
		runStateModel := runStateInterface.(string)
		runState = models.NewRunState(models.RunState(runStateModel))
	}
	var typeVar *models.VolumeInstanceType // VolumeInstanceType
	typeInterface, typeIsSet := d.GetOk("type")
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewVolumeInstanceType(models.VolumeInstanceType(typeModel))
	}
	return &models.VolInstStatusSummaryMsg{
		CreateTime:         createTime,
		DeviceID:           deviceID,
		DeviceName:         deviceName,
		DeviceState:        deviceState,
		ID:                 id,
		Name:               name,
		ProgressPercentage: progressPercentage,
		ProjectID:          projectID,
		ProjectName:        projectName,
		RunState:           runState,
		Type:               typeVar,
	}
}

func VolInstStatusSummaryMsgModelFromMap(m map[string]interface{}) *models.VolInstStatusSummaryMsg {
	createTime := m["create_time"].(strfmt.DateTime)
	deviceID := m["device_id"].(string)
	deviceName := m["device_name"].(string)
	var deviceState *models.SWState // SWState
	deviceStateInterface, deviceStateIsSet := m["device_state"]
	if deviceStateIsSet {
		deviceStateModel := deviceStateInterface.(string)
		deviceState = models.NewSWState(models.SWState(deviceStateModel))
	}
	id := m["id"].(string)
	name := m["name"].(string)
	progressPercentage := int64(m["progress_percentage"].(int)) // int64
	projectID := m["project_id"].(string)
	projectName := m["project_name"].(string)
	var runState *models.RunState // RunState
	runStateInterface, runStateIsSet := m["run_state"]
	if runStateIsSet {
		runStateModel := runStateInterface.(string)
		runState = models.NewRunState(models.RunState(runStateModel))
	}
	var typeVar *models.VolumeInstanceType // VolumeInstanceType
	typeInterface, typeIsSet := m["type"]
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewVolumeInstanceType(models.VolumeInstanceType(typeModel))
	}
	return &models.VolInstStatusSummaryMsg{
		CreateTime:         createTime,
		DeviceID:           deviceID,
		DeviceName:         deviceName,
		DeviceState:        deviceState,
		ID:                 id,
		Name:               name,
		ProgressPercentage: progressPercentage,
		ProjectID:          projectID,
		ProjectName:        projectName,
		RunState:           runState,
		Type:               typeVar,
	}
}

// Update the underlying VolInstStatusSummaryMsg resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetVolInstStatusSummaryMsgResourceData(d *schema.ResourceData, m *models.VolInstStatusSummaryMsg) {
	d.Set("create_time", m.CreateTime)
	d.Set("device_id", m.DeviceID)
	d.Set("device_name", m.DeviceName)
	d.Set("device_state", m.DeviceState)
	d.Set("id", m.ID)
	d.Set("name", m.Name)
	d.Set("progress_percentage", m.ProgressPercentage)
	d.Set("project_id", m.ProjectID)
	d.Set("project_name", m.ProjectName)
	d.Set("run_state", m.RunState)
	d.Set("type", m.Type)
}

// Iterate through and update the VolInstStatusSummaryMsg resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetVolInstStatusSummaryMsgSubResourceData(m []*models.VolInstStatusSummaryMsg) (d []*map[string]interface{}) {
	for _, VolInstStatusSummaryMsgModel := range m {
		if VolInstStatusSummaryMsgModel != nil {
			properties := make(map[string]interface{})
			properties["create_time"] = VolInstStatusSummaryMsgModel.CreateTime.String()
			properties["device_id"] = VolInstStatusSummaryMsgModel.DeviceID
			properties["device_name"] = VolInstStatusSummaryMsgModel.DeviceName
			properties["device_state"] = VolInstStatusSummaryMsgModel.DeviceState
			properties["id"] = VolInstStatusSummaryMsgModel.ID
			properties["name"] = VolInstStatusSummaryMsgModel.Name
			properties["progress_percentage"] = VolInstStatusSummaryMsgModel.ProgressPercentage
			properties["project_id"] = VolInstStatusSummaryMsgModel.ProjectID
			properties["project_name"] = VolInstStatusSummaryMsgModel.ProjectName
			properties["run_state"] = VolInstStatusSummaryMsgModel.RunState
			properties["type"] = VolInstStatusSummaryMsgModel.Type
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the VolInstStatusSummaryMsg resource defined in the Terraform configuration
func VolInstStatusSummaryMsgSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"create_time": {
			Description:  ``,
			Type:         schema.TypeString,
			ValidateFunc: validation.IsRFC3339Time,
			Optional:     true,
		},

		"device_id": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"device_name": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"device_state": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"id": {
			Description: ``,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"name": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"progress_percentage": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"project_id": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"project_name": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"run_state": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"type": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the VolInstStatusSummaryMsg resource
func GetVolInstStatusSummaryMsgPropertyFields() (t []string) {
	return []string{
		"create_time",
		"device_id",
		"device_name",
		"device_state",
		"id",
		"name",
		"progress_percentage",
		"project_id",
		"project_name",
		"run_state",
		"type",
	}
}
