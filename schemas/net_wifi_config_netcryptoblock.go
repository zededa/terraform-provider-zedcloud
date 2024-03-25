package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func NetWifiConfigNetcryptoblockModel(d *schema.ResourceData) *models.NetWifiConfigNetcryptoblock {
	identity, _ := d.Get("identity").(string)
	password, _ := d.Get("password").(string)
	return &models.NetWifiConfigNetcryptoblock{
		Identity: identity,
		Password: password,
	}
}

func NetWifiConfigNetcryptoblockModelFromMap(m map[string]interface{}) *models.NetWifiConfigNetcryptoblock {
	identity := m["identity"].(string)
	password := m["password"].(string)
	return &models.NetWifiConfigNetcryptoblock{
		Identity: identity,
		Password: password,
	}
}

// Update the underlying NetWifiConfigNetcryptoblock resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetNetWifiConfigNetcryptoblockResourceData(d *schema.ResourceData, m *models.NetWifiConfigNetcryptoblock) {
	d.Set("identity", m.Identity)
	d.Set("password", m.Password)
}

// Iterate through and update the NetWifiConfigNetcryptoblock resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetNetWifiConfigNetcryptoblockSubResourceData(m []*models.NetWifiConfigNetcryptoblock) (d []*map[string]interface{}) {
	for _, NetWifiConfigNetcryptoblockModel := range m {
		if NetWifiConfigNetcryptoblockModel != nil {
			properties := make(map[string]interface{})
			properties["identity"] = NetWifiConfigNetcryptoblockModel.Identity
			properties["password"] = NetWifiConfigNetcryptoblockModel.Password
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the NetWifiConfigNetcryptoblock resource defined in the Terraform configuration
func NetWifiConfigNetcryptoblockSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"identity": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
		},

		"password": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
		},
	}
}

// Retrieve property field names for updating the NetWifiConfigNetcryptoblock resource
func GetNetWifiConfigNetcryptoblockPropertyFields() (t []string) {
	return []string{
		"identity",
		"password",
	}
}
