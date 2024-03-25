package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

// Function to perform the following actions:
// (1) Translate DeviceSWInfo resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func DeviceSWInfoModel(d *schema.ResourceData) *models.DeviceSWInfo {
	activated, _ := d.Get("activated").(bool)
	downloadProgressInt, _ := d.Get("download_progress").(int)
	downloadProgress := int64(downloadProgressInt)
	longVersion, _ := d.Get("long_version").(string)
	partitionDevice, _ := d.Get("partition_device").(string)
	partitionLabel, _ := d.Get("partition_label").(string)
	partitionState, _ := d.Get("partition_state").(string)
	shortVersion, _ := d.Get("short_version").(string)
	var status *models.SWState // SWState
	statusInterface, statusIsSet := d.GetOk("status")
	if statusIsSet {
		statusModel := statusInterface.(string)
		status = models.NewSWState(models.SWState(statusModel))
	}
	subStatusProgressInt, _ := d.Get("sub_status_progress").(int)
	subStatusProgress := int64(subStatusProgressInt)
	var swError *models.DeviceError // DeviceError
	swErrorInterface, swErrorIsSet := d.GetOk("sw_error")
	if swErrorIsSet {
		swErrorMap := swErrorInterface.([]interface{})[0].(map[string]interface{})
		swError = DeviceErrorModelFromMap(swErrorMap)
	}
	var swStatus *models.DeviceSWStatus // DeviceSWStatus
	swStatusInterface, swStatusIsSet := d.GetOk("sw_status")
	if swStatusIsSet {
		swStatusModel := swStatusInterface.(string)
		swStatus = models.NewDeviceSWStatus(models.DeviceSWStatus(swStatusModel))
	}
	var swSubStatus *models.DeviceSWSubStatus // DeviceSWSubStatus
	swSubStatusInterface, swSubStatusIsSet := d.GetOk("sw_sub_status")
	if swSubStatusIsSet {
		swSubStatusModel := swSubStatusInterface.(string)
		swSubStatus = models.NewDeviceSWSubStatus(models.DeviceSWSubStatus(swSubStatusModel))
	}
	swSubStatusStr, _ := d.Get("sw_sub_status_str").(string)
	return &models.DeviceSWInfo{
		Activated:         activated,
		DownloadProgress:  downloadProgress,
		LongVersion:       longVersion,
		PartitionDevice:   partitionDevice,
		PartitionLabel:    partitionLabel,
		PartitionState:    partitionState,
		ShortVersion:      shortVersion,
		Status:            status,
		SubStatusProgress: subStatusProgress,
		SwError:           swError,
		SwStatus:          swStatus,
		SwSubStatus:       swSubStatus,
		SwSubStatusStr:    swSubStatusStr,
	}
}

func DeviceSWInfoModelFromMap(m map[string]interface{}) *models.DeviceSWInfo {
	activated := m["activated"].(bool)
	downloadProgress := int64(m["download_progress"].(int)) // int64 false false false
	longVersion := m["long_version"].(string)
	partitionDevice := m["partition_device"].(string)
	partitionLabel := m["partition_label"].(string)
	partitionState := m["partition_state"].(string)
	shortVersion := m["short_version"].(string)
	status := m["status"].(*models.SWState)                    // SWState
	subStatusProgress := int64(m["sub_status_progress"].(int)) // int64 false false false
	var swError *models.DeviceError                            // DeviceError
	swErrorInterface, swErrorIsSet := m["sw_error"]
	if swErrorIsSet {
		swErrorMap := swErrorInterface.([]interface{})[0].(map[string]interface{})
		swError = DeviceErrorModelFromMap(swErrorMap)
	}
	//
	swStatus := m["sw_status"].(*models.DeviceSWStatus)           // DeviceSWStatus
	swSubStatus := m["sw_sub_status"].(*models.DeviceSWSubStatus) // DeviceSWSubStatus
	swSubStatusStr := m["sw_sub_status_str"].(string)
	return &models.DeviceSWInfo{
		Activated:         activated,
		DownloadProgress:  downloadProgress,
		LongVersion:       longVersion,
		PartitionDevice:   partitionDevice,
		PartitionLabel:    partitionLabel,
		PartitionState:    partitionState,
		ShortVersion:      shortVersion,
		Status:            status,
		SubStatusProgress: subStatusProgress,
		SwError:           swError,
		SwStatus:          swStatus,
		SwSubStatus:       swSubStatus,
		SwSubStatusStr:    swSubStatusStr,
	}
}

// Update the underlying DeviceSWInfo resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDeviceSWInfoResourceData(d *schema.ResourceData, m *models.DeviceSWInfo) {
	d.Set("activated", m.Activated)
	d.Set("download_progress", m.DownloadProgress)
	d.Set("long_version", m.LongVersion)
	d.Set("partition_device", m.PartitionDevice)
	d.Set("partition_label", m.PartitionLabel)
	d.Set("partition_state", m.PartitionState)
	d.Set("short_version", m.ShortVersion)
	d.Set("status", m.Status)
	d.Set("sub_status_progress", m.SubStatusProgress)
	d.Set("sw_error", SetDeviceErrorSubResourceData([]*models.DeviceError{m.SwError}))
	d.Set("sw_status", m.SwStatus)
	d.Set("sw_sub_status", m.SwSubStatus)
	d.Set("sw_sub_status_str", m.SwSubStatusStr)
}

// Iterate through and update the DeviceSWInfo resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDeviceSWInfoSubResourceData(m []*models.DeviceSWInfo) (d []*map[string]interface{}) {
	for _, DeviceSWInfoModel := range m {
		if DeviceSWInfoModel != nil {
			properties := make(map[string]interface{})
			properties["activated"] = DeviceSWInfoModel.Activated
			properties["download_progress"] = DeviceSWInfoModel.DownloadProgress
			properties["long_version"] = DeviceSWInfoModel.LongVersion
			properties["partition_device"] = DeviceSWInfoModel.PartitionDevice
			properties["partition_label"] = DeviceSWInfoModel.PartitionLabel
			properties["partition_state"] = DeviceSWInfoModel.PartitionState
			properties["short_version"] = DeviceSWInfoModel.ShortVersion
			properties["status"] = DeviceSWInfoModel.Status
			properties["sub_status_progress"] = DeviceSWInfoModel.SubStatusProgress
			properties["sw_error"] = SetDeviceErrorSubResourceData([]*models.DeviceError{DeviceSWInfoModel.SwError})
			properties["sw_status"] = DeviceSWInfoModel.SwStatus
			properties["sw_sub_status"] = DeviceSWInfoModel.SwSubStatus
			properties["sw_sub_status_str"] = DeviceSWInfoModel.SwSubStatusStr
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the DeviceSWInfo resource defined in the Terraform configuration
func DeviceSWInfoSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"activated": {
			Description: ``,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"download_progress": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"long_version": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"partition_device": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"partition_label": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"partition_state": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"short_version": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"status": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"sub_status_progress": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"sw_error": {
			Description: ``,
			Type:        schema.TypeList, //GoType: DeviceError
			Elem: &schema.Resource{
				Schema: DeviceErrorSchema(),
			},
			Optional: true,
		},

		"sw_status": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"sw_sub_status": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"sw_sub_status_str": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the DeviceSWInfo resource
func GetDeviceSWInfoPropertyFields() (t []string) {
	return []string{
		"activated",
		"download_progress",
		"long_version",
		"partition_device",
		"partition_label",
		"partition_state",
		"short_version",
		"status",
		"sub_status_progress",
		"sw_error",
		"sw_status",
		"sw_sub_status",
		"sw_sub_status_str",
	}
}
