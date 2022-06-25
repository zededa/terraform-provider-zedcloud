// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"context"
	"fmt"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/schemas"
	zedcloudapi "github.com/zededa/zedcloud-api"
	"github.com/zededa/zedcloud-api/swagger_models"
	"log"
)

var networkUrlExtension = "networks"

var NetworkResourceSchema = &schema.Resource{
	CreateContext: createNetworkResource,
	ReadContext:   readResourceNetwork,
	UpdateContext: updateNetworkResource,
	DeleteContext: deleteNetworkResource,
	Schema:        schemas.NetworkSchema,
	Importer: &schema.ResourceImporter{
		StateContext: schema.ImportStatePassthroughContext,
	},
}

// The schema for a resource group data source
func getNetworkResourceSchema() *schema.Resource {
	return NetworkResourceSchema
}

func getNetworkUrl(name, id, urlType string) string {
	return zedcloudapi.UrlForObjectRequest(networkUrlExtension, name, id, urlType)
}

func rdMapNetCellularConfig(d map[string]interface{}) (interface{}, error) {
	return &swagger_models.NetCellularConfig{
		APN: rdEntryStr(d, "apn"),
	}, nil
}

func rdMapNetWifiConfigSecrets(d map[string]interface{}) (interface{}, error) {
	return &swagger_models.NetWifiConfigSecrets{
		WiFiPasswd: rdEntryStr(d, "wifi_passwd"),
	}, nil
}

func networkWiFiKeySchemePtr(strVal string) *swagger_models.NetworkWiFiKeyScheme {
	val := swagger_models.NetworkWiFiKeyScheme(strVal)
	return &val
}

func rdMapNetWifiConfig(d map[string]interface{}) (interface{}, error) {
	keyScheme := swagger_models.NetworkWiFiKeyScheme(rdEntryStr(d, "key_scheme"))
	val, err := rdEntryStructPtr(d, "secret", rdMapNetWifiConfigSecrets)
	if err != nil {
		return nil, err
	}
	var secret *swagger_models.NetWifiConfigSecrets
	if val != nil {
		secret = val.(*swagger_models.NetWifiConfigSecrets)
	}
	return &swagger_models.NetWifiConfig{
		Identity:  rdEntryStr(d, "identity"),
		KeyScheme: &keyScheme,
		Priority:  rdEntryInt32(d, "priority"),
		Secret:    secret,
		WifiSSID:  rdEntryStr(d, "wifi_ssid"),
	}, nil
}

func networkWirelessTypePtr(strVal string) *swagger_models.NetworkWirelessType {
	val := swagger_models.NetworkWirelessType(strVal)
	return &val
}

func rdMapNetWirelessConfig(d map[string]interface{}) (interface{}, error) {
	val, err := rdEntryStructPtr(d, "cellular_cfg", rdMapNetCellularConfig)
	if err != nil {
		return nil, err
	}
	var cellularCfg *swagger_models.NetCellularConfig
	if val != nil {
		cellularCfg = val.(*swagger_models.NetCellularConfig)
	}
	val, err = rdEntryStructPtr(d, "wifi_cfg", rdMapNetWifiConfig)
	if err != nil {
		return nil, err
	}
	var wifiCfg *swagger_models.NetWifiConfig
	if val != nil {
		wifiCfg = val.(*swagger_models.NetWifiConfig)
	}
	return &swagger_models.NetWirelessConfig{
		CellularCfg: cellularCfg,
		Type:        networkWirelessTypePtr(rdEntryStr(d, "type")),
		WifiCfg:     wifiCfg,
	}, nil
}
func rdMapNetWirelessConfigOrNil(d *schema.ResourceData) (*swagger_models.NetWirelessConfig, error) {
	val, err := rdEntryStructPtr(d, "wireless", rdMapNetWirelessConfig)
	if err != nil {
		return nil, err
	}
	if val == nil {
		return nil, nil
	}
	return val.(*swagger_models.NetWirelessConfig), nil
}

func networkProxyProtoPtr(strVal string) *swagger_models.NetworkProxyProto {
	val := swagger_models.NetworkProxyProto(strVal)
	return &val
}

func rdMapNetProxyServer(d map[string]interface{}, key string) (
	[]*swagger_models.NetProxyServer, error) {
	proxyList := make([]*swagger_models.NetProxyServer, 0)
	if val, exists := d[key]; exists {
		for _, v := range val.([]interface{}) {
			entry := v.(map[string]interface{})
			proxyList = append(proxyList, &swagger_models.NetProxyServer{
				Port:   rdEntryInt64(entry, "port"),
				Proto:  networkProxyProtoPtr(rdEntryStr(entry, "proto")),
				Server: rdEntryStr(entry, "server"),
			})
		}
	}
	return proxyList, nil
}

func rdMapNetProxyConfig(d map[string]interface{}) (interface{}, error) {
	certStrList := rdEntryStrList(d, "network_proxy_certs")
	certs := make([]strfmt.Base64, 0)
	for _, certStr := range certStrList {
		certs = append(certs, strfmt.Base64(certStr))
	}
	proxies, err := rdMapNetProxyServer(d, "proxy")
	if err != nil {
		return nil, err
	}
	return &swagger_models.NetProxyConfig{
		Exceptions:        rdEntryStr(d, "exceptions"),
		NetworkProxy:      rdEntryBool(d, "network_proxy"),
		NetworkProxyCerts: certs,
		NetworkProxyURL:   rdEntryStr(d, "network_proxy_url"),
		Pacfile:           rdEntryStr(d, "pacfile"),
		Proxies:           proxies,
	}, nil
}

func rdMapNetProxyConfigOrNil(d *schema.ResourceData) (*swagger_models.NetProxyConfig, error) {
	val, err := rdEntryStructPtr(d, "proxy", rdMapNetProxyConfig)
	if err != nil {
		return nil, err
	}
	if val == nil {
		return nil, nil
	}
	return val.(*swagger_models.NetProxyConfig), nil
}

func networkDHCPTypePtr(strVal string) *swagger_models.NetworkDHCPType {
	val := swagger_models.NetworkDHCPType(strVal)
	return &val
}

func rdMapIPSpec(d map[string]interface{}) (interface{}, error) {
	dhcpRange, err := rdMapDhcpIpRangeOrNil(d)
	if err != nil {
		return nil, err
	}
	return &swagger_models.IPSpec{
		Dhcp:      networkDHCPTypePtr(rdEntryStr(d, "dhcp")),
		DhcpRange: dhcpRange,
		DNS:       rdEntryStrList(d, "dns"),
		Domain:    rdEntryStr(d, "domain"),
		Gateway:   rdEntryStr(d, "gateway"),
		Mask:      rdEntryStr(d, "mask"),
		Ntp:       rdEntryStr(d, "ntp"),
		Subnet:    rdEntryStr(d, "subnet"),
	}, nil
}
func rdMapIPSpecOrNil(d *schema.ResourceData) (*swagger_models.IPSpec, error) {
	val, err := rdEntryStructPtr(d, "ip", rdMapIPSpec)
	if err != nil {
		return nil, err
	}
	if val == nil {
		return nil, nil
	}
	return val.(*swagger_models.IPSpec), nil
}

func networkKindPtr(strVal string) *swagger_models.NetworkKind {
	val := swagger_models.NetworkKind(strVal)
	return &val
}

func rdToNetConfig(cfg *swagger_models.NetConfig, d *schema.ResourceData) error {
	var err error
	cfg.Description = rdEntryStr(d, "description")
	cfg.DNSList, err = resourceDataToStaticDNSList(d)
	if err != nil {
		return err
	}
	cfg.EnterpriseDefault = rdEntryBool(d, "enterprise_default")
	cfg.IP, err = rdMapIPSpecOrNil(d)
	if err != nil {
		return err
	}
	cfg.Kind = networkKindPtr(rdEntryStr(d, "kind"))
	cfg.ProjectID = rdEntryStrPtrOrNil(d, "project_id")
	cfg.Proxy, err = rdMapNetProxyConfigOrNil(d)
	if err != nil {
		return err
	}
	cfg.Title = rdEntryStrPtrOrNil(d, "title")
	cfg.Wireless, err = rdMapNetWirelessConfigOrNil(d)
	if err != nil {
		return err
	}
	return nil
}

// Create the Resource Group
func createNetworkResource(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	client := (meta.(Client)).Client
	name := rdEntryStr(d, "name")
	errMsgPrefix := getErrMsgPrefix(name, "", "App Instance", "Create")
	log.Printf("[INFO] Creating Network: %s", name)
	cfg := &swagger_models.NetConfig{
		Name: &name,
	}
	err := rdToNetConfig(cfg, d)
	if err != nil {
		return diag.Errorf("%s Error: %s", errMsgPrefix, err.Error())
	}
	client.XRequestIdPrefix = "TF-create-network"
	rspData := &swagger_models.ZsrvResponse{}
	_, err = client.SendReq("POST", networkUrlExtension, cfg, rspData)
	if err != nil {
		return diag.Errorf("\n***[ERROR] Failed to create network. err: %s",
			err.Error())
	}
	log.Printf("Network %s (ID: %s) Successfully created\n",
		rspData.ObjectName, rspData.ObjectID)
	d.SetId(rspData.ObjectID)
	err = getNetworkAndPublishData(client, d, name, rspData.ObjectID, true)
	if err != nil {
		log.Printf("***[ERROR]- Failed to get Network: %s (ID: %s) after "+
			"creating it. Err: %s", name, rspData.ObjectID, err.Error())
	}
	return diags
}

// Update Network
func updateNetworkResource(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	client := (meta.(Client)).Client
	name := rdEntryStr(d, "name")
	id := rdEntryStr(d, "id")
	errMsgPrefix := getErrMsgPrefix(name, id, "Network ", "Update")
	if client == nil {
		return diag.Errorf("%s nil Client", errMsgPrefix)
	}
	cfg, err, _ := getNetworkConfig(client, name, id)
	if err != nil {
		return diag.Errorf("%s err: %s", errMsgPrefix, err.Error())
	}
	if cfg == nil {
		return diag.Errorf("%s Nil Config. Failed to find Network", errMsgPrefix)
	}
	err = checkInvalidAttrForUpdate(d, *cfg.Name, cfg.ID)
	if err != nil {
		return diag.Errorf("%s err: %s", errMsgPrefix, err.Error())
	}
	log.Printf("[INFO] Found Network: %s (ID: %s)", name, cfg.ID)
	err = rdToNetConfig(cfg, d)
	if err != nil {
		return diag.Errorf("%s Error: %s", errMsgPrefix, err.Error())
	}
	log.Printf("[INFO] Updating Network: %s", name)
	client.XRequestIdPrefix = "TF-nwinst-update"
	urlExtension := getNetworkUrl(name, id, "update")
	rspData := &swagger_models.ZsrvResponse{}
	_, err = client.SendReq("PUT", urlExtension, cfg, rspData)
	if err != nil {
		return diag.Errorf("%s err: %s", errMsgPrefix, err.Error())
	}
	log.Printf("Network Update Successful")
	err = getNetworkAndPublishData(client, d, name, id, true)
	if err != nil {
		log.Printf("***[ERROR]- Failed to get Network: %s (ID: %s) after "+
			"updating it. Err: %s", name, cfg.ID, err.Error())
	}
	return diags
}

// Delete the Resource Group
func deleteNetworkResource(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	client := (meta.(Client)).Client
	name := rdEntryStr(d, "name")
	id := rdEntryStr(d, "id")
	errMsgPrefix := fmt.Sprintf("[ERROR] Network %s ( id: %s) Delete Failed.",
		name, id)
	_, err, httpRsp := getNetworkConfig(client, name, id)
	if err != nil {
		if httpRsp != nil && zedcloudapi.IsObjectNotFound(httpRsp) {
			log.Printf("%s Not Found", errMsgPrefix)
			return diags
		}
		return diag.Errorf("%s Network not found. err: %s",
			errMsgPrefix, err.Error())
	}
	urlExtension := getNetworkUrl(name, id, "delete")
	rspData := &swagger_models.ZsrvResponse{}
	_, err = client.SendReq("DELETE", urlExtension, nil, rspData)
	if err != nil {
		return diag.Errorf("%s. Delete request failed. Err: %s", errMsgPrefix, err.Error())
	}
	log.Printf("[INFO] Network %s ( ID: %s) Delete Successful.", name, id)
	return diags
}

func readResourceNetwork(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return readNetwork(ctx, d, meta, true)
}
