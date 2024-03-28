package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func NetWifiConfigSecretsModel(d *schema.ResourceData) *models.NetWifiConfigSecrets {
	wiFiPasswd, _ := d.Get("wi_fi_passwd").(string)
	return &models.NetWifiConfigSecrets{
		WiFiPasswd: wiFiPasswd,
	}
}

func NetWifiConfigSecretsModelFromMap(m map[string]interface{}) *models.NetWifiConfigSecrets {
	wiFiPasswd := m["wi_fi_passwd"].(string)
	return &models.NetWifiConfigSecrets{
		WiFiPasswd: wiFiPasswd,
	}
}

// Update the underlying NetWifiConfigSecrets resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetNetWifiConfigSecretsResourceData(d *schema.ResourceData, m *models.NetWifiConfigSecrets) {
	d.Set("wi_fi_passwd", m.WiFiPasswd)
}

// Iterate through and update the NetWifiConfigSecrets resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetNetWifiConfigSecretsSubResourceData(m []*models.NetWifiConfigSecrets) (d []*map[string]interface{}) {
	for _, NetWifiConfigSecretsModel := range m {
		if NetWifiConfigSecretsModel != nil {
			properties := make(map[string]interface{})
			properties["wi_fi_passwd"] = NetWifiConfigSecretsModel.WiFiPasswd
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the NetWifiConfigSecrets resource defined in the Terraform configuration
func NetWifiConfigSecretsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"wi_fi_passwd": {
			Description: `Wifi Password`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the NetWifiConfigSecrets resource
func GetNetWifiConfigSecretsPropertyFields() (t []string) {
	return []string{
		"wi_fi_passwd",
	}
}
