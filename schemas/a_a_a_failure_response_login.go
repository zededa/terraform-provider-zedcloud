package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AAAFailureResponseLoginModel(d *schema.ResourceData) *models.AAAFailureResponseLogin {
	var cause *models.AAAFailureResponseLoginCause // AAAFailureResponseLoginCause
	causeInterface, causeIsSet := d.GetOk("cause")
	if causeIsSet {
		causeModel := causeInterface.(string)
		cause = models.NewAAAFailureResponseLoginCause(models.AAAFailureResponseLoginCause(causeModel))
	}
	error, _ := d.Get("error").(string)
	noOfLoginAttemptsLeftInt, _ := d.Get("no_of_login_attempts_left").(int)
	noOfLoginAttemptsLeft := int64(noOfLoginAttemptsLeftInt)
	var tempSuccessResponse *models.AAASuccessResponseLogin // AAASuccessResponseLogin
	tempSuccessResponseInterface, tempSuccessResponseIsSet := d.GetOk("temp_success_response")
	if tempSuccessResponseIsSet && tempSuccessResponseInterface != nil {
		tempSuccessResponseMap := tempSuccessResponseInterface.([]interface{})
		if len(tempSuccessResponseMap) > 0 {
			tempSuccessResponse = AAASuccessResponseLoginModelFromMap(tempSuccessResponseMap[0].(map[string]interface{}))
		}
	}
	var tempToken *models.Token64 // Token64
	tempTokenInterface, tempTokenIsSet := d.GetOk("temp_token")
	if tempTokenIsSet && tempTokenInterface != nil {
		tempTokenMap := tempTokenInterface.([]interface{})
		if len(tempTokenMap) > 0 {
			tempToken = Token64ModelFromMap(tempTokenMap[0].(map[string]interface{}))
		}
	}
	return &models.AAAFailureResponseLogin{
		Cause:                 cause,
		Error:                 error,
		NoOfLoginAttemptsLeft: noOfLoginAttemptsLeft,
		TempSuccessResponse:   tempSuccessResponse,
		TempToken:             tempToken,
	}
}

func AAAFailureResponseLoginModelFromMap(m map[string]interface{}) *models.AAAFailureResponseLogin {
	var cause *models.AAAFailureResponseLoginCause // AAAFailureResponseLoginCause
	causeInterface, causeIsSet := m["cause"]
	if causeIsSet {
		causeModel := causeInterface.(string)
		cause = models.NewAAAFailureResponseLoginCause(models.AAAFailureResponseLoginCause(causeModel))
	}
	error := m["error"].(string)
	noOfLoginAttemptsLeft := int64(m["no_of_login_attempts_left"].(int)) // int64
	var tempSuccessResponse *models.AAASuccessResponseLogin              // AAASuccessResponseLogin
	tempSuccessResponseInterface, tempSuccessResponseIsSet := m["temp_success_response"]
	if tempSuccessResponseIsSet && tempSuccessResponseInterface != nil {
		tempSuccessResponseMap := tempSuccessResponseInterface.([]interface{})
		if len(tempSuccessResponseMap) > 0 {
			tempSuccessResponse = AAASuccessResponseLoginModelFromMap(tempSuccessResponseMap[0].(map[string]interface{}))
		}
	}
	//
	var tempToken *models.Token64 // Token64
	tempTokenInterface, tempTokenIsSet := m["temp_token"]
	if tempTokenIsSet && tempTokenInterface != nil {
		tempTokenMap := tempTokenInterface.([]interface{})
		if len(tempTokenMap) > 0 {
			tempToken = Token64ModelFromMap(tempTokenMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.AAAFailureResponseLogin{
		Cause:                 cause,
		Error:                 error,
		NoOfLoginAttemptsLeft: noOfLoginAttemptsLeft,
		TempSuccessResponse:   tempSuccessResponse,
		TempToken:             tempToken,
	}
}

func SetAAAFailureResponseLoginResourceData(d *schema.ResourceData, m *models.AAAFailureResponseLogin) {
	d.Set("cause", m.Cause)
	d.Set("error", m.Error)
	d.Set("no_of_login_attempts_left", m.NoOfLoginAttemptsLeft)
	d.Set("temp_success_response", SetAAASuccessResponseLoginSubResourceData([]*models.AAASuccessResponseLogin{m.TempSuccessResponse}))
	d.Set("temp_token", SetToken64SubResourceData([]*models.Token64{m.TempToken}))
}

func SetAAAFailureResponseLoginSubResourceData(m []*models.AAAFailureResponseLogin) (d []*map[string]interface{}) {
	for _, AAAFailureResponseLoginModel := range m {
		if AAAFailureResponseLoginModel != nil {
			properties := make(map[string]interface{})
			properties["cause"] = AAAFailureResponseLoginModel.Cause
			properties["error"] = AAAFailureResponseLoginModel.Error
			properties["no_of_login_attempts_left"] = AAAFailureResponseLoginModel.NoOfLoginAttemptsLeft
			properties["temp_success_response"] = SetAAASuccessResponseLoginSubResourceData([]*models.AAASuccessResponseLogin{AAAFailureResponseLoginModel.TempSuccessResponse})
			properties["temp_token"] = SetToken64SubResourceData([]*models.Token64{AAAFailureResponseLoginModel.TempToken})
			d = append(d, &properties)
		}
	}
	return
}

func AAAFailureResponseLoginSchema() map[string]*schema.Schema {
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

		"no_of_login_attempts_left": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"temp_success_response": {
			Description: `Sessions depend heavily on AAASuccessResponseLogin. In case of password expired,
we need temporary token. We can not generate a temporary token for password reset with
AAAFailureResponseLogin itself. Therefore, adding this tempSuccessResponse, to be used to create new session.`,
			Type: schema.TypeList, //GoType: AAASuccessResponseLogin
			Elem: &schema.Resource{
				Schema: AAASuccessResponseLoginSchema(),
			},
			Optional: true,
		},

		"temp_token": {
			Description: ``,
			Type:        schema.TypeList, //GoType: Token64
			Elem: &schema.Resource{
				Schema: Token64Schema(),
			},
			Optional: true,
		},
	}
}

func GetAAAFailureResponseLoginPropertyFields() (t []string) {
	return []string{
		"cause",
		"error",
		"no_of_login_attempts_left",
		"temp_success_response",
		"temp_token",
	}
}
