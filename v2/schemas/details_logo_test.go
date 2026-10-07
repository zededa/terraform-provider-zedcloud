package schemas

import (
	"testing"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

func TestValidateAppLogo(t *testing.T) {
	path := cty.GetAttrPath("logo")
	cases := []struct {
		name          string
		in            map[string]interface{}
		errors, warns int
	}{
		{name: "empty", in: map[string]interface{}{}},
		{name: "artifact id under logo", in: map[string]interface{}{"logo": "0a12-c1ae_logo.png"}},
		{name: "unknown at plan time", in: map[string]interface{}{"logo": unknownValue}},
		{name: "https url", in: map[string]interface{}{"logo": "https://example.com/l.png"}, errors: 1},
		{name: "data url", in: map[string]interface{}{"logo": "data:image/png;base64,AAAA"}, errors: 1},
		{name: "uppercase scheme", in: map[string]interface{}{"logo": "HTTP://example.com/l.png"}, errors: 1},
		{name: "wrong key", in: map[string]interface{}{"logo.png": "0a12_logo.png"}, warns: 1},
		{name: "two entries", in: map[string]interface{}{"logo": "a_1", "zfill": "#fff"}, warns: 1},
		{name: "url and wrong key", in: map[string]interface{}{"icon": "https://x"}, errors: 1, warns: 1},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var errs, warns int
			for _, d := range ValidateAppLogo(tc.in, path) {
				switch d.Severity {
				case diag.Error:
					errs++
				case diag.Warning:
					warns++
				}
			}
			if errs != tc.errors || warns != tc.warns {
				t.Errorf("got %d errors, %d warnings; want %d, %d", errs, warns, tc.errors, tc.warns)
			}
		})
	}
}

func TestDetailsModelFromMap_Tolerant(t *testing.T) {
	// Missing optional keys must not panic (they used to be unchecked assertions).
	got := DetailsModelFromMap(map[string]interface{}{
		"app_category": "APP_CATEGORY_OTHERS",
		"logo":         map[string]interface{}{"logo": "0a12_logo.png", "nil": nil},
	})
	if got.Logo["logo"] != "0a12_logo.png" || len(got.Logo) != 1 {
		t.Errorf("logo = %v", got.Logo)
	}
	if got.Os != "" || got.Support != "" || got.Category == nil || *got.Category != "" {
		t.Errorf("unexpected defaults: os=%q support=%q category=%v", got.Os, got.Support, got.Category)
	}
}

func TestArtifactName(t *testing.T) {
	for name, ok := range map[string]bool{
		"logo": true, "myapp-logo.png": true, "a_b.c-d": true,
		"ab": false, "-logo": false, ".png": false, "has space": false, "": false,
	} {
		if got := artifactNameRegexp.MatchString(name); got != ok {
			t.Errorf("%q: match=%v, want %v", name, got, ok)
		}
	}
}
