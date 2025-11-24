package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func PrivateRepoStatusModel(d *schema.ResourceData) *models.PrivateRepoStatus {
	error, _ := d.Get("error").(bool)
	message, _ := d.Get("message").(string)
	name, _ := d.Get("name").(string)
	transitioning, _ := d.Get("transitioning").(bool)
	return &models.PrivateRepoStatus{
		Error:         error,
		Message:       message,
		Name:          name,
		Transitioning: transitioning,
	}
}

func PrivateRepoStatusModelFromMap(m map[string]interface{}) *models.PrivateRepoStatus {
	error := m["error"].(bool)
	message := m["message"].(string)
	name := m["name"].(string)
	transitioning := m["transitioning"].(bool)
	return &models.PrivateRepoStatus{
		Error:         error,
		Message:       message,
		Name:          name,
		Transitioning: transitioning,
	}
}

func SetPrivateRepoStatusResourceData(d *schema.ResourceData, m *models.PrivateRepoStatus) {
	d.Set("error", m.Error)
	d.Set("message", m.Message)
	d.Set("name", m.Name)
	d.Set("transitioning", m.Transitioning)
}

func SetPrivateRepoStatusSubResourceData(m []*models.PrivateRepoStatus) (d []*map[string]interface{}) {
	for _, PrivateRepoStatusModel := range m {
		if PrivateRepoStatusModel != nil {
			properties := make(map[string]interface{})
			properties["error"] = PrivateRepoStatusModel.Error
			properties["message"] = PrivateRepoStatusModel.Message
			properties["name"] = PrivateRepoStatusModel.Name
			properties["transitioning"] = PrivateRepoStatusModel.Transitioning
			d = append(d, &properties)
		}
	}
	return
}

func PrivateRepoStatusSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"error": {
			Description: `Indicates if the repository is in an error state`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"message": {
			Description: `Status message providing additional details about the repository state`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"name": {
			Description: `Name of the repository in its current state`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"transitioning": {
			Description: `Indicates if the repository is currently undergoing state transition`,
			Type:        schema.TypeBool,
			Optional:    true,
		},
	}
}

func GetPrivateRepoStatusPropertyFields() (t []string) {
	return []string{
		"error",
		"message",
		"name",
		"transitioning",
	}
}
