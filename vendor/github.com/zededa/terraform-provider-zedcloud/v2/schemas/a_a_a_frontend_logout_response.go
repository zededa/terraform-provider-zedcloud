package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func AAAFrontendLogoutResponseModel(d *schema.ResourceData) *models.AAAFrontendLogoutResponse {
	var cause *models.AAAFrontendLogoutResponseCause // AAAFrontendLogoutResponseCause
	causeInterface, causeIsSet := d.GetOk("cause")
	if causeIsSet {
		causeModel := causeInterface.(string)
		cause = models.NewAAAFrontendLogoutResponseCause(models.AAAFrontendLogoutResponseCause(causeModel))
	}
	var token *models.Token64 // Token64
	tokenInterface, tokenIsSet := d.GetOk("token")
	if tokenIsSet && tokenInterface != nil {
		tokenMap := tokenInterface.([]interface{})
		if len(tokenMap) > 0 {
			token = Token64ModelFromMap(tokenMap[0].(map[string]interface{}))
		}
	}
	return &models.AAAFrontendLogoutResponse{
		Cause: cause,
		Token: token,
	}
}

func AAAFrontendLogoutResponseModelFromMap(m map[string]interface{}) *models.AAAFrontendLogoutResponse {
	var cause *models.AAAFrontendLogoutResponseCause // AAAFrontendLogoutResponseCause
	causeInterface, causeIsSet := m["cause"]
	if causeIsSet {
		causeModel := causeInterface.(string)
		cause = models.NewAAAFrontendLogoutResponseCause(models.AAAFrontendLogoutResponseCause(causeModel))
	}
	var token *models.Token64 // Token64
	tokenInterface, tokenIsSet := m["token"]
	if tokenIsSet && tokenInterface != nil {
		tokenMap := tokenInterface.([]interface{})
		if len(tokenMap) > 0 {
			token = Token64ModelFromMap(tokenMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.AAAFrontendLogoutResponse{
		Cause: cause,
		Token: token,
	}
}

func SetAAAFrontendLogoutResponseResourceData(d *schema.ResourceData, m *models.AAAFrontendLogoutResponse) {
	d.Set("cause", m.Cause)
	d.Set("token", SetToken64SubResourceData([]*models.Token64{m.Token}))
}

func SetAAAFrontendLogoutResponseSubResourceData(m []*models.AAAFrontendLogoutResponse) (d []*map[string]interface{}) {
	for _, AAAFrontendLogoutResponseModel := range m {
		if AAAFrontendLogoutResponseModel != nil {
			properties := make(map[string]interface{})
			properties["cause"] = AAAFrontendLogoutResponseModel.Cause
			properties["token"] = SetToken64SubResourceData([]*models.Token64{AAAFrontendLogoutResponseModel.Token})
			d = append(d, &properties)
		}
	}
	return
}

func AAAFrontendLogoutResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cause": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"token": {
			Description: ``,
			Type:        schema.TypeList, //GoType: Token64
			Elem: &schema.Resource{
				Schema: Token64Schema(),
			},
			Optional: true,
		},
	}
}

func GetAAAFrontendLogoutResponsePropertyFields() (t []string) {
	return []string{
		"cause",
		"token",
	}
}
