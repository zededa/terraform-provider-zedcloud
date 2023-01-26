package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func NetworkWirelessModel(d *schema.ResourceData) *models.NetWirelessConfig {
	var cellularCfg *models.NetCellularConfig // NetCellularConfig
	cellularCfgInterface, cellularCfgIsSet := d.GetOk("cellular_cfg")
	if cellularCfgIsSet && cellularCfgInterface != nil {
		cellularCfgMap := cellularCfgInterface.([]interface{})
		if len(cellularCfgMap) > 0 {
			cellularCfg = NetCellularConfigModelFromMap(cellularCfgMap[0].(map[string]interface{}))
		}
	}
	var typeVar *models.NetworkWirelessType // NetworkWirelessType
	typeInterface, typeIsSet := d.GetOk("type")
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewNetworkWirelessType(models.NetworkWirelessType(typeModel))
	}
	var wifiCfg *models.NetWifiConfig // NetWifiConfig
	wifiCfgInterface, wifiCfgIsSet := d.GetOk("wifi_cfg")
	if wifiCfgIsSet && wifiCfgInterface != nil {
		wifiCfgMap := wifiCfgInterface.([]interface{})
		if len(wifiCfgMap) > 0 {
			wifiCfg = NetWifiConfigModelFromMap(wifiCfgMap[0].(map[string]interface{}))
		}
	}
	return &models.NetWirelessConfig{
		CellularCfg: cellularCfg,
		Type:        typeVar,
		WifiCfg:     wifiCfg,
	}
}

func NetworkWirelessModelFromMap(m map[string]interface{}) *models.NetWirelessConfig {
	var cellularCfg *models.NetCellularConfig // NetCellularConfig
	cellularCfgInterface, cellularCfgIsSet := m["cellular_cfg"]
	if cellularCfgIsSet && cellularCfgInterface != nil {
		cellularCfgMap := cellularCfgInterface.([]interface{})
		if len(cellularCfgMap) > 0 {
			cellularCfg = NetCellularConfigModelFromMap(cellularCfgMap[0].(map[string]interface{}))
		}
	}
	//
	var typeVar *models.NetworkWirelessType // NetworkWirelessType
	typeInterface, typeIsSet := m["type"]
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewNetworkWirelessType(models.NetworkWirelessType(typeModel))
	}
	var wifiCfg *models.NetWifiConfig // NetWifiConfig
	wifiCfgInterface, wifiCfgIsSet := m["wifi_cfg"]
	if wifiCfgIsSet && wifiCfgInterface != nil {
		wifiCfgMap := wifiCfgInterface.([]interface{})
		if len(wifiCfgMap) > 0 {
			wifiCfg = NetWifiConfigModelFromMap(wifiCfgMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.NetWirelessConfig{
		CellularCfg: cellularCfg,
		Type:        typeVar,
		WifiCfg:     wifiCfg,
	}
}

func SetNetworkWirelessResourceData(d *schema.ResourceData, m *models.NetWirelessConfig) {
	d.Set("cellular_cfg", SetNetCellularConfigSubResourceData([]*models.NetCellularConfig{m.CellularCfg}))
	d.Set("type", m.Type)
	d.Set("wifi_cfg", SetNetWifiConfigSubResourceData([]*models.NetWifiConfig{m.WifiCfg}))
}

// Iterate through and update the NetWirelessConfig resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetNetWirelessConfigSubResourceData(m []*models.NetWirelessConfig) (d []*map[string]interface{}) {
	for _, NetWirelessConfigModel := range m {
		if NetWirelessConfigModel != nil {
			properties := make(map[string]interface{})
			properties["cellular_cfg"] = SetNetCellularConfigSubResourceData([]*models.NetCellularConfig{NetWirelessConfigModel.CellularCfg})
			properties["type"] = NetWirelessConfigModel.Type
			properties["wifi_cfg"] = SetNetWifiConfigSubResourceData([]*models.NetWifiConfig{NetWirelessConfigModel.WifiCfg})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the NetWirelessConfig resource defined in the Terraform configuration
func NetworkWireless() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cellular_cfg": {
			Description: "Cellular configuration",
			Type:        schema.TypeList, //GoType: NetCellularConfig
			Elem: &schema.Resource{
				Schema: NetCellularConfigSchema(),
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

		"wifi_cfg": {
			Description: "Can be multiple APs on a single wlan, e.g. one for 2.5Ghz, other 5Ghz SSIDs",
			Type:        schema.TypeList, //GoType: NetWifiConfig
			Elem: &schema.Resource{
				Schema: NetWifiConfigSchema(),
			},
			Optional: true,
		},
	}
}

// Retrieve property field names for updating the NetWirelessConfig resource
func GetNetworkWirelessPropertyFields() (t []string) {
	return []string{
		"cellular_cfg",
		"type",
		"wifi_cfg",
	}
}
