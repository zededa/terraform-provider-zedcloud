package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func VolumeInstanceModel(d *schema.ResourceData) *models.VolumeInstance {
	var accessmode *models.VolumeInstanceAccessMode // VolumeInstanceAccessMode
	accessmodeInterface, accessmodeIsSet := d.GetOk("accessmode")
	if accessmodeIsSet {
		accessmodeModel := accessmodeInterface.(string)
		accessmode = models.NewVolumeInstanceAccessMode(models.VolumeInstanceAccessMode(accessmodeModel))
	}
	cleartext, _ := d.Get("cleartext").(bool)
	contentTreeID, _ := d.Get("content_tree_id").(string)
	description, _ := d.Get("description").(string)
	deviceID, _ := d.Get("device_id").(string)
	id, _ := d.Get("id").(string)
	image, _ := d.Get("image").(string)
	implicit, _ := d.Get("implicit").(bool)
	label, _ := d.Get("label").(string)
	multiattach, _ := d.Get("multiattach").(bool)
	name, _ := d.Get("name").(string)
	projectID, _ := d.Get("project_id").(string)
	// var purge *models.ZedCloudOpsCmd // ZedCloudOpsCmd
	// purgeInterface, purgeIsSet := d.GetOk("purge")
	// if purgeIsSet && purgeInterface != nil {
	// 	purgeMap := purgeInterface.([]interface{})
	// 	if len(purgeMap) > 0 {
	// 		purge = ZedCloudOpsCmdModelFromMap(purgeMap[0].(map[string]interface{}))
	// 	}
	// }
	sizeBytes, _ := d.Get("size_bytes").(string)
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
	var typeVar *models.VolumeInstanceType // VolumeInstanceType
	typeInterface, typeIsSet := d.GetOk("type")
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewVolumeInstanceType(models.VolumeInstanceType(typeModel))
	}
	return &models.VolumeInstance{
		Accessmode:    accessmode,
		Cleartext:     cleartext,
		ContentTreeID: contentTreeID,
		Description:   description,
		DeviceID:      deviceID,
		ID:            id,
		Image:         image,
		Implicit:      implicit,
		Label:         label,
		Multiattach:   multiattach,
		Name:          name,
		ProjectID:     projectID,
		// Purge:         purge,
		SizeBytes: sizeBytes,
		Tags:      tags,
		Title:     title,
		Type:      typeVar,
	}
}

func VolumeInstanceModelFromMap(m map[string]interface{}) *models.VolumeInstance {
	var accessmode *models.VolumeInstanceAccessMode // VolumeInstanceAccessMode
	accessmodeInterface, accessmodeIsSet := m["accessmode"]
	if accessmodeIsSet {
		accessmodeModel := accessmodeInterface.(string)
		accessmode = models.NewVolumeInstanceAccessMode(models.VolumeInstanceAccessMode(accessmodeModel))
	}
	cleartext := m["cleartext"].(bool)
	contentTreeID := m["content_tree_id"].(string)
	description := m["description"].(string)
	deviceID := m["device_id"].(string)
	id := m["id"].(string)
	image := m["image"].(string)
	implicit := m["implicit"].(bool)
	label := m["label"].(string)
	multiattach := m["multiattach"].(bool)
	name := m["name"].(string)
	projectID := m["project_id"].(string)
	// var purge *models.ZedCloudOpsCmd // ZedCloudOpsCmd
	// purgeInterface, purgeIsSet := m["purge"]
	// if purgeIsSet && purgeInterface != nil {
	// 	purgeMap := purgeInterface.([]interface{})
	// 	if len(purgeMap) > 0 {
	// 		purge = ZedCloudOpsCmdModelFromMap(purgeMap[0].(map[string]interface{}))
	// 	}
	// }
	//
	sizeBytes := m["size_bytes"].(string)
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
	var typeVar *models.VolumeInstanceType // VolumeInstanceType
	typeInterface, typeIsSet := m["type"]
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewVolumeInstanceType(models.VolumeInstanceType(typeModel))
	}
	return &models.VolumeInstance{
		Accessmode:    accessmode,
		Cleartext:     cleartext,
		ContentTreeID: contentTreeID,
		Description:   description,
		DeviceID:      deviceID,
		ID:            id,
		Image:         image,
		Implicit:      implicit,
		Label:         label,
		Multiattach:   multiattach,
		Name:          name,
		ProjectID:     projectID,
		// Purge:         purge,
		SizeBytes: sizeBytes,
		Tags:      tags,
		Title:     title,
		Type:      typeVar,
	}
}

func SetVolumeInstanceResourceData(d *schema.ResourceData, m *models.VolumeInstance) {
	d.Set("accessmode", m.Accessmode)
	d.Set("cleartext", m.Cleartext)
	d.Set("content_tree_id", m.ContentTreeID)
	d.Set("description", m.Description)
	d.Set("device_id", m.DeviceID)
	d.Set("id", m.ID)
	d.Set("image", m.Image)
	d.Set("implicit", m.Implicit)
	d.Set("label", m.Label)
	d.Set("multiattach", m.Multiattach)
	d.Set("name", m.Name)
	d.Set("project_id", m.ProjectID)
	// d.Set("purge", SetZedCloudOpsCmdSubResourceData([]*models.ZedCloudOpsCmd{m.Purge}))
	d.Set("revision", SetObjectRevisionSubResourceData([]*models.ObjectRevision{m.Revision}))
	d.Set("size_bytes", m.SizeBytes)
	d.Set("tags", m.Tags)
	d.Set("title", m.Title)
	d.Set("type", m.Type)
}

func SetVolumeInstanceSubResourceData(m []*models.VolumeInstance) (d []*map[string]interface{}) {
	for _, VolInstConfigModel := range m {
		if VolInstConfigModel != nil {
			properties := make(map[string]interface{})
			properties["accessmode"] = VolInstConfigModel.Accessmode
			properties["cleartext"] = VolInstConfigModel.Cleartext
			properties["content_tree_id"] = VolInstConfigModel.ContentTreeID
			properties["description"] = VolInstConfigModel.Description
			properties["device_id"] = VolInstConfigModel.DeviceID
			properties["id"] = VolInstConfigModel.ID
			properties["image"] = VolInstConfigModel.Image
			properties["implicit"] = VolInstConfigModel.Implicit
			properties["label"] = VolInstConfigModel.Label
			properties["multiattach"] = VolInstConfigModel.Multiattach
			properties["name"] = VolInstConfigModel.Name
			properties["project_id"] = VolInstConfigModel.ProjectID
			// properties["purge"] = SetZedCloudOpsCmdSubResourceData([]*models.ZedCloudOpsCmd{VolInstConfigModel.Purge})
			properties["revision"] = SetObjectRevisionSubResourceData([]*models.ObjectRevision{VolInstConfigModel.Revision})
			properties["size_bytes"] = VolInstConfigModel.SizeBytes
			properties["tags"] = VolInstConfigModel.Tags
			properties["title"] = VolInstConfigModel.Title
			properties["type"] = VolInstConfigModel.Type
			d = append(d, &properties)
		}
	}
	return
}

func VolumeInstance() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"accessmode": {
			Description: `Access mode`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"cleartext": {
			Description: `flag to keep the contents of the volume unencrypted (in clear text)`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"content_tree_id": {
			Description: `content tree ID`,
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
		},

		"description": {
			Description: `Detailed description of the volume instance.`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"device_id": {
			Description: `id of the device on which volume instance is created`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"id": {
			Description: `System defined universally unique Id of the volume instance.`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"image": {
			Description: `name of the image`,
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
		},

		"implicit": {
			Description: `flag to create implicit volumes`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"label": {
			Description: `label`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"multiattach": {
			Description: `flag to enable the volume to be attached to multiple app instances`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"name": {
			Description: `User defined name of the volume instance, unique across the enterprise. Once object is created, name can’t be changed.`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"project_id": {
			Description: `id of the project in which the volume instance is created`,
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
		},

		// "purge": {
		// 	Description: `Purge Counter information`,
		// 	Type:        schema.TypeList, //GoType: ZedCloudOpsCmd
		// 	Elem: &schema.Resource{
		// 		Schema: ZedCloudOpsCmdSchema(),
		// 	},
		// 	Optional: true,
		// },

		"revision": {
			Description: `system defined Revision info of the object`,
			Type:        schema.TypeList, //GoType: ObjectRevision
			Elem: &schema.Resource{
				Schema: ObjectRevision(),
			},
			Computed: true,
		},

		"size_bytes": {
			Description: `size of volume`,
			Type:        schema.TypeString,
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

		"title": {
			Description: `User defined title of the volume instance. Title can be changed at any time.`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"type": {
			Description: `type of Volume Instance`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the VolInstConfig resource
func GetVolInstConfigPropertyFields() (t []string) {
	return []string{
		"accessmode",
		"cleartext",
		"content_tree_id",
		"description",
		"device_id",
		"id",
		"image",
		"implicit",
		"label",
		"multiattach",
		"name",
		"project_id",
		"purge",
		"size_bytes",
		"tags",
		"title",
		"type",
	}
}
