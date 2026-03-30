package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func SetCertDistinguishedNameSubResourceData(m []*models.CertDistinguishedName) (d []*map[string]interface{}) {
	for _, v := range m {
		if v != nil {
			properties := make(map[string]interface{})
			properties["common_name"] = v.CommonName
			properties["serial_number"] = v.SerialNumber
			properties["organization"] = v.Organization
			properties["organizational_unit"] = v.OrganizationalUnit
			properties["country"] = v.Country
			properties["state_or_province"] = v.StateOrProvince
			properties["locality"] = v.Locality
			d = append(d, &properties)
		}
	}
	return
}

func CertDistinguishedNameSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"common_name": {
			Description: `Common Name (CN)`,
			Type:        schema.TypeString,
			Computed:    true,
		},
		"serial_number": {
			Description: `Serial Number`,
			Type:        schema.TypeString,
			Computed:    true,
		},
		"organization": {
			Description: `Organization (O)`,
			Type:        schema.TypeList,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Computed:    true,
		},
		"organizational_unit": {
			Description: `Organizational Unit (OU)`,
			Type:        schema.TypeList,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Computed:    true,
		},
		"country": {
			Description: `Country (C)`,
			Type:        schema.TypeList,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Computed:    true,
		},
		"state_or_province": {
			Description: `State or Province (ST)`,
			Type:        schema.TypeList,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Computed:    true,
		},
		"locality": {
			Description: `Locality (L)`,
			Type:        schema.TypeList,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Computed:    true,
		},
	}
}
