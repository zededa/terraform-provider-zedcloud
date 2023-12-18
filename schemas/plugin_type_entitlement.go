package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func PluginTypeEntitlementModel(d *schema.ResourceData) *models.PluginTypeEntitlement {
	var pluginType *models.PluginType // PluginType
	pluginTypeInterface, pluginTypeIsSet := d.GetOk("plugin_type")
	if pluginTypeIsSet {
		pluginTypeModel := pluginTypeInterface.(string)
		pluginType = models.NewPluginType(models.PluginType(pluginTypeModel))
	}
	pluginTypeLimitInt, _ := d.Get("plugin_type_limit").(int)
	pluginTypeLimit := int32(pluginTypeLimitInt)
	pluginTypeUsageInt, _ := d.Get("plugin_type_usage").(int)
	pluginTypeUsage := int64(pluginTypeUsageInt)
	return &models.PluginTypeEntitlement{
		PluginType:      pluginType,
		PluginTypeLimit: pluginTypeLimit,
		PluginTypeUsage: pluginTypeUsage,
	}
}

func PluginTypeEntitlementModelFromMap(m map[string]interface{}) *models.PluginTypeEntitlement {
	var pluginType *models.PluginType // PluginType
	pluginTypeInterface, pluginTypeIsSet := m["plugin_type"]
	if pluginTypeIsSet {
		pluginTypeModel := pluginTypeInterface.(string)
		pluginType = models.NewPluginType(models.PluginType(pluginTypeModel))
	}
	pluginTypeLimit := int32(m["plugin_type_limit"].(int)) // int32
	pluginTypeUsage := int64(m["plugin_type_usage"].(int)) // int64
	return &models.PluginTypeEntitlement{
		PluginType:      pluginType,
		PluginTypeLimit: pluginTypeLimit,
		PluginTypeUsage: pluginTypeUsage,
	}
}

func SetPluginTypeEntitlementResourceData(d *schema.ResourceData, m *models.PluginTypeEntitlement) {
	d.Set("plugin_type", m.PluginType)
	d.Set("plugin_type_limit", m.PluginTypeLimit)
	d.Set("plugin_type_usage", m.PluginTypeUsage)
}

func SetPluginTypeEntitlementSubResourceData(m []*models.PluginTypeEntitlement) (d []*map[string]interface{}) {
	for _, PluginTypeEntitlementModel := range m {
		if PluginTypeEntitlementModel != nil {
			properties := make(map[string]interface{})
			properties["plugin_type"] = PluginTypeEntitlementModel.PluginType
			properties["plugin_type_limit"] = PluginTypeEntitlementModel.PluginTypeLimit
			properties["plugin_type_usage"] = PluginTypeEntitlementModel.PluginTypeUsage
			d = append(d, &properties)
		}
	}
	return
}

func PluginTypeEntitlementSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"plugin_type": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"plugin_type_limit": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"plugin_type_usage": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},
	}
}

func GetPluginTypeEntitlementPropertyFields() (t []string) {
	return []string{
		"plugin_type",
		"plugin_type_limit",
		"plugin_type_usage",
	}
}
