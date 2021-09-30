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

var imageUrlExtension = "apps/images"

var ImageResourceSchema = &schema.Resource{
	CreateContext: createImageResource,
	ReadContext:   readImage,
	UpdateContext: updateImageResource,
	DeleteContext: deleteImageResource,
	Schema:        schemas.ImageSchema,
	Importer: &schema.ResourceImporter{
		StateContext: schema.ImportStatePassthroughContext,
	},
}

// The schema for a resource group data source
func getImageResourceSchema() *schema.Resource {
	return ImageResourceSchema
}

func getImageUrl(name, id, urlType string) string {
	return zedcloudapi.UrlForObjectRequest(imageUrlExtension, name, id, urlType)
}

func updateImageCfgFromResourceData(cfg *swagger_models.ImageConfig, d *schema.ResourceData) error {

	cfg.DatastoreID = rdEntryStrPtrOrNil(d, "datastore_id")
	if cfg.DatastoreID == nil {
		return fmt.Errorf("datastore_id must be specified. It cannot be an empty string.")
	}
	cfg.Description = rdEntryStr(d, "description")

	imageArch := swagger_models.ModelArchType(rdEntryStr(d, "image_arch"))
	if imageArch == "" {
		return fmt.Errorf("image_arch must be specified. It cannot be an empty string.")
	}
	cfg.ImageArch = &imageArch

	imageFormat := swagger_models.ConfigFormat(rdEntryStr(d, "image_format"))
	if imageFormat == "" {
		return fmt.Errorf("image_format must be specified. It cannot be an empty string.")
	}
	cfg.ImageFormat = &imageFormat

	cfg.ImageRelURL = rdEntryStr(d, "image_rel_url")
	cfg.ImageSha256 = rdEntryStr(d, "image_sha_256")
	cfg.ImageSizeBytes = rdEntryStr(d, "image_size_bytes")
	imageType := swagger_models.ImageType(rdEntryStr(d, "image_type"))
	cfg.ImageType = &imageType

	cfg.Title = rdEntryStrPtrOrNil(d, "title")
	if cfg.Title == nil {
		return fmt.Errorf("title must be specified. It cannot be an empty string.")
	}
	return nil
}

// Create the Resource Group
func createImageResource(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	client := (meta.(Client)).Client
	name := rdEntryStr(d, "name")
	id := rdEntryStr(d, "id")
	errMsgPrefix := fmt.Sprintf("[ERROR] Image %s (id: %s) Create Failed.",
		name, id)
	if client == nil {
		return diag.Errorf("%s nil Client", errMsgPrefix)
	}
	cfg := &swagger_models.ImageConfig{
		Name: &name,
	}
	err := updateImageCfgFromResourceData(cfg, d)
	if err != nil {
		return diag.Errorf("%s %s", errMsgPrefix, err.Error())
	}
	log.Printf("[INFO] Creating Image: %s", name)
	client.XRequestIdPrefix = "TF-image-create"
	rspData := &swagger_models.ZsrvResponse{}
	_, err = client.SendReq("POST", imageUrlExtension, cfg, rspData)
	if err != nil {
		return diag.Errorf("%s Err: %s", errMsgPrefix, err.Error())
	}
	log.Printf("Image %s (ID: %s) Successfully created\n",
		rspData.ObjectName, rspData.ObjectID)
	d.SetId(rspData.ObjectID)
	return diags
}

// Update the Resource Group
func updateImageResource(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	client := (meta.(Client)).Client
	name := rdEntryStr(d, "name")
	id := rdEntryStr(d, "id")
	errMsgPrefix := getErrMsgPrefix(name, id, "Image", "Update")
	if client == nil {
		return diag.Errorf("%s nil Client", errMsgPrefix)
	}
	cfg, err := getImage(client, name, id)
	if err != nil {
		return diag.Errorf("%s err: %s", errMsgPrefix, err.Error())
	}
	err = checkInvalidAttrForUpdate(d, *cfg.Name, cfg.ID)
	if err != nil {
		return diag.Errorf("%s err: %s", errMsgPrefix, err.Error())
	}
	log.Printf("[INFO] Updating Image: %s (ID: %s)", name, cfg.ID)
	err = updateImageCfgFromResourceData(cfg, d)
	if err != nil {
		return diag.Errorf("%s err: %s", errMsgPrefix, err.Error())
	}
	client.XRequestIdPrefix = "TF-image-update"
	urlExtension := getImageUrl(name, id, "update")
	rspData := &swagger_models.ZsrvResponse{}
	_, err = client.SendReq("PUT", urlExtension, cfg, rspData)
	if err != nil {
		return diag.Errorf("%s Request Failed. err: %s", errMsgPrefix, err.Error())
	}
	return diags
}

// Delete the Resource Group
func deleteImageResource(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	client := (meta.(Client)).Client
	name := rdEntryStr(d, "name")
	id := rdEntryStr(d, "id")
	errMsgPrefix := getErrMsgPrefix(name, id, "Image", "Delete")
	cfg, err := getImage(client, name, id)
	if err != nil {
		return diag.Errorf("%s Failed to get Image. err: %s", errMsgPrefix, err.Error())
	}
	if cfg == nil {
		log.Printf("%s Unexpected Error. nil config", errMsgPrefix)
		return diags
	}
	client.XRequestIdPrefix = "TF-image-delete"
	urlExtension := getImageUrl(name, id, "delete")
	rspData := &swagger_models.ZsrvResponse{}
	_, err = client.SendReq("DELETE", urlExtension, nil, rspData)
	if err != nil {
		return diag.Errorf("%s. Request Failed. err: %s", errMsgPrefix, err.Error())
	}
	log.Printf("[INFO] Image %s(id:%s) Delete Successful.", name, cfg.ID)
	return diags
}
