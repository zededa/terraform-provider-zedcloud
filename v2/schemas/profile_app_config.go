package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ProfileAppConfigModel(d *schema.ResourceData) *models.ProfileAppConfig {
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
	netInstName, _ := d.Get("net_inst_name").(string)
	netInstTags := map[string]string{}
	netInstTagsInterface, netInstTagsIsSet := d.GetOk("netInstTags")
	if netInstTagsIsSet {
		netInstTagsMap := netInstTagsInterface.(map[string]interface{})
		for k, v := range netInstTagsMap {
			if v == nil {
				continue
			}
			netInstTags[k] = v.(string)
		}
	}

	networksInt, _ := d.Get("networks").(int)
	networks := int64(networksInt)
	newBundleVersionAvailable, _ := d.Get("new_bundle_version_available").(bool)
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
	return &models.ProfileAppConfig{
		BundleVersion:             bundleVersion,
		Cpus:                      cpus,
		Interfaces:                interfaces,
		Logs:                      logs,
		ManifestJSON:              manifestJSON,
		Memory:                    memory,
		NameAppPart:               nameAppPart,
		NameProjectPart:           nameProjectPart,
		NetInstName:               netInstName,
		NetInstTags:               netInstTags,
		Networks:                  networks,
		NewBundleVersionAvailable: newBundleVersionAvailable,
		RemoteConsole:             remoteConsole,
		StartDelayInSeconds:       startDelayInSeconds,
		Storage:                   storage,
		Tags:                      tags,
		Vminfo:                    vminfo,
	}
}

func ProfileAppConfigModelFromMap(m map[string]interface{}) *models.ProfileAppConfig {
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
	netInstName := m["net_inst_name"].(string)
	netInstTags := map[string]string{}
	netInstTagsInterface, netInstTagsIsSet := m["net_inst_tags"]
	if netInstTagsIsSet {
		netInstTagsMap := netInstTagsInterface.(map[string]interface{})
		for k, v := range netInstTagsMap {
			if v == nil {
				continue
			}
			netInstTags[k] = v.(string)
		}
	}

	networks := int64(m["networks"].(int)) // int64
	newBundleVersionAvailable := m["new_bundle_version_available"].(bool)
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
	return &models.ProfileAppConfig{
		BundleVersion:             bundleVersion,
		Cpus:                      cpus,
		Interfaces:                interfaces,
		Logs:                      logs,
		ManifestJSON:              manifestJSON,
		Memory:                    memory,
		NameAppPart:               nameAppPart,
		NameProjectPart:           nameProjectPart,
		NetInstName:               netInstName,
		NetInstTags:               netInstTags,
		Networks:                  networks,
		NewBundleVersionAvailable: newBundleVersionAvailable,
		RemoteConsole:             remoteConsole,
		StartDelayInSeconds:       startDelayInSeconds,
		Storage:                   storage,
		Tags:                      tags,
		Vminfo:                    vminfo,
	}
}

func SetProfileAppConfigResourceData(d *schema.ResourceData, m *models.ProfileAppConfig) {
	d.Set("bundle_version", m.BundleVersion)
	d.Set("cpus", m.Cpus)
	d.Set("drives", m.Drives)
	d.Set("interfaces", SetAppInterfaceSubResourceData(m.Interfaces))
	d.Set("logs", SetAppInstanceLogsSubResourceData([]*models.AppInstanceLogs{m.Logs}))
	d.Set("manifest_json", SetVMManifestSubResourceData([]*models.VMManifest{m.ManifestJSON}))
	d.Set("memory", m.Memory)
	d.Set("name_app_part", m.NameAppPart)
	d.Set("name_project_part", m.NameProjectPart)
	d.Set("net_inst_name", m.NetInstName)
	d.Set("net_inst_tags", m.NetInstTags)
	d.Set("networks", m.Networks)
	d.Set("new_bundle_version_available", m.NewBundleVersionAvailable)
	d.Set("remote_console", m.RemoteConsole)
	d.Set("start_delay_in_seconds", m.StartDelayInSeconds)
	d.Set("storage", m.Storage)
	d.Set("tags", m.Tags)
	d.Set("vminfo", SetVMSubResourceData([]*models.VM{m.Vminfo}))
}

func SetProfileAppConfigSubResourceData(m []*models.ProfileAppConfig) (d []*map[string]interface{}) {
	for _, ProfileAppConfigModel := range m {
		if ProfileAppConfigModel != nil {
			properties := make(map[string]interface{})
			properties["bundle_version"] = ProfileAppConfigModel.BundleVersion
			properties["cpus"] = ProfileAppConfigModel.Cpus
			properties["drives"] = ProfileAppConfigModel.Drives
			properties["interfaces"] = SetAppInterfaceSubResourceData(ProfileAppConfigModel.Interfaces)
			properties["logs"] = SetAppInstanceLogsSubResourceData([]*models.AppInstanceLogs{ProfileAppConfigModel.Logs})
			properties["manifest_json"] = SetVMManifestSubResourceData([]*models.VMManifest{ProfileAppConfigModel.ManifestJSON})
			properties["memory"] = ProfileAppConfigModel.Memory
			properties["name_app_part"] = ProfileAppConfigModel.NameAppPart
			properties["name_project_part"] = ProfileAppConfigModel.NameProjectPart
			properties["net_inst_name"] = ProfileAppConfigModel.NetInstName
			properties["net_inst_tags"] = ProfileAppConfigModel.NetInstTags
			properties["networks"] = ProfileAppConfigModel.Networks
			properties["new_bundle_version_available"] = ProfileAppConfigModel.NewBundleVersionAvailable
			properties["remote_console"] = ProfileAppConfigModel.RemoteConsole
			properties["start_delay_in_seconds"] = ProfileAppConfigModel.StartDelayInSeconds
			properties["storage"] = ProfileAppConfigModel.Storage
			properties["tags"] = ProfileAppConfigModel.Tags
			properties["vminfo"] = SetVMSubResourceData([]*models.VM{ProfileAppConfigModel.Vminfo})
			d = append(d, &properties)
		}
	}
	return
}

func ProfileAppConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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
				Schema: AppInterfaceSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"logs": {
			Description: `App Instance logs`,
			Type:        schema.TypeList, //GoType: AppInstanceLogs
			Elem: &schema.Resource{
				Schema: AppInstanceLogsSchema(),
			},
			Optional: true,
		},

		"manifest_json": {
			Description: `user defined manifest in JSON format`,
			Type:        schema.TypeList, //GoType: VMManifest
			Elem: &schema.Resource{
				Schema: VMManifestSchema(),
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

		"net_inst_name": {
			Description: `network instance name`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"net_inst_tags": {
			Description: `Tags are name/value pairs that enable you to categorize resources. Tag names are case insensitive with max_length 512 and min_length 3. Tag values are case sensitive with max_length 256 and min_length 3.`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
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

func GetProfileAppConfigPropertyFields() (t []string) {
	return []string{
		"bundle_version",
		"cpus",
		"interfaces",
		"logs",
		"manifest_json",
		"memory",
		"name_app_part",
		"name_project_part",
		"net_inst_name",
		"net_inst_tags",
		"networks",
		"new_bundle_version_available",
		"remote_console",
		"start_delay_in_seconds",
		"storage",
		"tags",
		"vminfo",
	}
}
