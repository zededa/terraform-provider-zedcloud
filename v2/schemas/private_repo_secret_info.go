package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func PrivateRepoSecretInfoModel(d *schema.ResourceData) *models.PrivateRepoSecretInfo {
	var data *models.SecretData // SecretData
	dataInterface, dataIsSet := d.GetOk("data")
	if dataIsSet && dataInterface != nil {
		dataMap := dataInterface.([]interface{})
		if len(dataMap) > 0 {
			data = SecretDataModelFromMap(dataMap[0].(map[string]interface{}))
		}
	}
	id, _ := d.Get("id").(string)
	name, _ := d.Get("name").(string)
	var typeVar *models.SecretType // SecretType
	typeInterface, typeIsSet := d.GetOk("type")
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewSecretType(models.SecretType(typeModel))
	}
	return &models.PrivateRepoSecretInfo{
		Data: data,
		ID:   id,
		Name: name,
		Type: typeVar,
	}
}

func PrivateRepoSecretInfoModelFromMap(m map[string]interface{}) *models.PrivateRepoSecretInfo {
	var data *models.SecretData // SecretData
	dataInterface, dataIsSet := m["data"]
	if dataIsSet && dataInterface != nil {
		dataMap := dataInterface.([]interface{})
		if len(dataMap) > 0 {
			data = SecretDataModelFromMap(dataMap[0].(map[string]interface{}))
		}
	}
	//
	id := m["id"].(string)
	name := m["name"].(string)
	var typeVar *models.SecretType // SecretType
	typeInterface, typeIsSet := m["type"]
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewSecretType(models.SecretType(typeModel))
	}
	return &models.PrivateRepoSecretInfo{
		Data: data,
		ID:   id,
		Name: name,
		Type: typeVar,
	}
}

func SetPrivateRepoSecretInfoResourceData(d *schema.ResourceData, m *models.PrivateRepoSecretInfo) {
	d.Set("data", SetSecretDataSubResourceData([]*models.SecretData{m.Data}))
	d.Set("id", m.ID)
	d.Set("name", m.Name)
	d.Set("type", m.Type)
}

func SetPrivateRepoSecretInfoSubResourceData(m []*models.PrivateRepoSecretInfo) (d []*map[string]interface{}) {
	for _, PrivateRepoSecretInfoModel := range m {
		if PrivateRepoSecretInfoModel != nil {
			properties := make(map[string]interface{})
			properties["data"] = SetSecretDataSubResourceData([]*models.SecretData{PrivateRepoSecretInfoModel.Data})
			properties["id"] = PrivateRepoSecretInfoModel.ID
			properties["name"] = PrivateRepoSecretInfoModel.Name
			properties["type"] = PrivateRepoSecretInfoModel.Type
			d = append(d, &properties)
		}
	}
	return
}

func PrivateRepoSecretInfoSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"data": {
			Description: `Secret credentials data`,
			Type:        schema.TypeList, //GoType: SecretData
			Elem: &schema.Resource{
				Schema: SecretDataSchema(),
			},
			Optional: true,
		},

		"id": {
			Description: `Unique identifier of the secret`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"name": {
			Description: `Name of the secret resource`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"type": {
			Description: `Type of secret for authentication`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetPrivateRepoSecretInfoPropertyFields() (t []string) {
	return []string{
		"data",
		"id",
		"name",
		"type",
	}
}
