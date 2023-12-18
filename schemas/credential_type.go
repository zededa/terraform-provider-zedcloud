package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func CredentialTypeModel(d *schema.ResourceData) *models.CredentialType {
	credentialType, _ := d.Get("credential_type").(models.CredentialType)
	return &credentialType
}

func CredentialTypeModelFromMap(m map[string]interface{}) *models.CredentialType {
	credentialType := m["credential_type"].(models.CredentialType)
	return &credentialType
}

func SetCredentialTypeResourceData(d *schema.ResourceData, m *models.CredentialType) {
}

func SetCredentialTypeSubResourceData(m []*models.CredentialType) (d []*map[string]interface{}) {
	for _, CredentialTypeModel := range m {
		if CredentialTypeModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func CredentialTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func GetCredentialTypePropertyFields() (t []string) {
	return []string{}
}
