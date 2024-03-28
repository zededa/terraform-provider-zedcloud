package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

func ObjectTypeModel(d *schema.ResourceData) *models.ObjectType {
	objectType, _ := d.Get("object_type").(models.ObjectType)
	return &objectType
}

func ObjectTypeModelFromMap(m map[string]interface{}) *models.ObjectType {
	objectType := m["object_type"].(models.ObjectType)
	return &objectType
}

// Update the underlying ObjectType resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetObjectTypeResourceData(d *schema.ResourceData, m *models.ObjectType) {
}

// Iterate through and update the ObjectType resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetObjectTypeSubResourceData(m []*models.ObjectType) (d []*map[string]interface{}) {
	for _, ObjectTypeModel := range m {
		if ObjectTypeModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the ObjectType resource defined in the Terraform configuration
func ObjectTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the ObjectType resource
func GetObjectTypePropertyFields() (t []string) {
	return []string{}
}
