// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func pkey(parent, key string) string {
	if parent != "" {
		parent += "."
	}
	return parent + key
}

func doVerifyFlattenOutput(objSchema map[string]*schema.Schema,
	output map[string]interface{},
	checkAllKeysPresent bool,
	parentKeys string) string {
	errStr := ""
	if checkAllKeysPresent {
		if len(output) != len(objSchema) {
			for k := range objSchema {
				if _, ok := output[k]; !ok {
					errStr += fmt.Sprintf("Key (%s) not present in Key Chain: %s\n",
						k, parentKeys)
				}
			}
		}
	}
	for k, v := range output {
		schemaEntry, ok := objSchema[k]
		if !ok {
			errStr += fmt.Sprintf("Invalid Key ( %s ) in output. parent: %s\n",
				k, parentKeys)
			continue
		}
		if v == nil {
			continue
		}
		switch schemaEntry.Type {
		case schema.TypeList:
			listVal, ok := v.([]interface{})
			if !ok {
				errStr += fmt.Sprintf("Key: %s - Value not List. Value: %#v, type(v): %T",
					pkey(parentKeys, k), v, v)
				continue
			}
			parentKeyVal := pkey(parentKeys, k)
			for indx, entry := range listVal {
				entryParentKey := fmt.Sprintf("%s[%d]", parentKeyVal, indx)
				listEntryVal, ok := entry.(map[string]interface{})
				if ok {
					elem, ok := schemaEntry.Elem.(*schema.Resource)
					if ok {
						errStr += doVerifyFlattenOutput(elem.Schema, listEntryVal,
							checkAllKeysPresent, entryParentKey)
					}
				}
			}
		case schema.TypeMap:
			_, ok := schemaEntry.Elem.(*schema.Resource)
			if ok {
				_, ok := v.(map[string]interface{})
				if !ok {
					errStr += fmt.Sprintf("Key: %s - Value not map[string]interface{} "+
						"as expected. Value: %#v, Type: %T",
						pkey(parentKeys, k), v, v)
					continue
				}
			} else {
				elem := schemaEntry.Elem.(*schema.Schema)
				if elem.Type == schema.TypeString {
					_, ok := v.(map[string]interface{})
					if !ok {
						errStr += fmt.Sprintf("Key: %s - Value not map[string]string "+
							"as expected. Value: %#v, Type: %T",
							pkey(parentKeys, k), v, v)
						continue
					}
				}

			}
		}
	}
	return errStr
}

func verifyFlattenOutput(
	objSchema map[string]*schema.Schema,
	output map[string]interface{},
	checkAllKeysPresent bool) error {
	errStr := doVerifyFlattenOutput(objSchema, output, checkAllKeysPresent, "")
	if errStr == "" {
		return nil
	}
	return fmt.Errorf("%s", errStr)
}
