package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ZksInstanceK3sBulkUpgradeDataModel(d *schema.ResourceData) *models.ZksInstanceK3sBulkUpgradeData {
	version, _ := d.Get("version").(string)
	return &models.ZksInstanceK3sBulkUpgradeData{
		Version: version,
	}
}

func ZksInstanceK3sBulkUpgradeDataModelFromMap(m map[string]interface{}) *models.ZksInstanceK3sBulkUpgradeData {
	version := m["version"].(string)
	return &models.ZksInstanceK3sBulkUpgradeData{
		Version: version,
	}
}

func SetZksInstanceK3sBulkUpgradeDataResourceData(d *schema.ResourceData, m *models.ZksInstanceK3sBulkUpgradeData) {
	d.Set("version", m.Version)
}

func SetZksInstanceK3sBulkUpgradeDataSubResourceData(m []*models.ZksInstanceK3sBulkUpgradeData) (d []*map[string]interface{}) {
	for _, ZksInstanceK3sBulkUpgradeDataModel := range m {
		if ZksInstanceK3sBulkUpgradeDataModel != nil {
			properties := make(map[string]interface{})
			properties["version"] = ZksInstanceK3sBulkUpgradeDataModel.Version
			d = append(d, &properties)
		}
	}
	return
}

func ZksInstanceK3sBulkUpgradeDataSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"version": {
			Description: `K3s version to upgrade to`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetZksInstanceK3sBulkUpgradeDataPropertyFields() (t []string) {
	return []string{
		"version",
	}
}
