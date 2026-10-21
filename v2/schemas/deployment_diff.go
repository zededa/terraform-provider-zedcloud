package schemas

import (
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// zedcloud_deployment used to have perpetual diffs that did not matter while
// the resource could not be updated (UpdateDeployment was a no-op). Now that a
// change creates a new deployment version on the controller, every such diff
// would create a new version on every apply, so they are suppressed here.

// suppressEmptyNestedMapDiffs walks a schema tree and suppresses diffs between
// an absent and an empty map on every nested TypeMap that has no suppress
// function of its own. The controller returns empty maps (tags, port_tags,
// license_list, ...) that config leaves out, and the SDK plans "+ x = {}" for
// them. Top-level maps (nested=false) are left alone.
func suppressEmptyNestedMapDiffs(s map[string]*schema.Schema, nested bool) {
	for _, sch := range s {
		if sch == nil {
			continue
		}
		if sch.Type == schema.TypeMap && nested && sch.DiffSuppressFunc == nil && !sch.Computed {
			sch.DiffSuppressFunc = suppressEmptyMapDiff
		}
		if r, ok := sch.Elem.(*schema.Resource); ok {
			suppressEmptyNestedMapDiffs(r.Schema, true)
		}
	}
}

// suppressEmptyMapDiff suppresses the diff of a map whose old and new values
// are both empty. The SDK calls it for the count key ("<map>.%") and for each
// element key; only the count can be absent-vs-zero.
func suppressEmptyMapDiff(k, old, new string, d *schema.ResourceData) bool {
	if !strings.HasSuffix(k, ".%") {
		return false
	}
	return (old == "" || old == "0") && (new == "" || new == "0")
}

// The controller stores device policies as "<name>-<kind>-v<N>" (kind is
// attestation, pcr or pnac) and returns "<name>-<kind>" on read. Updates that
// send the returned name back grow it ("<name>-pcr-pcr"), so the provider
// strips the kind suffixes before sending (StripDevicePolicyNameDecoration)
// and ignores them when comparing.
var devicePolicyNameDecoration = regexp.MustCompile(`(-(attestation|pcr|pnac))+$`)

// StripDevicePolicyNameDecoration returns the user-chosen device policy name.
func StripDevicePolicyNameDecoration(name string) string {
	return devicePolicyNameDecoration.ReplaceAllString(name, "")
}

func suppressDecoratedDevicePolicyName(k, old, new string, d *schema.ResourceData) bool {
	if old == new {
		return true
	}
	return new != "" && StripDevicePolicyNameDecoration(old) == StripDevicePolicyNameDecoration(new)
}
