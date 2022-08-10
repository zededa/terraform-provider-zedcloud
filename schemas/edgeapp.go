// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var appACLMatchSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"type": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Type of Match. Allowed values are host, protocol, fport",
		},
		"value": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Value to match",
		},
	},
}

var appACLActionLimitParamsSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"limitburst": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Rate limit burst in ACL rule",
		},
		"limitrate": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Rate limit in ACL rule",
		},
		"limitunit": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Rate limit unit in ACL rule",
		},
	},
}

var appACLActionSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"drop": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Drop action on ACL rule",
		},
		"limit": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Limit action on ACL rule",
		},
		// limit_params -> LimitValue in API
		"limit_param": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Only valid if limit flag is set. Parameters for Limit action. ",
			Elem:        appACLActionLimitParamsSchema,
		},
		"portmap": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Enable device to app port mapping for incoming ACL " +
				"rule, implicitly added by ZedUI code",
		},
		// mapparams -> Portmapto in the API
		"mapparam": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     appMapParamsResourceSchema,
			MaxItems: 1,
			Description: "only valid if PortMap is set to true. " +
				"Device to app port mapping for incoming ACL rule",
		},
	},
}

var aclSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"action": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Chain of actions to be taken on matching network traffic",
			Elem:        appACLActionSchema,
		},
		"match": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     appACLMatchSchema,
			Description: "Network traffic matching criteria consistngs of one " +
				"or more of source IP address, destination IP address, protocol, " +
				"source port and destination port",
		},
		"name": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Name of the Access Control List",
		},
	},
}

var manifestInterfaceSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"acl": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     aclSchema,
			Description: "Traffic access control rules for this interface. " +
				"Applicable only when directattach flag is false.",
		},
		"directattach": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "If true, a physical adapter is assigned to the " +
				"edge application directly. If false, a network instance is " +
				"assigned to the edge application",
		},
		"name": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Interface name used by the edge application",
		},
		"optional": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Indicates if the interface is optional for edge application.",
		},
		"privateip": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "If true, DHCP network can't be assigned and user needs to provide a static IP address.",
		},
		"type": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Physical Adapter type for this interface. Applicable " +
				"only when directattach flag is true.",
		},
	},
}

var appCategoryValidValues = "" +
	"APP_CATEGORY_OPERATING_SYSTEM, " +
	"APP_CATEGORY_INDUSTRIAL, " +
	"APP_CATEGORY_EDGE_APPLICATION, " +
	"APP_CATEGORY_NETWORKING, " +
	"APP_CATEGORY_SECURITY, " +
	"APP_CATEGORY_DATA_ANALYTICS, " +
	"APP_CATEGORY_CLOUD_APPLICATION, " +
	"APP_CATEGORY_DEVOPS, " +
	"APP_CATEGORY_OTHERS, "

var manifestDescDetailsSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"agreement_list": {
			Type:     schema.TypeMap,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "Agreement List the the users of the App need to sign.",
		},
		"app_category": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Edge application category. The following are valid values: " +
				appCategoryValidValues,
		},
		"category": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Type of the Edge application",
		},
		"license_list": {
			Type:     schema.TypeMap,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "List of Licenses",
		},
		"logo": {
			Type:     schema.TypeMap,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "Application Logo",
		},
		"os": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Edge application's Operating System",
		},
		"support": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Support information for the Application",
		},
	},
}
var manifestUserDataTemplateSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"custom_config": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "User data template. ",
			Elem:        customConfigResourceSchema,
		},
	},
}

var manifestContainerDetailSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"container_create_option": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Base64 encoded container specific details. " +
				"Create options direct the creation of the Docker container",
		},
	},
}

var manifestImageSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"cleartext": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "cleartext",
		},
		"drvtype": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Type of Module. Valid Values are: CDROM, HDD, NET",
		},
		"ignorepurge": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Ignore Purge command for the App. ",
		},
		"imagename": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Name of the image",
		},
		"maxsize": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "maxsize of the image",
		},
		"mountpath": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "mountpath of the image inside the application instance",
		},
		"param": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "params for the Application",
			Elem:        nameValueEntrySchema,
		},
		"preserve": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Preserve the application image during purge",
		},
		"readonly": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "The image is readonly",
		},
		"target": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "enum: Disk, Kernel, Initrd, RamDisk",
		},
		"volumelabel": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "volumelabel",
		},
	},
}

var manifestModuleDetailSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"environment": {
			Type:     schema.TypeMap,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "Environment",
		},
		"module_type": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Type of Module. Valid Values are: " +
				"MODULE_TYPE_SYSTEM_DEFINED, MODULE_TYPE_CUSTOM",
		},
		"routes": {
			Type:     schema.TypeMap,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "Routes",
		},
		"twin_detail": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Base64 encoded module twin details, desired properties " +
				" of the module will be updated to reflect these values",
		},
	},
}
var manifestAuthorSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"company": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Company of the developer",
		},
		"email": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Contact email of the developer company",
		},
		"user": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Name of the developer",
		},
		"website": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Website of the developer company",
		},
	},
}

var appManifestSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"apptype": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "bundle type. Valid values are: APP_TYPE_VM, " +
				"APP_TYPE_VM_RUNTIME, APP_TYPE_CONTAINER, APP_TYPE_MODULE",
		},
		"configuration": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Custom Config for the APP to be provided for each instance",
			Elem:        manifestUserDataTemplateSchema,
		},
		"container_detail": {
			Type:     schema.TypeList,
			Optional: true,
			Description: "Container specific details. Create options direct the " +
				" creation of the Docker container",
			Elem: manifestContainerDetailSchema,
		},
		"deployment_type": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "type of deployment for the app. Valid Values: " +
				"DEPLOYMENT_TYPE_STAND_ALONE, DEPLOYMENT_TYPE_AZURE, " +
				"DEPLOYMENT_TYPE_K3S, DEPLOYMENT_TYPE_AWS, DEPLOYMENT_TYPE_K3S_AZURE" +
				"DEPLOYMENT_TYPE_K3S_AWS",
		},
		"desc_detail": {
			Type:     schema.TypeList,
			Optional: true,
			Description: "Meta data for the App for the Zededa App Marketplace " +
				"like agreements, category etc",
			Elem: manifestDescDetailsSchema,
		},
		"description": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Description of the App",
		},
		"display_name": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Display name of the App",
		},
		"enablevnc": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Enable VNC for the app",
		},
		"image": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Images used by the app",
			Elem:        manifestImageSchema,
		},
		"interface": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "I/O adapter settings",
			Elem:        manifestInterfaceSchema,
		},
		"module": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Azure module specific details like module twin, environment variable, routes",
			Elem:        manifestModuleDetailSchema,
		},
		"name": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Unique name of app manifest, should match object name",
		},
		"owner": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Information about Author of the App",
			Elem:        manifestAuthorSchema,
		},
		"resource": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Hardware resource requirement (CPU, Memory, Storage) for the app",
			Elem:        nameValueEntrySchema,
		},
		"vmmode": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "VM mode for VM-based app",
		},
		"cpu_pinning_enabled": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Enable CpuPinning",
		},
	},
}

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
	"manifest": {
		Type:     schema.TypeList,
		Optional: true,
		Description: "App Manifest definition. Only one of manifest or " +
			"manifest_file must be specified.",
		Elem: appManifestSchema,
	},
	"manifest_file": {
		Type:     schema.TypeString,
		Optional: true,
		Description: "Location of Edge Application Manifest file (JSON format). " +
			"Only one of manifest or manifest_file must be specified.",
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
