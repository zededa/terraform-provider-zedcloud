package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func StaticDNSListSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"addrs": {
			Description: `Addresses`,
			Type:        schema.TypeList, //GoType: []string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"hostname": {
			Description: `Host name`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}
