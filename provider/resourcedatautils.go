// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// This file has Utilities to convert Schema Resource Data Attributes to
//  appropriate values ( string, string Ptr, Bool, BoolPtr, map[string]string etc)
package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"reflect"
)

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

func ptrValStr(ptr interface{}) string {
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

// rdEntryByKey
//  Params:
//    rd - Resource Data. Must be *schema.ResourceData OR map[string]interface{}
//    key - Key of Item
//  Returns: (present, value)
//    present - bool - true if item is present. False otherwise
//    value - value of item if present, nil otherwise
func rdEntryByKey(rd interface{}, key string) (bool, interface{}) {
	d, ok := rd.(*schema.ResourceData)
	if ok {
		val, ok := d.GetOk(key)
		if ok {
			// Key present in Resource Data
			return true, val
		}
		return false, nil
	}
	// rd must be a map[string]interface{}
	dmap, ok := rd.(map[string]interface{})
	if !ok {
		// This is a bug - not a user mistake in .tf file. So panic
		err := fmt.Errorf("Unexpected rd type (%+v). rd is expected to be of type "+
			"*schema.ResourceData or map[string]interface{}", rd)
		panic(err)
	}
	if val, ok := dmap[key]; ok {
		return true, val
	}
	return false, nil
}

// rdEntryInt32
//  Returns int32 value for the Specified key from ResourceData. If key doesn't exist,
//      returns 0.
// Params:
//  rd can be *schema.ResourceData OR map[string]interface{}. Any other
//      type will cause a Panic
func rdEntryInt32(rd interface{}, key string) int32 {
	ok, val := rdEntryByKey(rd, key)
	if ok {
		return int32(val.(int))
	}
	return 0
}

// rdEntryInt32PtrOrNil
//  Returns int32Ptr value for the Specified key from ResourceData. If key doesn't exist,
//      returns nil.
// Params:
//  rd can be *schema.ResourceData OR map[string]interface{}. Any other
//      type will cause a Panic
func rdEntryInt32PtrOrNil(rd interface{}, key string) *int32 {
	if ok, val := rdEntryByKey(rd, key); ok {
		intVal := int32(val.(int))
		return &intVal
	}
	return nil
}

// rdEntryInt64
//  Returns int64 value for the Specified key from ResourceData. If key doesn't exist,
//      returns 0.
// Params:
//  rd can be *schema.ResourceData OR map[string]interface{}. Any other
//      type will cause a Panic
func rdEntryInt64(rd interface{}, key string) int64 {
	if ok, val := rdEntryByKey(rd, key); ok {
		return int64(val.(int))
	}
	return 0
}

// rdEntryInt64PtrOrNil
//  Returns *int64 value for the Specified key from ResourceData. If key doesn't exist,
//      returns ni  l.
// Params:
//  rd can be *schema.ResourceData OR map[string]interface{}. Any other
//      type will cause a Panic
func rdEntryInt64PtrOrNil(rd interface{}, key string) *int64 {
	if ok, val := rdEntryByKey(rd, key); ok {
		intVal := int64(val.(int))
		return &intVal
	}
	return nil
}

// rdEntryUint64PtrOrNil
//  Returns Uint64Ptr for the Specified key from ResourceData
// Params:
//  rd can be *schema.ResourceData OR map[string]interface{}. Any other
//      type will cause a Panic
func rdEntryUint64PtrOrNil(rd interface{}, key string) *uint64 {
	ok, val := rdEntryByKey(rd, key)
	if !ok {
		return nil
	}
	uint64Val := uint64(val.(int))
	return &uint64Val
}

// rdEntryStr
//  Returns string value for the Specified key from ResourceData. If key doesn't exist,
//      returns "".
// Params:
//  rd can be *schema.ResourceData OR map[string]interface{}. Any other
//      type will cause a Panic
func rdEntryStr(rd interface{}, key string) string {
	ok, val := rdEntryByKey(rd, key)
	if ok {
		return val.(string)
	}
	return ""
}

// rdEntryStrPtrOrNil
//  Returns *string for the Specified key from ResourceData. If key doesn't exist,
//      returns nil.
// Params:
//  rd can be *schema.ResourceData OR map[string]interface{}. Any other
//      type will cause a Panic
func rdEntryStrPtrOrNil(rd interface{}, key string) *string {
	ok, val := rdEntryByKey(rd, key)
	if ok {
		val := val.(string)
		return &val
	}
	return nil
}

// rdEntryStrList
//  Converts the specified entry of type []interface{} from ResourceData
//      into []string
// Params:
//  rd can be *schema.ResourceData OR map[string]interface{}. Any other
//      type will cause a Panic
func rdEntryStrList(rd interface{}, key string) []string {
	strList := make([]string, 0)
	ok, val := rdEntryByKey(rd, key)
	if !ok {
		return strList
	}
	for _, val := range val.([]interface{}) {
		strList = append(strList, val.(string))
	}
	return strList
}

// rdEntryStrMap
//  Converts the specified entry of type map[string]interface{} from ResourceData
//      into map[string]string
// Params:
//  rd can be *schema.ResourceData OR map[string]interface{}. Any other
//      type will cause a Panic
func rdEntryStrMap(rd interface{}, key string) map[string]string {
	strMap := make(map[string]string)
	ok, val := rdEntryByKey(rd, key)
	if !ok {
		return strMap
	}
	for key, val := range val.(map[string]interface{}) {
		strMap[key] = val.(string)
	}
	return strMap
}

// rdEntryBool
//  Returns bool for the Specified key from ResourceData
// Params:
//  rd can be *schema.ResourceData OR map[string]interface{}. Any other
//      type will cause a Panic
func rdEntryBool(rd interface{}, key string) bool {
	ok, val := rdEntryByKey(rd, key)
	if ok {
		return val.(bool)
	}
	return false
}

// rdEntryBoolPtrOrNil
//  Returns BoolPtr for the Specified key from ResourceData
// Params:
//  rd can be *schema.ResourceData OR map[string]interface{}. Any other
//      type will cause a Panic
func rdEntryBoolPtrOrNil(rd interface{}, key string) *bool {
	ok, val := rdEntryByKey(rd, key)
	if !ok {
		return nil
	}
	boolVal := val.(bool)
	return &boolVal
}

// rdEntryList
//  Returns []interface{} for the Specified key from ResourceData
//      If key doesn't exist, returns an empty list of type []interface{}
// Params:
//  rd can be *schema.ResourceData OR map[string]interface{}. Any other
//      type will cause a Panic
func rdEntryList(rd interface{}, key string) []interface{} {
	ok, val := rdEntryByKey(rd, key)
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
	panic("Key %s Value %+v - neither a list not a set")
}

type rdFuncMapToEntry func(d map[string]interface{}) (interface{}, error)

// rdEntryStructPtr
//  This is suposed to be used to convert attributes that translate to a
//  Ptr to some struct. The specified key, if present, is expected to have a value of
//  []interface{} of length 1.
//  Returns (interface{}, err) for the Specified key from ResourceData
//      If key doesn't exist, returns nil
// Params:
//  rd can be *schema.ResourceData OR map[string]interface{}. Any other
//      type will cause a Panic
func rdEntryStructPtr(rd interface{}, key string, entryFn rdFuncMapToEntry) (interface{}, error) {
	ok, val := rdEntryByKey(rd, key)
	if !ok {
		return nil, nil
	}
	entryList := val.([]interface{})
	if len(entryList) == 0 {
		return nil, nil
	}
	return entryFn(entryList[0].(map[string]interface{}))
}
