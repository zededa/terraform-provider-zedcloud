package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AAAFailureResponseCredentialChangeModel(d *schema.ResourceData) *models.AAAFailureResponseCredentialChange {
	var cause *models.AAAFailureResponseCredentialChangeCause // AAAFailureResponseCredentialChangeCause
	causeInterface, causeIsSet := d.GetOk("cause")
	if causeIsSet {
		causeModel := causeInterface.(string)
		cause = models.NewAAAFailureResponseCredentialChangeCause(models.AAAFailureResponseCredentialChangeCause(causeModel))
	}
	return &models.AAAFailureResponseCredentialChange{
		Cause: cause,
	}
}

func AAAFailureResponseCredentialChangeModelFromMap(m map[string]interface{}) *models.AAAFailureResponseCredentialChange {
	var cause *models.AAAFailureResponseCredentialChangeCause // AAAFailureResponseCredentialChangeCause
	causeInterface, causeIsSet := m["cause"]
	if causeIsSet {
		causeModel := causeInterface.(string)
		cause = models.NewAAAFailureResponseCredentialChangeCause(models.AAAFailureResponseCredentialChangeCause(causeModel))
	}
	return &models.AAAFailureResponseCredentialChange{
		Cause: cause,
	}
}

func SetAAAFailureResponseCredentialChangeResourceData(d *schema.ResourceData, m *models.AAAFailureResponseCredentialChange) {
	d.Set("cause", m.Cause)
}

func SetAAAFailureResponseCredentialChangeSubResourceData(m []*models.AAAFailureResponseCredentialChange) (d []*map[string]interface{}) {
	for _, AAAFailureResponseCredentialChangeModel := range m {
		if AAAFailureResponseCredentialChangeModel != nil {
			properties := make(map[string]interface{})
			properties["cause"] = AAAFailureResponseCredentialChangeModel.Cause
			d = append(d, &properties)
		}
	}
	return
}

func AAAFailureResponseCredentialChangeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cause": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetAAAFailureResponseCredentialChangePropertyFields() (t []string) {
	return []string{
		"cause",
	}
}
