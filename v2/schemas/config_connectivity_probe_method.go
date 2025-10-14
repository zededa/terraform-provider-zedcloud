package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ConfigConnectivityProbeMethodModel(d *schema.ResourceData) *models.ConfigConnectivityProbeMethod {
	configConnectivityProbeMethod, _ := d.Get("config_connectivity_probe_method").(models.ConfigConnectivityProbeMethod)
	return &configConnectivityProbeMethod
}

func ConfigConnectivityProbeMethodModelFromMap(m map[string]interface{}) *models.ConfigConnectivityProbeMethod {
	configConnectivityProbeMethod := m["config_connectivity_probe_method"].(models.ConfigConnectivityProbeMethod)
	return &configConnectivityProbeMethod
}

func SetConfigConnectivityProbeMethodResourceData(d *schema.ResourceData, m *models.ConfigConnectivityProbeMethod) {
}

func SetConfigConnectivityProbeMethodSubResourceData(m []*models.ConfigConnectivityProbeMethod) (d []*map[string]interface{}) {
	for _, ConfigConnectivityProbeMethodModel := range m {
		if ConfigConnectivityProbeMethodModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func ConfigConnectivityProbeMethodSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func GetConfigConnectivityProbeMethodPropertyFields() (t []string) {
	return []string{}
}
