package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate NetProxyServer resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func NetProxyServerModel(d *schema.ResourceData) *models.NetProxyServer {
	portInt, _ := d.Get("port").(int)
	port := int64(portInt)
	protoModel, ok := d.Get("proto").(models.NetworkProxyProto) // NetworkProxyProto
	proto := &protoModel
	if !ok {
		proto = nil
	}
	server, _ := d.Get("server").(string)
	return &models.NetProxyServer{
		Port:   port,
		Proto:  proto,
		Server: server,
	}
}

func NetProxyServerModelFromMap(m map[string]interface{}) *models.NetProxyServer {
	port := int64(m["port"].(int))                  // int64 false false false
	proto := m["proto"].(*models.NetworkProxyProto) // NetworkProxyProto
	server := m["server"].(string)
	return &models.NetProxyServer{
		Port:   port,
		Proto:  proto,
		Server: server,
	}
}

// Update the underlying NetProxyServer resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetNetProxyServerResourceData(d *schema.ResourceData, m *models.NetProxyServer) {
	d.Set("port", m.Port)
	d.Set("proto", m.Proto)
	d.Set("server", m.Server)
}

// Iterate through and update the NetProxyServer resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetNetProxyServerSubResourceData(m []*models.NetProxyServer) (d []*map[string]interface{}) {
	for _, NetProxyServerModel := range m {
		if NetProxyServerModel != nil {
			properties := make(map[string]interface{})
			properties["port"] = NetProxyServerModel.Port
			properties["proto"] = NetProxyServerModel.Proto
			properties["server"] = NetProxyServerModel.Server
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the NetProxyServer resource defined in the Terraform configuration
func NetProxyServerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"port": {
			Description: `Net Proxy Port`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"proto": {
			Description: `Net Proxy proto`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"server": {
			Description: `Net Proxy Server`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the NetProxyServer resource
func GetNetProxyServerPropertyFields() (t []string) {
	return []string{
		"port",
		"proto",
		"server",
	}
}
