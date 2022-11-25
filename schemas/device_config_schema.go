package schemas

import (
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate DeviceConfig resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func DeviceConfigModel(d *schema.ResourceData) *models.DeviceConfig {
	adminState := d.Get("admin_state").(*models.AdminState) // AdminState
	assetID := d.Get("asset_id").(string)
	baseImage := d.Get("base_image").([]*models.BaseOSImage) // []*BaseOSImage
	clientIP := d.Get("client_ip").(string)
	clusterID := d.Get("cluster_id").(string)
	configItem := d.Get("config_item").([]*models.EDConfigItem) // []*EDConfigItem
	cpu := int64(d.Get("cpu").(int))
	deprecated := d.Get("deprecated").(string)
	description := d.Get("description").(string)
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
	id := d.Get("id").(string)
	identity := d.Get("identity").(strfmt.Base64)
	interfaces := d.Get("interfaces").([]*models.SysInterface) // []*SysInterface
	location := d.Get("location").(string)
	memory := int64(d.Get("memory").(int))
	modelID := d.Get("model_id").(string)
	name := d.Get("name").(string)
	obkey := d.Get("obkey").(string)
	var onboarding *models.DeviceCerts // DeviceCerts
	onboardingInterface, onboardingIsSet := d.GetOk("onboarding")
	if onboardingIsSet {
		onboardingMap := onboardingInterface.([]interface{})[0].(map[string]interface{})
		onboarding = DeviceCertsModelFromMap(onboardingMap)
	}
	projectID := d.Get("project_id").(string)
	resetCounter := int64(d.Get("reset_counter").(int))
	resetTime := d.Get("reset_time").(string)
	var revision *models.ObjectRevision // ObjectRevision
	revisionInterface, revisionIsSet := d.GetOk("revision")
	if revisionIsSet {
		revisionMap := revisionInterface.([]interface{})[0].(map[string]interface{})
		revision = ObjectRevisionModelFromMap(revisionMap)
	}
	serialno := d.Get("serialno").(string)
	sitePictures := d.Get("site_pictures").([]string)
	storage := int64(d.Get("storage").(int))
	tags := d.Get("tags").(map[string]string) // map[string]string
	thread := int64(d.Get("thread").(int))
	title := d.Get("title").(string)
	token := d.Get("token").(string)
	utype := d.Get("utype").(*models.ModelArchType) // ModelArchType
	return &models.DeviceConfig{
		AdminState:   adminState,
		AssetID:      assetID,
		BaseImage:    baseImage,
		ClientIP:     clientIP,
		ClusterID:    clusterID,
		ConfigItem:   configItem,
		CPU:          cpu,
		Deprecated:   deprecated,
		Description:  description,
		DevLocation:  devLocation,
		Dlisp:        dlisp,
		ID:           id,
		Identity:     identity,
		Interfaces:   interfaces,
		Location:     location,
		Memory:       memory,
		ModelID:      &modelID, // string true false false
		Name:         &name,    // string true false false
		Obkey:        obkey,
		Onboarding:   onboarding,
		ProjectID:    &projectID, // string true false false
		ResetCounter: resetCounter,
		ResetTime:    resetTime,
		Revision:     revision,
		Serialno:     serialno,
		SitePictures: sitePictures,
		Storage:      storage,
		Tags:         tags,
		Thread:       thread,
		Title:        &title, // string true false false
		Token:        token,
		Utype:        utype,
	}
}

func DeviceConfigModelFromMap(m map[string]interface{}) *models.DeviceConfig {
	adminState := m["admin_state"].(*models.AdminState) // AdminState
	assetID := m["asset_id"].(string)
	baseImage := m["base_image"].([]*models.BaseOSImage) // []*BaseOSImage
	clientIP := m["client_ip"].(string)
	clusterID := m["cluster_id"].(string)
	configItem := m["config_item"].([]*models.EDConfigItem) // []*EDConfigItem
	cpu := int64(m["cpu"].(int))                            // int64 false false false
	deprecated := m["deprecated"].(string)
	description := m["description"].(string)
	var devLocation *models.GeoLocation // GeoLocation
	devLocationInterface, devLocationIsSet := m["dev_location"]
	if devLocationIsSet {
		devLocationMap := devLocationInterface.([]interface{})[0].(map[string]interface{})
		devLocation = GeoLocationModelFromMap(devLocationMap)
	}
	//
	var dlisp *models.DeviceLisp // DeviceLisp
	dlispInterface, dlispIsSet := m["dlisp"]
	if dlispIsSet {
		dlispMap := dlispInterface.([]interface{})[0].(map[string]interface{})
		dlisp = DeviceLispModelFromMap(dlispMap)
	}
	//
	id := m["id"].(string)
	identity := m["identity"].(strfmt.Base64)
	interfaces := m["interfaces"].([]*models.SysInterface) // []*SysInterface
	location := m["location"].(string)
	memory := int64(m["memory"].(int)) // int64 false false false
	modelID := m["model_id"].(string)
	name := m["name"].(string)
	obkey := m["obkey"].(string)
	var onboarding *models.DeviceCerts // DeviceCerts
	onboardingInterface, onboardingIsSet := m["onboarding"]
	if onboardingIsSet {
		onboardingMap := onboardingInterface.([]interface{})[0].(map[string]interface{})
		onboarding = DeviceCertsModelFromMap(onboardingMap)
	}
	//
	projectID := m["project_id"].(string)
	resetCounter := int64(m["reset_counter"].(int)) // int64 false false false
	resetTime := m["reset_time"].(string)
	var revision *models.ObjectRevision // ObjectRevision
	revisionInterface, revisionIsSet := m["revision"]
	if revisionIsSet {
		revisionMap := revisionInterface.([]interface{})[0].(map[string]interface{})
		revision = ObjectRevisionModelFromMap(revisionMap)
	}
	//
	serialno := m["serialno"].(string)
	sitePictures := m["site_pictures"].([]string)
	storage := int64(m["storage"].(int)) // int64 false false false
	tags := m["tags"].(map[string]string)
	thread := int64(m["thread"].(int)) // int64 false false false
	title := m["title"].(string)
	token := m["token"].(string)
	utype := m["utype"].(*models.ModelArchType) // ModelArchType
	return &models.DeviceConfig{
		AdminState:   adminState,
		AssetID:      assetID,
		BaseImage:    baseImage,
		ClientIP:     clientIP,
		ClusterID:    clusterID,
		ConfigItem:   configItem,
		CPU:          cpu,
		Deprecated:   deprecated,
		Description:  description,
		DevLocation:  devLocation,
		Dlisp:        dlisp,
		ID:           id,
		Identity:     identity,
		Interfaces:   interfaces,
		Location:     location,
		Memory:       memory,
		ModelID:      &modelID,
		Name:         &name,
		Obkey:        obkey,
		Onboarding:   onboarding,
		ProjectID:    &projectID,
		ResetCounter: resetCounter,
		ResetTime:    resetTime,
		Revision:     revision,
		Serialno:     serialno,
		SitePictures: sitePictures,
		Storage:      storage,
		Tags:         tags,
		Thread:       thread,
		Title:        &title,
		Token:        token,
		Utype:        utype,
	}
}

// Update the underlying DeviceConfig resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDeviceConfigResourceData(d *schema.ResourceData, m *models.DeviceConfig) {
	d.Set("admin_state", m.AdminState)
	d.Set("asset_id", m.AssetID)
	d.Set("base_image", SetBaseOSImageSubResourceData(m.BaseImage))
	d.Set("client_ip", m.ClientIP)
	d.Set("cluster_id", m.ClusterID)
	d.Set("config_item", SetEDConfigItemSubResourceData(m.ConfigItem))
	d.Set("cpu", m.CPU)
	d.Set("deprecated", m.Deprecated)
	d.Set("description", m.Description)
	d.Set("dev_location", SetGeoLocationSubResourceData([]*models.GeoLocation{m.DevLocation}))
	d.Set("dlisp", SetDeviceLispSubResourceData([]*models.DeviceLisp{m.Dlisp}))
	d.Set("id", m.ID)
	d.Set("identity", m.Identity)
	d.Set("interfaces", SetSysInterfaceSubResourceData(m.Interfaces))
	d.Set("location", m.Location)
	d.Set("memory", m.Memory)
	d.Set("model_id", m.ModelID)
	d.Set("name", m.Name)
	d.Set("obkey", m.Obkey)
	d.Set("onboarding", SetDeviceCertsSubResourceData([]*models.DeviceCerts{m.Onboarding}))
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

// Iterate throught and update the DeviceConfig resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDeviceConfigSubResourceData(m []*models.DeviceConfig) (d []*map[string]interface{}) {
	for _, DeviceConfigModel := range m {
		if DeviceConfigModel != nil {
			properties := make(map[string]interface{})
			properties["admin_state"] = DeviceConfigModel.AdminState
			properties["asset_id"] = DeviceConfigModel.AssetID
			properties["base_image"] = SetBaseOSImageSubResourceData(DeviceConfigModel.BaseImage)
			properties["client_ip"] = DeviceConfigModel.ClientIP
			properties["cluster_id"] = DeviceConfigModel.ClusterID
			properties["config_item"] = SetEDConfigItemSubResourceData(DeviceConfigModel.ConfigItem)
			properties["cpu"] = DeviceConfigModel.CPU
			properties["deprecated"] = DeviceConfigModel.Deprecated
			properties["description"] = DeviceConfigModel.Description
			properties["dev_location"] = SetGeoLocationSubResourceData([]*models.GeoLocation{DeviceConfigModel.DevLocation})
			properties["dlisp"] = SetDeviceLispSubResourceData([]*models.DeviceLisp{DeviceConfigModel.Dlisp})
			properties["id"] = DeviceConfigModel.ID
			properties["identity"] = DeviceConfigModel.Identity
			properties["interfaces"] = SetSysInterfaceSubResourceData(DeviceConfigModel.Interfaces)
			properties["location"] = DeviceConfigModel.Location
			properties["memory"] = DeviceConfigModel.Memory
			properties["model_id"] = DeviceConfigModel.ModelID
			properties["name"] = DeviceConfigModel.Name
			properties["obkey"] = DeviceConfigModel.Obkey
			properties["onboarding"] = SetDeviceCertsSubResourceData([]*models.DeviceCerts{DeviceConfigModel.Onboarding})
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
			Type:        schema.TypeList, //GoType: AdminState
			Elem: &schema.Resource{
				Schema: AdminStateSchema(),
			},
			Optional: true,
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
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
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
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"cpu": {
			Description: `CPU (configured values)`,
			Type:        schema.TypeInt,
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

		"id": {
			Description: `system generated unique id for a device`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"identity": {
			Description:  `Device identity`,
			Type:         schema.TypeString,
			ValidateFunc: validation.IsRFC3339Time,
			Optional:     true,
		},

		"interfaces": {
			Description: `System Interface list`,
			Type:        schema.TypeList, //GoType: []*SysInterface
			Elem: &schema.Resource{
				Schema: SysInterfaceSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
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
		},

		"onboarding": {
			Description: `Device level certificates used while onboarding`,
			Type:        schema.TypeList, //GoType: DeviceCerts
			Elem: &schema.Resource{
				Schema: DeviceCertsSchema(),
			},
			Optional: true,
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
		},

		"reset_time": {
			Description: `device reset time`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"revision": {
			Description: `Object revision details`,
			Type:        schema.TypeList, //GoType: ObjectRevision
			Elem: &schema.Resource{
				Schema: ObjectRevisionSchema(),
			},
			Optional: true,
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
			Description: ``,
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
			Type:        schema.TypeList, //GoType: ModelArchType
			Elem: &schema.Resource{
				Schema: ModelArchTypeSchema(),
			},
			Optional: true,
		},
	}
}

// Retrieve property field names for updating the DeviceConfig resource
func GetDeviceConfigPropertyFields() (t []string) {
	return []string{
		"admin_state",
		"asset_id",
		"base_image",
		"client_ip",
		"cluster_id",
		"config_item",
		"cpu",
		"deprecated",
		"description",
		"dev_location",
		"dlisp",
		"id",
		"identity",
		"interfaces",
		"location",
		"memory",
		"model_id",
		"name",
		"obkey",
		"onboarding",
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
