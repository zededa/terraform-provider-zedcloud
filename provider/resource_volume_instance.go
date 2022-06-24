// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/schemas"
	zedcloudapi "github.com/zededa/zedcloud-api"
	"github.com/zededa/zedcloud-api/swagger_models"
	"log"
)

var volumeInstanceUrlExtension = "volumes/instances"

var VolumeInstanceResourceSchema = &schema.Resource{
	CreateContext: createVolumeInstanceResource,
	ReadContext:   readResourceVolumeInstance,
	UpdateContext: updateVolumeInstanceResource,
	DeleteContext: deleteVolumeInstanceResource,
	Schema:        schemas.VolumeInstanceSchema,
	Importer: &schema.ResourceImporter{
		StateContext: schema.ImportStatePassthroughContext,
	},
}

// The schema for a resource group data source
func getVolumeInstanceResourceSchema() *schema.Resource {
	return VolumeInstanceResourceSchema
}

func getVolumeInstanceUrl(name, id, urlType string) string {
	return zedcloudapi.UrlForObjectRequest(volumeInstanceUrlExtension, name, id, urlType)
}

func rdUpdateVolumeInstanceCfg(cfg *swagger_models.VolInstConfig, d *schema.ResourceData) error {
	accessMode := swagger_models.VolumeInstanceAccessMode(
		rdEntryStr(d, "accessmode"))
	cfg.Accessmode = &accessMode
	cfg.Cleartext = rdEntryBool(d, "cleartext")
	cfg.ContentTreeID = rdEntryStr(d, "content_tree_id")
	cfg.Description = rdEntryStr(d, "description")
	cfg.DeviceID = rdEntryStr(d, "device_id")
	cfg.Image = rdEntryStr(d, "image")
	cfg.Label = rdEntryStr(d, "label")
	cfg.Multiattach = rdEntryBool(d, "multiattach")
	cfg.SizeBytes = rdEntryStr(d, "size_bytes")
	cfg.Title = rdEntryStr(d, "title")
	volType := swagger_models.VolumeInstanceType(rdEntryStr(d, "type"))
	cfg.Type = &volType
	return nil
}

// Create the Resource Group
func createVolumeInstanceResource(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	client := (meta.(Client)).Client
	name := rdEntryStr(d, "name")
	id := rdEntryStr(d, "id")
	errMsgPrefix := fmt.Sprintf("[ERROR] VolumeInstance %s (id: %s) Create Failed.",
		name, id)
	if client == nil {
		return diag.Errorf("%s nil Client", errMsgPrefix)
	}
	cfg := &swagger_models.VolInstConfig{
		Name:      name,
		ProjectID: rdEntryStr(d, "project_id"),
	}
	err := rdUpdateVolumeInstanceCfg(cfg, d)
	if err != nil {
		return diag.Errorf("%s %s", errMsgPrefix, err.Error())
	}
	log.Printf("[INFO] Creating VolumeInstance: %s", name)
	client.XRequestIdPrefix = "TF-volumeInstance-create"
	rspData := &swagger_models.ZsrvResponse{}
	_, err = client.SendReq("POST", volumeInstanceUrlExtension, cfg, rspData)
	if err != nil {
		return diag.Errorf("%s Err: %s", errMsgPrefix, err.Error())
	}
	log.Printf("VolumeInstance %s (ID: %s) Successfully created\n",
		rspData.ObjectName, rspData.ObjectID)
	id = rspData.ObjectID
	d.SetId(id)
	err = getVolumeInstanceAndPublishData(client, d, name, id, true)
	if err != nil {
		log.Printf("***[ERROR]- Failed to get VolumeInstance: %s (ID: %s) after "+
			"creating it. Err: %s", name, id, err.Error())
	}
	return diags
}

// Update the Resource Group
func updateVolumeInstanceResource(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	client := (meta.(Client)).Client
	name := rdEntryStr(d, "name")
	id := rdEntryStr(d, "id")
	errMsgPrefix := getErrMsgPrefix(name, id, "VolumeInstance", "Update")
	if client == nil {
		return diag.Errorf("%s nil Client", errMsgPrefix)
	}
	cfg, err, _ := getVolumeInstance(client, name, id)
	if err != nil {
		return diag.Errorf("%s err: %s", errMsgPrefix, err.Error())
	}
	err = checkInvalidAttrForUpdate(d, cfg.Name, cfg.ID)
	if err != nil {
		return diag.Errorf("%s err: %s", errMsgPrefix, err.Error())
	}
	log.Printf("[INFO] Updating VolumeInstance: %s (ID: %s)", name, cfg.ID)
	client.XRequestIdPrefix = "TF-volumeInstance-update"
	urlExtension := getVolumeInstanceUrl(name, id, "update")
	rspData := &swagger_models.ZsrvResponse{}
	_, err = client.SendReq("PUT", urlExtension, cfg, rspData)
	if err != nil {
		return diag.Errorf("%s Request Failed. err: %s", errMsgPrefix, err.Error())
	}
	err = getVolumeInstanceAndPublishData(client, d, name, id, true)
	if err != nil {
		log.Printf("***[ERROR]- Failed to get VolumeInstance: %s (ID: %s) after "+
			"updating it. Err: %s", name, id, err.Error())
	}
	return diags
}

// Delete the Resource Group
func deleteVolumeInstanceResource(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	client := (meta.(Client)).Client
	name := rdEntryStr(d, "name")
	id := rdEntryStr(d, "id")
	errMsgPrefix := getErrMsgPrefix(name, id, "VolumeInstance", "Delete")
	cfg, err, httpRsp := getVolumeInstance(client, name, id)
	if err != nil {
		if httpRsp != nil && zedcloudapi.IsObjectNotFound(httpRsp) {
			log.Printf("%s Not Found", errMsgPrefix)
			return diags
		}
		return diag.Errorf("%s Failed to get VolumeInstance. err: %s", errMsgPrefix, err.Error())
	}
	if cfg == nil {
		log.Printf("%s Unexpected Error. nil config", errMsgPrefix)
		return diags
	}
	client.XRequestIdPrefix = "TF-volumeInstance-delete"
	urlExtension := getVolumeInstanceUrl(name, id, "delete")
	rspData := &swagger_models.ZsrvResponse{}
	_, err = client.SendReq("DELETE", urlExtension, nil, rspData)
	if err != nil {
		return diag.Errorf("%s. Request Failed. err: %s", errMsgPrefix, err.Error())
	}
	log.Printf("[INFO] VolumeInstance %s(id:%s) Delete Successful.", name, cfg.ID)
	return diags
}

func readResourceVolumeInstance(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return readVolumeInstance(ctx, d, meta, true)
}
