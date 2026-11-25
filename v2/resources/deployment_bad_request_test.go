package resources

import (
	"context"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	zschema "github.com/zededa/terraform-provider-zedcloud/v2/schemas"
)

// ENG-3119: the controller answers a rejected deployment create or update with
// a 400 whose body is a ZsrvResponse and carries the reason in httpStatusMsg.
// The generated client had no case for 400, decoded the body as a
// googlerpc.Status and the user saw only {"details":null}.

const (
	depErrProject = "0b6f4b1e-3f6c-4c36-9d9f-4a4a3b3c1f11"
	depErrOldID   = "8e9c48c7-0514-458a-b414-7ab2de1e9199"
	depErrReason  = `failed to parse json request body: proto: (line 1:2482): invalid value for enum field namingScheme: "APP_NAMING_SCHEME_BOGUS"`
)

// depErrController answers every deployment write with status and body, and
// the version listing with one existing version (so an update is allowed).
func depErrController(status int, body string) http.Handler {
	return depErrControllerWithVersions(status, body, true)
}

// depErrControllerWithVersions: existing=false lists no version, so a create
// goes to POST (the first deployment of a project).
func depErrControllerWithVersions(status int, body string, existing bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == http.MethodGet && !existing {
			_, _ = w.Write([]byte(`{"list":[]}`))
			return
		}
		if r.Method == http.MethodGet {
			_, _ = w.Write([]byte(`{"list":[{"id":"` + depErrOldID + `","revision":{"curr":"1","createdAt":"2026-11-25T00:00:00Z","createdBy":"tf","updatedAt":"2026-11-25T00:00:00Z","updatedBy":"tf"}}]}`))
			return
		}
		w.WriteHeader(status)
		_, _ = w.Write([]byte(body))
	})
}

const depErrBody400 = `{"httpStatusCode":400,"httpStatusMsg":"` + `failed to parse json request body: proto: (line 1:2482): invalid value for enum field namingScheme: \"APP_NAMING_SCHEME_BOGUS\"` + `","error":[{"ec":"BadReqBody","details":"request body parsing failed"}]}`

func depErrData(t *testing.T, id string) *schema.ResourceData {
	t.Helper()
	d := schema.TestResourceDataRaw(t, zschema.Deployment(), map[string]interface{}{
		"name": "dep", "title": "dep", "project_id": depErrProject, "deployment_tag": "v2",
	})
	if id != "" {
		d.SetId(id)
		_ = d.Set("id", id)
	}
	return d
}

func joinDiags(diags diag.Diagnostics) string {
	var parts []string
	for _, d := range diags {
		parts = append(parts, d.Summary+" "+d.Detail)
	}
	return strings.Join(parts, " | ")
}

// zschemaValidateNamingScheme runs the validation of the deployment's
// app_inst_config.naming_scheme attribute.
func zschemaValidateNamingScheme(v string) diag.Diagnostics {
	cfg := zschema.AppInstConfig()["naming_scheme"]
	return cfg.ValidateDiagFunc(v, cty.GetAttrPath("naming_scheme"))
}

func TestDeployment_BadRequest_ReasonReachesUser(t *testing.T) {
	cases := []struct {
		name string
		run  func(m interface{}) string
	}{
		{"create", func(m interface{}) string {
			diags := CreateDeployment(context.Background(), depErrData(t, ""), m)
			return joinDiags(diags)
		}},
		{"update", func(m interface{}) string {
			// the list in depErrController has no list for create; create sees an
			// existing version too and takes the new-version path, which is fine.
			diags := UpdateDeployment(context.Background(), depErrData(t, depErrOldID), m)
			return joinDiags(diags)
		}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			m := newNodeMeta(t, depErrController(http.StatusBadRequest, depErrBody400))
			got := c.run(m)
			if !strings.Contains(got, depErrReason) {
				t.Errorf("the controller's httpStatusMsg must reach the user, got: %s", got)
			}
			if strings.Contains(got, `{"details":null}`) {
				t.Errorf("must not show the empty googlerpc status, got: %s", got)
			}
			if !strings.Contains(got, "400") {
				t.Errorf("the status code should be shown, got: %s", got)
			}
		})
	}
}

func TestDeployment_BadRequest_FirstCreate_ReasonReachesUser(t *testing.T) {
	m := newNodeMeta(t, depErrControllerWithVersions(http.StatusBadRequest, depErrBody400, false))
	got := joinDiags(CreateDeployment(context.Background(), depErrData(t, ""), m))
	if !strings.Contains(got, depErrReason) || strings.Contains(got, `{"details":null}`) {
		t.Errorf("the POST create path must show the controller's reason, got: %s", got)
	}
}

func TestDeployment_NotFoundAndConflict_ReasonReachesUser(t *testing.T) {
	for _, c := range []struct {
		status int
		msg    string
	}{{http.StatusNotFound, "project not found"}, {http.StatusConflict, "deployment tag already used"}} {
		body := `{"httpStatusCode":` + itoa(c.status) + `,"httpStatusMsg":"` + c.msg + `"}`
		m := newNodeMeta(t, depErrController(c.status, body))
		got := joinDiags(UpdateDeployment(context.Background(), depErrData(t, depErrOldID), m))
		if !strings.Contains(got, c.msg) {
			t.Errorf("status %d: reason %q missing from: %s", c.status, c.msg, got)
		}
	}
}

func TestDeploymentSchema_NamingSchemeIsValidated(t *testing.T) {
	for _, tc := range []struct {
		value string
		ok    bool
	}{
		{"APP_NAMING_SCHEME_BOGUS", false},
		{"app_naming_scheme_device", false},
		{"APP_NAMING_SCHEME_PROJECT_APP_DEVICE", true},
		{"APP_NAMING_SCHEME_DEVICE", true},
		{"", true},
	} {
		diags := zschemaValidateNamingScheme(tc.value)
		if tc.ok && diags.HasError() {
			t.Errorf("%q must be accepted, got %v", tc.value, diags)
		}
		if !tc.ok && !diags.HasError() {
			t.Errorf("%q must be rejected at plan time", tc.value)
		}
	}
}
