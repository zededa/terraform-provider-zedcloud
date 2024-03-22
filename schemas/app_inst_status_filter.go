package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

func AppInstStatusFilterModel(d *schema.ResourceData) *models.AppInstStatusFilter {
	appName, _ := d.Get("app_name").(string)
	var appType *models.AppType // AppType
	appTypeInterface, appTypeIsSet := d.GetOk("app_type")
	if appTypeIsSet {
		appTypeModel := appTypeInterface.(string)
		appType = models.NewAppType(models.AppType(appTypeModel))
	}
	var deploymentType *models.DeploymentType // DeploymentType
	deploymentTypeInterface, deploymentTypeIsSet := d.GetOk("deployment_type")
	if deploymentTypeIsSet {
		deploymentTypeModel := deploymentTypeInterface.(string)
		deploymentType = models.NewDeploymentType(models.DeploymentType(deploymentTypeModel))
	}
	deviceName, _ := d.Get("device_name").(string)
	deviceNamePattern, _ := d.Get("device_name_pattern").(string)
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

	return &models.AppInstStatusFilter{
		AppName:            appName,
		AppType:            appType,
		DeploymentType:     deploymentType,
		DeviceName:         deviceName,
		DeviceNamePattern:  deviceNamePattern,
		NamePattern:        namePattern,
		ProjectName:        projectName,
		ProjectNamePattern: projectNamePattern,
		RunState:           runState,
		Tags:               tags,
	}
}

func AppInstStatusFilterModelFromMap(m map[string]interface{}) *models.AppInstStatusFilter {
	appName := m["app_name"].(string)
	var appType *models.AppType // AppType
	appTypeInterface, appTypeIsSet := m["app_type"]
	if appTypeIsSet {
		appTypeModel := appTypeInterface.(string)
		appType = models.NewAppType(models.AppType(appTypeModel))
	}
	var deploymentType *models.DeploymentType // DeploymentType
	deploymentTypeInterface, deploymentTypeIsSet := m["deployment_type"]
	if deploymentTypeIsSet {
		deploymentTypeModel := deploymentTypeInterface.(string)
		deploymentType = models.NewDeploymentType(models.DeploymentType(deploymentTypeModel))
	}
	deviceName := m["device_name"].(string)
	deviceNamePattern := m["device_name_pattern"].(string)
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

	return &models.AppInstStatusFilter{
		AppName:            appName,
		AppType:            appType,
		DeploymentType:     deploymentType,
		DeviceName:         deviceName,
		DeviceNamePattern:  deviceNamePattern,
		NamePattern:        namePattern,
		ProjectName:        projectName,
		ProjectNamePattern: projectNamePattern,
		RunState:           runState,
		Tags:               tags,
	}
}

// Update the underlying AppInstStatusFilter resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetAppInstStatusFilterResourceData(d *schema.ResourceData, m *models.AppInstStatusFilter) {
	d.Set("app_name", m.AppName)
	d.Set("app_type", m.AppType)
	d.Set("deployment_type", m.DeploymentType)
	d.Set("device_name", m.DeviceName)
	d.Set("device_name_pattern", m.DeviceNamePattern)
	d.Set("name_pattern", m.NamePattern)
	d.Set("project_name", m.ProjectName)
	d.Set("project_name_pattern", m.ProjectNamePattern)
	d.Set("run_state", m.RunState)
	d.Set("tags", m.Tags)
}

// Iterate through and update the AppInstStatusFilter resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetAppInstStatusFilterSubResourceData(m []*models.AppInstStatusFilter) (d []*map[string]interface{}) {
	for _, AppInstStatusFilterModel := range m {
		if AppInstStatusFilterModel != nil {
			properties := make(map[string]interface{})
			properties["app_name"] = AppInstStatusFilterModel.AppName
			properties["app_type"] = AppInstStatusFilterModel.AppType
			properties["deployment_type"] = AppInstStatusFilterModel.DeploymentType
			properties["device_name"] = AppInstStatusFilterModel.DeviceName
			properties["device_name_pattern"] = AppInstStatusFilterModel.DeviceNamePattern
			properties["name_pattern"] = AppInstStatusFilterModel.NamePattern
			properties["project_name"] = AppInstStatusFilterModel.ProjectName
			properties["project_name_pattern"] = AppInstStatusFilterModel.ProjectNamePattern
			properties["run_state"] = AppInstStatusFilterModel.RunState
			properties["tags"] = AppInstStatusFilterModel.Tags
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the AppInstStatusFilter resource defined in the Terraform configuration
func AppInstStatusFilterSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"app_name": {
			Description: `User defined name of the app bundle, unique across the enterprise. Once app bundle is created, name can’t be changed`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"app_type": {
			Description: `type of app`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"deployment_type": {
			Description: `type of deployment for the app, eg: azure, k3s, standalone`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"device_name": {
			Description: `User defined name of the device, unique across the enterprise. Once device is created, name can’t be changed`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"device_name_pattern": {
			Description: `device name pattern`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"name_pattern": {
			Description: `name pattern`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"project_name": {
			Description: `User defined name of the project, unique across the enterprise. Once project is created, name can’t be changed`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"project_name_pattern": {
			Description: `project name pattern`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"run_state": {
			Description: `aperation status`,
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

// Retrieve property field names for updating the AppInstStatusFilter resource
func GetAppInstStatusFilterPropertyFields() (t []string) {
	return []string{
		"app_name",
		"app_type",
		"deployment_type",
		"device_name",
		"device_name_pattern",
		"name_pattern",
		"project_name",
		"project_name_pattern",
		"run_state",
		"tags",
	}
}
