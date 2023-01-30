package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func ImageModel(d *schema.ResourceData) *models.Image {
	datastoreID, _ := d.Get("datastore_id").(string)
	description, _ := d.Get("description").(string)
	id, _ := d.Get("id").(string)
	var imageArch *models.ModelArchType // ModelArchType
	imageArchInterface, imageArchIsSet := d.GetOk("image_arch")
	if imageArchIsSet {
		imageArchModel := imageArchInterface.(string)
		imageArch = models.NewModelArchType(models.ModelArchType(imageArchModel))
	}
	var imageFormat *models.ConfigFormat // ConfigFormat
	imageFormatInterface, imageFormatIsSet := d.GetOk("image_format")
	if imageFormatIsSet {
		imageFormatModel := imageFormatInterface.(string)
		imageFormat = models.NewConfigFormat(models.ConfigFormat(imageFormatModel))
	}
	imageRelURL, _ := d.Get("image_rel_url").(string)
	imageSha256, _ := d.Get("image_sha256").(string)
	imageSizeBytes, _ := d.Get("image_size_bytes").(string)
	var imageType *models.ImageType // ImageType
	imageTypeInterface, imageTypeIsSet := d.GetOk("image_type")
	if imageTypeIsSet {
		imageTypeModel := imageTypeInterface.(string)
		imageType = models.NewImageType(models.ImageType(imageTypeModel))
	}
	imageVersion, _ := d.Get("image_version").(string)
	name, _ := d.Get("name").(string)
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
	title, _ := d.Get("title").(string)
	return &models.Image{
		DatastoreID:       &datastoreID, // string true false false
		Description:       description,
		ID:                id,
		ImageArch:         imageArch,
		ImageFormat:       imageFormat,
		ImageRelURL:       imageRelURL,
		ImageSha256:       imageSha256,
		ImageSizeBytes:    imageSizeBytes,
		ImageType:         imageType,
		ImageVersion:      imageVersion,
		Name:              &name, // string true false false
		ProjectAccessList: projectAccessList,
		Title:             &title, // string true false false
	}
}

func ImageModelFromMap(m map[string]interface{}) *models.Image {
	datastoreID := m["datastore_id"].(string)
	description := m["description"].(string)
	id := m["id"].(string)
	var imageArch *models.ModelArchType // ModelArchType
	imageArchInterface, imageArchIsSet := m["image_arch"]
	if imageArchIsSet {
		imageArchModel := imageArchInterface.(string)
		imageArch = models.NewModelArchType(models.ModelArchType(imageArchModel))
	}
	var imageFormat *models.ConfigFormat // ConfigFormat
	imageFormatInterface, imageFormatIsSet := m["image_format"]
	if imageFormatIsSet {
		imageFormatModel := imageFormatInterface.(string)
		imageFormat = models.NewConfigFormat(models.ConfigFormat(imageFormatModel))
	}
	imageRelURL := m["image_rel_url"].(string)
	imageSha256 := m["image_sha256"].(string)
	imageSizeBytes := m["image_size_bytes"].(string)
	var imageType *models.ImageType // ImageType
	imageTypeInterface, imageTypeIsSet := m["image_type"]
	if imageTypeIsSet {
		imageTypeModel := imageTypeInterface.(string)
		imageType = models.NewImageType(models.ImageType(imageTypeModel))
	}
	imageVersion := m["image_version"].(string)
	name := m["name"].(string)
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
	title := m["title"].(string)
	return &models.Image{
		DatastoreID:       &datastoreID,
		Description:       description,
		ID:                id,
		ImageArch:         imageArch,
		ImageFormat:       imageFormat,
		ImageRelURL:       imageRelURL,
		ImageSha256:       imageSha256,
		ImageSizeBytes:    imageSizeBytes,
		ImageType:         imageType,
		ImageVersion:      imageVersion,
		Name:              &name,
		ProjectAccessList: projectAccessList,
		Title:             &title,
	}
}

func SetImageResourceData(d *schema.ResourceData, m *models.Image) {
	d.Set("datastore_id", m.DatastoreID)
	d.Set("description", m.Description)
	d.Set("id", m.ID)
	d.Set("image_arch", m.ImageArch)
	d.Set("image_error", m.ImageError)
	d.Set("image_format", m.ImageFormat)
	d.Set("image_local", m.ImageLocal)
	d.Set("image_rel_url", m.ImageRelURL)
	d.Set("image_sha256", m.ImageSha256)
	d.Set("image_size_bytes", m.ImageSizeBytes)
	d.Set("image_status", m.ImageStatus)
	d.Set("image_type", m.ImageType)
	d.Set("image_version", m.ImageVersion)
	d.Set("name", m.Name)
	d.Set("origin_type", m.OriginType)
	d.Set("project_access_list", m.ProjectAccessList)
	d.Set("revision", SetObjectRevisionSubResourceData([]*models.ObjectRevision{m.Revision}))
	d.Set("title", m.Title)
}

func SetImageSubResourceData(m []*models.Image) (d []*map[string]interface{}) {
	for _, ImageConfigModel := range m {
		if ImageConfigModel != nil {
			properties := make(map[string]interface{})
			properties["datastore_id"] = ImageConfigModel.DatastoreID
			properties["description"] = ImageConfigModel.Description
			properties["id"] = ImageConfigModel.ID
			properties["image_arch"] = ImageConfigModel.ImageArch
			properties["image_error"] = ImageConfigModel.ImageError
			properties["image_format"] = ImageConfigModel.ImageFormat
			properties["image_local"] = ImageConfigModel.ImageLocal
			properties["image_rel_url"] = ImageConfigModel.ImageRelURL
			properties["image_sha256"] = ImageConfigModel.ImageSha256
			properties["image_size_bytes"] = ImageConfigModel.ImageSizeBytes
			properties["image_status"] = ImageConfigModel.ImageStatus
			properties["image_type"] = ImageConfigModel.ImageType
			properties["image_version"] = ImageConfigModel.ImageVersion
			properties["name"] = ImageConfigModel.Name
			properties["origin_type"] = ImageConfigModel.OriginType
			properties["project_access_list"] = ImageConfigModel.ProjectAccessList
			properties["revision"] = SetObjectRevisionSubResourceData([]*models.ObjectRevision{ImageConfigModel.Revision})
			properties["title"] = ImageConfigModel.Title
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the ImageConfig resource defined in the Terraform configuration
func Image() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"datastore_id": {
			Description: `Datastore Id where image binary is located.`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"description": {
			Description: `Detailed description of the image.`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"id": {
			Description: `System defined universally unique Id of the image.`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"image_arch": {
			Description: `Image Architecture.`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"image_error": {
			Description: `Image upload/uplink detailed error/status message`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"image_format": {
			Description: `Image binary format.`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"image_local": {
			Description: `Internal image location.`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"image_rel_url": {
			Description: `Image relative path w.r.t. Datastore`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"image_sha256": {
			Description: `Image checksum in SHA256 format`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"image_size_bytes": {
			Description: `Image size in KBytes.`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"image_status": {
			Description: `Image status`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"image_type": {
			Description: `Image type`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"image_version": {
			Description: `system defined info`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"name": {
			Description: `User defined name of the image, unique across the enterprise. Once image is created, name can’t be changed.`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"origin_type": {
			Description: `Origin type of image.`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"project_access_list": {
			Description: `project access list of the image`,
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
			Computed: true,
		},

		"title": {
			Description: `User defined title of the image. Title can be changed at any time.`,
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

func GetImagePropertyFields() (t []string) {
	return []string{
		"datastore_id",
		"description",
		"id",
		"image_arch",
		"image_format",
		"image_rel_url",
		"image_sha256",
		"image_size_bytes",
		"image_type",
		"image_version",
		"name",
		"project_access_list",
		"title",
	}
}
