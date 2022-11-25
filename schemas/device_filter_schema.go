package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate DeviceFilter resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func DeviceFilterModel(d *schema.ResourceData) *models.DeviceFilter {
	namePattern := d.Get("name_pattern").(string)
	project := d.Get("project").(string)
	return &models.DeviceFilter{
		NamePattern: &namePattern, // string true false false
		Project:     &project,     // string true false false
	}
}

func DeviceFilterModelFromMap(m map[string]interface{}) *models.DeviceFilter {
	namePattern := m["name_pattern"].(string)
	project := m["project"].(string)
	return &models.DeviceFilter{
		NamePattern: &namePattern,
		Project:     &project,
	}
}

// Update the underlying DeviceFilter resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDeviceFilterResourceData(d *schema.ResourceData, m *models.DeviceFilter) {
	d.Set("name_pattern", m.NamePattern)
	d.Set("project", m.Project)
}

// Iterate throught and update the DeviceFilter resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDeviceFilterSubResourceData(m []*models.DeviceFilter) (d []*map[string]interface{}) {
	for _, DeviceFilterModel := range m {
		if DeviceFilterModel != nil {
			properties := make(map[string]interface{})
			properties["name_pattern"] = DeviceFilterModel.NamePattern
			properties["project"] = DeviceFilterModel.Project
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the DeviceFilter resource defined in the Terraform configuration
func DeviceFilterSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name_pattern": {
			Description: `name pattern`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"project": {
			Description: `project`,
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

// Retrieve property field names for updating the DeviceFilter resource
func GetDeviceFilterPropertyFields() (t []string) {
	return []string{
		"name_pattern",
		"project",
	}
}
