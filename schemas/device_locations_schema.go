package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate DeviceLocations resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func DeviceLocationsModel(d *schema.ResourceData) *models.DeviceLocations {
	deviceLocations := d.Get("device_locations").([]*models.DeviceLocation) // []*DeviceLocation
	var next *models.Cursor                                                 // Cursor
	nextInterface, nextIsSet := d.GetOk("next")
	if nextIsSet {
		nextMap := nextInterface.([]interface{})[0].(map[string]interface{})
		next = CursorModelFromMap(nextMap)
	}
	return &models.DeviceLocations{
		DeviceLocations: deviceLocations,
		Next:            next,
	}
}

func DeviceLocationsModelFromMap(m map[string]interface{}) *models.DeviceLocations {
	deviceLocations := m["device_locations"].([]*models.DeviceLocation) // []*DeviceLocation
	var next *models.Cursor                                             // Cursor
	nextInterface, nextIsSet := m["next"]
	if nextIsSet {
		nextMap := nextInterface.([]interface{})[0].(map[string]interface{})
		next = CursorModelFromMap(nextMap)
	}
	//
	return &models.DeviceLocations{
		DeviceLocations: deviceLocations,
		Next:            next,
	}
}

// Update the underlying DeviceLocations resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDeviceLocationsResourceData(d *schema.ResourceData, m *models.DeviceLocations) {
	d.Set("device_locations", SetDeviceLocationSubResourceData(m.DeviceLocations))
	d.Set("next", SetCursorSubResourceData([]*models.Cursor{m.Next}))
}

// Iterate throught and update the DeviceLocations resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDeviceLocationsSubResourceData(m []*models.DeviceLocations) (d []*map[string]interface{}) {
	for _, DeviceLocationsModel := range m {
		if DeviceLocationsModel != nil {
			properties := make(map[string]interface{})
			properties["device_locations"] = SetDeviceLocationSubResourceData(DeviceLocationsModel.DeviceLocations)
			properties["next"] = SetCursorSubResourceData([]*models.Cursor{DeviceLocationsModel.Next})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the DeviceLocations resource defined in the Terraform configuration
func DeviceLocationsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"device_locations": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []*DeviceLocation
			Elem: &schema.Resource{
				Schema: DeviceLocationSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"next": {
			Description: ``,
			Type:        schema.TypeList, //GoType: Cursor
			Elem: &schema.Resource{
				Schema: CursorSchema(),
			},
			Optional: true,
		},
	}
}

// Retrieve property field names for updating the DeviceLocations resource
func GetDeviceLocationsPropertyFields() (t []string) {
	return []string{
		"device_locations",
		"next",
	}
}
