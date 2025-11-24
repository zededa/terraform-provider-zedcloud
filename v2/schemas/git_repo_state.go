package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func GitRepoStateModel(d *schema.ResourceData) *models.GitRepoState {
	error, _ := d.Get("error").(bool)
	message, _ := d.Get("message").(string)
	name, _ := d.Get("name").(string)
	transitioning, _ := d.Get("transitioning").(bool)
	return &models.GitRepoState{
		Error:         error,
		Message:       message,
		Name:          name,
		Transitioning: transitioning,
	}
}

func GitRepoStateModelFromMap(m map[string]interface{}) *models.GitRepoState {
	error := m["error"].(bool)
	message := m["message"].(string)
	name := m["name"].(string)
	transitioning := m["transitioning"].(bool)
	return &models.GitRepoState{
		Error:         error,
		Message:       message,
		Name:          name,
		Transitioning: transitioning,
	}
}

func SetGitRepoStateResourceData(d *schema.ResourceData, m *models.GitRepoState) {
	d.Set("error", m.Error)
	d.Set("message", m.Message)
	d.Set("name", m.Name)
	d.Set("transitioning", m.Transitioning)
}

func SetGitRepoStateSubResourceData(m []*models.GitRepoState) (d []*map[string]interface{}) {
	for _, GitRepoStateModel := range m {
		if GitRepoStateModel != nil {
			properties := make(map[string]interface{})
			properties["error"] = GitRepoStateModel.Error
			properties["message"] = GitRepoStateModel.Message
			properties["name"] = GitRepoStateModel.Name
			properties["transitioning"] = GitRepoStateModel.Transitioning
			d = append(d, &properties)
		}
	}
	return
}

func GitRepoStateSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"error": {
			Description: `Indicates if the GitRepo is in an error state`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"message": {
			Description: `Status message providing additional details about the GitRepo state`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"name": {
			Description: `current state name`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"transitioning": {
			Description: `Indicates if the GitRepo is currently undergoing state transition`,
			Type:        schema.TypeBool,
			Optional:    true,
		},
	}
}

func GetGitRepoStatePropertyFields() (t []string) {
	return []string{
		"error",
		"message",
		"name",
		"transitioning",
	}
}
