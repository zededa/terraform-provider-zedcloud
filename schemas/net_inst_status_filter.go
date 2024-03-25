package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

func NetInstStatusFilterModel(d *schema.ResourceData) *models.NetInstStatusFilter {
	deviceName, _ := d.Get("device_name").(string)
	deviceNamePattern, _ := d.Get("device_name_pattern").(string)
	var kind *models.NetworkInstanceKind // NetworkInstanceKind
	kindInterface, kindIsSet := d.GetOk("kind")
	if kindIsSet {
		kindModel := kindInterface.(string)
		kind = models.NewNetworkInstanceKind(models.NetworkInstanceKind(kindModel))
	}
	namePattern, _ := d.Get("name_pattern").(string)
	projectName, _ := d.Get("project_name").(string)
	projectNamePattern, _ := d.Get("project_name_pattern").(string)
	var runState *models.RunState // RunState
	runStateInterface, runStateIsSet := d.GetOk("run_state")
	if runStateIsSet {
		runStateModel := runStateInterface.(string)
		runState = models.NewRunState(models.RunState(runStateModel))
	}
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

	var typeVar *models.NetworkInstanceDhcpType // NetworkInstanceDhcpType
	typeInterface, typeIsSet := d.GetOk("type")
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewNetworkInstanceDhcpType(models.NetworkInstanceDhcpType(typeModel))
	}
	return &models.NetInstStatusFilter{
		DeviceName:         deviceName,
		DeviceNamePattern:  deviceNamePattern,
		Kind:               kind,
		NamePattern:        namePattern,
		ProjectName:        projectName,
		ProjectNamePattern: projectNamePattern,
		RunState:           runState,
		Tags:               tags,
		Type:               typeVar,
	}
}

func NetInstStatusFilterModelFromMap(m map[string]interface{}) *models.NetInstStatusFilter {
	deviceName := m["device_name"].(string)
	deviceNamePattern := m["device_name_pattern"].(string)
	var kind *models.NetworkInstanceKind // NetworkInstanceKind
	kindInterface, kindIsSet := m["kind"]
	if kindIsSet {
		kindModel := kindInterface.(string)
		kind = models.NewNetworkInstanceKind(models.NetworkInstanceKind(kindModel))
	}
	namePattern := m["name_pattern"].(string)
	projectName := m["project_name"].(string)
	projectNamePattern := m["project_name_pattern"].(string)
	var runState *models.RunState // RunState
	runStateInterface, runStateIsSet := m["run_state"]
	if runStateIsSet {
		runStateModel := runStateInterface.(string)
		runState = models.NewRunState(models.RunState(runStateModel))
	}
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

	var typeVar *models.NetworkInstanceDhcpType // NetworkInstanceDhcpType
	typeInterface, typeIsSet := m["type"]
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewNetworkInstanceDhcpType(models.NetworkInstanceDhcpType(typeModel))
	}
	return &models.NetInstStatusFilter{
		DeviceName:         deviceName,
		DeviceNamePattern:  deviceNamePattern,
		Kind:               kind,
		NamePattern:        namePattern,
		ProjectName:        projectName,
		ProjectNamePattern: projectNamePattern,
		RunState:           runState,
		Tags:               tags,
		Type:               typeVar,
	}
}

// Update the underlying NetInstStatusFilter resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetNetInstStatusFilterResourceData(d *schema.ResourceData, m *models.NetInstStatusFilter) {
	d.Set("device_name", m.DeviceName)
	d.Set("device_name_pattern", m.DeviceNamePattern)
	d.Set("kind", m.Kind)
	d.Set("name_pattern", m.NamePattern)
	d.Set("project_name", m.ProjectName)
	d.Set("project_name_pattern", m.ProjectNamePattern)
	d.Set("run_state", m.RunState)
	d.Set("tags", m.Tags)
	d.Set("type", m.Type)
}

// Iterate through and update the NetInstStatusFilter resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetNetInstStatusFilterSubResourceData(m []*models.NetInstStatusFilter) (d []*map[string]interface{}) {
	for _, NetInstStatusFilterModel := range m {
		if NetInstStatusFilterModel != nil {
			properties := make(map[string]interface{})
			properties["device_name"] = NetInstStatusFilterModel.DeviceName
			properties["device_name_pattern"] = NetInstStatusFilterModel.DeviceNamePattern
			properties["kind"] = NetInstStatusFilterModel.Kind
			properties["name_pattern"] = NetInstStatusFilterModel.NamePattern
			properties["project_name"] = NetInstStatusFilterModel.ProjectName
			properties["project_name_pattern"] = NetInstStatusFilterModel.ProjectNamePattern
			properties["run_state"] = NetInstStatusFilterModel.RunState
			properties["tags"] = NetInstStatusFilterModel.Tags
			properties["type"] = NetInstStatusFilterModel.Type
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the NetInstStatusFilter resource defined in the Terraform configuration
func NetInstStatusFilterSchema() map[string]*schema.Schema {
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

		"kind": {
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

		"tags": {
			Description: ``,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"type": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the NetInstStatusFilter resource
func GetNetInstStatusFilterPropertyFields() (t []string) {
	return []string{
		"device_name",
		"device_name_pattern",
		"kind",
		"name_pattern",
		"project_name",
		"project_name_pattern",
		"run_state",
		"tags",
		"type",
	}
}
