package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func VolInstFilterModel(d *schema.ResourceData) *models.VolInstFilter {
	appInstName, _ := d.Get("app_inst_name").(string)
	deviceName, _ := d.Get("device_name").(string)
	deviceNamePattern, _ := d.Get("device_name_pattern").(string)
	labelName, _ := d.Get("label_name").(string)
	namePattern, _ := d.Get("name_pattern").(string)
	projectName, _ := d.Get("project_name").(string)
	projectNamePattern, _ := d.Get("project_name_pattern").(string)
	var typeVar *models.VolumeInstanceType // VolumeInstanceType
	typeInterface, typeIsSet := d.GetOk("type")
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewVolumeInstanceType(models.VolumeInstanceType(typeModel))
	}
	return &models.VolInstFilter{
		AppInstName:        appInstName,
		DeviceName:         deviceName,
		DeviceNamePattern:  deviceNamePattern,
		LabelName:          labelName,
		NamePattern:        namePattern,
		ProjectName:        projectName,
		ProjectNamePattern: projectNamePattern,
		Type:               typeVar,
	}
}

func VolInstFilterModelFromMap(m map[string]interface{}) *models.VolInstFilter {
	appInstName := m["app_inst_name"].(string)
	deviceName := m["device_name"].(string)
	deviceNamePattern := m["device_name_pattern"].(string)
	labelName := m["label_name"].(string)
	namePattern := m["name_pattern"].(string)
	projectName := m["project_name"].(string)
	projectNamePattern := m["project_name_pattern"].(string)
	var typeVar *models.VolumeInstanceType // VolumeInstanceType
	typeInterface, typeIsSet := m["type"]
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewVolumeInstanceType(models.VolumeInstanceType(typeModel))
	}
	return &models.VolInstFilter{
		AppInstName:        appInstName,
		DeviceName:         deviceName,
		DeviceNamePattern:  deviceNamePattern,
		LabelName:          labelName,
		NamePattern:        namePattern,
		ProjectName:        projectName,
		ProjectNamePattern: projectNamePattern,
		Type:               typeVar,
	}
}

// Update the underlying VolInstFilter resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetVolInstFilterResourceData(d *schema.ResourceData, m *models.VolInstFilter) {
	d.Set("app_inst_name", m.AppInstName)
	d.Set("device_name", m.DeviceName)
	d.Set("device_name_pattern", m.DeviceNamePattern)
	d.Set("label_name", m.LabelName)
	d.Set("name_pattern", m.NamePattern)
	d.Set("project_name", m.ProjectName)
	d.Set("project_name_pattern", m.ProjectNamePattern)
	d.Set("type", m.Type)
}

// Iterate through and update the VolInstFilter resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetVolInstFilterSubResourceData(m []*models.VolInstFilter) (d []*map[string]interface{}) {
	for _, VolInstFilterModel := range m {
		if VolInstFilterModel != nil {
			properties := make(map[string]interface{})
			properties["app_inst_name"] = VolInstFilterModel.AppInstName
			properties["device_name"] = VolInstFilterModel.DeviceName
			properties["device_name_pattern"] = VolInstFilterModel.DeviceNamePattern
			properties["label_name"] = VolInstFilterModel.LabelName
			properties["name_pattern"] = VolInstFilterModel.NamePattern
			properties["project_name"] = VolInstFilterModel.ProjectName
			properties["project_name_pattern"] = VolInstFilterModel.ProjectNamePattern
			properties["type"] = VolInstFilterModel.Type
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the VolInstFilter resource defined in the Terraform configuration
func VolInstFilterSchema() map[string]*schema.Schema {
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

		"label_name": {
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

		"type": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the VolInstFilter resource
func GetVolInstFilterPropertyFields() (t []string) {
	return []string{
		"app_inst_name",
		"device_name",
		"device_name_pattern",
		"label_name",
		"name_pattern",
		"project_name",
		"project_name_pattern",
		"type",
	}
}
