package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

func AppInstConfigModel(d *schema.ResourceData) *models.AppInstConfig {
	bundleID, _ := d.Get("bundle_id").(string)
	bundleVersionInt, _ := d.Get("bundle_version").(int)
	bundleVersion := int64(bundleVersionInt)
	cpusInt, _ := d.Get("cpus").(int)
	cpus := int64(cpusInt)
	var interfaces []*models.AppInterface // []*AppInterface
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
			m := AppInterfaceModelFromMap(v.(map[string]interface{}))
			interfaces = append(interfaces, m)
		}
	}
	var logs *models.AppInstanceLogs // AppInstanceLogs
	logsInterface, logsIsSet := d.GetOk("logs")
	if logsIsSet && logsInterface != nil {
		logsMap := logsInterface.([]interface{})
		if len(logsMap) > 0 {
			logs = AppInstanceLogsModelFromMap(logsMap[0].(map[string]interface{}))
		}
	}
	var manifestJSON *models.VMManifest // VMManifest
	manifestJSONInterface, manifestJSONIsSet := d.GetOk("manifest_json")
	if manifestJSONIsSet && manifestJSONInterface != nil {
		manifestJSONMap := manifestJSONInterface.([]interface{})
		if len(manifestJSONMap) > 0 {
			manifestJSON = VMManifestModelFromMap(manifestJSONMap[0].(map[string]interface{}))
		}
	}
	memoryInt, _ := d.Get("memory").(int)
	memory := int64(memoryInt)
	nameAppPart, _ := d.Get("name_app_part").(string)
	nameProjectPart, _ := d.Get("name_project_part").(string)
	var namingScheme *models.AppNamingSchemeV2 // AppNamingSchemeV2
	namingSchemeInterface, namingSchemeIsSet := d.GetOk("naming_scheme")
	if namingSchemeIsSet {
		namingSchemeModel := namingSchemeInterface.(string)
		namingScheme = models.NewAppNamingSchemeV2(models.AppNamingSchemeV2(namingSchemeModel))
	}
	networksInt, _ := d.Get("networks").(int)
	networks := int64(networksInt)
	newBundleVersionAvailable, _ := d.Get("new_bundle_version_available").(bool)
	var originType *models.Origin // Origin
	originTypeInterface, originTypeIsSet := d.GetOk("origin_type")
	if originTypeIsSet {
		originTypeModel := originTypeInterface.(string)
		originType = models.NewOrigin(models.Origin(originTypeModel))
	}
	var parentDetail *models.ObjectParentDetail // ObjectParentDetail
	parentDetailInterface, parentDetailIsSet := d.GetOk("parent_detail")
	if parentDetailIsSet && parentDetailInterface != nil {
		parentDetailMap := parentDetailInterface.([]interface{})
		if len(parentDetailMap) > 0 {
			parentDetail = ObjectParentDetailModelFromMap(parentDetailMap[0].(map[string]interface{}))
		}
	}
	remoteConsole, _ := d.Get("remote_console").(bool)
	startDelayInSecondsInt, _ := d.Get("start_delay_in_seconds").(int)
	startDelayInSeconds := int64(startDelayInSecondsInt)
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

	var vminfo *models.VM // VM
	vminfoInterface, vminfoIsSet := d.GetOk("vminfo")
	if vminfoIsSet && vminfoInterface != nil {
		vminfoMap := vminfoInterface.([]interface{})
		if len(vminfoMap) > 0 {
			vminfo = VMModelFromMap(vminfoMap[0].(map[string]interface{}))
		}
	}
	return &models.AppInstConfig{
		BundleID:                  bundleID,
		BundleVersion:             bundleVersion,
		Cpus:                      cpus,
		Interfaces:                interfaces,
		Logs:                      logs,
		ManifestJSON:              manifestJSON,
		Memory:                    memory,
		NameAppPart:               nameAppPart,
		NameProjectPart:           nameProjectPart,
		NamingScheme:              namingScheme,
		Networks:                  networks,
		NewBundleVersionAvailable: newBundleVersionAvailable,
		OriginType:                originType,
		ParentDetail:              parentDetail,
		RemoteConsole:             remoteConsole,
		StartDelayInSeconds:       startDelayInSeconds,
		Storage:                   storage,
		Tags:                      tags,
		Vminfo:                    vminfo,
	}
}

func AppInstConfigModelFromMap(m map[string]interface{}) *models.AppInstConfig {
	bundleID := m["bundle_id"].(string)
	bundleVersion := int64(m["bundle_version"].(int)) // int64
	cpus := int64(m["cpus"].(int))                    // int64
	var interfaces []*models.AppInterface             // []*AppInterface
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
			m := AppInterfaceModelFromMap(v.(map[string]interface{}))
			interfaces = append(interfaces, m)
		}
	}
	var logs *models.AppInstanceLogs // AppInstanceLogs
	logsInterface, logsIsSet := m["logs"]
	if logsIsSet && logsInterface != nil {
		logsMap := logsInterface.([]interface{})
		if len(logsMap) > 0 {
			logs = AppInstanceLogsModelFromMap(logsMap[0].(map[string]interface{}))
		}
	}
	//
	var manifestJSON *models.VMManifest // VMManifest
	manifestJSONInterface, manifestJSONIsSet := m["manifest_json"]
	if manifestJSONIsSet && manifestJSONInterface != nil {
		manifestJSONMap := manifestJSONInterface.([]interface{})
		if len(manifestJSONMap) > 0 {
			manifestJSON = VMManifestModelFromMap(manifestJSONMap[0].(map[string]interface{}))
		}
	}
	//
	memory := int64(m["memory"].(int)) // int64
	nameAppPart := m["name_app_part"].(string)
	nameProjectPart := m["name_project_part"].(string)
	var namingScheme *models.AppNamingSchemeV2 // AppNamingSchemeV2
	namingSchemeInterface, namingSchemeIsSet := m["naming_scheme"]
	if namingSchemeIsSet {
		namingSchemeModel := namingSchemeInterface.(string)
		namingScheme = models.NewAppNamingSchemeV2(models.AppNamingSchemeV2(namingSchemeModel))
	}
	networks := int64(m["networks"].(int)) // int64
	newBundleVersionAvailable := m["new_bundle_version_available"].(bool)
	var originType *models.Origin // Origin
	originTypeInterface, originTypeIsSet := m["origin_type"]
	if originTypeIsSet {
		originTypeModel := originTypeInterface.(string)
		originType = models.NewOrigin(models.Origin(originTypeModel))
	}
	var parentDetail *models.ObjectParentDetail // ObjectParentDetail
	parentDetailInterface, parentDetailIsSet := m["parent_detail"]
	if parentDetailIsSet && parentDetailInterface != nil {
		parentDetailMap := parentDetailInterface.([]interface{})
		if len(parentDetailMap) > 0 {
			parentDetail = ObjectParentDetailModelFromMap(parentDetailMap[0].(map[string]interface{}))
		}
	}
	//
	remoteConsole := m["remote_console"].(bool)
	startDelayInSeconds := int64(m["start_delay_in_seconds"].(int)) // int64
	storage := int64(m["storage"].(int))                            // int64
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

	var vminfo *models.VM // VM
	vminfoInterface, vminfoIsSet := m["vminfo"]
	if vminfoIsSet && vminfoInterface != nil {
		vminfoMap := vminfoInterface.([]interface{})
		if len(vminfoMap) > 0 {
			vminfo = VMModelFromMap(vminfoMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.AppInstConfig{
		BundleID:                  bundleID,
		BundleVersion:             bundleVersion,
		Cpus:                      cpus,
		Interfaces:                interfaces,
		Logs:                      logs,
		ManifestJSON:              manifestJSON,
		Memory:                    memory,
		NameAppPart:               nameAppPart,
		NameProjectPart:           nameProjectPart,
		NamingScheme:              namingScheme,
		Networks:                  networks,
		NewBundleVersionAvailable: newBundleVersionAvailable,
		OriginType:                originType,
		ParentDetail:              parentDetail,
		RemoteConsole:             remoteConsole,
		StartDelayInSeconds:       startDelayInSeconds,
		Storage:                   storage,
		Tags:                      tags,
		Vminfo:                    vminfo,
	}
}

func SetAppInstConfigResourceData(d *schema.ResourceData, m *models.AppInstConfig) {
	d.Set("bundle_id", m.BundleID)
	d.Set("bundle_version", m.BundleVersion)
	d.Set("cpus", m.Cpus)
	d.Set("drives", m.Drives)
	d.Set("interfaces", SetAppInterfaceSubResourceData(m.Interfaces))
	d.Set("logs", SetAppInstanceLogsSubResourceData([]*models.AppInstanceLogs{m.Logs}))
	d.Set("manifest_json", SetVMManifestSubResourceData([]*models.VMManifest{m.ManifestJSON}))
	d.Set("memory", m.Memory)
	d.Set("name_app_part", m.NameAppPart)
	d.Set("name_project_part", m.NameProjectPart)
	d.Set("naming_scheme", m.NamingScheme)
	d.Set("networks", m.Networks)
	d.Set("new_bundle_version_available", m.NewBundleVersionAvailable)
	d.Set("origin_type", m.OriginType)
	d.Set("parent_detail", SetObjectParentDetailSubResourceData([]*models.ObjectParentDetail{m.ParentDetail}))
	d.Set("remote_console", m.RemoteConsole)
	d.Set("start_delay_in_seconds", m.StartDelayInSeconds)
	d.Set("storage", m.Storage)
	d.Set("tags", m.Tags)
	d.Set("vminfo", SetVMSubResourceData([]*models.VM{m.Vminfo}))
}

func SetAppInstConfigSubResourceData(m []*models.AppInstConfig) (d []*map[string]interface{}) {
	for _, AppInstConfigModel := range m {
		if AppInstConfigModel != nil {
			properties := make(map[string]interface{})
			properties["bundle_id"] = AppInstConfigModel.BundleID
			properties["bundle_version"] = AppInstConfigModel.BundleVersion
			properties["cpus"] = AppInstConfigModel.Cpus
			properties["drives"] = AppInstConfigModel.Drives
			properties["interfaces"] = SetAppInterfaceSubResourceData(AppInstConfigModel.Interfaces)
			properties["logs"] = SetAppInstanceLogsSubResourceData([]*models.AppInstanceLogs{AppInstConfigModel.Logs})
			properties["manifest_json"] = SetVMManifestSubResourceData([]*models.VMManifest{AppInstConfigModel.ManifestJSON})
			properties["memory"] = AppInstConfigModel.Memory
			properties["name_app_part"] = AppInstConfigModel.NameAppPart
			properties["name_project_part"] = AppInstConfigModel.NameProjectPart
			properties["naming_scheme"] = AppInstConfigModel.NamingScheme
			properties["networks"] = AppInstConfigModel.Networks
			properties["new_bundle_version_available"] = AppInstConfigModel.NewBundleVersionAvailable
			properties["origin_type"] = AppInstConfigModel.OriginType
			properties["parent_detail"] = SetObjectParentDetailSubResourceData([]*models.ObjectParentDetail{AppInstConfigModel.ParentDetail})
			properties["remote_console"] = AppInstConfigModel.RemoteConsole
			properties["start_delay_in_seconds"] = AppInstConfigModel.StartDelayInSeconds
			properties["storage"] = AppInstConfigModel.Storage
			properties["tags"] = AppInstConfigModel.Tags
			properties["vminfo"] = SetVMSubResourceData([]*models.VM{AppInstConfigModel.Vminfo})
			d = append(d, &properties)
		}
	}
	return
}

func AppInstConfig() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"bundle_id": {
			Description: `User defined name of the edge app, unique across the enterprise. Once app name is created, name can’t be changed`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"bundle_version": {
			Description: `current bundle version`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"cpus": {
			Description: `user defined cpus for bundle`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"drives": {
			Description: `user defined drives`,
			Type:        schema.TypeInt,
			Computed:    true,
		},

		"interfaces": {
			Description: `application interfaces`,
			Type:        schema.TypeList, //GoType: []*AppInterface
			Elem: &schema.Resource{
				Schema: AppInterface(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"logs": {
			Description: `App Instance logs`,
			Type:        schema.TypeList, //GoType: AppInstanceLogs
			Elem: &schema.Resource{
				Schema: AppInstanceLogs(),
			},
			Optional: true,
		},

		"manifest_json": {
			Description: `user defined manifest in JSON format`,
			Type:        schema.TypeList, //GoType: VMManifest
			Elem: &schema.Resource{
				Schema: VMManifest(),
			},
			Optional: true,
		},

		"memory": {
			Description: `user defined memory for bundle`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"name_app_part": {
			Description: `User provided name part  for the auto deployed app`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"name_project_part": {
			Description: `User provided name part  for the auto deployed app`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"naming_scheme": {
			Description: `app naming scheme`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"networks": {
			Description: `user defined network options`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"new_bundle_version_available": {
			Description: `this flag denotes whether there is latest bundle available in the marketplace or not`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"origin_type": {
			Description: `origin of object`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"parent_detail": {
			Description: `origin and parent related details`,
			Type:        schema.TypeList, //GoType: ObjectParentDetail
			Elem: &schema.Resource{
				Schema: ObjectParentDetail(),
			},
			Optional: true,
		},

		"remote_console": {
			Description: `Remote console flag`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"start_delay_in_seconds": {
			Description: `start delay is the time in seconds EVE should wait after boot before starting the application instance`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"storage": {
			Description: `user defined storage for bundle`,
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

		"vminfo": {
			Description: `virtual machine info`,
			Type:        schema.TypeList, //GoType: VM
			Elem: &schema.Resource{
				Schema: VMSchema(),
			},
			Optional: true,
		},
	}
}

func GetAppInstConfigPropertyFields() (t []string) {
	return []string{
		"bundle_id",
		"bundle_version",
		"cpus",
		"interfaces",
		"logs",
		"manifest_json",
		"memory",
		"name_app_part",
		"name_project_part",
		"naming_scheme",
		"networks",
		"new_bundle_version_available",
		"origin_type",
		"parent_detail",
		"remote_console",
		"start_delay_in_seconds",
		"storage",
		"tags",
		"vminfo",
	}
}
