package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate ResourceGroupDetail resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ResourceGroupDetailModel(d *schema.ResourceData) *models.ResourceGroupDetail {
	name, _ := d.Get("name").(string)
	region, _ := d.Get("region").(string)
	subscriptionID, _ := d.Get("subscription_id").(string)
	return &models.ResourceGroupDetail{
		Name:           &name, // string true false false
		Region:         region,
		SubscriptionID: &subscriptionID, // string true false false
	}
}

func ResourceGroupDetailModelFromMap(m map[string]interface{}) *models.ResourceGroupDetail {
	name := m["name"].(string)
	region := m["region"].(string)
	subscriptionID := m["subscription_id"].(string)
	return &models.ResourceGroupDetail{
		Name:           &name,
		Region:         region,
		SubscriptionID: &subscriptionID,
	}
}

// Update the underlying ResourceGroupDetail resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetResourceGroupDetailResourceData(d *schema.ResourceData, m *models.ResourceGroupDetail) {
	d.Set("name", m.Name)
	d.Set("region", m.Region)
	d.Set("subscription_id", m.SubscriptionID)
}

// Iterate through and update the ResourceGroupDetail resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetResourceGroupDetailSubResourceData(m []*models.ResourceGroupDetail) (d []*map[string]interface{}) {
	for _, ResourceGroupDetailModel := range m {
		if ResourceGroupDetailModel != nil {
			properties := make(map[string]interface{})
			properties["name"] = ResourceGroupDetailModel.Name
			properties["region"] = ResourceGroupDetailModel.Region
			properties["subscription_id"] = ResourceGroupDetailModel.SubscriptionID
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the ResourceGroupDetail resource defined in the Terraform configuration
func ResourceGroupDetail() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Description: `resource group name`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"region": {
			Description: `resource group region`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"subscription_id": {
			Description: `azure subscription id to which resource group is attached`,
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

// Retrieve property field names for updating the ResourceGroupDetail resource
func GetResourceGroupDetailPropertyFields() (t []string) {
	return []string{
		"name",
		"region",
		"subscription_id",
	}
}
