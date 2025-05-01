package schemas

import (
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func AppInstStatusMsgModel(d *schema.ResourceData) *models.AppInstStatusMsg {
	var cpu *models.CPUSummary // CPUSummary
	CpuInterface, CpuIsSet := d.GetOk("cpu")
	if CpuIsSet && CpuInterface != nil {
		CpuMap := CpuInterface.([]interface{})
		if len(CpuMap) > 0 {
			cpu = CPUSummaryModelFromMap(CpuMap[0].(map[string]interface{}))
		}
	}
	var ioStatusList []*models.IoBundleStatus // []*IoBundleStatus
	IoStatusListInterface, IoStatusListIsSet := d.GetOk("io_status_list")
	if IoStatusListIsSet {
		var items []interface{}
		if listItems, isList := IoStatusListInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = IoStatusListInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := IoBundleStatusModelFromMap(v.(map[string]interface{}))
			ioStatusList = append(ioStatusList, m)
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
	var appStatusFromTPController *models.AppStatusFromTPController // AppStatusFromTPController
	appStatusFromTPControllerInterface, appStatusFromTPControllerIsSet := d.GetOk("app_status_from_t_p_controller")
	if appStatusFromTPControllerIsSet && appStatusFromTPControllerInterface != nil {
		appStatusFromTPControllerMap := appStatusFromTPControllerInterface.([]interface{})
		if len(appStatusFromTPControllerMap) > 0 {
			appStatusFromTPController = AppStatusFromTPControllerModelFromMap(appStatusFromTPControllerMap[0].(map[string]interface{}))
		}
	}
	var appType *models.AppType // AppType
	appTypeInterface, appTypeIsSet := d.GetOk("app_type")
	if appTypeIsSet {
		appTypeModel := appTypeInterface.(string)
		appType = models.NewAppType(models.AppType(appTypeModel))
	}
	bootTime, _ := d.Get("boot_time").(strfmt.DateTime)
	clusterID, _ := d.Get("cluster_id").(string)
	var containerStatusList []*models.ComposeContainerStatus // []*ComposeContainerStatus
	containerStatusListInterface, containerStatusListIsSet := d.GetOk("container_status_list")
	if containerStatusListIsSet {
		var items []interface{}
		if listItems, isList := containerStatusListInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = containerStatusListInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := ComposeContainerStatusModelFromMap(v.(map[string]interface{}))
			containerStatusList = append(containerStatusList, m)
		}
	}
	var deploymentType *models.DeploymentType // DeploymentType
	deploymentTypeInterface, deploymentTypeIsSet := d.GetOk("deployment_type")
	if deploymentTypeIsSet {
		deploymentTypeModel := deploymentTypeInterface.(string)
		deploymentType = models.NewDeploymentType(models.DeploymentType(deploymentTypeModel))
	}
	deviceID, _ := d.Get("device_id").(string)
	var errInfo []*models.DeviceError // []*DeviceError
	errInfoInterface, errInfoIsSet := d.GetOk("err_info")
	if errInfoIsSet {
		var items []interface{}
		if listItems, isList := errInfoInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = errInfoInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := DeviceErrorModelFromMap(v.(map[string]interface{}))
			errInfo = append(errInfo, m)
		}
	}
	id, _ := d.Get("id").(string)
	lastUpdate, _ := d.Get("last_update").(strfmt.DateTime)
	var memorySummary *models.AppInstMemorySummary // AppInstMemorySummary
	memorySummaryInterface, memorySummaryIsSet := d.GetOk("memory_summary")
	if memorySummaryIsSet && memorySummaryInterface != nil {
		memorySummaryMap := memorySummaryInterface.([]interface{})
		if len(memorySummaryMap) > 0 {
			memorySummary = AppInstMemorySummaryModelFromMap(memorySummaryMap[0].(map[string]interface{}))
		}
	}
	name, _ := d.Get("name").(string)
	var netCounterList []*models.NetworkCounters // []*NetworkCounters
	netCounterListInterface, netCounterListIsSet := d.GetOk("net_counter_list")
	if netCounterListIsSet {
		var items []interface{}
		if listItems, isList := netCounterListInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = netCounterListInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := NetworkCountersModelFromMap(v.(map[string]interface{}))
			netCounterList = append(netCounterList, m)
		}
	}
	var netStatusList []*models.NetworkStatus // []*NetworkStatus
	netStatusListInterface, netStatusListIsSet := d.GetOk("net_status_list")
	if netStatusListIsSet {
		var items []interface{}
		if listItems, isList := netStatusListInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = netStatusListInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := NetworkStatusModelFromMap(v.(map[string]interface{}))
			netStatusList = append(netStatusList, m)
		}
	}
	projectID, _ := d.Get("project_id").(string)
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
	upTime, _ := d.Get("up_time").(strfmt.DateTime)
	var volumeRefs []string
	volumeRefsInterface, volumeRefsIsSet := d.GetOk("volumeRefs")
	if volumeRefsIsSet {
		var items []interface{}
		if listItems, isList := volumeRefsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = volumeRefsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			volumeRefs = append(volumeRefs, v.(string))
		}
	}
	var zpoolMetrics *models.StorageDeviceMetrics // StorageDeviceMetrics
	zpoolMetricsInterface, zpoolMetricsIsSet := d.GetOk("zpool_metrics")
	if zpoolMetricsIsSet && zpoolMetricsInterface != nil {
		zpoolMetricsMap := zpoolMetricsInterface.([]interface{})
		if len(zpoolMetricsMap) > 0 {
			zpoolMetrics = StorageDeviceMetricsModelFromMap(zpoolMetricsMap[0].(map[string]interface{}))
		}
	}
	return &models.AppInstStatusMsg{
		CPU:                       cpu,
		IoStatusList:              ioStatusList,
		Memory:                    memory,
		Storage:                   storage,
		AdminState:                adminState,
		AppID:                     appID,
		AppStatusFromTPController: appStatusFromTPController,
		AppType:                   appType,
		BootTime:                  bootTime,
		ClusterID:                 clusterID,
		ContainerStatusList:       containerStatusList,
		DeploymentType:            deploymentType,
		DeviceID:                  deviceID,
		ErrInfo:                   errInfo,
		ID:                        id,
		LastUpdate:                lastUpdate,
		MemorySummary:             memorySummary,
		Name:                      name,
		NetCounterList:            netCounterList,
		NetStatusList:             netStatusList,
		ProjectID:                 projectID,
		RunState:                  runState,
		SwInfo:                    swInfo,
		SwState:                   swState,
		Tags:                      tags,
		Title:                     title,
		UpTime:                    upTime,
		VolumeRefs:                volumeRefs,
		ZpoolMetrics:              zpoolMetrics,
	}
}

func AppInstStatusMsgModelFromMap(m map[string]interface{}) *models.AppInstStatusMsg {
	var cpu *models.CPUSummary // CPUSummary
	CpuInterface, CpuIsSet := m["cpu"]
	if CpuIsSet && CpuInterface != nil {
		CpuMap := CpuInterface.([]interface{})
		if len(CpuMap) > 0 {
			cpu = CPUSummaryModelFromMap(CpuMap[0].(map[string]interface{}))
		}
	}
	//
	var ioStatusList []*models.IoBundleStatus // []*IoBundleStatus
	IoStatusListInterface, IoStatusListIsSet := m["io_status_list"]
	if IoStatusListIsSet {
		var items []interface{}
		if listItems, isList := IoStatusListInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = IoStatusListInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := IoBundleStatusModelFromMap(v.(map[string]interface{}))
			ioStatusList = append(ioStatusList, m)
		}
	}
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
	var appStatusFromTPController *models.AppStatusFromTPController // AppStatusFromTPController
	appStatusFromTPControllerInterface, appStatusFromTPControllerIsSet := m["app_status_from_t_p_controller"]
	if appStatusFromTPControllerIsSet && appStatusFromTPControllerInterface != nil {
		appStatusFromTPControllerMap := appStatusFromTPControllerInterface.([]interface{})
		if len(appStatusFromTPControllerMap) > 0 {
			appStatusFromTPController = AppStatusFromTPControllerModelFromMap(appStatusFromTPControllerMap[0].(map[string]interface{}))
		}
	}
	//
	var appType *models.AppType // AppType
	appTypeInterface, appTypeIsSet := m["app_type"]
	if appTypeIsSet {
		appTypeModel := appTypeInterface.(string)
		appType = models.NewAppType(models.AppType(appTypeModel))
	}
	bootTime := m["boot_time"].(strfmt.DateTime)
	clusterID := m["cluster_id"].(string)
	var deploymentType *models.DeploymentType // DeploymentType
	deploymentTypeInterface, deploymentTypeIsSet := m["deployment_type"]
	if deploymentTypeIsSet {
		deploymentTypeModel := deploymentTypeInterface.(string)
		deploymentType = models.NewDeploymentType(models.DeploymentType(deploymentTypeModel))
	}
	var containerStatusList []*models.ComposeContainerStatus // []*ComposeContainerStatus
	containerStatusListInterface, containerStatusListIsSet := m["container_status_list"]
	if containerStatusListIsSet {
		var items []interface{}
		if listItems, isList := containerStatusListInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = containerStatusListInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := ComposeContainerStatusModelFromMap(v.(map[string]interface{}))
			containerStatusList = append(containerStatusList, m)
		}
	}
	deviceID := m["device_id"].(string)
	var errInfo []*models.DeviceError // []*DeviceError
	errInfoInterface, errInfoIsSet := m["err_info"]
	if errInfoIsSet {
		var items []interface{}
		if listItems, isList := errInfoInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = errInfoInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := DeviceErrorModelFromMap(v.(map[string]interface{}))
			errInfo = append(errInfo, m)
		}
	}
	id := m["id"].(string)
	lastUpdate := m["last_update"].(strfmt.DateTime)
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
	var netCounterList []*models.NetworkCounters // []*NetworkCounters
	netCounterListInterface, netCounterListIsSet := m["net_counter_list"]
	if netCounterListIsSet {
		var items []interface{}
		if listItems, isList := netCounterListInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = netCounterListInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := NetworkCountersModelFromMap(v.(map[string]interface{}))
			netCounterList = append(netCounterList, m)
		}
	}
	var netStatusList []*models.NetworkStatus // []*NetworkStatus
	netStatusListInterface, netStatusListIsSet := m["net_status_list"]
	if netStatusListIsSet {
		var items []interface{}
		if listItems, isList := netStatusListInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = netStatusListInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := NetworkStatusModelFromMap(v.(map[string]interface{}))
			netStatusList = append(netStatusList, m)
		}
	}
	projectID := m["project_id"].(string)
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
	upTime := m["up_time"].(strfmt.DateTime)
	var volumeRefs []string
	volumeRefsInterface, volumeRefsIsSet := m["volumeRefs"]
	if volumeRefsIsSet {
		var items []interface{}
		if listItems, isList := volumeRefsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = volumeRefsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			volumeRefs = append(volumeRefs, v.(string))
		}
	}
	var zpoolMetrics *models.StorageDeviceMetrics // StorageDeviceMetrics
	zpoolMetricsInterface, zpoolMetricsIsSet := m["zpool_metrics"]
	if zpoolMetricsIsSet && zpoolMetricsInterface != nil {
		zpoolMetricsMap := zpoolMetricsInterface.([]interface{})
		if len(zpoolMetricsMap) > 0 {
			zpoolMetrics = StorageDeviceMetricsModelFromMap(zpoolMetricsMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.AppInstStatusMsg{
		CPU:                       cpu,
		IoStatusList:              ioStatusList,
		Memory:                    memory,
		Storage:                   storage,
		AdminState:                adminState,
		AppID:                     appID,
		AppStatusFromTPController: appStatusFromTPController,
		AppType:                   appType,
		BootTime:                  bootTime,
		ClusterID:                 clusterID,
		ContainerStatusList:       containerStatusList,
		DeploymentType:            deploymentType,
		DeviceID:                  deviceID,
		ErrInfo:                   errInfo,
		ID:                        id,
		LastUpdate:                lastUpdate,
		MemorySummary:             memorySummary,
		Name:                      name,
		NetCounterList:            netCounterList,
		NetStatusList:             netStatusList,
		ProjectID:                 projectID,
		RunState:                  runState,
		SwInfo:                    swInfo,
		SwState:                   swState,
		Tags:                      tags,
		Title:                     title,
		UpTime:                    upTime,
		VolumeRefs:                volumeRefs,
		ZpoolMetrics:              zpoolMetrics,
	}
}

// Update the underlying AppInstStatusMsg resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetAppInstStatusMsgResourceData(d *schema.ResourceData, m *models.AppInstStatusMsg) {
	d.Set("cpu", SetCPUSummarySubResourceData([]*models.CPUSummary{m.CPU}))
	d.Set("io_status_list", SetIoBundleStatusSubResourceData(m.IoStatusList))
	d.Set("memory", SetMemorySummarySubResourceData([]*models.MemorySummary{m.Memory}))
	d.Set("storage", SetStorageSummarySubResourceData([]*models.StorageSummary{m.Storage}))
	d.Set("admin_state", m.AdminState)
	d.Set("app_id", m.AppID)
	d.Set("app_status_from_t_p_controller", SetAppStatusFromTPControllerSubResourceData([]*models.AppStatusFromTPController{m.AppStatusFromTPController}))
	d.Set("app_type", m.AppType)
	d.Set("boot_time", m.BootTime)
	d.Set("cluster_id", m.ClusterID)
	d.Set("container_status_list", SetComposeContainerStatusSubResourceData(m.ContainerStatusList))
	d.Set("deployment_type", m.DeploymentType)
	d.Set("device_id", m.DeviceID)
	d.Set("err_info", SetDeviceErrorSubResourceData(m.ErrInfo))
	d.Set("id", m.ID)
	d.Set("last_update", m.LastUpdate)
	d.Set("memory_summary", SetAppInstMemorySummarySubResourceData([]*models.AppInstMemorySummary{m.MemorySummary}))
	d.Set("name", m.Name)
	d.Set("net_counter_list", SetNetworkCountersSubResourceData(m.NetCounterList))
	d.Set("net_status_list", SetNetworkStatusSubResourceData(m.NetStatusList))
	d.Set("project_id", m.ProjectID)
	d.Set("run_state", m.RunState)
	d.Set("sw_info", SetSWInfoSubResourceData(m.SwInfo))
	d.Set("sw_state", m.SwState)
	d.Set("tags", m.Tags)
	d.Set("title", m.Title)
	d.Set("up_time", m.UpTime)
	d.Set("volume_refs", m.VolumeRefs)
	d.Set("zpool_metrics", SetStorageDeviceMetricsSubResourceData([]*models.StorageDeviceMetrics{m.ZpoolMetrics}))
}

// Iterate through and update the AppInstStatusMsg resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetAppInstStatusMsgSubResourceData(m []*models.AppInstStatusMsg) (d []*map[string]interface{}) {
	for _, AppInstStatusMsgModel := range m {
		if AppInstStatusMsgModel != nil {
			properties := make(map[string]interface{})
			properties["cpu"] = SetCPUSummarySubResourceData([]*models.CPUSummary{AppInstStatusMsgModel.CPU})
			properties["io_status_list"] = SetIoBundleStatusSubResourceData(AppInstStatusMsgModel.IoStatusList)
			properties["memory"] = SetMemorySummarySubResourceData([]*models.MemorySummary{AppInstStatusMsgModel.Memory})
			properties["storage"] = SetStorageSummarySubResourceData([]*models.StorageSummary{AppInstStatusMsgModel.Storage})
			properties["admin_state"] = AppInstStatusMsgModel.AdminState
			properties["app_id"] = AppInstStatusMsgModel.AppID
			properties["app_status_from_t_p_controller"] = SetAppStatusFromTPControllerSubResourceData([]*models.AppStatusFromTPController{AppInstStatusMsgModel.AppStatusFromTPController})
			properties["app_type"] = AppInstStatusMsgModel.AppType
			properties["boot_time"] = AppInstStatusMsgModel.BootTime.String()
			properties["cluster_id"] = AppInstStatusMsgModel.ClusterID
			properties["container_status_list"] = SetComposeContainerStatusSubResourceData(AppInstStatusMsgModel.ContainerStatusList)
			properties["deployment_type"] = AppInstStatusMsgModel.DeploymentType
			properties["device_id"] = AppInstStatusMsgModel.DeviceID
			properties["err_info"] = SetDeviceErrorSubResourceData(AppInstStatusMsgModel.ErrInfo)
			properties["id"] = AppInstStatusMsgModel.ID
			properties["last_update"] = AppInstStatusMsgModel.LastUpdate.String()
			properties["memory_summary"] = SetAppInstMemorySummarySubResourceData([]*models.AppInstMemorySummary{AppInstStatusMsgModel.MemorySummary})
			properties["name"] = AppInstStatusMsgModel.Name
			properties["net_counter_list"] = SetNetworkCountersSubResourceData(AppInstStatusMsgModel.NetCounterList)
			properties["net_status_list"] = SetNetworkStatusSubResourceData(AppInstStatusMsgModel.NetStatusList)
			properties["project_id"] = AppInstStatusMsgModel.ProjectID
			properties["run_state"] = AppInstStatusMsgModel.RunState
			properties["sw_info"] = SetSWInfoSubResourceData(AppInstStatusMsgModel.SwInfo)
			properties["sw_state"] = AppInstStatusMsgModel.SwState
			properties["tags"] = AppInstStatusMsgModel.Tags
			properties["title"] = AppInstStatusMsgModel.Title
			properties["up_time"] = AppInstStatusMsgModel.UpTime.String()
			properties["volume_refs"] = AppInstStatusMsgModel.VolumeRefs
			properties["zpool_metrics"] = SetStorageDeviceMetricsSubResourceData([]*models.StorageDeviceMetrics{AppInstStatusMsgModel.ZpoolMetrics})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the AppInstStatusMsg resource defined in the Terraform configuration
func AppInstStatusMsgSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cpu": {
			Description: `SPU details`,
			Type:        schema.TypeList, //GoType: CPUSummary
			Elem: &schema.Resource{
				Schema: CPUSummarySchema(),
			},
			Optional: true,
		},

		"io_status_list": {
			Description: `Io bundle status list`,
			Type:        schema.TypeList, //GoType: []*IoBundleStatus
			Elem: &schema.Resource{
				Schema: IoBundleStatusSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
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
			Description: `App instance status`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"app_id": {
			Description: `User defined name of the edge app, unique across the enterprise. Once edge app is created, name can’t be changed`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"app_status_from_t_p_controller": {
			Description: `app status from third party controller`,
			Type:        schema.TypeList, //GoType: AppStatusFromTPController
			Elem: &schema.Resource{
				Schema: AppStatusFromTPControllerSchema(),
			},
			Optional: true,
		},

		"app_type": {
			Description: `type of app`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"boot_time": {
			Description:  `device boot time`,
			Type:         schema.TypeString,
			ValidateFunc: validation.IsRFC3339Time,
			Optional:     true,
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

		"container_status_list": {
			Description: `docker-compose container status`,
			Type:        schema.TypeList, //GoType: []*ComposeContainerStatus
			Elem: &schema.Resource{
				Schema: ComposeContainerStatusSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"device_id": {
			Description: `User defined name of the device, unique across the enterprise. Once device is created, name can’t be changed`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"err_info": {
			Description: `Device error details`,
			Type:        schema.TypeList, //GoType: []*DeviceError
			Elem: &schema.Resource{
				Schema: DeviceErrorSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"id": {
			Description: `app instance id`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"last_update": {
			Description:  `device last update time`,
			Type:         schema.TypeString,
			ValidateFunc: validation.IsRFC3339Time,
			Optional:     true,
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

		"net_counter_list": {
			Description: `network counter list`,
			Type:        schema.TypeList, //GoType: []*NetworkCounters
			Elem: &schema.Resource{
				Schema: NetworkCountersSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"net_status_list": {
			Description: `network status list`,
			Type:        schema.TypeList, //GoType: []*NetworkStatus
			Elem: &schema.Resource{
				Schema: NetworkStatusSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"project_id": {
			Description: `User defined name of the project, unique across the enterprise. Once project is created, name can’t be changed`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"run_state": {
			Description: `app instance state`,
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
			Description: `software state`,
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
			Description: `app instance status title`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"up_time": {
			Description:  `device up time`,
			Type:         schema.TypeString,
			ValidateFunc: validation.IsRFC3339Time,
			Optional:     true,
		},

		"volume_refs": {
			Description: `list of volume ids attached`,
			Type:        schema.TypeList, //GoType: []string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"zpool_metrics": {
			Description: `Last received counters for zvol metrics`,
			Type:        schema.TypeList, //GoType: StorageDeviceMetrics
			Elem: &schema.Resource{
				Schema: StorageDeviceMetricsSchema(),
			},
			Optional: true,
		},
	}
}

// Retrieve property field names for updating the AppInstStatusMsg resource
func GetAppInstStatusMsgPropertyFields() (t []string) {
	return []string{
		"cpu",
		"io_status_list",
		"memory",
		"storage",
		"admin_state",
		"app_id",
		"app_status_from_t_p_controller",
		"app_type",
		"boot_time",
		"cluster_id",
		"container_status_list",
		"deployment_type",
		"device_id",
		"err_info",
		"id",
		"last_update",
		"memory_summary",
		"name",
		"net_counter_list",
		"net_status_list",
		"project_id",
		"run_state",
		"sw_info",
		"sw_state",
		"tags",
		"title",
		"up_time",
		"volume_refs",
		"zpool_metrics",
	}
}
