package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AAARedirectResponseModel(d *schema.ResourceData) *models.AAARedirectResponse {
	codeInt, _ := d.Get("code").(int)
	code := int64(codeInt)
	redirectURL, _ := d.Get("redirect_url").(string)
	return &models.AAARedirectResponse{
		Code:        code,
		RedirectURL: redirectURL,
	}
}

func AAARedirectResponseModelFromMap(m map[string]interface{}) *models.AAARedirectResponse {
	code := int64(m["code"].(int)) // int64
	redirectURL := m["redirect_url"].(string)
	return &models.AAARedirectResponse{
		Code:        code,
		RedirectURL: redirectURL,
	}
}

func SetAAARedirectResponseResourceData(d *schema.ResourceData, m *models.AAARedirectResponse) {
	d.Set("code", m.Code)
	d.Set("redirect_url", m.RedirectURL)
}

func SetAAARedirectResponseSubResourceData(m []*models.AAARedirectResponse) (d []*map[string]interface{}) {
	for _, AAARedirectResponseModel := range m {
		if AAARedirectResponseModel != nil {
			properties := make(map[string]interface{})
			properties["code"] = AAARedirectResponseModel.Code
			properties["redirect_url"] = AAARedirectResponseModel.RedirectURL
			d = append(d, &properties)
		}
	}
	return
}

func AAARedirectResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"code": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"redirect_url": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetAAARedirectResponsePropertyFields() (t []string) {
	return []string{
		"code",
		"redirect_url",
	}
}
