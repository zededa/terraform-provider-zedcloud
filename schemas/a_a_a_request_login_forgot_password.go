package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AAARequestLoginForgotPasswordModel(d *schema.ResourceData) *models.AAARequestLoginForgotPassword {
	username, _ := d.Get("username").(string)
	return &models.AAARequestLoginForgotPassword{
		Username: username,
	}
}

func AAARequestLoginForgotPasswordModelFromMap(m map[string]interface{}) *models.AAARequestLoginForgotPassword {
	username := m["username"].(string)
	return &models.AAARequestLoginForgotPassword{
		Username: username,
	}
}

func SetAAARequestLoginForgotPasswordResourceData(d *schema.ResourceData, m *models.AAARequestLoginForgotPassword) {
	d.Set("username", m.Username)
}

func SetAAARequestLoginForgotPasswordSubResourceData(m []*models.AAARequestLoginForgotPassword) (d []*map[string]interface{}) {
	for _, AAARequestLoginForgotPasswordModel := range m {
		if AAARequestLoginForgotPasswordModel != nil {
			properties := make(map[string]interface{})
			properties["username"] = AAARequestLoginForgotPasswordModel.Username
			d = append(d, &properties)
		}
	}
	return
}

func AAARequestLoginForgotPasswordSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"username": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetAAARequestLoginForgotPasswordPropertyFields() (t []string) {
	return []string{
		"username",
	}
}
