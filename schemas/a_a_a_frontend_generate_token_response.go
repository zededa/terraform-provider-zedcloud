package schemas

import (
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/zededa/terraform-provider/models"
)

func AAAFrontendGenerateTokenResponseModel(d *schema.ResourceData) *models.AAAFrontendGenerateTokenResponse {
	var cause *models.AAAFrontendGenerateTokenResponseCause // AAAFrontendGenerateTokenResponseCause
	causeInterface, causeIsSet := d.GetOk("cause")
	if causeIsSet {
		causeModel := causeInterface.(string)
		cause = models.NewAAAFrontendGenerateTokenResponseCause(models.AAAFrontendGenerateTokenResponseCause(causeModel))
	}
	clientIP, _ := d.Get("client_ip").(string)
	error, _ := d.Get("error").(string)
	expiresAt, _ := d.Get("expires_at").(strfmt.DateTime)
	token, _ := d.Get("token").(string)
	userAgent, _ := d.Get("user_agent").(string)
	userName, _ := d.Get("user_name").(string)
	return &models.AAAFrontendGenerateTokenResponse{
		Cause:     cause,
		ClientIP:  clientIP,
		Error:     error,
		ExpiresAt: expiresAt,
		Token:     token,
		UserAgent: userAgent,
		UserName:  userName,
	}
}

func AAAFrontendGenerateTokenResponseModelFromMap(m map[string]interface{}) *models.AAAFrontendGenerateTokenResponse {
	var cause *models.AAAFrontendGenerateTokenResponseCause // AAAFrontendGenerateTokenResponseCause
	causeInterface, causeIsSet := m["cause"]
	if causeIsSet {
		causeModel := causeInterface.(string)
		cause = models.NewAAAFrontendGenerateTokenResponseCause(models.AAAFrontendGenerateTokenResponseCause(causeModel))
	}
	clientIP := m["client_ip"].(string)
	error := m["error"].(string)
	expiresAt := m["expires_at"].(strfmt.DateTime)
	token := m["token"].(string)
	userAgent := m["user_agent"].(string)
	userName := m["user_name"].(string)
	return &models.AAAFrontendGenerateTokenResponse{
		Cause:     cause,
		ClientIP:  clientIP,
		Error:     error,
		ExpiresAt: expiresAt,
		Token:     token,
		UserAgent: userAgent,
		UserName:  userName,
	}
}

func SetAAAFrontendGenerateTokenResponseResourceData(d *schema.ResourceData, m *models.AAAFrontendGenerateTokenResponse) {
	d.Set("cause", m.Cause)
	d.Set("client_ip", m.ClientIP)
	d.Set("error", m.Error)
	d.Set("expires_at", m.ExpiresAt)
	d.Set("token", m.Token)
	d.Set("user_agent", m.UserAgent)
	d.Set("user_name", m.UserName)
}

func SetAAAFrontendGenerateTokenResponseSubResourceData(m []*models.AAAFrontendGenerateTokenResponse) (d []*map[string]interface{}) {
	for _, AAAFrontendGenerateTokenResponseModel := range m {
		if AAAFrontendGenerateTokenResponseModel != nil {
			properties := make(map[string]interface{})
			properties["cause"] = AAAFrontendGenerateTokenResponseModel.Cause
			properties["client_ip"] = AAAFrontendGenerateTokenResponseModel.ClientIP
			properties["error"] = AAAFrontendGenerateTokenResponseModel.Error
			properties["expires_at"] = AAAFrontendGenerateTokenResponseModel.ExpiresAt.String()
			properties["token"] = AAAFrontendGenerateTokenResponseModel.Token
			properties["user_agent"] = AAAFrontendGenerateTokenResponseModel.UserAgent
			properties["user_name"] = AAAFrontendGenerateTokenResponseModel.UserName
			d = append(d, &properties)
		}
	}
	return
}

func AAAFrontendGenerateTokenResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cause": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"client_ip": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"error": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"expires_at": {
			Description:  ``,
			Type:         schema.TypeString,
			ValidateFunc: validation.IsRFC3339Time,
			Optional:     true,
		},

		"token": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"user_agent": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"user_name": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetAAAFrontendGenerateTokenResponsePropertyFields() (t []string) {
	return []string{
		"cause",
		"client_ip",
		"error",
		"expires_at",
		"token",
		"user_agent",
		"user_name",
	}
}
