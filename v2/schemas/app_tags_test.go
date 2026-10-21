package schemas

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestValidateAppTags(t *testing.T) {
	many := map[string]interface{}{}
	for i := 0; i < AppTagLimit+1; i++ {
		many[fmt.Sprintf("key%02d", i)] = "value"
	}
	cases := []struct {
		name    string
		in      map[string]interface{}
		wantErr string
	}{
		{"empty", map[string]interface{}{}, ""},
		{"valid", map[string]interface{}{"customer": "acme", "tier": "gold"}, ""},
		{"key too short", map[string]interface{}{"ab": "value"}, "key must be"},
		{"value too short", map[string]interface{}{"key": "ab"}, "value must be"},
		{"key with slash", map[string]interface{}{"a/b": "value"}, "key must not contain"},
		{"key with quote", map[string]interface{}{"ab'c": "value"}, "key must not contain"},
		{"value with quote", map[string]interface{}{"key": "it's"}, "value must not contain"},
		{"value with backslash", map[string]interface{}{"key": `a\b`}, "value must not contain"},
		{"value with slash ok", map[string]interface{}{"key": "a/b/c"}, ""},
		{"unknown value skipped", map[string]interface{}{"key": unknownValue}, ""},
		{"too many", many, "Too many tags"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			diags := ValidateAppTags(tc.in, cty.GetAttrPath("tags"))
			if tc.wantErr == "" {
				if diags.HasError() {
					t.Fatalf("unexpected diagnostics: %+v", diags)
				}
				return
			}
			if !diags.HasError() {
				t.Fatalf("expected an error containing %q", tc.wantErr)
			}
			var all []string
			for _, d := range diags {
				all = append(all, d.Summary+": "+d.Detail)
			}
			if !strings.Contains(strings.Join(all, "\n"), tc.wantErr) {
				t.Fatalf("diagnostics %q do not contain %q", all, tc.wantErr)
			}
		})
	}
}

func TestApplicationTagsRoundTrip(t *testing.T) {
	d := schema.TestResourceDataRaw(t, Application(), map[string]interface{}{
		"name":  "app1",
		"title": "app1",
		"tags":  map[string]interface{}{"customer": "acme"},
	})
	app := ApplicationModel(d)
	if !reflect.DeepEqual(app.Tags, map[string]string{"customer": "acme"}) {
		t.Fatalf("tags = %v", app.Tags)
	}

	app.Tags = map[string]string{"customer": "acme", "tier": "gold"}
	SetApplicationResourceData(d, app)
	if got := d.Get("tags").(map[string]interface{}); len(got) != 2 || got["tier"] != "gold" {
		t.Fatalf("state tags = %v", got)
	}

	empty := schema.TestResourceDataRaw(t, Application(), map[string]interface{}{"name": "a", "title": "a"})
	if tags := ApplicationModel(empty).Tags; tags != nil {
		t.Fatalf("no tags configured must send no tags field, got %v", tags)
	}
}
