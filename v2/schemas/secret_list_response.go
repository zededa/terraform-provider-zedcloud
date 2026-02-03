package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func SecretListResponseModel(d *schema.ResourceData) *models.SecretListResponse {
	var secrets []*models.SecretInfoResponse // []*SecretInfoResponse
	secretsInterface, secretsIsSet := d.GetOk("secrets")
	if secretsIsSet {
		var items []interface{}
		if listItems, isList := secretsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = secretsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := SecretInfoResponseModelFromMap(v.(map[string]interface{}))
			secrets = append(secrets, m)
		}
	}
	totalCountInt, _ := d.Get("total_count").(int)
	totalCount := int32(totalCountInt)
	return &models.SecretListResponse{
		Secrets:    secrets,
		TotalCount: totalCount,
	}
}

func SecretListResponseModelFromMap(m map[string]interface{}) *models.SecretListResponse {
	var secrets []*models.SecretInfoResponse // []*SecretInfoResponse
	secretsInterface, secretsIsSet := m["secrets"]
	if secretsIsSet {
		var items []interface{}
		if listItems, isList := secretsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = secretsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := SecretInfoResponseModelFromMap(v.(map[string]interface{}))
			secrets = append(secrets, m)
		}
	}
	totalCount := int32(m["total_count"].(int)) // int32
	return &models.SecretListResponse{
		Secrets:    secrets,
		TotalCount: totalCount,
	}
}

func SetSecretListResponseResourceData(d *schema.ResourceData, m *models.SecretListResponse) {
	d.Set("secrets", SetSecretInfoResponseSubResourceData(m.Secrets))
	d.Set("total_count", m.TotalCount)
}

func SetSecretListResponseSubResourceData(m []*models.SecretListResponse) (d []*map[string]interface{}) {
	for _, SecretListResponseModel := range m {
		if SecretListResponseModel != nil {
			properties := make(map[string]interface{})
			properties["secrets"] = SetSecretInfoResponseSubResourceData(SecretListResponseModel.Secrets)
			properties["total_count"] = SecretListResponseModel.TotalCount
			d = append(d, &properties)
		}
	}
	return
}

func SecretListResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"secrets": {
			Description: `List of secrets with their detailed information`,
			Type:        schema.TypeList, //GoType: []*SecretInfoResponse
			Elem: &schema.Resource{
				Schema: SecretInfoResponseSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"total_count": {
			Description: `Total number of secrets available`,
			Type:        schema.TypeInt,
			Optional:    true,
		},
	}
}

func GetSecretListResponsePropertyFields() (t []string) {
	return []string{
		"secrets",
		"total_count",
	}
}
