// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package resources

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	zschemas "github.com/zededa/terraform-provider-zedcloud/schemas"
	zedcloudapi "github.com/zededa/zedcloud-api"
	"github.com/zededa/zedcloud-api/swagger_models"
)

var deviceUrlExtension = "devices"

var EdgeNodeResourceSchema = &schema.Resource{
	CreateContext: createEdgeNodeResource,
	ReadContext:   readResourceEdgeNode,
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
	_, err := client.GetObj(deviceUrlExtension, name, id, false, rspData)
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

func adminStatePtr(strVal string) *swagger_models.AdminState {
	val := swagger_models.AdminState(strVal)
	return &val
}

func edgeNodeUpdateAdminState(client *zedcloudapi.Client, d *schema.ResourceData,
	id string) error {
	cfg, err := getEdgeNodeConfig(client, "", id)
	if err != nil {
		return fmt.Errorf("Failed to find Edge Node to set Admin state. err: %s",
			err.Error())
	}
	var action string
	action, err = setAdminState(cfg, d)
	if err != nil {
		return fmt.Errorf("Failed to set Admin state in Config. err: %s", err.Error())
	}
	if action == "" {
		return nil
	}
	_, err = edgeNodeSendPutReq(client, cfg, action)
	if err != nil {
		return fmt.Errorf("Failed to set Admin state. Err: %s", err.Error())
	}
	return nil
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
		return fmt.Errorf("BaseOsImage Publish failed. eve_image_version: %s, Err: %s",
			eve_image_version, err.Error())
	}
	// We need to apply the configuration
	rsp, err := edgeNodeSendPutReq(client, cfg, "apply")
	// Ignore HTTP_STATUS_CONFLICT
	if err != nil && rsp.StatusCode != http.StatusConflict {
		return fmt.Errorf("BaseOsUpdate Apply request failed. eve_image_version: %s, "+
			"Err: %s.", eve_image_version, err.Error())
	}
	return nil
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
	locationInfo := getList(d, "dev_location")
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

func adapterUsagePtr(strVal string) *swagger_models.AdapterUsage {
	val := swagger_models.AdapterUsage(strVal)
	return &val
}

func setSystemInterface(cfg *swagger_models.DeviceConfig, d *schema.ResourceData) error {
	interfaceArray := getList(d, "interface")
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
		intf.Tags = getStrMap(entry, "tags")
		cfg.Interfaces = append(cfg.Interfaces, &intf)
	}
	return nil
}

func checkAdminStateValue(d *schema.ResourceData) (string, error) {
	strVal := getStr(d, "adminstate_config")
	switch strVal {
	case "ADMIN_STATE_ACTIVE":
		return "activate", nil
	case "ADMIN_STATE_INACTIVE":
		return "deactivate", nil
	default:
		return "", fmt.Errorf("adminstate_config must be specified and be one of " +
			"ADMIN_STATE_ACTIVE, ADMIN_STATE_INACTIVE")
	}
}

func setAdminState(cfg *swagger_models.DeviceConfig,
	d *schema.ResourceData) (string, error) {
	action, err := checkAdminStateValue(d)
	if err != nil {
		return "", err
	}
	strVal := getStr(d, "adminstate_config")
	adminstate := swagger_models.AdminState(strVal)
	if cfg.AdminState != nil {
		if adminstate == *cfg.AdminState {
			// Admin State same as configured value - no action needed.
			return "", nil
		}
		// If adminstate_config is ACTIVE and the device is already in
		// Registered state. No action needed.
		if strVal == "ADMIN_STATE_ACTIVE" &&
			*cfg.AdminState == swagger_models.AdminStateADMINSTATEREGISTERED {
			return "", nil
		}
	}
	cfg.AdminState = &adminstate
	return action, nil
}

func rdDeviceConfig(cfg *swagger_models.DeviceConfig, d *schema.ResourceData, create bool) error {
	var err error
	cfg.Description = getStr(d, "description")
	cfg.Title = getStrPtrOrNil(d, "title")

	_, err = checkAdminStateValue(d)
	if err != nil {
		return err
	}
	cfg.AssetID = getStr(d, "asset_id")
	cfg.ClientIP = getStr(d, "client_ip")
	cfg.ClusterID = getStr(d, "cluster_id")
	cfg.ConfigItem, err = rdConfigItems(d)
	if err != nil {
		return err
	}
	err = setDeviceLocation(cfg, d)
	if err != nil {
		return err
	}
	eve_image_version := getStr(d, "eve_image_version")
	if eve_image_version == "" {
		return fmt.Errorf("eve_image_version must be specified.")
	}
	err = setSystemInterface(cfg, d)
	if err != nil {
		return err
	}
	if create {
		// model id and project id will be set only during create.
		cfg.ModelID = getStrPtrOrNil(d, "model_id")
		if cfg.ModelID == nil || *cfg.ModelID == "" {
			return fmt.Errorf("model_id must be specified for the EdgeNode.")
		}
		cfg.Obkey = getStr(d, "onboard_key")
	}
	projectIdPtr := getStrPtrOrNil(d, "project_id")
	if create {
		if projectIdPtr == nil || *projectIdPtr == "" {
			return fmt.Errorf("project_id must be specified for the EdgeNode.")
		}
		cfg.ProjectID = projectIdPtr
	} else {
		// Change of Project for a device is NOT supported with Update operation.
		// Check if the configured project is different from the config we got
		// from Cloud. If different - error out.
		// If project_id is not configured, ignore the check. We cannot verify
		// the case of device created with project_id set to non-default project,
		// but changed to remove set the project_id to empty, though this also is
		// not supported. Such an update request would be rejected by ZEDCloud.
		if projectIdPtr != nil && cfg.ProjectID != nil &&
			*cfg.ProjectID != *projectIdPtr {
			// Update. Project cannot be changed
			return fmt.Errorf("project_id cannot be changed after EdgeNode is "+
				"created. Current: %s, New: %s", *cfg.ProjectID, *projectIdPtr)
		}
	}
	cfg.Serialno = getStr(d, "serialno")
	cfg.Tags = getStrMap(d, "tags")
	return nil
}

// Create the Edge Node
func createEdgeNodeResource(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := (meta.(Client)).Client
	name := getStr(d, "name")
	id := getStr(d, "id")
	errMsgPrefix := getErrMsgPrefix(name, id, "EdgeNode", "Create")
	if client == nil {
		return diag.Errorf("%s err: %s", errMsgPrefix, "nil Client")
	}
	cfg := &swagger_models.DeviceConfig{
		Name: &name,
	}
	err := rdDeviceConfig(cfg, d, true)
	if err != nil {
		return diag.Errorf("%s Error: %s", errMsgPrefix, err.Error())
	}

	log.Printf("[INFO] Creating EdgeNode: %s", name)
	client.XRequestIdPrefix = "TF-edgenode-create"
	rspData := &swagger_models.ZsrvResponse{}
	_, err = client.SendReq("POST", deviceUrlExtension, cfg, rspData)
	if err != nil {
		return diag.Errorf("%s err: %s", errMsgPrefix, err.Error())
	}
	id = rspData.ObjectID
	log.Printf("EdgeNode %s (ID: %s) Successfully created\n", rspData.ObjectName, id)
	d.SetId(id)

	// Get Edgenode Config and update BaseOs.
	cfg, err = getEdgeNodeConfig(client, name, id)
	if err != nil {
		return diag.Errorf("%s Failed to find Edge Node. err: %s",
			errMsgPrefix, err.Error())
	}
	err = edgeNodeUpdateBaseOs(client, cfg, getStr(d, "eve_image_version"))
	if err != nil {
		return diag.Errorf("%s %s", errMsgPrefix, err.Error())
	}

	// Activating / De-Activating the device requires activate / deactivate call
	err = edgeNodeUpdateAdminState(client, d, id)
	if err != nil {
		return diag.Errorf("%s Failed to update Admin state. Err: %s",
			errMsgPrefix, err.Error())
	}

	// Get Edge node config and publish the latest version. This is mainly to
	// published the computed fields. Object rev. changes for every update
	err = getEdgeNodeAndPublishData(client, d, name, id)
	if err != nil {
		log.Printf("***[ERROR]- Failed to get EdgeNode: %s (ID: %s) after "+
			"creating it. Err: %s", name, id, err.Error())
	}
	return diags
}

// Update the Resource Group
func updateEdgeNodeResource(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := (meta.(Client)).Client
	id := getStr(d, "id")
	name := getStr(d, "name")
	errMsgPrefix := getErrMsgPrefix(name, id, "Edge Node", "Update")
	if client == nil {
		return diag.Errorf("%s nil Client", errMsgPrefix)
	}
	cfg, err := getEdgeNodeConfig(client, name, id)
	if err != nil {
		return diag.Errorf("%s Failed to find Edge Node. err: %s",
			errMsgPrefix, err.Error())
	}
	err = checkInvalidAttrForUpdate(d, *cfg.Name, cfg.ID)
	if err != nil {
		return diag.Errorf("%s err: %s", errMsgPrefix, err.Error())
	}
	if cfg.ID != id {
		return diag.Errorf("%s ID of edge node changed. Expected: %s, actual: %s. "+
			"Object in zedcontrol is not same as expected by Terraform.",
			errMsgPrefix, id, cfg.ID)
	}
	err = rdDeviceConfig(cfg, d, false)
	if err != nil {
		return diag.Errorf("%s Error: %s", errMsgPrefix, err.Error())
	}
	_, err = edgeNodeSendPutReq(client, cfg, "update")
	if err != nil {
		return diag.Errorf("%s Update request failed. Err: %s",
			errMsgPrefix, err.Error())
	}
	err = edgeNodeUpdateBaseOs(client, cfg, getStr(d, "eve_image_version"))
	if err != nil {
		return diag.Errorf("%s %s", errMsgPrefix, err.Error())
	}
	err = edgeNodeUpdateAdminState(client, d, id)
	if err != nil {
		return diag.Errorf("%s %s", errMsgPrefix, err.Error())
	}
	err = getEdgeNodeAndPublishData(client, d, name, id)
	if err != nil {
		return diag.Errorf("%s", err.Error())
	}
	log.Printf("[INFO] EdgeNode %s (ID: %s) Update Successful.", name, cfg.ID)
	return diags
}

// Delete the Resource Group
func deleteEdgeNodeResource(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := (meta.(Client)).Client
	name := getStr(d, "name")
	id := getStr(d, "id")
	errMsgPrefix := getErrMsgPrefix(name, id, "Edge Node", "Delete")
	client.XRequestIdPrefix = "TF-edgenode-delete"
	urlExtension := getEdgeNodeUrl(name, id, "delete")
	rspData := &swagger_models.ZsrvResponse{}
	_, err := client.SendReq("DELETE", urlExtension, nil, rspData)
	if err != nil {
		return diag.Errorf("%s. Err: %s", errMsgPrefix, err.Error())
	}
	log.Printf("[INFO] EdgeNode Delete Successful.")
	return diags
}

func readResourceEdgeNode(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return readEdgeNode(ctx, d, meta)
}
