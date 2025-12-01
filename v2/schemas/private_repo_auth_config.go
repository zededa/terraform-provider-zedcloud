package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func PrivateRepoAuthConfigModel(d *schema.ResourceData) *models.PrivateRepoAuthConfig {
	var secret *models.PrivateRepoSecretConfig // PrivateRepoSecretConfig
	secretInterface, secretIsSet := d.GetOk("secret")
	if secretIsSet && secretInterface != nil {
		secretMap := secretInterface.([]interface{})
		if len(secretMap) > 0 {
			secret = PrivateRepoSecretConfigModelFromMap(secretMap[0].(map[string]interface{}))
		}
	}
	return &models.PrivateRepoAuthConfig{
		Secret: secret,
	}
}

func PrivateRepoAuthConfigModelFromMap(m map[string]interface{}) *models.PrivateRepoAuthConfig {
	var secret *models.PrivateRepoSecretConfig // PrivateRepoSecretConfig
	secretInterface, secretIsSet := m["secret"]
	if secretIsSet && secretInterface != nil {
		secretMap := secretInterface.([]interface{})
		if len(secretMap) > 0 {
			secret = PrivateRepoSecretConfigModelFromMap(secretMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.PrivateRepoAuthConfig{
		Secret: secret,
	}
}

func SetPrivateRepoAuthConfigResourceData(d *schema.ResourceData, m *models.PrivateRepoAuthConfig) {
	d.Set("secret", SetPrivateRepoSecretConfigSubResourceData([]*models.PrivateRepoSecretConfig{m.Secret}))
}

func SetPrivateRepoAuthConfigSubResourceData(m []*models.PrivateRepoAuthConfig) (d []*map[string]interface{}) {
	for _, PrivateRepoAuthConfigModel := range m {
		if PrivateRepoAuthConfigModel != nil {
			properties := make(map[string]interface{})
			properties["secret"] = SetPrivateRepoSecretConfigSubResourceData([]*models.PrivateRepoSecretConfig{PrivateRepoAuthConfigModel.Secret})
			d = append(d, &properties)
		}
	}
	return
}

func PrivateRepoAuthConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"secret": {
			Description: `Secret configuration for authentication`,
			Type:        schema.TypeList, //GoType: PrivateRepoSecretConfig
			Elem: &schema.Resource{
				Schema: PrivateRepoSecretConfigSchema(),
			},
			Optional: true,
		},
	}
}

func GetPrivateRepoAuthConfigPropertyFields() (t []string) {
	return []string{
		"secret",
	}
}
