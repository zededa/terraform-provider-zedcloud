// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func getErrMsgPrefix(name, id, objectType, action string) string {
	return fmt.Sprintf("[ERROR] %s %s ( id: %s) %s Failed.",
		objectType, name, id, action)
}

// marshalData is used to ensure the data is put into a format Terraform can output
func marshalData(d *schema.ResourceData, vals map[string]interface{}) {
	for k, v := range vals {
		if k == "id" {
			d.SetId(v.(string))
		} else {
			str, ok := v.(string)
			if ok {
				d.Set(k, str)
			} else {
				d.Set(k, v)
			}
		}
	}
}

func checkInvalidAttrForUpdate(d *schema.ResourceData, name, id string) error {
	newName := rdEntryStr(d, "name")
	if newName != name {
		return fmt.Errorf("Name of an object cannot be changed. current: %s, new: %s",
			name, newName)
	}
	idInState := rdEntryStr(d, "id")
	if idInState != id {
		return fmt.Errorf("ID of the object has Changed. Cannot update the object. "+
			"ID in state: %s, ID of Object in Cloud: %s", idInState, id)
	}
	return nil
}
