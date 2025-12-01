package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func HelmAuthConfigModel(d *schema.ResourceData) *models.HelmAuthConfig {
	helmRepoURLRegex, _ := d.Get("helm_repo_url_regex").(string)
	var secret *models.GitRepoSecretConfig // GitRepoSecretConfig
	secretInterface, secretIsSet := d.GetOk("secret")
	if secretIsSet && secretInterface != nil {
		secretMap := secretInterface.([]interface{})
		if len(secretMap) > 0 {
			secret = GitRepoSecretConfigModelFromMap(secretMap[0].(map[string]interface{}))
		}
	}
	return &models.HelmAuthConfig{
		HelmRepoURLRegex: helmRepoURLRegex,
		Secret:           secret,
	}
}

func HelmAuthConfigModelFromMap(m map[string]interface{}) *models.HelmAuthConfig {
	helmRepoURLRegex := m["helm_repo_url_regex"].(string)
	var secret *models.GitRepoSecretConfig // GitRepoSecretConfig
	secretInterface, secretIsSet := m["secret"]
	if secretIsSet && secretInterface != nil {
		secretMap := secretInterface.([]interface{})
		if len(secretMap) > 0 {
			secret = GitRepoSecretConfigModelFromMap(secretMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.HelmAuthConfig{
		HelmRepoURLRegex: helmRepoURLRegex,
		Secret:           secret,
	}
}

func SetHelmAuthConfigResourceData(d *schema.ResourceData, m *models.HelmAuthConfig) {
	d.Set("helm_repo_url_regex", m.HelmRepoURLRegex)
	d.Set("secret", SetGitRepoSecretConfigSubResourceData([]*models.GitRepoSecretConfig{m.Secret}))
}

func SetHelmAuthConfigSubResourceData(m []*models.HelmAuthConfig) (d []*map[string]interface{}) {
	for _, HelmAuthConfigModel := range m {
		if HelmAuthConfigModel != nil {
			properties := make(map[string]interface{})
			properties["helm_repo_url_regex"] = HelmAuthConfigModel.HelmRepoURLRegex
			properties["secret"] = SetGitRepoSecretConfigSubResourceData([]*models.GitRepoSecretConfig{HelmAuthConfigModel.Secret})
			d = append(d, &properties)
		}
	}
	return
}

func HelmAuthConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"helm_repo_url_regex": {
			Description: `Regular expression pattern to match Helm repository URLs`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"secret": {
			Description: `Secret configuration for Helm authentication`,
			Type:        schema.TypeList, //GoType: GitRepoSecretConfig
			Elem: &schema.Resource{
				Schema: GitRepoSecretConfigSchema(),
			},
			Optional: true,
		},
	}
}

func GetHelmAuthConfigPropertyFields() (t []string) {
	return []string{
		"helm_repo_url_regex",
		"secret",
	}
}
