package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ZserviceAssetTypeModel(d *schema.ResourceData) *models.ZserviceAssetType {
	zserviceAssetType, _ := d.Get("zservice_asset_type").(models.ZserviceAssetType)
	return &zserviceAssetType
}

func ZserviceAssetTypeModelFromMap(m map[string]interface{}) *models.ZserviceAssetType {
	zserviceAssetType := m["zservice_asset_type"].(models.ZserviceAssetType)
	return &zserviceAssetType
}

func SetZserviceAssetTypeResourceData(d *schema.ResourceData, m *models.ZserviceAssetType) {
}

func SetZserviceAssetTypeSubResourceData(m []*models.ZserviceAssetType) (d []*map[string]interface{}) {
	for _, ZserviceAssetTypeModel := range m {
		if ZserviceAssetTypeModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func ZserviceAssetTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func GetZserviceAssetTypePropertyFields() (t []string) {
	return []string{}
}
