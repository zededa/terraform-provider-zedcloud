package resources

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
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

// TestProvider_InternalValidate checks the provider schema the way the SDK
// does at plugin start-up: every resource, data source, nested block and
// Importer.
//
// It needs no controller and no credentials, so a malformed schema fails here
// in seconds rather than at `terraform apply` against a live tenant.
func TestProvider_InternalValidate(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("provider schema is invalid: %s", err)
	}
}

// TestProviderResourcesAreImportable pins the NFR-165 outcome: the resources
// that represent a real, readable object all support `terraform import`.
//
// The exclusions are deliberate and documented, not a backlog:
//   - zedcloud_credential has ReadContext: schema.NoopContext because the API
//     never returns credential material. There is nothing to import.
//   - the manifest and status resources are read-only projections whose
//     Create/Update/Delete are all no-ops.
func TestProviderResourcesAreImportable(t *testing.T) {
	notImportable := map[string]string{
		"zedcloud_credential":                "write-only; the API never returns credential material",
		"zedcloud_cluster_group_manifest":    "read-only projection, not a managed object",
		"zedcloud_kubernetes_cluster_status": "read-only projection, not a managed object",

		// NOT permanent. Held back pending NFR-165 §7: OAUTHProfile carries
		// ClientSecret, CryptoKey and EncryptedSecrets, and import would write
		// whatever the API returns into plaintext state. Confirm against a live
		// tenant what GET /v1/authorization/profiles/id/{id} actually returns
		// for those fields, mark them Sensitive if they come back populated,
		// then add the Importer and delete this entry -- this test will tell
		// you to.
		"zedcloud_auth_profile": "pending the NFR-165 §7 secret-exposure audit",
	}

	for name, res := range Provider().ResourcesMap {
		reason, excluded := notImportable[name]

		switch {
		case excluded && res.Importer != nil:
			t.Errorf("%s has an Importer but is listed as not importable (%s); "+
				"remove it from the exclusion list", name, reason)
		case !excluded && res.Importer == nil:
			t.Errorf("%s has no Importer. Either add one, or add it to the "+
				"exclusion list here with the reason why it cannot be imported.", name)
		}
	}
}

// importIDFromAttribute builds an ImportStateIdFunc that feeds the value of a
// state attribute -- a username or a role name -- to `terraform import`,
// exercising the by-name resolution in NFR-165 §4.1.
func importIDFromAttribute(resourceName, attribute string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("%s not found in state", resourceName)
		}
		value := rs.Primary.Attributes[attribute]
		if value == "" {
			return "", fmt.Errorf("%s has no %s in state", resourceName, attribute)
		}
		return value, nil
	}
}
