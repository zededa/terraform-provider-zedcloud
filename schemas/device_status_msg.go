package schemas

import (
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

// Function to perform the following actions:
// (1) Translate DeviceStatusMsg resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func DeviceStatusMsgModel(d *schema.ResourceData) *models.DeviceStatusMsg {
	var cpu *models.CPUSummary // CPUSummary
	CpuInterface, CpuIsSet := d.GetOk("cpu")
	if CpuIsSet {
		CpuMap := CpuInterface.([]interface{})[0].(map[string]interface{})
		cpu = CPUSummaryModelFromMap(CpuMap)
	}
	var memory *models.MemorySummary // MemorySummary
	MemoryInterface, MemoryIsSet := d.GetOk("memory")
	if MemoryIsSet {
		MemoryMap := MemoryInterface.([]interface{})[0].(map[string]interface{})
		memory = MemorySummaryModelFromMap(MemoryMap)
	}
	var storage *models.StorageSummary // StorageSummary
	StorageInterface, StorageIsSet := d.GetOk("storage")
	if StorageIsSet {
		StorageMap := StorageInterface.([]interface{})[0].(map[string]interface{})
		storage = StorageSummaryModelFromMap(StorageMap)
	}
	var adminState *models.AdminState // AdminState
	adminStateInterface, adminStateIsSet := d.GetOk("admin_state")
	if adminStateIsSet {
		adminStateModel := adminStateInterface.(string)
		adminState = models.NewAdminState(models.AdminState(adminStateModel))
	}
	var attestState *models.AttestState // AttestState
	attestStateInterface, attestStateIsSet := d.GetOk("attest_state")
	if attestStateIsSet {
		attestStateModel := attestStateInterface.(string)
		attestState = models.NewAttestState(models.AttestState(attestStateModel))
	}
	blobList, _ := d.Get("blob_list").([]*models.BlobStatus) // []*BlobStatus
	bootTime, _ := d.Get("boot_time").(strfmt.DateTime)
	var capabilities *models.Capabilities // Capabilities
	capabilitiesInterface, capabilitiesIsSet := d.GetOk("capabilities")
	if capabilitiesIsSet {
		capabilitiesMap := capabilitiesInterface.([]interface{})[0].(map[string]interface{})
		capabilities = CapabilitiesModelFromMap(capabilitiesMap)
	}
	clusterID, _ := d.Get("cluster_id").(string)
	dataSecInfo, _ := d.Get("data_sec_info").([]*models.DevDataSecAtRest) // []*DevDataSecAtRest
	devError, _ := d.Get("dev_error").([]*models.DeviceError)             // []*DeviceError
	var deviceRebootReason *models.DeviceBootReason                       // DeviceBootReason
	deviceRebootReasonInterface, deviceRebootReasonIsSet := d.GetOk("device_reboot_reason")
	if deviceRebootReasonIsSet {
		deviceRebootReasonModel := deviceRebootReasonInterface.(string)
		deviceRebootReason = models.NewDeviceBootReason(models.DeviceBootReason(deviceRebootReasonModel))
	}
	var dinfo *models.DeviceInfo // DeviceInfo
	dinfoInterface, dinfoIsSet := d.GetOk("dinfo")
	if dinfoIsSet {
		dinfoMap := dinfoInterface.([]interface{})[0].(map[string]interface{})
		dinfo = DeviceInfoModelFromMap(dinfoMap)
	}
	var dns *models.DNSInfo // DNSInfo
	dnsInterface, dnsIsSet := d.GetOk("dns")
	if dnsIsSet {
		dnsMap := dnsInterface.([]interface{})[0].(map[string]interface{})
		dns = DNSInfoModelFromMap(dnsMap)
	}
	hostName, _ := d.Get("host_name").(string)
	id, _ := d.Get("id").(string)
	ioStatusList, _ := d.Get("io_status_list").([]*models.IoBundleStatus) // []*IoBundleStatus
	lastRebootReason, _ := d.Get("last_reboot_reason").(string)
	lastRebootTime, _ := d.Get("last_reboot_time").(strfmt.DateTime)
	lastUpdate, _ := d.Get("last_update").(strfmt.DateTime)
	var memorySummary *models.DeviceMemorySummary // DeviceMemorySummary
	memorySummaryInterface, memorySummaryIsSet := d.GetOk("memory_summary")
	if memorySummaryIsSet {
		memorySummaryMap := memorySummaryInterface.([]interface{})[0].(map[string]interface{})
		memorySummary = DeviceMemorySummaryModelFromMap(memorySummaryMap)
	}
	var minfo *models.ZManufacturerInfo // ZManufacturerInfo
	minfoInterface, minfoIsSet := d.GetOk("minfo")
	if minfoIsSet {
		minfoMap := minfoInterface.([]interface{})[0].(map[string]interface{})
		minfo = ZManufacturerInfoModelFromMap(minfoMap)
	}
	name, _ := d.Get("name").(string)
	netCounterList, _ := d.Get("net_counter_list").([]*models.NetworkCounters) // []*NetworkCounters
	netStatusList, _ := d.Get("net_status_list").([]*models.NetworkStatus)     // []*NetworkStatus
	projectID, _ := d.Get("project_id").(string)
	rawMetrics, _ := d.Get("raw_metrics").(string)
	rawStatus, _ := d.Get("raw_status").(string)
	var runState *models.RunState // RunState
	runStateInterface, runStateIsSet := d.GetOk("run_state")
	if runStateIsSet {
		runStateModel := runStateInterface.(string)
		runState = models.NewRunState(models.RunState(runStateModel))
	}
	storageList, _ := d.Get("storage_list").([]*models.StorageStatus) // []*StorageStatus
	swInfo, _ := d.Get("sw_info").([]*models.DeviceSWInfo)            // []*DeviceSWInfo
	tags, _ := d.Get("tags").(map[string]string)                      // map[string]string
	title, _ := d.Get("title").(string)
	upTime, _ := d.Get("up_time").(strfmt.DateTime)
	zcCounters, _ := d.Get("zc_counters").([]*models.ZedcloudCounters) // []*ZedcloudCounters
	var zpoolMetrics *models.StorageDeviceMetrics                      // StorageDeviceMetrics
	zpoolMetricsInterface, zpoolMetricsIsSet := d.GetOk("zpool_metrics")
	if zpoolMetricsIsSet {
		zpoolMetricsMap := zpoolMetricsInterface.([]interface{})[0].(map[string]interface{})
		zpoolMetrics = StorageDeviceMetricsModelFromMap(zpoolMetricsMap)
	}
	return &models.DeviceStatusMsg{
		CPU:                cpu,
		Memory:             memory,
		Storage:            storage,
		AdminState:         adminState,
		AttestState:        attestState,
		BlobList:           blobList,
		BootTime:           bootTime,
		Capabilities:       capabilities,
		ClusterID:          clusterID,
		DataSecInfo:        dataSecInfo,
		DevError:           devError,
		DeviceRebootReason: deviceRebootReason,
		Dinfo:              dinfo,
		DNS:                dns,
		HostName:           hostName,
		ID:                 id,
		IoStatusList:       ioStatusList,
		LastRebootReason:   lastRebootReason,
		LastRebootTime:     lastRebootTime,
		LastUpdate:         lastUpdate,
		MemorySummary:      memorySummary,
		Minfo:              minfo,
		Name:               name,
		NetCounterList:     netCounterList,
		NetStatusList:      netStatusList,
		ProjectID:          projectID,
		RawMetrics:         rawMetrics,
		RawStatus:          rawStatus,
		RunState:           runState,
		StorageList:        storageList,
		SwInfo:             swInfo,
		Tags:               tags,
		Title:              title,
		UpTime:             upTime,
		ZcCounters:         zcCounters,
		ZpoolMetrics:       zpoolMetrics,
	}
}

func DeviceStatusMsgModelFromMap(m map[string]interface{}) *models.DeviceStatusMsg {
	var cpu *models.CPUSummary // CPUSummary
	CpuInterface, CpuIsSet := m["cpu"]
	if CpuIsSet {
		CpuMap := CpuInterface.([]interface{})[0].(map[string]interface{})
		cpu = CPUSummaryModelFromMap(CpuMap)
	}
	//
	var memory *models.MemorySummary // MemorySummary
	MemoryInterface, MemoryIsSet := m["memory"]
	if MemoryIsSet {
		MemoryMap := MemoryInterface.([]interface{})[0].(map[string]interface{})
		memory = MemorySummaryModelFromMap(MemoryMap)
	}
	//
	var storage *models.StorageSummary // StorageSummary
	StorageInterface, StorageIsSet := m["storage"]
	if StorageIsSet {
		StorageMap := StorageInterface.([]interface{})[0].(map[string]interface{})
		storage = StorageSummaryModelFromMap(StorageMap)
	}
	//
	adminState := m["admin_state"].(*models.AdminState)    // AdminState
	attestState := m["attest_state"].(*models.AttestState) // AttestState
	blobList := m["blob_list"].([]*models.BlobStatus)      // []*BlobStatus
	bootTime := m["boot_time"].(strfmt.DateTime)
	var capabilities *models.Capabilities // Capabilities
	capabilitiesInterface, capabilitiesIsSet := m["capabilities"]
	if capabilitiesIsSet {
		capabilitiesMap := capabilitiesInterface.([]interface{})[0].(map[string]interface{})
		capabilities = CapabilitiesModelFromMap(capabilitiesMap)
	}
	//
	clusterID := m["cluster_id"].(string)
	dataSecInfo := m["data_sec_info"].([]*models.DevDataSecAtRest)             // []*DevDataSecAtRest
	devError := m["dev_error"].([]*models.DeviceError)                         // []*DeviceError
	deviceRebootReason := m["device_reboot_reason"].(*models.DeviceBootReason) // DeviceBootReason
	var dinfo *models.DeviceInfo                                               // DeviceInfo
	dinfoInterface, dinfoIsSet := m["dinfo"]
	if dinfoIsSet {
		dinfoMap := dinfoInterface.([]interface{})[0].(map[string]interface{})
		dinfo = DeviceInfoModelFromMap(dinfoMap)
	}
	//
	var dns *models.DNSInfo // DNSInfo
	dnsInterface, dnsIsSet := m["dns"]
	if dnsIsSet {
		dnsMap := dnsInterface.([]interface{})[0].(map[string]interface{})
		dns = DNSInfoModelFromMap(dnsMap)
	}
	//
	hostName := m["host_name"].(string)
	id := m["id"].(string)
	ioStatusList := m["io_status_list"].([]*models.IoBundleStatus) // []*IoBundleStatus
	lastRebootReason := m["last_reboot_reason"].(string)
	lastRebootTime := m["last_reboot_time"].(strfmt.DateTime)
	lastUpdate := m["last_update"].(strfmt.DateTime)
	var memorySummary *models.DeviceMemorySummary // DeviceMemorySummary
	memorySummaryInterface, memorySummaryIsSet := m["memory_summary"]
	if memorySummaryIsSet {
		memorySummaryMap := memorySummaryInterface.([]interface{})[0].(map[string]interface{})
		memorySummary = DeviceMemorySummaryModelFromMap(memorySummaryMap)
	}
	//
	var minfo *models.ZManufacturerInfo // ZManufacturerInfo
	minfoInterface, minfoIsSet := m["minfo"]
	if minfoIsSet {
		minfoMap := minfoInterface.([]interface{})[0].(map[string]interface{})
		minfo = ZManufacturerInfoModelFromMap(minfoMap)
	}
	//
	name := m["name"].(string)
	netCounterList := m["net_counter_list"].([]*models.NetworkCounters) // []*NetworkCounters
	netStatusList := m["net_status_list"].([]*models.NetworkStatus)     // []*NetworkStatus
	projectID := m["project_id"].(string)
	rawMetrics := m["raw_metrics"].(string)
	rawStatus := m["raw_status"].(string)
	runState := m["run_state"].(*models.RunState)              // RunState
	storageList := m["storage_list"].([]*models.StorageStatus) // []*StorageStatus
	swInfo := m["sw_info"].([]*models.DeviceSWInfo)            // []*DeviceSWInfo
	tags := m["tags"].(map[string]string)
	title := m["title"].(string)
	upTime := m["up_time"].(strfmt.DateTime)
	zcCounters := m["zc_counters"].([]*models.ZedcloudCounters) // []*ZedcloudCounters
	var zpoolMetrics *models.StorageDeviceMetrics               // StorageDeviceMetrics
	zpoolMetricsInterface, zpoolMetricsIsSet := m["zpool_metrics"]
	if zpoolMetricsIsSet {
		zpoolMetricsMap := zpoolMetricsInterface.([]interface{})[0].(map[string]interface{})
		zpoolMetrics = StorageDeviceMetricsModelFromMap(zpoolMetricsMap)
	}
	//
	return &models.DeviceStatusMsg{
		CPU:                cpu,
		Memory:             memory,
		Storage:            storage,
		AdminState:         adminState,
		AttestState:        attestState,
		BlobList:           blobList,
		BootTime:           bootTime,
		Capabilities:       capabilities,
		ClusterID:          clusterID,
		DataSecInfo:        dataSecInfo,
		DevError:           devError,
		DeviceRebootReason: deviceRebootReason,
		Dinfo:              dinfo,
		DNS:                dns,
		HostName:           hostName,
		ID:                 id,
		IoStatusList:       ioStatusList,
		LastRebootReason:   lastRebootReason,
		LastRebootTime:     lastRebootTime,
		LastUpdate:         lastUpdate,
		MemorySummary:      memorySummary,
		Minfo:              minfo,
		Name:               name,
		NetCounterList:     netCounterList,
		NetStatusList:      netStatusList,
		ProjectID:          projectID,
		RawMetrics:         rawMetrics,
		RawStatus:          rawStatus,
		RunState:           runState,
		StorageList:        storageList,
		SwInfo:             swInfo,
		Tags:               tags,
		Title:              title,
		UpTime:             upTime,
		ZcCounters:         zcCounters,
		ZpoolMetrics:       zpoolMetrics,
	}
}

// Update the underlying DeviceStatusMsg resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDeviceStatusMsgResourceData(d *schema.ResourceData, m *models.DeviceStatusMsg) {
	d.Set("cpu", SetCPUSummarySubResourceData([]*models.CPUSummary{m.CPU}))
	d.Set("memory", SetMemorySummarySubResourceData([]*models.MemorySummary{m.Memory}))
	d.Set("storage", SetStorageSummarySubResourceData([]*models.StorageSummary{m.Storage}))
	d.Set("admin_state", m.AdminState)
	d.Set("attest_state", m.AttestState)
	d.Set("blob_list", SetBlobStatusSubResourceData(m.BlobList))
	d.Set("boot_time", m.BootTime)
	d.Set("capabilities", SetCapabilitiesSubResourceData([]*models.Capabilities{m.Capabilities}))
	d.Set("cluster_id", m.ClusterID)
	d.Set("data_sec_info", SetDevDataSecAtRestSubResourceData(m.DataSecInfo))
	d.Set("dev_error", SetDeviceErrorSubResourceData(m.DevError))
	d.Set("device_reboot_reason", m.DeviceRebootReason)
	d.Set("dinfo", SetDeviceInfoSubResourceData([]*models.DeviceInfo{m.Dinfo}))
	d.Set("dns", SetDNSInfoSubResourceData([]*models.DNSInfo{m.DNS}))
	d.Set("host_name", m.HostName)
	d.Set("id", m.ID)
	d.Set("io_status_list", SetIoBundleStatusSubResourceData(m.IoStatusList))
	d.Set("last_reboot_reason", m.LastRebootReason)
	d.Set("last_reboot_time", m.LastRebootTime)
	d.Set("last_update", m.LastUpdate)
	d.Set("memory_summary", SetDeviceMemorySummarySubResourceData([]*models.DeviceMemorySummary{m.MemorySummary}))
	d.Set("minfo", SetZManufacturerInfoSubResourceData([]*models.ZManufacturerInfo{m.Minfo}))
	d.Set("name", m.Name)
	d.Set("net_counter_list", SetNetworkCountersSubResourceData(m.NetCounterList))
	d.Set("net_status_list", SetNetworkStatusSubResourceData(m.NetStatusList))
	d.Set("project_id", m.ProjectID)
	d.Set("raw_metrics", m.RawMetrics)
	d.Set("raw_status", m.RawStatus)
	d.Set("run_state", m.RunState)
	d.Set("storage_list", SetStorageStatusSubResourceData(m.StorageList))
	d.Set("sw_info", SetDeviceSWInfoSubResourceData(m.SwInfo))
	d.Set("tags", m.Tags)
	d.Set("title", m.Title)
	d.Set("up_time", m.UpTime)
	d.Set("zc_counters", SetZedcloudCountersSubResourceData(m.ZcCounters))
	d.Set("zpool_metrics", SetStorageDeviceMetricsSubResourceData([]*models.StorageDeviceMetrics{m.ZpoolMetrics}))
}

// Iterate through and update the DeviceStatusMsg resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDeviceStatusMsgSubResourceData(m []*models.DeviceStatusMsg) (d []*map[string]interface{}) {
	for _, DeviceStatusMsgModel := range m {
		if DeviceStatusMsgModel != nil {
			properties := make(map[string]interface{})
			properties["cpu"] = SetCPUSummarySubResourceData([]*models.CPUSummary{DeviceStatusMsgModel.CPU})
			properties["memory"] = SetMemorySummarySubResourceData([]*models.MemorySummary{DeviceStatusMsgModel.Memory})
			properties["storage"] = SetStorageSummarySubResourceData([]*models.StorageSummary{DeviceStatusMsgModel.Storage})
			properties["admin_state"] = DeviceStatusMsgModel.AdminState
			properties["attest_state"] = DeviceStatusMsgModel.AttestState
			properties["blob_list"] = SetBlobStatusSubResourceData(DeviceStatusMsgModel.BlobList)
			properties["boot_time"] = DeviceStatusMsgModel.BootTime.String()
			properties["capabilities"] = SetCapabilitiesSubResourceData([]*models.Capabilities{DeviceStatusMsgModel.Capabilities})
			properties["cluster_id"] = DeviceStatusMsgModel.ClusterID
			properties["data_sec_info"] = SetDevDataSecAtRestSubResourceData(DeviceStatusMsgModel.DataSecInfo)
			properties["dev_error"] = SetDeviceErrorSubResourceData(DeviceStatusMsgModel.DevError)
			properties["device_reboot_reason"] = DeviceStatusMsgModel.DeviceRebootReason
			properties["dinfo"] = SetDeviceInfoSubResourceData([]*models.DeviceInfo{DeviceStatusMsgModel.Dinfo})
			properties["dns"] = SetDNSInfoSubResourceData([]*models.DNSInfo{DeviceStatusMsgModel.DNS})
			properties["host_name"] = DeviceStatusMsgModel.HostName
			properties["id"] = DeviceStatusMsgModel.ID
			properties["io_status_list"] = SetIoBundleStatusSubResourceData(DeviceStatusMsgModel.IoStatusList)
			properties["last_reboot_reason"] = DeviceStatusMsgModel.LastRebootReason
			properties["last_reboot_time"] = DeviceStatusMsgModel.LastRebootTime.String()
			properties["last_update"] = DeviceStatusMsgModel.LastUpdate.String()
			properties["memory_summary"] = SetDeviceMemorySummarySubResourceData([]*models.DeviceMemorySummary{DeviceStatusMsgModel.MemorySummary})
			properties["minfo"] = SetZManufacturerInfoSubResourceData([]*models.ZManufacturerInfo{DeviceStatusMsgModel.Minfo})
			properties["name"] = DeviceStatusMsgModel.Name
			properties["net_counter_list"] = SetNetworkCountersSubResourceData(DeviceStatusMsgModel.NetCounterList)
			properties["net_status_list"] = SetNetworkStatusSubResourceData(DeviceStatusMsgModel.NetStatusList)
			properties["project_id"] = DeviceStatusMsgModel.ProjectID
			properties["raw_metrics"] = DeviceStatusMsgModel.RawMetrics
			properties["raw_status"] = DeviceStatusMsgModel.RawStatus
			properties["run_state"] = DeviceStatusMsgModel.RunState
			properties["storage_list"] = SetStorageStatusSubResourceData(DeviceStatusMsgModel.StorageList)
			properties["sw_info"] = SetDeviceSWInfoSubResourceData(DeviceStatusMsgModel.SwInfo)
			properties["tags"] = DeviceStatusMsgModel.Tags
			properties["title"] = DeviceStatusMsgModel.Title
			properties["up_time"] = DeviceStatusMsgModel.UpTime.String()
			properties["zc_counters"] = SetZedcloudCountersSubResourceData(DeviceStatusMsgModel.ZcCounters)
			properties["zpool_metrics"] = SetStorageDeviceMetricsSubResourceData([]*models.StorageDeviceMetrics{DeviceStatusMsgModel.ZpoolMetrics})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the DeviceStatusMsg resource defined in the Terraform configuration
func DeviceStatusMsgSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cpu": {
			Description: ``,
			Type:        schema.TypeList, //GoType: CPUSummary
			Elem: &schema.Resource{
				Schema: CPUSummarySchema(),
			},
			Optional: true,
		},

		"memory": {
			Description: `memory - OBSOLETE. Use memorySummary instead.`,
			Type:        schema.TypeList, //GoType: MemorySummary
			Elem: &schema.Resource{
				Schema: MemorySummarySchema(),
			},
			Optional: true,
		},

		"storage": {
			Description: ``,
			Type:        schema.TypeList, //GoType: StorageSummary
			Elem: &schema.Resource{
				Schema: StorageSummarySchema(),
			},
			Optional: true,
		},

		"admin_state": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"attest_state": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"blob_list": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []*BlobStatus
			Elem: &schema.Resource{
				Schema: BlobStatusSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"boot_time": {
			Description:  ``,
			Type:         schema.TypeString,
			ValidateFunc: validation.IsRFC3339Time,
			Optional:     true,
		},

		"capabilities": {
			Description: `Edge node virtualization capabilities.`,
			Type:        schema.TypeList, //GoType: Capabilities
			Elem: &schema.Resource{
				Schema: CapabilitiesSchema(),
			},
			Optional: true,
		},

		"cluster_id": {
			Description: `System defined universally unique clusterInstance ID, unique across the enterprise.`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"data_sec_info": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []*DevDataSecAtRest
			Elem: &schema.Resource{
				Schema: DevDataSecAtRestSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"dev_error": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []*DeviceError
			Elem: &schema.Resource{
				Schema: DeviceErrorSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"device_reboot_reason": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"dinfo": {
			Description: ``,
			Type:        schema.TypeList, //GoType: DeviceInfo
			Elem: &schema.Resource{
				Schema: DeviceInfoSchema(),
			},
			Optional: true,
		},

		"dns": {
			Description: ``,
			Type:        schema.TypeList, //GoType: DNSInfo
			Elem: &schema.Resource{
				Schema: DNSInfoSchema(),
			},
			Optional: true,
		},

		"host_name": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"id": {
			Description: ``,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"io_status_list": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []*IoBundleStatus
			Elem: &schema.Resource{
				Schema: IoBundleStatusSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"last_reboot_reason": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"last_reboot_time": {
			Description:  ``,
			Type:         schema.TypeString,
			ValidateFunc: validation.IsRFC3339Time,
			Optional:     true,
		},

		"last_update": {
			Description:  ``,
			Type:         schema.TypeString,
			ValidateFunc: validation.IsRFC3339Time,
			Optional:     true,
		},

		"memory_summary": {
			Description: `Device memory Info`,
			Type:        schema.TypeList, //GoType: DeviceMemorySummary
			Elem: &schema.Resource{
				Schema: DeviceMemorySummarySchema(),
			},
			Optional: true,
		},

		"minfo": {
			Description: ``,
			Type:        schema.TypeList, //GoType: ZManufacturerInfo
			Elem: &schema.Resource{
				Schema: ZManufacturerInfoSchema(),
			},
			Optional: true,
		},

		"name": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"net_counter_list": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []*NetworkCounters
			Elem: &schema.Resource{
				Schema: NetworkCountersSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"net_status_list": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []*NetworkStatus
			Elem: &schema.Resource{
				Schema: NetworkStatusSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"project_id": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"raw_metrics": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"raw_status": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"run_state": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"storage_list": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []*StorageStatus
			Elem: &schema.Resource{
				Schema: StorageStatusSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"sw_info": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []*DeviceSWInfo
			Elem: &schema.Resource{
				Schema: DeviceSWInfoSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
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
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"up_time": {
			Description:  ``,
			Type:         schema.TypeString,
			ValidateFunc: validation.IsRFC3339Time,
			Optional:     true,
		},

		"zc_counters": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []*ZedcloudCounters
			Elem: &schema.Resource{
				Schema: ZedcloudCountersSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"zpool_metrics": {
			Description: `Last received counters for zpool metrics.`,
			Type:        schema.TypeList, //GoType: StorageDeviceMetrics
			Elem: &schema.Resource{
				Schema: StorageDeviceMetricsSchema(),
			},
			Optional: true,
		},
	}
}

// Retrieve property field names for updating the DeviceStatusMsg resource
func GetDeviceStatusMsgPropertyFields() (t []string) {
	return []string{
		"cpu",
		"memory",
		"storage",
		"admin_state",
		"attest_state",
		"blob_list",
		"boot_time",
		"capabilities",
		"cluster_id",
		"data_sec_info",
		"dev_error",
		"device_reboot_reason",
		"dinfo",
		"dns",
		"host_name",
		"id",
		"io_status_list",
		"last_reboot_reason",
		"last_reboot_time",
		"last_update",
		"memory_summary",
		"minfo",
		"name",
		"net_counter_list",
		"net_status_list",
		"project_id",
		"raw_metrics",
		"raw_status",
		"run_state",
		"storage_list",
		"sw_info",
		"tags",
		"title",
		"up_time",
		"zc_counters",
		"zpool_metrics",
	}
}
