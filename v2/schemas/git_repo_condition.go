package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func GitRepoConditionModel(d *schema.ResourceData) *models.GitRepoCondition {
	error, _ := d.Get("error").(bool)
	lastUpdateTime, _ := d.Get("last_update_time").(string)
	message, _ := d.Get("message").(string)
	status, _ := d.Get("status").(string)
	transitioning, _ := d.Get("transitioning").(bool)
	typeVar, _ := d.Get("type").(string)
	return &models.GitRepoCondition{
		Error:          error,
		LastUpdateTime: lastUpdateTime,
		Message:        message,
		Status:         status,
		Transitioning:  transitioning,
		Type:           typeVar,
	}
}

func GitRepoConditionModelFromMap(m map[string]interface{}) *models.GitRepoCondition {
	error := m["error"].(bool)
	lastUpdateTime := m["last_update_time"].(string)
	message := m["message"].(string)
	status := m["status"].(string)
	transitioning := m["transitioning"].(bool)
	typeVar := m["type"].(string)
	return &models.GitRepoCondition{
		Error:          error,
		LastUpdateTime: lastUpdateTime,
		Message:        message,
		Status:         status,
		Transitioning:  transitioning,
		Type:           typeVar,
	}
}

func SetGitRepoConditionResourceData(d *schema.ResourceData, m *models.GitRepoCondition) {
	d.Set("error", m.Error)
	d.Set("last_update_time", m.LastUpdateTime)
	d.Set("message", m.Message)
	d.Set("status", m.Status)
	d.Set("transitioning", m.Transitioning)
	d.Set("type", m.Type)
}

func SetGitRepoConditionSubResourceData(m []*models.GitRepoCondition) (d []*map[string]interface{}) {
	for _, GitRepoConditionModel := range m {
		if GitRepoConditionModel != nil {
			properties := make(map[string]interface{})
			properties["error"] = GitRepoConditionModel.Error
			properties["last_update_time"] = GitRepoConditionModel.LastUpdateTime
			properties["message"] = GitRepoConditionModel.Message
			properties["status"] = GitRepoConditionModel.Status
			properties["transitioning"] = GitRepoConditionModel.Transitioning
			properties["type"] = GitRepoConditionModel.Type
			d = append(d, &properties)
		}
	}
	return
}

func GitRepoConditionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"error": {
			Description: `Whether the condition represents an error state`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"last_update_time": {
			Description: `Last time the condition was updated`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"message": {
			Description: `Human readable message indicating details about the condition`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"status": {
			Description: `Status of the condition (True, False, Unknown)`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"transitioning": {
			Description: `Whether the condition is in a transitioning state`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"type": {
			Description: `Type of the condition`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetGitRepoConditionPropertyFields() (t []string) {
	return []string{
		"error",
		"last_update_time",
		"message",
		"status",
		"transitioning",
		"type",
	}
}
