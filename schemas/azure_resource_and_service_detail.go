package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AzureResourceAndServiceDetailModel(d *schema.ResourceData) *models.AzureResourceAndServiceDetail {
	var sKU *models.SKUDetail // SKUDetail
	SKUInterface, SKUIsSet := d.GetOk("s_k_u")
	if SKUIsSet && SKUInterface != nil {
		SKUMap := SKUInterface.([]interface{})
		if len(SKUMap) > 0 {
			sKU = SKUDetailModelFromMap(SKUMap[0].(map[string]interface{}))
		}
	}
	createByDefault, _ := d.Get("create_by_default").(bool)
	name, _ := d.Get("name").(string)
	region, _ := d.Get("region").(string)
	resourceGroupName, _ := d.Get("resource_group_name").(string)
	subscriptionID, _ := d.Get("subscription_id").(string)
	return &models.AzureResourceAndServiceDetail{
		SKU:               sKU,
		CreateByDefault:   createByDefault,
		Name:              name,
		Region:            region,
		ResourceGroupName: resourceGroupName,
		SubscriptionID:    subscriptionID,
	}
}

func AzureResourceAndServiceDetailModelFromMap(m map[string]interface{}) *models.AzureResourceAndServiceDetail {
	var sKU *models.SKUDetail // SKUDetail
	SKUInterface, SKUIsSet := m["s_k_u"]
	if SKUIsSet && SKUInterface != nil {
		SKUMap := SKUInterface.([]interface{})
		if len(SKUMap) > 0 {
			sKU = SKUDetailModelFromMap(SKUMap[0].(map[string]interface{}))
		}
	}
	//
	createByDefault := m["create_by_default"].(bool)
	name := m["name"].(string)
	region := m["region"].(string)
	resourceGroupName := m["resource_group_name"].(string)
	subscriptionID := m["subscription_id"].(string)
	return &models.AzureResourceAndServiceDetail{
		SKU:               sKU,
		CreateByDefault:   createByDefault,
		Name:              name,
		Region:            region,
		ResourceGroupName: resourceGroupName,
		SubscriptionID:    subscriptionID,
	}
}

func SetAzureResourceAndServiceDetailResourceData(d *schema.ResourceData, m *models.AzureResourceAndServiceDetail) {
	d.Set("s_k_u", SetSKUDetailSubResourceData([]*models.SKUDetail{m.SKU}))
	d.Set("create_by_default", m.CreateByDefault)
	d.Set("name", m.Name)
	d.Set("region", m.Region)
	d.Set("resource_group_name", m.ResourceGroupName)
	d.Set("subscription_id", m.SubscriptionID)
}

func SetAzureResourceAndServiceDetailSubResourceData(m []*models.AzureResourceAndServiceDetail) (d []*map[string]interface{}) {
	for _, AzureResourceAndServiceDetailModel := range m {
		if AzureResourceAndServiceDetailModel != nil {
			properties := make(map[string]interface{})
			properties["s_k_u"] = SetSKUDetailSubResourceData([]*models.SKUDetail{AzureResourceAndServiceDetailModel.SKU})
			properties["create_by_default"] = AzureResourceAndServiceDetailModel.CreateByDefault
			properties["name"] = AzureResourceAndServiceDetailModel.Name
			properties["region"] = AzureResourceAndServiceDetailModel.Region
			properties["resource_group_name"] = AzureResourceAndServiceDetailModel.ResourceGroupName
			properties["subscription_id"] = AzureResourceAndServiceDetailModel.SubscriptionID
			d = append(d, &properties)
		}
	}
	return
}

func AzureResourceAndServiceDetail() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"s_k_u": {
			Description: ``,
			Type:        schema.TypeList, //GoType: SKUDetail
			Elem: &schema.Resource{
				Schema: SKUDetailSchema(),
			},
			Optional: true,
		},

		"create_by_default": {
			Description: ``,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"name": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"region": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"resource_group_name": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"subscription_id": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetAzureResourceAndServiceDetailPropertyFields() (t []string) {
	return []string{
		"s_k_u",
		"create_by_default",
		"name",
		"region",
		"resource_group_name",
		"subscription_id",
	}
}
