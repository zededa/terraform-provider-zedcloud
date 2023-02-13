package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AppInstStatusSummaryMsgModel(d *schema.ResourceData) *models.AppInstStatusSummaryMsg {
	var cpu *models.CPUSummary // CPUSummary
	CpuInterface, CpuIsSet := d.GetOk("cpu")
	if CpuIsSet && CpuInterface != nil {
		CpuMap := CpuInterface.([]interface{})
		if len(CpuMap) > 0 {
			cpu = CPUSummaryModelFromMap(CpuMap[0].(map[string]interface{}))
		}
	}
	var memory *models.MemorySummary // MemorySummary
	MemoryInterface, MemoryIsSet := d.GetOk("memory")
	if MemoryIsSet && MemoryInterface != nil {
		MemoryMap := MemoryInterface.([]interface{})
		if len(MemoryMap) > 0 {
			memory = MemorySummaryModelFromMap(MemoryMap[0].(map[string]interface{}))
		}
	}
	var storage *models.StorageSummary // StorageSummary
	StorageInterface, StorageIsSet := d.GetOk("storage")
	if StorageIsSet && StorageInterface != nil {
		StorageMap := StorageInterface.([]interface{})
		if len(StorageMap) > 0 {
			storage = StorageSummaryModelFromMap(StorageMap[0].(map[string]interface{}))
		}
	}
	var adminState *models.AdminState // AdminState
	adminStateInterface, adminStateIsSet := d.GetOk("admin_state")
	if adminStateIsSet {
		adminStateModel := adminStateInterface.(string)
		adminState = models.NewAdminState(models.AdminState(adminStateModel))
	}
	appID, _ := d.Get("app_id").(string)
	appName, _ := d.Get("app_name").(string)
	var appType *models.AppType // AppType
	appTypeInterface, appTypeIsSet := d.GetOk("app_type")
	if appTypeIsSet {
		appTypeModel := appTypeInterface.(string)
		appType = models.NewAppType(models.AppType(appTypeModel))
	}
	clusterID, _ := d.Get("cluster_id").(string)
	var deploymentType *models.DeploymentType // DeploymentType
	deploymentTypeInterface, deploymentTypeIsSet := d.GetOk("deployment_type")
	if deploymentTypeIsSet {
		deploymentTypeModel := deploymentTypeInterface.(string)
		deploymentType = models.NewDeploymentType(models.DeploymentType(deploymentTypeModel))
	}
	deviceID, _ := d.Get("device_id").(string)
	deviceName, _ := d.Get("device_name").(string)
	id, _ := d.Get("id").(string)
	var memorySummary *models.AppInstMemorySummary // AppInstMemorySummary
	memorySummaryInterface, memorySummaryIsSet := d.GetOk("memory_summary")
	if memorySummaryIsSet && memorySummaryInterface != nil {
		memorySummaryMap := memorySummaryInterface.([]interface{})
		if len(memorySummaryMap) > 0 {
			memorySummary = AppInstMemorySummaryModelFromMap(memorySummaryMap[0].(map[string]interface{}))
		}
	}
	name, _ := d.Get("name").(string)
	projectID, _ := d.Get("project_id").(string)
	projectName, _ := d.Get("project_name").(string)
	var runState *models.RunState // RunState
	runStateInterface, runStateIsSet := d.GetOk("run_state")
	if runStateIsSet {
		runStateModel := runStateInterface.(string)
		runState = models.NewRunState(models.RunState(runStateModel))
	}
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

	title, _ := d.Get("title").(string)
	return &models.AppInstStatusSummaryMsg{
		CPU:            cpu,
		Memory:         memory,
		Storage:        storage,
		AdminState:     adminState,
		AppID:          appID,
		AppName:        appName,
		AppType:        appType,
		ClusterID:      clusterID,
		DeploymentType: deploymentType,
		DeviceID:       deviceID,
		DeviceName:     deviceName,
		ID:             id,
		MemorySummary:  memorySummary,
		Name:           name,
		ProjectID:      projectID,
		ProjectName:    projectName,
		RunState:       runState,
		SwInfo:         swInfo,
		SwState:        swState,
		Tags:           tags,
		Title:          title,
	}
}

func AppInstStatusSummaryMsgModelFromMap(m map[string]interface{}) *models.AppInstStatusSummaryMsg {
	var cpu *models.CPUSummary // CPUSummary
	CpuInterface, CpuIsSet := m["cpu"]
	if CpuIsSet && CpuInterface != nil {
		CpuMap := CpuInterface.([]interface{})
		if len(CpuMap) > 0 {
			cpu = CPUSummaryModelFromMap(CpuMap[0].(map[string]interface{}))
		}
	}
	//
	var memory *models.MemorySummary // MemorySummary
	MemoryInterface, MemoryIsSet := m["memory"]
	if MemoryIsSet && MemoryInterface != nil {
		MemoryMap := MemoryInterface.([]interface{})
		if len(MemoryMap) > 0 {
			memory = MemorySummaryModelFromMap(MemoryMap[0].(map[string]interface{}))
		}
	}
	//
	var storage *models.StorageSummary // StorageSummary
	StorageInterface, StorageIsSet := m["storage"]
	if StorageIsSet && StorageInterface != nil {
		StorageMap := StorageInterface.([]interface{})
		if len(StorageMap) > 0 {
			storage = StorageSummaryModelFromMap(StorageMap[0].(map[string]interface{}))
		}
	}
	//
	var adminState *models.AdminState // AdminState
	adminStateInterface, adminStateIsSet := m["admin_state"]
	if adminStateIsSet {
		adminStateModel := adminStateInterface.(string)
		adminState = models.NewAdminState(models.AdminState(adminStateModel))
	}
	appID := m["app_id"].(string)
	appName := m["app_name"].(string)
	var appType *models.AppType // AppType
	appTypeInterface, appTypeIsSet := m["app_type"]
	if appTypeIsSet {
		appTypeModel := appTypeInterface.(string)
		appType = models.NewAppType(models.AppType(appTypeModel))
	}
	clusterID := m["cluster_id"].(string)
	var deploymentType *models.DeploymentType // DeploymentType
	deploymentTypeInterface, deploymentTypeIsSet := m["deployment_type"]
	if deploymentTypeIsSet {
		deploymentTypeModel := deploymentTypeInterface.(string)
		deploymentType = models.NewDeploymentType(models.DeploymentType(deploymentTypeModel))
	}
	deviceID := m["device_id"].(string)
	deviceName := m["device_name"].(string)
	id := m["id"].(string)
	var memorySummary *models.AppInstMemorySummary // AppInstMemorySummary
	memorySummaryInterface, memorySummaryIsSet := m["memory_summary"]
	if memorySummaryIsSet && memorySummaryInterface != nil {
		memorySummaryMap := memorySummaryInterface.([]interface{})
		if len(memorySummaryMap) > 0 {
			memorySummary = AppInstMemorySummaryModelFromMap(memorySummaryMap[0].(map[string]interface{}))
		}
	}
	//
	name := m["name"].(string)
	projectID := m["project_id"].(string)
	projectName := m["project_name"].(string)
	var runState *models.RunState // RunState
	runStateInterface, runStateIsSet := m["run_state"]
	if runStateIsSet {
		runStateModel := runStateInterface.(string)
		runState = models.NewRunState(models.RunState(runStateModel))
	}
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

	title := m["title"].(string)
	return &models.AppInstStatusSummaryMsg{
		CPU:            cpu,
		Memory:         memory,
		Storage:        storage,
		AdminState:     adminState,
		AppID:          appID,
		AppName:        appName,
		AppType:        appType,
		ClusterID:      clusterID,
		DeploymentType: deploymentType,
		DeviceID:       deviceID,
		DeviceName:     deviceName,
		ID:             id,
		MemorySummary:  memorySummary,
		Name:           name,
		ProjectID:      projectID,
		ProjectName:    projectName,
		RunState:       runState,
		SwInfo:         swInfo,
		SwState:        swState,
		Tags:           tags,
		Title:          title,
	}
}

// Update the underlying AppInstStatusSummaryMsg resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetAppInstStatusSummaryMsgResourceData(d *schema.ResourceData, m *models.AppInstStatusSummaryMsg) {
	d.Set("cpu", SetCPUSummarySubResourceData([]*models.CPUSummary{m.CPU}))
	d.Set("memory", SetMemorySummarySubResourceData([]*models.MemorySummary{m.Memory}))
	d.Set("storage", SetStorageSummarySubResourceData([]*models.StorageSummary{m.Storage}))
	d.Set("admin_state", m.AdminState)
	d.Set("app_id", m.AppID)
	d.Set("app_name", m.AppName)
	d.Set("app_type", m.AppType)
	d.Set("cluster_id", m.ClusterID)
	d.Set("deployment_type", m.DeploymentType)
	d.Set("device_id", m.DeviceID)
	d.Set("device_name", m.DeviceName)
	d.Set("id", m.ID)
	d.Set("memory_summary", SetAppInstMemorySummarySubResourceData([]*models.AppInstMemorySummary{m.MemorySummary}))
	d.Set("name", m.Name)
	d.Set("project_id", m.ProjectID)
	d.Set("project_name", m.ProjectName)
	d.Set("run_state", m.RunState)
	d.Set("sw_info", SetSWInfoSubResourceData(m.SwInfo))
	d.Set("sw_state", m.SwState)
	d.Set("tags", m.Tags)
	d.Set("title", m.Title)
}

// Iterate through and update the AppInstStatusSummaryMsg resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetAppInstStatusSummaryMsgSubResourceData(m []*models.AppInstStatusSummaryMsg) (d []*map[string]interface{}) {
	for _, AppInstStatusSummaryMsgModel := range m {
		if AppInstStatusSummaryMsgModel != nil {
			properties := make(map[string]interface{})
			properties["cpu"] = SetCPUSummarySubResourceData([]*models.CPUSummary{AppInstStatusSummaryMsgModel.CPU})
			properties["memory"] = SetMemorySummarySubResourceData([]*models.MemorySummary{AppInstStatusSummaryMsgModel.Memory})
			properties["storage"] = SetStorageSummarySubResourceData([]*models.StorageSummary{AppInstStatusSummaryMsgModel.Storage})
			properties["admin_state"] = AppInstStatusSummaryMsgModel.AdminState
			properties["app_id"] = AppInstStatusSummaryMsgModel.AppID
			properties["app_name"] = AppInstStatusSummaryMsgModel.AppName
			properties["app_type"] = AppInstStatusSummaryMsgModel.AppType
			properties["cluster_id"] = AppInstStatusSummaryMsgModel.ClusterID
			properties["deployment_type"] = AppInstStatusSummaryMsgModel.DeploymentType
			properties["device_id"] = AppInstStatusSummaryMsgModel.DeviceID
			properties["device_name"] = AppInstStatusSummaryMsgModel.DeviceName
			properties["id"] = AppInstStatusSummaryMsgModel.ID
			properties["memory_summary"] = SetAppInstMemorySummarySubResourceData([]*models.AppInstMemorySummary{AppInstStatusSummaryMsgModel.MemorySummary})
			properties["name"] = AppInstStatusSummaryMsgModel.Name
			properties["project_id"] = AppInstStatusSummaryMsgModel.ProjectID
			properties["project_name"] = AppInstStatusSummaryMsgModel.ProjectName
			properties["run_state"] = AppInstStatusSummaryMsgModel.RunState
			properties["sw_info"] = SetSWInfoSubResourceData(AppInstStatusSummaryMsgModel.SwInfo)
			properties["sw_state"] = AppInstStatusSummaryMsgModel.SwState
			properties["tags"] = AppInstStatusSummaryMsgModel.Tags
			properties["title"] = AppInstStatusSummaryMsgModel.Title
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the AppInstStatusSummaryMsg resource defined in the Terraform configuration
func AppInstStatusSummaryMsgSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cpu": {
			Description: `cpu summary`,
			Type:        schema.TypeList, //GoType: CPUSummary
			Elem: &schema.Resource{
				Schema: CPUSummarySchema(),
			},
			Optional: true,
		},

		"memory": {
			Description: `memory summary`,
			Type:        schema.TypeList, //GoType: MemorySummary
			Elem: &schema.Resource{
				Schema: MemorySummarySchema(),
			},
			Optional: true,
		},

		"storage": {
			Description: `storage summary`,
			Type:        schema.TypeList, //GoType: StorageSummary
			Elem: &schema.Resource{
				Schema: StorageSummarySchema(),
			},
			Optional: true,
		},

		"admin_state": {
			Description: `SPU details`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"app_id": {
			Description: `User defined name of the edge app, unique across the enterprise. Once edge app is created, name can’t be changed`,
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

		"cluster_id": {
			Description: `System defined universally unique clusterInstance ID, unique across the enterprise.`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"deployment_type": {
			Description: `type of deployment for the app, eg: azure, k3s, standalone`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"device_id": {
			Description: `User defined name of the device, unique across the enterprise. Once device is created, name can’t be changed`,
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

		"memory_summary": {
			Description: `memory summary`,
			Type:        schema.TypeList, //GoType: AppInstMemorySummary
			Elem: &schema.Resource{
				Schema: AppInstMemorySummarySchema(),
			},
			Optional: true,
		},

		"name": {
			Description: `User defined name of the app instance, unique across the enterprise. Once app instance is created, name can’t be changed`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"project_id": {
			Description: `User defined name of the project, unique across the enterprise. Once project is created, name can’t be changed`,
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

		"tags": {
			Description: `Tags are name/value pairs that enable you to categorize resources. Tag names are case insensitive with max_length 512 and min_length 3. Tag values are case sensitive with max_length 256 and min_length 3.`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"title": {
			Description: `User defined title of the app instance. Title can be changed at any time`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the AppInstStatusSummaryMsg resource
func GetAppInstStatusSummaryMsgPropertyFields() (t []string) {
	return []string{
		"cpu",
		"memory",
		"storage",
		"admin_state",
		"app_id",
		"app_name",
		"app_type",
		"cluster_id",
		"deployment_type",
		"device_id",
		"device_name",
		"id",
		"memory_summary",
		"name",
		"project_id",
		"project_name",
		"run_state",
		"sw_info",
		"sw_state",
		"tags",
		"title",
	}
}
