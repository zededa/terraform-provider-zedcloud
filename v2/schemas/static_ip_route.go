package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func StaticIPRouteModel(d *schema.ResourceData) *models.StaticIPRoute {
	gateway, _ := d.Get("gateway").(string)
	outputPort, _ := d.Get("output_port").(string)
	prefix, _ := d.Get("prefix").(string)
	var probeConfig *models.ProbeConfig // ProbeConfig
	probeConfigInterface, probeConfigIsSet := d.GetOk("probe_config")
	if probeConfigIsSet && probeConfigInterface != nil {
		probeConfigMap := probeConfigInterface.([]interface{})
		if len(probeConfigMap) > 0 {
			probeConfig = ProbeConfigModelFromMap(probeConfigMap[0].(map[string]interface{}))
		}
	}
	return &models.StaticIPRoute{
		Gateway:     gateway,
		OutputPort:  outputPort,
		Prefix:      prefix,
		ProbeConfig: probeConfig,
	}
}

func StaticIPRouteModelFromMap(m map[string]interface{}) *models.StaticIPRoute {
	gateway := m["gateway"].(string)
	outputPort := m["output_port"].(string)
	prefix := m["prefix"].(string)
	var probeConfig *models.ProbeConfig // ProbeConfig
	probeConfigInterface, probeConfigIsSet := m["probe_config"]
	if probeConfigIsSet && probeConfigInterface != nil {
		probeConfigMap := probeConfigInterface.([]interface{})
		if len(probeConfigMap) > 0 {
			probeConfig = ProbeConfigModelFromMap(probeConfigMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.StaticIPRoute{
		Gateway:     gateway,
		OutputPort:  outputPort,
		Prefix:      prefix,
		ProbeConfig: probeConfig,
	}
}

func SetStaticIPRouteResourceData(d *schema.ResourceData, m *models.StaticIPRoute) {
	d.Set("gateway", m.Gateway)
	d.Set("output_port", m.OutputPort)
	d.Set("prefix", m.Prefix)
	d.Set("probe_config", SetProbeConfigSubResourceData([]*models.ProbeConfig{m.ProbeConfig}))
}

func SetStaticIPRouteSubResourceData(m []*models.StaticIPRoute) (d []*map[string]interface{}) {
	for _, StaticIPRouteModel := range m {
		if StaticIPRouteModel != nil {
			properties := make(map[string]interface{})
			properties["gateway"] = StaticIPRouteModel.Gateway
			properties["output_port"] = StaticIPRouteModel.OutputPort
			properties["prefix"] = StaticIPRouteModel.Prefix
			properties["probe_config"] = SetProbeConfigSubResourceData([]*models.ProbeConfig{StaticIPRouteModel.ProbeConfig})
			d = append(d, &properties)
		}
	}
	return
}

func StaticIPRouteSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"gateway": {
			Description: `Gateway IP`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"output_port": {
			Description: `Output Port`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"prefix": {
			Description: `IP Prefix`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"probe_config": {
			Description: `Probe Configuration`,
			Type:        schema.TypeList, //GoType: ProbeConfig
			Elem: &schema.Resource{
				Schema: ProbeConfigSchema(),
			},
			Optional: true,
		},
	}
}

func GetStaticIPRoutePropertyFields() (t []string) {
	return []string{
		"gateway",
		"output_port",
		"prefix",
		"probe_config",
	}
}
