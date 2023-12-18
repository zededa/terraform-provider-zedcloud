package schemas

import (
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/zededa/terraform-provider/models"
)

func SessionDetailsModel(d *schema.ResourceData) *models.SessionDetails {
	clientIP, _ := d.Get("client_ip").(string)
	expiresAt, _ := d.Get("expires_at").(strfmt.DateTime)
	userAgent, _ := d.Get("user_agent").(string)
	userName, _ := d.Get("user_name").(string)
	return &models.SessionDetails{
		ClientIP:  clientIP,
		ExpiresAt: expiresAt,
		UserAgent: userAgent,
		UserName:  userName,
	}
}

func SessionDetailsModelFromMap(m map[string]interface{}) *models.SessionDetails {
	clientIP := m["client_ip"].(string)
	expiresAt := m["expires_at"].(strfmt.DateTime)
	userAgent := m["user_agent"].(string)
	userName := m["user_name"].(string)
	return &models.SessionDetails{
		ClientIP:  clientIP,
		ExpiresAt: expiresAt,
		UserAgent: userAgent,
		UserName:  userName,
	}
}

func SetSessionDetailsResourceData(d *schema.ResourceData, m *models.SessionDetails) {
	d.Set("client_ip", m.ClientIP)
	d.Set("expires_at", m.ExpiresAt)
	d.Set("session_id", m.SessionID)
	d.Set("user_agent", m.UserAgent)
	d.Set("user_name", m.UserName)
}

func SetSessionDetailsSubResourceData(m []*models.SessionDetails) (d []*map[string]interface{}) {
	for _, SessionDetailsModel := range m {
		if SessionDetailsModel != nil {
			properties := make(map[string]interface{})
			properties["client_ip"] = SessionDetailsModel.ClientIP
			properties["expires_at"] = SessionDetailsModel.ExpiresAt.String()
			properties["session_id"] = SessionDetailsModel.SessionID
			properties["user_agent"] = SessionDetailsModel.UserAgent
			properties["user_name"] = SessionDetailsModel.UserName
			d = append(d, &properties)
		}
	}
	return
}

func SessionDetailsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"client_ip": {
			Description: `IP address of the logged in user`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"expires_at": {
			Description:  `session expiry time`,
			Type:         schema.TypeString,
			ValidateFunc: validation.IsRFC3339Time,
			Optional:     true,
		},

		"session_id": {
			Description: `DEPRECATED. sessionId is obfuscated for security reasons`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"user_agent": {
			Description: `user agent details`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"user_name": {
			Description: `Logged in user name`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetSessionDetailsPropertyFields() (t []string) {
	return []string{
		"client_ip",
		"expires_at",
		"user_agent",
		"user_name",
	}
}
