package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate AzureResourceAndServiceDetail resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func AzureResourceAndServiceDetailModel(d *schema.ResourceData) *models.AzureResourceAndServiceDetail {
	var sKU *models.SKUDetail // SKUDetail
	SKUInterface, SKUIsSet := d.GetOk("s_k_u")
	if SKUIsSet {
		SKUMap := SKUInterface.([]interface{})[0].(map[string]interface{})
		sKU = SKUDetailModelFromMap(SKUMap)
	}
	createByDefault := d.Get("create_by_default").(bool)
	name := d.Get("name").(string)
	region := d.Get("region").(string)
	resourceGroupName := d.Get("resource_group_name").(string)
	subscriptionID := d.Get("subscription_id").(string)
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
	if SKUIsSet {
		SKUMap := SKUInterface.([]interface{})[0].(map[string]interface{})
		sKU = SKUDetailModelFromMap(SKUMap)
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

// Update the underlying AzureResourceAndServiceDetail resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetAzureResourceAndServiceDetailResourceData(d *schema.ResourceData, m *models.AzureResourceAndServiceDetail) {
	d.Set("s_k_u", SetSKUDetailSubResourceData([]*models.SKUDetail{m.SKU}))
	d.Set("create_by_default", m.CreateByDefault)
	d.Set("name", m.Name)
	d.Set("region", m.Region)
	d.Set("resource_group_name", m.ResourceGroupName)
	d.Set("subscription_id", m.SubscriptionID)
}

// Iterate throught and update the AzureResourceAndServiceDetail resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
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

// Schema mapping representing the AzureResourceAndServiceDetail resource defined in the Terraform configuration
func AzureResourceAndServiceDetailSchema() map[string]*schema.Schema {
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

// Retrieve property field names for updating the AzureResourceAndServiceDetail resource
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
