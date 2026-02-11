package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func GitRepoSecretConfigModel(d *schema.ResourceData) *models.GitRepoSecretConfig {
	name, _ := d.Get("name").(string)
	var typeVar *models.SecretType // SecretType
	typeInterface, typeIsSet := d.GetOk("type")
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewSecretType(models.SecretType(typeModel))
	}
	return &models.GitRepoSecretConfig{
		Name: name,
		Type: typeVar,
	}
}

func GitRepoSecretConfigModelFromMap(m map[string]interface{}) *models.GitRepoSecretConfig {
	name := m["name"].(string)
	var typeVar *models.SecretType // SecretType
	typeInterface, typeIsSet := m["type"]
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewSecretType(models.SecretType(typeModel))
	}
	return &models.GitRepoSecretConfig{
		Name: name,
		Type: typeVar,
	}
}

func SetGitRepoSecretConfigResourceData(d *schema.ResourceData, m *models.GitRepoSecretConfig) {
	d.Set("name", m.Name)
	d.Set("type", m.Type)
}

func SetGitRepoSecretConfigSubResourceData(m []*models.GitRepoSecretConfig) (d []*map[string]interface{}) {
	for _, GitRepoSecretConfigModel := range m {
		if GitRepoSecretConfigModel != nil {
			properties := make(map[string]interface{})
			properties["name"] = GitRepoSecretConfigModel.Name
			properties["type"] = GitRepoSecretConfigModel.Type
			d = append(d, &properties)
		}
	}
	return
}

func GitRepoSecretConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Description: `Name of the secret`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"type": {
			Description: `Type of secret for authentication(SECRET_TYPE_SSH, SECRET_TYPE_BASIC_AUTH, SECRET_TYPE_UNSPECIFIED, SECRET_TYPE_NONE). For a Public Repository, set type as SECRET_TYPE_NONE`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetGitRepoSecretConfigPropertyFields() (t []string) {
	return []string{
		"name",
		"type",
	}
}
