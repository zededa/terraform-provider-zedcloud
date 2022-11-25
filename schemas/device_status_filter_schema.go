package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate DeviceStatusFilter resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func DeviceStatusFilterModel(d *schema.ResourceData) *models.DeviceStatusFilter {
	load := d.Get("load").(*models.DeviceLoad) // DeviceLoad
	namePattern := d.Get("name_pattern").(string)
	projectName := d.Get("project_name").(string)
	runState := d.Get("run_state").(*models.RunState) // RunState
	tags := d.Get("tags").(map[string]string)         // map[string]string
	return &models.DeviceStatusFilter{
		Load:        load,
		NamePattern: namePattern,
		ProjectName: projectName,
		RunState:    runState,
		Tags:        tags,
	}
}

func DeviceStatusFilterModelFromMap(m map[string]interface{}) *models.DeviceStatusFilter {
	load := m["load"].(*models.DeviceLoad) // DeviceLoad
	namePattern := m["name_pattern"].(string)
	projectName := m["project_name"].(string)
	runState := m["run_state"].(*models.RunState) // RunState
	tags := m["tags"].(map[string]string)
	return &models.DeviceStatusFilter{
		Load:        load,
		NamePattern: namePattern,
		ProjectName: projectName,
		RunState:    runState,
		Tags:        tags,
	}
}

// Update the underlying DeviceStatusFilter resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDeviceStatusFilterResourceData(d *schema.ResourceData, m *models.DeviceStatusFilter) {
	d.Set("load", m.Load)
	d.Set("name_pattern", m.NamePattern)
	d.Set("project_name", m.ProjectName)
	d.Set("run_state", m.RunState)
	d.Set("tags", m.Tags)
}

// Iterate throught and update the DeviceStatusFilter resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDeviceStatusFilterSubResourceData(m []*models.DeviceStatusFilter) (d []*map[string]interface{}) {
	for _, DeviceStatusFilterModel := range m {
		if DeviceStatusFilterModel != nil {
			properties := make(map[string]interface{})
			properties["load"] = DeviceStatusFilterModel.Load
			properties["name_pattern"] = DeviceStatusFilterModel.NamePattern
			properties["project_name"] = DeviceStatusFilterModel.ProjectName
			properties["run_state"] = DeviceStatusFilterModel.RunState
			properties["tags"] = DeviceStatusFilterModel.Tags
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the DeviceStatusFilter resource defined in the Terraform configuration
func DeviceStatusFilterSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"load": {
			Description: ``,
			Type:        schema.TypeList, //GoType: DeviceLoad
			Elem: &schema.Resource{
				Schema: DeviceLoadSchema(),
			},
			Optional: true,
		},

		"name_pattern": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"project_name": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"run_state": {
			Description: ``,
			Type:        schema.TypeList, //GoType: RunState
			Elem: &schema.Resource{
				Schema: RunStateSchema(),
			},
			Optional: true,
		},

		"tags": {
			Description: ``,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},
	}
}

// Retrieve property field names for updating the DeviceStatusFilter resource
func GetDeviceStatusFilterPropertyFields() (t []string) {
	return []string{
		"load",
		"name_pattern",
		"project_name",
		"run_state",
		"tags",
	}
}
