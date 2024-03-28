package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

func IoTypeModel(d *schema.ResourceData) *models.IoType {
	ioType, _ := d.Get("io_type").(models.IoType)
	return &ioType
}

func IoTypeModelFromMap(m map[string]interface{}) *models.IoType {
	ioType := m["io_type"].(models.IoType)
	return &ioType
}

// Update the underlying IoType resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetIoTypeResourceData(d *schema.ResourceData, m *models.IoType) {
}

// Iterate through and update the IoType resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetIoTypeSubResourceData(m []*models.IoType) (d []*map[string]interface{}) {
	for _, IoTypeModel := range m {
		if IoTypeModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the IoType resource defined in the Terraform configuration
func IoTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the IoType resource
func GetIoTypePropertyFields() (t []string) {
	return []string{}
}
