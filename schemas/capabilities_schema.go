package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate Capabilities resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func CapabilitiesModel(d *schema.ResourceData) *models.Capabilities {
	hWAssistedVirtualization := d.Get("h_w_assisted_virtualization").(bool)
	iOVirtualization := d.Get("i_o_virtualization").(bool)
	return &models.Capabilities{
		HWAssistedVirtualization: hWAssistedVirtualization,
		IOVirtualization:         iOVirtualization,
	}
}

func CapabilitiesModelFromMap(m map[string]interface{}) *models.Capabilities {
	hWAssistedVirtualization := m["h_w_assisted_virtualization"].(bool)
	iOVirtualization := m["i_o_virtualization"].(bool)
	return &models.Capabilities{
		HWAssistedVirtualization: hWAssistedVirtualization,
		IOVirtualization:         iOVirtualization,
	}
}

// Update the underlying Capabilities resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetCapabilitiesResourceData(d *schema.ResourceData, m *models.Capabilities) {
	d.Set("h_w_assisted_virtualization", m.HWAssistedVirtualization)
	d.Set("i_o_virtualization", m.IOVirtualization)
}

// Iterate throught and update the Capabilities resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetCapabilitiesSubResourceData(m []*models.Capabilities) (d []*map[string]interface{}) {
	for _, CapabilitiesModel := range m {
		if CapabilitiesModel != nil {
			properties := make(map[string]interface{})
			properties["h_w_assisted_virtualization"] = CapabilitiesModel.HWAssistedVirtualization
			properties["i_o_virtualization"] = CapabilitiesModel.IOVirtualization
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the Capabilities resource defined in the Terraform configuration
func CapabilitiesSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"h_w_assisted_virtualization": {
			Description: ``,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"i_o_virtualization": {
			Description: ``,
			Type:        schema.TypeBool,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the Capabilities resource
func GetCapabilitiesPropertyFields() (t []string) {
	return []string{
		"h_w_assisted_virtualization",
		"i_o_virtualization",
	}
}
