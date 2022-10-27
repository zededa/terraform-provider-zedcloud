// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// This file has Utilities to convert Schema Resource Data Attributes to
//
//	appropriate values ( string, string Ptr, Bool, BoolPtr, map[string]string etc)
package utils

import (
	"fmt"
	"reflect"
)

func StrPtr(val string) *string {
	return &val
}

func Int64Ptr(val int64) *int64 {
	return &val
}

func Int32Ptr(val int32) *int32 {
	return &val
}

func BoolPtr(val bool) *bool {
	return &val
}

func PtrValStr(ptr interface{}) string {
	value := reflect.ValueOf(ptr)
	if value.Type().Kind() != reflect.Ptr {
		panic(fmt.Errorf("Non-Ptr - Expecting a Ptr. %+v", ptr))
	}
	if value.IsNil() {
		return ""
	}
	return fmt.Sprintf("%v", value.Elem())
}

func PtrValBool(ptr *bool) bool {
	if ptr == nil {
		return false
	}
	return *ptr
}
