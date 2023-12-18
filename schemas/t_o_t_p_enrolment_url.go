package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func TOTPEnrolmentURLModel(d *schema.ResourceData) *models.TOTPEnrolmentURL {
	secret, _ := d.Get("secret").(string)
	url, _ := d.Get("url").(string)
	return &models.TOTPEnrolmentURL{
		Secret: secret,
		URL:    url,
	}
}

func TOTPEnrolmentURLModelFromMap(m map[string]interface{}) *models.TOTPEnrolmentURL {
	secret := m["secret"].(string)
	url := m["url"].(string)
	return &models.TOTPEnrolmentURL{
		Secret: secret,
		URL:    url,
	}
}

func SetTOTPEnrolmentURLResourceData(d *schema.ResourceData, m *models.TOTPEnrolmentURL) {
	d.Set("secret", m.Secret)
	d.Set("url", m.URL)
}

func SetTOTPEnrolmentURLSubResourceData(m []*models.TOTPEnrolmentURL) (d []*map[string]interface{}) {
	for _, TOTPEnrolmentURLModel := range m {
		if TOTPEnrolmentURLModel != nil {
			properties := make(map[string]interface{})
			properties["secret"] = TOTPEnrolmentURLModel.Secret
			properties["url"] = TOTPEnrolmentURLModel.URL
			d = append(d, &properties)
		}
	}
	return
}

func TOTPEnrolmentURLSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"secret": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"url": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetTOTPEnrolmentURLPropertyFields() (t []string) {
	return []string{
		"secret",
		"url",
	}
}
