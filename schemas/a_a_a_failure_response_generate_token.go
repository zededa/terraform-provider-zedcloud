package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AAAFailureResponseGenerateTokenModel(d *schema.ResourceData) *models.AAAFailureResponseGenerateToken {
	var cause *models.AAAFailureResponseGenerateTokenCause // AAAFailureResponseGenerateTokenCause
	causeInterface, causeIsSet := d.GetOk("cause")
	if causeIsSet {
		causeModel := causeInterface.(string)
		cause = models.NewAAAFailureResponseGenerateTokenCause(models.AAAFailureResponseGenerateTokenCause(causeModel))
	}
	error, _ := d.Get("error").(string)
	return &models.AAAFailureResponseGenerateToken{
		Cause: cause,
		Error: error,
	}
}

func AAAFailureResponseGenerateTokenModelFromMap(m map[string]interface{}) *models.AAAFailureResponseGenerateToken {
	var cause *models.AAAFailureResponseGenerateTokenCause // AAAFailureResponseGenerateTokenCause
	causeInterface, causeIsSet := m["cause"]
	if causeIsSet {
		causeModel := causeInterface.(string)
		cause = models.NewAAAFailureResponseGenerateTokenCause(models.AAAFailureResponseGenerateTokenCause(causeModel))
	}
	error := m["error"].(string)
	return &models.AAAFailureResponseGenerateToken{
		Cause: cause,
		Error: error,
	}
}

func SetAAAFailureResponseGenerateTokenResourceData(d *schema.ResourceData, m *models.AAAFailureResponseGenerateToken) {
	d.Set("cause", m.Cause)
	d.Set("error", m.Error)
}

func SetAAAFailureResponseGenerateTokenSubResourceData(m []*models.AAAFailureResponseGenerateToken) (d []*map[string]interface{}) {
	for _, AAAFailureResponseGenerateTokenModel := range m {
		if AAAFailureResponseGenerateTokenModel != nil {
			properties := make(map[string]interface{})
			properties["cause"] = AAAFailureResponseGenerateTokenModel.Cause
			properties["error"] = AAAFailureResponseGenerateTokenModel.Error
			d = append(d, &properties)
		}
	}
	return
}

func AAAFailureResponseGenerateTokenSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cause": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"error": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetAAAFailureResponseGenerateTokenPropertyFields() (t []string) {
	return []string{
		"cause",
		"error",
	}
}
