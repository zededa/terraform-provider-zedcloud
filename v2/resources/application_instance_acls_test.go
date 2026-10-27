package resources

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	zschema "github.com/zededa/terraform-provider-zedcloud/v2/schemas"
)

// UE-203 / ENG-3000: interface ACLs on zedcloud_application_instance.

func ace(matches [][2]string, mapPort int) map[string]interface{} {
	ms := make([]interface{}, 0, len(matches))
	for _, m := range matches {
		ms = append(ms, map[string]interface{}{"type": m[0], "value": m[1]})
	}
	a := map[string]interface{}{"matches": ms}
	if mapPort != 0 {
		a["actions"] = []interface{}{map[string]interface{}{
			"portmap":   true,
			"mapparams": []interface{}{map[string]interface{}{"port": mapPort}},
		}}
	}
	return a
}

var (
	aceAllow = ace([][2]string{{"ip", "0.0.0.0/0"}}, 0)
	ace9022  = ace([][2]string{{"protocol", "tcp"}, {"lport", "9022"}, {"ip", "0.0.0.0/0"}}, 22)
	ace8022  = ace([][2]string{{"protocol", "tcp"}, {"lport", "8022"}, {"ip", "0.0.0.0/0"}}, 22)
	ace9080  = ace([][2]string{{"protocol", "tcp"}, {"lport", "9080"}, {"ip", "0.0.0.0/0"}}, 80)
)

// intf builds one interfaces element; acls == nil leaves the block out.
func intf(name string, acls ...map[string]interface{}) map[string]interface{} {
	i := map[string]interface{}{"intfname": name, "netinstname": "ni-1"}
	if acls != nil {
		l := make([]interface{}, 0, len(acls))
		for _, a := range acls {
			l = append(l, a)
		}
		i["acls"] = l
	}
	return i
}

func instanceConfig(intfs ...map[string]interface{}) map[string]interface{} {
	l := make([]interface{}, 0, len(intfs))
	for _, i := range intfs {
		l = append(l, i)
	}
	return map[string]interface{}{
		"name":       "ai",
		"title":      "ai",
		"app_id":     "app-1",
		"device_id":  "dev-1",
		"project_id": "proj-1",
		"interfaces": l,
	}
}

// stubOrdering replaces the project lookup for one test and counts calls.
func stubOrdering(t *testing.T, enabled bool, err error) *int {
	t.Helper()
	calls := 0
	orig := projectInterfaceOrdering
	projectInterfaceOrdering = func(_ context.Context, _ interface{}, projectID string) (bool, error) {
		calls++
		if projectID != "proj-1" {
			t.Errorf("project lookup for %q, want proj-1", projectID)
		}
		return enabled, err
	}
	t.Cleanup(func() { projectInterfaceOrdering = orig })
	return &calls
}

// planFrom plans cfg against an existing instance whose state is stored.
func planFrom(t *testing.T, stored, cfg map[string]interface{}) *terraform.InstanceDiff {
	t.Helper()
	r := ApplicationInstanceResource()
	d := schema.TestResourceDataRaw(t, r.Schema, stored)
	d.SetId("ai-1")
	diff, err := r.Diff(context.Background(), d.State(), terraform.NewResourceConfigRaw(cfg), nil)
	if err != nil {
		t.Fatalf("diff: %s", err)
	}
	return diff
}

func diffKeys(diff *terraform.InstanceDiff, prefix string) []string {
	var keys []string
	if diff == nil {
		return nil
	}
	for k := range diff.Attributes {
		if strings.HasPrefix(k, prefix) {
			keys = append(keys, k)
		}
	}
	return keys
}

func TestAppInstanceACLs_EditIsPlannedWithOrdering(t *testing.T) {
	calls := stubOrdering(t, true, nil)
	diff := planFrom(t,
		instanceConfig(intf("eth0", aceAllow, ace9022, ace9080)),
		instanceConfig(intf("eth0", aceAllow, ace8022, ace9080)))
	if keys := diffKeys(diff, "interfaces.0.acls.1."); len(keys) == 0 {
		t.Fatalf("no diff for the edited rule; diff: %v", diff)
	}
	if *calls != 1 {
		t.Errorf("project lookups = %d, want 1", *calls)
	}
}

func TestAppInstanceACLs_RemovedRuleIsPlanned(t *testing.T) {
	stubOrdering(t, true, nil)
	diff := planFrom(t,
		instanceConfig(intf("eth0", aceAllow, ace9022, ace9080)),
		instanceConfig(intf("eth0", aceAllow, ace9080)))
	if keys := diffKeys(diff, "interfaces.0.acls.#"); len(keys) == 0 {
		t.Fatalf("no count diff for the removed rule; diff: %v", diff)
	}
}

func TestAppInstanceACLs_EditIsClearedWithoutOrdering(t *testing.T) {
	calls := stubOrdering(t, false, nil)
	diff := planFrom(t,
		instanceConfig(intf("eth0", aceAllow, ace9022, ace9080)),
		instanceConfig(intf("eth0", aceAllow, ace8022, ace9080)))
	if keys := diffKeys(diff, "interfaces"); len(keys) != 0 {
		t.Fatalf("ACL diff kept on an ordering-disabled project: %v", keys)
	}
	if *calls != 1 {
		t.Errorf("project lookups = %d, want 1", *calls)
	}
}

func TestAppInstanceACLs_LookupErrorKeepsDiff(t *testing.T) {
	stubOrdering(t, false, errors.New("boom"))
	diff := planFrom(t,
		instanceConfig(intf("eth0", aceAllow, ace9022)),
		instanceConfig(intf("eth0", aceAllow, ace8022)))
	if keys := diffKeys(diff, "interfaces.0.acls."); len(keys) == 0 {
		t.Fatal("ACL diff dropped although the project could not be read")
	}
}

func TestAppInstanceACLs_OmittedACLsKeepStoredList(t *testing.T) {
	calls := stubOrdering(t, true, nil)
	diff := planFrom(t,
		instanceConfig(intf("eth0", aceAllow, ace9022, ace9080)),
		instanceConfig(intf("eth0")))
	if keys := diffKeys(diff, "interfaces"); len(keys) != 0 {
		t.Fatalf("omitted acls planned a change: %v", keys)
	}
	if *calls != 0 {
		t.Errorf("project lookups = %d, want 0 without an ACL change", *calls)
	}
}

func TestAppInstanceACLs_EqualListNoDiffNoLookup(t *testing.T) {
	calls := stubOrdering(t, true, nil)
	diff := planFrom(t,
		instanceConfig(intf("eth0", aceAllow, ace9022)),
		instanceConfig(intf("eth0", aceAllow, ace9022)))
	if keys := diffKeys(diff, "interfaces"); len(keys) != 0 {
		t.Fatalf("unexpected diff: %v", keys)
	}
	if *calls != 0 {
		t.Errorf("project lookups = %d, want 0", *calls)
	}
}

func TestAppInstanceACLs_OtherInterfaceChangesStaySuppressed(t *testing.T) {
	calls := stubOrdering(t, true, nil)
	stored := instanceConfig(intf("eth0", aceAllow))
	cfg := instanceConfig(intf("eth0", aceAllow))
	cfg["interfaces"].([]interface{})[0].(map[string]interface{})["netinstname"] = "ni-2"
	if keys := diffKeys(planFrom(t, stored, cfg), "interfaces"); len(keys) != 0 {
		t.Fatalf("non-ACL interface change planned: %v", keys)
	}
	// An added interface, ACLs included, stays suppressed too (NFR-116).
	added := instanceConfig(intf("eth0", aceAllow), intf("eth1", aceAllow))
	if keys := diffKeys(planFrom(t, stored, added), "interfaces"); len(keys) != 0 {
		t.Fatalf("interface add planned: %v", keys)
	}
	// A renamed interface at the same index is not an ACL edit.
	renamed := instanceConfig(intf("eth1", ace9022))
	if keys := diffKeys(planFrom(t, stored, renamed), "interfaces"); len(keys) != 0 {
		t.Fatalf("renamed interface planned: %v", keys)
	}
	if *calls != 0 {
		t.Errorf("project lookups = %d, want 0", *calls)
	}
}

func TestAppInstanceACLs_OnlyChangedInterfaceIsCleared(t *testing.T) {
	stubOrdering(t, false, nil)
	diff := planFrom(t,
		instanceConfig(intf("eth0", aceAllow), intf("eth1", aceAllow, ace9022)),
		instanceConfig(intf("eth0", aceAllow), intf("eth1", aceAllow, ace8022)))
	if keys := diffKeys(diff, "interfaces"); len(keys) != 0 {
		t.Fatalf("ACL diff kept: %v", keys)
	}
}

func TestAppInstanceACLs_SchemaComputedOnResourceOnly(t *testing.T) {
	if zschema.AppInterface()["acls"].Computed {
		t.Error("shared AppInterface() acls is Computed; deployments and app profiles must not change")
	}
	r := ApplicationInstanceResource()
	acls := r.Schema["interfaces"].Elem.(*schema.Resource).Schema["acls"]
	if !acls.Optional || !acls.Computed {
		t.Errorf("resource interfaces.acls Optional=%v Computed=%v, want both", acls.Optional, acls.Computed)
	}
	ds := ApplicationInstanceDataSource().Schema["interfaces"].Elem.(*schema.Resource).Schema["acls"]
	if ds.Computed {
		t.Error("data source interfaces.acls became Computed")
	}
}

func readBackData(t *testing.T, stored ...map[string]interface{}) *schema.ResourceData {
	t.Helper()
	d := schema.TestResourceDataRaw(t, ApplicationInstanceResource().Schema, instanceConfig(stored...))
	d.SetId("ai-1")
	return d
}

// sentInterfaces is d.Get("interfaces") of a configuration, as Create and
// Update capture it before the write.
func sentInterfaces(t *testing.T, intfs ...map[string]interface{}) []interface{} {
	t.Helper()
	d := schema.TestResourceDataRaw(t, ApplicationInstanceResource().Schema, instanceConfig(intfs...))
	return d.Get("interfaces").([]interface{})
}

func hasWarning(diags diag.Diagnostics) bool {
	for _, d := range diags {
		if d.Severity == diag.Warning {
			return true
		}
	}
	return false
}

func TestAppInstanceACLs_ReadBackWarnsOnAppendedBundleRules(t *testing.T) {
	stubOrdering(t, true, nil)
	// a controller without #8548 appends the bundle's 9022 rule
	d := readBackData(t, intf("eth0", aceAllow, ace8022, ace9022, ace9080))
	diags := aclReadBackDiags(context.Background(), d, nil, sentInterfaces(t, intf("eth0", aceAllow, ace8022, ace9080)))
	if !hasWarning(diags) {
		t.Fatalf("no warning, diags: %v", diags)
	}
	if !strings.Contains(diags[0].Detail, "eth0") {
		t.Errorf("warning does not name the interface: %s", diags[0].Detail)
	}
}

func TestAppInstanceACLs_ReadBackQuietCases(t *testing.T) {
	sent := sentInterfaces(t, intf("eth0", aceAllow, ace8022))
	t.Run("stored as sent", func(t *testing.T) {
		calls := stubOrdering(t, true, nil)
		d := readBackData(t, intf("eth0", aceAllow, ace8022))
		if diags := aclReadBackDiags(context.Background(), d, nil, sent); len(diags) != 0 {
			t.Fatalf("diags: %v", diags)
		}
		if *calls != 0 {
			t.Errorf("project lookups = %d, want 0", *calls)
		}
	})
	t.Run("ordering disabled", func(t *testing.T) {
		stubOrdering(t, false, nil)
		d := readBackData(t, intf("eth0", aceAllow, ace8022, ace9022))
		if diags := aclReadBackDiags(context.Background(), d, nil, sent); len(diags) != 0 {
			t.Fatalf("diags: %v", diags)
		}
	})
	t.Run("acls omitted", func(t *testing.T) {
		stubOrdering(t, true, nil)
		d := readBackData(t, intf("eth0", aceAllow, ace9022))
		if diags := aclReadBackDiags(context.Background(), d, nil, sentInterfaces(t, intf("eth0"))); len(diags) != 0 {
			t.Fatalf("diags: %v", diags)
		}
	})
}

func TestAppInstanceACLs_ProjectOrderingLookup(t *testing.T) {
	fc := newFakeController()
	fc.projectOrdering = map[string]string{
		"p-on":  "INTERFACE_ORDERING_ENABLED",
		"p-off": "INTERFACE_ORDERING_DISABLED",
	}
	meta := newFakeMeta(t, fc, true)
	ctx := context.Background()
	for _, tc := range []struct {
		project string
		want    bool
	}{{"p-on", true}, {"p-off", false}, {"p-on", true}} {
		got, err := projectInterfaceOrdering(ctx, meta, tc.project)
		if err != nil || got != tc.want {
			t.Errorf("%s: got %v, %v; want %v", tc.project, got, err, tc.want)
		}
	}
	if n := strings.Count(fc.requestLog(), "GET /v1/projects/id/p-on"); n != 1 {
		t.Errorf("p-on read %d times, want 1 (cached)", n)
	}
	if _, err := projectInterfaceOrdering(ctx, meta, "p-missing"); err == nil {
		t.Error("missing project: no error")
	}
}
