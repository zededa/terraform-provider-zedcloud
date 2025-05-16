package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ZserviceAssetGroupCreateResponseModel(d *schema.ResourceData) *models.ZserviceAssetGroupCreateResponse {
	assetGroupID, _ := d.Get("asset_group_id").(string)
	error, _ := d.Get("error").(string)
	return &models.ZserviceAssetGroupCreateResponse{
		AssetGroupID: assetGroupID,
		Error:        error,
	}
}

func ZserviceAssetGroupCreateResponseModelFromMap(m map[string]interface{}) *models.ZserviceAssetGroupCreateResponse {
	assetGroupID := m["asset_group_id"].(string)
	error := m["error"].(string)
	return &models.ZserviceAssetGroupCreateResponse{
		AssetGroupID: assetGroupID,
		Error:        error,
	}
}

func SetZserviceAssetGroupCreateResponseResourceData(d *schema.ResourceData, m *models.ZserviceAssetGroupCreateResponse) {
	d.Set("asset_group_id", m.AssetGroupID)
	d.Set("error", m.Error)
}

func SetZserviceAssetGroupCreateResponseSubResourceData(m []*models.ZserviceAssetGroupCreateResponse) (d []*map[string]interface{}) {
	for _, ZserviceAssetGroupCreateResponseModel := range m {
		if ZserviceAssetGroupCreateResponseModel != nil {
			properties := make(map[string]interface{})
			properties["asset_group_id"] = ZserviceAssetGroupCreateResponseModel.AssetGroupID
			properties["error"] = ZserviceAssetGroupCreateResponseModel.Error
			d = append(d, &properties)
		}
	}
	return
}

func ZserviceAssetGroupCreateResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"asset_group_id": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"error": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetZserviceAssetGroupCreateResponsePropertyFields() (t []string) {
	return []string{
		"asset_group_id",
		"error",
	}
}
