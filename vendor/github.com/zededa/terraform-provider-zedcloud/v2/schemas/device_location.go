package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

// Function to perform the following actions:
// (1) Translate DeviceLocation resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func DeviceLocationModel(d *schema.ResourceData) *models.DeviceLocation {
	countInt, _ := d.Get("count").(int)
	count := int64(countInt)
	location, _ := d.Get("location").(string)
	return &models.DeviceLocation{
		Count:    count,
		Location: location,
	}
}

func DeviceLocationModelFromMap(m map[string]interface{}) *models.DeviceLocation {
	count := int64(m["count"].(int)) // int64 false false false
	location := m["location"].(string)
	return &models.DeviceLocation{
		Count:    count,
		Location: location,
	}
}

// Update the underlying DeviceLocation resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDeviceLocationResourceData(d *schema.ResourceData, m *models.DeviceLocation) {
	d.Set("count", m.Count)
	d.Set("location", m.Location)
}

// Iterate through and update the DeviceLocation resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDeviceLocationSubResourceData(m []*models.DeviceLocation) (d []*map[string]interface{}) {
	for _, DeviceLocationModel := range m {
		if DeviceLocationModel != nil {
			properties := make(map[string]interface{})
			properties["count"] = DeviceLocationModel.Count
			properties["location"] = DeviceLocationModel.Location
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the DeviceLocation resource defined in the Terraform configuration
func DeviceLocationSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"count": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"location": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the DeviceLocation resource
func GetDeviceLocationPropertyFields() (t []string) {
	return []string{
		"count",
		"location",
	}
}
