package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func RSAModel(d *schema.ResourceData) *models.RSA {
	rsaBits, _ := d.Get("rsa_bits").(string)
	return &models.RSA{
		RsaBits: rsaBits,
	}
}

func RSAModelFromMap(m map[string]interface{}) *models.RSA {
	rsaBits := m["rsa_bits"].(string)
	return &models.RSA{
		RsaBits: rsaBits,
	}
}

func SetRSAResourceData(d *schema.ResourceData, m *models.RSA) {
	d.Set("rsa_bits", m.RsaBits)
}

func SetRSASubResourceData(m []*models.RSA) (d []*map[string]interface{}) {
	for _, RSAModel := range m {
		if RSAModel != nil {
			properties := make(map[string]interface{})
			properties["rsa_bits"] = RSAModel.RsaBits
			d = append(d, &properties)
		}
	}
	return
}

func RSA() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"rsa_bits": {
			Description: `RSA Encryption Key bit size.`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetRSAPropertyFields() (t []string) {
	return []string{
		"rsa_bits",
	}
}
