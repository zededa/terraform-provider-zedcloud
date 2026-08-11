package schemas

import (
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func whiteLabelingBlock(primary, secondary, logo, productName string) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"primary_color":   primary,
			"secondary_color": secondary,
			"logo_url":        logo,
			"product_name":    productName,
		},
	}
}

// TestEnterpriseWhiteLabelingAttributes_ExactKeys pins the attribute keys the block
// writes. The enterprise API validates every key against an allowlist and rejects
// the entire create/update request on an unknown one, so a typo here does not
// degrade gracefully - it breaks every apply that sets white-labeling.
func TestEnterpriseWhiteLabelingAttributes_ExactKeys(t *testing.T) {
	got := EnterpriseWhiteLabelingAttributes(
		whiteLabelingBlock("#0A2540", "#00B3A4", "https://acme.example.com/logo.svg", "Acme Edge"),
	)

	want := map[string]string{
		"$ztag.entp.zui.ux.color.primary":   "#0A2540",
		"$ztag.entp.zui.ux.color.secondary": "#00B3A4",
		"$ztag.entp.zui.ux.logo":            "https://acme.example.com/logo.svg",
		"$ztag.entp.zui.ux.product.name":    "Acme Edge",
	}

	if len(got) != len(want) {
		t.Fatalf("expected %d attributes, got %d: %v", len(want), len(got), got)
	}
	for key, value := range want {
		if got[key] != value {
			t.Errorf("attribute %q: expected %q, got %q", key, value, got[key])
		}
	}
}

// TestEnterpriseWhiteLabelingAttributes_OmitsEmpty covers how a single white-label
// value gets cleared. The API replaces the whole attribute map on every update, so
// an unset field has to be absent from the map - sending it as an empty string would
// both fail the API's 3-character minimum and leave the value set.
func TestEnterpriseWhiteLabelingAttributes_OmitsEmpty(t *testing.T) {
	got := EnterpriseWhiteLabelingAttributes(whiteLabelingBlock("#0A2540", "", "", ""))

	if len(got) != 1 {
		t.Fatalf("expected only the primary color, got %v", got)
	}
	if got[whiteLabelPrimaryColorKey] != "#0A2540" {
		t.Errorf("expected primary color to be kept, got %q", got[whiteLabelPrimaryColorKey])
	}
}

// TestEnterpriseWhiteLabelingAttributes_NoBlock guards the write path against an
// absent or empty block. Removing the block must produce no attributes at all,
// which is what clears white-labeling.
func TestEnterpriseWhiteLabelingAttributes_NoBlock(t *testing.T) {
	for name, input := range map[string]interface{}{
		"nil":           nil,
		"empty list":    []interface{}{},
		"nil block":     []interface{}{nil},
		"not a list":    "unexpected",
		"empty block":   whiteLabelingBlock("", "", "", ""),
		"wrong element": []interface{}{"unexpected"},
	} {
		t.Run(name, func(t *testing.T) {
			if got := EnterpriseWhiteLabelingAttributes(input); len(got) != 0 {
				t.Errorf("expected no attributes, got %v", got)
			}
		})
	}
}

// TestEnterpriseWhiteLabeling_RoundTrip is the no-perpetual-diff guarantee: what we
// read back from the API and project into state must convert back to exactly the
// attributes we sent, otherwise every plan shows a change.
func TestEnterpriseWhiteLabeling_RoundTrip(t *testing.T) {
	sent := EnterpriseWhiteLabelingAttributes(
		whiteLabelingBlock("#0A2540", "#00B3A4", "https://acme.example.com/logo.svg", "Acme Edge"),
	)

	block := SetEnterpriseWhiteLabelingSubResourceData(sent)
	if len(block) != 1 {
		t.Fatalf("expected one white_labeling block, got %d", len(block))
	}

	roundTripped := EnterpriseWhiteLabelingAttributes([]interface{}{*block[0]})
	if len(roundTripped) != len(sent) {
		t.Fatalf("expected %d attributes after round trip, got %v", len(sent), roundTripped)
	}
	for key, value := range sent {
		if roundTripped[key] != value {
			t.Errorf("attribute %q: expected %q after round trip, got %q", key, value, roundTripped[key])
		}
	}
}

// TestSetEnterpriseWhiteLabelingSubResourceData_NotWhiteLabeled keeps state clean for
// the overwhelmingly common case of an enterprise with no white-labeling. Returning a
// block of empty strings instead of nil would show up as a spurious block in state.
func TestSetEnterpriseWhiteLabelingSubResourceData_NotWhiteLabeled(t *testing.T) {
	for name, attributes := range map[string]map[string]string{
		"nil":               nil,
		"empty":             {},
		"unrelated key":     {"$ztag.entp.something.else": "value"},
		"empty white label": {whiteLabelLogoKey: ""},
	} {
		t.Run(name, func(t *testing.T) {
			if got := SetEnterpriseWhiteLabelingSubResourceData(attributes); got != nil {
				t.Errorf("expected no block, got %v", *got[0])
			}
		})
	}
}

// TestNonWhiteLabelAttributes checks the split between the block and the deprecated
// raw attributes map. A key owned by the block must not also surface in attributes,
// or both fields would claim the same value and fight over it on every plan.
func TestNonWhiteLabelAttributes(t *testing.T) {
	got := nonWhiteLabelAttributes(map[string]string{
		whiteLabelPrimaryColorKey:   "#0A2540",
		whiteLabelSecondaryColorKey: "#00B3A4",
		whiteLabelLogoKey:           "https://acme.example.com/logo.svg",
		whiteLabelProductNameKey:    "Acme Edge",
		"$ztag.entp.something.else": "value",
	})

	if len(got) != 1 || got["$ztag.entp.something.else"] != "value" {
		t.Errorf("expected only the non-white-label key to remain, got %v", got)
	}
}

// TestSetEnterpriseAttributes_ProjectsIntoBlock is the default read path: an
// enterprise whose white-labeling Terraform does not yet track in state - including a
// fresh import - lands in the typed block, not the deprecated map.
func TestSetEnterpriseAttributes_ProjectsIntoBlock(t *testing.T) {
	d := schema.TestResourceDataRaw(t, EnterpriseSchema(), map[string]interface{}{})

	setEnterpriseAttributes(d, map[string]string{
		whiteLabelPrimaryColorKey: "#0A2540",
		whiteLabelProductNameKey:  "Acme Edge",
	})

	block := d.Get("white_labeling").([]interface{})
	if len(block) != 1 {
		t.Fatalf("expected white-labeling to be projected into the block, got %v", block)
	}
	values := block[0].(map[string]interface{})
	if values["primary_color"] != "#0A2540" {
		t.Errorf("expected primary_color to be set, got %q", values["primary_color"])
	}
	if values["product_name"] != "Acme Edge" {
		t.Errorf("expected product_name to be set, got %q", values["product_name"])
	}
	if attributes := d.Get("attributes").(map[string]interface{}); len(attributes) != 0 {
		t.Errorf("expected white-label keys to be kept out of attributes, got %v", attributes)
	}
}

// TestSetEnterpriseAttributes_KeepsLegacyMap protects configs written before the
// white_labeling block existed, which set the $ztag keys in attributes directly.
// Migrating their values into the block would leave attributes empty in state while
// the config still lists them, i.e. a diff that never converges.
func TestSetEnterpriseAttributes_KeepsLegacyMap(t *testing.T) {
	d := schema.TestResourceDataRaw(t, EnterpriseSchema(), map[string]interface{}{
		"attributes": map[string]interface{}{
			whiteLabelPrimaryColorKey: "#0A2540",
		},
	})

	setEnterpriseAttributes(d, map[string]string{
		whiteLabelPrimaryColorKey: "#0A2540",
	})

	attributes := d.Get("attributes").(map[string]interface{})
	if attributes[whiteLabelPrimaryColorKey] != "#0A2540" {
		t.Errorf("expected the legacy attributes entry to be kept, got %v", attributes)
	}
	if block := d.Get("white_labeling").([]interface{}); len(block) != 0 {
		t.Errorf("expected no white_labeling block for a legacy config, got %v", block)
	}
}

// TestEnterpriseWhiteLabelingSchema_ValueLength pins the bounds the API enforces on
// enterprise tag values. The 256-character cap is the reason logo_url is a URL: an
// inlined image cannot fit, and without this check the config only fails at apply
// time with an opaque API error.
func TestEnterpriseWhiteLabelingSchema_ValueLength(t *testing.T) {
	logo := EnterpriseWhiteLabelingSchema()["logo_url"]

	if _, errs := logo.ValidateFunc(strings.Repeat("a", 257), "logo_url"); len(errs) == 0 {
		t.Error("expected a value longer than 256 characters to be rejected")
	}
	if _, errs := logo.ValidateFunc("ab", "logo_url"); len(errs) == 0 {
		t.Error("expected a value shorter than 3 characters to be rejected")
	}
	if _, errs := logo.ValidateFunc("https://acme.example.com/logo.svg", "logo_url"); len(errs) != 0 {
		t.Errorf("expected a normal logo URL to be accepted, got %v", errs)
	}
}

// TestEnterpriseSchema_WhiteLabelingConflictsWithAttributes keeps the two write paths
// from overlapping. Both feed the same attribute map, so allowing both in one config
// would make the resulting values depend on merge order rather than on the config.
func TestEnterpriseSchema_WhiteLabelingConflictsWithAttributes(t *testing.T) {
	s := EnterpriseSchema()

	if got := s["white_labeling"].ConflictsWith; len(got) != 1 || got[0] != "attributes" {
		t.Errorf("expected white_labeling to conflict with attributes, got %v", got)
	}
	if got := s["attributes"].ConflictsWith; len(got) != 1 || got[0] != "white_labeling" {
		t.Errorf("expected attributes to conflict with white_labeling, got %v", got)
	}
	if s["attributes"].Deprecated == "" {
		t.Error("expected attributes to be deprecated in favor of white_labeling")
	}
}

// TestEnterpriseModel_MergesWhiteLabeling checks the write path end to end: the block
// has to reach models.Enterprise.Attributes, since that map is the only thing the API
// stores white-labeling in.
func TestEnterpriseModel_MergesWhiteLabeling(t *testing.T) {
	d := schema.TestResourceDataRaw(t, EnterpriseSchema(), map[string]interface{}{
		"name":           "acme",
		"title":          "Acme",
		"white_labeling": whiteLabelingBlock("#0A2540", "", "https://acme.example.com/logo.svg", ""),
	})

	got := EnterpriseModel(d).Attributes

	if got[whiteLabelPrimaryColorKey] != "#0A2540" {
		t.Errorf("expected primary color in attributes, got %v", got)
	}
	if got[whiteLabelLogoKey] != "https://acme.example.com/logo.svg" {
		t.Errorf("expected logo in attributes, got %v", got)
	}
	if _, isSet := got[whiteLabelSecondaryColorKey]; isSet {
		t.Errorf("expected unset secondary color to be absent, got %v", got)
	}
}
