package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func IoMemberStatusModel(d *schema.ResourceData) *models.IoMemberStatus {
	var ioAddress *models.IoAddresses // IoAddresses
	ioAddressInterface, ioAddressIsSet := d.GetOk("io_address")
	if ioAddressIsSet && ioAddressInterface != nil {
		ioAddressMap := ioAddressInterface.([]interface{})
		if len(ioAddressMap) > 0 {
			ioAddress = IoAddressesModelFromMap(ioAddressMap[0].(map[string]interface{}))
		}
	}
	name, _ := d.Get("name").(string)
	return &models.IoMemberStatus{
		IoAddress: ioAddress,
		Name:      name,
	}
}

func IoMemberStatusModelFromMap(m map[string]interface{}) *models.IoMemberStatus {
	var ioAddress *models.IoAddresses // IoAddresses
	ioAddressInterface, ioAddressIsSet := m["io_address"]
	if ioAddressIsSet && ioAddressInterface != nil {
		ioAddressMap := ioAddressInterface.([]interface{})
		if len(ioAddressMap) > 0 {
			ioAddress = IoAddressesModelFromMap(ioAddressMap[0].(map[string]interface{}))
		}
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

// Iterate through and update the IoMemberStatus resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
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
