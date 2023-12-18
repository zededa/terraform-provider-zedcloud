package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func Token64Model(d *schema.ResourceData) *models.Token64 {
	base64, _ := d.Get("base64").(string)
	expires, _ := d.Get("expires").(string)
	return &models.Token64{
		Base64:  base64,
		Expires: expires,
	}
}

func Token64ModelFromMap(m map[string]interface{}) *models.Token64 {
	base64 := m["base64"].(string)
	expires := m["expires"].(string)
	return &models.Token64{
		Base64:  base64,
		Expires: expires,
	}
}

func SetToken64ResourceData(d *schema.ResourceData, m *models.Token64) {
	d.Set("base64", m.Base64)
	d.Set("expires", m.Expires)
}

func SetToken64SubResourceData(m []*models.Token64) (d []*map[string]interface{}) {
	for _, Token64Model := range m {
		if Token64Model != nil {
			properties := make(map[string]interface{})
			properties["base64"] = Token64Model.Base64
			properties["expires"] = Token64Model.Expires
			d = append(d, &properties)
		}
	}
	return
}

func Token64Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"base64": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"expires": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetToken64PropertyFields() (t []string) {
	return []string{
		"base64",
		"expires",
	}
}
