package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func OAUTHProfileJWTAuthProfileModel(d *schema.ResourceData) *models.OAUTHProfileJWTAuthProfile {
	var alg *models.JWTAuthProfileAlgorithm // JWTAuthProfileAlgorithm
	algInterface, algIsSet := d.GetOk("alg")
	if algIsSet {
		algModel := algInterface.(string)
		alg = models.NewJWTAuthProfileAlgorithm(models.JWTAuthProfileAlgorithm(algModel))
	}
	return &models.OAUTHProfileJWTAuthProfile{
		Alg: alg,
	}
}

func OAUTHProfileJWTAuthProfileModelFromMap(m map[string]interface{}) *models.OAUTHProfileJWTAuthProfile {
	var alg *models.JWTAuthProfileAlgorithm // JWTAuthProfileAlgorithm
	algInterface, algIsSet := m["alg"]
	if algIsSet {
		algModel := algInterface.(string)
		alg = models.NewJWTAuthProfileAlgorithm(models.JWTAuthProfileAlgorithm(algModel))
	}
	return &models.OAUTHProfileJWTAuthProfile{
		Alg: alg,
	}
}

func SetOAUTHProfileJWTAuthProfileResourceData(d *schema.ResourceData, m *models.OAUTHProfileJWTAuthProfile) {
	d.Set("alg", m.Alg)
}

func SetOAUTHProfileJWTAuthProfileSubResourceData(m []*models.OAUTHProfileJWTAuthProfile) (d []*map[string]interface{}) {
	for _, OAUTHProfileJWTAuthProfileModel := range m {
		if OAUTHProfileJWTAuthProfileModel != nil {
			properties := make(map[string]interface{})
			properties["alg"] = OAUTHProfileJWTAuthProfileModel.Alg
			d = append(d, &properties)
		}
	}
	return
}

func OAUTHProfileJWTAuthProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"alg": {
			Description: `Algorithm for JWT signature verification`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetOAUTHProfileJWTAuthProfilePropertyFields() (t []string) {
	return []string{
		"alg",
	}
}
