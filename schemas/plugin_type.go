package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func PluginTypeModel(d *schema.ResourceData) *models.PluginType {
	pluginType, _ := d.Get("plugin_type").(models.PluginType)
	return &pluginType
}

func PluginTypeModelFromMap(m map[string]interface{}) *models.PluginType {
	pluginType := m["plugin_type"].(models.PluginType)
	return &pluginType
}

func SetPluginTypeResourceData(d *schema.ResourceData, m *models.PluginType) {
}

func SetPluginTypeSubResourceData(m []*models.PluginType) (d []*map[string]interface{}) {
	for _, PluginTypeModel := range m {
		if PluginTypeModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func PluginTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func GetPluginTypePropertyFields() (t []string) {
	return []string{}
}
