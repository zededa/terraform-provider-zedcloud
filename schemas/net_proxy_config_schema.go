package schemas

import (
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func NetProxyConfigModel(d *schema.ResourceData) *models.NetProxyConfig {
	exceptions, _ := d.Get("exceptions").(string)
	networkProxy, _ := d.Get("network_proxy").(bool)
	networkProxyCerts, _ := d.Get("network_proxy_certs").([]models.StrfmtBase64) // []strfmt.Base64
	networkProxyURL, _ := d.Get("network_proxy_url").(string)
	pacfile, _ := d.Get("pacfile").(string)
	var proxies []*models.NetProxyServer // []*NetProxyServer
	proxiesInterface, proxiesIsSet := d.GetOk("proxies")
	if proxiesIsSet {
		var items []interface{}
		if listItems, isList := proxiesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = proxiesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := NetProxyServerModelFromMap(v.(map[string]interface{}))
			proxies = append(proxies, m)
		}
	}
	return &models.NetProxyConfig{
		Exceptions:        exceptions,
		NetworkProxy:      networkProxy,
		NetworkProxyCerts: networkProxyCerts,
		NetworkProxyURL:   networkProxyURL,
		Pacfile:           pacfile,
		Proxies:           proxies,
	}
}

func NetProxyConfigModelFromMap(m map[string]interface{}) *models.NetProxyConfig {
	exceptions := m["exceptions"].(string)
	networkProxy := m["network_proxy"].(bool)
	networkProxyCerts := m["network_proxy_certs"].([]models.StrfmtBase64) // []strfmt.Base64
	networkProxyURL := m["network_proxy_url"].(string)
	pacfile := m["pacfile"].(string)
	var proxies []*models.NetProxyServer // []*NetProxyServer
	proxiesInterface, proxiesIsSet := m["proxies"]
	if proxiesIsSet {
		var items []interface{}
		if listItems, isList := proxiesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = proxiesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := NetProxyServerModelFromMap(v.(map[string]interface{}))
			proxies = append(proxies, m)
		}
	}
	return &models.NetProxyConfig{
		Exceptions:        exceptions,
		NetworkProxy:      networkProxy,
		NetworkProxyCerts: networkProxyCerts,
		NetworkProxyURL:   networkProxyURL,
		Pacfile:           pacfile,
		Proxies:           proxies,
	}
}

// Update the underlying NetProxyConfig resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetNetProxyConfigResourceData(d *schema.ResourceData, m *models.NetProxyConfig) {
	d.Set("exceptions", m.Exceptions)
	d.Set("network_proxy", m.NetworkProxy)
	d.Set("network_proxy_certs", m.NetworkProxyCerts)
	d.Set("network_proxy_url", m.NetworkProxyURL)
	d.Set("pacfile", m.Pacfile)
	d.Set("proxies", SetNetProxyServerSubResourceData(m.Proxies))
}

// Iterate through and update the NetProxyConfig resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetNetProxyConfigSubResourceData(m []*models.NetProxyConfig) (d []*map[string]interface{}) {
	for _, NetProxyConfigModel := range m {
		if NetProxyConfigModel != nil {
			properties := make(map[string]interface{})
			properties["exceptions"] = NetProxyConfigModel.Exceptions
			properties["network_proxy"] = NetProxyConfigModel.NetworkProxy
			properties["network_proxy_certs"] = NetProxyConfigModel.NetworkProxyCerts
			properties["network_proxy_url"] = NetProxyConfigModel.NetworkProxyURL
			properties["pacfile"] = NetProxyConfigModel.Pacfile
			properties["proxies"] = SetNetProxyServerSubResourceData(NetProxyConfigModel.Proxies)
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the NetProxyConfig resource defined in the Terraform configuration
func NetProxyConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"exceptions": {
			Description: `Proxy exceptions`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"network_proxy": {
			Description: `Network proxy`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"network_proxy_certs": {
			Description: `Network Proxy Certificates`,
			Type:        schema.TypeList, //GoType: []strfmt.Base64
			Elem: &schema.Resource{
				Schema: strfmt.Base64Schema(),
			},
			Optional: true,
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
			Description: `Net Proxy: protocol level proxies`,
			Type:        schema.TypeList, //GoType: []*NetProxyServer
			Elem: &schema.Resource{
				Schema: NetProxyServerSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
	}
}

// Retrieve property field names for updating the NetProxyConfig resource
func GetNetProxyConfigPropertyFields() (t []string) {
	return []string{
		"exceptions",
		"network_proxy",
		"network_proxy_certs",
		"network_proxy_url",
		"pacfile",
		"proxies",
	}
}
