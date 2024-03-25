package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func EntityModel(d *schema.ResourceData) *models.Entity {
	entity, _ := d.Get("entity").(models.Entity)
	return &entity
}

func EntityModelFromMap(m map[string]interface{}) *models.Entity {
	entity := m["entity"].(models.Entity)
	return &entity
}

// Update the underlying Entity resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetEntityResourceData(d *schema.ResourceData, m *models.Entity) {
}

// Iterate through and update the Entity resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetEntitySubResourceData(m []*models.Entity) (d []*map[string]interface{}) {
	for _, EntityModel := range m {
		if EntityModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the Entity resource defined in the Terraform configuration
func EntitySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the Entity resource
func GetEntityPropertyFields() (t []string) {
	return []string{}
}
