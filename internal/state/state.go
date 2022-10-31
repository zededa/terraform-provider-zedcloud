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

func RemoveLocal(d *schema.ResourceData, resourceType, id, name string) {
	log.Printf("remove %s %s from state (id=%s)", resourceType, name, id)
	schema.RemoveFromState(d, nil)
}

func CheckInvalidAttrForUpdate(d *schema.ResourceData, remoteName, remoteID string) error {
	localName := resourcedata.GetStr(d, "name")
	if localName != remoteName {
		return fmt.Errorf(
			"cannot change name of object after creation, remote: %s, local: %s",
			remoteName,
			localName,
		)
	}
	localID := resourcedata.GetStr(d, "id")
	if localID != remoteID {
		return fmt.Errorf(
			"cannot update the object, ID of the object has changed. remote: %s, local: %s",
			remoteID,
			localID,
		)
	}
	return nil
}
