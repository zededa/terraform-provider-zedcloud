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

var ImageDataSourceSchema = &schema.Resource{
	ReadContext: readImage,
	Schema:      zschemas.ImageSchema,
	Description: "Schema for data source zedcloud_image. Must specify id or name",
}

// The schema for a resource group data source
func getImageDataSourceSchema() *schema.Resource {
	return ImageDataSourceSchema
}

func getImage(client *zedcloudapi.Client,
	name, id string) (*swagger_models.ImageConfig, error) {
	rspData := &swagger_models.ImageConfig{}
	err := client.GetObj(imageUrlExtension, name, id, false, rspData)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] FAILED to get Image %s ( id: %s). Err: %s",
			name, id, err.Error())
	}
	return rspData, nil
}

func flattenImageConfig(cfg *swagger_models.ImageConfig) map[string]interface{} {
	return map[string]interface{}{
		"datastore_id":     ptrValStr(cfg.DatastoreID),
		"description":      cfg.Description,
		"id":               cfg.ID,
		"image_arch":       ptrValStr(cfg.ImageArch),
		"image_error":      cfg.ImageError,
		"image_format":     ptrValStr(cfg.ImageFormat),
		"image_local":      cfg.ImageLocal,
		"image_rel_url":    cfg.ImageRelURL,
		"image_sha_256":    cfg.ImageSha256,
		"image_size_bytes": cfg.ImageSizeBytes,
		"image_status":     ptrValStr(cfg.ImageStatus),
		"image_type":       ptrValStr(cfg.ImageType),
		"image_version":    cfg.ImageVersion,
		"name":             ptrValStr(cfg.Name),
		"origin_type":      ptrValStr(cfg.OriginType),
		"revision":         flattenObjectRevision(cfg.Revision),
		"title":            ptrValStr(cfg.Title),
	}
}

// Read the Resource Group
func readImage(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := (meta.(Client)).Client
	id := rdEntryStr(d, "id")
	name := rdEntryStr(d, "name")

	if client == nil {
		return diag.Errorf("nil Client.")
	}
	log.Printf("PROVIDER:  readImageResource - id: %s, name: %s", id, name)
	if (id == "") && (name == "") {
		return diag.Errorf("The arguments \"id\" or \"name\" are required, but no definition was found.")
	}
	cfg, err := getImage(client, name, id)
	if err != nil {
		return diag.Errorf("[ERROR] Image %s (id: %s) not found. Err: %s",
			name, id, err.Error())
	}

	// Take the Config and convert it to terraform object
	marshalData(d, flattenImageConfig(cfg))
	return diags
}
