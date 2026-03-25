package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func LacpRateModel(d *schema.ResourceData) *models.LacpRate {
	lacpRate, _ := d.Get("lacp_rate").(models.LacpRate)
	return &lacpRate
}

func LacpRateModelFromMap(m map[string]interface{}) *models.LacpRate {
	lacpRate := m["lacp_rate"].(models.LacpRate)
	return &lacpRate
}

func SetLacpRateResourceData(d *schema.ResourceData, m *models.LacpRate) {
}

func SetLacpRateSubResourceData(m []*models.LacpRate) (d []*map[string]interface{}) {
	for _, LacpRateModel := range m {
		if LacpRateModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func LacpRateSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func GetLacpRatePropertyFields() (t []string) {
	return []string{}
}
