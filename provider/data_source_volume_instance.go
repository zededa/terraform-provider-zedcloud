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

var VolumeInstanceDataSourceSchema = &schema.Resource{
	ReadContext: readVolumeInstance,
	Schema:      zschemas.VolumeInstanceSchema,
	Description: "Schema for data source zedcloud_volume_instance. Must specify " +
		"id or name",
}

// The schema for a resource group data source
func getVolumeInstanceDataSourceSchema() *schema.Resource {
	return VolumeInstanceDataSourceSchema
}

func getVolumeInstance(client *zedcloudapi.Client,
	name, id string) (*swagger_models.VolInstConfig, error) {
	rspData := &swagger_models.VolInstConfig{}
	err := client.GetObj(volumeInstanceUrlExtension, name, id, false, rspData)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] FAILED to get VolumeInstance %s ( id: %s). Err: %s",
			name, id, err.Error())
	}
	return rspData, nil
}

func flattenVolInstConfig(cfg *swagger_models.VolInstConfig) map[string]interface{} {
	if cfg == nil {
		return map[string]interface{}{}
	}
	return map[string]interface{}{
		"accessmode":      ptrValStr(cfg.Accessmode),
		"cleartext":       cfg.Cleartext,
		"content_tree_id": cfg.ContentTreeID,
		"description":     cfg.Description,
		"device_id":       cfg.DeviceID,
		"id":              cfg.ID,
		"image":           cfg.Image,
		"implicit":        cfg.Implicit,
		"label":           cfg.Label,
		"multiattach":     cfg.Multiattach,
		"name":            cfg.Name,
		"project_id":      cfg.ProjectID,
		"purge":           flattenZedCloudOpsCmd(cfg.Purge),
		"revision":        flattenObjectRevision(cfg.Revision),
		"size_bytes":      cfg.SizeBytes,
		"title":           cfg.Title,
		"type":            ptrValStr(cfg.Type),
	}
}

// Read the Resource Group
func readVolumeInstance(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := (meta.(Client)).Client
	id := rdEntryStr(d, "id")
	name := rdEntryStr(d, "name")

	if client == nil {
		return diag.Errorf("nil Client.")
	}
	log.Printf("PROVIDER:  readVolumeInstanceResource - id: %s, name: %s", id, name)
	if (id == "") && (name == "") {
		return diag.Errorf("The arguments \"id\" or \"name\" are required, but no definition was found.")
	}
	cfg, err := getVolumeInstance(client, name, id)
	if err != nil {
		return diag.Errorf("[ERROR] VolumeInstance %s (id: %s) not found. Err: %s",
			name, id, err.Error())
	}
	marshalData(d, flattenVolInstConfig(cfg))
	return diags
}
