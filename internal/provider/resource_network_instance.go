// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/internal/client"
	"github.com/zededa/terraform-provider-zedcloud/internal/resourcedata"
	"github.com/zededa/terraform-provider-zedcloud/schemas"
	zedcloudapi "github.com/zededa/zedcloud-api"
	"github.com/zededa/zedcloud-api/swagger_models"
)

var netInstUrlExtension = "netinsts"

var NetInstResourceSchema = &schema.Resource{
	CreateContext: createNetInstResource,
	ReadContext:   readResourceNetInst,
	UpdateContext: updateNetInstResource,
	DeleteContext: deleteNetInstResource,
	Schema:        schemas.NetworkInstanceSchema,
	Importer: &schema.ResourceImporter{
		StateContext: schema.ImportStatePassthroughContext,
	},
}

// The schema for a resource group data source
func getNetInstResourceSchema() *schema.Resource {
	return NetInstResourceSchema
}

func getNetInstUrl(name, id, urlType string) string {
	return zedcloudapi.UrlForObjectRequest(netInstUrlExtension, name, id, urlType)
}

func resourceDataMapToNetInstOpaqueConfig(d map[string]interface{}) (interface{}, error) {
	oconfigType := swagger_models.OpaqueConfigType(resourcedata.GetStr(d, "type"))
	return &swagger_models.NetInstOpaqueConfig{
		Oconfig: resourcedata.GetStr(d, "oconfig"),
		Type:    &oconfigType,
	}, nil
}

func resourceDataToNetInstOpaqueConfig(d *schema.ResourceData) (
	*swagger_models.NetInstOpaqueConfig, error) {
	val, err := resourcedata.GetStructPtr(d, "opaque", resourceDataMapToNetInstOpaqueConfig)
	if val == nil {
		return nil, nil
	}
	return val.(*swagger_models.NetInstOpaqueConfig), err
}

func rdMapDhcpIpRange(d map[string]interface{}) (interface{}, error) {
	return &swagger_models.DhcpIPRange{
		End:   resourcedata.GetStr(d, "end"),
		Start: resourcedata.GetStr(d, "start"),
	}, nil
}

func rdMapDhcpIpRangeOrNil(d interface{}) (*swagger_models.DhcpIPRange, error) {
	val, err := resourcedata.GetStructPtr(d, "dhcp_range", rdMapDhcpIpRange)
	if err != nil {
		return nil, err
	}
	if val == nil {
		return nil, nil
	}
	return val.(*swagger_models.DhcpIPRange), nil
}

func resourceDataMapToDhcpServerConfig(d map[string]interface{}) (interface{}, error) {
	dhcpRange, err := rdMapDhcpIpRangeOrNil(d)
	if err != nil {
		return nil, err
	}
	return &swagger_models.DhcpServerConfig{
		DhcpRange: dhcpRange,
		DNS:       resourcedata.GetStrList(d, "dns"),
		Domain:    resourcedata.GetStr(d, "domain"),
		Gateway:   resourcedata.GetStr(d, "gateway"),
		Mask:      resourcedata.GetStr(d, "mask"),
		Ntp:       resourcedata.GetStr(d, "ntp"),
		Subnet:    resourcedata.GetStr(d, "subnet"),
	}, nil
}

func resourceDataToDhcpServerConfig(d *schema.ResourceData) (*swagger_models.DhcpServerConfig, error) {
	val, err := resourcedata.GetStructPtr(d, "ip", resourceDataMapToDhcpServerConfig)
	if err != nil {
		return nil, err
	}
	if val == nil {
		return nil, nil
	}
	return val.(*swagger_models.DhcpServerConfig), nil
}

func resourceDataToStaticDNSList(d *schema.ResourceData) ([]*swagger_models.StaticDNSList, error) {
	dnsList := make([]*swagger_models.StaticDNSList, 0)
	data := rdEntryList(d, "dns_list")
	for _, entry := range data {
		dataEntry := entry.(map[string]interface{})
		dnsEntry := &swagger_models.StaticDNSList{
			Addrs:    resourcedata.GetStrSet(dataEntry, "addrs"),
			Hostname: resourcedata.GetStr(dataEntry, "hostname"),
		}
		dnsList = append(dnsList, dnsEntry)
	}
	return dnsList, nil
}

func rdUpdateNetInstConfig(d *schema.ResourceData,
	cfg *swagger_models.NetInstConfig) error {
	dnsList, err := resourceDataToStaticDNSList(d)
	if err != nil {
		return err
	}
	dhcpServerCfg, err := resourceDataToDhcpServerConfig(d)
	if err != nil {
		return err
	}
	kind := swagger_models.NetworkInstanceKind(resourcedata.GetStr(d, "kind"))
	opaque, err := resourceDataToNetInstOpaqueConfig(d)
	if err != nil {
		return err
	}
	dhcpType := swagger_models.NetworkInstanceDhcpType(resourcedata.GetStr(d, "type"))

	cfg.Description = resourcedata.GetStr(d, "description")
	cfg.DeviceDefault = rdEntryBool(d, "device_default")
	cfg.DeviceID = resourcedata.GetStrPtrOrNil(d, "device_id")
	cfg.DNSList = dnsList
	cfg.IP = dhcpServerCfg
	cfg.Kind = &kind
	cfg.NetworkPolicyID = resourcedata.GetStr(d, "network_policy_id")
	cfg.Opaque = opaque
	cfg.Port = resourcedata.GetStrPtrOrNil(d, "port")
	cfg.PortTags = resourcedata.GetStrMap(d, "port_tags")
	cfg.ProjectID = resourcedata.GetStr(d, "project_id")
	cfg.Tags = resourcedata.GetStrMap(d, "tags")
	cfg.Title = resourcedata.GetStrPtrOrNil(d, "title")
	cfg.Type = &dhcpType

	return nil
}

// Create the Resource Group
func createNetInstResource(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	client, ok := meta.(*client.Client)
	if !ok {
		return diag.Errorf("expect meta to be of type client.Client{} but is %T", meta)
	}
	name := resourcedata.GetStr(d, "name")
	errMsgPrefix := getErrMsgPrefix(name, "", "NetworkInstance", "Create")
	if name == "" {
		return diag.Errorf("%s \"name\" must be specified.", errMsgPrefix)
	}
	// NetInst Config doesn't exist. Create it
	log.Printf("[INFO] Creating NetInst: %s", name)
	cfg := &swagger_models.NetInstConfig{
		Name: &name,
	}
	err := rdUpdateNetInstConfig(d, cfg)
	if err != nil {
		return diag.Errorf("%s %s", errMsgPrefix, err.Error())
	}
	client.XRequestIdPrefix = "TF-nwinst-create"
	rspData := &swagger_models.ZsrvResponse{}
	_, err = client.SendReq("POST", "netinsts", cfg, rspData)
	if err != nil {
		return diag.Errorf("%s %s", errMsgPrefix, err.Error())
	}
	log.Printf("Network Instance %s (ID: %s) Successfully created\n",
		rspData.ObjectName, rspData.ObjectID)

	id := rspData.ObjectID
	d.SetId(id)
	err = getNetInstAndPublishData(client, d, name, id)
	if err != nil {
		log.Printf("***[ERROR]- Failed to get Network Instance: %s (ID: %s) after "+
			"creating it. Err: %s", name, id, err.Error())
	}
	return diags
}

func validateNetInstUpdateOperation(client *zedcloudapi.Client,
	d *schema.ResourceData) (*swagger_models.NetInstConfig, error) {
	if client == nil {
		return nil, fmt.Errorf("nil Client")
	}
	id := resourcedata.GetStr(d, "id")
	if id == "" {
		return nil, fmt.Errorf("id cannot be empty string for Update operation")
	}
	cfg, err, _ := getNetInstConfig(client, "", id)
	if err != nil {
		return nil, err
	}
	err = checkInvalidAttrForUpdate(d, *cfg.Name, cfg.ID)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

// Update the Resource Group
func updateNetInstResource(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	client, ok := meta.(*client.Client)
	if !ok {
		return diag.Errorf("expect meta to be of type client.Client{} but is %T", meta)
	}
	name := resourcedata.GetStr(d, "name")
	id := resourcedata.GetStr(d, "id")
	errMsgPrefix := getErrMsgPrefix(name, id, "Network Instance", "Update")
	cfg, err := validateNetInstUpdateOperation(client, d)
	if err != nil {
		return diag.Errorf("%s err: %s", errMsgPrefix, err.Error())
	}
	err = rdUpdateNetInstConfig(d, cfg)
	if err != nil {
		return diag.Errorf("%s %s", errMsgPrefix, err.Error())
	}
	log.Printf("[INFO] Updating NetInst: %s", name)
	client.XRequestIdPrefix = "TF-nwinst-update"
	urlExtension := getNetInstUrl(name, id, "update")
	rspData := &swagger_models.ZsrvResponse{}
	_, err = client.SendReq("PUT", urlExtension, cfg, rspData)
	if err != nil {
		return diag.Errorf("%s err: %s", errMsgPrefix, err.Error())
	}
	log.Printf("Network Instance Update Successful")
	err = getNetInstAndPublishData(client, d, name, id)
	if err != nil {
		log.Printf("***[ERROR]- Failed to get Network Instance: %s (ID: %s) after "+
			"updating it. Err: %s", name, id, err.Error())
	}
	return diags
}

// Delete the Resource Group
func deleteNetInstResource(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	client, ok := meta.(*client.Client)
	if !ok {
		return diag.Errorf("expect meta to be of type client.Client{} but is %T", meta)
	}
	name := resourcedata.GetStr(d, "name")
	id := resourcedata.GetStr(d, "id")
	errMsgPrefix := getErrMsgPrefix(name, id, "Network Instance", "Delete")
	_, err, httpRsp := getNetInstConfig(client, name, id)
	if err != nil {
		if httpRsp != nil && zedcloudapi.IsObjectNotFound(httpRsp) {
			log.Printf("%s Not Found", errMsgPrefix)
			return diags
		}
		return diag.Errorf("%s Network Instance not found. err: %s",
			errMsgPrefix, err.Error())
	}
	client.XRequestIdPrefix = "TF-nwinst-delete"
	urlExtension := getNetInstUrl(name, id, "delete")
	rspData := &swagger_models.ZsrvResponse{}
	_, err = client.SendReq("DELETE", urlExtension, nil, rspData)
	if err != nil {
		return diag.Errorf("%s. Err: %s", errMsgPrefix, err.Error())
	}
	log.Printf("[INFO] Network Instance %s ( ID: %s) Delete Successful.", name, id)
	return diags
}
func readResourceNetInst(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return readNetInst(ctx, d, meta)
}
