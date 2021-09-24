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
)

var NetInstDataSourceSchema = &schema.Resource{
	ReadContext: readNetInst,
	Schema:      zschemas.NetworkInstanceSchema,
	Description: "Schema for data source zedcloud_network_instance. Must specify " +
		"id or name",
}

// The schema for a resource group data source
func getNetInstDataSourceSchema() *schema.Resource {
	return NetInstDataSourceSchema
}

func getNetInstConfig(client *zedcloudapi.Client,
	name, id string) (*swagger_models.NetInstConfig, error) {
	rspData := &swagger_models.NetInstConfig{}
	err := client.GetObj(netInstUrlExtension, name, id, false, rspData)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] FAILED to get NetInst %s ( id: %s). Err: %s",
			name, id, err.Error())
	}
	return rspData, nil
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
	return map[string]interface{}{
		"description":       cfg.Description,
		"device_default":    cfg.DeviceDefault,
		"device_id":         ptrValStr(cfg.DeviceID),
		"dns_list":          flattenStaticDNSList(cfg.DNSList),
		"id":                cfg.ID,
		"ip":                flattenDhcpServerConfig(cfg.IP),
		"kind":              ptrValStr(cfg.Kind),
		"name":              ptrValStr(cfg.Name),
		"network_policy_id": cfg.NetworkPolicyID,
		"opaque":            flattenNetInstOpaqueConfig(cfg.Opaque),
		"port":              ptrValStr(cfg.Port),
		"port_tags":         flattenStringMap(cfg.PortTags),
		"project_id":        cfg.ProjectID,
		"revision":          flattenObjectRevision(cfg.Revision),
		"tags":              flattenStringMap(cfg.Tags),
		"title":             ptrValStr(cfg.Title),
		"type":              ptrValStr(cfg.Type),
	}
}

// Read the Resource Group
func readNetInst(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	cfg, err := getNetInstConfig(client, name, id)
	if err != nil {
		errStr := fmt.Sprintf("[ERROR] network instance %s (id: %s) not found. Err: %s",
			name, id, err.Error())
		log.Printf(errStr)
		return diag.Errorf(errStr)
	}
	// Take the Config and convert it to terraform object
	marshalData(d, flattenNetInstConfig(cfg))
	return diags
}
