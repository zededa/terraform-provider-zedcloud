package schemas

import (
	"strings"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

func NodeModel(d *schema.ResourceData) *models.Node {
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
		var items []interface{}
		if listItems, isList := configItemInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = configItemInterface.(*schema.Set).List()
		}
		for _, v := range items {
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
	var defaultNetInst *models.NetworkInstance // NetInstConfig
	defaultNetInstInterface, defaultNetInstIsSet := d.GetOk("default_net_inst")
	if defaultNetInstIsSet {
		defaultNetInstMap := defaultNetInstInterface.([]interface{})[0].(map[string]interface{})
		defaultNetInst = NetworkInstanceModelFromMap(defaultNetInstMap)
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
	var interfaces []*models.SystemInterface // []*SysInterface
	interfacesInterface, interfacesIsSet := d.GetOk("interfaces")
	if interfacesIsSet {
		var items []interface{}
		if listItems, isList := interfacesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = interfacesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := SystemInterfaceModelFromMap(v.(map[string]interface{}))
			interfaces = append(interfaces, m)
		}
	}
	location, _ := d.Get("location").(string)
	memoryInt, _ := d.Get("memory").(int)
	memory := int64(memoryInt)
	modelID, _ := d.Get("model_id").(string)
	name, _ := d.Get("name").(string)
	obkey, _ := d.Get("onboarding_key").(string)
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
	return &models.Node{
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
		ModelID:                &modelID,
		Name:                   &name,
		Obkey:                  obkey,
		Onboarding:             onboarding,
		PreparePowerOffCounter: preparePowerOffCounter,
		PreparePowerOffTime:    preparePowerOffTime,
		ProjectID:              &projectID,
		ResetCounter:           resetCounter,
		ResetTime:              resetTime,
		Revision:               revision,
		Serialno:               serialno,
		SitePictures:           sitePictures,
		Storage:                storage,
		Tags:                   tags,
		Thread:                 thread,
		Title:                  &title,
		Token:                  token,
		Utype:                  utype,
	}
}

func EdgeNodeModelFromMap(m map[string]interface{}) *models.Node {
	var adminState *models.AdminState // AdminState
	adminStateInterface, adminStateIsSet := m["admin_state"]
	if adminStateIsSet {
		adminStateModel := adminStateInterface.(string)
		adminState = models.NewAdminState(models.AdminState(adminStateModel))
	}
	assetID := m["asset_id"].(string)
	var baseImage []*models.BaseOSImage // []*BaseOSImage
	baseImageInterface, baseImageIsSet := m["base_image"]
	if baseImageIsSet {
		var items []interface{}
		if listItems, isList := baseImageInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = baseImageInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := BaseOSImageModelFromMap(v.(map[string]interface{}))
			baseImage = append(baseImage, m)
		}
	}
	baseOsRetryCounter := int64(m["base_os_retry_counter"].(int)) // int64
	baseOsRetryTime := m["base_os_retry_time"].(string)
	clientIP := m["client_ip"].(string)
	clusterID := m["cluster_id"].(string)
	var configItem []*models.EDConfigItem // []*EDConfigItem
	configItemInterface, configItemIsSet := m["config_item"]
	if configItemIsSet {
		var items []interface{}
		if listItems, isList := configItemInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = configItemInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := EDConfigItemModelFromMap(v.(map[string]interface{}))
			configItem = append(configItem, m)
		}
	}
	cpu := int64(m["cpu"].(int))          // int64
	var debugKnob *models.DebugKnobDetail // DebugKnobDetail
	debugKnobInterface, debugKnobIsSet := m["debug_knob"]
	if debugKnobIsSet && debugKnobInterface != nil {
		debugKnobMap := debugKnobInterface.([]interface{})
		if len(debugKnobMap) > 0 {
			debugKnob = DebugKnobDetailModelFromMap(debugKnobMap[0].(map[string]interface{}))
		}
	}
	var defaultNetInst *models.NetworkInstance // NetInstConfig
	defaultNetInstInterface, defaultNetInstIsSet := m["default_net_inst"]
	if defaultNetInstIsSet && defaultNetInstInterface != nil {
		defaultNetInstMap := defaultNetInstInterface.([]interface{})
		if len(defaultNetInstMap) > 0 {
			defaultNetInst = NetworkInstanceModelFromMap(defaultNetInstMap[0].(map[string]interface{}))
		}
	}
	deploymentTag := m["deployment_tag"].(string)
	deprecated := m["deprecated"].(string)
	description := m["description"].(string)
	var devLocation *models.GeoLocation // GeoLocation
	devLocationInterface, devLocationIsSet := m["dev_location"]
	if devLocationIsSet && devLocationInterface != nil {
		devLocationMap := devLocationInterface.([]interface{})
		if len(devLocationMap) > 0 {
			devLocation = GeoLocationModelFromMap(devLocationMap[0].(map[string]interface{}))
		}
	}
	var dlisp *models.DeviceLisp // DeviceLisp
	dlispInterface, dlispIsSet := m["dlisp"]
	if dlispIsSet && dlispInterface != nil {
		dlispMap := dlispInterface.([]interface{})
		if len(dlispMap) > 0 {
			dlisp = DeviceLispModelFromMap(dlispMap[0].(map[string]interface{}))
		}
	}
	var edgeviewconfig *models.EdgeviewCfg // EdgeviewCfg
	edgeviewconfigInterface, edgeviewconfigIsSet := m["edgeviewconfig"]
	if edgeviewconfigIsSet && edgeviewconfigInterface != nil {
		edgeviewconfigMap := edgeviewconfigInterface.([]interface{})
		if len(edgeviewconfigMap) > 0 {
			edgeviewconfig = EdgeviewCfgModelFromMap(edgeviewconfigMap[0].(map[string]interface{}))
		}
	}
	generateSoftSerial := m["generate_soft_serial"].(bool)
	id := m["id"].(string)
	identity := m["identity"].(strfmt.Base64)
	var interfaces []*models.SystemInterface // []*SysInterface
	interfacesInterface, interfacesIsSet := m["interfaces"]
	if interfacesIsSet {
		var items []interface{}
		if listItems, isList := interfacesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = interfacesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := SystemInterfaceModelFromMap(v.(map[string]interface{}))
			interfaces = append(interfaces, m)
		}
	}
	location := m["location"].(string)
	memory := int64(m["memory"].(int)) // int64
	modelID := m["model_id"].(string)
	name := m["name"].(string)
	obkey := m["onboarding_key"].(string)
	var onboarding *models.DeviceCerts // DeviceCerts
	onboardingInterface, onboardingIsSet := m["onboarding"]
	if onboardingIsSet && onboardingInterface != nil {
		onboardingMap := onboardingInterface.([]interface{})
		if len(onboardingMap) > 0 {
			onboarding = DeviceCertsModelFromMap(onboardingMap[0].(map[string]interface{}))
		}
	}
	preparePowerOffCounter := int64(m["prepare_power_off_counter"].(int)) // int64
	preparePowerOffTime := m["prepare_power_off_time"].(string)
	projectID := m["project_id"].(string)
	resetCounter := int64(m["reset_counter"].(int)) // int64
	resetTime := m["reset_time"].(string)
	var revision *models.ObjectRevision // ObjectRevision
	revisionInterface, revisionIsSet := m["revision"]
	if revisionIsSet && revisionInterface != nil {
		revisionMap := revisionInterface.([]interface{})
		if len(revisionMap) > 0 {
			revision = ObjectRevisionModelFromMap(revisionMap[0].(map[string]interface{}))
		}
	}
	serialno := m["serialno"].(string)
	var sitePictures []string
	sitePicturesInterface, sitePicturesIsSet := m["sitePictures"]
	if sitePicturesIsSet {
		sitePicturesSlice := sitePicturesInterface.([]interface{})
		for _, i := range sitePicturesSlice {
			sitePicturesSlice = append(sitePicturesSlice, i.(string))
		}
	}
	storage := int64(m["storage"].(int)) // int64
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

	thread := int64(m["thread"].(int)) // int64
	title := m["title"].(string)
	token := m["token"].(string)
	var utype *models.ModelArchType // ModelArchType
	utypeInterface, utypeIsSet := m["utype"]
	if utypeIsSet {
		utypeModel := utypeInterface.(string)
		utype = models.NewModelArchType(models.ModelArchType(utypeModel))
	}
	return &models.Node{
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
		ModelID:                &modelID,
		Name:                   &name,
		Obkey:                  obkey,
		Onboarding:             onboarding,
		PreparePowerOffCounter: preparePowerOffCounter,
		PreparePowerOffTime:    preparePowerOffTime,
		ProjectID:              &projectID,
		ResetCounter:           resetCounter,
		ResetTime:              resetTime,
		Revision:               revision,
		Serialno:               serialno,
		SitePictures:           sitePictures,
		Storage:                storage,
		Tags:                   tags,
		Thread:                 thread,
		Title:                  &title,
		Token:                  token,
		Utype:                  utype,
	}
}

func SetNodeResourceData(d *schema.ResourceData, m *models.Node) {
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
	d.Set("default_net_inst", SetNetworkInstanceSubResourceData([]*models.NetworkInstance{m.DefaultNetInst}))
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
	d.Set("onboarding_key", m.Obkey)
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

func SetEdgeNodeSubResourceData(m []*models.Node) (d []*map[string]interface{}) {
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
			properties["default_net_inst"] = SetNetworkInstanceSubResourceData([]*models.NetworkInstance{DeviceConfigModel.DefaultNetInst})
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
			properties["onboarding_key"] = DeviceConfigModel.Obkey
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
func Node() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"admin_state": {
			Description: `administrative state of device`,
			Type:        schema.TypeString,
			Optional:    true,
			Default:     models.ADMINSTATE_CREATED,
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
				Schema: BaseOSImage(),
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
				Schema: EDConfigItem(),
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
				Schema: DebugKnobDetail(),
			},
			Computed:  true,
			Sensitive: true,
		},

		"default_net_inst": {
			Description: `default network instance details`,
			Type:        schema.TypeList, //GoType: NetInstConfig
			Elem: &schema.Resource{
				Schema: NetworkInstance(),
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
				Schema: DeviceLisp(),
			},
			Optional: true,
		},

		"edgeviewconfig": {
			Description: `edgeview configuration for device`,
			Type:        schema.TypeList, //GoType: EdgeviewCfg
			Elem: &schema.Resource{
				Schema: EdgeView(),
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
			Computed:    true,
		},

		"interfaces": {
			Description: `System Interface list`,
			Type:        schema.TypeList, //GoType: []*SysInterface
			Elem: &schema.Resource{
				Schema: SystemInterface(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional:         true,
			DiffSuppressFunc: diffSuppressSystemInterfaceListOrder("interfaces"),
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
			ForceNew:    true,
		},

		"onboarding_key": {
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
			Optional:    true,
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
				Schema: ObjectRevision(),
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
			Type:        schema.TypeSet, //GoType: []string
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
func GetEdgeNodePropertyFields() (t []string) {
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
		"onboarding_key",
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
