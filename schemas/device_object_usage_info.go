package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func DeviceObjectUsageInfoModel(d *schema.ResourceData) *models.DeviceObjectUsageInfo {
	createTime, _ := d.Get("create_time").(interface{}) // interface{}

	lastRefCountChangeTime, _ := d.Get("last_ref_count_change_time").(interface{}) // interface{}

	refCountInt, _ := d.Get("ref_count").(int)
	refCount := int64(refCountInt)
	return &models.DeviceObjectUsageInfo{
		CreateTime:             createTime,
		LastRefCountChangeTime: lastRefCountChangeTime,
		RefCount:               &refCount, // int64 true false false
	}
}

func DeviceObjectUsageInfoModelFromMap(m map[string]interface{}) *models.DeviceObjectUsageInfo {
	createTime := m["create_time"].(interface{})

	lastRefCountChangeTime := m["last_ref_count_change_time"].(interface{})

	refCount := int64(m["ref_count"].(int)) // int64
	return &models.DeviceObjectUsageInfo{
		CreateTime:             createTime,
		LastRefCountChangeTime: lastRefCountChangeTime,
		RefCount:               &refCount,
	}
}

// Update the underlying DeviceObjectUsageInfo resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDeviceObjectUsageInfoResourceData(d *schema.ResourceData, m *models.DeviceObjectUsageInfo) {
	d.Set("create_time", m.CreateTime)
	d.Set("last_ref_count_change_time", m.LastRefCountChangeTime)
	d.Set("ref_count", m.RefCount)
}

// Iterate through and update the DeviceObjectUsageInfo resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDeviceObjectUsageInfoSubResourceData(m []*models.DeviceObjectUsageInfo) (d []*map[string]interface{}) {
	for _, DeviceObjectUsageInfoModel := range m {
		if DeviceObjectUsageInfoModel != nil {
			properties := make(map[string]interface{})
			properties["create_time"] = DeviceObjectUsageInfoModel.CreateTime
			properties["last_ref_count_change_time"] = DeviceObjectUsageInfoModel.LastRefCountChangeTime
			properties["ref_count"] = DeviceObjectUsageInfoModel.RefCount
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the DeviceObjectUsageInfo resource defined in the Terraform configuration
func DeviceObjectUsageInfoSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"create_time": {
			Description: `Timestamp at which object was created`,
			Type:        schema.TypeMap, //GoType: interface{}
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Required: true,
		},

		"last_ref_count_change_time": {
			Description: `Timestamp at which object refcount was last changed`,
			Type:        schema.TypeMap, //GoType: interface{}
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Required: true,
		},

		"ref_count": {
			Description: `Object refcount`,
			Type:        schema.TypeInt,
			Required:    true,
		},
	}
}

// Retrieve property field names for updating the DeviceObjectUsageInfo resource
func GetDeviceObjectUsageInfoPropertyFields() (t []string) {
	return []string{
		"create_time",
		"last_ref_count_change_time",
		"ref_count",
	}
}
