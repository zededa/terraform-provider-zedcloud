package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func EntitlementModel(d *schema.ResourceData) *models.Entitlement {
	var devices []*models.DeviceEntitlement // []*DeviceEntitlement
	devicesInterface, devicesIsSet := d.GetOk("devices")
	if devicesIsSet {
		var items []interface{}
		if listItems, isList := devicesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = devicesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := DeviceEntitlementModelFromMap(v.(map[string]interface{}))
			devices = append(devices, m)
		}
	}
	entpID, _ := d.Get("entp_id").(string)
	entpName, _ := d.Get("entp_name").(string)
	var plugins *models.PluginEntitlement // PluginEntitlement
	pluginsInterface, pluginsIsSet := d.GetOk("plugins")
	if pluginsIsSet && pluginsInterface != nil {
		pluginsMap := pluginsInterface.([]interface{})
		if len(pluginsMap) > 0 {
			plugins = PluginEntitlementModelFromMap(pluginsMap[0].(map[string]interface{}))
		}
	}
	var userCount *models.UserEntitlement // UserEntitlement
	userCountInterface, userCountIsSet := d.GetOk("user_count")
	if userCountIsSet && userCountInterface != nil {
		userCountMap := userCountInterface.([]interface{})
		if len(userCountMap) > 0 {
			userCount = UserEntitlementModelFromMap(userCountMap[0].(map[string]interface{}))
		}
	}
	return &models.Entitlement{
		Devices:   devices,
		EntpID:    entpID,
		EntpName:  entpName,
		Plugins:   plugins,
		UserCount: userCount,
	}
}

func EntitlementModelFromMap(m map[string]interface{}) *models.Entitlement {
	var devices []*models.DeviceEntitlement // []*DeviceEntitlement
	devicesInterface, devicesIsSet := m["devices"]
	if devicesIsSet {
		var items []interface{}
		if listItems, isList := devicesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = devicesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := DeviceEntitlementModelFromMap(v.(map[string]interface{}))
			devices = append(devices, m)
		}
	}
	entpID := m["entp_id"].(string)
	entpName := m["entp_name"].(string)
	var plugins *models.PluginEntitlement // PluginEntitlement
	pluginsInterface, pluginsIsSet := m["plugins"]
	if pluginsIsSet && pluginsInterface != nil {
		pluginsMap := pluginsInterface.([]interface{})
		if len(pluginsMap) > 0 {
			plugins = PluginEntitlementModelFromMap(pluginsMap[0].(map[string]interface{}))
		}
	}
	//
	var userCount *models.UserEntitlement // UserEntitlement
	userCountInterface, userCountIsSet := m["user_count"]
	if userCountIsSet && userCountInterface != nil {
		userCountMap := userCountInterface.([]interface{})
		if len(userCountMap) > 0 {
			userCount = UserEntitlementModelFromMap(userCountMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.Entitlement{
		Devices:   devices,
		EntpID:    entpID,
		EntpName:  entpName,
		Plugins:   plugins,
		UserCount: userCount,
	}
}

func SetEntitlementResourceData(d *schema.ResourceData, m *models.Entitlement) {
	d.Set("devices", SetDeviceEntitlementSubResourceData(m.Devices))
	d.Set("entp_id", m.EntpID)
	d.Set("entp_name", m.EntpName)
	d.Set("plugins", SetPluginEntitlementSubResourceData([]*models.PluginEntitlement{m.Plugins}))
	d.Set("user_count", SetUserEntitlementSubResourceData([]*models.UserEntitlement{m.UserCount}))
}

func SetEntitlementSubResourceData(m []*models.Entitlement) (d []*map[string]interface{}) {
	for _, EntitlementModel := range m {
		if EntitlementModel != nil {
			properties := make(map[string]interface{})
			properties["devices"] = SetDeviceEntitlementSubResourceData(EntitlementModel.Devices)
			properties["entp_id"] = EntitlementModel.EntpID
			properties["entp_name"] = EntitlementModel.EntpName
			properties["plugins"] = SetPluginEntitlementSubResourceData([]*models.PluginEntitlement{EntitlementModel.Plugins})
			properties["user_count"] = SetUserEntitlementSubResourceData([]*models.UserEntitlement{EntitlementModel.UserCount})
			d = append(d, &properties)
		}
	}
	return
}

func EntitlementSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"devices": {
			Description: `Device entitlement data`,
			Type:        schema.TypeList, //GoType: []*DeviceEntitlement
			Elem: &schema.Resource{
				Schema: DeviceEntitlementSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"entp_id": {
			Description: `Enterprise id for which we want to post/get enforce entitlement`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"entp_name": {
			Description: `Enterprise name for which we want to post/get enforce entitlement`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"plugins": {
			Description: `Plugin entitlement data`,
			Type:        schema.TypeList, //GoType: PluginEntitlement
			Elem: &schema.Resource{
				Schema: PluginEntitlementSchema(),
			},
			Optional: true,
		},

		"user_count": {
			Description: `User entitlement data`,
			Type:        schema.TypeList, //GoType: UserEntitlement
			Elem: &schema.Resource{
				Schema: UserEntitlementSchema(),
			},
			Optional: true,
		},
	}
}

func GetEntitlementPropertyFields() (t []string) {
	return []string{
		"devices",
		"entp_id",
		"entp_name",
		"plugins",
		"user_count",
	}
}
