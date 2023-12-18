package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func PluginEntitlementModel(d *schema.ResourceData) *models.PluginEntitlement {
	pluginTotalLimitInt, _ := d.Get("plugin_total_limit").(int)
	pluginTotalLimit := int32(pluginTotalLimitInt)
	pluginTotalUsageInt, _ := d.Get("plugin_total_usage").(int)
	pluginTotalUsage := int64(pluginTotalUsageInt)
	var pluginTypeEntitlement []*models.PluginTypeEntitlement // []*PluginTypeEntitlement
	pluginTypeEntitlementInterface, pluginTypeEntitlementIsSet := d.GetOk("plugin_type_entitlement")
	if pluginTypeEntitlementIsSet {
		var items []interface{}
		if listItems, isList := pluginTypeEntitlementInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = pluginTypeEntitlementInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := PluginTypeEntitlementModelFromMap(v.(map[string]interface{}))
			pluginTypeEntitlement = append(pluginTypeEntitlement, m)
		}
	}
	return &models.PluginEntitlement{
		PluginTotalLimit:      pluginTotalLimit,
		PluginTotalUsage:      pluginTotalUsage,
		PluginTypeEntitlement: pluginTypeEntitlement,
	}
}

func PluginEntitlementModelFromMap(m map[string]interface{}) *models.PluginEntitlement {
	pluginTotalLimit := int32(m["plugin_total_limit"].(int))  // int32
	pluginTotalUsage := int64(m["plugin_total_usage"].(int))  // int64
	var pluginTypeEntitlement []*models.PluginTypeEntitlement // []*PluginTypeEntitlement
	pluginTypeEntitlementInterface, pluginTypeEntitlementIsSet := m["plugin_type_entitlement"]
	if pluginTypeEntitlementIsSet {
		var items []interface{}
		if listItems, isList := pluginTypeEntitlementInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = pluginTypeEntitlementInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := PluginTypeEntitlementModelFromMap(v.(map[string]interface{}))
			pluginTypeEntitlement = append(pluginTypeEntitlement, m)
		}
	}
	return &models.PluginEntitlement{
		PluginTotalLimit:      pluginTotalLimit,
		PluginTotalUsage:      pluginTotalUsage,
		PluginTypeEntitlement: pluginTypeEntitlement,
	}
}

func SetPluginEntitlementResourceData(d *schema.ResourceData, m *models.PluginEntitlement) {
	d.Set("plugin_total_limit", m.PluginTotalLimit)
	d.Set("plugin_total_usage", m.PluginTotalUsage)
	d.Set("plugin_type_entitlement", SetPluginTypeEntitlementSubResourceData(m.PluginTypeEntitlement))
}

func SetPluginEntitlementSubResourceData(m []*models.PluginEntitlement) (d []*map[string]interface{}) {
	for _, PluginEntitlementModel := range m {
		if PluginEntitlementModel != nil {
			properties := make(map[string]interface{})
			properties["plugin_total_limit"] = PluginEntitlementModel.PluginTotalLimit
			properties["plugin_total_usage"] = PluginEntitlementModel.PluginTotalUsage
			properties["plugin_type_entitlement"] = SetPluginTypeEntitlementSubResourceData(PluginEntitlementModel.PluginTypeEntitlement)
			d = append(d, &properties)
		}
	}
	return
}

func PluginEntitlementSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"plugin_total_limit": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"plugin_total_usage": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"plugin_type_entitlement": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []*PluginTypeEntitlement
			Elem: &schema.Resource{
				Schema: PluginTypeEntitlementSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
	}
}

func GetPluginEntitlementPropertyFields() (t []string) {
	return []string{
		"plugin_total_limit",
		"plugin_total_usage",
		"plugin_type_entitlement",
	}
}
