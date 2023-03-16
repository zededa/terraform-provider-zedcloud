package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AppSummaryModel(d *schema.ResourceData) *models.AppSummary {
	appInstCountInt, _ := d.Get("app_inst_count").(int)
	appInstCount := int32(appInstCountInt)
	cpusInt, _ := d.Get("cpus").(int)
	cpus := int64(cpusInt)
	description, _ := d.Get("description").(string)
	id, _ := d.Get("id").(string)
	isImported, _ := d.Get("is_imported").(bool)
	var manifestJSON *models.VMManifestSummary // VMManifestSummary
	manifestJSONInterface, manifestJSONIsSet := d.GetOk("manifest_json")
	if manifestJSONIsSet && manifestJSONInterface != nil {
		manifestJSONMap := manifestJSONInterface.([]interface{})
		if len(manifestJSONMap) > 0 {
			manifestJSON = VMManifestSummaryModelFromMap(manifestJSONMap[0].(map[string]interface{}))
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
	storageInt, _ := d.Get("storage").(int)
	storage := int64(storageInt)
	title, _ := d.Get("title").(string)
	userDefinedVersion, _ := d.Get("user_defined_version").(string)
	return &models.AppSummary{
		AppInstCount:       appInstCount,
		Cpus:               cpus,
		Description:        description,
		ID:                 id,
		IsImported:         isImported,
		ManifestJSON:       manifestJSON,
		Memory:             memory,
		Name:               name,
		Networks:           networks,
		OriginType:         originType,
		ParentDetail:       parentDetail,
		Storage:            storage,
		Title:              title,
		UserDefinedVersion: userDefinedVersion,
	}
}

func AppSummaryModelFromMap(m map[string]interface{}) *models.AppSummary {
	appInstCount := int32(m["app_inst_count"].(int)) // int32
	cpus := int64(m["cpus"].(int))                   // int64
	description := m["description"].(string)
	id := m["id"].(string)
	isImported := m["is_imported"].(bool)
	var manifestJSON *models.VMManifestSummary // VMManifestSummary
	manifestJSONInterface, manifestJSONIsSet := m["manifest_json"]
	if manifestJSONIsSet && manifestJSONInterface != nil {
		manifestJSONMap := manifestJSONInterface.([]interface{})
		if len(manifestJSONMap) > 0 {
			manifestJSON = VMManifestSummaryModelFromMap(manifestJSONMap[0].(map[string]interface{}))
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
	storage := int64(m["storage"].(int)) // int64
	title := m["title"].(string)
	userDefinedVersion := m["user_defined_version"].(string)
	return &models.AppSummary{
		AppInstCount:       appInstCount,
		Cpus:               cpus,
		Description:        description,
		ID:                 id,
		IsImported:         isImported,
		ManifestJSON:       manifestJSON,
		Memory:             memory,
		Name:               name,
		Networks:           networks,
		OriginType:         originType,
		ParentDetail:       parentDetail,
		Storage:            storage,
		Title:              title,
		UserDefinedVersion: userDefinedVersion,
	}
}

// Update the underlying AppSummary resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetAppSummaryResourceData(d *schema.ResourceData, m *models.AppSummary) {
	d.Set("app_inst_count", m.AppInstCount)
	d.Set("cpus", m.Cpus)
	d.Set("description", m.Description)
	d.Set("drives", m.Drives)
	d.Set("id", m.ID)
	d.Set("is_imported", m.IsImported)
	d.Set("manifest_json", SetVMManifestSummarySubResourceData([]*models.VMManifestSummary{m.ManifestJSON}))
	d.Set("memory", m.Memory)
	d.Set("name", m.Name)
	d.Set("networks", m.Networks)
	d.Set("origin_type", m.OriginType)
	d.Set("parent_detail", SetObjectParentDetailSubResourceData([]*models.ObjectParentDetail{m.ParentDetail}))
	d.Set("storage", m.Storage)
	d.Set("title", m.Title)
	d.Set("user_defined_version", m.UserDefinedVersion)
}

// Iterate through and update the AppSummary resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetAppSummarySubResourceData(m []*models.AppSummary) (d []*map[string]interface{}) {
	for _, AppSummaryModel := range m {
		if AppSummaryModel != nil {
			properties := make(map[string]interface{})
			properties["app_inst_count"] = AppSummaryModel.AppInstCount
			properties["cpus"] = AppSummaryModel.Cpus
			properties["description"] = AppSummaryModel.Description
			properties["drives"] = AppSummaryModel.Drives
			properties["id"] = AppSummaryModel.ID
			properties["is_imported"] = AppSummaryModel.IsImported
			properties["manifest_json"] = SetVMManifestSummarySubResourceData([]*models.VMManifestSummary{AppSummaryModel.ManifestJSON})
			properties["memory"] = AppSummaryModel.Memory
			properties["name"] = AppSummaryModel.Name
			properties["networks"] = AppSummaryModel.Networks
			properties["origin_type"] = AppSummaryModel.OriginType
			properties["parent_detail"] = SetObjectParentDetailSubResourceData([]*models.ObjectParentDetail{AppSummaryModel.ParentDetail})
			properties["storage"] = AppSummaryModel.Storage
			properties["title"] = AppSummaryModel.Title
			properties["user_defined_version"] = AppSummaryModel.UserDefinedVersion
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the AppSummary resource defined in the Terraform configuration
func AppSummarySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"app_inst_count": {
			Description: `App instance count`,
			Type:        schema.TypeInt,
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

		"is_imported": {
			Description: `Flag to represent where app bundle is already imported`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"manifest_json": {
			Description: ``,
			Type:        schema.TypeList, //GoType: VMManifestSummary
			Elem: &schema.Resource{
				Schema: VMManifestSummarySchema(),
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
			Optional:    true,
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

		"storage": {
			Description: `user defined storage for bundle`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"title": {
			Description: `User defined title of the edge application. Title can be changed at any time`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"user_defined_version": {
			Description: `User defined version for the given edge-app`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the AppSummary resource
func GetAppSummaryPropertyFields() (t []string) {
	return []string{
		"app_inst_count",
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
		"storage",
		"title",
		"user_defined_version",
	}
}
