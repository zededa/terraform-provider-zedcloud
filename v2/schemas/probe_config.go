package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ProbeConfigModel(d *schema.ResourceData) *models.ProbeConfig {
	var customProbeConfig *models.ConfigConnectivityProbe // ConfigConnectivityProbe
	customProbeConfigInterface, customProbeConfigIsSet := d.GetOk("custom_probe_config")
	if customProbeConfigIsSet && customProbeConfigInterface != nil {
		customProbeConfigMap := customProbeConfigInterface.([]interface{})
		if len(customProbeConfigMap) > 0 {
			customProbeConfig = ConfigConnectivityProbeModelFromMap(customProbeConfigMap[0].(map[string]interface{}))
		}
	}
	enableGatewayPing, _ := d.Get("enable_gateway_ping").(bool)
	pingMaxCostInt, _ := d.Get("ping_max_cost").(int)
	pingMaxCost := int64(pingMaxCostInt)
	preferLowerCost, _ := d.Get("prefer_lower_cost").(bool)
	preferStrongerWwanSignal, _ := d.Get("prefer_stronger_wwan_signal").(bool)
	return &models.ProbeConfig{
		CustomProbeConfig:        customProbeConfig,
		EnableGatewayPing:        enableGatewayPing,
		PingMaxCost:              pingMaxCost,
		PreferLowerCost:          preferLowerCost,
		PreferStrongerWwanSignal: preferStrongerWwanSignal,
	}
}

func ProbeConfigModelFromMap(m map[string]interface{}) *models.ProbeConfig {
	var customProbeConfig *models.ConfigConnectivityProbe // ConfigConnectivityProbe
	customProbeConfigInterface, customProbeConfigIsSet := m["custom_probe_config"]
	if customProbeConfigIsSet && customProbeConfigInterface != nil {
		customProbeConfigMap := customProbeConfigInterface.([]interface{})
		if len(customProbeConfigMap) > 0 {
			customProbeConfig = ConfigConnectivityProbeModelFromMap(customProbeConfigMap[0].(map[string]interface{}))
		}
	}
	//
	enableGatewayPing := m["enable_gateway_ping"].(bool)
	pingMaxCost := int64(m["ping_max_cost"].(int)) // int64
	preferLowerCost := m["prefer_lower_cost"].(bool)
	preferStrongerWwanSignal := m["prefer_stronger_wwan_signal"].(bool)
	return &models.ProbeConfig{
		CustomProbeConfig:        customProbeConfig,
		EnableGatewayPing:        enableGatewayPing,
		PingMaxCost:              pingMaxCost,
		PreferLowerCost:          preferLowerCost,
		PreferStrongerWwanSignal: preferStrongerWwanSignal,
	}
}

func SetProbeConfigResourceData(d *schema.ResourceData, m *models.ProbeConfig) {
	d.Set("custom_probe_config", SetConfigConnectivityProbeSubResourceData([]*models.ConfigConnectivityProbe{m.CustomProbeConfig}))
	d.Set("enable_gateway_ping", m.EnableGatewayPing)
	d.Set("ping_max_cost", m.PingMaxCost)
	d.Set("prefer_lower_cost", m.PreferLowerCost)
	d.Set("prefer_stronger_wwan_signal", m.PreferStrongerWwanSignal)
}

func SetProbeConfigSubResourceData(m []*models.ProbeConfig) (d []*map[string]interface{}) {
	for _, ProbeConfigModel := range m {
		if ProbeConfigModel != nil {
			properties := make(map[string]interface{})
			properties["custom_probe_config"] = SetConfigConnectivityProbeSubResourceData([]*models.ConfigConnectivityProbe{ProbeConfigModel.CustomProbeConfig})
			properties["enable_gateway_ping"] = ProbeConfigModel.EnableGatewayPing
			properties["ping_max_cost"] = ProbeConfigModel.PingMaxCost
			properties["prefer_lower_cost"] = ProbeConfigModel.PreferLowerCost
			properties["prefer_stronger_wwan_signal"] = ProbeConfigModel.PreferStrongerWwanSignal
			d = append(d, &properties)
		}
	}
	return
}

func ProbeConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"custom_probe_config": {
			Description: `Custom probe configuration`,
			Type:        schema.TypeList, //GoType: ConfigConnectivityProbe
			Elem: &schema.Resource{
				Schema: ConfigConnectivityProbeSchema(),
			},
			Optional: true,
		},

		"enable_gateway_ping": {
			Description: `Enable gateway ping`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"ping_max_cost": {
			Description: `Ping max cost`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"prefer_lower_cost": {
			Description: `Prefer lower cost`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"prefer_stronger_wwan_signal": {
			Description: `Prefer stronger WWAN signal`,
			Type:        schema.TypeBool,
			Optional:    true,
		},
	}
}

func GetProbeConfigPropertyFields() (t []string) {
	return []string{
		"custom_probe_config",
		"enable_gateway_ping",
		"ping_max_cost",
		"prefer_lower_cost",
		"prefer_stronger_wwan_signal",
	}
}
