package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func ControllerTypeModel(d *schema.ResourceData) *models.ControllerType {
	controllerType, _ := d.Get("controller_type").(models.ControllerType)
	return &controllerType
}

func ControllerTypeModelFromMap(m map[string]interface{}) *models.ControllerType {
	controllerType := m["controller_type"].(models.ControllerType)
	return &controllerType
}

// Update the underlying ControllerType resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetControllerTypeResourceData(d *schema.ResourceData, m *models.ControllerType) {
}

// Iterate through and update the ControllerType resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetControllerTypeSubResourceData(m []*models.ControllerType) (d []*map[string]interface{}) {
	for _, ControllerTypeModel := range m {
		if ControllerTypeModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the ControllerType resource defined in the Terraform configuration
func ControllerTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the ControllerType resource
func GetControllerTypePropertyFields() (t []string) {
	return []string{}
}
