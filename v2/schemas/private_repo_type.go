package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func PrivateRepoTypeModel(d *schema.ResourceData) *models.PrivateRepoType {
	privateRepoType, _ := d.Get("private_repo_type").(models.PrivateRepoType)
	return &privateRepoType
}

func PrivateRepoTypeModelFromMap(m map[string]interface{}) *models.PrivateRepoType {
	privateRepoType := m["private_repo_type"].(models.PrivateRepoType)
	return &privateRepoType
}

func SetPrivateRepoTypeResourceData(d *schema.ResourceData, m *models.PrivateRepoType) {
}

func SetPrivateRepoTypeSubResourceData(m []*models.PrivateRepoType) (d []*map[string]interface{}) {
	for _, PrivateRepoTypeModel := range m {
		if PrivateRepoTypeModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func PrivateRepoTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func GetPrivateRepoTypePropertyFields() (t []string) {
	return []string{}
}
