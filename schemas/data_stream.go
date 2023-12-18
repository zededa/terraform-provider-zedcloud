package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func DataStreamModel(d *schema.ResourceData) *models.DataStream {
	enabled, _ := d.Get("enabled").(bool)
	pluginID, _ := d.Get("plugin_id").(string)
	return &models.DataStream{
		Enabled:  enabled,
		PluginID: pluginID,
	}
}

func DataStreamModelFromMap(m map[string]interface{}) *models.DataStream {
	enabled := m["enabled"].(bool)
	pluginID := m["plugin_id"].(string)
	return &models.DataStream{
		Enabled:  enabled,
		PluginID: pluginID,
	}
}

func SetDataStreamResourceData(d *schema.ResourceData, m *models.DataStream) {
	d.Set("enabled", m.Enabled)
	d.Set("plugin_id", m.PluginID)
}

func SetDataStreamSubResourceData(m []*models.DataStream) (d []*map[string]interface{}) {
	for _, DataStreamModel := range m {
		if DataStreamModel != nil {
			properties := make(map[string]interface{})
			properties["enabled"] = DataStreamModel.Enabled
			properties["plugin_id"] = DataStreamModel.PluginID
			d = append(d, &properties)
		}
	}
	return
}

func DataStreamSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"enabled": {
			Description: ``,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"plugin_id": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetDataStreamPropertyFields() (t []string) {
	return []string{
		"enabled",
		"plugin_id",
	}
}
