package resources

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Some attributes (zedcloud_application.tags, zedcloud_edgenode.pcr_policy)
// can also be set outside Terraform (UI, zcli, API), and the controller clears
// them when a full-object PUT leaves them out. For those attributes the
// provider uses "managed only once configured" semantics:
//
//   - read: the attribute is written to state only if it is already in state,
//     i.e. it came from config (create/update put it there before the read).
//     A value set outside Terraform never enters state, so it causes no diff.
//   - update: if the attribute is in neither the old nor the new state, the
//     provider sends the value the object has on the controller, so it
//     survives the apply. If it was in state and is removed from config, the
//     provider sends nothing, which clears it.
//
// d.Get returns the prior state during read and the planned value during
// create/update, which is what these helpers rely on.

// asMap returns v as a map, or nil.
func asMap(v interface{}) map[string]interface{} {
	m, _ := v.(map[string]interface{})
	return m
}

// asList returns v as a list, or nil.
func asList(v interface{}) []interface{} {
	l, _ := v.([]interface{})
	return l
}

// isManagedMap reports whether the TypeMap attribute key is managed by
// Terraform (non-empty in state/plan).
func isManagedMap(d *schema.ResourceData, key string) bool {
	return len(asMap(d.Get(key))) > 0
}

// isManagedList reports whether the TypeList attribute key is managed by
// Terraform (non-empty in state/plan).
func isManagedList(d *schema.ResourceData, key string) bool {
	return len(asList(d.Get(key))) > 0
}
