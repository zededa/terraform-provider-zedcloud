package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate DeviceFilter resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func DeviceFilterModel(d *schema.ResourceData) *models.DeviceFilter {
	var adminState *models.AdminState // AdminState
	adminStateInterface, adminStateIsSet := d.GetOk("admin_state")
	if adminStateIsSet {
		adminStateModel := adminStateInterface.(string)
		adminState = models.NewAdminState(models.AdminState(adminStateModel))
	}
	namePattern, _ := d.Get("name_pattern").(string)
	project, _ := d.Get("project").(string)
	projectNamePattern, _ := d.Get("project_name_pattern").(string)
	return &models.DeviceFilter{
		AdminState:         adminState,
		NamePattern:        &namePattern, // string true false false
		Project:            &project,     // string true false false
		ProjectNamePattern: projectNamePattern,
	}
}

func DeviceFilterModelFromMap(m map[string]interface{}) *models.DeviceFilter {
	adminState := m["admin_state"].(*models.AdminState) // AdminState
	namePattern := m["name_pattern"].(string)
	project := m["project"].(string)
	projectNamePattern := m["project_name_pattern"].(string)
	return &models.DeviceFilter{
		AdminState:         adminState,
		NamePattern:        &namePattern,
		Project:            &project,
		ProjectNamePattern: projectNamePattern,
	}
}

// Update the underlying DeviceFilter resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDeviceFilterResourceData(d *schema.ResourceData, m *models.DeviceFilter) {
	d.Set("admin_state", m.AdminState)
	d.Set("name_pattern", m.NamePattern)
	d.Set("project", m.Project)
	d.Set("project_name_pattern", m.ProjectNamePattern)
}

// Iterate through and update the DeviceFilter resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDeviceFilterSubResourceData(m []*models.DeviceFilter) (d []*map[string]interface{}) {
	for _, DeviceFilterModel := range m {
		if DeviceFilterModel != nil {
			properties := make(map[string]interface{})
			properties["admin_state"] = DeviceFilterModel.AdminState
			properties["name_pattern"] = DeviceFilterModel.NamePattern
			properties["project"] = DeviceFilterModel.Project
			properties["project_name_pattern"] = DeviceFilterModel.ProjectNamePattern
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the DeviceFilter resource defined in the Terraform configuration
func DeviceFilterSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"admin_state": {
			Description: `admin state of the device`,
			Type:        schema.TypeString,
			Optional:    true,
		},

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

		"project_name_pattern": {
			Description: `project name pattern`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the DeviceFilter resource
func GetDeviceFilterPropertyFields() (t []string) {
	return []string{
		"admin_state",
		"name_pattern",
		"project",
		"project_name_pattern",
	}
}
