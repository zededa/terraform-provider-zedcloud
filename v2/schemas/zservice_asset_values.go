package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ZserviceAssetValuesModel(d *schema.ResourceData) *models.ZserviceAssetValues {
	datastoreID, _ := d.Get("datastore_id").(string)
	url, _ := d.Get("url").(string)
	return &models.ZserviceAssetValues{
		DatastoreID: datastoreID,
		URL:         url,
	}
}

func ZserviceAssetValuesModelFromMap(m map[string]interface{}) *models.ZserviceAssetValues {
	datastoreID := m["datastore_id"].(string)
	url := m["url"].(string)
	return &models.ZserviceAssetValues{
		DatastoreID: datastoreID,
		URL:         url,
	}
}

func SetZserviceAssetValuesResourceData(d *schema.ResourceData, m *models.ZserviceAssetValues) {
	d.Set("datastore_id", m.DatastoreID)
	d.Set("url", m.URL)
}

func SetZserviceAssetValuesSubResourceData(m []*models.ZserviceAssetValues) (d []*map[string]interface{}) {
	for _, ZserviceAssetValuesModel := range m {
		if ZserviceAssetValuesModel != nil {
			properties := make(map[string]interface{})
			properties["datastore_id"] = ZserviceAssetValuesModel.DatastoreID
			properties["url"] = ZserviceAssetValuesModel.URL
			d = append(d, &properties)
		}
	}
	return
}

func ZserviceAssetValuesSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"datastore_id": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"url": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetZserviceAssetValuesPropertyFields() (t []string) {
	return []string{
		"datastore_id",
		"url",
	}
}
