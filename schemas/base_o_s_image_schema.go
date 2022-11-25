package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate BaseOSImage resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func BaseOSImageModel(d *schema.ResourceData) *models.BaseOSImage {
	activate := d.Get("activate").(bool)
	imageName := d.Get("image_name").(string)
	uuid := d.Get("uuid").(string)
	version := d.Get("version").(string)
	return &models.BaseOSImage{
		Activate:  activate,
		ImageName: &imageName, // string true false false
		UUID:      &uuid,      // string true false false
		Version:   &version,   // string true false false
	}
}

func BaseOSImageModelFromMap(m map[string]interface{}) *models.BaseOSImage {
	activate := m["activate"].(bool)
	imageName := m["image_name"].(string)
	uuid := m["uuid"].(string)
	version := m["version"].(string)
	return &models.BaseOSImage{
		Activate:  activate,
		ImageName: &imageName,
		UUID:      &uuid,
		Version:   &version,
	}
}

// Update the underlying BaseOSImage resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetBaseOSImageResourceData(d *schema.ResourceData, m *models.BaseOSImage) {
	d.Set("activate", m.Activate)
	d.Set("image_name", m.ImageName)
	d.Set("uuid", m.UUID)
	d.Set("version", m.Version)
}

// Iterate throught and update the BaseOSImage resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetBaseOSImageSubResourceData(m []*models.BaseOSImage) (d []*map[string]interface{}) {
	for _, BaseOSImageModel := range m {
		if BaseOSImageModel != nil {
			properties := make(map[string]interface{})
			properties["activate"] = BaseOSImageModel.Activate
			properties["image_name"] = BaseOSImageModel.ImageName
			properties["uuid"] = BaseOSImageModel.UUID
			properties["version"] = BaseOSImageModel.Version
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the BaseOSImage resource defined in the Terraform configuration
func BaseOSImageSchema() map[string]*schema.Schema {
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
		"uuid",
		"version",
	}
}
