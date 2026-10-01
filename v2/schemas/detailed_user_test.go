package schemas

import (
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

// userData builds a ResourceData the way the SDK does at runtime, from an
// InstanceState, so that nested blocks and maps are visible to d.GetOk the
// same way they are during a real plan.
func userData(attrs map[string]string) *schema.ResourceData {
	r := &schema.Resource{Schema: DetailedUserSchema()}
	return r.Data(&terraform.InstanceState{
		ID:         "AAGFABAEqnH4je5PHZTXSmHOs-XC",
		Attributes: attrs,
	})
}

func baseUserAttrs() map[string]string {
	return map[string]string{
		"id":       "AAGFABAEqnH4je5PHZTXSmHOs-XC",
		"username": "alice@corp.com",
		"email":    "alice@corp.com",
		"role_id":  "BBGFABAEqnH4je5PHZTXSmHOs-YD",
		"type":     "AUTH_TYPE_LOCAL",
	}
}

// NFR-165 §3.4. DetailedUserModel read the camelCase swagger key instead of the
// snake_case schema key, so custom_user_input was silently dropped on every
// create and update. After an import populated it from the API this produced a
// diff that never converged.
func TestDetailedUserModel_CustomUserInput(t *testing.T) {
	t.Run("round-trips from configuration", func(t *testing.T) {
		attrs := baseUserAttrs()
		attrs["custom_user_input.%"] = "2"
		attrs["custom_user_input.department"] = "networking"
		attrs["custom_user_input.costCenter"] = "42"

		got := DetailedUserModel(userData(attrs))

		if len(got.CustomUserInput) != 2 {
			t.Fatalf("CustomUserInput = %v, want 2 entries", got.CustomUserInput)
		}
		if got.CustomUserInput["department"] != "networking" {
			t.Errorf("CustomUserInput[department] = %q, want networking", got.CustomUserInput["department"])
		}
		if got.CustomUserInput["costCenter"] != "42" {
			t.Errorf("CustomUserInput[costCenter] = %q, want 42", got.CustomUserInput["costCenter"])
		}
	})

	t.Run("absent from configuration", func(t *testing.T) {
		got := DetailedUserModel(userData(baseUserAttrs()))

		if len(got.CustomUserInput) != 0 {
			t.Errorf("CustomUserInput = %v, want empty", got.CustomUserInput)
		}
	})
}

// NFR-165 §3.6. last_login_time and last_logout_time are server-observed and
// must never be sent back to the API. The schema declares them as strings, so
// the old strfmt.DateTime type assertion in the model builder could not have
// worked in any case.
func TestDetailedUserModel_OmitsServerObservedTimestamps(t *testing.T) {
	attrs := baseUserAttrs()
	attrs["last_login_time"] = "2026-09-01T10:00:00.000Z"
	attrs["last_logout_time"] = "2026-09-01T18:00:00.000Z"

	got := DetailedUserModel(userData(attrs))

	if !got.LastLoginTime.Equal(strfmt.DateTime{}) {
		t.Errorf("LastLoginTime = %v, want the zero value", got.LastLoginTime)
	}
	if !got.LastLogoutTime.Equal(strfmt.DateTime{}) {
		t.Errorf("LastLogoutTime = %v, want the zero value", got.LastLogoutTime)
	}
}

// NFR-165 §4.2. username identifies the account and cannot be changed in place,
// so a change must plan a replacement rather than an update the API rejects.
func TestDetailedUserSchema_ImmutableIdentifiers(t *testing.T) {
	s := DetailedUserSchema()

	if !s["username"].ForceNew {
		t.Error("username must be ForceNew: the API cannot rename a user in place")
	}

	// Server-observed fields must not be settable, or Terraform will plan a
	// removal for every imported user whose configuration omits them.
	for _, key := range []string{"last_login_time", "last_logout_time"} {
		if !s[key].Computed {
			t.Errorf("%s must be Computed", key)
		}
		if s[key].Optional {
			t.Errorf("%s must not be Optional", key)
		}
		if s[key].DiffSuppressFunc != nil {
			t.Errorf("%s must not need a DiffSuppressFunc once it is Computed", key)
		}
	}
}

// NFR-165 §3. An imported user is only usable if what Read writes into state is
// what the model builder sends back. Anything the setter misses is silently
// dropped on the next apply.
func TestSetDetailedUserResourceData_RoundTrip(t *testing.T) {
	username := "alice@corp.com"
	email := "alice@corp.com"
	roleID := "BBGFABAEqnH4je5PHZTXSmHOs-YD"
	authType := models.NewAuthType(models.AuthTypeAUTHTYPELOCAL)

	from := &models.DetailedUser{
		ID:              "AAGFABAEqnH4je5PHZTXSmHOs-XC",
		Username:        &username,
		Email:           &email,
		RoleID:          &roleID,
		Type:            authType,
		FirstName:       "Alice",
		FullName:        "Alice Example",
		Locale:          "en_US",
		NotifyPref:      "email",
		Phone:           "+15550100",
		TimeZone:        "America/Los_Angeles",
		CustomUserInput: map[string]string{"department": "networking"},
	}

	d := userData(map[string]string{})
	SetDetailedUserResourceData(d, from)

	back := DetailedUserModel(d)

	if back.Username == nil || *back.Username != username {
		t.Errorf("Username = %v, want %q", back.Username, username)
	}
	if back.Email == nil || *back.Email != email {
		t.Errorf("Email = %v, want %q", back.Email, email)
	}
	if back.RoleID == nil || *back.RoleID != roleID {
		t.Errorf("RoleID = %v, want %q", back.RoleID, roleID)
	}
	if back.Type == nil || *back.Type != *authType {
		t.Errorf("Type = %v, want %v", back.Type, *authType)
	}
	if back.FirstName != from.FirstName {
		t.Errorf("FirstName = %q, want %q", back.FirstName, from.FirstName)
	}
	if back.FullName != from.FullName {
		t.Errorf("FullName = %q, want %q", back.FullName, from.FullName)
	}
	if back.Locale != from.Locale {
		t.Errorf("Locale = %q, want %q", back.Locale, from.Locale)
	}
	if back.NotifyPref != from.NotifyPref {
		t.Errorf("NotifyPref = %q, want %q", back.NotifyPref, from.NotifyPref)
	}
	if back.Phone != from.Phone {
		t.Errorf("Phone = %q, want %q", back.Phone, from.Phone)
	}
	if back.TimeZone != from.TimeZone {
		t.Errorf("TimeZone = %q, want %q", back.TimeZone, from.TimeZone)
	}
	// The defect this whole ticket turns on: without the key fix the map comes
	// back empty here and the next plan shows the same diff for ever.
	if back.CustomUserInput["department"] != "networking" {
		t.Errorf("CustomUserInput = %v, want department=networking to survive the round trip", back.CustomUserInput)
	}
}

// NFR-165 §12.7. The controller requires a custom-parameter key to be three
// '_'-separated segments (ValidateCustomParam, zedcloud libs/zutils). Before
// this validator the only signal was "HTTP status code: 400 / Invalid
// characters in the key" at apply time, naming neither the attribute nor the
// key.
func TestValidateCustomUserInput(t *testing.T) {
	cases := []struct {
		name    string
		entries map[string]interface{}
		wantErr bool
	}{
		{"three segments", map[string]interface{}{"acme_ui_theme": "dark"}, false},
		{"punctuation allowed by the controller", map[string]interface{}{"a-b_c.d_e@f": "x"}, false},
		{"several valid keys", map[string]interface{}{"a_b_c": "1", "d_e_f": "2"}, false},
		{"empty map", map[string]interface{}{}, false},

		{"one segment", map[string]interface{}{"department": "networking"}, true},
		{"two segments", map[string]interface{}{"acme_theme": "dark"}, true},
		{"four segments", map[string]interface{}{"a_b_c_d": "x"}, true},
		{"space in key", map[string]interface{}{"a_b c_d": "x"}, true},
		{"empty value", map[string]interface{}{"a_b_c": ""}, true},
		{"one bad key among good ones", map[string]interface{}{"a_b_c": "1", "bad": "2"}, true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			diags := validateCustomUserInput(tc.entries, cty.Path{})
			if got := diags.HasError(); got != tc.wantErr {
				t.Errorf("validateCustomUserInput(%v) error = %v, want %v (%v)",
					tc.entries, got, tc.wantErr, diags)
			}
		})
	}
}
