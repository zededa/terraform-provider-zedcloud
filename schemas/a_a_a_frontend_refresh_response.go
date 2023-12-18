package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AAAFrontendRefreshResponseModel(d *schema.ResourceData) *models.AAAFrontendRefreshResponse {
	var cause *models.AAAFrontendRefreshResponseCause // AAAFrontendRefreshResponseCause
	causeInterface, causeIsSet := d.GetOk("cause")
	if causeIsSet {
		causeModel := causeInterface.(string)
		cause = models.NewAAAFrontendRefreshResponseCause(models.AAAFrontendRefreshResponseCause(causeModel))
	}
	var token *models.Token64 // Token64
	tokenInterface, tokenIsSet := d.GetOk("token")
	if tokenIsSet && tokenInterface != nil {
		tokenMap := tokenInterface.([]interface{})
		if len(tokenMap) > 0 {
			token = Token64ModelFromMap(tokenMap[0].(map[string]interface{}))
		}
	}
	userID, _ := d.Get("user_id").(string)
	return &models.AAAFrontendRefreshResponse{
		Cause:  cause,
		Token:  token,
		UserID: userID,
	}
}

func AAAFrontendRefreshResponseModelFromMap(m map[string]interface{}) *models.AAAFrontendRefreshResponse {
	var cause *models.AAAFrontendRefreshResponseCause // AAAFrontendRefreshResponseCause
	causeInterface, causeIsSet := m["cause"]
	if causeIsSet {
		causeModel := causeInterface.(string)
		cause = models.NewAAAFrontendRefreshResponseCause(models.AAAFrontendRefreshResponseCause(causeModel))
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
	userID := m["user_id"].(string)
	return &models.AAAFrontendRefreshResponse{
		Cause:  cause,
		Token:  token,
		UserID: userID,
	}
}

func SetAAAFrontendRefreshResponseResourceData(d *schema.ResourceData, m *models.AAAFrontendRefreshResponse) {
	d.Set("cause", m.Cause)
	d.Set("token", SetToken64SubResourceData([]*models.Token64{m.Token}))
	d.Set("user_id", m.UserID)
}

func SetAAAFrontendRefreshResponseSubResourceData(m []*models.AAAFrontendRefreshResponse) (d []*map[string]interface{}) {
	for _, AAAFrontendRefreshResponseModel := range m {
		if AAAFrontendRefreshResponseModel != nil {
			properties := make(map[string]interface{})
			properties["cause"] = AAAFrontendRefreshResponseModel.Cause
			properties["token"] = SetToken64SubResourceData([]*models.Token64{AAAFrontendRefreshResponseModel.Token})
			properties["user_id"] = AAAFrontendRefreshResponseModel.UserID
			d = append(d, &properties)
		}
	}
	return
}

func AAAFrontendRefreshResponseSchema() map[string]*schema.Schema {
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

		"user_id": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetAAAFrontendRefreshResponsePropertyFields() (t []string) {
	return []string{
		"cause",
		"token",
		"user_id",
	}
}
