package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func JWTAuthProfileAlgorithmModel(d *schema.ResourceData) *models.JWTAuthProfileAlgorithm {
	jWTAuthProfileAlgorithm, _ := d.Get("j_w_t_auth_profile_algorithm").(models.JWTAuthProfileAlgorithm)
	return &jWTAuthProfileAlgorithm
}

func JWTAuthProfileAlgorithmModelFromMap(m map[string]interface{}) *models.JWTAuthProfileAlgorithm {
	jWTAuthProfileAlgorithm := m["j_w_t_auth_profile_algorithm"].(models.JWTAuthProfileAlgorithm)
	return &jWTAuthProfileAlgorithm
}

func SetJWTAuthProfileAlgorithmResourceData(d *schema.ResourceData, m *models.JWTAuthProfileAlgorithm) {
}

func SetJWTAuthProfileAlgorithmSubResourceData(m []*models.JWTAuthProfileAlgorithm) (d []*map[string]interface{}) {
	for _, JWTAuthProfileAlgorithmModel := range m {
		if JWTAuthProfileAlgorithmModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func JWTAuthProfileAlgorithmSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func GetJWTAuthProfileAlgorithmPropertyFields() (t []string) {
	return []string{}
}
