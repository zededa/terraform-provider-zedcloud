package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

func NetworkWirelessModel(d *schema.ResourceData) *models.Wireless {
	var cellularCfg *models.Cellular // NetCellularConfig
	cellularCfgInterface, cellularCfgIsSet := d.GetOk("cellular")
	if cellularCfgIsSet && cellularCfgInterface != nil {
		cellularCfgMap := cellularCfgInterface.([]interface{})
		if len(cellularCfgMap) > 0 {
			cellularCfg = NetworkCellularModelFromMap(cellularCfgMap[0].(map[string]interface{}))
		}
	}
	var typeVar *models.Type // NetworkWirelessType
	typeInterface, typeIsSet := d.GetOk("type")
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewNetworkWirelessType(models.Type(typeModel))
	}
	var wifiCfg *models.Wifi // NetWifiConfig
	wifiCfgInterface, wifiCfgIsSet := d.GetOk("wifi")
	if wifiCfgIsSet && wifiCfgInterface != nil {
		wifiCfgMap := wifiCfgInterface.([]interface{})
		if len(wifiCfgMap) > 0 {
			wifiCfg = NetworkWifiModelFromMap(wifiCfgMap[0].(map[string]interface{}))
		}
	}
	return &models.Wireless{
		CellularCfg: cellularCfg,
		Type:        typeVar,
		WifiCfg:     wifiCfg,
	}
}

func NetworkWirelessModelFromMap(m map[string]interface{}) *models.Wireless {
	var cellularCfg *models.Cellular // NetCellularConfig
	cellularCfgInterface, cellularCfgIsSet := m["cellular"]
	if cellularCfgIsSet && cellularCfgInterface != nil {
		cellularCfgMap := cellularCfgInterface.([]interface{})
		if len(cellularCfgMap) > 0 {
			cellularCfg = NetworkCellularModelFromMap(cellularCfgMap[0].(map[string]interface{}))
		}
	}
	var typeVar *models.Type // NetworkWirelessType
	typeInterface, typeIsSet := m["type"]
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewNetworkWirelessType(models.Type(typeModel))
	}
	var wifiCfg *models.Wifi // NetWifiConfig
	wifiCfgInterface, wifiCfgIsSet := m["wifi"]
	if wifiCfgIsSet && wifiCfgInterface != nil {
		wifiCfgMap := wifiCfgInterface.([]interface{})
		if len(wifiCfgMap) > 0 {
			wifiCfg = NetworkWifiModelFromMap(wifiCfgMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.Wireless{
		CellularCfg: cellularCfg,
		Type:        typeVar,
		WifiCfg:     wifiCfg,
	}
}

func SetNetworkWirelessResourceData(d *schema.ResourceData, m *models.Wireless) {
	d.Set("cellular", SetNetworkCellularSubResourceData([]*models.Cellular{m.CellularCfg}))
	d.Set("type", m.Type)
	d.Set("wifi", SetNetworkWifiSubResourceData([]*models.Wifi{m.WifiCfg}))
}

func SetNetWirelessConfigSubResourceData(m []*models.Wireless) (d []*map[string]interface{}) {
	for _, NetWirelessConfigModel := range m {
		if NetWirelessConfigModel != nil {
			properties := make(map[string]interface{})
			properties["cellular"] = SetNetworkCellularSubResourceData([]*models.Cellular{NetWirelessConfigModel.CellularCfg})
			properties["type"] = NetWirelessConfigModel.Type
			properties["wifi"] = SetNetworkWifiSubResourceData([]*models.Wifi{NetWirelessConfigModel.WifiCfg})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the NetWirelessConfig resource defined in the Terraform configuration
func NetworkWireless() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cellular": {
			Description: "Cellular configuration",
			Type:        schema.TypeList, //GoType: NetCellularConfig
			Elem: &schema.Resource{
				Schema: NetworkCellular(),
			},
			Optional: true,
		},

		"type": {
			Description: `Type of Wireless Network:
NETWORK_WIRELESS_TYPE_WIFI
NETWORK_WIRELESS_TYPE_CELLULAR`,
			Type:     schema.TypeString,
			Optional: true,
		},

		"wifi": {
			Description: "Can be multiple APs on a single wlan, e.g. one for 2.5Ghz, other 5Ghz SSIDs",
			Type:        schema.TypeList, //GoType: NetWifiConfig
			Elem: &schema.Resource{
				Schema: NetworkWifi(),
			},
			Optional: true,
		},
	}
}

// Retrieve property field names for updating the NetWirelessConfig resource
func GetNetworkWirelessPropertyFields() (t []string) {
	return []string{
		"cellular",
		"type",
		"wifi",
	}
}
