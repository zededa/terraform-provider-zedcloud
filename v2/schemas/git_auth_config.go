package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func GitAuthConfigModel(d *schema.ResourceData) *models.GitAuthConfig {
	var secret *models.GitRepoSecretConfig // GitRepoSecretConfig
	secretInterface, secretIsSet := d.GetOk("secret")
	if secretIsSet && secretInterface != nil {
		secretMap := secretInterface.([]interface{})
		if len(secretMap) > 0 {
			secret = GitRepoSecretConfigModelFromMap(secretMap[0].(map[string]interface{}))
		}
	}
	return &models.GitAuthConfig{
		Secret: secret,
	}
}

func GitAuthConfigModelFromMap(m map[string]interface{}) *models.GitAuthConfig {
	var secret *models.GitRepoSecretConfig // GitRepoSecretConfig
	secretInterface, secretIsSet := m["secret"]
	if secretIsSet && secretInterface != nil {
		secretMap := secretInterface.([]interface{})
		if len(secretMap) > 0 {
			secret = GitRepoSecretConfigModelFromMap(secretMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.GitAuthConfig{
		Secret: secret,
	}
}

func SetGitAuthConfigResourceData(d *schema.ResourceData, m *models.GitAuthConfig) {
	d.Set("secret", SetGitRepoSecretConfigSubResourceData([]*models.GitRepoSecretConfig{m.Secret}))
}

func SetGitAuthConfigSubResourceData(m []*models.GitAuthConfig) (d []*map[string]interface{}) {
	for _, GitAuthConfigModel := range m {
		if GitAuthConfigModel != nil {
			properties := make(map[string]interface{})
			properties["secret"] = SetGitRepoSecretConfigSubResourceData([]*models.GitRepoSecretConfig{GitAuthConfigModel.Secret})
			d = append(d, &properties)
		}
	}
	return
}

func GitAuthConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"secret": {
			Description: `Secret configuration for Git authentication`,
			Type:        schema.TypeList, //GoType: GitRepoSecretConfig
			Elem: &schema.Resource{
				Schema: GitRepoSecretConfigSchema(),
			},
			Optional: true,
		},
	}
}

func GetGitAuthConfigPropertyFields() (t []string) {
	return []string{
		"secret",
	}
}
