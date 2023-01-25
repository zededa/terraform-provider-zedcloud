package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func NetInstFilterModel(d *schema.ResourceData) *models.NetInstFilter {
	deviceName, _ := d.Get("device_name").(string)
	deviceNamePattern, _ := d.Get("device_name_pattern").(string)
	namePattern, _ := d.Get("name_pattern").(string)
	projectName, _ := d.Get("project_name").(string)
	projectNamePattern, _ := d.Get("project_name_pattern").(string)
	tags := map[string]string{}
	tagsInterface, tagsIsSet := d.GetOk("tags")
	if tagsIsSet {
		tagsMap := tagsInterface.(map[string]interface{})
		for k, v := range tagsMap {
			if v == nil {
				continue
			}
			tags[k] = v.(string)
		}
	}

	return &models.NetInstFilter{
		DeviceName:         deviceName,
		DeviceNamePattern:  deviceNamePattern,
		NamePattern:        namePattern,
		ProjectName:        projectName,
		ProjectNamePattern: projectNamePattern,
		Tags:               tags,
	}
}

func NetInstFilterModelFromMap(m map[string]interface{}) *models.NetInstFilter {
	deviceName := m["device_name"].(string)
	deviceNamePattern := m["device_name_pattern"].(string)
	namePattern := m["name_pattern"].(string)
	projectName := m["project_name"].(string)
	projectNamePattern := m["project_name_pattern"].(string)
	tags := map[string]string{}
	tagsInterface, tagsIsSet := m["tags"]
	if tagsIsSet {
		tagsMap := tagsInterface.(map[string]interface{})
		for k, v := range tagsMap {
			if v == nil {
				continue
			}
			tags[k] = v.(string)
		}
	}

	return &models.NetInstFilter{
		DeviceName:         deviceName,
		DeviceNamePattern:  deviceNamePattern,
		NamePattern:        namePattern,
		ProjectName:        projectName,
		ProjectNamePattern: projectNamePattern,
		Tags:               tags,
	}
}

// Update the underlying NetInstFilter resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetNetInstFilterResourceData(d *schema.ResourceData, m *models.NetInstFilter) {
	d.Set("device_name", m.DeviceName)
	d.Set("device_name_pattern", m.DeviceNamePattern)
	d.Set("name_pattern", m.NamePattern)
	d.Set("project_name", m.ProjectName)
	d.Set("project_name_pattern", m.ProjectNamePattern)
	d.Set("tags", m.Tags)
}

// Iterate through and update the NetInstFilter resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetNetInstFilterSubResourceData(m []*models.NetInstFilter) (d []*map[string]interface{}) {
	for _, NetInstFilterModel := range m {
		if NetInstFilterModel != nil {
			properties := make(map[string]interface{})
			properties["device_name"] = NetInstFilterModel.DeviceName
			properties["device_name_pattern"] = NetInstFilterModel.DeviceNamePattern
			properties["name_pattern"] = NetInstFilterModel.NamePattern
			properties["project_name"] = NetInstFilterModel.ProjectName
			properties["project_name_pattern"] = NetInstFilterModel.ProjectNamePattern
			properties["tags"] = NetInstFilterModel.Tags
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the NetInstFilter resource defined in the Terraform configuration
func NetInstFilterSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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

// Retrieve property field names for updating the NetInstFilter resource
func GetNetInstFilterPropertyFields() (t []string) {
	return []string{
		"device_name",
		"device_name_pattern",
		"name_pattern",
		"project_name",
		"project_name_pattern",
		"tags",
	}
}
