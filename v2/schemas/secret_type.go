package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func SecretTypeModel(d *schema.ResourceData) *models.SecretType {
	secretType, _ := d.Get("secret_type").(models.SecretType)
	return &secretType
}

func SecretTypeModelFromMap(m map[string]interface{}) *models.SecretType {
	secretType := m["secret_type"].(models.SecretType)
	return &secretType
}

func SetSecretTypeResourceData(d *schema.ResourceData, m *models.SecretType) {
}

func SetSecretTypeSubResourceData(m []*models.SecretType) (d []*map[string]interface{}) {
	for _, SecretTypeModel := range m {
		if SecretTypeModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func SecretTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func GetSecretTypePropertyFields() (t []string) {
	return []string{}
}
