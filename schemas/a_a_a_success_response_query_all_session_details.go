package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AAASuccessResponseQueryAllSessionDetailsModel(d *schema.ResourceData) *models.AAASuccessResponseQueryAllSessionDetails {
	var cause *models.AAASuccessResponseQueryAllSessionDetailsCause // AAASuccessResponseQueryAllSessionDetailsCause
	causeInterface, causeIsSet := d.GetOk("cause")
	if causeIsSet {
		causeModel := causeInterface.(string)
		cause = models.NewAAASuccessResponseQueryAllSessionDetailsCause(models.AAASuccessResponseQueryAllSessionDetailsCause(causeModel))
	}
	var sessionDetails []*models.SessionDetails // []*SessionDetails
	sessionDetailsInterface, sessionDetailsIsSet := d.GetOk("session_details")
	if sessionDetailsIsSet {
		var items []interface{}
		if listItems, isList := sessionDetailsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = sessionDetailsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := SessionDetailsModelFromMap(v.(map[string]interface{}))
			sessionDetails = append(sessionDetails, m)
		}
	}
	return &models.AAASuccessResponseQueryAllSessionDetails{
		Cause:          cause,
		SessionDetails: sessionDetails,
	}
}

func AAASuccessResponseQueryAllSessionDetailsModelFromMap(m map[string]interface{}) *models.AAASuccessResponseQueryAllSessionDetails {
	var cause *models.AAASuccessResponseQueryAllSessionDetailsCause // AAASuccessResponseQueryAllSessionDetailsCause
	causeInterface, causeIsSet := m["cause"]
	if causeIsSet {
		causeModel := causeInterface.(string)
		cause = models.NewAAASuccessResponseQueryAllSessionDetailsCause(models.AAASuccessResponseQueryAllSessionDetailsCause(causeModel))
	}
	var sessionDetails []*models.SessionDetails // []*SessionDetails
	sessionDetailsInterface, sessionDetailsIsSet := m["session_details"]
	if sessionDetailsIsSet {
		var items []interface{}
		if listItems, isList := sessionDetailsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = sessionDetailsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := SessionDetailsModelFromMap(v.(map[string]interface{}))
			sessionDetails = append(sessionDetails, m)
		}
	}
	return &models.AAASuccessResponseQueryAllSessionDetails{
		Cause:          cause,
		SessionDetails: sessionDetails,
	}
}

func SetAAASuccessResponseQueryAllSessionDetailsResourceData(d *schema.ResourceData, m *models.AAASuccessResponseQueryAllSessionDetails) {
	d.Set("cause", m.Cause)
	d.Set("session_details", SetSessionDetailsSubResourceData(m.SessionDetails))
}

func SetAAASuccessResponseQueryAllSessionDetailsSubResourceData(m []*models.AAASuccessResponseQueryAllSessionDetails) (d []*map[string]interface{}) {
	for _, AAASuccessResponseQueryAllSessionDetailsModel := range m {
		if AAASuccessResponseQueryAllSessionDetailsModel != nil {
			properties := make(map[string]interface{})
			properties["cause"] = AAASuccessResponseQueryAllSessionDetailsModel.Cause
			properties["session_details"] = SetSessionDetailsSubResourceData(AAASuccessResponseQueryAllSessionDetailsModel.SessionDetails)
			d = append(d, &properties)
		}
	}
	return
}

func AAASuccessResponseQueryAllSessionDetailsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cause": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"session_details": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []*SessionDetails
			Elem: &schema.Resource{
				Schema: SessionDetailsSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
	}
}

func GetAAASuccessResponseQueryAllSessionDetailsPropertyFields() (t []string) {
	return []string{
		"cause",
		"session_details",
	}
}
