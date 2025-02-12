package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func AppConfigModel(d *schema.ResourceData) *models.AppConfig {
	appID, _ := d.Get("app_id").(string)
	appVersion, _ := d.Get("app_version").(string)
	cpusInt, _ := d.Get("cpus").(int)
	cpus := int64(cpusInt)
	description, _ := d.Get("description").(string)
	id, _ := d.Get("id").(string)
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
	name, _ := d.Get("name").(string)
	nameAppPart, _ := d.Get("name_app_part").(string)
	nameProjectPart, _ := d.Get("name_project_part").(string)
	var namingScheme *models.AppNamingScheme // AppNamingScheme
	namingSchemeInterface, namingSchemeIsSet := d.GetOk("naming_scheme")
	if namingSchemeIsSet {
		namingSchemeModel := namingSchemeInterface.(string)
		namingScheme = models.NewAppNamingScheme(models.AppNamingScheme(namingSchemeModel))
	}
	networksInt, _ := d.Get("networks").(int)
	networks := int64(networksInt)
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

	title, _ := d.Get("title").(string)
	return &models.AppConfig{
		AppID:               appID,
		AppVersion:          appVersion,
		Cpus:                &cpus, // int64
		Description:         description,
		ID:                  id,
		Interfaces:          interfaces,
		ManifestJSON:        manifestJSON,
		Memory:              &memory, // int64
		Name:                &name,   // string
		NameAppPart:         nameAppPart,
		NameProjectPart:     nameProjectPart,
		NamingScheme:        namingScheme,
		Networks:            &networks, // int64
		OriginType:          originType,
		ParentDetail:        parentDetail,
		StartDelayInSeconds: startDelayInSeconds,
		Storage:             storage,
		Tags:                tags,
		Title:               &title, // string
	}
}

func AppConfigModelFromMap(m map[string]interface{}) *models.AppConfig {
	appID := m["app_id"].(string)
	appVersion := m["app_version"].(string)
	cpus := int64(m["cpus"].(int)) // int64
	description := m["description"].(string)
	id := m["id"].(string)
	var interfaces []*models.AppInterface // []*AppInterface
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
	name := m["name"].(string)
	nameAppPart := m["name_app_part"].(string)
	nameProjectPart := m["name_project_part"].(string)
	var namingScheme *models.AppNamingScheme // AppNamingScheme
	namingSchemeInterface, namingSchemeIsSet := m["naming_scheme"]
	if namingSchemeIsSet {
		namingSchemeModel := namingSchemeInterface.(string)
		namingScheme = models.NewAppNamingScheme(models.AppNamingScheme(namingSchemeModel))
	}
	networks := int64(m["networks"].(int)) // int64
	var originType *models.Origin          // Origin
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

	title := m["title"].(string)
	return &models.AppConfig{
		AppID:               appID,
		AppVersion:          appVersion,
		Cpus:                &cpus,
		Description:         description,
		ID:                  id,
		Interfaces:          interfaces,
		ManifestJSON:        manifestJSON,
		Memory:              &memory,
		Name:                &name,
		NameAppPart:         nameAppPart,
		NameProjectPart:     nameProjectPart,
		NamingScheme:        namingScheme,
		Networks:            &networks,
		OriginType:          originType,
		ParentDetail:        parentDetail,
		StartDelayInSeconds: startDelayInSeconds,
		Storage:             storage,
		Tags:                tags,
		Title:               &title,
	}
}

func SetAppConfigResourceData(d *schema.ResourceData, m *models.AppConfig) {
	d.Set("app_id", m.AppID)
	d.Set("app_version", m.AppVersion)
	d.Set("cpus", m.Cpus)
	d.Set("description", m.Description)
	d.Set("drives", m.Drives)
	d.Set("id", m.ID)
	d.Set("interfaces", SetAppInterfaceSubResourceData(m.Interfaces))
	d.Set("manifest_json", SetVMManifestSubResourceData([]*models.VMManifest{m.ManifestJSON}))
	d.Set("memory", m.Memory)
	d.Set("name", m.Name)
	d.Set("name_app_part", m.NameAppPart)
	d.Set("name_project_part", m.NameProjectPart)
	d.Set("naming_scheme", m.NamingScheme)
	d.Set("networks", m.Networks)
	d.Set("origin_type", m.OriginType)
	d.Set("parent_detail", SetObjectParentDetailSubResourceData([]*models.ObjectParentDetail{m.ParentDetail}))
	d.Set("start_delay_in_seconds", m.StartDelayInSeconds)
	d.Set("storage", m.Storage)
	d.Set("tags", m.Tags)
	d.Set("title", m.Title)
}

func SetAppConfigSubResourceData(m []*models.AppConfig) (d []*map[string]interface{}) {
	for _, AppConfigModel := range m {
		if AppConfigModel != nil {
			properties := make(map[string]interface{})
			properties["app_id"] = AppConfigModel.AppID
			properties["app_version"] = AppConfigModel.AppVersion
			properties["cpus"] = AppConfigModel.Cpus
			properties["description"] = AppConfigModel.Description
			properties["drives"] = AppConfigModel.Drives
			properties["id"] = AppConfigModel.ID
			properties["interfaces"] = SetAppInterfaceSubResourceData(AppConfigModel.Interfaces)
			properties["manifest_json"] = SetVMManifestSubResourceData([]*models.VMManifest{AppConfigModel.ManifestJSON})
			properties["memory"] = AppConfigModel.Memory
			properties["name"] = AppConfigModel.Name
			properties["name_app_part"] = AppConfigModel.NameAppPart
			properties["name_project_part"] = AppConfigModel.NameProjectPart
			properties["naming_scheme"] = AppConfigModel.NamingScheme
			properties["networks"] = AppConfigModel.Networks
			properties["origin_type"] = AppConfigModel.OriginType
			properties["parent_detail"] = SetObjectParentDetailSubResourceData([]*models.ObjectParentDetail{AppConfigModel.ParentDetail})
			properties["start_delay_in_seconds"] = AppConfigModel.StartDelayInSeconds
			properties["storage"] = AppConfigModel.Storage
			properties["tags"] = AppConfigModel.Tags
			properties["title"] = AppConfigModel.Title
			d = append(d, &properties)
		}
	}
	return
}

func AppConfig() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"app_id": {
			Description: `User defined name of the edge app, unique across the enterprise. Once app name is created, name can’t be changed`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"app_version": {
			Description: `Current version of the attached bundle`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"cpus": {
			Description: `user defined cpus for bundle`,
			Type:        schema.TypeInt,
			Required:    true,
		},

		"description": {
			Description: `Detailed description of the edge application`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"drives": {
			Description: `user defined drives`,
			Type:        schema.TypeInt,
			Computed:    true,
		},

		"id": {
			Description: `System defined universally unique Id of the edge application`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"interfaces": {
			Description: `application interfaces`,
			Type:        schema.TypeList, // GoType: []*AppInterface
			Elem: &schema.Resource{
				Schema: AppInterface(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"manifest_json": {
			Description: `Manifest data`,
			Type:        schema.TypeList, // GoType: VMManifest
			Elem: &schema.Resource{
				Schema: VMManifest(),
			},
			Optional: true,
		},

		"memory": {
			Description: `user defined memory for bundle`,
			Type:        schema.TypeInt,
			Required:    true,
		},

		"name": {
			Description: `User defined name of the edge application, unique across the enterprise. Once object is created, name can’t be changed`,
			Type:        schema.TypeString,
			Required:    true,
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
			Required:    true,
		},

		"origin_type": {
			Description: `origin of object`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"parent_detail": {
			Description: `origin and parent related details`,
			Type:        schema.TypeList, // GoType: ObjectParentDetail
			Elem: &schema.Resource{
				Schema: ObjectParentDetail(),
			},
			Optional: true,
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
			Type:        schema.TypeMap, // GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"title": {
			Description: `User defined title of the edge application. Title can be changed at any time`,
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

func GetAppConfigPropertyFields() (t []string) {
	return []string{
		"app_id",
		"app_version",
		"cpus",
		"description",
		"id",
		"interfaces",
		"manifest_json",
		"memory",
		"name",
		"name_app_part",
		"name_project_part",
		"naming_scheme",
		"networks",
		"origin_type",
		"parent_detail",
		"start_delay_in_seconds",
		"storage",
		"tags",
		"title",
	}
}
