package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func NetworkCellularModel(d *schema.ResourceData) *models.NetCellularConfig {
	aPN, _ := d.Get("apn").(string)
	locationTracking, _ := d.Get("location_tracking").(bool)
	return &models.NetCellularConfig{
		APN:              aPN,
		LocationTracking: locationTracking,
	}
}

func NetworkCellularModelFromMap(m map[string]interface{}) *models.NetCellularConfig {
	aPN := m["apn"].(string)
	locationTracking := m["location_tracking"].(bool)
	return &models.NetCellularConfig{
		APN:              aPN,
		LocationTracking: locationTracking,
	}
}

func NetworkCellularResourceData(d *schema.ResourceData, m *models.NetCellularConfig) {
	d.Set("apn", m.APN)
	d.Set("location_tracking", m.LocationTracking)
}

func SetNetworkCellularSubResourceData(m []*models.NetCellularConfig) (d []*map[string]interface{}) {
	for _, NetCellularConfigModel := range m {
		if NetCellularConfigModel != nil {
			properties := make(map[string]interface{})
			properties["apn"] = NetCellularConfigModel.APN
			properties["location_tracking"] = NetCellularConfigModel.LocationTracking
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the NetCellularConfig resource defined in the Terraform configuration
func NetworkCellular() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"apn": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"location_tracking": {
			Description: ``,
			Type:        schema.TypeBool,
			Optional:    true,
			Computed:    true,
		},
	}
}

// Retrieve property field names for updating the NetCellularConfig resource
func GetNetworkCellularPropertyFields() (t []string) {
	return []string{
		"apn",
		"location_tracking",
	}
}
