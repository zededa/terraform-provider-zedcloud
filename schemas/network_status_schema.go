package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate NetworkStatus resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func NetworkStatusModel(d *schema.ResourceData) *models.NetworkStatus {
	defaultRouters := d.Get("default_routers").([]string)
	var dns *models.DNSInfo // DNSInfo
	dnsInterface, dnsIsSet := d.GetOk("dns")
	if dnsIsSet {
		dnsMap := dnsInterface.([]interface{})[0].(map[string]interface{})
		dns = DNSInfoModelFromMap(dnsMap)
	}
	var errInfo *models.DeviceError // DeviceError
	errInfoInterface, errInfoIsSet := d.GetOk("err_info")
	if errInfoIsSet {
		errInfoMap := errInfoInterface.([]interface{})[0].(map[string]interface{})
		errInfo = DeviceErrorModelFromMap(errInfoMap)
	}
	var gpsLocation *models.GPSLocation // GPSLocation
	gpsLocationInterface, gpsLocationIsSet := d.GetOk("gps_location")
	if gpsLocationIsSet {
		gpsLocationMap := gpsLocationInterface.([]interface{})[0].(map[string]interface{})
		gpsLocation = GPSLocationModelFromMap(gpsLocationMap)
	}
	ifName := d.Get("if_name").(string)
	iPAddrs := d.Get("ip_addrs").([]string)
	var location *models.GeoLocation // GeoLocation
	locationInterface, locationIsSet := d.GetOk("location")
	if locationIsSet {
		locationMap := locationInterface.([]interface{})[0].(map[string]interface{})
		location = GeoLocationModelFromMap(locationMap)
	}
	macAddr := d.Get("mac_addr").(string)
	var proxy *models.NetProxyStatus // NetProxyStatus
	proxyInterface, proxyIsSet := d.GetOk("proxy")
	if proxyIsSet {
		proxyMap := proxyInterface.([]interface{})[0].(map[string]interface{})
		proxy = NetProxyStatusModelFromMap(proxyMap)
	}
	up := d.Get("up").(bool)
	uplink := d.Get("uplink").(bool)
	return &models.NetworkStatus{
		DefaultRouters: defaultRouters,
		DNS:            dns,
		ErrInfo:        errInfo,
		GpsLocation:    gpsLocation,
		IfName:         ifName,
		IPAddrs:        iPAddrs,
		Location:       location,
		MacAddr:        macAddr,
		Proxy:          proxy,
		Up:             up,
		Uplink:         uplink,
	}
}

func NetworkStatusModelFromMap(m map[string]interface{}) *models.NetworkStatus {
	defaultRouters := m["default_routers"].([]string)
	var dns *models.DNSInfo // DNSInfo
	dnsInterface, dnsIsSet := m["dns"]
	if dnsIsSet {
		dnsMap := dnsInterface.([]interface{})[0].(map[string]interface{})
		dns = DNSInfoModelFromMap(dnsMap)
	}
	//
	var errInfo *models.DeviceError // DeviceError
	errInfoInterface, errInfoIsSet := m["err_info"]
	if errInfoIsSet {
		errInfoMap := errInfoInterface.([]interface{})[0].(map[string]interface{})
		errInfo = DeviceErrorModelFromMap(errInfoMap)
	}
	//
	var gpsLocation *models.GPSLocation // GPSLocation
	gpsLocationInterface, gpsLocationIsSet := m["gps_location"]
	if gpsLocationIsSet {
		gpsLocationMap := gpsLocationInterface.([]interface{})[0].(map[string]interface{})
		gpsLocation = GPSLocationModelFromMap(gpsLocationMap)
	}
	//
	ifName := m["if_name"].(string)
	iPAddrs := m["ip_addrs"].([]string)
	var location *models.GeoLocation // GeoLocation
	locationInterface, locationIsSet := m["location"]
	if locationIsSet {
		locationMap := locationInterface.([]interface{})[0].(map[string]interface{})
		location = GeoLocationModelFromMap(locationMap)
	}
	//
	macAddr := m["mac_addr"].(string)
	var proxy *models.NetProxyStatus // NetProxyStatus
	proxyInterface, proxyIsSet := m["proxy"]
	if proxyIsSet {
		proxyMap := proxyInterface.([]interface{})[0].(map[string]interface{})
		proxy = NetProxyStatusModelFromMap(proxyMap)
	}
	//
	up := m["up"].(bool)
	uplink := m["uplink"].(bool)
	return &models.NetworkStatus{
		DefaultRouters: defaultRouters,
		DNS:            dns,
		ErrInfo:        errInfo,
		GpsLocation:    gpsLocation,
		IfName:         ifName,
		IPAddrs:        iPAddrs,
		Location:       location,
		MacAddr:        macAddr,
		Proxy:          proxy,
		Up:             up,
		Uplink:         uplink,
	}
}

// Update the underlying NetworkStatus resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetNetworkStatusResourceData(d *schema.ResourceData, m *models.NetworkStatus) {
	d.Set("default_routers", m.DefaultRouters)
	d.Set("dns", SetDNSInfoSubResourceData([]*models.DNSInfo{m.DNS}))
	d.Set("err_info", SetDeviceErrorSubResourceData([]*models.DeviceError{m.ErrInfo}))
	d.Set("gps_location", SetGPSLocationSubResourceData([]*models.GPSLocation{m.GpsLocation}))
	d.Set("if_name", m.IfName)
	d.Set("ip_addrs", m.IPAddrs)
	d.Set("location", SetGeoLocationSubResourceData([]*models.GeoLocation{m.Location}))
	d.Set("mac_addr", m.MacAddr)
	d.Set("proxy", SetNetProxyStatusSubResourceData([]*models.NetProxyStatus{m.Proxy}))
	d.Set("up", m.Up)
	d.Set("uplink", m.Uplink)
}

// Iterate throught and update the NetworkStatus resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetNetworkStatusSubResourceData(m []*models.NetworkStatus) (d []*map[string]interface{}) {
	for _, NetworkStatusModel := range m {
		if NetworkStatusModel != nil {
			properties := make(map[string]interface{})
			properties["default_routers"] = NetworkStatusModel.DefaultRouters
			properties["dns"] = SetDNSInfoSubResourceData([]*models.DNSInfo{NetworkStatusModel.DNS})
			properties["err_info"] = SetDeviceErrorSubResourceData([]*models.DeviceError{NetworkStatusModel.ErrInfo})
			properties["gps_location"] = SetGPSLocationSubResourceData([]*models.GPSLocation{NetworkStatusModel.GpsLocation})
			properties["if_name"] = NetworkStatusModel.IfName
			properties["ip_addrs"] = NetworkStatusModel.IPAddrs
			properties["location"] = SetGeoLocationSubResourceData([]*models.GeoLocation{NetworkStatusModel.Location})
			properties["mac_addr"] = NetworkStatusModel.MacAddr
			properties["proxy"] = SetNetProxyStatusSubResourceData([]*models.NetProxyStatus{NetworkStatusModel.Proxy})
			properties["up"] = NetworkStatusModel.Up
			properties["uplink"] = NetworkStatusModel.Uplink
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the NetworkStatus resource defined in the Terraform configuration
func NetworkStatusSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"default_routers": {
			Description: `Default Routers`,
			Type:        schema.TypeList, //GoType: []string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"dns": {
			Description: `DNS Configuration`,
			Type:        schema.TypeList, //GoType: DNSInfo
			Elem: &schema.Resource{
				Schema: DNSInfoSchema(),
			},
			Optional: true,
		},

		"err_info": {
			Description: `Network error details`,
			Type:        schema.TypeList, //GoType: DeviceError
			Elem: &schema.Resource{
				Schema: DeviceErrorSchema(),
			},
			Optional: true,
		},

		"gps_location": {
			Description: `Location from GNSS receivers on WWAN type adapters`,
			Type:        schema.TypeList, //GoType: GPSLocation
			Elem: &schema.Resource{
				Schema: GPSLocationSchema(),
			},
			Optional: true,
		},

		"if_name": {
			Description: `ifName`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"ip_addrs": {
			Description: `Array of IP addresses`,
			Type:        schema.TypeList, //GoType: []string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"location": {
			Description: `Geo Location Details`,
			Type:        schema.TypeList, //GoType: GeoLocation
			Elem: &schema.Resource{
				Schema: GeoLocationSchema(),
			},
			Optional: true,
		},

		"mac_addr": {
			Description: `mac Address`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"proxy": {
			Description: `Network Proxy status`,
			Type:        schema.TypeList, //GoType: NetProxyStatus
			Elem: &schema.Resource{
				Schema: NetProxyStatusSchema(),
			},
			Optional: true,
		},

		"up": {
			Description: `Network Status flag`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"uplink": {
			Description: `Uplink flag`,
			Type:        schema.TypeBool,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the NetworkStatus resource
func GetNetworkStatusPropertyFields() (t []string) {
	return []string{
		"default_routers",
		"dns",
		"err_info",
		"gps_location",
		"if_name",
		"ip_addrs",
		"location",
		"mac_addr",
		"proxy",
		"up",
		"uplink",
	}
}
