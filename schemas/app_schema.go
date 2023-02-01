package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func ApplicationModel(d *schema.ResourceData) *models.Application {
	cpusInt, _ := d.Get("cpus").(int)
	cpus := int64(cpusInt)
	description, _ := d.Get("description").(string)
	id, _ := d.Get("id").(string)
	isImported, _ := d.Get("is_imported").(bool)
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
	var projectAccessList []string
	projectAccessListInterface, projectAccessListIsSet := d.GetOk("projectAccessList")
	if projectAccessListIsSet {
		var items []interface{}
		if listItems, isList := projectAccessListInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = projectAccessListInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			projectAccessList = append(projectAccessList, v.(string))
		}
	}
	var revision *models.ObjectRevision // ObjectRevision
	revisionInterface, revisionIsSet := d.GetOk("revision")
	if revisionIsSet && revisionInterface != nil {
		revisionMap := revisionInterface.([]interface{})
		if len(revisionMap) > 0 {
			revision = ObjectRevisionModelFromMap(revisionMap[0].(map[string]interface{}))
		}
	}
	storageInt, _ := d.Get("storage").(int)
	storage := int64(storageInt)
	title, _ := d.Get("title").(string)
	userDefinedVersion, _ := d.Get("user_defined_version").(string)
	return &models.Application{
		Cpus:               cpus,
		Description:        description,
		ID:                 id,
		IsImported:         isImported,
		ManifestJSON:       manifestJSON,
		Memory:             memory,
		Name:               &name, // string true false false
		Networks:           networks,
		OriginType:         originType,
		ParentDetail:       parentDetail,
		ProjectAccessList:  projectAccessList,
		Revision:           revision,
		Storage:            storage,
		Title:              &title, // string true false false
		UserDefinedVersion: userDefinedVersion,
	}
}

func EdgeApplicationModelFromMap(m map[string]interface{}) *models.Application {
	cpus := int64(m["cpus"].(int)) // int64
	description := m["description"].(string)
	id := m["id"].(string)
	isImported := m["is_imported"].(bool)
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
	var projectAccessList []string
	projectAccessListInterface, projectAccessListIsSet := m["projectAccessList"]
	if projectAccessListIsSet {
		var items []interface{}
		if listItems, isList := projectAccessListInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = projectAccessListInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			projectAccessList = append(projectAccessList, v.(string))
		}
	}
	var revision *models.ObjectRevision // ObjectRevision
	revisionInterface, revisionIsSet := m["revision"]
	if revisionIsSet && revisionInterface != nil {
		revisionMap := revisionInterface.([]interface{})
		if len(revisionMap) > 0 {
			revision = ObjectRevisionModelFromMap(revisionMap[0].(map[string]interface{}))
		}
	}
	//
	storage := int64(m["storage"].(int)) // int64
	title := m["title"].(string)
	userDefinedVersion := m["user_defined_version"].(string)
	return &models.Application{
		Cpus:               cpus,
		Description:        description,
		ID:                 id,
		IsImported:         isImported,
		ManifestJSON:       manifestJSON,
		Memory:             memory,
		Name:               &name,
		Networks:           networks,
		OriginType:         originType,
		ParentDetail:       parentDetail,
		ProjectAccessList:  projectAccessList,
		Revision:           revision,
		Storage:            storage,
		Title:              &title,
		UserDefinedVersion: userDefinedVersion,
	}
}

func SetApplicationResourceData(d *schema.ResourceData, m *models.Application) {
	d.Set("cpus", m.Cpus)
	d.Set("description", m.Description)
	d.Set("drives", m.Drives)
	d.Set("id", m.ID)
	d.Set("is_imported", m.IsImported)
	d.Set("manifest_json", SetVMManifestSubResourceData([]*models.VMManifest{m.ManifestJSON}))
	d.Set("memory", m.Memory)
	d.Set("name", m.Name)
	d.Set("networks", m.Networks)
	d.Set("origin_type", m.OriginType)
	d.Set("parent_detail", SetObjectParentDetailSubResourceData([]*models.ObjectParentDetail{m.ParentDetail}))
	d.Set("project_access_list", m.ProjectAccessList)
	d.Set("revision", SetObjectRevisionSubResourceData([]*models.ObjectRevision{m.Revision}))
	d.Set("storage", m.Storage)
	d.Set("title", m.Title)
	d.Set("user_defined_version", m.UserDefinedVersion)
}

func SetApplicationSubResourceData(m []*models.Application) (d []*map[string]interface{}) {
	for _, AppModel := range m {
		if AppModel != nil {
			properties := make(map[string]interface{})
			properties["cpus"] = AppModel.Cpus
			properties["description"] = AppModel.Description
			properties["drives"] = AppModel.Drives
			properties["id"] = AppModel.ID
			properties["is_imported"] = AppModel.IsImported
			properties["manifest_json"] = SetVMManifestSubResourceData([]*models.VMManifest{AppModel.ManifestJSON})
			properties["memory"] = AppModel.Memory
			properties["name"] = AppModel.Name
			properties["networks"] = AppModel.Networks
			properties["origin_type"] = AppModel.OriginType
			properties["parent_detail"] = SetObjectParentDetailSubResourceData([]*models.ObjectParentDetail{AppModel.ParentDetail})
			properties["project_access_list"] = AppModel.ProjectAccessList
			properties["revision"] = SetObjectRevisionSubResourceData([]*models.ObjectRevision{AppModel.Revision})
			properties["storage"] = AppModel.Storage
			properties["title"] = AppModel.Title
			properties["user_defined_version"] = AppModel.UserDefinedVersion
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the App resource defined in the Terraform configuration
func Application() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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

		"is_imported": {
			Description: `Flag to represent where app bundle is already imported`,
			Type:        schema.TypeBool,
			Optional:    true,
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

		"name": {
			Description: `User defined name of the edge application, unique across the enterprise. Once object is created, name can’t be changed`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"networks": {
			Description: `user defined network options`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"origin_type": {
			Description: `origin of object`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"parent_detail": {
			Description: `origin and parent related details`,
			Type:        schema.TypeList, //GoType: ObjectParentDetail
			Elem: &schema.Resource{
				Schema: ObjectParentDetailSchema(),
			},
			Optional: true,
		},

		"project_access_list": {
			Description: `project access list of the app bundle`,
			Type:        schema.TypeList, //GoType: []string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"revision": {
			Description: `system defined info`,
			Type:        schema.TypeList, //GoType: ObjectRevision
			Elem: &schema.Resource{
				Schema: ObjectRevision(),
			},
			Optional: true,
		},

		"storage": {
			Description: `user defined storage for bundle`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"title": {
			Description: `User defined title of the edge application. Title can be changed at any time`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"user_defined_version": {
			Description: `User defined version for the given edge-app`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the App resource
func GetApplicationPropertyFields() (t []string) {
	return []string{
		"cpus",
		"description",
		"id",
		"is_imported",
		"manifest_json",
		"memory",
		"name",
		"networks",
		"origin_type",
		"parent_detail",
		"project_access_list",
		"revision",
		"storage",
		"title",
		"user_defined_version",
	}
}
