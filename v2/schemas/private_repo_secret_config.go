package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func PrivateRepoSecretConfigModel(d *schema.ResourceData) *models.PrivateRepoSecretConfig {
	name, _ := d.Get("name").(string)
	var typeVar *models.SecretType // SecretType
	typeInterface, typeIsSet := d.GetOk("type")
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewSecretType(models.SecretType(typeModel))
	}
	return &models.PrivateRepoSecretConfig{
		Name: name,
		Type: typeVar,
	}
}

func PrivateRepoSecretConfigModelFromMap(m map[string]interface{}) *models.PrivateRepoSecretConfig {
	name := m["name"].(string)
	var typeVar *models.SecretType // SecretType
	typeInterface, typeIsSet := m["type"]
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewSecretType(models.SecretType(typeModel))
	}
	return &models.PrivateRepoSecretConfig{
		Name: name,
		Type: typeVar,
	}
}

func SetPrivateRepoSecretConfigResourceData(d *schema.ResourceData, m *models.PrivateRepoSecretConfig) {
	d.Set("name", m.Name)
	d.Set("type", m.Type)
}

func SetPrivateRepoSecretConfigSubResourceData(m []*models.PrivateRepoSecretConfig) (d []*map[string]interface{}) {
	for _, PrivateRepoSecretConfigModel := range m {
		if PrivateRepoSecretConfigModel != nil {
			properties := make(map[string]interface{})
			properties["name"] = PrivateRepoSecretConfigModel.Name
			properties["type"] = PrivateRepoSecretConfigModel.Type
			d = append(d, &properties)
		}
	}
	return
}

func PrivateRepoSecretConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Description: `Name of the secret resource`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"type": {
			Description: `Type of the secret (SECRET_TYPE_SSH, SECRET_TYPE_BASIC_AUTH, SECRET_TYPE_UNSPECIFIED, SECRET_TYPE_NONE). This field is required for private repositories and must be a valid Kubernetes secret type. For a public repository, set type as SECRET_TYPE_NONE`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetPrivateRepoSecretConfigPropertyFields() (t []string) {
	return []string{
		"name",
		"type",
	}
}
