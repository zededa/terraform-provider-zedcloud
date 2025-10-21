package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ConfigConnectivityProbeModel(d *schema.ResourceData) *models.ConfigConnectivityProbe {
	var probeEndpoint *models.ConfigProbeEndpoint // ConfigProbeEndpoint
	probeEndpointInterface, probeEndpointIsSet := d.GetOk("probe_endpoint")
	if probeEndpointIsSet && probeEndpointInterface != nil {
		probeEndpointMap := probeEndpointInterface.([]interface{})
		if len(probeEndpointMap) > 0 {
			probeEndpoint = ConfigProbeEndpointModelFromMap(probeEndpointMap[0].(map[string]interface{}))
		}
	}
	var probeMethod *models.ConfigConnectivityProbeMethod // ConfigConnectivityProbeMethod
	probeMethodInterface, probeMethodIsSet := d.GetOk("probe_method")
	if probeMethodIsSet {
		probeMethodModel := probeMethodInterface.(string)
		probeMethod = models.NewConfigConnectivityProbeMethod(models.ConfigConnectivityProbeMethod(probeMethodModel))
	}
	return &models.ConfigConnectivityProbe{
		ProbeEndpoint: probeEndpoint,
		ProbeMethod:   probeMethod,
	}
}

func ConfigConnectivityProbeModelFromMap(m map[string]interface{}) *models.ConfigConnectivityProbe {
	var probeEndpoint *models.ConfigProbeEndpoint // ConfigProbeEndpoint
	probeEndpointInterface, probeEndpointIsSet := m["probe_endpoint"]
	if probeEndpointIsSet && probeEndpointInterface != nil {
		probeEndpointMap := probeEndpointInterface.([]interface{})
		if len(probeEndpointMap) > 0 {
			probeEndpoint = ConfigProbeEndpointModelFromMap(probeEndpointMap[0].(map[string]interface{}))
		}
	}
	//
	var probeMethod *models.ConfigConnectivityProbeMethod // ConfigConnectivityProbeMethod
	probeMethodInterface, probeMethodIsSet := m["probe_method"]
	if probeMethodIsSet {
		probeMethodModel := probeMethodInterface.(string)
		probeMethod = models.NewConfigConnectivityProbeMethod(models.ConfigConnectivityProbeMethod(probeMethodModel))
	}
	return &models.ConfigConnectivityProbe{
		ProbeEndpoint: probeEndpoint,
		ProbeMethod:   probeMethod,
	}
}

func SetConfigConnectivityProbeResourceData(d *schema.ResourceData, m *models.ConfigConnectivityProbe) {
	d.Set("probe_endpoint", SetConfigProbeEndpointSubResourceData([]*models.ConfigProbeEndpoint{m.ProbeEndpoint}))
	d.Set("probe_method", m.ProbeMethod)
}

func SetConfigConnectivityProbeSubResourceData(m []*models.ConfigConnectivityProbe) (d []*map[string]interface{}) {
	for _, ConfigConnectivityProbeModel := range m {
		if ConfigConnectivityProbeModel != nil {
			properties := make(map[string]interface{})
			properties["probe_endpoint"] = SetConfigProbeEndpointSubResourceData([]*models.ConfigProbeEndpoint{ConfigConnectivityProbeModel.ProbeEndpoint})
			properties["probe_method"] = ConfigConnectivityProbeModel.ProbeMethod
			d = append(d, &properties)
		}
	}
	return
}

func ConfigConnectivityProbeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"probe_endpoint": {
			Description: `Endpoint to probe using the selected probing mechanism to determine
the connectivity status.`,
			Type: schema.TypeList, //GoType: ConfigProbeEndpoint
			Elem: &schema.Resource{
				Schema: ConfigProbeEndpointSchema(),
			},
			Optional: true,
		},

		"probe_method": {
			Description: `Method to use to determine the connectivity status.`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetConfigConnectivityProbePropertyFields() (t []string) {
	return []string{
		"probe_endpoint",
		"probe_method",
	}
}
