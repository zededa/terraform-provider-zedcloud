// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package state

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/internal/resourcedata"
)

// SetLocal is used to ensure the data is put into a format Terraform can output
func SetLocal(d *schema.ResourceData, vals map[string]interface{}) {
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

func CheckInvalidAttrForUpdate(d *schema.ResourceData, name, id string) error {
	newName := resourcedata.GetStr(d, "name")
	if newName != name {
		return fmt.Errorf("Name of an object cannot be changed. current: %s, new: %s", name, newName)
	}
	idInState := resourcedata.GetStr(d, "id")
	if idInState != id {
		return fmt.Errorf(
			"ID of the object has Changed. Cannot update the object. ID in state: %s, ID of Object in Cloud: %s",
			idInState,
			id,
		)
	}
	return nil
}

func RemoveLocal(d *schema.ResourceData, resourceType, id, name string) {
	log.Printf("remove %s %s from state (id=%s)", resourceType, name, id)
	schema.RemoveFromState(d, nil)
}
