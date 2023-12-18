package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AAASuccessResponseCredentialChangeModel(d *schema.ResourceData) *models.AAASuccessResponseCredentialChange {
	credentialID, _ := d.Get("credential_id").(string)
	return &models.AAASuccessResponseCredentialChange{
		CredentialID: credentialID,
	}
}

func AAASuccessResponseCredentialChangeModelFromMap(m map[string]interface{}) *models.AAASuccessResponseCredentialChange {
	credentialID := m["credential_id"].(string)
	return &models.AAASuccessResponseCredentialChange{
		CredentialID: credentialID,
	}
}

func SetAAASuccessResponseCredentialChangeResourceData(d *schema.ResourceData, m *models.AAASuccessResponseCredentialChange) {
	d.Set("credential_id", m.CredentialID)
}

func SetAAASuccessResponseCredentialChangeSubResourceData(m []*models.AAASuccessResponseCredentialChange) (d []*map[string]interface{}) {
	for _, AAASuccessResponseCredentialChangeModel := range m {
		if AAASuccessResponseCredentialChangeModel != nil {
			properties := make(map[string]interface{})
			properties["credential_id"] = AAASuccessResponseCredentialChangeModel.CredentialID
			d = append(d, &properties)
		}
	}
	return
}

func AAASuccessResponseCredentialChangeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"credential_id": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetAAASuccessResponseCredentialChangePropertyFields() (t []string) {
	return []string{
		"credential_id",
	}
}
