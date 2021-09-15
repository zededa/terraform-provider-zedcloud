// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var EdgeAppSchema = map[string]*schema.Schema{
	// Keep the following common fields at the top of schema definitions for all
	//  objects.
	"name":        nameSchema,
	"id":          idSchema,
	"description": descriptionSchema,
	"title":       titleSchema,

	// Rest of the fields must be in the alphabetical order of keys
	"cpus": {
		Type:        schema.TypeInt,
		Optional:    true,
		Description: "Number of CPUs to be assigned to an App Instance of this app",
	},
	"drives": {
		Type:        schema.TypeInt,
		Optional:    true,
		Description: "Number of drives required for an App Instance of this App",
	},
	"manifest_file": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Location of Edge Application Manifest file (JSON format)",
	},
	"memory": {
		Type:        schema.TypeInt,
		Optional:    true,
		Description: "Memory to be assigned to an App Instance of this App",
	},
	"networks": {
		Type:        schema.TypeInt,
		Optional:    true,
		Description: "Number of Networks required by the App",
	},
	"origin_type": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Origin of object",
	},
	"parent_detail": {
		Type:        schema.TypeList,
		Computed:    true,
		Description: "Details of Parent Object",
		Elem:        objectParentDetailSchema,
	},
	"revision": revisionSchema,
	"storage": {
		Type:        schema.TypeInt,
		Optional:    true,
		Description: "Amount of Storage required for App Instance of this App",
	},
	"user_defined_version": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "User defined version of the edge-app",
	},
}
