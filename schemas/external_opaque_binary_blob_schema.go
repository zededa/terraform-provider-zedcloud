package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func ExternalOpaqueBinaryBlobModel(d *schema.ResourceData) *models.ExternalOpaqueBinaryBlob {
	blobMetaData, _ := d.Get("blob_meta_data").(string)
	fileNameToUse, _ := d.Get("file_name_to_use").(string)
	imageID, _ := d.Get("image_id").(string)
	imageName, _ := d.Get("image_name").(string)
	return &models.ExternalOpaqueBinaryBlob{
		BlobMetaData:  blobMetaData,
		FileNameToUse: fileNameToUse,
		ImageID:       imageID,
		ImageName:     imageName,
	}
}

func ExternalOpaqueBinaryBlobModelFromMap(m map[string]interface{}) *models.ExternalOpaqueBinaryBlob {
	blobMetaData := m["blob_meta_data"].(string)
	fileNameToUse := m["file_name_to_use"].(string)
	imageID := m["image_id"].(string)
	imageName := m["image_name"].(string)
	return &models.ExternalOpaqueBinaryBlob{
		BlobMetaData:  blobMetaData,
		FileNameToUse: fileNameToUse,
		ImageID:       imageID,
		ImageName:     imageName,
	}
}

func SetExternalOpaqueBinaryBlobResourceData(d *schema.ResourceData, m *models.ExternalOpaqueBinaryBlob) {
	d.Set("blob_meta_data", m.BlobMetaData)
	d.Set("file_name_to_use", m.FileNameToUse)
	d.Set("image_id", m.ImageID)
	d.Set("image_name", m.ImageName)
}

func SetExternalOpaqueBinaryBlobSubResourceData(m []*models.ExternalOpaqueBinaryBlob) (d []*map[string]interface{}) {
	for _, ExternalOpaqueBinaryBlobModel := range m {
		if ExternalOpaqueBinaryBlobModel != nil {
			properties := make(map[string]interface{})
			properties["blob_meta_data"] = ExternalOpaqueBinaryBlobModel.BlobMetaData
			properties["file_name_to_use"] = ExternalOpaqueBinaryBlobModel.FileNameToUse
			properties["image_id"] = ExternalOpaqueBinaryBlobModel.ImageID
			properties["image_name"] = ExternalOpaqueBinaryBlobModel.ImageName
			d = append(d, &properties)
		}
	}
	return
}

func ExternalOpaqueBinaryBlobSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"blob_meta_data": {
			Description: `Eve shall use the image name when fileNameToUse is empty
optional - this can be image type or size etc encoded into a single string.
only the application instance agents will interpret this.`,
			Type:     schema.TypeString,
			Optional: true,
		},

		"file_name_to_use": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"image_id": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"image_name": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetExternalOpaqueBinaryBlobPropertyFields() (t []string) {
	return []string{
		"blob_meta_data",
		"file_name_to_use",
		"image_id",
		"image_name",
	}
}
