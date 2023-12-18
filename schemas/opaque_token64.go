package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func OpaqueToken64Model(d *schema.ResourceData) *models.OpaqueToken64 {
	base64, _ := d.Get("base64").(string)
	return &models.OpaqueToken64{
		Base64: base64,
	}
}

func OpaqueToken64ModelFromMap(m map[string]interface{}) *models.OpaqueToken64 {
	base64 := m["base64"].(string)
	return &models.OpaqueToken64{
		Base64: base64,
	}
}

func SetOpaqueToken64ResourceData(d *schema.ResourceData, m *models.OpaqueToken64) {
	d.Set("base64", m.Base64)
}

func SetOpaqueToken64SubResourceData(m []*models.OpaqueToken64) (d []*map[string]interface{}) {
	for _, OpaqueToken64Model := range m {
		if OpaqueToken64Model != nil {
			properties := make(map[string]interface{})
			properties["base64"] = OpaqueToken64Model.Base64
			d = append(d, &properties)
		}
	}
	return
}

func OpaqueToken64Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"base64": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetOpaqueToken64PropertyFields() (t []string) {
	return []string{
		"base64",
	}
}
