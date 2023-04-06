package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func ECDSAModel(d *schema.ResourceData) *models.ECDSA {
	curve, _ := d.Get("curve").(string)
	return &models.ECDSA{
		Curve: curve,
	}
}

func ECDSAModelFromMap(m map[string]interface{}) *models.ECDSA {
	curve := m["curve"].(string)
	return &models.ECDSA{
		Curve: curve,
	}
}

func SetECDSAResourceData(d *schema.ResourceData, m *models.ECDSA) {
	d.Set("curve", m.Curve)
}

func SetECDSASubResourceData(m []*models.ECDSA) (d []*map[string]interface{}) {
	for _, ECDSAModel := range m {
		if ECDSAModel != nil {
			properties := make(map[string]interface{})
			properties["curve"] = ECDSAModel.Curve
			d = append(d, &properties)
		}
	}
	return
}

func ECDSA() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"curve": {
			Description: `ECDSA curve to be used while signing the certificate.`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetECDSAPropertyFields() (t []string) {
	return []string{
		"curve",
	}
}
