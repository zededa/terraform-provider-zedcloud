package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func PrivateRepoAuthInfoModel(d *schema.ResourceData) *models.PrivateRepoAuthInfo {
	var secret *models.PrivateRepoSecretInfo // PrivateRepoSecretInfo
	secretInterface, secretIsSet := d.GetOk("secret")
	if secretIsSet && secretInterface != nil {
		secretMap := secretInterface.([]interface{})
		if len(secretMap) > 0 {
			secret = PrivateRepoSecretInfoModelFromMap(secretMap[0].(map[string]interface{}))
		}
	}
	return &models.PrivateRepoAuthInfo{
		Secret: secret,
	}
}

func PrivateRepoAuthInfoModelFromMap(m map[string]interface{}) *models.PrivateRepoAuthInfo {
	var secret *models.PrivateRepoSecretInfo // PrivateRepoSecretInfo
	secretInterface, secretIsSet := m["secret"]
	if secretIsSet && secretInterface != nil {
		secretMap := secretInterface.([]interface{})
		if len(secretMap) > 0 {
			secret = PrivateRepoSecretInfoModelFromMap(secretMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.PrivateRepoAuthInfo{
		Secret: secret,
	}
}

func SetPrivateRepoAuthInfoResourceData(d *schema.ResourceData, m *models.PrivateRepoAuthInfo) {
	d.Set("secret", SetPrivateRepoSecretInfoSubResourceData([]*models.PrivateRepoSecretInfo{m.Secret}))
}

func SetPrivateRepoAuthInfoSubResourceData(m []*models.PrivateRepoAuthInfo) (d []*map[string]interface{}) {
	for _, PrivateRepoAuthInfoModel := range m {
		if PrivateRepoAuthInfoModel != nil {
			properties := make(map[string]interface{})
			properties["secret"] = SetPrivateRepoSecretInfoSubResourceData([]*models.PrivateRepoSecretInfo{PrivateRepoAuthInfoModel.Secret})
			d = append(d, &properties)
		}
	}
	return
}

func PrivateRepoAuthInfoSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"secret": {
			Description: `Enhanced secret information for authentication`,
			Type:        schema.TypeList, //GoType: PrivateRepoSecretInfo
			Elem: &schema.Resource{
				Schema: PrivateRepoSecretInfoSchema(),
			},
			Optional: true,
		},
	}
}

func GetPrivateRepoAuthInfoPropertyFields() (t []string) {
	return []string{
		"secret",
	}
}
