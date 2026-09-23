// Copyright (c) Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package resources

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
	zschema "github.com/zededa/terraform-provider-zedcloud/v2/schemas"
)

// ===========================================================================
// CI-752, second defect.
//
// The ticket's plan output carries a diff nobody has yet explained, on a
// resource that has nothing to do with clusters:
//
//	~ resource "zedcloud_datastore" "demo_az_ds" {
//	    + api_key = (sensitive value)
//	    + secret {
//	        + api_passwd = (sensitive value)
//	      }
//
// `+` on a value the configuration has always supplied means state does not
// have it. `api_key` and `secret` are write-only on the API: gilas nils the
// secrets on every read (srvs/gilas/datastore_handlers.go:345 and :411) and
// seine never populates the legacy top-level `apiKey` on a response
// (srvs/seine/datastoreproc.go:127-216). SetDatastoreResourceData set both
// unconditionally, so refresh wrote the redacted emptiness into state and the
// configuration re-added them on the next plan. Forever: the value is never
// returned, so the diff cannot converge.
//
// Unlike the device_id half of CI-752 this one is not merely noisy. The
// controller's update path writes the whole row with no merge
// (srvs/seine/datastoreproc.go:86-87, and neither DsSecrets nor CryptoKeyCtx
// carries `nullzero`/`update_invalid` in srvs/seine/datastoredata.go:385-386),
// so a PUT built from blanked state overwrites the stored credentials with
// empty.
// ===========================================================================

const (
	testDSKey    = "AKIAIOSFODNN7EXAMPLE"
	testDSPasswd = "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"
)

// datastoreStateWithCredentials is a state that has been through a successful
// create: the credentials the configuration supplied are in it.
func datastoreStateWithCredentials(t *testing.T) *schema.ResourceData {
	t.Helper()

	d := schema.TestResourceDataRaw(t, zschema.Datastore(), map[string]interface{}{
		"name":    "TF-KUBCON-AZURE-DS",
		"title":   "TF-KUBCON-AZURE-DS",
		"ds_type": "DATASTORE_TYPE_AZUREBLOB",
		"ds_fqdn": "https://example.blob.core.windows.net",
		"api_key": testDSKey,
		"secret": []interface{}{
			map[string]interface{}{
				"api_key":    testDSKey,
				"api_passwd": testDSPasswd,
			},
		},
	})
	d.SetId("060d1566-6f51-43ed-8b6e-86ac2823c338")

	// Guard the fixture itself: if TestResourceDataRaw ever stops populating
	// these, the assertions below would pass for the wrong reason.
	if got := d.Get("api_key").(string); got != testDSKey {
		t.Fatalf("fixture: api_key is %q before the refresh, want %q", got, testDSKey)
	}
	if n := len(d.Get("secret").([]interface{})); n != 1 {
		t.Fatalf("fixture: secret has %d blocks before the refresh, want 1", n)
	}
	return d
}

// redactedDatastore is what a GET actually returns: everything but the
// credentials, which the controller has stripped.
func redactedDatastore() *models.Datastore {
	name := "TF-KUBCON-AZURE-DS"
	title := "TF-KUBCON-AZURE-DS"
	return &models.Datastore{
		ID:     "060d1566-6f51-43ed-8b6e-86ac2823c338",
		Name:   &name,
		Title:  &title,
		DsFQDN: &[]string{"https://example.blob.core.windows.net"}[0],
		// APIKey: "" and Secret: nil -- redacted, NOT removed.
		EncryptedSecrets: map[string]string{"cipherData": "<vault ciphertext>"},
		CryptoKey:        "<crypto key context>",
	}
}

func TestDatastore_RefreshKeepsWriteOnlyCredentials(t *testing.T) {
	d := datastoreStateWithCredentials(t)

	zschema.SetDatastoreResourceData(d, redactedDatastore())

	if got := d.Get("api_key").(string); got != testDSKey {
		t.Errorf("CI-752 regression: refresh erased api_key from state (got %q, want %q).\n"+
			"The API never returns it, so the configuration re-adds it on every plan "+
			"and the diff never converges.", got, testDSKey)
	}

	secret, _ := d.Get("secret").([]interface{})
	if len(secret) != 1 {
		t.Fatalf("CI-752 regression: refresh erased the secret block from state "+
			"(%d blocks, want 1)", len(secret))
	}
	if got := secret[0].(map[string]interface{})["api_passwd"].(string); got != testDSPasswd {
		t.Errorf("CI-752 regression: refresh erased secret.api_passwd (got %q)", got)
	}

	// The fields the controller DOES return must still be written through --
	// the fix must not turn into "never update anything sensitive".
	if got := d.Get("crypto_key").(string); got != "<crypto key context>" {
		t.Errorf("crypto_key was not refreshed: got %q. It is returned by the API "+
			"(srvs/seine/datastoreproc.go:177-180) and must keep tracking the server.", got)
	}
}

// NEGATIVE CONTROL: the unconditional Set this replaced. Without it the test
// above could pass simply because nothing ever writes to those keys.
func TestDatastore_UnconditionalSetErasesCredentials(t *testing.T) {
	d := datastoreStateWithCredentials(t)
	m := redactedDatastore()

	// Exactly what SetDatastoreResourceData used to do for these two fields.
	d.Set("api_key", m.APIKey)
	d.Set("secret", zschema.SetDatastoreSecretsSubResourceData(
		[]*models.DatastoreInfoSecrets{m.Secret}))

	if d.Get("api_key").(string) != "" || len(d.Get("secret").([]interface{})) != 0 {
		t.Fatalf("negative control did not reproduce CI-752: the unconditional Set left "+
			"api_key=%q and %d secret blocks, so the assertion in the test above proves "+
			"nothing. Fix this test before trusting it.",
			d.Get("api_key").(string), len(d.Get("secret").([]interface{})))
	}
	t.Log("negative control reproduced the diff: api_key and secret both cleared by refresh")
}
