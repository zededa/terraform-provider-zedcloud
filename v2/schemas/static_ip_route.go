package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func StaticIPRouteModel(d *schema.ResourceData) *models.StaticIPRoute {
	gateway, _ := d.Get("gateway").(string)
	prefix, _ := d.Get("prefix").(string)
	return &models.StaticIPRoute{
		Gateway: gateway,
		Prefix:  prefix,
	}
}

func StaticIPRouteModelFromMap(m map[string]interface{}) *models.StaticIPRoute {
	gateway := m["gateway"].(string)
	prefix := m["prefix"].(string)
	return &models.StaticIPRoute{
		Gateway: gateway,
		Prefix:  prefix,
	}
}

func SetStaticIPRouteResourceData(d *schema.ResourceData, m *models.StaticIPRoute) {
	d.Set("gateway", m.Gateway)
	d.Set("prefix", m.Prefix)
}

func SetStaticIPRouteSubResourceData(m []*models.StaticIPRoute) (d []*map[string]interface{}) {
	for _, StaticIPRouteModel := range m {
		if StaticIPRouteModel != nil {
			properties := make(map[string]interface{})
			properties["gateway"] = StaticIPRouteModel.Gateway
			properties["prefix"] = StaticIPRouteModel.Prefix
			d = append(d, &properties)
		}
	}
	return
}

func StaticIPRouteSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"gateway": {
			Description: `Gateway IP`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"prefix": {
			Description: `IP Prefix`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetStaticIPRoutePropertyFields() (t []string) {
	return []string{
		"gateway",
		"prefix",
	}
}
