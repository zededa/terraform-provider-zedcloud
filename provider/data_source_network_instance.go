// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	zschemas "github.com/zededa/terraform-provider-zedcloud/schemas"
	zedcloudapi "github.com/zededa/zedcloud-api"
	"github.com/zededa/zedcloud-api/swagger_models"
	"log"
	"net/http"
)

var NetInstDataSourceSchema = &schema.Resource{
	ReadContext: readDataSourceNetInst,
	Schema:      zschemas.NetworkInstanceSchema,
	Description: "Schema for data source zedcloud_network_instance. Must specify " +
		"id or name",
}

// The schema for a resource group data source
func getNetInstDataSourceSchema() *schema.Resource {
	return NetInstDataSourceSchema
}

func getNetInstConfig(client *zedcloudapi.Client,
	name, id string) (*swagger_models.NetInstConfig, error, *http.Response) {
	rspData := &swagger_models.NetInstConfig{}
	httpResp, err := client.GetObj(netInstUrlExtension, name, id, false, rspData)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] FAILED to get NetInst %s ( id: %s). Err: %s",
			name, id, err.Error()), httpResp
	}
	return rspData, nil, httpResp
}

func flattenNetInstOpaqueConfig(cfg *swagger_models.NetInstOpaqueConfig) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	return []interface{}{map[string]interface{}{
		"oconfig": cfg.Oconfig,
		"type":    ptrValStr(cfg.Type),
	}}
}

func flattenDhcpIPRange(cfg *swagger_models.DhcpIPRange) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	return []interface{}{map[string]interface{}{
		"end":   cfg.End,
		"start": cfg.Start,
	}}
}

func flattenDhcpServerConfig(cfg *swagger_models.DhcpServerConfig) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	dns := make([]interface{}, 0)
	for _, entry := range cfg.DNS {
		dns = append(dns, entry)
	}
	return []interface{}{map[string]interface{}{
		"dhcp_range": flattenDhcpIPRange(cfg.DhcpRange),
		"dns":        dns,
		"domain":     cfg.Domain,
		"gateway":    cfg.Gateway,
		"mask":       cfg.Mask,
		"ntp":        cfg.Ntp,
		"subnet":     cfg.Subnet,
	}}
}

func flattenStaticDNSList(cfgList []*swagger_models.StaticDNSList) []interface{} {
	entryList := make([]interface{}, 0)
	for _, cfg := range cfgList {
		entryList = append(entryList, map[string]interface{}{
			"addrs":    flattenStringList(cfg.Addrs),
			"hostname": cfg.Hostname,
		})
	}
	return entryList
}

func flattenNetInstConfig(cfg *swagger_models.NetInstConfig) map[string]interface{} {
	if cfg == nil {
		return map[string]interface{}{}
	}
	data := map[string]interface{}{
		"id":         cfg.ID,
		"cluster_id": cfg.ClusterID,
		"project_id": cfg.ProjectID,
		"revision":   flattenObjectRevision(cfg.Revision),
	}
	data["description"] = cfg.Description
	data["device_default"] = cfg.DeviceDefault
	data["device_id"] = ptrValStr(cfg.DeviceID)
	data["dns_list"] = flattenStaticDNSList(cfg.DNSList)
	data["ip"] = flattenDhcpServerConfig(cfg.IP)
	data["kind"] = ptrValStr(cfg.Kind)
	data["name"] = ptrValStr(cfg.Name)
	data["network_policy_id"] = cfg.NetworkPolicyID
	data["opaque"] = flattenNetInstOpaqueConfig(cfg.Opaque)
	data["port"] = ptrValStr(cfg.Port)
	data["port_tags"] = flattenStringMap(cfg.PortTags)
	data["tags"] = flattenStringMap(cfg.Tags)
	data["title"] = ptrValStr(cfg.Title)
	data["type"] = ptrValStr(cfg.Type)
	flattenedDataCheckKeys(zschemas.NetworkInstanceSchema, data)
	return data
}

func getNetInstAndPublishData(client *zedcloudapi.Client, d *schema.ResourceData,
	name, id string) error {
	cfg, err, httpRsp := getNetInstConfig(client, name, id)
	if err != nil {
		err = fmt.Errorf("[ERROR] Network Instance %s (id: %s) not found. Err: %s",
			name, id, err.Error())
		if httpRsp != nil && zedcloudapi.IsObjectNotFound(httpRsp) {
			log.Printf("PROVIDER:  REMOVED Network Instance %s (id: %s) from State", id, name)
			schema.RemoveFromState(d, nil)
			return nil
		}
		return err
	}
	marshalData(d, flattenNetInstConfig(cfg))
	return nil
}

// Read the Resource Group
func readNetInst(ctx context.Context, d *schema.ResourceData,
	meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := (meta.(Client)).Client
	id := rdEntryStr(d, "id")
	name := rdEntryStr(d, "name")
	if client == nil {
		return diag.Errorf("nil Client")
	}
	log.Printf("PROVIDER:  readNetInst - id: %s, name: %s", id, name)
	if (id == "") && (name == "") {
		return diag.Errorf("The arguments \"id\" or \"name\" are required, but no definition was found.")
	}
	err := getNetInstAndPublishData(client, d, name, id)
	if err != nil {
		log.Printf(err.Error())
		return diag.Errorf("%s", err.Error())
	}
	return diags
}

func readDataSourceNetInst(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return readNetInst(ctx, d, meta)
}
