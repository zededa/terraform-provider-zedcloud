package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate DiskDescription resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func DiskDescriptionModel(d *schema.ResourceData) *models.DiskDescription {
	logicalName, _ := d.Get("logical_name").(string)
	name, _ := d.Get("name").(string)
	serial, _ := d.Get("serial").(string)
	return &models.DiskDescription{
		LogicalName: logicalName,
		Name:        name,
		Serial:      serial,
	}
}

func DiskDescriptionModelFromMap(m map[string]interface{}) *models.DiskDescription {
	logicalName := m["logical_name"].(string)
	name := m["name"].(string)
	serial := m["serial"].(string)
	return &models.DiskDescription{
		LogicalName: logicalName,
		Name:        name,
		Serial:      serial,
	}
}

// Update the underlying DiskDescription resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDiskDescriptionResourceData(d *schema.ResourceData, m *models.DiskDescription) {
	d.Set("logical_name", m.LogicalName)
	d.Set("name", m.Name)
	d.Set("serial", m.Serial)
}

// Iterate through and update the DiskDescription resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDiskDescriptionSubResourceData(m []*models.DiskDescription) (d []*map[string]interface{}) {
	for _, DiskDescriptionModel := range m {
		if DiskDescriptionModel != nil {
			properties := make(map[string]interface{})
			properties["logical_name"] = DiskDescriptionModel.LogicalName
			properties["name"] = DiskDescriptionModel.Name
			properties["serial"] = DiskDescriptionModel.Serial
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the DiskDescription resource defined in the Terraform configuration
func DiskDescriptionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"logical_name": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"name": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"serial": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the DiskDescription resource
func GetDiskDescriptionPropertyFields() (t []string) {
	return []string{
		"logical_name",
		"name",
		"serial",
	}
}
