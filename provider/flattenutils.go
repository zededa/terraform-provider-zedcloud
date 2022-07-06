// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/zedcloud-api/swagger_models"
)

func flattenStringList(cfgList []string) []interface{} {
	entryList := make([]interface{}, 0)
	for _, v := range cfgList {
		entryList = append(entryList, v)
	}
	return entryList
}

func flattenStringMap(entryMap map[string]string) map[string]interface{} {
	dataMap := make(map[string]interface{})
	for k, v := range entryMap {
		dataMap[k] = v
	}
	return dataMap
}

func flattenObjectRevision(entry interface{}) []interface{} {
	if entry == nil {
		return []interface{}{}
	}
	object := entry.(*swagger_models.ObjectRevision)
	if object == nil {
		return []interface{}{}
	}
	data := map[string]interface{}{
		"created_at": "",
		"created_by": ptrValStr(object.CreatedBy),
		"curr":       ptrValStr(object.Curr),
		"prev":       object.Prev,
		"updated_at": "",
		"updated_by": ptrValStr(object.UpdatedBy),
	}
	if object.CreatedAt != nil {
		data["created_at"] = fmt.Sprintf("%+v", object.CreatedAt)
	}
	if object.UpdatedAt != nil {
		data["updated_at"] = fmt.Sprintf("%+v", object.UpdatedAt)
	}
	return []interface{}{data}
}

func flattenObjectParentDetail(entry interface{}) []interface{} {
	object := entry.(*swagger_models.ObjectParentDetail)
	if object == nil {
		return []interface{}{}
	}
	return []interface{}{map[string]interface{}{
		"id_of_parent":             object.IDOfParentObject,
		"reference_exists":         object.ReferenceExists,
		"update_available":         object.UpdateAvailable,
		"version_of_parent_object": int(object.VersionOfParentObject),
	}}
}

func flattenZedCloudOpsCmd(entry interface{}) []interface{} {
	object := entry.(*swagger_models.ZedCloudOpsCmd)
	if object == nil {
		return []interface{}{}
	}
	return []interface{}{map[string]interface{}{
		"counter":  int(object.Counter),
		"ops_time": fmt.Sprintf("%v", object.OpsTime),
	}}
}

func flattenedDataCheckKeys(schemaMap map[string]*schema.Schema,
	data map[string]interface{}) {
	// Verify all keys published in data re computed fields.
	for k, _ := range data {
		_, ok := schemaMap[k]
		if !ok {
			panic(fmt.Errorf("Flattened Key %s doesn't exist in Schema", k))
		}
	}
	// Verify all fields in the schema are published
	for k, v := range schemaMap {
		_, ok := data[k]
		if !ok {
			if v.Computed {
				// Computed fields must always be piblished.
				panic(fmt.Errorf("Computed Key %s not flattened", k))
			}
			// Sensitive keys are not published for data sources.
			// So can't really assert for missing keys
		}
	}
}
