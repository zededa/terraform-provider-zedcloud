package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func VolInstStatusFilterModel(d *schema.ResourceData) *models.VolInstStatusFilter {
	appInstName, _ := d.Get("app_inst_name").(string)
	deviceName, _ := d.Get("device_name").(string)
	deviceNamePattern, _ := d.Get("device_name_pattern").(string)
	imageName, _ := d.Get("image_name").(string)
	namePattern, _ := d.Get("name_pattern").(string)
	projectName, _ := d.Get("project_name").(string)
	projectNamePattern, _ := d.Get("project_name_pattern").(string)
	var runState *models.RunState // RunState
	runStateInterface, runStateIsSet := d.GetOk("run_state")
	if runStateIsSet {
		runStateModel := runStateInterface.(string)
		runState = models.NewRunState(models.RunState(runStateModel))
	}
	var typeVar *models.VolumeInstanceType // VolumeInstanceType
	typeInterface, typeIsSet := d.GetOk("type")
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewVolumeInstanceType(models.VolumeInstanceType(typeModel))
	}
	return &models.VolInstStatusFilter{
		AppInstName:        appInstName,
		DeviceName:         deviceName,
		DeviceNamePattern:  deviceNamePattern,
		ImageName:          imageName,
		NamePattern:        namePattern,
		ProjectName:        projectName,
		ProjectNamePattern: projectNamePattern,
		RunState:           runState,
		Type:               typeVar,
	}
}

func VolInstStatusFilterModelFromMap(m map[string]interface{}) *models.VolInstStatusFilter {
	appInstName := m["app_inst_name"].(string)
	deviceName := m["device_name"].(string)
	deviceNamePattern := m["device_name_pattern"].(string)
	imageName := m["image_name"].(string)
	namePattern := m["name_pattern"].(string)
	projectName := m["project_name"].(string)
	projectNamePattern := m["project_name_pattern"].(string)
	var runState *models.RunState // RunState
	runStateInterface, runStateIsSet := m["run_state"]
	if runStateIsSet {
		runStateModel := runStateInterface.(string)
		runState = models.NewRunState(models.RunState(runStateModel))
	}
	var typeVar *models.VolumeInstanceType // VolumeInstanceType
	typeInterface, typeIsSet := m["type"]
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewVolumeInstanceType(models.VolumeInstanceType(typeModel))
	}
	return &models.VolInstStatusFilter{
		AppInstName:        appInstName,
		DeviceName:         deviceName,
		DeviceNamePattern:  deviceNamePattern,
		ImageName:          imageName,
		NamePattern:        namePattern,
		ProjectName:        projectName,
		ProjectNamePattern: projectNamePattern,
		RunState:           runState,
		Type:               typeVar,
	}
}

// Update the underlying VolInstStatusFilter resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetVolInstStatusFilterResourceData(d *schema.ResourceData, m *models.VolInstStatusFilter) {
	d.Set("app_inst_name", m.AppInstName)
	d.Set("device_name", m.DeviceName)
	d.Set("device_name_pattern", m.DeviceNamePattern)
	d.Set("image_name", m.ImageName)
	d.Set("name_pattern", m.NamePattern)
	d.Set("project_name", m.ProjectName)
	d.Set("project_name_pattern", m.ProjectNamePattern)
	d.Set("run_state", m.RunState)
	d.Set("type", m.Type)
}

// Iterate through and update the VolInstStatusFilter resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetVolInstStatusFilterSubResourceData(m []*models.VolInstStatusFilter) (d []*map[string]interface{}) {
	for _, VolInstStatusFilterModel := range m {
		if VolInstStatusFilterModel != nil {
			properties := make(map[string]interface{})
			properties["app_inst_name"] = VolInstStatusFilterModel.AppInstName
			properties["device_name"] = VolInstStatusFilterModel.DeviceName
			properties["device_name_pattern"] = VolInstStatusFilterModel.DeviceNamePattern
			properties["image_name"] = VolInstStatusFilterModel.ImageName
			properties["name_pattern"] = VolInstStatusFilterModel.NamePattern
			properties["project_name"] = VolInstStatusFilterModel.ProjectName
			properties["project_name_pattern"] = VolInstStatusFilterModel.ProjectNamePattern
			properties["run_state"] = VolInstStatusFilterModel.RunState
			properties["type"] = VolInstStatusFilterModel.Type
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the VolInstStatusFilter resource defined in the Terraform configuration
func VolInstStatusFilterSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"app_inst_name": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"device_name": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"device_name_pattern": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"image_name": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
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

		"project_name_pattern": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"run_state": {
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

// Retrieve property field names for updating the VolInstStatusFilter resource
func GetVolInstStatusFilterPropertyFields() (t []string) {
	return []string{
		"app_inst_name",
		"device_name",
		"device_name_pattern",
		"image_name",
		"name_pattern",
		"project_name",
		"project_name_pattern",
		"run_state",
		"type",
	}
}
