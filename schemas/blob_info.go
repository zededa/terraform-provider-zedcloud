package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func BlobInfoModel(d *schema.ResourceData) *models.BlobInfo {
	var state *models.SWState // SWState
	StateInterface, StateIsSet := d.GetOk("state")
	if StateIsSet {
		StateModel := StateInterface.(string)
		state = models.NewSWState(models.SWState(StateModel))
	}
	var errInfo *models.DeviceError // DeviceError
	errInfoInterface, errInfoIsSet := d.GetOk("err_info")
	if errInfoIsSet && errInfoInterface != nil {
		errInfoMap := errInfoInterface.([]interface{})
		if len(errInfoMap) > 0 {
			errInfo = DeviceErrorModelFromMap(errInfoMap[0].(map[string]interface{}))
		}
	}
	progressPercentageInt, _ := d.Get("progress_percentage").(int)
	progressPercentage := int64(progressPercentageInt)
	var resource *models.VolInstResource // VolInstResource
	resourceInterface, resourceIsSet := d.GetOk("resource")
	if resourceIsSet && resourceInterface != nil {
		resourceMap := resourceInterface.([]interface{})
		if len(resourceMap) > 0 {
			resource = VolInstResourceModelFromMap(resourceMap[0].(map[string]interface{}))
		}
	}
	sha256, _ := d.Get("sha256").(string)
	var usage *models.DeviceObjectUsageInfo // DeviceObjectUsageInfo
	usageInterface, usageIsSet := d.GetOk("usage")
	if usageIsSet && usageInterface != nil {
		usageMap := usageInterface.([]interface{})
		if len(usageMap) > 0 {
			usage = DeviceObjectUsageInfoModelFromMap(usageMap[0].(map[string]interface{}))
		}
	}
	return &models.BlobInfo{
		State:              state,
		ErrInfo:            errInfo,
		ProgressPercentage: progressPercentage,
		Resource:           resource,
		Sha256:             sha256,
		Usage:              usage,
	}
}

func BlobInfoModelFromMap(m map[string]interface{}) *models.BlobInfo {
	var state *models.SWState // SWState
	StateInterface, StateIsSet := m["state"]
	if StateIsSet {
		StateModel := StateInterface.(string)
		state = models.NewSWState(models.SWState(StateModel))
	}
	var errInfo *models.DeviceError // DeviceError
	errInfoInterface, errInfoIsSet := m["err_info"]
	if errInfoIsSet && errInfoInterface != nil {
		errInfoMap := errInfoInterface.([]interface{})
		if len(errInfoMap) > 0 {
			errInfo = DeviceErrorModelFromMap(errInfoMap[0].(map[string]interface{}))
		}
	}
	//
	progressPercentage := int64(m["progress_percentage"].(int)) // int64
	var resource *models.VolInstResource                        // VolInstResource
	resourceInterface, resourceIsSet := m["resource"]
	if resourceIsSet && resourceInterface != nil {
		resourceMap := resourceInterface.([]interface{})
		if len(resourceMap) > 0 {
			resource = VolInstResourceModelFromMap(resourceMap[0].(map[string]interface{}))
		}
	}
	//
	sha256 := m["sha256"].(string)
	var usage *models.DeviceObjectUsageInfo // DeviceObjectUsageInfo
	usageInterface, usageIsSet := m["usage"]
	if usageIsSet && usageInterface != nil {
		usageMap := usageInterface.([]interface{})
		if len(usageMap) > 0 {
			usage = DeviceObjectUsageInfoModelFromMap(usageMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.BlobInfo{
		State:              state,
		ErrInfo:            errInfo,
		ProgressPercentage: progressPercentage,
		Resource:           resource,
		Sha256:             sha256,
		Usage:              usage,
	}
}

// Update the underlying BlobInfo resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetBlobInfoResourceData(d *schema.ResourceData, m *models.BlobInfo) {
	d.Set("state", m.State)
	d.Set("err_info", SetDeviceErrorSubResourceData([]*models.DeviceError{m.ErrInfo}))
	d.Set("progress_percentage", m.ProgressPercentage)
	d.Set("resource", SetVolInstResourceSubResourceData([]*models.VolInstResource{m.Resource}))
	d.Set("sha256", m.Sha256)
	d.Set("usage", SetDeviceObjectUsageInfoSubResourceData([]*models.DeviceObjectUsageInfo{m.Usage}))
}

// Iterate through and update the BlobInfo resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetBlobInfoSubResourceData(m []*models.BlobInfo) (d []*map[string]interface{}) {
	for _, BlobInfoModel := range m {
		if BlobInfoModel != nil {
			properties := make(map[string]interface{})
			properties["state"] = BlobInfoModel.State
			properties["err_info"] = SetDeviceErrorSubResourceData([]*models.DeviceError{BlobInfoModel.ErrInfo})
			properties["progress_percentage"] = BlobInfoModel.ProgressPercentage
			properties["resource"] = SetVolInstResourceSubResourceData([]*models.VolInstResource{BlobInfoModel.Resource})
			properties["sha256"] = BlobInfoModel.Sha256
			properties["usage"] = SetDeviceObjectUsageInfoSubResourceData([]*models.DeviceObjectUsageInfo{BlobInfoModel.Usage})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the BlobInfo resource defined in the Terraform configuration
func BlobInfoSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"state": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"err_info": {
			Description: ``,
			Type:        schema.TypeList, //GoType: DeviceError
			Elem: &schema.Resource{
				Schema: DeviceErrorSchema(),
			},
			Optional: true,
		},

		"progress_percentage": {
			Description: ``,
			Type:        schema.TypeInt,
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

		"sha256": {
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

// Retrieve property field names for updating the BlobInfo resource
func GetBlobInfoPropertyFields() (t []string) {
	return []string{
		"state",
		"err_info",
		"progress_percentage",
		"resource",
		"sha256",
		"usage",
	}
}
