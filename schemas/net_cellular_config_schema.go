package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func NetCellularConfigModel(d *schema.ResourceData) *models.NetCellularConfig {
	aPN, _ := d.Get("a_p_n").(string)
	locationTracking, _ := d.Get("location_tracking").(bool)
	return &models.NetCellularConfig{
		APN:              aPN,
		LocationTracking: locationTracking,
	}
}

func NetCellularConfigModelFromMap(m map[string]interface{}) *models.NetCellularConfig {
	aPN := m["a_p_n"].(string)
	locationTracking := m["location_tracking"].(bool)
	return &models.NetCellularConfig{
		APN:              aPN,
		LocationTracking: locationTracking,
	}
}

// Update the underlying NetCellularConfig resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetNetCellularConfigResourceData(d *schema.ResourceData, m *models.NetCellularConfig) {
	d.Set("a_p_n", m.APN)
	d.Set("location_tracking", m.LocationTracking)
}

// Iterate through and update the NetCellularConfig resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetNetCellularConfigSubResourceData(m []*models.NetCellularConfig) (d []*map[string]interface{}) {
	for _, NetCellularConfigModel := range m {
		if NetCellularConfigModel != nil {
			properties := make(map[string]interface{})
			properties["a_p_n"] = NetCellularConfigModel.APN
			properties["location_tracking"] = NetCellularConfigModel.LocationTracking
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the NetCellularConfig resource defined in the Terraform configuration
func NetCellularConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"a_p_n": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"location_tracking": {
			Description: ``,
			Type:        schema.TypeBool,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the NetCellularConfig resource
func GetNetCellularConfigPropertyFields() (t []string) {
	return []string{
		"a_p_n",
		"location_tracking",
	}
}
