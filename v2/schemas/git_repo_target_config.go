package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func GitRepoTargetConfigModel(d *schema.ResourceData) *models.GitRepoTargetConfig {
	var typeVar *models.GitRepoTargetType // GitRepoTargetType
	typeInterface, typeIsSet := d.GetOk("type")
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewGitRepoTargetType(models.GitRepoTargetType(typeModel))
	}
	value, _ := d.Get("value").(string)
	return &models.GitRepoTargetConfig{
		Type:  typeVar,
		Value: value,
	}
}

func GitRepoTargetConfigModelFromMap(m map[string]interface{}) *models.GitRepoTargetConfig {
	var typeVar *models.GitRepoTargetType // GitRepoTargetType
	typeInterface, typeIsSet := m["type"]
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewGitRepoTargetType(models.GitRepoTargetType(typeModel))
	}
	value := m["value"].(string)
	return &models.GitRepoTargetConfig{
		Type:  typeVar,
		Value: value,
	}
}

func SetGitRepoTargetConfigResourceData(d *schema.ResourceData, m *models.GitRepoTargetConfig) {
	d.Set("type", m.Type)
	d.Set("value", m.Value)
}

func SetGitRepoTargetConfigSubResourceData(m []*models.GitRepoTargetConfig) (d []*map[string]interface{}) {
	for _, GitRepoTargetConfigModel := range m {
		if GitRepoTargetConfigModel != nil {
			properties := make(map[string]interface{})
			properties["type"] = GitRepoTargetConfigModel.Type
			properties["value"] = GitRepoTargetConfigModel.Value
			d = append(d, &properties)
		}
	}
	return
}

func GitRepoTargetConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"type": {
			Description: `Type of target for GitRepo deployment (required)`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"value": {
			Description: `Target value (cluster ID, cluster group ID, etc.). Required when type is CLUSTER or CLUSTER_GROUP, not required for ALL_CLUSTER`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetGitRepoTargetConfigPropertyFields() (t []string) {
	return []string{
		"type",
		"value",
	}
}
