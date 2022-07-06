// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var ImageSchema = map[string]*schema.Schema{
	// Keep the following common fields at the top of schema definitions for all
	//  objects.
	"name":        nameSchema,
	"id":          idSchema,
	"description": descriptionSchema,
	"title":       titleSchema,

	// Rest of the fields must be in the alphabetical order of keys
	"datastore_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Datastore Id where image binary is located",
	},
	"image_arch": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Image Architecture. Valid Values: AMD64, ARM64",
	},
	"image_error": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Image upload/uplink detailed error/status message",
	},
	"image_format": {
		Type:     schema.TypeString,
		Optional: true,
		Description: "Image binary format. Valid Values: FmtUnknown, RAW, QCOW, QCOW2, " +
			"VHD, VMDK, OVA, VHDX, CONTAINER",
	},
	"image_local": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Internal image location.",
	},
	"image_rel_url": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Image relative path w.r.t. Datastore",
	},
	"image_sha_256": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Image checksum in SHA256 format",
	},
	"image_size_bytes": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Image size in KBytes.",
		Default:     "0",
	},
	"image_status": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Image status",
	},
	"image_type": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "IMAGE_TYPE_EVE, IMAGE_TYPE_APPLICATION, IMAGE_TYPE_EVEPRIVATE",
	},
	"image_version": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "System defined Version of the Object",
	},
	"origin_type": {
		Type:     schema.TypeString,
		Computed: true,
		Description: "Origin type of image, Valid Values: ORIGIN_UNSPECIFIED, " +
			"ORIGIN_IMPORTED, ORIGIN_LOCAL, ORIGIN_GLOBAL",
	},
	"revision": revisionSchema,
}
