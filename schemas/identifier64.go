package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func Identifier64Model(d *schema.ResourceData) *models.Identifier64 {
	base64, _ := d.Get("base64").(string)
	return &models.Identifier64{
		Base64: base64,
	}
}

func Identifier64ModelFromMap(m map[string]interface{}) *models.Identifier64 {
	base64 := m["base64"].(string)
	return &models.Identifier64{
		Base64: base64,
	}
}

func SetIdentifier64ResourceData(d *schema.ResourceData, m *models.Identifier64) {
	d.Set("base64", m.Base64)
}

func SetIdentifier64SubResourceData(m []*models.Identifier64) (d []*map[string]interface{}) {
	for _, Identifier64Model := range m {
		if Identifier64Model != nil {
			properties := make(map[string]interface{})
			properties["base64"] = Identifier64Model.Base64
			d = append(d, &properties)
		}
	}
	return
}

func Identifier64Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"base64": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetIdentifier64PropertyFields() (t []string) {
	return []string{
		"base64",
	}
}
