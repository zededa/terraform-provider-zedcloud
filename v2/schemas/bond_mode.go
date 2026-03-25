package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func BondModeModel(d *schema.ResourceData) *models.BondMode {
	bondMode, _ := d.Get("bond_mode").(models.BondMode)
	return &bondMode
}

func BondModeModelFromMap(m map[string]interface{}) *models.BondMode {
	bondMode := m["bond_mode"].(models.BondMode)
	return &bondMode
}

func SetBondModeResourceData(d *schema.ResourceData, m *models.BondMode) {
}

func SetBondModeSubResourceData(m []*models.BondMode) (d []*map[string]interface{}) {
	for _, BondModeModel := range m {
		if BondModeModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func BondModeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func GetBondModePropertyFields() (t []string) {
	return []string{}
}
