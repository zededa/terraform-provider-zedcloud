package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate AppInstConfig resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func AppInstConfigModel(d *schema.ResourceData) *models.AppInstConfig {
	bundleID := d.Get("bundle_id").(string)
	bundleVersion := int64(d.Get("bundle_version").(int))
	cpus := int64(d.Get("cpus").(int))
	interfaces := d.Get("interfaces").([]*models.AppInterface) // []*AppInterface
	var logs *models.AppInstanceLogs                           // AppInstanceLogs
	logsInterface, logsIsSet := d.GetOk("logs")
	if logsIsSet {
		logsMap := logsInterface.([]interface{})[0].(map[string]interface{})
		logs = AppInstanceLogsModelFromMap(logsMap)
	}
	var manifestJSON *models.VMManifest // VMManifest
	manifestJSONInterface, manifestJSONIsSet := d.GetOk("manifest_json")
	if manifestJSONIsSet {
		manifestJSONMap := manifestJSONInterface.([]interface{})[0].(map[string]interface{})
		manifestJSON = VMManifestModelFromMap(manifestJSONMap)
	}
	memory := int64(d.Get("memory").(int))
	nameAppPart := d.Get("name_app_part").(string)
	nameProjectPart := d.Get("name_project_part").(string)
	namingScheme := d.Get("naming_scheme").(*models.AppNamingSchemeV2) // AppNamingSchemeV2
	networks := int64(d.Get("networks").(int))
	newBundleVersionAvailable := d.Get("new_bundle_version_available").(bool)
	originType := d.Get("origin_type").(*models.Origin) // Origin
	var parentDetail *models.ObjectParentDetail         // ObjectParentDetail
	parentDetailInterface, parentDetailIsSet := d.GetOk("parent_detail")
	if parentDetailIsSet {
		parentDetailMap := parentDetailInterface.([]interface{})[0].(map[string]interface{})
		parentDetail = ObjectParentDetailModelFromMap(parentDetailMap)
	}
	remoteConsole := d.Get("remote_console").(bool)
	startDelayInSeconds := int64(d.Get("start_delay_in_seconds").(int))
	storage := int64(d.Get("storage").(int))
	tags := d.Get("tags").(map[string]string) // map[string]string
	var vminfo *models.VM                     // VM
	vminfoInterface, vminfoIsSet := d.GetOk("vminfo")
	if vminfoIsSet {
		vminfoMap := vminfoInterface.([]interface{})[0].(map[string]interface{})
		vminfo = VMModelFromMap(vminfoMap)
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
	bundleVersion := int64(m["bundle_version"].(int))      // int64 false false false
	cpus := int64(m["cpus"].(int))                         // int64 false false false
	interfaces := m["interfaces"].([]*models.AppInterface) // []*AppInterface
	var logs *models.AppInstanceLogs                       // AppInstanceLogs
	logsInterface, logsIsSet := m["logs"]
	if logsIsSet {
		logsMap := logsInterface.([]interface{})[0].(map[string]interface{})
		logs = AppInstanceLogsModelFromMap(logsMap)
	}
	//
	var manifestJSON *models.VMManifest // VMManifest
	manifestJSONInterface, manifestJSONIsSet := m["manifest_json"]
	if manifestJSONIsSet {
		manifestJSONMap := manifestJSONInterface.([]interface{})[0].(map[string]interface{})
		manifestJSON = VMManifestModelFromMap(manifestJSONMap)
	}
	//
	memory := int64(m["memory"].(int)) // int64 false false false
	nameAppPart := m["name_app_part"].(string)
	nameProjectPart := m["name_project_part"].(string)
	namingScheme := m["naming_scheme"].(*models.AppNamingSchemeV2) // AppNamingSchemeV2
	networks := int64(m["networks"].(int))                         // int64 false false false
	newBundleVersionAvailable := m["new_bundle_version_available"].(bool)
	originType := m["origin_type"].(*models.Origin) // Origin
	var parentDetail *models.ObjectParentDetail     // ObjectParentDetail
	parentDetailInterface, parentDetailIsSet := m["parent_detail"]
	if parentDetailIsSet {
		parentDetailMap := parentDetailInterface.([]interface{})[0].(map[string]interface{})
		parentDetail = ObjectParentDetailModelFromMap(parentDetailMap)
	}
	//
	remoteConsole := m["remote_console"].(bool)
	startDelayInSeconds := int64(m["start_delay_in_seconds"].(int)) // int64 false false false
	storage := int64(m["storage"].(int))                            // int64 false false false
	tags := m["tags"].(map[string]string)
	var vminfo *models.VM // VM
	vminfoInterface, vminfoIsSet := m["vminfo"]
	if vminfoIsSet {
		vminfoMap := vminfoInterface.([]interface{})[0].(map[string]interface{})
		vminfo = VMModelFromMap(vminfoMap)
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

// Update the underlying AppInstConfig resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
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

// Iterate throught and update the AppInstConfig resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
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

// Schema mapping representing the AppInstConfig resource defined in the Terraform configuration
func AppInstConfigSchema() map[string]*schema.Schema {
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
				Schema: AppInterfaceSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
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

		"naming_scheme": {
			Description: `app naming scheme`,
			Type:        schema.TypeList, //GoType: AppNamingSchemeV2
			Elem: &schema.Resource{
				Schema: AppNamingSchemeV2Schema(),
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

		"origin_type": {
			Description: `origin of object`,
			Type:        schema.TypeList, //GoType: Origin
			Elem: &schema.Resource{
				Schema: OriginSchema(),
			},
			Optional: true,
		},

		"parent_detail": {
			Description: `origin and parent related details`,
			Type:        schema.TypeList, //GoType: ObjectParentDetail
			Elem: &schema.Resource{
				Schema: ObjectParentDetailSchema(),
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

// Retrieve property field names for updating the AppInstConfig resource
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
