package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

func IPAssignmentModel(d *schema.ResourceData) *models.IPAssignment {
	var iPAddress []string
	iPAddressInterface, iPAddressIsSet := d.GetOk("iPAddress")
	if iPAddressIsSet {
		iPAddressSlice := iPAddressInterface.([]interface{})
		for _, i := range iPAddressSlice {
			iPAddressSlice = append(iPAddressSlice, i.(string))
		}
	}
	macAddress, _ := d.Get("mac_address").(string)
	return &models.IPAssignment{
		IPAddress:  iPAddress,
		MacAddress: macAddress,
	}
}

func IPAssignmentModelFromMap(m map[string]interface{}) *models.IPAssignment {
	var iPAddress []string
	iPAddressInterface, iPAddressIsSet := m["iPAddress"]
	if iPAddressIsSet {
		iPAddressSlice := iPAddressInterface.([]interface{})
		for _, i := range iPAddressSlice {
			iPAddressSlice = append(iPAddressSlice, i.(string))
		}
	}
	macAddress := m["mac_address"].(string)
	return &models.IPAssignment{
		IPAddress:  iPAddress,
		MacAddress: macAddress,
	}
}

// Update the underlying IPAssignment resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetIPAssignmentResourceData(d *schema.ResourceData, m *models.IPAssignment) {
	d.Set("ip_address", m.IPAddress)
	d.Set("mac_address", m.MacAddress)
}

// Iterate through and update the IPAssignment resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetIPAssignmentSubResourceData(m []*models.IPAssignment) (d []*map[string]interface{}) {
	for _, IPAssignmentModel := range m {
		if IPAssignmentModel != nil {
			properties := make(map[string]interface{})
			properties["ip_address"] = IPAssignmentModel.IPAddress
			properties["mac_address"] = IPAssignmentModel.MacAddress
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the IPAssignment resource defined in the Terraform configuration
func IPAssignmentSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"ip_address": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"mac_address": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the IPAssignment resource
func GetIPAssignmentPropertyFields() (t []string) {
	return []string{
		"ip_address",
		"mac_address",
	}
}
