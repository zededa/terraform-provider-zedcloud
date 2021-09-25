// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema for swagger_models.GeoLocation
var GeoLocationResourceSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"city": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "City",
		},
		"country": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Country code consisting of 2 capital letters as per ISO 3166-1 alpha2 standard",
		},
		"freeloc": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Free formatted location string",
		},
		"hostname": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Host name",
		},
		"loc": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Ordered pair of (latitude, longitude) separated by comma (,). " +
				"Latitude is the horizontal component used for geographic positioning. " +
				"It is the angle between 0° (the equator) and ±90° (north or south) at " +
				"the poles measured in decimal degrees. It is the first value in an " +
				"ordered pair. A negative number denotes a location south of the equator. " +
				"a positive number is north. Longitude is the vertical component used for " +
				" geographic positioning; it is the angle between 0° (the Prime Meridian) " +
				"and ±180° (westward or eastward) measured in decimal degrees. It is the " +
				"second number in an ordered pair. A negative number indicates a location " +
				"west of Greenwich, England; a positive number east.",
		},
		"org": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The name of the recipient, firm, or company at this geographical location.",
		},
		"postal": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Postal code (ZIP code for USA) of the geographical location",
		},
		"region": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Region",
		},
		"underlay_ip": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Single IP address, either in IPv4 or in IPv6 format",
		},
	},
}

var SysInterfaceResourceSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"cost": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Cost of using this interface. Default is 0. Maximum: 255",
		},
		"intfname": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Name of interface in the manifest to which this network or adapter maps to",
		},
		"intf_usage": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Adapter Usage. Valid values are: ADAPTER_USAGE_UNSPECIFIED, " +
				"ADAPTER_USAGE_MANAGEMENT, ADAPTER_USAGE_APP_DIRECT, " +
				"ADAPTER_USAGE_APP_SHARED, ADAPTER_USAGE_DISABLED ",
		},
		"ipaddr": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "IP address. Required if network is configured for static address",
		},
		"macaddr": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Mac address to be used for the interface. Could be over-written in some cases",
		},
		"netname": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Name of network object attached to the interface",
		},
		"tags": tagsSchema,
	},
}

// Schema Generated from swagger_models.DeviceConfig
var EdgeNodeSchema = map[string]*schema.Schema{
	// Keep the following common fields at the top of schema definitions for all
	//  objects.
	"name":        nameSchema,
	"id":          idSchema,
	"description": descriptionSchema,
	"title":       titleSchema,

	// Rest of the fields must be in the alphabetical order of keys
	"adminstate": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Administrative state of device",
	},
	"asset_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Device asset ID",
	},
	"client_ip": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Client IP",
	},
	"cluster_id": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "ID of the Cluster to which the Edge Node belongs.",
	},
	"config_items": {
		Type:        schema.TypeMap,
		Optional:    true,
		Description: "EVE Runtime Configuration Properties",
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},
	"cpu": {
		Type:        schema.TypeInt,
		Computed:    true,
		Description: "CPU (configured values)",
	},
	"dev_location": {
		Type:        schema.TypeList,
		Optional:    true,
		Description: "User specified geo location",
		Elem:        GeoLocationResourceSchema,
	},
	"eve_image_version": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Verion of EVE-OS image to be used by the Edge Node",
	},
	"interface": {
		Type:        schema.TypeSet,
		Optional:    true,
		Description: "System Interface Set",
		Elem:        SysInterfaceResourceSchema,
	},
	"memory": {
		Type:        schema.TypeInt,
		Computed:    true,
		Description: "Device memory in MBs",
	},
	"model_id": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "ID of device model object for the Edge Node",
	},
	// Project ID for the device cannot be changed. Since the device is created
	//  in ZEDCloud, set Computed=true for project_id. This is different from
	//  other objects.
	"project_id": &schema.Schema{
		Type:        schema.TypeString,
		Computed:    true,
		Description: "ID of the project to which the Object belongs",
	},
	"reset_counter": {
		Type:        schema.TypeInt,
		Computed:    true,
		Description: "Reset Counter Value",
	},
	"reset_time": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Edge Node Last Reset Time",
	},
	"revision": revisionSchema,
	"serialno": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Edge Node serial number",
	},
	"storage": {
		Type:        schema.TypeInt,
		Computed:    true,
		Description: "Device storage in GBs",
	},
	"tags": tagsSchema,
	"thread": {
		Type:        schema.TypeInt,
		Computed:    true,
		Description: "Threads",
	},
	"utype": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Device Model Archecture Type",
	},
}
