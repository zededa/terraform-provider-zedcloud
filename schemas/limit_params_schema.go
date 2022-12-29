package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate LimitParams resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func LimitParamsModel(d *schema.ResourceData) *models.LimitParams {
	limitburstInt, _ := d.Get("limitburst").(int)
	limitburst := int64(limitburstInt)
	limitrateInt, _ := d.Get("limitrate").(int)
	limitrate := int64(limitrateInt)
	limitunit, _ := d.Get("limitunit").(string)
	return &models.LimitParams{
		Limitburst: limitburst,
		Limitrate:  limitrate,
		Limitunit:  limitunit,
	}
}

func LimitParamsModelFromMap(m map[string]interface{}) *models.LimitParams {
	limitburst := int64(m["limitburst"].(int)) // int64 false false false
	limitrate := int64(m["limitrate"].(int))   // int64 false false false
	limitunit := m["limitunit"].(string)
	return &models.LimitParams{
		Limitburst: limitburst,
		Limitrate:  limitrate,
		Limitunit:  limitunit,
	}
}

// Update the underlying LimitParams resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetLimitParamsResourceData(d *schema.ResourceData, m *models.LimitParams) {
	d.Set("limitburst", m.Limitburst)
	d.Set("limitrate", m.Limitrate)
	d.Set("limitunit", m.Limitunit)
}

// Iterate through and update the LimitParams resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetLimitParamsSubResourceData(m []*models.LimitParams) (d []*map[string]interface{}) {
	for _, LimitParamsModel := range m {
		if LimitParamsModel != nil {
			properties := make(map[string]interface{})
			properties["limitburst"] = LimitParamsModel.Limitburst
			properties["limitrate"] = LimitParamsModel.Limitrate
			properties["limitunit"] = LimitParamsModel.Limitunit
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the LimitParams resource defined in the Terraform configuration
func LimitParamsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"limitburst": {
			Description: `UI map: AppDetailsPage:EnvironmentsPane, AppDetailsPage:EnvironmentsPane`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"limitrate": {
			Description: `UI map: AppDetailsPage:EnvironmentsPane, AppDetailsPage:EnvironmentsPane`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"limitunit": {
			Description: `UI map: AppDetailsPage:EnvironmentsPane, AppDetailsPage:EnvironmentsPane`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the LimitParams resource
func GetLimitParamsPropertyFields() (t []string) {
	return []string{
		"limitburst",
		"limitrate",
		"limitunit",
	}
}
