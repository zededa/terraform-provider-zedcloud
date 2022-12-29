package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate SKUDetail resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func SKUDetailModel(d *schema.ResourceData) *models.SKUDetail {
	capacity, _ := d.Get("capacity").(string)
	name, _ := d.Get("name").(string)
	tier, _ := d.Get("tier").(string)
	return &models.SKUDetail{
		Capacity: capacity,
		Name:     name,
		Tier:     tier,
	}
}

func SKUDetailModelFromMap(m map[string]interface{}) *models.SKUDetail {
	capacity := m["capacity"].(string)
	name := m["name"].(string)
	tier := m["tier"].(string)
	return &models.SKUDetail{
		Capacity: capacity,
		Name:     name,
		Tier:     tier,
	}
}

// Update the underlying SKUDetail resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetSKUDetailResourceData(d *schema.ResourceData, m *models.SKUDetail) {
	d.Set("capacity", m.Capacity)
	d.Set("name", m.Name)
	d.Set("tier", m.Tier)
}

// Iterate through and update the SKUDetail resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetSKUDetailSubResourceData(m []*models.SKUDetail) (d []*map[string]interface{}) {
	for _, SKUDetailModel := range m {
		if SKUDetailModel != nil {
			properties := make(map[string]interface{})
			properties["capacity"] = SKUDetailModel.Capacity
			properties["name"] = SKUDetailModel.Name
			properties["tier"] = SKUDetailModel.Tier
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the SKUDetail resource defined in the Terraform configuration
func SKUDetailSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"capacity": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"name": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"tier": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the SKUDetail resource
func GetSKUDetailPropertyFields() (t []string) {
	return []string{
		"capacity",
		"name",
		"tier",
	}
}
