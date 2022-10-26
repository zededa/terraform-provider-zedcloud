// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// This file has Utilities to convert Schema Resource Data Attributes to
//
//	appropriate values ( string, string Ptr, Bool, BoolPtr, map[string]string etc)
package provider

import (
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// FIXME(fabian): remove all panics!

func strPtr(val string) *string {
	return &val
}

func int64Ptr(val int64) *int64 {
	return &val
}

func int32Ptr(val int32) *int32 {
	return &val
}

func boolPtr(val bool) *bool {
	return &val
}

func strPtrVal(ptr interface{}) string {
	value := reflect.ValueOf(ptr)
	if value.Type().Kind() != reflect.Ptr {
		panic(fmt.Errorf("Non-Ptr - Expecting a Ptr. %+v", ptr))
	}
	if value.IsNil() {
		return ""
	}
	return fmt.Sprintf("%v", value.Elem())
}

func ptrValBool(ptr *bool) bool {
	if ptr == nil {
		return false
	}
	return *ptr
}

// getByKey
//
//	Params:
//	  resourceData - Must be *schema.ResourceData OR map[string]interface{}
//	  key - Key of Item
//	Returns: (present, value)
//	  present - bool - true if item is present. False otherwise
//	  value - value of item if present, nil otherwise
func getByKey(resourceData interface{}, key string) (bool, interface{}) {
	d, ok := resourceData.(*schema.ResourceData)
	if ok {
		val, ok := d.GetOk(key)
		if ok {
			// Key present in Resource Data
			return true, val
		}
		return false, nil
	}
	// resourceData must be a map[string]interface{}
	dmap, ok := resourceData.(map[string]interface{})
	if !ok {
		// This is a bug - not a user mistake in .tf file. So panic
		err := fmt.Errorf(
			"expect resourceData of type %T or map[string]interface{} but got %T",
			&schema.ResourceData{},
			resourceData,
		)
		panic(err)
	}
	if val, ok := dmap[key]; ok {
		return true, val
	}
	return false, nil
}

// getInt32
//
//	Returns int32 value for the Specified key from ResourceData. If key doesn't exist,
//	    returns 0.
//
// Params:
//
//	resourceData can be *schema.ResourceData OR map[string]interface{}. Any other
//	    type will cause a Panic
func getInt32(resourceData interface{}, key string) int32 {
	ok, val := getByKey(resourceData, key)
	if ok {
		return int32(val.(int))
	}
	return 0
}

// getInt32PtrOrNil
//
//	Returns int32Ptr value for the Specified key from ResourceData. If key doesn't exist,
//	    returns nil.
//
// Params:
//
//	resourceData can be *schema.ResourceData OR map[string]interface{}. Any other
//	    type will cause a Panic
func getInt32PtrOrNil(resourceData interface{}, key string) *int32 {
	if ok, val := getByKey(resourceData, key); ok {
		intVal := int32(val.(int))
		return &intVal
	}
	return nil
}

// getInt64
//
//	Returns int64 value for the Specified key from ResourceData. If key doesn't exist,
//	    returns 0.
//
// Params:
//
//	resourceData can be *schema.ResourceData OR map[string]interface{}. Any other
//	    type will cause a Panic
func getInt64(resourceData interface{}, key string) int64 {
	if ok, val := getByKey(resourceData, key); ok {
		return int64(val.(int))
	}
	return 0
}

// getInt64PtrOrNil
//
//	Returns *int64 value for the Specified key from ResourceData. If key doesn't exist,
//	    returns ni  l.
//
// Params:
//
//	resourceData can be *schema.ResourceData OR map[string]interface{}. Any other
//	    type will cause a Panic
func getInt64PtrOrNil(resourceData interface{}, key string) *int64 {
	if ok, val := getByKey(resourceData, key); ok {
		intVal := int64(val.(int))
		return &intVal
	}
	return nil
}

// getUint64PtrOrNil
//
//	Returns Uint64Ptr for the Specified key from ResourceData
//
// Params:
//
//	resourceData can be *schema.ResourceData OR map[string]interface{}. Any other
//	    type will cause a Panic
func getUint64PtrOrNil(resourceData interface{}, key string) *uint64 {
	ok, val := getByKey(resourceData, key)
	if !ok {
		return nil
	}
	uint64Val := uint64(val.(int))
	return &uint64Val
}

// getStr
//
//	Returns string value for the Specified key from ResourceData. If key doesn't exist,
//	    returns "".
//
// Params:
//
//	resourceData can be *schema.ResourceData OR map[string]interface{}. Any other
//	    type will cause a Panic
func getStr(resourceData interface{}, key string) string {
	ok, val := getByKey(resourceData, key)
	if ok {
		return val.(string)
	}
	return ""
}

// getStrPtrOrNil
//
//	Returns *string for the Specified key from ResourceData. If key doesn't exist,
//	    returns nil.
//
// Params:
//
//	resourceData can be *schema.ResourceData OR map[string]interface{}. Any other
//	    type will cause a Panic
func getStrPtrOrNil(resourceData interface{}, key string) *string {
	ok, val := getByKey(resourceData, key)
	if ok {
		val := val.(string)
		return &val
	}
	return nil
}

// getStrList
//
//	Converts the specified entry of type []interface{} from ResourceData
//	    into []string
//
// Params:
//
//	resourceData can be *schema.ResourceData OR map[string]interface{}. Any other
//	    type will cause a Panic
func getStrList(resourceData interface{}, key string) []string {
	strList := make([]string, 0)
	ok, val := getByKey(resourceData, key)
	if !ok {
		return strList
	}
	for _, val := range val.([]interface{}) {
		strList = append(strList, val.(string))
	}
	return strList
}

// getStrSet
//
//	Converts the specified entry of type []interface{} from ResourceData
//	    into []string
//
// Params:
//
//	resourceData can be *schema.ResourceData OR map[string]interface{}. Any other
//	    type will cause a Panic
func getStrSet(resourceData interface{}, key string) []string {
	strList := make([]string, 0)
	ok, val := getByKey(resourceData, key)
	if !ok {
		return strList
	}
	rdList := (val.(*schema.Set)).List()
	for _, val := range rdList {
		strList = append(strList, val.(string))
	}
	return strList
}

// getStrMap
//
//	Converts the specified entry of type map[string]interface{} from ResourceData
//	    into map[string]string
//
// Params:
//
//	resourceData can be *schema.ResourceData OR map[string]interface{}. Any other
//	    type will cause a Panic
func getStrMap(resourceData interface{}, key string) map[string]string {
	strMap := make(map[string]string)
	ok, val := getByKey(resourceData, key)
	if !ok {
		return strMap
	}
	for key, val := range val.(map[string]interface{}) {
		strMap[key] = val.(string)
	}
	return strMap
}

// getBool
//
//	Returns bool for the Specified key from ResourceData
//
// Params:
//
//	resourceData can be *schema.ResourceData OR map[string]interface{}. Any other
//	    type will cause a Panic
func getBool(resourceData interface{}, key string) bool {
	ok, val := getByKey(resourceData, key)
	if ok {
		return val.(bool)
	}
	return false
}

// getBoolPtrOrNil
//
//	Returns BoolPtr for the Specified key from ResourceData
//
// Params:
//
//	resourceData can be *schema.ResourceData OR map[string]interface{}. Any other
//	    type will cause a Panic
func getBoolPtrOrNil(resourceData interface{}, key string) *bool {
	ok, val := getByKey(resourceData, key)
	if !ok {
		return nil
	}
	boolVal := val.(bool)
	return &boolVal
}

// getList
//
//	Returns []interface{} for the Specified key from ResourceData
//	    If key doesn't exist, returns an empty list of type []interface{}
//
// Params:
//
//	resourceData can be *schema.ResourceData OR map[string]interface{}. Any other
//	    type will cause a Panic
func getList(resourceData interface{}, key string) []interface{} {
	ok, val := getByKey(resourceData, key)
	if !ok {
		return []interface{}{}
	}
	listVal, ok := val.([]interface{})
	if ok {
		return listVal
	}
	setVal, ok := val.(*schema.Set)
	if ok {
		return setVal.List()
	}
	panic("Key %s Value %+v - neither a list nor a set")
}

type rdFuncMapToEntry func(d map[string]interface{}) (interface{}, error)

// getStructPtr
//
//	This is suposed to be used to convert attributes that translate to a
//	Ptr to some struct. The specified key, if present, is expected to have a value of
//	[]interface{} of length 1.
//	Returns (interface{}, err) for the Specified key from ResourceData
//	    If key doesn't exist, returns nil
//
// Params:
//
//	resourceData can be *schema.ResourceData OR map[string]interface{}. Any other
//	    type will cause a Panic
func getStructPtr(resourceData interface{}, key string, entryFn rdFuncMapToEntry) (interface{}, error) {
	ok, val := getByKey(resourceData, key)
	if !ok {
		return nil, nil
	}
	entryList := val.([]interface{})
	if len(entryList) == 0 {
		return nil, nil
	}
	return entryFn(entryList[0].(map[string]interface{}))
}
