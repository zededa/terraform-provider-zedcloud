package resources

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	testhelper "github.com/zededa/terraform-provider-zedcloud/v2/testing"
)

var testAccProviders map[string]*schema.Provider
var testProvider *schema.Provider

func init() {
	testProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"zedcloud": testProvider,
	}
}

// suffixed appends the per-run suffix to an expected object name, matching what
// the fixtures render through testhelper.MustGetTestInput.
//
// Any assertion on a `name`/`title` the fixture tokenised has to go through
// this. Assertions on values the fixtures leave literal -- structural keys, and
// filler titles like "title" -- must NOT: see
// v2/testing/scripts/suffix_fixtures.py for which is which.
func suffixed(name string) string {
	return name + testhelper.Suffix()
}
