package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func AppInstFilterModel(d *schema.ResourceData) *models.AppInstFilter {
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
	return &models.AppInstFilter{
		AppName:            appName,
		AppType:            appType,
		DeploymentType:     deploymentType,
		DeviceName:         deviceName,
		DeviceNamePattern:  deviceNamePattern,
		NamePattern:        namePattern,
		ProjectName:        projectName,
		ProjectNamePattern: projectNamePattern,
	}
}

func AppInstFilterModelFromMap(m map[string]interface{}) *models.AppInstFilter {
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
	return &models.AppInstFilter{
		AppName:            appName,
		AppType:            appType,
		DeploymentType:     deploymentType,
		DeviceName:         deviceName,
		DeviceNamePattern:  deviceNamePattern,
		NamePattern:        namePattern,
		ProjectName:        projectName,
		ProjectNamePattern: projectNamePattern,
	}
}

// Update the underlying AppInstFilter resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetAppInstFilterResourceData(d *schema.ResourceData, m *models.AppInstFilter) {
	d.Set("app_name", m.AppName)
	d.Set("app_type", m.AppType)
	d.Set("deployment_type", m.DeploymentType)
	d.Set("device_name", m.DeviceName)
	d.Set("device_name_pattern", m.DeviceNamePattern)
	d.Set("name_pattern", m.NamePattern)
	d.Set("project_name", m.ProjectName)
	d.Set("project_name_pattern", m.ProjectNamePattern)
}

// Iterate through and update the AppInstFilter resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetAppInstFilterSubResourceData(m []*models.AppInstFilter) (d []*map[string]interface{}) {
	for _, AppInstFilterModel := range m {
		if AppInstFilterModel != nil {
			properties := make(map[string]interface{})
			properties["app_name"] = AppInstFilterModel.AppName
			properties["app_type"] = AppInstFilterModel.AppType
			properties["deployment_type"] = AppInstFilterModel.DeploymentType
			properties["device_name"] = AppInstFilterModel.DeviceName
			properties["device_name_pattern"] = AppInstFilterModel.DeviceNamePattern
			properties["name_pattern"] = AppInstFilterModel.NamePattern
			properties["project_name"] = AppInstFilterModel.ProjectName
			properties["project_name_pattern"] = AppInstFilterModel.ProjectNamePattern
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the AppInstFilter resource defined in the Terraform configuration
func AppInstFilterSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"app_name": {
			Description: `User defined name of the app instance, unique across the enterprise. Once app instance is created, name can’t be changed`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"app_type": {
			Description: `type of bundle`,
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
	}
}

// Retrieve property field names for updating the AppInstFilter resource
func GetAppInstFilterPropertyFields() (t []string) {
	return []string{
		"app_name",
		"app_type",
		"deployment_type",
		"device_name",
		"device_name_pattern",
		"name_pattern",
		"project_name",
		"project_name_pattern",
	}
}
