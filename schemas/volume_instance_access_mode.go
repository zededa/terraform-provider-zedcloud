package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func VolumeInstanceAccessModeModel(d *schema.ResourceData) *models.VolumeInstanceAccessMode {
	volumeInstanceAccessMode, _ := d.Get("volume_instance_access_mode").(models.VolumeInstanceAccessMode)
	return &volumeInstanceAccessMode
}

func VolumeInstanceAccessModeModelFromMap(m map[string]interface{}) *models.VolumeInstanceAccessMode {
	volumeInstanceAccessMode := m["volume_instance_access_mode"].(models.VolumeInstanceAccessMode)
	return &volumeInstanceAccessMode
}

// Update the underlying VolumeInstanceAccessMode resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetVolumeInstanceAccessModeResourceData(d *schema.ResourceData, m *models.VolumeInstanceAccessMode) {
}

// Iterate through and update the VolumeInstanceAccessMode resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetVolumeInstanceAccessModeSubResourceData(m []*models.VolumeInstanceAccessMode) (d []*map[string]interface{}) {
	for _, VolumeInstanceAccessModeModel := range m {
		if VolumeInstanceAccessModeModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the VolumeInstanceAccessMode resource defined in the Terraform configuration
func VolumeInstanceAccessModeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the VolumeInstanceAccessMode resource
func GetVolumeInstanceAccessModePropertyFields() (t []string) {
	return []string{}
}
