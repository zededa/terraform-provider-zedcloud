// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/internal/resourcedata"
	zschemas "github.com/zededa/terraform-provider-zedcloud/schemas"
	"github.com/zededa/terraform-provider-zedcloud/utils"
	zedcloudapi "github.com/zededa/zedcloud-api"
	"github.com/zededa/zedcloud-api/swagger_models"
)

var NetworkDataSchema = &schema.Resource{
	ReadContext: readDataSourceNetwork,
	Schema:      zschemas.NetworkSchema,
	Description: "Schema for data source zedcloud_network. Must specify id or name",
}

// The schema for a resource group data source
func getNetworkDataSourceSchema() *schema.Resource {
	return NetworkDataSchema
}

func getNetworkConfig(client *zedcloudapi.Client,
	name, id string) (*swagger_models.NetConfig, error, *http.Response) {
	rspData := &swagger_models.NetConfig{}
	httpRsp, err := client.GetObj(networkUrlExtension, name, id, false, rspData)
	if err != nil {
		err = fmt.Errorf("[ERROR] FAILED to get Network %s ( id: %s). Err: %s",
			name, id, err.Error())
	}
	return rspData, err, httpRsp
}

func flattenNetWifiConfig(cfg *swagger_models.NetWifiConfig) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	// Do not publish sensitive fields - "secret", identity
	return []interface{}{map[string]interface{}{
		"key_scheme": utils.PtrValStr(cfg.KeyScheme),
		"priority":   int(cfg.Priority),
		"wifi_ssid":  cfg.WifiSSID,
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
	if cfg.Type != nil && *cfg.Type == "NETWORK_WIRELESS_TYPE_UNSPECIFIED" {
		// Ideally, ZEDCloud should not send this block in GETConfig API.
		// Till that is fixed in ZEDCloud, handle it in provider
		return []interface{}{}
	}
	return []interface{}{map[string]interface{}{
		"cellular_cfg": flattenNetCellularConfig(cfg.CellularCfg),
		"type":         utils.PtrValStr(cfg.Type),
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
			"proto":  utils.PtrValStr(cfg.Proto),
			"server": cfg.Server,
		})
	}
	return entryList
}

func flattenNetProxyConfig(cfg *swagger_models.NetProxyConfig) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	if !cfg.NetworkProxy {
		// Ideally, ZEDCloud should not send this block in GETConfig API.
		// Till that is fixed in ZEDCloud, handle it in provider
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
		"dhcp":       utils.PtrValStr(cfg.Dhcp),
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
	data := map[string]interface{}{
		"id":       cfg.ID,
		"revision": utils.FlattenObjectRevision(cfg.Revision),
	}
	data["description"] = cfg.Description
	data["dns_list"] = flattenStaticDNSList(cfg.DNSList)
	data["enterprise_default"] = cfg.EnterpriseDefault
	data["ip"] = flattenIPSpec(cfg.IP)
	data["kind"] = utils.PtrValStr(cfg.Kind)
	data["name"] = utils.PtrValStr(cfg.Name)
	data["project_id"] = utils.PtrValStr(cfg.ProjectID)
	data["proxy"] = flattenNetProxyConfig(cfg.Proxy)
	data["title"] = utils.PtrValStr(cfg.Title)
	data["wireless"] = flattenNetWirelessConfig(cfg.Wireless)
	utils.CheckIfAllKeysExist(zschemas.NetworkSchema, data)
	return data
}

func getNetworkAndPublishData(client *zedcloudapi.Client, d *schema.ResourceData,
	name, id string) error {
	cfg, err, httpRsp := getNetworkConfig(client, name, id)
	if err != nil {
		err = fmt.Errorf("[ERROR] Network %s (id: %s) not found. Err: %s",
			name, id, err.Error())
		if httpRsp != nil && zedcloudapi.IsObjectNotFound(httpRsp) {
			log.Printf("PROVIDER:  REMOVED Network %s (id: %s) from State", id, name)
			schema.RemoveFromState(d, nil)
			return nil
		}
		return err
	}
	marshalData(d, flattenNetConfig(cfg))
	return nil
}

func readNetwork(ctx context.Context, d *schema.ResourceData,
	meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := (meta.(Client)).Client
	id := resourcedata.GetStr(d, "id")
	name := resourcedata.GetStr(d, "name")

	log.Printf("PROVIDER:  readNetworkResource - id: %s, name: %s", id, name)
	if (id == "") && (name == "") {
		return diag.Errorf("The arguments \"id\" or \"name\" are required, but no definition was found.")
	}
	err := getNetworkAndPublishData(client, d, name, id)
	if err != nil {
		return diag.Errorf("%s", err.Error())
	}
	return diags
}

func readDataSourceNetwork(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return readNetwork(ctx, d, meta)
}
