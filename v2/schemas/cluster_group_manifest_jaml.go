package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ClusterGroupManifestJamlModel(d *schema.ResourceData) *models.ClusterGroupManifestJaml {
	yamlContent, _ := d.Get("yaml_content").(string)
	return &models.ClusterGroupManifestJaml{
		YamlContent: yamlContent,
	}
}

func ClusterGroupManifestJamlModelFromMap(m map[string]interface{}) *models.ClusterGroupManifestJaml {
	yamlContent := m["yaml_content"].(string)
	return &models.ClusterGroupManifestJaml{
		YamlContent: yamlContent,
	}
}

func SetClusterGroupManifestJamlResourceData(d *schema.ResourceData, m *models.ClusterGroupManifestJaml) {
	d.Set("yaml_content", m.YamlContent)
}

func SetClusterGroupManifestJamlSubResourceData(m []*models.ClusterGroupManifestJaml) (d []*map[string]interface{}) {
	for _, ClusterGroupManifestJamlModel := range m {
		if ClusterGroupManifestJamlModel != nil {
			properties := make(map[string]interface{})
			properties["yaml_content"] = ClusterGroupManifestJamlModel.YamlContent
			d = append(d, &properties)
		}
	}
	return
}

func ClusterGroupManifestJamlSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"yaml_content": {
			Description: `The content of an jaml manifest`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetClusterGroupManifestJamlPropertyFields() (t []string) {
	return []string{
		"yaml_content",
	}
}
