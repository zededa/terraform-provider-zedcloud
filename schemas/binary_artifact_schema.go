package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func BinaryArtifactModel(d *schema.ResourceData) *models.BinaryArtifact {
	artifactMetaData, _ := d.Get("artifact_meta_data").(string)
	var base64Artifact *models.InlineOpaqueBase64Data // InlineOpaqueBase64Data
	base64ArtifactInterface, base64ArtifactIsSet := d.GetOk("base64_artifact")
	if base64ArtifactIsSet && base64ArtifactInterface != nil {
		base64ArtifactMap := base64ArtifactInterface.([]interface{})
		if len(base64ArtifactMap) > 0 {
			base64Artifact = InlineOpaqueBase64DataModelFromMap(base64ArtifactMap[0].(map[string]interface{}))
		}
	}
	var binaryArtifact *models.ExternalOpaqueBinaryBlob // ExternalOpaqueBinaryBlob
	binaryArtifactInterface, binaryArtifactIsSet := d.GetOk("binary_artifact")
	if binaryArtifactIsSet && binaryArtifactInterface != nil {
		binaryArtifactMap := binaryArtifactInterface.([]interface{})
		if len(binaryArtifactMap) > 0 {
			binaryArtifact = ExternalOpaqueBinaryBlobModelFromMap(binaryArtifactMap[0].(map[string]interface{}))
		}
	}
	var format *models.OpaqueObjectCategory // OpaqueObjectCategory
	formatInterface, formatIsSet := d.GetOk("format")
	if formatIsSet {
		formatModel := formatInterface.(string)
		format = models.NewOpaqueObjectCategory(models.OpaqueObjectCategory(formatModel))
	}
	id, _ := d.Get("id").(string)
	return &models.BinaryArtifact{
		ArtifactMetaData: artifactMetaData,
		Base64Artifact:   base64Artifact,
		BinaryArtifact:   binaryArtifact,
		Format:           format,
		ID:               id,
	}
}

func BinaryArtifactModelFromMap(m map[string]interface{}) *models.BinaryArtifact {
	artifactMetaData := m["artifact_meta_data"].(string)
	var base64Artifact *models.InlineOpaqueBase64Data // InlineOpaqueBase64Data
	base64ArtifactInterface, base64ArtifactIsSet := m["base64_artifact"]
	if base64ArtifactIsSet && base64ArtifactInterface != nil {
		base64ArtifactMap := base64ArtifactInterface.([]interface{})
		if len(base64ArtifactMap) > 0 {
			base64Artifact = InlineOpaqueBase64DataModelFromMap(base64ArtifactMap[0].(map[string]interface{}))
		}
	}
	//
	var binaryArtifact *models.ExternalOpaqueBinaryBlob // ExternalOpaqueBinaryBlob
	binaryArtifactInterface, binaryArtifactIsSet := m["binary_artifact"]
	if binaryArtifactIsSet && binaryArtifactInterface != nil {
		binaryArtifactMap := binaryArtifactInterface.([]interface{})
		if len(binaryArtifactMap) > 0 {
			binaryArtifact = ExternalOpaqueBinaryBlobModelFromMap(binaryArtifactMap[0].(map[string]interface{}))
		}
	}
	//
	var format *models.OpaqueObjectCategory // OpaqueObjectCategory
	formatInterface, formatIsSet := m["format"]
	if formatIsSet {
		formatModel := formatInterface.(string)
		format = models.NewOpaqueObjectCategory(models.OpaqueObjectCategory(formatModel))
	}
	id := m["id"].(string)
	return &models.BinaryArtifact{
		ArtifactMetaData: artifactMetaData,
		Base64Artifact:   base64Artifact,
		BinaryArtifact:   binaryArtifact,
		Format:           format,
		ID:               id,
	}
}

func SetBinaryArtifactResourceData(d *schema.ResourceData, m *models.BinaryArtifact) {
	d.Set("artifact_meta_data", m.ArtifactMetaData)
	d.Set("base64_artifact", SetInlineOpaqueBase64DataSubResourceData([]*models.InlineOpaqueBase64Data{m.Base64Artifact}))
	d.Set("binary_artifact", SetExternalOpaqueBinaryBlobSubResourceData([]*models.ExternalOpaqueBinaryBlob{m.BinaryArtifact}))
	d.Set("format", m.Format)
	d.Set("id", m.ID)
}

func SetBinaryArtifactSubResourceData(m []*models.BinaryArtifact) (d []*map[string]interface{}) {
	for _, BinaryArtifactModel := range m {
		if BinaryArtifactModel != nil {
			properties := make(map[string]interface{})
			properties["artifact_meta_data"] = BinaryArtifactModel.ArtifactMetaData
			properties["base64_artifact"] = SetInlineOpaqueBase64DataSubResourceData([]*models.InlineOpaqueBase64Data{BinaryArtifactModel.Base64Artifact})
			properties["binary_artifact"] = SetExternalOpaqueBinaryBlobSubResourceData([]*models.ExternalOpaqueBinaryBlob{BinaryArtifactModel.BinaryArtifact})
			properties["format"] = BinaryArtifactModel.Format
			properties["id"] = BinaryArtifactModel.ID
			d = append(d, &properties)
		}
	}
	return
}

func BinaryArtifactSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"artifact_meta_data": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"base64_artifact": {
			Description: ``,
			Type:        schema.TypeList, //GoType: InlineOpaqueBase64Data
			Elem: &schema.Resource{
				Schema: InlineOpaqueBase64DataSchema(),
			},
			Optional: true,
		},

		"binary_artifact": {
			Description: ``,
			Type:        schema.TypeList, //GoType: ExternalOpaqueBinaryBlob
			Elem: &schema.Resource{
				Schema: ExternalOpaqueBinaryBlobSchema(),
			},
			Optional: true,
		},

		"format": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"id": {
			Description: `System defined universally unique Id of the artifact.`,
			Type:        schema.TypeString,
			Computed:    true,
		},
	}
}

func GetBinaryArtifactPropertyFields() (t []string) {
	return []string{
		"artifact_meta_data",
		"base64_artifact",
		"binary_artifact",
		"format",
		"id",
	}
}
