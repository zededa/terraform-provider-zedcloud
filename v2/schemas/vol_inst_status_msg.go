package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func VolInstStatusMsgModel(d *schema.ResourceData) *models.VolInstStatusMsg {
	var blobs []*models.BlobInfo // []*BlobInfo
	blobsInterface, blobsIsSet := d.GetOk("blobs")
	if blobsIsSet {
		var items []interface{}
		if listItems, isList := blobsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = blobsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := BlobInfoModelFromMap(v.(map[string]interface{}))
			blobs = append(blobs, m)
		}
	}
	deviceID, _ := d.Get("device_id").(string)
	var deviceState *models.SWState // SWState
	deviceStateInterface, deviceStateIsSet := d.GetOk("device_state")
	if deviceStateIsSet {
		deviceStateModel := deviceStateInterface.(string)
		deviceState = models.NewSWState(models.SWState(deviceStateModel))
	}
	var errInfo []*models.DeviceError // []*DeviceError
	errInfoInterface, errInfoIsSet := d.GetOk("err_info")
	if errInfoIsSet {
		var items []interface{}
		if listItems, isList := errInfoInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = errInfoInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := DeviceErrorModelFromMap(v.(map[string]interface{}))
			errInfo = append(errInfo, m)
		}
	}
	id, _ := d.Get("id").(string)
	name, _ := d.Get("name").(string)
	progressPercentageInt, _ := d.Get("progress_percentage").(int)
	progressPercentage := int64(progressPercentageInt)
	projectID, _ := d.Get("project_id").(string)
	rawStatus, _ := d.Get("raw_status").(string)
	var resource *models.VolInstResource // VolInstResource
	resourceInterface, resourceIsSet := d.GetOk("resource")
	if resourceIsSet && resourceInterface != nil {
		resourceMap := resourceInterface.([]interface{})
		if len(resourceMap) > 0 {
			resource = VolInstResourceModelFromMap(resourceMap[0].(map[string]interface{}))
		}
	}
	var runState *models.RunState // RunState
	runStateInterface, runStateIsSet := d.GetOk("run_state")
	if runStateIsSet {
		runStateModel := runStateInterface.(string)
		runState = models.NewRunState(models.RunState(runStateModel))
	}
	sha256, _ := d.Get("sha256").(string)
	var typeVar *models.VolumeInstanceType // VolumeInstanceType
	typeInterface, typeIsSet := d.GetOk("type")
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewVolumeInstanceType(models.VolumeInstanceType(typeModel))
	}
	var usage *models.DeviceObjectUsageInfo // DeviceObjectUsageInfo
	usageInterface, usageIsSet := d.GetOk("usage")
	if usageIsSet && usageInterface != nil {
		usageMap := usageInterface.([]interface{})
		if len(usageMap) > 0 {
			usage = DeviceObjectUsageInfoModelFromMap(usageMap[0].(map[string]interface{}))
		}
	}
	return &models.VolInstStatusMsg{
		Blobs:              blobs,
		DeviceID:           deviceID,
		DeviceState:        deviceState,
		ErrInfo:            errInfo,
		ID:                 id,
		Name:               name,
		ProgressPercentage: progressPercentage,
		ProjectID:          projectID,
		RawStatus:          rawStatus,
		Resource:           resource,
		RunState:           runState,
		Sha256:             sha256,
		Type:               typeVar,
		Usage:              usage,
	}
}

func VolInstStatusMsgModelFromMap(m map[string]interface{}) *models.VolInstStatusMsg {
	var blobs []*models.BlobInfo // []*BlobInfo
	blobsInterface, blobsIsSet := m["blobs"]
	if blobsIsSet {
		var items []interface{}
		if listItems, isList := blobsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = blobsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := BlobInfoModelFromMap(v.(map[string]interface{}))
			blobs = append(blobs, m)
		}
	}
	deviceID := m["device_id"].(string)
	var deviceState *models.SWState // SWState
	deviceStateInterface, deviceStateIsSet := m["device_state"]
	if deviceStateIsSet {
		deviceStateModel := deviceStateInterface.(string)
		deviceState = models.NewSWState(models.SWState(deviceStateModel))
	}
	var errInfo []*models.DeviceError // []*DeviceError
	errInfoInterface, errInfoIsSet := m["err_info"]
	if errInfoIsSet {
		var items []interface{}
		if listItems, isList := errInfoInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = errInfoInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := DeviceErrorModelFromMap(v.(map[string]interface{}))
			errInfo = append(errInfo, m)
		}
	}
	id := m["id"].(string)
	name := m["name"].(string)
	progressPercentage := int64(m["progress_percentage"].(int)) // int64
	projectID := m["project_id"].(string)
	rawStatus := m["raw_status"].(string)
	var resource *models.VolInstResource // VolInstResource
	resourceInterface, resourceIsSet := m["resource"]
	if resourceIsSet && resourceInterface != nil {
		resourceMap := resourceInterface.([]interface{})
		if len(resourceMap) > 0 {
			resource = VolInstResourceModelFromMap(resourceMap[0].(map[string]interface{}))
		}
	}
	//
	var runState *models.RunState // RunState
	runStateInterface, runStateIsSet := m["run_state"]
	if runStateIsSet {
		runStateModel := runStateInterface.(string)
		runState = models.NewRunState(models.RunState(runStateModel))
	}
	sha256 := m["sha256"].(string)
	var typeVar *models.VolumeInstanceType // VolumeInstanceType
	typeInterface, typeIsSet := m["type"]
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewVolumeInstanceType(models.VolumeInstanceType(typeModel))
	}
	var usage *models.DeviceObjectUsageInfo // DeviceObjectUsageInfo
	usageInterface, usageIsSet := m["usage"]
	if usageIsSet && usageInterface != nil {
		usageMap := usageInterface.([]interface{})
		if len(usageMap) > 0 {
			usage = DeviceObjectUsageInfoModelFromMap(usageMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.VolInstStatusMsg{
		Blobs:              blobs,
		DeviceID:           deviceID,
		DeviceState:        deviceState,
		ErrInfo:            errInfo,
		ID:                 id,
		Name:               name,
		ProgressPercentage: progressPercentage,
		ProjectID:          projectID,
		RawStatus:          rawStatus,
		Resource:           resource,
		RunState:           runState,
		Sha256:             sha256,
		Type:               typeVar,
		Usage:              usage,
	}
}

// Update the underlying VolInstStatusMsg resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetVolInstStatusMsgResourceData(d *schema.ResourceData, m *models.VolInstStatusMsg) {
	d.Set("blobs", SetBlobInfoSubResourceData(m.Blobs))
	d.Set("device_id", m.DeviceID)
	d.Set("device_state", m.DeviceState)
	d.Set("err_info", SetDeviceErrorSubResourceData(m.ErrInfo))
	d.Set("id", m.ID)
	d.Set("name", m.Name)
	d.Set("progress_percentage", m.ProgressPercentage)
	d.Set("project_id", m.ProjectID)
	d.Set("raw_status", m.RawStatus)
	d.Set("resource", SetVolInstResourceSubResourceData([]*models.VolInstResource{m.Resource}))
	d.Set("run_state", m.RunState)
	d.Set("sha256", m.Sha256)
	d.Set("type", m.Type)
	d.Set("usage", SetDeviceObjectUsageInfoSubResourceData([]*models.DeviceObjectUsageInfo{m.Usage}))
}

// Iterate through and update the VolInstStatusMsg resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetVolInstStatusMsgSubResourceData(m []*models.VolInstStatusMsg) (d []*map[string]interface{}) {
	for _, VolInstStatusMsgModel := range m {
		if VolInstStatusMsgModel != nil {
			properties := make(map[string]interface{})
			properties["blobs"] = SetBlobInfoSubResourceData(VolInstStatusMsgModel.Blobs)
			properties["device_id"] = VolInstStatusMsgModel.DeviceID
			properties["device_state"] = VolInstStatusMsgModel.DeviceState
			properties["err_info"] = SetDeviceErrorSubResourceData(VolInstStatusMsgModel.ErrInfo)
			properties["id"] = VolInstStatusMsgModel.ID
			properties["name"] = VolInstStatusMsgModel.Name
			properties["progress_percentage"] = VolInstStatusMsgModel.ProgressPercentage
			properties["project_id"] = VolInstStatusMsgModel.ProjectID
			properties["raw_status"] = VolInstStatusMsgModel.RawStatus
			properties["resource"] = SetVolInstResourceSubResourceData([]*models.VolInstResource{VolInstStatusMsgModel.Resource})
			properties["run_state"] = VolInstStatusMsgModel.RunState
			properties["sha256"] = VolInstStatusMsgModel.Sha256
			properties["type"] = VolInstStatusMsgModel.Type
			properties["usage"] = SetDeviceObjectUsageInfoSubResourceData([]*models.DeviceObjectUsageInfo{VolInstStatusMsgModel.Usage})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the VolInstStatusMsg resource defined in the Terraform configuration
func VolInstStatusMsgSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"blobs": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []*BlobInfo
			Elem: &schema.Resource{
				Schema: BlobInfoSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"device_id": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"device_state": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"err_info": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []*DeviceError
			Elem: &schema.Resource{
				Schema: DeviceErrorSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
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

		"raw_status": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"resource": {
			Description: ``,
			Type:        schema.TypeList, //GoType: VolInstResource
			Elem: &schema.Resource{
				Schema: VolInstResourceSchema(),
			},
			Optional: true,
		},

		"run_state": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"sha256": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"type": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"usage": {
			Description: ``,
			Type:        schema.TypeList, //GoType: DeviceObjectUsageInfo
			Elem: &schema.Resource{
				Schema: DeviceObjectUsageInfoSchema(),
			},
			Optional: true,
		},
	}
}

// Retrieve property field names for updating the VolInstStatusMsg resource
func GetVolInstStatusMsgPropertyFields() (t []string) {
	return []string{
		"blobs",
		"device_id",
		"device_state",
		"err_info",
		"id",
		"name",
		"progress_percentage",
		"project_id",
		"raw_status",
		"resource",
		"run_state",
		"sha256",
		"type",
		"usage",
	}
}
