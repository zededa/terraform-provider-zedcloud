package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func DatastoreSecretsModel(d *schema.ResourceData) *models.DatastoreInfoSecrets {
	aPIKey, _ := d.Get("api_key").(string)
	aPIPasswd, _ := d.Get("api_passwd").(string)
	return &models.DatastoreInfoSecrets{
		APIKey:    aPIKey,
		APIPasswd: aPIPasswd,
	}
}

func DatastoreSecretsModelFromMap(m map[string]interface{}) *models.DatastoreInfoSecrets {
	aPIKey := m["api_key"].(string)
	aPIPasswd := m["api_passwd"].(string)
	return &models.DatastoreInfoSecrets{
		APIKey:    aPIKey,
		APIPasswd: aPIPasswd,
	}
}

func SetDatastoreSecretsResourceData(d *schema.ResourceData, m *models.DatastoreInfoSecrets) {
	d.Set("api_key", m.APIKey)
	d.Set("api_passwd", m.APIPasswd)
}

func SetDatastoreSecretsSubResourceData(m []*models.DatastoreInfoSecrets) (d []*map[string]interface{}) {
	for _, DatastoreInfoSecretsModel := range m {
		if DatastoreInfoSecretsModel != nil {
			properties := make(map[string]interface{})
			properties["api_key"] = DatastoreInfoSecretsModel.APIKey
			properties["api_passwd"] = DatastoreInfoSecretsModel.APIPasswd
			d = append(d, &properties)
		}
	}
	return
}

func DatastoreSecrets() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"api_key": {
			Description: `Datastore access API key in plain-text`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"api_passwd": {
			Description: `Datastore access API password in plain-text`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the DatastoreInfoSecrets resource
func GetDatastoreSecretsPropertyFields() (t []string) {
	return []string{
		"api_key",
		"api_passwd",
	}
}
