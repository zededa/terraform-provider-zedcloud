package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func DeviceEntityModel(d *schema.ResourceData) *models.DeviceEntity {
	var entity *models.Entity // Entity
	entityInterface, entityIsSet := d.GetOk("entity")
	if entityIsSet {
		entityModel := entityInterface.(string)
		entity = models.NewEntity(models.Entity(entityModel))
	}
	entityID, _ := d.Get("entity_id").(string)
	entityName, _ := d.Get("entity_name").(string)
	return &models.DeviceEntity{
		Entity:     entity,
		EntityID:   entityID,
		EntityName: entityName,
	}
}

func DeviceEntityModelFromMap(m map[string]interface{}) *models.DeviceEntity {
	var entity *models.Entity // Entity
	entityInterface, entityIsSet := m["entity"]
	if entityIsSet {
		entityModel := entityInterface.(string)
		entity = models.NewEntity(models.Entity(entityModel))
	}
	entityID := m["entity_id"].(string)
	entityName := m["entity_name"].(string)
	return &models.DeviceEntity{
		Entity:     entity,
		EntityID:   entityID,
		EntityName: entityName,
	}
}

// Update the underlying DeviceEntity resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDeviceEntityResourceData(d *schema.ResourceData, m *models.DeviceEntity) {
	d.Set("entity", m.Entity)
	d.Set("entity_id", m.EntityID)
	d.Set("entity_name", m.EntityName)
}

// Iterate through and update the DeviceEntity resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDeviceEntitySubResourceData(m []*models.DeviceEntity) (d []*map[string]interface{}) {
	for _, DeviceEntityModel := range m {
		if DeviceEntityModel != nil {
			properties := make(map[string]interface{})
			properties["entity"] = DeviceEntityModel.Entity
			properties["entity_id"] = DeviceEntityModel.EntityID
			properties["entity_name"] = DeviceEntityModel.EntityName
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the DeviceEntity resource defined in the Terraform configuration
func DeviceEntitySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"entity": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"entity_id": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"entity_name": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the DeviceEntity resource
func GetDeviceEntityPropertyFields() (t []string) {
	return []string{
		"entity",
		"entity_id",
		"entity_name",
	}
}
