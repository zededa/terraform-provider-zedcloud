// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"net/http"

	zschemas "github.com/zededa/terraform-provider-zedcloud/schemas"
	zedcloudapi "github.com/zededa/zedcloud-api"
	"github.com/zededa/zedcloud-api/swagger_models"
)

var deviceUrlExtension = "devices"

var EdgeNodeResourceSchema = &schema.Resource{
	CreateContext: createEdgeNodeResource,
	ReadContext:   readEdgeNode,
	UpdateContext: updateEdgeNodeResource,
	DeleteContext: deleteEdgeNodeResource,
	Schema:        zschemas.EdgeNodeSchema,
	Importer: &schema.ResourceImporter{
		StateContext: schema.ImportStatePassthroughContext,
	},
}

func getEdgeNodeUrl(name, id, urlType string) string {
	return zedcloudapi.UrlForObjectRequest(deviceUrlExtension, name, id, urlType)
}

// The schema for a resource group data source
func getEdgeNodeResourceSchema() *schema.Resource {
	return EdgeNodeResourceSchema
}

func getEdgeNodeConfig(client *zedcloudapi.Client,
	name, id string) (*swagger_models.DeviceConfig, error) {
	rspData := &swagger_models.DeviceConfig{}
	err := client.GetObj(deviceUrlExtension, name, id, false, rspData)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] FAILED to get edge node %s ( id: %s). Err: %s",
			name, id, err.Error())
	}
	return rspData, nil
}

func edgeNodeSendPutReq(client *zedcloudapi.Client, cfg *swagger_models.DeviceConfig,
	reqType string) (*http.Response, error) {
	client.XRequestIdPrefix = "TF-edgenode-" + reqType
	urlExtension := getEdgeNodeUrl(*cfg.Name, cfg.ID, reqType)
	rspData := &swagger_models.ZsrvResponse{}
	return client.SendReq("PUT", urlExtension, cfg, rspData)
}

func cfgBaseosForEveVersionStr(eve_image_version string) []*swagger_models.BaseOSImage {
	return []*swagger_models.BaseOSImage{&swagger_models.BaseOSImage{
		ImageName: &eve_image_version,
		Version:   &eve_image_version,
		Activate:  true,
	}}
}

func edgeNodeUpdateBaseOs(client *zedcloudapi.Client, cfg *swagger_models.DeviceConfig,
	eve_image_version string) error {
	// BaseImage is supposed to have only one entry. If there are multiple,
	// reset the BaseOS config
	if cfg.BaseImage != nil && len(cfg.BaseImage) == 1 &&
		*cfg.BaseImage[0].ImageName == eve_image_version &&
		cfg.BaseImage[0].Activate {
		// Baseos update not required
		return nil
	}
	cfg.BaseImage = cfgBaseosForEveVersionStr(eve_image_version)
	log.Printf("[TRACE]: Update BaseOsImage. ImageName: %s, Activate: %t",
		*cfg.BaseImage[0].ImageName, cfg.BaseImage[0].Activate)
	// BaseOs update is Special case - Publish Config followed by ApplyConfig
	_, err := edgeNodeSendPutReq(client, cfg, "publish")
	if err != nil {
		return fmt.Errorf("BaseOsImage Publish failed. Err: %s", err.Error())
	}
	// We need to apply the configuration
	rsp, err := edgeNodeSendPutReq(client, cfg, "apply")
	// Ignore HTTP_STATUS_CONFLICT
	if err != nil && rsp.StatusCode != http.StatusConflict {
		return fmt.Errorf("BaseOsUpdate Apply request failed. Err: %s", err.Error())
	}
	return nil
}

// Create the Edge Node
func createEdgeNodeResource(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Errorf("[ERROR] terraform-provider-zededa does not support Creation of Edge Nodes")
}

func rdConfigItems(d *schema.ResourceData) ([]*swagger_models.EDConfigItem, error) {
	cfgItems := make([]*swagger_models.EDConfigItem, 0)
	val, exists := d.GetOk("config_items")
	if !exists {
		return cfgItems, nil
	}
	rdCfgItems := val.(map[string]interface{})
	for key, val := range rdCfgItems {
		strVal, ok := val.(string)
		if !ok {
			return cfgItems, fmt.Errorf("Value (%+v) for Config Item (%s) must be a valid string.",
				val, key)
		}
		cfgItems = append(cfgItems, &swagger_models.EDConfigItem{
			Key:         key,
			StringValue: strVal,
		})
	}
	return cfgItems, nil
}

func setDeviceLocation(cfg *swagger_models.DeviceConfig, d *schema.ResourceData) error {
	locationInfo := rdEntryList(d, "dev_location")
	if len(locationInfo) == 0 {
		// Location Info not specified.
		cfg.DevLocation = nil
		return nil
	}
	if len(locationInfo) > 1 {
		return fmt.Errorf("Location info must be specified only Once. The "+
			"configuration has %d blocks for device location", len(locationInfo))
	}
	entry := locationInfo[0].(map[string]interface{})
	cfgLocation := swagger_models.GeoLocation{}
	if val, ok := entry["city"]; ok {
		cfgLocation.City = val.(string)
	}
	if val, ok := entry["country"]; ok {
		cfgLocation.Country = val.(string)
	}
	if val, ok := entry["freeloc"]; ok {
		cfgLocation.Freeloc = val.(string)
	}
	if val, ok := entry["hostname"]; ok {
		cfgLocation.Hostname = val.(string)
	}
	if val, ok := entry["loc"]; ok {
		cfgLocation.Loc = val.(string)
	}
	if val, ok := entry["org"]; ok {
		cfgLocation.Org = val.(string)
	}
	if val, ok := entry["postal"]; ok {
		cfgLocation.Postal = val.(string)
	}
	if val, ok := entry["region"]; ok {
		cfgLocation.Region = val.(string)
	}
	if val, ok := entry["underlay_ip"]; ok {
		cfgLocation.UnderlayIP = val.(string)
	}
	cfg.DevLocation = &cfgLocation
	return nil
}

func setSystemInterface(cfg *swagger_models.DeviceConfig, d *schema.ResourceData) error {
	interfaceArray := rdEntryList(d, "interface")
	cfg.Interfaces = make([]*swagger_models.SysInterface, 0)
	for _, val := range interfaceArray {
		entry := val.(map[string]interface{})
		intf := swagger_models.SysInterface{}
		if val, ok := entry["cost"]; ok {
			intf.Cost = int64(val.(int))
		}
		if val, ok := entry["intf_usage"]; ok {
			usage := swagger_models.AdapterUsage(val.(string))
			intf.IntfUsage = &usage
		}
		if val, ok := entry["intfname"]; ok {
			intf.Intfname = val.(string)
		} else {
			return fmt.Errorf("intfname must be specified for Interfaces")
		}
		if val, ok := entry["ipaddr"]; ok {
			intf.Ipaddr = val.(string)
		}
		if val, ok := entry["macaddr"]; ok {
			intf.Macaddr = val.(string)
		}
		if val, ok := entry["netname"]; ok {
			intf.Netname = val.(string)
		}
		intf.Tags = rdEntryStrMap(entry, "tags")
		cfg.Interfaces = append(cfg.Interfaces, &intf)
	}
	return nil
}

func rdDeviceConfig(cfg *swagger_models.DeviceConfig, d *schema.ResourceData) error {
	cfg.AssetID = rdEntryStr(d, "asset_id")
	cfg.ClientIP = rdEntryStr(d, "client_ip")
	cfg.ClusterID = rdEntryStr(d, "cluster_id")
	cfg.Description = rdEntryStr(d, "description")
	eve_image_version := rdEntryStr(d, "eve_image_version")
	if eve_image_version == "" {
		return fmt.Errorf("eve_image_version must be specified.")
	}
	var err error
	cfg.ConfigItem, err = rdConfigItems(d)
	if err != nil {
		return err
	}
	err = setDeviceLocation(cfg, d)
	if err != nil {
		return err
	}
	err = setSystemInterface(cfg, d)
	if err != nil {
		return err
	}
	cfg.Tags = rdEntryStrMap(d, "tags")
	cfg.Title = rdEntryStrPtrOrNil(d, "title")
	return nil
}

// Update the Resource Group
func updateEdgeNodeResource(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	client := (meta.(Client)).Client
	id := rdEntryStr(d, "id")
	name := rdEntryStr(d, "name")
	errMsgPrefix := getErrMsgPrefix(name, id, "Edge Node", "Update")
	if client == nil {
		return diag.Errorf("%s nil Client", errMsgPrefix)
	}
	cfg, err := getEdgeNodeConfig(client, name, id)
	if err != nil {
		return diag.Errorf("%s Failed to find Edge Node. err: %s",
			errMsgPrefix, err.Error())
	}
	if cfg.ID != id {
		return diag.Errorf("%s ID of edge node changed. Expected: %s, actual: %s. "+
			"Object in zedcontrol is not same as expected by Terraform.",
			errMsgPrefix, id, cfg.ID)
	}
	err = rdDeviceConfig(cfg, d)
	if err != nil {
		return diag.Errorf("%s Error: %s", errMsgPrefix, err.Error())
	}
	_, err = edgeNodeSendPutReq(client, cfg, "update")
	if err != nil {
		return diag.Errorf("%s Update request failed. Err: %s",
			errMsgPrefix, err.Error())
	}
	err = edgeNodeUpdateBaseOs(client, cfg, rdEntryStr(d, "eve_image_version"))
	if err != nil {
		return diag.Errorf("%s %s", errMsgPrefix, err.Error())
	}
	log.Printf("[INFO] EdgeNode %s (ID: %s) Update Successful.", name, cfg.ID)
	return diags
}

// Delete the Resource Group
func deleteEdgeNodeResource(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	client := (meta.(Client)).Client
	name := rdEntryStr(d, "name")
	id := rdEntryStr(d, "id")
	errMsgPrefix := fmt.Sprintf("[ERROR] Edge Node %s ( id: %s) Delete Failed.",
		name, id)
	client.XRequestIdPrefix = "TF-edgenode-delete"
	urlExtension := getEdgeNodeUrl(name, id, "delete")
	rspData := &swagger_models.ZsrvResponse{}
	resp, err := client.SendReq("delete", urlExtension, nil, rspData)
	if err != nil {
		return diag.Errorf("%s. Err: %s", errMsgPrefix, err.Error())
	}
	if resp.StatusCode == 200 {
		log.Printf("[INFO] EdgeNode Delete Successful.")
		return diags
	}
	return diag.Errorf("%s. Status: %s\n resp: %++v", errMsgPrefix, resp.Status,
		resp)
}
