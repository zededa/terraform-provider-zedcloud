package resources

import (
	"context"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	zschema "github.com/zededa/terraform-provider-zedcloud/v2/schemas"
)

// ENG-3120: built-in roles (SysAdmin is USER_ROLE_SYSTEM_DEFINED, SysRoot is
// USER_ROLE_CLUSTER) must not be importable into zedcloud_role.

const (
	roleImportID  = "AAGFABAvyOcsKThCdYCiNNZA1chw" // 28 characters, like a real role ID
	roleMissingID = "AAGFABAvyOcsKThCdYCiNNZA1chx" // same shape, unknown to the controller
)

func roleImportController(typ string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if !strings.Contains(r.URL.Path, "/roles/") {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if strings.Contains(r.URL.Path, "/roles/id/"+roleMissingID) {
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte(`{"httpStatusCode":404,"httpStatusMsg":"not found"}`))
			return
		}
		_, _ = w.Write([]byte(`{"id":"` + roleImportID + `","name":"SysAdmin","title":"SysAdmin","type":"` + typ + `","state":"ROLE_STATE_ACTIVE"}`))
	})
}

func importRole(t *testing.T, typ, importID string) (string, error) {
	t.Helper()
	m := newNodeMeta(t, roleImportController(typ))
	d := (&schema.Resource{Schema: zschema.RoleSchema()}).Data(nil)
	d.SetId(importID)
	out, err := importRoleStateContext(context.Background(), d, m)
	if err != nil {
		return "", err
	}
	return out[0].Id(), nil
}

func TestRoleImport_RefusesBuiltInRoles(t *testing.T) {
	for _, typ := range []string{"USER_ROLE_SYSTEM_DEFINED", "USER_ROLE_CLUSTER"} {
		for name, id := range map[string]string{"by id": roleImportID, "by name": "SysAdmin"} {
			t.Run(typ+" "+name, func(t *testing.T) {
				_, err := importRole(t, typ, id)
				if err == nil {
					t.Fatal("importing a built-in role must fail")
				}
				for _, want := range []string{"SysAdmin", "built-in", typ, "data source"} {
					if !strings.Contains(err.Error(), want) {
						t.Errorf("error %q should mention %q", err, want)
					}
				}
			})
		}
	}
}

func TestRoleImport_AcceptsCustomRoles(t *testing.T) {
	for _, typ := range []string{"USER_ROLE_USER_DEFINED", "USER_ROLE_UNSPECIFIED"} {
		for name, id := range map[string]string{"by id": roleImportID, "by name": "SysAdmin"} {
			t.Run(typ+" "+name, func(t *testing.T) {
				got, err := importRole(t, typ, id)
				if err != nil {
					t.Fatalf("custom role import failed: %v", err)
				}
				if got != roleImportID {
					t.Errorf("imported id %q, want %q", got, roleImportID)
				}
			})
		}
	}
}

func TestRoleImport_UnknownIDIsReported(t *testing.T) {
	_, err := importRole(t, "USER_ROLE_USER_DEFINED", roleMissingID)
	if err == nil || !strings.Contains(err.Error(), "no role found") {
		t.Errorf("want a 'no role found' error, got %v", err)
	}
}
