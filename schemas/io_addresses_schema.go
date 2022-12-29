package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate IoAddresses resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func IoAddressesModel(d *schema.ResourceData) *models.IoAddresses {
	macAddress, _ := d.Get("mac_address").(string)
	return &models.IoAddresses{
		MacAddress: macAddress,
	}
}

func IoAddressesModelFromMap(m map[string]interface{}) *models.IoAddresses {
	macAddress := m["mac_address"].(string)
	return &models.IoAddresses{
		MacAddress: macAddress,
	}
}

// Update the underlying IoAddresses resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetIoAddressesResourceData(d *schema.ResourceData, m *models.IoAddresses) {
	d.Set("mac_address", m.MacAddress)
}

// Iterate through and update the IoAddresses resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetIoAddressesSubResourceData(m []*models.IoAddresses) (d []*map[string]interface{}) {
	for _, IoAddressesModel := range m {
		if IoAddressesModel != nil {
			properties := make(map[string]interface{})
			properties["mac_address"] = IoAddressesModel.MacAddress
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the IoAddresses resource defined in the Terraform configuration
func IoAddressesSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"mac_address": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the IoAddresses resource
func GetIoAddressesPropertyFields() (t []string) {
	return []string{
		"mac_address",
	}
}
