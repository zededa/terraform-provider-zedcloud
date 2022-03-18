// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var variableOptionValResourceSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"label": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Label of option",
		},
		"value": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Value of option",
		},
	},
}

var variableGroupVariableResourceSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"default": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Default value for the Variable Group Variable",
		},
		"encode": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Encoding Type. Valid Values: FILE_ENCODING_UNSPECIFIED, FILE_ENCODING_BASE64",
		},
		"format": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Format of the Variable. Valid Values: VARIABLE_FORMAT_UNSPECIFIED, " +
				"VARIABLE_FORMAT_TEXT, VARIABLE_FORMAT_NUMBER, VARIABLE_FORMAT_FILE, " +
				"VARIABLE_FORMAT_DROPDOWN, VARIABLE_FORMAT_BOOLEAN, VARIABLE_FORMAT_PASSWORD",
		},
		"label": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Label",
		},
		"max_length": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Max length of the variable",
		},
		"name": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Name of Variable Group Variable",
		},
		"option": {
			Type:        schema.TypeList,
			Optional:    true,
			Elem:        variableOptionValResourceSchema,
			Description: "Variable Options",
		},
		"required": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Flag indicates if this is a required variable",
		},
		"value": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Value of variable group variable.",
		},
	},
}

var variableGroupConditionResourceSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Name of variable group condition",
		},
		"operator": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Condition Operator",
		},
		"value": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Value",
		},
	},
}

var customConfigVariableGroupResourceSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"condition": {
			Type:        schema.TypeList,
			Optional:    true,
			Elem:        variableGroupConditionResourceSchema,
			MaxItems:    1,
			Description: "Condition",
		},
		"name": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Name of Variable group",
		},
		"required": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Indicates if this is a required condition",
		},
		"variable": {
			Type:        schema.TypeList,
			Optional:    true,
			Elem:        variableGroupVariableResourceSchema,
			Description: "List of variables",
		},
	},
}

var customConfigResourceSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"add": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Add Custom Config",
		},
		"allow_storage_resize": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Allow storage resize",
		},
		"field_delimiter": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Field delimiter",
		},
		"name": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "User Defined name of Custom Config",
		},
		"override": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Override",
		},
		"template": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Template",
		},
		"variable_group": {
			Type:        schema.TypeList,
			Optional:    true,
			Elem:        customConfigVariableGroupResourceSchema,
			Description: "Variable group",
		},
	},
}
