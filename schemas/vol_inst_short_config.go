package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

func VolInstShortConfigModel(d *schema.ResourceData) *models.VolInstShortConfig {
	deviceID, _ := d.Get("device_id").(string)
	id, _ := d.Get("id").(string)
	name, _ := d.Get("name").(string)
	projectID, _ := d.Get("project_id").(string)
	var typeVar *models.VolumeInstanceType // VolumeInstanceType
	typeInterface, typeIsSet := d.GetOk("type")
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewVolumeInstanceType(models.VolumeInstanceType(typeModel))
	}
	return &models.VolInstShortConfig{
		DeviceID:  deviceID,
		ID:        id,
		Name:      name,
		ProjectID: projectID,
		Type:      typeVar,
	}
}

func VolInstShortConfigModelFromMap(m map[string]interface{}) *models.VolInstShortConfig {
	deviceID := m["device_id"].(string)
	id := m["id"].(string)
	name := m["name"].(string)
	projectID := m["project_id"].(string)
	var typeVar *models.VolumeInstanceType // VolumeInstanceType
	typeInterface, typeIsSet := m["type"]
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewVolumeInstanceType(models.VolumeInstanceType(typeModel))
	}
	return &models.VolInstShortConfig{
		DeviceID:  deviceID,
		ID:        id,
		Name:      name,
		ProjectID: projectID,
		Type:      typeVar,
	}
}

// Update the underlying VolInstShortConfig resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetVolInstShortConfigResourceData(d *schema.ResourceData, m *models.VolInstShortConfig) {
	d.Set("device_id", m.DeviceID)
	d.Set("id", m.ID)
	d.Set("name", m.Name)
	d.Set("project_id", m.ProjectID)
	d.Set("type", m.Type)
}

// Iterate through and update the VolInstShortConfig resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetVolInstShortConfigSubResourceData(m []*models.VolInstShortConfig) (d []*map[string]interface{}) {
	for _, VolInstShortConfigModel := range m {
		if VolInstShortConfigModel != nil {
			properties := make(map[string]interface{})
			properties["device_id"] = VolInstShortConfigModel.DeviceID
			properties["id"] = VolInstShortConfigModel.ID
			properties["name"] = VolInstShortConfigModel.Name
			properties["project_id"] = VolInstShortConfigModel.ProjectID
			properties["type"] = VolInstShortConfigModel.Type
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the VolInstShortConfig resource defined in the Terraform configuration
func VolInstShortConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"device_id": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"id": {
			Description: ``,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"name": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"project_id": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"type": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the VolInstShortConfig resource
func GetVolInstShortConfigPropertyFields() (t []string) {
	return []string{
		"device_id",
		"id",
		"name",
		"project_id",
		"type",
	}
}
