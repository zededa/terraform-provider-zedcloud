package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

// Function to perform the following actions:
// (1) Translate NetProxyStatus resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func NetProxyStatusModel(d *schema.ResourceData) *models.NetProxyStatus {
	exceptions, _ := d.Get("exceptions").(string)
	networkProxy, _ := d.Get("network_proxy").(bool)
	networkProxyURL, _ := d.Get("network_proxy_url").(string)
	pacfile, _ := d.Get("pacfile").(string)
	proxies, _ := d.Get("proxies").([]*models.Server) // []*NetProxyServer
	wpadProxyURL, _ := d.Get("wpad_proxy_url").(string)
	return &models.NetProxyStatus{
		Exceptions:      exceptions,
		NetworkProxy:    networkProxy,
		NetworkProxyURL: networkProxyURL,
		Pacfile:         pacfile,
		Proxies:         proxies,
		WpadProxyURL:    wpadProxyURL,
	}
}

func NetProxyStatusModelFromMap(m map[string]interface{}) *models.NetProxyStatus {
	exceptions := m["exceptions"].(string)
	networkProxy := m["network_proxy"].(bool)
	networkProxyURL := m["network_proxy_url"].(string)
	pacfile := m["pacfile"].(string)
	proxies := m["proxies"].([]*models.Server) // []*NetProxyServer
	wpadProxyURL := m["wpad_proxy_url"].(string)
	return &models.NetProxyStatus{
		Exceptions:      exceptions,
		NetworkProxy:    networkProxy,
		NetworkProxyURL: networkProxyURL,
		Pacfile:         pacfile,
		Proxies:         proxies,
		WpadProxyURL:    wpadProxyURL,
	}
}

// Update the underlying NetProxyStatus resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetNetProxyStatusResourceData(d *schema.ResourceData, m *models.NetProxyStatus) {
	d.Set("exceptions", m.Exceptions)
	d.Set("network_proxy", m.NetworkProxy)
	d.Set("network_proxy_url", m.NetworkProxyURL)
	d.Set("pacfile", m.Pacfile)
	d.Set("proxies", SetNetworkProxyServerSubResourceData(m.Proxies))
	d.Set("wpad_proxy_url", m.WpadProxyURL)
}

// Iterate through and update the NetProxyStatus resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetNetProxyStatusSubResourceData(m []*models.NetProxyStatus) (d []*map[string]interface{}) {
	for _, NetProxyStatusModel := range m {
		if NetProxyStatusModel != nil {
			properties := make(map[string]interface{})
			properties["exceptions"] = NetProxyStatusModel.Exceptions
			properties["network_proxy"] = NetProxyStatusModel.NetworkProxy
			properties["network_proxy_url"] = NetProxyStatusModel.NetworkProxyURL
			properties["pacfile"] = NetProxyStatusModel.Pacfile
			properties["proxies"] = SetNetworkProxyServerSubResourceData(NetProxyStatusModel.Proxies)
			properties["wpad_proxy_url"] = NetProxyStatusModel.WpadProxyURL
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the NetProxyStatus resource defined in the Terraform configuration
func NetProxyStatusSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"exceptions": {
			Description: `exceptions`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"network_proxy": {
			Description: `Use pacfile (Auto discover or manual upload)`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"network_proxy_url": {
			Description: `Network Proxy URL`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"pacfile": {
			Description: `proxy configuration in a pacfile`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"proxies": {
			Description: `protocol level proxies`,
			Type:        schema.TypeList, //GoType: []*NetProxyServer
			Elem: &schema.Resource{
				Schema: NetworkProxyServer(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"wpad_proxy_url": {
			Description: `WPAD Proxy URL`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the NetProxyStatus resource
func GetNetProxyStatusPropertyFields() (t []string) {
	return []string{
		"exceptions",
		"network_proxy",
		"network_proxy_url",
		"pacfile",
		"proxies",
		"wpad_proxy_url",
	}
}
