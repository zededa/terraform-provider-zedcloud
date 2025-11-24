package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/go-openapi/strfmt"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func GitRepoEventsModel(d *schema.ResourceData) *models.GitRepoEvents {
	lastUpdateTime, _ := d.Get("last_update_time").(strfmt.DateTime)
	message, _ := d.Get("message").(string)
	reason, _ := d.Get("reason").(string)
	typeVar, _ := d.Get("type").(string)
	return &models.GitRepoEvents{
		LastUpdateTime: lastUpdateTime,
		Message:        message,
		Reason:         reason,
		Type:           typeVar,
	}
}

func GitRepoEventsModelFromMap(m map[string]interface{}) *models.GitRepoEvents {
	lastUpdateTime := m["last_update_time"].(strfmt.DateTime)
	message := m["message"].(string)
	reason := m["reason"].(string)
	typeVar := m["type"].(string)
	return &models.GitRepoEvents{
		LastUpdateTime: lastUpdateTime,
		Message:        message,
		Reason:         reason,
		Type:           typeVar,
	}
}

func SetGitRepoEventsResourceData(d *schema.ResourceData, m *models.GitRepoEvents) {
	d.Set("last_update_time", m.LastUpdateTime)
	d.Set("message", m.Message)
	d.Set("reason", m.Reason)
	d.Set("type", m.Type)
}

func SetGitRepoEventsSubResourceData(m []*models.GitRepoEvents) (d []*map[string]interface{}) {
	for _, GitRepoEventsModel := range m {
		if GitRepoEventsModel != nil {
			properties := make(map[string]interface{})
			properties["last_update_time"] = GitRepoEventsModel.LastUpdateTime.String()
			properties["message"] = GitRepoEventsModel.Message
			properties["reason"] = GitRepoEventsModel.Reason
			properties["type"] = GitRepoEventsModel.Type
			d = append(d, &properties)
		}
	}
	return
}

func GitRepoEventsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"last_update_time": {
			Description:  `Last time the event was updated`,
			Type:         schema.TypeString,
			ValidateFunc: validation.IsRFC3339Time,
			Optional:     true,
		},

		"message": {
			Description: `Detailed message about the event`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"reason": {
			Description: `Reason for the event`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"type": {
			Description: `Type of the event`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetGitRepoEventsPropertyFields() (t []string) {
	return []string{
		"last_update_time",
		"message",
		"reason",
		"type",
	}
}
