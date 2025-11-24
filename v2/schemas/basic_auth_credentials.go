package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func BasicAuthCredentialsModel(d *schema.ResourceData) *models.BasicAuthCredentials {
	password, _ := d.Get("password").(string)
	username, _ := d.Get("username").(string)
	return &models.BasicAuthCredentials{
		Password: password,
		Username: username,
	}
}

func BasicAuthCredentialsModelFromMap(m map[string]interface{}) *models.BasicAuthCredentials {
	password := m["password"].(string)
	username := m["username"].(string)
	return &models.BasicAuthCredentials{
		Password: password,
		Username: username,
	}
}

func SetBasicAuthCredentialsResourceData(d *schema.ResourceData, m *models.BasicAuthCredentials) {
	d.Set("password", m.Password)
	d.Set("username", m.Username)
}

func SetBasicAuthCredentialsSubResourceData(m []*models.BasicAuthCredentials) (d []*map[string]interface{}) {
	for _, BasicAuthCredentialsModel := range m {
		if BasicAuthCredentialsModel != nil {
			properties := make(map[string]interface{})
			properties["password"] = BasicAuthCredentialsModel.Password
			properties["username"] = BasicAuthCredentialsModel.Username
			d = append(d, &properties)
		}
	}
	return
}

func BasicAuthCredentialsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"password": {
			Description: `Password for basic authentication`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"username": {
			Description: `Username for basic authentication`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetBasicAuthCredentialsPropertyFields() (t []string) {
	return []string{
		"password",
		"username",
	}
}
