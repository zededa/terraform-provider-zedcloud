package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate IoMemberStatus resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func IoMemberStatusModel(d *schema.ResourceData) *models.IoMemberStatus {
	var ioAddress *models.IoAddresses // IoAddresses
	ioAddressInterface, ioAddressIsSet := d.GetOk("io_address")
	if ioAddressIsSet {
		ioAddressMap := ioAddressInterface.([]interface{})[0].(map[string]interface{})
		ioAddress = IoAddressesModelFromMap(ioAddressMap)
	}
	name := d.Get("name").(string)
	return &models.IoMemberStatus{
		IoAddress: ioAddress,
		Name:      name,
	}
}

func IoMemberStatusModelFromMap(m map[string]interface{}) *models.IoMemberStatus {
	var ioAddress *models.IoAddresses // IoAddresses
	ioAddressInterface, ioAddressIsSet := m["io_address"]
	if ioAddressIsSet {
		ioAddressMap := ioAddressInterface.([]interface{})[0].(map[string]interface{})
		ioAddress = IoAddressesModelFromMap(ioAddressMap)
	}
	//
	name := m["name"].(string)
	return &models.IoMemberStatus{
		IoAddress: ioAddress,
		Name:      name,
	}
}

// Update the underlying IoMemberStatus resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetIoMemberStatusResourceData(d *schema.ResourceData, m *models.IoMemberStatus) {
	d.Set("io_address", SetIoAddressesSubResourceData([]*models.IoAddresses{m.IoAddress}))
	d.Set("name", m.Name)
}

// Iterate throught and update the IoMemberStatus resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetIoMemberStatusSubResourceData(m []*models.IoMemberStatus) (d []*map[string]interface{}) {
	for _, IoMemberStatusModel := range m {
		if IoMemberStatusModel != nil {
			properties := make(map[string]interface{})
			properties["io_address"] = SetIoAddressesSubResourceData([]*models.IoAddresses{IoMemberStatusModel.IoAddress})
			properties["name"] = IoMemberStatusModel.Name
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the IoMemberStatus resource defined in the Terraform configuration
func IoMemberStatusSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"io_address": {
			Description: `IO addresses of the member. Each address corresponds to the member in the members array`,
			Type:        schema.TypeList, //GoType: IoAddresses
			Elem: &schema.Resource{
				Schema: IoAddressesSchema(),
			},
			Optional: true,
		},

		"name": {
			Description: `IoBundle Member name`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the IoMemberStatus resource
func GetIoMemberStatusPropertyFields() (t []string) {
	return []string{
		"io_address",
		"name",
	}
}
