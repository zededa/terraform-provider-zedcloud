package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ConfigProbeEndpointModel(d *schema.ResourceData) *models.ConfigProbeEndpoint {
	host, _ := d.Get("host").(string)
	portInt, _ := d.Get("port").(int)
	port := int64(portInt)
	return &models.ConfigProbeEndpoint{
		Host: host,
		Port: port,
	}
}

func ConfigProbeEndpointModelFromMap(m map[string]interface{}) *models.ConfigProbeEndpoint {
	host := m["host"].(string)
	port := int64(m["port"].(int)) // int64
	return &models.ConfigProbeEndpoint{
		Host: host,
		Port: port,
	}
}

func SetConfigProbeEndpointResourceData(d *schema.ResourceData, m *models.ConfigProbeEndpoint) {
	d.Set("host", m.Host)
	d.Set("port", m.Port)
}

func SetConfigProbeEndpointSubResourceData(m []*models.ConfigProbeEndpoint) (d []*map[string]interface{}) {
	for _, ConfigProbeEndpointModel := range m {
		if ConfigProbeEndpointModel != nil {
			properties := make(map[string]interface{})
			properties["host"] = ConfigProbeEndpointModel.Host
			properties["port"] = ConfigProbeEndpointModel.Port
			d = append(d, &properties)
		}
	}
	return
}

func ConfigProbeEndpointSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"host": {
			Description: `IP address or FQDN.`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"port": {
			Description: `TCP port required for CONNECTIVITY_PROBE_METHOD_TCP.
Leave empty for CONNECTIVITY_PROBE_METHOD_ICMP.`,
			Type:     schema.TypeInt,
			Optional: true,
		},
	}
}

func GetConfigProbeEndpointPropertyFields() (t []string) {
	return []string{
		"host",
		"port",
	}
}
