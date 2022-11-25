package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate AppConfig resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func AppConfigModel(d *schema.ResourceData) *models.AppConfig {
	appID := d.Get("app_id").(string)
	appVersion := d.Get("app_version").(string)
	cpus := int64(d.Get("cpus").(int))
	description := d.Get("description").(string)
	id := d.Get("id").(string)
	interfaces := d.Get("interfaces").([]*models.AppInterface) // []*AppInterface
	var manifestJSON *models.VMManifest                        // VMManifest
	manifestJSONInterface, manifestJSONIsSet := d.GetOk("manifest_json")
	if manifestJSONIsSet {
		manifestJSONMap := manifestJSONInterface.([]interface{})[0].(map[string]interface{})
		manifestJSON = VMManifestModelFromMap(manifestJSONMap)
	}
	memory := int64(d.Get("memory").(int))
	name := d.Get("name").(string)
	nameAppPart := d.Get("name_app_part").(string)
	nameProjectPart := d.Get("name_project_part").(string)
	namingScheme := d.Get("naming_scheme").(*models.AppNamingScheme) // AppNamingScheme
	networks := int64(d.Get("networks").(int))
	originType := d.Get("origin_type").(*models.Origin) // Origin
	var parentDetail *models.ObjectParentDetail         // ObjectParentDetail
	parentDetailInterface, parentDetailIsSet := d.GetOk("parent_detail")
	if parentDetailIsSet {
		parentDetailMap := parentDetailInterface.([]interface{})[0].(map[string]interface{})
		parentDetail = ObjectParentDetailModelFromMap(parentDetailMap)
	}
	storage := int64(d.Get("storage").(int))
	tags := d.Get("tags").(map[string]string) // map[string]string
	title := d.Get("title").(string)
	return &models.AppConfig{
		AppID:           appID,
		AppVersion:      appVersion,
		Cpus:            cpus,
		Description:     description,
		ID:              id,
		Interfaces:      interfaces,
		ManifestJSON:    manifestJSON,
		Memory:          memory,
		Name:            &name, // string true false false
		NameAppPart:     nameAppPart,
		NameProjectPart: nameProjectPart,
		NamingScheme:    namingScheme,
		Networks:        networks,
		OriginType:      originType,
		ParentDetail:    parentDetail,
		Storage:         storage,
		Tags:            tags,
		Title:           &title, // string true false false
	}
}

func AppConfigModelFromMap(m map[string]interface{}) *models.AppConfig {
	appID := m["app_id"].(string)
	appVersion := m["app_version"].(string)
	cpus := int64(m["cpus"].(int)) // int64 false false false
	description := m["description"].(string)
	id := m["id"].(string)
	interfaces := m["interfaces"].([]*models.AppInterface) // []*AppInterface
	var manifestJSON *models.VMManifest                    // VMManifest
	manifestJSONInterface, manifestJSONIsSet := m["manifest_json"]
	if manifestJSONIsSet {
		manifestJSONMap := manifestJSONInterface.([]interface{})[0].(map[string]interface{})
		manifestJSON = VMManifestModelFromMap(manifestJSONMap)
	}
	//
	memory := int64(m["memory"].(int)) // int64 false false false
	name := m["name"].(string)
	nameAppPart := m["name_app_part"].(string)
	nameProjectPart := m["name_project_part"].(string)
	namingScheme := m["naming_scheme"].(*models.AppNamingScheme) // AppNamingScheme
	networks := int64(m["networks"].(int))                       // int64 false false false
	originType := m["origin_type"].(*models.Origin)              // Origin
	var parentDetail *models.ObjectParentDetail                  // ObjectParentDetail
	parentDetailInterface, parentDetailIsSet := m["parent_detail"]
	if parentDetailIsSet {
		parentDetailMap := parentDetailInterface.([]interface{})[0].(map[string]interface{})
		parentDetail = ObjectParentDetailModelFromMap(parentDetailMap)
	}
	//
	storage := int64(m["storage"].(int)) // int64 false false false
	tags := m["tags"].(map[string]string)
	title := m["title"].(string)
	return &models.AppConfig{
		AppID:           appID,
		AppVersion:      appVersion,
		Cpus:            cpus,
		Description:     description,
		ID:              id,
		Interfaces:      interfaces,
		ManifestJSON:    manifestJSON,
		Memory:          memory,
		Name:            &name,
		NameAppPart:     nameAppPart,
		NameProjectPart: nameProjectPart,
		NamingScheme:    namingScheme,
		Networks:        networks,
		OriginType:      originType,
		ParentDetail:    parentDetail,
		Storage:         storage,
		Tags:            tags,
		Title:           &title,
	}
}

// Update the underlying AppConfig resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
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
	d.Set("storage", m.Storage)
	d.Set("tags", m.Tags)
	d.Set("title", m.Title)
}

// Iterate throught and update the AppConfig resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
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
			properties["storage"] = AppConfigModel.Storage
			properties["tags"] = AppConfigModel.Tags
			properties["title"] = AppConfigModel.Title
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the AppConfig resource defined in the Terraform configuration
func AppConfigSchema() map[string]*schema.Schema {
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
			Optional:    true,
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
			Type:        schema.TypeList, //GoType: []*AppInterface
			Elem: &schema.Resource{
				Schema: AppInterfaceSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"manifest_json": {
			Description: `Manifest data`,
			Type:        schema.TypeList, //GoType: VMManifest
			Elem: &schema.Resource{
				Schema: VMManifestSchema(),
			},
			Required: true,
		},

		"memory": {
			Description: `user defined memory for bundle`,
			Type:        schema.TypeInt,
			Optional:    true,
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
			Type:        schema.TypeList, //GoType: AppNamingScheme
			Elem: &schema.Resource{
				Schema: AppNamingSchemeSchema(),
			},
			Optional: true,
		},

		"networks": {
			Description: `user defined network options`,
			Type:        schema.TypeInt,
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

		"storage": {
			Description: `user defined storage for bundle`,
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

		"title": {
			Description: `User defined title of the edge application. Title can be changed at any time`,
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

// Retrieve property field names for updating the AppConfig resource
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
		"storage",
		"tags",
		"title",
	}
}
