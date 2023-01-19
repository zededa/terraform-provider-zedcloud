package schemas

import (
	"strings"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func DeviceConfigModel(d *schema.ResourceData) *models.DeviceConfig {
	var adminState *models.AdminState // AdminState
	adminStateInterface, adminStateIsSet := d.GetOk("admin_state")
	if adminStateIsSet {
		adminStateModel := adminStateInterface.(string)
		adminState = models.NewAdminState(models.AdminState(adminStateModel))
	}
	assetID, _ := d.Get("asset_id").(string)
	var baseImage []*models.BaseOSImage // []*BaseOSImage
	baseImageInterface, baseImageIsSet := d.GetOk("base_image")
	if baseImageIsSet {
		for _, v := range baseImageInterface.([]interface{}) {
			if v == nil {
				continue
			}
			m := BaseOSImageModelFromMap(v.(map[string]interface{}))
			baseImage = append(baseImage, m)
		}
	}
	baseOsRetryCounterInt, _ := d.Get("base_os_retry_counter").(int)
	baseOsRetryCounter := int64(baseOsRetryCounterInt)
	baseOsRetryTime, _ := d.Get("base_os_retry_time").(string)
	clientIP, _ := d.Get("client_ip").(string)
	clusterID, _ := d.Get("cluster_id").(string)
	var configItem []*models.EDConfigItem // []*EDConfigItem
	configItemInterface, configItemIsSet := d.GetOk("config_item")
	if configItemIsSet {
		for _, v := range configItemInterface.([]interface{}) {
			if v == nil {
				continue
			}
			m := EDConfigItemModelFromMap(v.(map[string]interface{}))
			configItem = append(configItem, m)
		}
	}
	cpuInt, _ := d.Get("cpu").(int)
	cpu := int64(cpuInt)
	var debugKnob *models.DebugKnobDetail // DebugKnobDetail
	debugKnobInterface, debugKnobIsSet := d.GetOk("debug_knob")
	if debugKnobIsSet {
		debugKnobMap := debugKnobInterface.([]interface{})[0].(map[string]interface{})
		debugKnob = DebugKnobDetailModelFromMap(debugKnobMap)
	}
	var defaultNetInst *models.NetInstConfig // NetInstConfig
	defaultNetInstInterface, defaultNetInstIsSet := d.GetOk("default_net_inst")
	if defaultNetInstIsSet {
		defaultNetInstMap := defaultNetInstInterface.([]interface{})[0].(map[string]interface{})
		defaultNetInst = NetInstConfigModelFromMap(defaultNetInstMap)
	}
	deploymentTag, _ := d.Get("deployment_tag").(string)
	deprecated, _ := d.Get("deprecated").(string)
	description, _ := d.Get("description").(string)
	var devLocation *models.GeoLocation // GeoLocation
	devLocationInterface, devLocationIsSet := d.GetOk("dev_location")
	if devLocationIsSet {
		devLocationMap := devLocationInterface.([]interface{})[0].(map[string]interface{})
		devLocation = GeoLocationModelFromMap(devLocationMap)
	}
	var dlisp *models.DeviceLisp // DeviceLisp
	dlispInterface, dlispIsSet := d.GetOk("dlisp")
	if dlispIsSet {
		dlispMap := dlispInterface.([]interface{})[0].(map[string]interface{})
		dlisp = DeviceLispModelFromMap(dlispMap)
	}
	var edgeviewconfig *models.EdgeviewCfg // EdgeviewCfg
	edgeviewconfigInterface, edgeviewconfigIsSet := d.GetOk("edgeviewconfig")
	if edgeviewconfigIsSet {
		edgeviewconfigMap := edgeviewconfigInterface.([]interface{})[0].(map[string]interface{})
		edgeviewconfig = EdgeviewCfgModelFromMap(edgeviewconfigMap)
	}
	generateSoftSerial, _ := d.Get("generate_soft_serial").(bool)
	id, _ := d.Get("id").(string)
	identity, _ := d.Get("identity").(strfmt.Base64)
	var interfaces []*models.SysInterface // []*SysInterface
	interfacesInterface, interfacesIsSet := d.GetOk("interfaces")
	if interfacesIsSet {
		for _, v := range interfacesInterface.([]interface{}) {
			if v == nil {
				continue
			}
			m := SysInterfaceModelFromMap(v.(map[string]interface{}))
			interfaces = append(interfaces, m)
		}
	}
	location, _ := d.Get("location").(string)
	memoryInt, _ := d.Get("memory").(int)
	memory := int64(memoryInt)
	modelID, _ := d.Get("model_id").(string)
	name, _ := d.Get("name").(string)
	obkey, _ := d.Get("obkey").(string)
	var onboarding *models.DeviceCerts // DeviceCerts
	onboardingInterface, onboardingIsSet := d.GetOk("onboarding")
	if onboardingIsSet {
		onboardingMap := onboardingInterface.([]interface{})[0].(map[string]interface{})
		onboarding = DeviceCertsModelFromMap(onboardingMap)
	}
	preparePowerOffCounterInt, _ := d.Get("prepare_power_off_counter").(int)
	preparePowerOffCounter := int64(preparePowerOffCounterInt)
	preparePowerOffTime, _ := d.Get("prepare_power_off_time").(string)
	projectID, _ := d.Get("project_id").(string)
	resetCounterInt, _ := d.Get("reset_counter").(int)
	resetCounter := int64(resetCounterInt)
	resetTime, _ := d.Get("reset_time").(string)
	var revision *models.ObjectRevision // ObjectRevision
	revisionInterface, revisionIsSet := d.GetOk("revision")
	if revisionIsSet {
		revisionMap := revisionInterface.([]interface{})[0].(map[string]interface{})
		revision = ObjectRevisionModelFromMap(revisionMap)
	}
	serialno, _ := d.Get("serialno").(string)
	sitePictures, _ := d.Get("site_pictures").([]string)
	storageInt, _ := d.Get("storage").(int)
	storage := int64(storageInt)
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
	threadInt, _ := d.Get("thread").(int)
	thread := int64(threadInt)
	title, _ := d.Get("title").(string)
	token, _ := d.Get("token").(string)
	var utype *models.ModelArchType // ModelArchType
	utypeInterface, utypeIsSet := d.GetOk("utype")
	if utypeIsSet {
		utypeModel := utypeInterface.(string)
		utype = models.NewModelArchType(models.ModelArchType(utypeModel))
	}
	return &models.DeviceConfig{
		AdminState:             adminState,
		AssetID:                assetID,
		BaseImage:              baseImage,
		BaseOsRetryCounter:     baseOsRetryCounter,
		BaseOsRetryTime:        baseOsRetryTime,
		ClientIP:               clientIP,
		ClusterID:              clusterID,
		ConfigItem:             configItem,
		CPU:                    cpu,
		DebugKnob:              debugKnob,
		DefaultNetInst:         defaultNetInst,
		DeploymentTag:          deploymentTag,
		Deprecated:             deprecated,
		Description:            description,
		DevLocation:            devLocation,
		Dlisp:                  dlisp,
		Edgeviewconfig:         edgeviewconfig,
		GenerateSoftSerial:     generateSoftSerial,
		ID:                     id,
		Identity:               identity,
		Interfaces:             interfaces,
		Location:               location,
		Memory:                 memory,
		ModelID:                &modelID, // string true false false
		Name:                   &name,    // string true false false
		Obkey:                  obkey,
		Onboarding:             onboarding,
		PreparePowerOffCounter: preparePowerOffCounter,
		PreparePowerOffTime:    preparePowerOffTime,
		ProjectID:              &projectID, // string true false false
		ResetCounter:           resetCounter,
		ResetTime:              resetTime,
		Revision:               revision,
		Serialno:               serialno,
		SitePictures:           sitePictures,
		Storage:                storage,
		Tags:                   tags,
		Thread:                 thread,
		Title:                  &title, // string true false false
		Token:                  token,
		Utype:                  utype,
	}
}

//func DeviceConfigModelFromMap(m map[string]interface{}) *models.DeviceConfig {
//	adminState := m["admin_state"].(*models.AdminState) // AdminState
//	assetID := m["asset_id"].(string)
//	baseImage := m["base_image"].([]*models.BaseOSImage)          // []*BaseOSImage
//	baseOsRetryCounter := int64(m["base_os_retry_counter"].(int)) // int64 false false false
//	baseOsRetryTime := m["base_os_retry_time"].(string)
//	clientIP := m["client_ip"].(string)
//	clusterID := m["cluster_id"].(string)
//	configItem := m["config_item"].([]*models.EDConfigItem) // []*EDConfigItem
//	cpu := int64(m["cpu"].(int))                            // int64 false false false
//	var debugKnob *models.DebugKnobDetail                   // DebugKnobDetail
//	debugKnobInterface, debugKnobIsSet := m["debug_knob"]
//	if debugKnobIsSet {
//		debugKnobMap := debugKnobInterface.([]interface{})[0].(map[string]interface{})
//		debugKnob = DebugKnobDetailModelFromMap(debugKnobMap)
//	}
//	//
//	var defaultNetInst *models.NetInstConfig // NetInstConfig
//	defaultNetInstInterface, defaultNetInstIsSet := m["default_net_inst"]
//	if defaultNetInstIsSet {
//		defaultNetInstMap := defaultNetInstInterface.([]interface{})[0].(map[string]interface{})
//		defaultNetInst = NetInstConfigModelFromMap(defaultNetInstMap)
//	}
//	//
//	deploymentTag := m["deployment_tag"].(string)
//	deprecated := m["deprecated"].(string)
//	description := m["description"].(string)
//	var devLocation *models.GeoLocation // GeoLocation
//	devLocationInterface, devLocationIsSet := m["dev_location"]
//	if devLocationIsSet {
//		devLocationMap := devLocationInterface.([]interface{})[0].(map[string]interface{})
//		devLocation = GeoLocationModelFromMap(devLocationMap)
//	}
//	//
//	var dlisp *models.DeviceLisp // DeviceLisp
//	dlispInterface, dlispIsSet := m["dlisp"]
//	if dlispIsSet {
//		dlispMap := dlispInterface.([]interface{})[0].(map[string]interface{})
//		dlisp = DeviceLispModelFromMap(dlispMap)
//	}
//	//
//	var edgeviewconfig *models.EdgeviewCfg // EdgeviewCfg
//	edgeviewconfigInterface, edgeviewconfigIsSet := m["edgeviewconfig"]
//	if edgeviewconfigIsSet {
//		edgeviewconfigMap := edgeviewconfigInterface.([]interface{})[0].(map[string]interface{})
//		edgeviewconfig = EdgeviewCfgModelFromMap(edgeviewconfigMap)
//	}
//	//
//	generateSoftSerial := m["generate_soft_serial"].(bool)
//	id := m["id"].(string)
//	identity := m["identity"].(strfmt.Base64)
//	interfaces := m["interfaces"].([]*models.SysInterface) // []*SysInterface
//	location := m["location"].(string)
//	memory := int64(m["memory"].(int)) // int64 false false false
//	modelID := m["model_id"].(string)
//	name := m["name"].(string)
//	obkey := m["obkey"].(string)
//	var onboarding *models.DeviceCerts // DeviceCerts
//	onboardingInterface, onboardingIsSet := m["onboarding"]
//	if onboardingIsSet {
//		onboardingMap := onboardingInterface.([]interface{})[0].(map[string]interface{})
//		onboarding = DeviceCertsModelFromMap(onboardingMap)
//	}
//	//
//	preparePowerOffCounter := int64(m["prepare_power_off_counter"].(int)) // int64 false false false
//	preparePowerOffTime := m["prepare_power_off_time"].(string)
//	projectID := m["project_id"].(string)
//	resetCounter := int64(m["reset_counter"].(int)) // int64 false false false
//	resetTime := m["reset_time"].(string)
//	var revision *models.ObjectRevision // ObjectRevision
//	revisionInterface, revisionIsSet := m["revision"]
//	if revisionIsSet {
//		revisionMap := revisionInterface.([]interface{})[0].(map[string]interface{})
//		revision = ObjectRevisionModelFromMap(revisionMap)
//	}
//	//
//	serialno := m["serialno"].(string)
//	sitePictures := m["site_pictures"].([]string)
//	storage := int64(m["storage"].(int)) // int64 false false false
//	tags := m["tags"].(map[string]string)
//	thread := int64(m["thread"].(int)) // int64 false false false
//	title := m["title"].(string)
//	token := m["token"].(string)
//	utype := m["utype"].(*models.ModelArchType) // ModelArchType
//	return &models.DeviceConfig{
//		AdminState:             adminState,
//		AssetID:                assetID,
//		BaseImage:              baseImage,
//		BaseOsRetryCounter:     baseOsRetryCounter,
//		BaseOsRetryTime:        baseOsRetryTime,
//		ClientIP:               clientIP,
//		ClusterID:              clusterID,
//		ConfigItem:             configItem,
//		CPU:                    cpu,
//		DebugKnob:              debugKnob,
//		DefaultNetInst:         defaultNetInst,
//		DeploymentTag:          deploymentTag,
//		Deprecated:             deprecated,
//		Description:            description,
//		DevLocation:            devLocation,
//		Dlisp:                  dlisp,
//		Edgeviewconfig:         edgeviewconfig,
//		GenerateSoftSerial:     generateSoftSerial,
//		ID:                     id,
//		Identity:               identity,
//		Interfaces:             interfaces,
//		Location:               location,
//		Memory:                 memory,
//		ModelID:                &modelID,
//		Name:                   &name,
//		Obkey:                  obkey,
//		Onboarding:             onboarding,
//		PreparePowerOffCounter: preparePowerOffCounter,
//		PreparePowerOffTime:    preparePowerOffTime,
//		ProjectID:              &projectID,
//		ResetCounter:           resetCounter,
//		ResetTime:              resetTime,
//		Revision:               revision,
//		Serialno:               serialno,
//		SitePictures:           sitePictures,
//		Storage:                storage,
//		Tags:                   tags,
//		Thread:                 thread,
//		Title:                  &title,
//		Token:                  token,
//		Utype:                  utype,
//	}
//}

// Update the underlying DeviceConfig resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDeviceConfigResourceData(d *schema.ResourceData, m *models.DeviceConfig) {
	d.Set("admin_state", m.AdminState)
	d.Set("asset_id", m.AssetID)
	d.Set("base_image", SetBaseOSImageSubResourceData(m.BaseImage))
	d.Set("base_os_retry_counter", m.BaseOsRetryCounter)
	d.Set("base_os_retry_time", m.BaseOsRetryTime)
	d.Set("client_ip", m.ClientIP)
	d.Set("cluster_id", m.ClusterID)
	d.Set("config_item", SetEDConfigItemSubResourceData(m.ConfigItem))
	d.Set("cpu", m.CPU)
	d.Set("debug_knob", SetDebugKnobDetailSubResourceData([]*models.DebugKnobDetail{m.DebugKnob}))
	d.Set("default_net_inst", SetNetInstConfigSubResourceData([]*models.NetInstConfig{m.DefaultNetInst}))
	d.Set("deployment_tag", m.DeploymentTag)
	d.Set("deprecated", m.Deprecated)
	d.Set("description", m.Description)
	d.Set("dev_location", SetGeoLocationSubResourceData([]*models.GeoLocation{m.DevLocation}))
	d.Set("dlisp", SetDeviceLispSubResourceData([]*models.DeviceLisp{m.Dlisp}))
	d.Set("edgeviewconfig", SetEdgeviewCfgSubResourceData([]*models.EdgeviewCfg{m.Edgeviewconfig}))
	d.Set("generate_soft_serial", m.GenerateSoftSerial)
	d.Set("id", m.ID)
	d.Set("identity", m.Identity.String())
	d.Set("interfaces", SetSysInterfaceSubResourceData(m.Interfaces))
	d.Set("location", m.Location)
	d.Set("memory", m.Memory)
	d.Set("model_id", m.ModelID)
	d.Set("name", m.Name)
	d.Set("obkey", m.Obkey)
	d.Set("prepare_power_off_counter", m.PreparePowerOffCounter)
	d.Set("prepare_power_off_time", m.PreparePowerOffTime)
	d.Set("project_id", m.ProjectID)
	d.Set("reset_counter", m.ResetCounter)
	d.Set("reset_time", m.ResetTime)
	d.Set("revision", SetObjectRevisionSubResourceData([]*models.ObjectRevision{m.Revision}))
	d.Set("serialno", m.Serialno)
	d.Set("site_pictures", m.SitePictures)
	d.Set("storage", m.Storage)
	d.Set("tags", m.Tags)
	d.Set("thread", m.Thread)
	d.Set("title", m.Title)
	d.Set("token", m.Token)
	d.Set("utype", m.Utype)
}

// Iterate through and update the DeviceConfig resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDeviceConfigSubResourceData(m []*models.DeviceConfig) (d []*map[string]interface{}) {
	for _, DeviceConfigModel := range m {
		if DeviceConfigModel != nil {
			properties := make(map[string]interface{})
			properties["admin_state"] = DeviceConfigModel.AdminState
			properties["asset_id"] = DeviceConfigModel.AssetID
			properties["base_image"] = SetBaseOSImageSubResourceData(DeviceConfigModel.BaseImage)
			properties["base_os_retry_counter"] = DeviceConfigModel.BaseOsRetryCounter
			properties["base_os_retry_time"] = DeviceConfigModel.BaseOsRetryTime
			properties["client_ip"] = DeviceConfigModel.ClientIP
			properties["cluster_id"] = DeviceConfigModel.ClusterID
			properties["config_item"] = SetEDConfigItemSubResourceData(DeviceConfigModel.ConfigItem)
			properties["cpu"] = DeviceConfigModel.CPU
			properties["debug_knob"] = SetDebugKnobDetailSubResourceData([]*models.DebugKnobDetail{DeviceConfigModel.DebugKnob})
			properties["default_net_inst"] = SetNetInstConfigSubResourceData([]*models.NetInstConfig{DeviceConfigModel.DefaultNetInst})
			properties["deployment_tag"] = DeviceConfigModel.DeploymentTag
			properties["deprecated"] = DeviceConfigModel.Deprecated
			properties["description"] = DeviceConfigModel.Description
			properties["dev_location"] = SetGeoLocationSubResourceData([]*models.GeoLocation{DeviceConfigModel.DevLocation})
			properties["dlisp"] = SetDeviceLispSubResourceData([]*models.DeviceLisp{DeviceConfigModel.Dlisp})
			properties["edgeviewconfig"] = SetEdgeviewCfgSubResourceData([]*models.EdgeviewCfg{DeviceConfigModel.Edgeviewconfig})
			properties["generate_soft_serial"] = DeviceConfigModel.GenerateSoftSerial
			properties["id"] = DeviceConfigModel.ID
			properties["identity"] = DeviceConfigModel.Identity.String()
			properties["interfaces"] = SetSysInterfaceSubResourceData(DeviceConfigModel.Interfaces)
			properties["location"] = DeviceConfigModel.Location
			properties["memory"] = DeviceConfigModel.Memory
			properties["model_id"] = DeviceConfigModel.ModelID
			properties["name"] = DeviceConfigModel.Name
			properties["obkey"] = DeviceConfigModel.Obkey
			properties["prepare_power_off_counter"] = DeviceConfigModel.PreparePowerOffCounter
			properties["prepare_power_off_time"] = DeviceConfigModel.PreparePowerOffTime
			properties["project_id"] = DeviceConfigModel.ProjectID
			properties["reset_counter"] = DeviceConfigModel.ResetCounter
			properties["reset_time"] = DeviceConfigModel.ResetTime
			properties["revision"] = SetObjectRevisionSubResourceData([]*models.ObjectRevision{DeviceConfigModel.Revision})
			properties["serialno"] = DeviceConfigModel.Serialno
			properties["site_pictures"] = DeviceConfigModel.SitePictures
			properties["storage"] = DeviceConfigModel.Storage
			properties["tags"] = DeviceConfigModel.Tags
			properties["thread"] = DeviceConfigModel.Thread
			properties["title"] = DeviceConfigModel.Title
			properties["token"] = DeviceConfigModel.Token
			properties["utype"] = DeviceConfigModel.Utype
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the DeviceConfig resource defined in the Terraform configuration
func DeviceConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"admin_state": {
			Description: `administrative state of device`,
			Type:        schema.TypeString,
			Optional:    true,
			Default:     models.ADMINSTATE_CREATED, // FIXME: name properly in API spec
			DiffSuppressFunc: func(k, oldValue, newValue string, d *schema.ResourceData) bool {
				if oldValue == string(models.ADMINSTATE_REGISTERED) &&
					newValue == string(models.ADMINSTATE_ACTIVE) {
					return true
				}
				return false
			},
		},

		"asset_id": {
			Description: `Device asset ID`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"base_image": {
			Description: `base images`,
			Type:        schema.TypeList, //GoType: []*BaseOSImage
			Elem: &schema.Resource{
				Schema: BaseOSImageSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"base_os_retry_counter": {
			Description: `device baseos retry counter`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"base_os_retry_time": {
			Description: `device baseos retry time`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"client_ip": {
			Description: `Client IP`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"cluster_id": {
			Description: `System defined universally unique clusterInstance ID, unique across the enterprise.`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"config_item": {
			Description: `ED configurations`,
			Type:        schema.TypeList, //GoType: []*EDConfigItem
			Elem: &schema.Resource{
				Schema: EDConfigItemSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"cpu": {
			Description: `CPU (configured values)`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"debug_knob": {
			Description: `debug knob details for the device`,
			Type:        schema.TypeList, //GoType: DebugKnobDetail
			Elem: &schema.Resource{
				Schema: DebugKnobDetailSchema(),
			},
			Computed:  true,
			Sensitive: true,
		},

		"default_net_inst": {
			Description: `default network instance details`,
			Type:        schema.TypeList, //GoType: NetInstConfig
			Elem: &schema.Resource{
				Schema: NetInstConfigSchema(),
			},
			Optional: true,
		},

		"deployment_tag": {
			Description: `user defined tag for the device, which is used while deploying policies.`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"deprecated": {
			Description: `deprecated field`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"description": {
			Description: `user specified description`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"dev_location": {
			Description: `User specified geo location`,
			Type:        schema.TypeList, //GoType: GeoLocation
			Elem: &schema.Resource{
				Schema: GeoLocationSchema(),
			},
			Optional: true,
		},

		"dlisp": {
			Description: `device Lisp`,
			Type:        schema.TypeList, //GoType: DeviceLisp
			Elem: &schema.Resource{
				Schema: DeviceLispSchema(),
			},
			Optional: true,
		},

		"edgeviewconfig": {
			Description: `edgeview configuration for device`,
			Type:        schema.TypeList, //GoType: EdgeviewCfg
			Elem: &schema.Resource{
				Schema: EdgeviewCfgSchema(),
			},
			Optional: true,
		},

		"generate_soft_serial": {
			Description: `indicates whether a soft serial should be generated; it will work ONLY when device is created`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"id": {
			Description: `system generated unique id for a device`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"identity": {
			Description: `Device identity`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"interfaces": {
			Description: `System Interface list`,
			Type:        schema.TypeList, //GoType: []*SysInterface
			Elem: &schema.Resource{
				Schema: SysInterfaceSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"location": {
			Description: `Device location: deprecated`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"memory": {
			Description: `Device memory in MBs`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"model_id": {
			Description: `device model`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"name": {
			Description: `user specified device name`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"obkey": {
			Description: `Object key`,
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			DiffSuppressFunc: func(k, oldValue, newValue string, d *schema.ResourceData) bool {
				if oldValue == "" {
					return false
				}
				parts := strings.Split(newValue, ":")
				if len(parts) != 2 {
					return true
				}
				if parts[0] != oldValue {
					return true
				}
				return false
			},
		},

		"prepare_power_off_counter": {
			Description: `prepare poweroff counter`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"prepare_power_off_time": {
			Description: `prepare poweroff time`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"project_id": {
			Description: `project name`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"reset_counter": {
			Description: `devicereset counter`,
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
		},

		"reset_time": {
			Description: `device reset time`,
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
		},

		"revision": {
			Description: `Object revision details`,
			Type:        schema.TypeList, //GoType: ObjectRevision
			Elem: &schema.Resource{
				Schema: ObjectRevisionSchema(),
			},
			Optional: true,
			Computed: true,
		},

		"serialno": {
			Description: `Device serial number`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"site_pictures": {
			Description: `Site captured pictures`,
			Type:        schema.TypeList, //GoType: []string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"storage": {
			Description: `Device storage in GBs`,
			Type:        schema.TypeInt,
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

		"thread": {
			Description: `Threads`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"title": {
			Description: `user specified title`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"token": {
			Description: `Single use token`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"utype": {
			Description: `device model arch type`,
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
		},
	}
}

// Retrieve property field names for updating the DeviceConfig resource
func GetDeviceConfigPropertyFields() (t []string) {
	return []string{
		"admin_state",
		"asset_id",
		"base_image",
		"base_os_retry_counter",
		"base_os_retry_time",
		"client_ip",
		"cluster_id",
		"config_item",
		"cpu",
		"debug_knob",
		"default_net_inst",
		"deployment_tag",
		"deprecated",
		"description",
		"dev_location",
		"dlisp",
		"edgeviewconfig",
		"generate_soft_serial",
		"id",
		"identity",
		"interfaces",
		"location",
		"memory",
		"model_id",
		"name",
		"obkey",
		"onboarding",
		"prepare_power_off_counter",
		"prepare_power_off_time",
		"project_id",
		"reset_counter",
		"reset_time",
		"revision",
		"serialno",
		"site_pictures",
		"storage",
		"tags",
		"thread",
		"title",
		"token",
		"utype",
	}
}
