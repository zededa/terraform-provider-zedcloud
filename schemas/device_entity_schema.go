package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate DeviceEntity resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func DeviceEntityModel(d *schema.ResourceData) *models.DeviceEntity {
	entity := d.Get("entity").(*models.Entity) // Entity
	entityID := d.Get("entity_id").(string)
	return &models.DeviceEntity{
		Entity:   entity,
		EntityID: entityID,
	}
}

func DeviceEntityModelFromMap(m map[string]interface{}) *models.DeviceEntity {
	entity := m["entity"].(*models.Entity) // Entity
	entityID := m["entity_id"].(string)
	return &models.DeviceEntity{
		Entity:   entity,
		EntityID: entityID,
	}
}

// Update the underlying DeviceEntity resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDeviceEntityResourceData(d *schema.ResourceData, m *models.DeviceEntity) {
	d.Set("entity", m.Entity)
	d.Set("entity_id", m.EntityID)
}

// Iterate throught and update the DeviceEntity resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDeviceEntitySubResourceData(m []*models.DeviceEntity) (d []*map[string]interface{}) {
	for _, DeviceEntityModel := range m {
		if DeviceEntityModel != nil {
			properties := make(map[string]interface{})
			properties["entity"] = DeviceEntityModel.Entity
			properties["entity_id"] = DeviceEntityModel.EntityID
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
			Type:        schema.TypeList, //GoType: Entity
			Elem: &schema.Resource{
				Schema: EntitySchema(),
			},
			Optional: true,
		},

		"entity_id": {
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
	}
}
