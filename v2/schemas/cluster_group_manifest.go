package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ClusterGroupManifestModel(d *schema.ResourceData) *models.ClusterGroupManifest {
	yamlContent, _ := d.Get("yaml_content").(string)
	return &models.ClusterGroupManifest{
		YamlContent: &yamlContent, // string
	}
}

func ClusterGroupManifestModelFromMap(m map[string]interface{}) *models.ClusterGroupManifest {
	yamlContent := m["yaml_content"].(string)
	return &models.ClusterGroupManifest{
		YamlContent: &yamlContent,
	}
}

func SetClusterGroupManifestResourceData(d *schema.ResourceData, m *models.ClusterGroupManifest) {
	d.Set("yaml_content", m.YamlContent)
}

func SetClusterGroupManifestSubResourceData(m []*models.ClusterGroupManifest) (d []*map[string]interface{}) {
	for _, ClusterGroupManifestModel := range m {
		if ClusterGroupManifestModel != nil {
			properties := make(map[string]interface{})
			properties["yaml_content"] = ClusterGroupManifestModel.YamlContent
			d = append(d, &properties)
		}
	}
	return
}

func ClusterGroupManifestSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"yaml_content": {
			Description: `The content of an yaml manifest`,
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

func GetClusterGroupManifestPropertyFields() (t []string) {
	return []string{
		"yaml_content",
	}
}
