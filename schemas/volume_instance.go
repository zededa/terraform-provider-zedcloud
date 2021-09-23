// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strings"
)

func volInstTypeValuesStr() string {
	values := []string{
		"VOLUME_INSTANCE_TYPE_UNSPECIFIED",
		"VOLUME_INSTANCE_TYPE_EMPTYDIR",
		"VOLUME_INSTANCE_TYPE_BLOCKSTORAGE",
		"VOLUME_INSTANCE_TYPE_HOSTFS",
		"VOLUME_INSTANCE_TYPE_TMPFS",
		"VOLUME_INSTANCE_TYPE_SECRET",
		"VOLUME_INSTANCE_TYPE_NFS",
		"VOLUME_INSTANCE_TYPE_AWS_BLOCK_STORAGE",
		"VOLUME_INSTANCE_TYPE_CONTENT_TREE"}
	return strings.Join(values, ", ")
}

func volInstAccessModeValuesStr() string {
	values := []string{
		"VOLUME_INSTANCE_ACCESS_MODE_INVALID",
		"VOLUME_INSTANCE_ACCESS_MODE_READWRITE",
		"VOLUME_INSTANCE_ACCESS_MODE_READONLY",
		"VOLUME_INSTANCE_ACCESS_MODE_MULTIREAD_SINGLEWRITE",
	}
	return strings.Join(values, ", ")
}

var VolumeInstanceSchema = map[string]*schema.Schema{
	// Keep the following common fields at the top of schema definitions for all
	//  objects.
	"name":        nameSchema,
	"id":          idSchema,
	"description": descriptionSchema,
	"title":       titleSchema,

	// Rest of the fields must be in the alphabetical order of keys
	"accessmode": {
		Type:     schema.TypeString,
		Optional: true,
		Description: "Access mode. Valid Values: VOLUME_INSTANCE_ACCESS_MODE_INVALID, " +
			"VOLUME_INSTANCE_ACCESS_MODE_READWRITE, VOLUME_INSTANCE_ACCESS_MODE_READONLY, " +
			"VOLUME_INSTANCE_ACCESS_MODE_MULTIREAD_SINGLEWRITE",
	},
	"cleartext": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Flag to keep the contents of the volume unencrypted (in clear text)",
	},
	"content_tree_id": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Content tree ID",
	},
	"device_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "ID of the device on which volume instance is created",
	},
	"image": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Name of the image",
	},
	"implicit": {
		Type:        schema.TypeBool,
		Computed:    true,
		Description: "Flag to create implicit volumes",
	},
	"label": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Label",
	},
	"multiattach": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Flag to enable the volume to be attached to multiple app instances",
	},
	"project_id": projectIdSchema,
	"purge": {
		Type:        schema.TypeList,
		Computed:    true,
		Elem:        zedcloudOpsCmdSchema,
		Description: "Purge Counter information",
	},
	"revision": revisionSchema,
	"size_bytes": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Size of volume",
	},
	"type": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Type of Volume Instance. Valid Values: " + volInstTypeValuesStr(),
	},
}
