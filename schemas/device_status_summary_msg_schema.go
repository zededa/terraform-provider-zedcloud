package schemas

import (
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate DeviceStatusSummaryMsg resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func DeviceStatusSummaryMsgModel(d *schema.ResourceData) *models.DeviceStatusSummaryMsg {
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
	adminStateModel, ok := d.Get("admin_state").(models.AdminState) // AdminState
	adminState := &adminStateModel
	if !ok {
		adminState = nil
	}
	appInstCountInt, _ := d.Get("app_inst_count").(int)
	appInstCount := int64(appInstCountInt)
	clusterID, _ := d.Get("cluster_id").(string)
	debugKnob, _ := d.Get("debug_knob").(bool)
	debugKnobExpiryTime, _ := d.Get("debug_knob_expiry_time").(strfmt.DateTime)
	devError, _ := d.Get("dev_error").([]*models.DeviceError) // []*DeviceError
	var dinfo *models.DeviceInfo                              // DeviceInfo
	dinfoInterface, dinfoIsSet := d.GetOk("dinfo")
	if dinfoIsSet {
		dinfoMap := dinfoInterface.([]interface{})[0].(map[string]interface{})
		dinfo = DeviceInfoModelFromMap(dinfoMap)
	}
	edgeviewActive, _ := d.Get("edgeview_active").(bool)
	id, _ := d.Get("id").(string)
	location, _ := d.Get("location").(string)
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
	netStatusList, _ := d.Get("net_status_list").([]*models.NetworkStatus) // []*NetworkStatus
	projectID, _ := d.Get("project_id").(string)
	projectName, _ := d.Get("project_name").(string)
	runStateModel, ok := d.Get("run_state").(models.RunState) // RunState
	runState := &runStateModel
	if !ok {
		runState = nil
	}
	swInfo, _ := d.Get("sw_info").([]*models.DeviceSWInfo) // []*DeviceSWInfo
	tags, _ := d.Get("tags").(map[string]string)           // map[string]string
	title, _ := d.Get("title").(string)
	return &models.DeviceStatusSummaryMsg{
		CPU:                 cpu,
		Memory:              memory,
		Storage:             storage,
		AdminState:          adminState,
		AppInstCount:        appInstCount,
		ClusterID:           clusterID,
		DebugKnob:           debugKnob,
		DebugKnobExpiryTime: debugKnobExpiryTime,
		DevError:            devError,
		Dinfo:               dinfo,
		EdgeviewActive:      edgeviewActive,
		ID:                  id,
		Location:            location,
		MemorySummary:       memorySummary,
		Minfo:               minfo,
		Name:                name,
		NetStatusList:       netStatusList,
		ProjectID:           projectID,
		ProjectName:         projectName,
		RunState:            runState,
		SwInfo:              swInfo,
		Tags:                tags,
		Title:               title,
	}
}

func DeviceStatusSummaryMsgModelFromMap(m map[string]interface{}) *models.DeviceStatusSummaryMsg {
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
	adminState := m["admin_state"].(*models.AdminState) // AdminState
	appInstCount := int64(m["app_inst_count"].(int))    // int64 false false false
	clusterID := m["cluster_id"].(string)
	debugKnob := m["debug_knob"].(bool)
	debugKnobExpiryTime := m["debug_knob_expiry_time"].(strfmt.DateTime)
	devError := m["dev_error"].([]*models.DeviceError) // []*DeviceError
	var dinfo *models.DeviceInfo                       // DeviceInfo
	dinfoInterface, dinfoIsSet := m["dinfo"]
	if dinfoIsSet {
		dinfoMap := dinfoInterface.([]interface{})[0].(map[string]interface{})
		dinfo = DeviceInfoModelFromMap(dinfoMap)
	}
	//
	edgeviewActive := m["edgeview_active"].(bool)
	id := m["id"].(string)
	location := m["location"].(string)
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
	netStatusList := m["net_status_list"].([]*models.NetworkStatus) // []*NetworkStatus
	projectID := m["project_id"].(string)
	projectName := m["project_name"].(string)
	runState := m["run_state"].(*models.RunState)   // RunState
	swInfo := m["sw_info"].([]*models.DeviceSWInfo) // []*DeviceSWInfo
	tags := m["tags"].(map[string]string)
	title := m["title"].(string)
	return &models.DeviceStatusSummaryMsg{
		CPU:                 cpu,
		Memory:              memory,
		Storage:             storage,
		AdminState:          adminState,
		AppInstCount:        appInstCount,
		ClusterID:           clusterID,
		DebugKnob:           debugKnob,
		DebugKnobExpiryTime: debugKnobExpiryTime,
		DevError:            devError,
		Dinfo:               dinfo,
		EdgeviewActive:      edgeviewActive,
		ID:                  id,
		Location:            location,
		MemorySummary:       memorySummary,
		Minfo:               minfo,
		Name:                name,
		NetStatusList:       netStatusList,
		ProjectID:           projectID,
		ProjectName:         projectName,
		RunState:            runState,
		SwInfo:              swInfo,
		Tags:                tags,
		Title:               title,
	}
}

// Update the underlying DeviceStatusSummaryMsg resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDeviceStatusSummaryMsgResourceData(d *schema.ResourceData, m *models.DeviceStatusSummaryMsg) {
	d.Set("cpu", SetCPUSummarySubResourceData([]*models.CPUSummary{m.CPU}))
	d.Set("memory", SetMemorySummarySubResourceData([]*models.MemorySummary{m.Memory}))
	d.Set("storage", SetStorageSummarySubResourceData([]*models.StorageSummary{m.Storage}))
	d.Set("admin_state", m.AdminState)
	d.Set("app_inst_count", m.AppInstCount)
	d.Set("cluster_id", m.ClusterID)
	d.Set("debug_knob", m.DebugKnob)
	d.Set("debug_knob_expiry_time", m.DebugKnobExpiryTime)
	d.Set("dev_error", SetDeviceErrorSubResourceData(m.DevError))
	d.Set("dinfo", SetDeviceInfoSubResourceData([]*models.DeviceInfo{m.Dinfo}))
	d.Set("edgeview_active", m.EdgeviewActive)
	d.Set("id", m.ID)
	d.Set("location", m.Location)
	d.Set("memory_summary", SetDeviceMemorySummarySubResourceData([]*models.DeviceMemorySummary{m.MemorySummary}))
	d.Set("minfo", SetZManufacturerInfoSubResourceData([]*models.ZManufacturerInfo{m.Minfo}))
	d.Set("name", m.Name)
	d.Set("net_status_list", SetNetworkStatusSubResourceData(m.NetStatusList))
	d.Set("project_id", m.ProjectID)
	d.Set("project_name", m.ProjectName)
	d.Set("run_state", m.RunState)
	d.Set("sw_info", SetDeviceSWInfoSubResourceData(m.SwInfo))
	d.Set("tags", m.Tags)
	d.Set("title", m.Title)
}

// Iterate through and update the DeviceStatusSummaryMsg resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDeviceStatusSummaryMsgSubResourceData(m []*models.DeviceStatusSummaryMsg) (d []*map[string]interface{}) {
	for _, DeviceStatusSummaryMsgModel := range m {
		if DeviceStatusSummaryMsgModel != nil {
			properties := make(map[string]interface{})
			properties["cpu"] = SetCPUSummarySubResourceData([]*models.CPUSummary{DeviceStatusSummaryMsgModel.CPU})
			properties["memory"] = SetMemorySummarySubResourceData([]*models.MemorySummary{DeviceStatusSummaryMsgModel.Memory})
			properties["storage"] = SetStorageSummarySubResourceData([]*models.StorageSummary{DeviceStatusSummaryMsgModel.Storage})
			properties["admin_state"] = DeviceStatusSummaryMsgModel.AdminState
			properties["app_inst_count"] = DeviceStatusSummaryMsgModel.AppInstCount
			properties["cluster_id"] = DeviceStatusSummaryMsgModel.ClusterID
			properties["debug_knob"] = DeviceStatusSummaryMsgModel.DebugKnob
			properties["debug_knob_expiry_time"] = DeviceStatusSummaryMsgModel.DebugKnobExpiryTime.String()
			properties["dev_error"] = SetDeviceErrorSubResourceData(DeviceStatusSummaryMsgModel.DevError)
			properties["dinfo"] = SetDeviceInfoSubResourceData([]*models.DeviceInfo{DeviceStatusSummaryMsgModel.Dinfo})
			properties["edgeview_active"] = DeviceStatusSummaryMsgModel.EdgeviewActive
			properties["id"] = DeviceStatusSummaryMsgModel.ID
			properties["location"] = DeviceStatusSummaryMsgModel.Location
			properties["memory_summary"] = SetDeviceMemorySummarySubResourceData([]*models.DeviceMemorySummary{DeviceStatusSummaryMsgModel.MemorySummary})
			properties["minfo"] = SetZManufacturerInfoSubResourceData([]*models.ZManufacturerInfo{DeviceStatusSummaryMsgModel.Minfo})
			properties["name"] = DeviceStatusSummaryMsgModel.Name
			properties["net_status_list"] = SetNetworkStatusSubResourceData(DeviceStatusSummaryMsgModel.NetStatusList)
			properties["project_id"] = DeviceStatusSummaryMsgModel.ProjectID
			properties["project_name"] = DeviceStatusSummaryMsgModel.ProjectName
			properties["run_state"] = DeviceStatusSummaryMsgModel.RunState
			properties["sw_info"] = SetDeviceSWInfoSubResourceData(DeviceStatusSummaryMsgModel.SwInfo)
			properties["tags"] = DeviceStatusSummaryMsgModel.Tags
			properties["title"] = DeviceStatusSummaryMsgModel.Title
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the DeviceStatusSummaryMsg resource defined in the Terraform configuration
func DeviceStatusSummaryMsgSchema() map[string]*schema.Schema {
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
			Description: `Memory - OBSOLETE. Use memorySummary instead.`,
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

		"app_inst_count": {
			Description: `App instance count that is actively running on the device`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"cluster_id": {
			Description: `System defined universally unique clusterInstance ID, unique across the enterprise.`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"debug_knob": {
			Description: ``,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"debug_knob_expiry_time": {
			Description:  ``,
			Type:         schema.TypeString,
			ValidateFunc: validation.IsRFC3339Time,
			Optional:     true,
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

		"dinfo": {
			Description: ``,
			Type:        schema.TypeList, //GoType: DeviceInfo
			Elem: &schema.Resource{
				Schema: DeviceInfoSchema(),
			},
			Optional: true,
		},

		"edgeview_active": {
			Description: `Device edgeview session active`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"id": {
			Description: ``,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"location": {
			Description: `Device location`,
			Type:        schema.TypeString,
			Optional:    true,
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

		"project_name": {
			Description: `Project name to which device is associated with`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"run_state": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
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
	}
}

// Retrieve property field names for updating the DeviceStatusSummaryMsg resource
func GetDeviceStatusSummaryMsgPropertyFields() (t []string) {
	return []string{
		"cpu",
		"memory",
		"storage",
		"admin_state",
		"app_inst_count",
		"cluster_id",
		"debug_knob",
		"debug_knob_expiry_time",
		"dev_error",
		"dinfo",
		"edgeview_active",
		"id",
		"location",
		"memory_summary",
		"minfo",
		"name",
		"net_status_list",
		"project_id",
		"project_name",
		"run_state",
		"sw_info",
		"tags",
		"title",
	}
}
