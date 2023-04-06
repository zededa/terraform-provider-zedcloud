package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AppInstConfigStatusModel(d *schema.ResourceData) *models.AppInstConfigStatus {
	var storage *models.StorageSummary // StorageSummary
	StorageInterface, StorageIsSet := d.GetOk("storage")
	if StorageIsSet && StorageInterface != nil {
		StorageMap := StorageInterface.([]interface{})
		if len(StorageMap) > 0 {
			storage = StorageSummaryModelFromMap(StorageMap[0].(map[string]interface{}))
		}
	}
	appID, _ := d.Get("app_id").(string)
	appName, _ := d.Get("app_name").(string)
	var appType *models.AppType // AppType
	appTypeInterface, appTypeIsSet := d.GetOk("app_type")
	if appTypeIsSet {
		appTypeModel := appTypeInterface.(string)
		appType = models.NewAppType(models.AppType(appTypeModel))
	}
	cPUUtilization, _ := d.Get("cpu_utilization").(float64)
	deviceID, _ := d.Get("device_id").(string)
	deviceName, _ := d.Get("device_name").(string)
	id, _ := d.Get("id").(string)
	var manifestInfo *models.ManifestInfo // ManifestInfo
	manifestInfoInterface, manifestInfoIsSet := d.GetOk("manifest_info")
	if manifestInfoIsSet && manifestInfoInterface != nil {
		manifestInfoMap := manifestInfoInterface.([]interface{})
		if len(manifestInfoMap) > 0 {
			manifestInfo = ManifestInfoModelFromMap(manifestInfoMap[0].(map[string]interface{}))
		}
	}
	memoryUtilization, _ := d.Get("memory_utilization").(float64)
	name, _ := d.Get("name").(string)
	projectID, _ := d.Get("project_id").(string)
	projectName, _ := d.Get("project_name").(string)
	var runState *models.RunState // RunState
	runStateInterface, runStateIsSet := d.GetOk("run_state")
	if runStateIsSet {
		runStateModel := runStateInterface.(string)
		runState = models.NewRunState(models.RunState(runStateModel))
	}
	storageUtilization, _ := d.Get("storage_utilization").(float64)
	var swInfo []*models.SWInfo // []*SWInfo
	swInfoInterface, swInfoIsSet := d.GetOk("sw_info")
	if swInfoIsSet {
		var items []interface{}
		if listItems, isList := swInfoInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = swInfoInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := SWInfoModelFromMap(v.(map[string]interface{}))
			swInfo = append(swInfo, m)
		}
	}
	var swState *models.SWState // SWState
	swStateInterface, swStateIsSet := d.GetOk("sw_state")
	if swStateIsSet {
		swStateModel := swStateInterface.(string)
		swState = models.NewSWState(models.SWState(swStateModel))
	}
	userDefinedVersion, _ := d.Get("user_defined_version").(string)
	return &models.AppInstConfigStatus{
		Storage:            storage,
		AppID:              appID,
		AppName:            appName,
		AppType:            appType,
		CPUUtilization:     cPUUtilization,
		DeviceID:           deviceID,
		DeviceName:         deviceName,
		ID:                 id,
		ManifestInfo:       manifestInfo,
		MemoryUtilization:  memoryUtilization,
		Name:               name,
		ProjectID:          projectID,
		ProjectName:        projectName,
		RunState:           runState,
		StorageUtilization: storageUtilization,
		SwInfo:             swInfo,
		SwState:            swState,
		UserDefinedVersion: userDefinedVersion,
	}
}

func AppInstConfigStatusModelFromMap(m map[string]interface{}) *models.AppInstConfigStatus {
	var storage *models.StorageSummary // StorageSummary
	StorageInterface, StorageIsSet := m["storage"]
	if StorageIsSet && StorageInterface != nil {
		StorageMap := StorageInterface.([]interface{})
		if len(StorageMap) > 0 {
			storage = StorageSummaryModelFromMap(StorageMap[0].(map[string]interface{}))
		}
	}
	//
	appID := m["app_id"].(string)
	appName := m["app_name"].(string)
	var appType *models.AppType // AppType
	appTypeInterface, appTypeIsSet := m["app_type"]
	if appTypeIsSet {
		appTypeModel := appTypeInterface.(string)
		appType = models.NewAppType(models.AppType(appTypeModel))
	}
	cPUUtilization := m["cpu_utilization"].(float64)
	deviceID := m["device_id"].(string)
	deviceName := m["device_name"].(string)
	id := m["id"].(string)
	var manifestInfo *models.ManifestInfo // ManifestInfo
	manifestInfoInterface, manifestInfoIsSet := m["manifest_info"]
	if manifestInfoIsSet && manifestInfoInterface != nil {
		manifestInfoMap := manifestInfoInterface.([]interface{})
		if len(manifestInfoMap) > 0 {
			manifestInfo = ManifestInfoModelFromMap(manifestInfoMap[0].(map[string]interface{}))
		}
	}
	//
	memoryUtilization := m["memory_utilization"].(float64)
	name := m["name"].(string)
	projectID := m["project_id"].(string)
	projectName := m["project_name"].(string)
	var runState *models.RunState // RunState
	runStateInterface, runStateIsSet := m["run_state"]
	if runStateIsSet {
		runStateModel := runStateInterface.(string)
		runState = models.NewRunState(models.RunState(runStateModel))
	}
	storageUtilization := m["storage_utilization"].(float64)
	var swInfo []*models.SWInfo // []*SWInfo
	swInfoInterface, swInfoIsSet := m["sw_info"]
	if swInfoIsSet {
		var items []interface{}
		if listItems, isList := swInfoInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = swInfoInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := SWInfoModelFromMap(v.(map[string]interface{}))
			swInfo = append(swInfo, m)
		}
	}
	var swState *models.SWState // SWState
	swStateInterface, swStateIsSet := m["sw_state"]
	if swStateIsSet {
		swStateModel := swStateInterface.(string)
		swState = models.NewSWState(models.SWState(swStateModel))
	}
	userDefinedVersion := m["user_defined_version"].(string)
	return &models.AppInstConfigStatus{
		Storage:            storage,
		AppID:              appID,
		AppName:            appName,
		AppType:            appType,
		CPUUtilization:     cPUUtilization,
		DeviceID:           deviceID,
		DeviceName:         deviceName,
		ID:                 id,
		ManifestInfo:       manifestInfo,
		MemoryUtilization:  memoryUtilization,
		Name:               name,
		ProjectID:          projectID,
		ProjectName:        projectName,
		RunState:           runState,
		StorageUtilization: storageUtilization,
		SwInfo:             swInfo,
		SwState:            swState,
		UserDefinedVersion: userDefinedVersion,
	}
}

// Update the underlying AppInstConfigStatus resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetAppInstConfigStatusResourceData(d *schema.ResourceData, m *models.AppInstConfigStatus) {
	d.Set("storage", SetStorageSummarySubResourceData([]*models.StorageSummary{m.Storage}))
	d.Set("app_id", m.AppID)
	d.Set("app_name", m.AppName)
	d.Set("app_type", m.AppType)
	d.Set("cpu_utilization", m.CPUUtilization)
	d.Set("device_id", m.DeviceID)
	d.Set("device_name", m.DeviceName)
	d.Set("id", m.ID)
	d.Set("manifest_info", SetManifestInfoSubResourceData([]*models.ManifestInfo{m.ManifestInfo}))
	d.Set("memory_utilization", m.MemoryUtilization)
	d.Set("name", m.Name)
	d.Set("project_id", m.ProjectID)
	d.Set("project_name", m.ProjectName)
	d.Set("run_state", m.RunState)
	d.Set("storage_utilization", m.StorageUtilization)
	d.Set("sw_info", SetSWInfoSubResourceData(m.SwInfo))
	d.Set("sw_state", m.SwState)
	d.Set("user_defined_version", m.UserDefinedVersion)
}

// Iterate through and update the AppInstConfigStatus resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetAppInstConfigStatusSubResourceData(m []*models.AppInstConfigStatus) (d []*map[string]interface{}) {
	for _, AppInstConfigStatusModel := range m {
		if AppInstConfigStatusModel != nil {
			properties := make(map[string]interface{})
			properties["storage"] = SetStorageSummarySubResourceData([]*models.StorageSummary{AppInstConfigStatusModel.Storage})
			properties["app_id"] = AppInstConfigStatusModel.AppID
			properties["app_name"] = AppInstConfigStatusModel.AppName
			properties["app_type"] = AppInstConfigStatusModel.AppType
			properties["cpu_utilization"] = AppInstConfigStatusModel.CPUUtilization
			properties["device_id"] = AppInstConfigStatusModel.DeviceID
			properties["device_name"] = AppInstConfigStatusModel.DeviceName
			properties["id"] = AppInstConfigStatusModel.ID
			properties["manifest_info"] = SetManifestInfoSubResourceData([]*models.ManifestInfo{AppInstConfigStatusModel.ManifestInfo})
			properties["memory_utilization"] = AppInstConfigStatusModel.MemoryUtilization
			properties["name"] = AppInstConfigStatusModel.Name
			properties["project_id"] = AppInstConfigStatusModel.ProjectID
			properties["project_name"] = AppInstConfigStatusModel.ProjectName
			properties["run_state"] = AppInstConfigStatusModel.RunState
			properties["storage_utilization"] = AppInstConfigStatusModel.StorageUtilization
			properties["sw_info"] = SetSWInfoSubResourceData(AppInstConfigStatusModel.SwInfo)
			properties["sw_state"] = AppInstConfigStatusModel.SwState
			properties["user_defined_version"] = AppInstConfigStatusModel.UserDefinedVersion
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the AppInstConfigStatus resource defined in the Terraform configuration
func AppInstConfigStatusSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"storage": {
			Description: `storage summary`,
			Type:        schema.TypeList, //GoType: StorageSummary
			Elem: &schema.Resource{
				Schema: StorageSummarySchema(),
			},
			Optional: true,
		},

		"app_id": {
			Description: `System generated identifier for the app bundle`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"app_name": {
			Description: `User defined name of the edge app, unique across the enterprise. Once edge app is created, name can’t be changed`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"app_type": {
			Description: `type of app`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"cpu_utilization": {
			Description: `cpu utilization`,
			Type:        schema.TypeFloat,
			Optional:    true,
		},

		"device_id": {
			Description: `System generated identifier for the device`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"device_name": {
			Description: `User defined name of the device, unique across the enterprise. Once device is created, name can’t be changed`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"id": {
			Description: `System defined universally unique Id of the app instance`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"manifest_info": {
			Description: `app manifest Info`,
			Type:        schema.TypeList, //GoType: ManifestInfo
			Elem: &schema.Resource{
				Schema: ManifestInfoSchema(),
			},
			Optional: true,
		},

		"memory_utilization": {
			Description: `memory utilization`,
			Type:        schema.TypeFloat,
			Optional:    true,
		},

		"name": {
			Description: `User defined name of the app instance, unique across the enterprise. Once app instance is created, name can’t be changed`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"project_id": {
			Description: `System generated identifier for the project`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"project_name": {
			Description: `User defined name of the project, unique across the enterprise. Once project is created, name can’t be changed`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"run_state": {
			Description: `operation status`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"storage_utilization": {
			Description: `storage utilization`,
			Type:        schema.TypeFloat,
			Optional:    true,
		},

		"sw_info": {
			Description: `Software details`,
			Type:        schema.TypeList, //GoType: []*SWInfo
			Elem: &schema.Resource{
				Schema: SWInfoSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"sw_state": {
			Description: `sotware state`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"user_defined_version": {
			Description: `Instance version tells which edge app does this instance is running`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the AppInstConfigStatus resource
func GetAppInstConfigStatusPropertyFields() (t []string) {
	return []string{
		"storage",
		"app_id",
		"app_name",
		"app_type",
		"cpu_utilization",
		"device_id",
		"device_name",
		"id",
		"manifest_info",
		"memory_utilization",
		"name",
		"project_id",
		"project_name",
		"run_state",
		"storage_utilization",
		"sw_info",
		"sw_state",
		"user_defined_version",
	}
}
