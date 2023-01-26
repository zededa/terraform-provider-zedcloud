package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func BaseOSImageModel(d *schema.ResourceData) *models.BaseOSImage {
	activate, _ := d.Get("activate").(bool)
	imageName, _ := d.Get("image_name").(string)
	imvolID, _ := d.Get("imvol_id").(string)
	uuid, _ := d.Get("uuid").(string)
	version, _ := d.Get("version").(string)
	return &models.BaseOSImage{
		Activate:  &activate,
		ImageName: &imageName,
		ImvolID:   imvolID,
		UUID:      &uuid,
		Version:   &version,
	}
}

func BaseOSImageModelFromMap(m map[string]interface{}) *models.BaseOSImage {
	activate := m["activate"].(bool)
	imageName := m["image_name"].(string)
	imvolID := m["imvol_id"].(string)
	uuid := m["uuid"].(string)
	version := m["version"].(string)
	return &models.BaseOSImage{
		Activate:  &activate,
		ImageName: &imageName,
		ImvolID:   imvolID,
		UUID:      &uuid,
		Version:   &version,
	}
}

func SetBaseOSImageResourceData(d *schema.ResourceData, m *models.BaseOSImage) {
	d.Set("activate", m.Activate)
	d.Set("image_name", m.ImageName)
	d.Set("imvol_id", m.ImvolID)
	d.Set("uuid", m.UUID)
	d.Set("version", m.Version)
}

func SetBaseOSImageSubResourceData(m []*models.BaseOSImage) (d []*map[string]interface{}) {
	for _, BaseOSImageModel := range m {
		if BaseOSImageModel != nil {
			properties := make(map[string]interface{})
			properties["activate"] = BaseOSImageModel.Activate
			properties["image_name"] = BaseOSImageModel.ImageName
			properties["imvol_id"] = BaseOSImageModel.ImvolID
			properties["uuid"] = BaseOSImageModel.UUID
			properties["version"] = BaseOSImageModel.Version
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the BaseOSImage resource defined in the Terraform configuration
func BaseOSImage() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"activate": {
			Description: `activation flag`,
			Type:        schema.TypeBool,
			Required:    true,
		},

		"image_name": {
			Description: `image name`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"imvol_id": {
			Description: `immutable Volume for this base image`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"uuid": {
			Description: `system generated unique id for an image`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"version": {
			Description: `image version`,
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

// Retrieve property field names for updating the BaseOSImage resource
func GetBaseOSImagePropertyFields() (t []string) {
	return []string{
		"activate",
		"image_name",
		"imvol_id",
		"uuid",
		"version",
	}
}
