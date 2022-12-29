package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate VolumeInstConfig resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func VolumeInstConfigModel(d *schema.ResourceData) *models.VolumeInstConfig {
	accessmodeModel, _ := d.Get("accessmode").(models.VolumeInstanceAccessMode) // VolumeInstanceAccessMode
	accessmode := &accessmodeModel
	if !ok {
		accessmode = nil
	}
	blockStorageTags, _ := d.Get("block_storage_tags").(map[string]string) // map[string]string
	cleartext, _ := d.Get("cleartext").(bool)
	contentTreeID, _ := d.Get("content_tree_id").(string)
	contentTreeTargetCondition, _ := d.Get("content_tree_target_condition").(map[string]string) // map[string]string
	image, _ := d.Get("image").(string)
	label, _ := d.Get("label").(string)
	multiattach, _ := d.Get("multiattach").(bool)
	var purge *models.ZedCloudOpsCmd // ZedCloudOpsCmd
	purgeInterface, purgeIsSet := d.GetOk("purge")
	if purgeIsSet {
		purgeMap := purgeInterface.([]interface{})[0].(map[string]interface{})
		purge = ZedCloudOpsCmdModelFromMap(purgeMap)
	}
	sizeBytes, _ := d.Get("size_bytes").(uint64)
	typeVarModel, _ := d.Get("type").(models.VolumeInstanceType) // VolumeInstanceType
	typeVar := &typeVarModel
	if !ok {
		typeVar = nil
	}
	return &models.VolumeInstConfig{
		Accessmode:                 accessmode,
		BlockStorageTags:           blockStorageTags,
		Cleartext:                  cleartext,
		ContentTreeID:              contentTreeID,
		ContentTreeTargetCondition: contentTreeTargetCondition,
		Image:                      image,
		Label:                      label,
		Multiattach:                multiattach,
		Purge:                      purge,
		SizeBytes:                  sizeBytes,
		Type:                       typeVar,
	}
}

func VolumeInstConfigModelFromMap(m map[string]interface{}) *models.VolumeInstConfig {
	accessmode := m["accessmode"].(*models.VolumeInstanceAccessMode) // VolumeInstanceAccessMode
	blockStorageTags := m["block_storage_tags"].(map[string]string)
	cleartext := m["cleartext"].(bool)
	contentTreeID := m["content_tree_id"].(string)
	contentTreeTargetCondition := m["content_tree_target_condition"].(map[string]string)
	image := m["image"].(string)
	label := m["label"].(string)
	multiattach := m["multiattach"].(bool)
	var purge *models.ZedCloudOpsCmd // ZedCloudOpsCmd
	purgeInterface, purgeIsSet := m["purge"]
	if purgeIsSet {
		purgeMap := purgeInterface.([]interface{})[0].(map[string]interface{})
		purge = ZedCloudOpsCmdModelFromMap(purgeMap)
	}
	//
	sizeBytes := m["size_bytes"].(uint64)
	typeVar := m["type"].(*models.VolumeInstanceType) // VolumeInstanceType
	return &models.VolumeInstConfig{
		Accessmode:                 accessmode,
		BlockStorageTags:           blockStorageTags,
		Cleartext:                  cleartext,
		ContentTreeID:              contentTreeID,
		ContentTreeTargetCondition: contentTreeTargetCondition,
		Image:                      image,
		Label:                      label,
		Multiattach:                multiattach,
		Purge:                      purge,
		SizeBytes:                  sizeBytes,
		Type:                       typeVar,
	}
}

// Update the underlying VolumeInstConfig resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetVolumeInstConfigResourceData(d *schema.ResourceData, m *models.VolumeInstConfig) {
	d.Set("accessmode", m.Accessmode)
	d.Set("block_storage_tags", m.BlockStorageTags)
	d.Set("cleartext", m.Cleartext)
	d.Set("content_tree_id", m.ContentTreeID)
	d.Set("content_tree_target_condition", m.ContentTreeTargetCondition)
	d.Set("image", m.Image)
	d.Set("label", m.Label)
	d.Set("multiattach", m.Multiattach)
	d.Set("purge", SetZedCloudOpsCmdSubResourceData([]*models.ZedCloudOpsCmd{m.Purge}))
	d.Set("size_bytes", m.SizeBytes)
	d.Set("type", m.Type)
}

// Iterate through and update the VolumeInstConfig resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetVolumeInstConfigSubResourceData(m []*models.VolumeInstConfig) (d []*map[string]interface{}) {
	for _, VolumeInstConfigModel := range m {
		if VolumeInstConfigModel != nil {
			properties := make(map[string]interface{})
			properties["accessmode"] = VolumeInstConfigModel.Accessmode
			properties["block_storage_tags"] = VolumeInstConfigModel.BlockStorageTags
			properties["cleartext"] = VolumeInstConfigModel.Cleartext
			properties["content_tree_id"] = VolumeInstConfigModel.ContentTreeID
			properties["content_tree_target_condition"] = VolumeInstConfigModel.ContentTreeTargetCondition
			properties["image"] = VolumeInstConfigModel.Image
			properties["label"] = VolumeInstConfigModel.Label
			properties["multiattach"] = VolumeInstConfigModel.Multiattach
			properties["purge"] = SetZedCloudOpsCmdSubResourceData([]*models.ZedCloudOpsCmd{VolumeInstConfigModel.Purge})
			properties["size_bytes"] = VolumeInstConfigModel.SizeBytes
			properties["type"] = VolumeInstConfigModel.Type
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the VolumeInstConfig resource defined in the Terraform configuration
func VolumeInstConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"accessmode": {
			Description: `Access mode`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"block_storage_tags": {
			Description: `user defined key-value pairs of a block storage, will be used for targeting content tree`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
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
		},

		"content_tree_target_condition": {
			Description: `user defined key-value pairs of a content tree that will be used for targeting by block storage`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"image": {
			Description: `name of the image`,
			Type:        schema.TypeString,
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

		"purge": {
			Description: `Purge Counter information`,
			Type:        schema.TypeList, //GoType: ZedCloudOpsCmd
			Elem: &schema.Resource{
				Schema: ZedCloudOpsCmdSchema(),
			},
			Optional: true,
		},

		"size_bytes": {
			Description: `size of volume`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"type": {
			Description: `type of Volume Instance`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the VolumeInstConfig resource
func GetVolumeInstConfigPropertyFields() (t []string) {
	return []string{
		"accessmode",
		"block_storage_tags",
		"cleartext",
		"content_tree_id",
		"content_tree_target_condition",
		"image",
		"label",
		"multiattach",
		"purge",
		"size_bytes",
		"type",
	}
}
