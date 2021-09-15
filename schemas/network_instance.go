// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var netInstOpaqueConfigSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"oconfig": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Base64 encoded string of opaque config",
		},
		"type": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Type of Opaque config. Valid Values: OPAQUE_CONFIG_TYPE_VPN",
		},
	},
}

var staticDNSListSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"addrs": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Addresses",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"hostname": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Host name",
		},
	},
}

var dhcpIpRangeSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"end": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "End of IP Address range",
		},
		"start": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Start of IP Address range",
		},
	},
}

var dhcpServerConfigSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"dhcp_range": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Range of IP addresses to be used for DHCP",
			MaxItems:    1,
			Elem:        dhcpIpRangeSchema,
		},
		"dns": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "List of IP Addresses of DNS servers",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"domain": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Network domain",
		},
		"gateway": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "IP Address of Network Gateway",
		},
		"mask": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Subnet Mask",
		},
		"ntp": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "IP Address of NTP Server",
		},
		"subnet": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Subnet address",
		},
	},
}

var NetworkInstanceSchema = map[string]*schema.Schema{
	// Keep the following common fields at the top of schema definitions for all
	//  objects.
	"name":        nameSchema,
	"id":          idSchema,
	"description": descriptionSchema,
	"title":       titleSchema,

	// Rest of the fields must be in the alphabetical order of keys
	"cluster_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "ID of the Cluster in which the network instance is configured",
	},
	"device_default": {
		Type:        schema.TypeBool,
		Default:     false,
		Optional:    true,
		Description: "Flag to indicate if this is the default network instance for the device",
	},
	"device_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "ID of the device on which network instance is created",
	},
	"dhcp": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Flag to enable / disable dhcp on this network instance",
	},
	"dns_list": {
		Type:        schema.TypeList,
		Optional:    true,
		Description: "List of Static DNS entries",
		Elem:        staticDNSListSchema,
	},
	"ip": {
		Type:        schema.TypeList,
		Optional:    true,
		Description: "Dhcp Server Configuration",
		Elem:        dhcpServerConfigSchema,
	},
	"kind": {
		Type:     schema.TypeString,
		Optional: true,
		Description: "Kind of Network Instance. Valid Values: NETWORK_INSTANCE_KIND_UNSPECIFIED, " +
			"NETWORK_INSTANCE_KIND_TRANSPARENT, NETWORK_INSTANCE_KIND_SWITCH, NETWORK_INSTANCE_KIND_LOCAL " +
			"NETWORK_INSTANCE_KIND_CLOUD, NETWORK_INSTANCE_KIND_MESH, NETWORK_INSTANCE_KIND_HONEYPOT",
	},
	"network_policy_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "ID of the network policy to be attached to this network instance",
	},
	"opaque": {
		Type:        schema.TypeList,
		Optional:    true,
		Description: "Service specific Config",
		MaxItems:    1,
		Elem:        netInstOpaqueConfigSchema,
	},
	"port": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Name of port mapping in the model",
	},
	"port_tags": {
		Type:     schema.TypeMap,
		Optional: true,
		Description: "Port Tags are name/value pairs that enable you to categorize " +
			"resources. Tag names are case insensitive with max_length 512 and min_length 3. " +
			" Tag values are case sensitive with max_length 256 and min_length 3.",
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},
	"project_id": projectIdSchema,
	"revision":   revisionSchema,
	"tags":       tagsSchema,
	"type": {
		Type:     schema.TypeString,
		Optional: true,
		Description: "Network Instance DHCP type. Valid Values: NETWORK_INSTANCE_DHCP_TYPE_V4, " +
			"NETWORK_INSTANCE_DHCP_TYPE_V6, NETWORK_INSTANCE_DHCP_TYPE_CRYPTOEID, " +
			"NETWORK_INSTANCE_DHCP_TYPE_CRYPTOV4, NETWORK_INSTANCE_DHCP_TYPE_CRYPTOV6",
	},
}
