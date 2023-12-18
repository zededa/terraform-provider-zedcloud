package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AAAFailureResponseQueryAllSessionDetailsModel(d *schema.ResourceData) *models.AAAFailureResponseQueryAllSessionDetails {
	var cause *models.AAAFailureResponseQueryAllSessionDetailsCause // AAAFailureResponseQueryAllSessionDetailsCause
	causeInterface, causeIsSet := d.GetOk("cause")
	if causeIsSet {
		causeModel := causeInterface.(string)
		cause = models.NewAAAFailureResponseQueryAllSessionDetailsCause(models.AAAFailureResponseQueryAllSessionDetailsCause(causeModel))
	}
	error, _ := d.Get("error").(string)
	return &models.AAAFailureResponseQueryAllSessionDetails{
		Cause: cause,
		Error: error,
	}
}

func AAAFailureResponseQueryAllSessionDetailsModelFromMap(m map[string]interface{}) *models.AAAFailureResponseQueryAllSessionDetails {
	var cause *models.AAAFailureResponseQueryAllSessionDetailsCause // AAAFailureResponseQueryAllSessionDetailsCause
	causeInterface, causeIsSet := m["cause"]
	if causeIsSet {
		causeModel := causeInterface.(string)
		cause = models.NewAAAFailureResponseQueryAllSessionDetailsCause(models.AAAFailureResponseQueryAllSessionDetailsCause(causeModel))
	}
	error := m["error"].(string)
	return &models.AAAFailureResponseQueryAllSessionDetails{
		Cause: cause,
		Error: error,
	}
}

func SetAAAFailureResponseQueryAllSessionDetailsResourceData(d *schema.ResourceData, m *models.AAAFailureResponseQueryAllSessionDetails) {
	d.Set("cause", m.Cause)
	d.Set("error", m.Error)
}

func SetAAAFailureResponseQueryAllSessionDetailsSubResourceData(m []*models.AAAFailureResponseQueryAllSessionDetails) (d []*map[string]interface{}) {
	for _, AAAFailureResponseQueryAllSessionDetailsModel := range m {
		if AAAFailureResponseQueryAllSessionDetailsModel != nil {
			properties := make(map[string]interface{})
			properties["cause"] = AAAFailureResponseQueryAllSessionDetailsModel.Cause
			properties["error"] = AAAFailureResponseQueryAllSessionDetailsModel.Error
			d = append(d, &properties)
		}
	}
	return
}

func AAAFailureResponseQueryAllSessionDetailsSchema() map[string]*schema.Schema {
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

func GetAAAFailureResponseQueryAllSessionDetailsPropertyFields() (t []string) {
	return []string{
		"cause",
		"error",
	}
}
