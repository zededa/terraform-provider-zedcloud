// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"context"
	"fmt"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	zschemas "github.com/zededa/terraform-provider-zedcloud/schemas"
	zedcloudapi "github.com/zededa/zedcloud-api"
	"github.com/zededa/zedcloud-api/swagger_models"
	"log"
)

var NetworkDataSchema = &schema.Resource{
	ReadContext: readNetwork,
	Schema:      zschemas.NetworkSchema,
	Description: "Schema for data source zedcloud_network. Must specify id or name",
}

// The schema for a resource group data source
func getNetworkDataSourceSchema() *schema.Resource {
	return NetworkDataSchema
}

func getNetworkConfig(client *zedcloudapi.Client,
	name, id string) (*swagger_models.NetConfig, error) {
	rspData := &swagger_models.NetConfig{}
	err := client.GetObj(networkUrlExtension, name, id, false, rspData)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] FAILED to get Network %s ( id: %s). Err: %s",
			name, id, err.Error())
	}
	return rspData, nil
}

func flattenNetWifiConfigNetcryptoblock(cfg *swagger_models.NetWifiConfigNetcryptoblock) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	return []interface{}{map[string]interface{}{
		"identity": cfg.Identity,
		"password": cfg.Password,
	}}
}

func flattenNetWifiConfigSecrets(cfg *swagger_models.NetWifiConfigSecrets) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	return []interface{}{map[string]interface{}{
		"wifi_passwd": cfg.WiFiPasswd,
	}}
}

func flattenNetWifiConfig(cfg *swagger_models.NetWifiConfig) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	return []interface{}{map[string]interface{}{
		"crypto":            flattenNetWifiConfigNetcryptoblock(cfg.Crypto),
		"crypto_key":        cfg.CryptoKey,
		"encrypted_secrets": flattenStringMap(cfg.EncryptedSecrets),
		"identity":          cfg.Identity,
		"key_scheme":        ptrValStr(cfg.KeyScheme),
		"priority":          int(cfg.Priority),
		"secret":            flattenNetWifiConfigSecrets(cfg.Secret),
		"wifi_ssid":         cfg.WifiSSID,
	}}
}

func flattenNetCellularConfig(cfg *swagger_models.NetCellularConfig) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	return []interface{}{map[string]interface{}{
		"apn": cfg.APN,
	}}
}

func flattenNetWirelessConfig(cfg *swagger_models.NetWirelessConfig) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	return []interface{}{map[string]interface{}{
		"cellular_cfg": flattenNetCellularConfig(cfg.CellularCfg),
		"type":         ptrValStr(cfg.Type),
		"wifi_cfg":     flattenNetWifiConfig(cfg.WifiCfg),
	}}
}

func flattenBase64List(cfgList []strfmt.Base64) []interface{} {
	entryList := make([]interface{}, 0)
	for _, cfg := range cfgList {
		entryList = append(entryList, string(cfg))
	}
	return entryList
}

func flattenNetProxyServerList(cfgList []*swagger_models.NetProxyServer) []interface{} {
	entryList := make([]interface{}, 0)
	for _, cfg := range cfgList {
		entryList = append(entryList, map[string]interface{}{
			"port":   int(cfg.Port),
			"proto":  ptrValStr(cfg.Proto),
			"server": cfg.Server,
		})
	}
	return entryList
}

func flattenNetProxyConfig(cfg *swagger_models.NetProxyConfig) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	return []interface{}{map[string]interface{}{
		"exceptions":          cfg.Exceptions,
		"network_proxy":       cfg.NetworkProxy,
		"network_proxy_certs": flattenBase64List(cfg.NetworkProxyCerts),
		"network_proxy_url":   cfg.NetworkProxyURL,
		"pacfile":             cfg.Pacfile,
		"proxy":               flattenNetProxyServerList(cfg.Proxies),
	}}
}

func flattenIPSpec(cfg *swagger_models.IPSpec) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	return []interface{}{map[string]interface{}{
		"dhcp":       ptrValStr(cfg.Dhcp),
		"dhcp_range": flattenDhcpIPRange(cfg.DhcpRange),
		"dns":        flattenStringList(cfg.DNS),
		"domain":     cfg.Domain,
		"gateway":    cfg.Gateway,
		"mask":       cfg.Mask,
		"ntp":        cfg.Ntp,
		"subnet":     cfg.Subnet,
	}}
}

func flattenNetConfig(cfg *swagger_models.NetConfig) map[string]interface{} {
	if cfg == nil {
		return map[string]interface{}{}
	}
	return map[string]interface{}{
		"description":        cfg.Description,
		"dns_list":           flattenStaticDNSList(cfg.DNSList),
		"enterprise_default": cfg.EnterpriseDefault,
		"id":                 cfg.ID,
		"ip":                 flattenIPSpec(cfg.IP),
		"kind":               ptrValStr(cfg.Kind),
		"name":               ptrValStr(cfg.Name),
		"project_id":         ptrValStr(cfg.ProjectID),
		"proxy":              flattenNetProxyConfig(cfg.Proxy),
		"revision":           flattenObjectRevision(cfg.Revision),
		"title":              ptrValStr(cfg.Title),
		"wireless":           flattenNetWirelessConfig(cfg.Wireless),
	}
}

func readNetwork(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := (meta.(Client)).Client
	id := rdEntryStr(d, "id")
	name := rdEntryStr(d, "name")

	log.Printf("PROVIDER:  readNetworkResource - id: %s, name: %s", id, name)
	if (id == "") && (name == "") {
		return diag.Errorf("The arguments \"id\" or \"name\" are required, but no definition was found.")
	}
	cfg, err := getNetworkConfig(client, name, id)
	if err != nil {
		return diag.Errorf("[ERROR] Network %s (id: %s) not found. Err: %s",
			name, id, err.Error())
	}
	// Take the Config and convert it to terraform object
	marshalData(d, flattenNetConfig(cfg))
	return diags
}
