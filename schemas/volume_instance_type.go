package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

func VolumeInstanceTypeModel(d *schema.ResourceData) *models.VolumeInstanceType {
	volumeInstanceType, _ := d.Get("volume_instance_type").(models.VolumeInstanceType)
	return &volumeInstanceType
}

func VolumeInstanceTypeModelFromMap(m map[string]interface{}) *models.VolumeInstanceType {
	volumeInstanceType := m["volume_instance_type"].(models.VolumeInstanceType)
	return &volumeInstanceType
}

// Update the underlying VolumeInstanceType resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetVolumeInstanceTypeResourceData(d *schema.ResourceData, m *models.VolumeInstanceType) {
}

// Iterate through and update the VolumeInstanceType resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetVolumeInstanceTypeSubResourceData(m []*models.VolumeInstanceType) (d []*map[string]interface{}) {
	for _, VolumeInstanceTypeModel := range m {
		if VolumeInstanceTypeModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the VolumeInstanceType resource defined in the Terraform configuration
func VolumeInstanceTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the VolumeInstanceType resource
func GetVolumeInstanceTypePropertyFields() (t []string) {
	return []string{}
}
