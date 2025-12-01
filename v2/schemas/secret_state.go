package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func SecretStateModel(d *schema.ResourceData) *models.SecretState {
	error, _ := d.Get("error").(bool)
	message, _ := d.Get("message").(string)
	name, _ := d.Get("name").(string)
	transitioning, _ := d.Get("transitioning").(bool)
	return &models.SecretState{
		Error:         error,
		Message:       message,
		Name:          name,
		Transitioning: transitioning,
	}
}

func SecretStateModelFromMap(m map[string]interface{}) *models.SecretState {
	error := m["error"].(bool)
	message := m["message"].(string)
	name := m["name"].(string)
	transitioning := m["transitioning"].(bool)
	return &models.SecretState{
		Error:         error,
		Message:       message,
		Name:          name,
		Transitioning: transitioning,
	}
}

func SetSecretStateResourceData(d *schema.ResourceData, m *models.SecretState) {
	d.Set("error", m.Error)
	d.Set("message", m.Message)
	d.Set("name", m.Name)
	d.Set("transitioning", m.Transitioning)
}

func SetSecretStateSubResourceData(m []*models.SecretState) (d []*map[string]interface{}) {
	for _, SecretStateModel := range m {
		if SecretStateModel != nil {
			properties := make(map[string]interface{})
			properties["error"] = SecretStateModel.Error
			properties["message"] = SecretStateModel.Message
			properties["name"] = SecretStateModel.Name
			properties["transitioning"] = SecretStateModel.Transitioning
			d = append(d, &properties)
		}
	}
	return
}

func SecretStateSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"error": {
			Description: `Indicates if the secret is in an error state`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"message": {
			Description: `Status message providing additional details about the secret state`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"name": {
			Description: `Name of the secret in its current state`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"transitioning": {
			Description: `Indicates if the secret is currently undergoing state transition`,
			Type:        schema.TypeBool,
			Optional:    true,
		},
	}
}

func GetSecretStatePropertyFields() (t []string) {
	return []string{
		"error",
		"message",
		"name",
		"transitioning",
	}
}
