// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	tagsDescription = "Tags are name/value pairs that enable you to categorize resources. " +
		"Tag names are case insensitive with max_length 512 and min_length 3. Tag " +
		"values are case sensitive with max_length 256 and min_length 3"
)

var description = &schema.Schema{
	Type:        schema.TypeString,
	Optional:    true,
	Description: "Detailed description of the Object",
}

var id = &schema.Schema{
	Type:        schema.TypeString,
	Computed:    true,
	Description: "System defined unique ID of the Object",
}

var name = &schema.Schema{
	Type:     schema.TypeString,
	Required: true,
	Description: "User defined name of the object. Must be unique across the enterprise. " +
		"Once object is created, name can’t be changed",
}

var projectID = &schema.Schema{
	Type:        schema.TypeString,
	Optional:    true,
	Description: "ID of the project to which the Object belongs",
	DiffSuppressFunc: func(k, oldVal, newVal string, d *schema.ResourceData) bool {
		if newVal == "" {
			// If new is not specified - old one was published by get.
			// This is the case of .tf not having the project id for the object,
			// but zedcloud setting it - may be default project id.
			return true
		}
		return false
	},
}

var projectIDComputed = &schema.Schema{
	Type:        schema.TypeString,
	Computed:    true,
	Description: "ID of the project to which the Object belongs",
}

var tags = &schema.Schema{
	Type:     schema.TypeMap,
	Optional: true,
	Elem: &schema.Schema{
		Type: schema.TypeString,
	},
	Description: tagsDescription,
}

var revision = &schema.Schema{
	Type:        schema.TypeList,
	Computed:    true,
	Description: "System defined revision information of the object",
	Elem:        objectRevision,
}

var title = &schema.Schema{
	Type:        schema.TypeString,
	Optional:    true,
	Description: "User defined title of the object. title can be changed any time.",
}

var zedcloudOpsCmd = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"counter": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Operation Counter",
		},
		"ops_time": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Timestamp: Operational time",
		},
	},
}

var objectParentDetail = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"id_of_parent": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "ID of parent object",
		},
		"reference_exists": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Relation with parent object exists or not",
		},
		"update_available": {
			Type:     schema.TypeBool,
			Computed: true,
			Description: "Flag to indicate if Update required for Child object " +
				"due to an update to parent object",
		},
		"version_of_parent_object": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Version of Parent Object currently used by the child object",
		},
	},
}

var objectRevision = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"created_at": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The time, in milliseconds since the epoch, when the object was created",
		},
		"created_by": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "User who Created the object",
		},
		"curr": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Current version of the object",
		},
		"prev": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Prev version of the object",
		},
		"updated_at": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The time, in milliseconds since the epoch, when the object was last updated",
		},
		"updated_by": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "User who last updated the object",
		},
	},
}

var nameValueEntry = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Name of the entry",
		},
		"value": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Value of the entry",
		},
	},
}
