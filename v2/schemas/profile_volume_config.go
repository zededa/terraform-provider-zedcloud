package schemas

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ProfileVolumeConfigModel(d *schema.ResourceData) *models.ProfileVolumeConfig {
	var accessmode *models.VolumeInstanceAccessMode // VolumeInstanceAccessMode
	accessmodeInterface, accessmodeIsSet := d.GetOk("accessmode")
	if accessmodeIsSet {
		accessmodeModel := accessmodeInterface.(string)
		accessmode = models.NewVolumeInstanceAccessMode(models.VolumeInstanceAccessMode(accessmodeModel))
	}
	blockStorageTags := map[string]string{}
	blockStorageTagsInterface, blockStorageTagsIsSet := d.GetOk("blockStorageTags")
	if blockStorageTagsIsSet {
		blockStorageTagsMap := blockStorageTagsInterface.(map[string]interface{})
		for k, v := range blockStorageTagsMap {
			if v == nil {
				continue
			}
			blockStorageTags[k] = v.(string)
		}
	}

	cleartext, _ := d.Get("cleartext").(bool)
	contentTreeID, _ := d.Get("content_tree_id").(string)
	contentTreeTargetCondition := map[string]string{}
	contentTreeTargetConditionInterface, contentTreeTargetConditionIsSet := d.GetOk("contentTreeTargetCondition")
	if contentTreeTargetConditionIsSet {
		contentTreeTargetConditionMap := contentTreeTargetConditionInterface.(map[string]interface{})
		for k, v := range contentTreeTargetConditionMap {
			if v == nil {
				continue
			}
			contentTreeTargetCondition[k] = v.(string)
		}
	}

	image, _ := d.Get("image").(string)
	label, _ := d.Get("label").(string)
	multiattach, _ := d.Get("multiattach").(bool)
	var purge *models.ZedCloudOpsCmd // ZedCloudOpsCmd
	purgeInterface, purgeIsSet := d.GetOk("purge")
	if purgeIsSet && purgeInterface != nil {
		purgeMap := purgeInterface.([]interface{})
		if len(purgeMap) > 0 {
			purge = ZedCloudOpsCmdModelFromMap(purgeMap[0].(map[string]interface{}))
		}
	}
	sizeBytes, _ := d.Get("size_bytes").(string)
	var typeVar *models.VolumeInstanceType // VolumeInstanceType
	typeInterface, typeIsSet := d.GetOk("type")
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewVolumeInstanceType(models.VolumeInstanceType(typeModel))
	}
	return &models.ProfileVolumeConfig{
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

func ProfileVolumeConfigModelFromMap(m map[string]interface{}) *models.ProfileVolumeConfig {
	var accessmode *models.VolumeInstanceAccessMode // VolumeInstanceAccessMode
	accessmodeInterface, accessmodeIsSet := m["accessmode"]
	if accessmodeIsSet {
		accessmodeModel := accessmodeInterface.(string)
		accessmode = models.NewVolumeInstanceAccessMode(models.VolumeInstanceAccessMode(accessmodeModel))
	}
	blockStorageTags := map[string]string{}
	blockStorageTagsInterface, blockStorageTagsIsSet := m["block_storage_tags"]
	if blockStorageTagsIsSet {
		blockStorageTagsMap := blockStorageTagsInterface.(map[string]interface{})
		for k, v := range blockStorageTagsMap {
			if v == nil {
				continue
			}
			blockStorageTags[k] = v.(string)
		}
	}

	cleartext := m["cleartext"].(bool)
	contentTreeID := m["content_tree_id"].(string)
	contentTreeTargetCondition := map[string]string{}
	contentTreeTargetConditionInterface, contentTreeTargetConditionIsSet := m["content_tree_target_condition"]
	if contentTreeTargetConditionIsSet {
		contentTreeTargetConditionMap := contentTreeTargetConditionInterface.(map[string]interface{})
		for k, v := range contentTreeTargetConditionMap {
			if v == nil {
				continue
			}
			contentTreeTargetCondition[k] = v.(string)
		}
	}

	image := m["image"].(string)
	label := m["label"].(string)
	multiattach := m["multiattach"].(bool)
	var purge *models.ZedCloudOpsCmd // ZedCloudOpsCmd
	purgeInterface, purgeIsSet := m["purge"]
	if purgeIsSet && purgeInterface != nil {
		purgeMap := purgeInterface.([]interface{})
		if len(purgeMap) > 0 {
			purge = ZedCloudOpsCmdModelFromMap(purgeMap[0].(map[string]interface{}))
		}
	}
	//
	sizeBytesInt := m["size_bytes"].(int)
	sizeBytes := fmt.Sprintf("%d", sizeBytesInt)
	var typeVar *models.VolumeInstanceType // VolumeInstanceType
	typeInterface, typeIsSet := m["type"]
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewVolumeInstanceType(models.VolumeInstanceType(typeModel))
	}
	return &models.ProfileVolumeConfig{
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

func SetProfileVolumeConfigResourceData(d *schema.ResourceData, m *models.ProfileVolumeConfig) {
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

func SetProfileVolumeConfigSubResourceData(m []*models.ProfileVolumeConfig) (d []*map[string]interface{}) {
	for _, ProfileVolumeConfigModel := range m {
		if ProfileVolumeConfigModel != nil {
			properties := make(map[string]interface{})
			properties["accessmode"] = ProfileVolumeConfigModel.Accessmode
			properties["block_storage_tags"] = ProfileVolumeConfigModel.BlockStorageTags
			properties["cleartext"] = ProfileVolumeConfigModel.Cleartext
			properties["content_tree_id"] = ProfileVolumeConfigModel.ContentTreeID
			properties["content_tree_target_condition"] = ProfileVolumeConfigModel.ContentTreeTargetCondition
			properties["image"] = ProfileVolumeConfigModel.Image
			properties["label"] = ProfileVolumeConfigModel.Label
			properties["multiattach"] = ProfileVolumeConfigModel.Multiattach
			properties["purge"] = SetZedCloudOpsCmdSubResourceData([]*models.ZedCloudOpsCmd{ProfileVolumeConfigModel.Purge})
			properties["size_bytes"] = ProfileVolumeConfigModel.SizeBytes
			properties["type"] = ProfileVolumeConfigModel.Type
			d = append(d, &properties)
		}
	}
	return
}

func ProfileVolumeConfigSchema() map[string]*schema.Schema {
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

func GetProfileVolumeConfigPropertyFields() (t []string) {
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
