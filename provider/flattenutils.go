// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"fmt"
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
	object := entry.(*swagger_models.ObjectRevision)
	if object == nil {
		return []interface{}{}
	}
	return []interface{}{map[string]interface{}{
		"created_at": fmt.Sprintf("%+v", object.CreatedAt),
		"created_by": ptrValStr(object.CreatedBy),
		"curr":       ptrValStr(object.Curr),
		"prev":       object.Prev,
		"updated_at": fmt.Sprintf("%s", object.UpdatedAt),
		"updated_by": ptrValStr(object.UpdatedBy),
	}}
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
