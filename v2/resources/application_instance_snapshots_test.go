package resources

import (
	"context"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
	zschema "github.com/zededa/terraform-provider-zedcloud/v2/schemas"
)

// Unit tests for allow_snapshots, max_snapshots and snapshot_on_update on
// zedcloud_application_instance (PLE-576). They drive the resource through
// Diff and Apply, so HasChange and CustomizeDiff behave as in a real plan.

const snapInstID = "55555555-5555-5555-5555-555555555555"

func baseInstanceConfig() map[string]interface{} {
	return map[string]interface{}{
		"name": "inst-s", "title": "title", "app_id": "app-1", "device_id": "dev-1", "project_id": "proj-1",
	}
}

// seedInstance puts an instance on the fake controller and returns its
// refreshed Terraform state.
func seedInstance(t *testing.T, fc *fakeController, meta *api_client.ZedcloudAPI, extra map[string]interface{}) *terraform.InstanceState {
	t.Helper()
	inst := map[string]interface{}{
		"id": snapInstID, "name": "inst-s", "title": "title", "appId": "app-1", "deviceId": "dev-1",
		"projectId": "proj-1", "activate": true, "maxSnapshots": 1, "PersistentRuntimeSizeBytes": "0",
	}
	for k, v := range extra {
		inst[k] = v
	}
	fc.mu.Lock()
	fc.instances["inst-s"] = inst
	fc.mu.Unlock()

	r := ApplicationInstanceResource()
	d := r.Data(&terraform.InstanceState{ID: snapInstID, Attributes: map[string]string{"id": snapInstID, "name": "inst-s"}})
	if ds := resourceReadApplicationInstance(context.Background(), d, meta); ds.HasError() {
		t.Fatalf("seed read: %v", ds)
	}
	return d.State()
}

// plan runs Diff for config against state.
func plan(t *testing.T, meta *api_client.ZedcloudAPI, state *terraform.InstanceState, cfg map[string]interface{}) (*terraform.InstanceDiff, error) {
	t.Helper()
	return ApplicationInstanceResource().Diff(context.Background(), state, terraform.NewResourceConfigRaw(cfg), meta)
}

// apply plans and applies config against state.
func apply(t *testing.T, meta *api_client.ZedcloudAPI, state *terraform.InstanceState, cfg map[string]interface{}) *terraform.InstanceState {
	t.Helper()
	diff, err := plan(t, meta, state, cfg)
	if err != nil {
		t.Fatalf("plan: %v", err)
	}
	if diff == nil || diff.Empty() {
		return state
	}
	newState, ds := ApplicationInstanceResource().Apply(context.Background(), state, diff, meta)
	if ds.HasError() {
		t.Fatalf("apply: %v", ds)
	}
	return newState
}

func withCfg(extra map[string]interface{}) map[string]interface{} {
	c := baseInstanceConfig()
	for k, v := range extra {
		c[k] = v
	}
	return c
}

func countPrefix(log, prefix string) int {
	n := 0
	for _, l := range strings.Split(log, "\n") {
		if strings.HasPrefix(l, prefix) {
			n++
		}
	}
	return n
}

func TestApplicationInstanceSnapshots_ModelRoundTrip(t *testing.T) {
	d := schema.TestResourceDataRaw(t, zschema.ApplicationInstanceResourceSchema(), withCfg(map[string]interface{}{
		"allow_snapshots": true, "max_snapshots": 2,
	}))
	m := zschema.ApplicationInstanceModel(d)
	if !m.AllowSnapshots || m.MaxSnapshots != 2 {
		t.Fatalf("model: allow %v, max %d", m.AllowSnapshots, m.MaxSnapshots)
	}
	d2 := schema.TestResourceDataRaw(t, zschema.ApplicationInstanceResourceSchema(), baseInstanceConfig())
	zschema.SetApplicationInstanceResourceData(d2, &models.AppInstance{AllowSnapshots: true, MaxSnapshots: 3})
	if d2.Get("allow_snapshots") != true || d2.Get("max_snapshots") != 3 {
		t.Fatalf("state: allow %v, max %v", d2.Get("allow_snapshots"), d2.Get("max_snapshots"))
	}
}

func TestApplicationInstanceSnapshots_UpdateKeepsAllowSnapshotsFromUI(t *testing.T) {
	// Regression: before PLE-576 every update sent no allowSnapshots and the
	// controller reset it to false.
	fc := newFakeController()
	meta := newFakeMeta(t, fc, true)
	state := seedInstance(t, fc, meta, map[string]interface{}{"allowSnapshots": true})

	apply(t, meta, state, withCfg(map[string]interface{}{"title": "new title"}))
	if len(fc.puts) != 1 || fc.puts[0]["allowSnapshots"] != true {
		t.Fatalf("update body must keep allowSnapshots=true: %v", fc.puts)
	}
}

func TestApplicationInstanceSnapshots_NoDiffForReportedMaxSnapshots(t *testing.T) {
	// The controller reports maxSnapshots 0 as 1; an untouched value must not
	// plan a change.
	fc := newFakeController()
	meta := newFakeMeta(t, fc, true)
	state := seedInstance(t, fc, meta, nil)

	diff, err := plan(t, meta, state, baseInstanceConfig())
	if err != nil {
		t.Fatal(err)
	}
	if diff != nil {
		for k := range diff.Attributes {
			if strings.Contains(k, "snapshot") {
				t.Fatalf("unexpected change of %s: %+v", k, diff.Attributes[k])
			}
		}
	}
}

func TestApplicationInstanceSnapshots_SnapshotOnUpdate(t *testing.T) {
	for _, async := range []bool{true, false} {
		fc := newFakeController()
		meta := newFakeMeta(t, fc, async)
		state := seedInstance(t, fc, meta, map[string]interface{}{"allowSnapshots": true})
		state = apply(t, meta, state, withCfg(map[string]interface{}{"allow_snapshots": true, "snapshot_on_update": true}))
		if n := countPrefix(fc.requestLog(), "PUT"); n != 0 {
			t.Fatalf("async=%v: turning on snapshot_on_update alone must not update:\n%s", async, fc.requestLog())
		}

		apply(t, meta, state, withCfg(map[string]interface{}{
			"allow_snapshots": true, "snapshot_on_update": true, "title": "changed",
		}))
		log := fc.requestLog()
		if countPrefix(log, "PUT /v2/apps/instances/id/"+snapInstID+"?makeSnapshot=true") != 1 ||
			countPrefix(log, "PUT") != 1 {
			t.Fatalf("async=%v: expected one v2 update with makeSnapshot:\n%s", async, log)
		}
	}
}

func TestApplicationInstanceSnapshots_TurnOnAllowAndSnapshotInOneApply(t *testing.T) {
	fc := newFakeController()
	meta := newFakeMeta(t, fc, true)
	state := seedInstance(t, fc, meta, nil)

	apply(t, meta, state, withCfg(map[string]interface{}{
		"allow_snapshots": true, "snapshot_on_update": true, "title": "changed",
	}))
	log := fc.requestLog()
	if countPrefix(log, "PUT /v2/apps/instances/id/"+snapInstID+"?makeSnapshot=true") != 1 || countPrefix(log, "PUT") != 2 {
		t.Fatalf("expected two updates, the second with makeSnapshot:\n%s", log)
	}
	first, second := fc.puts[0], fc.puts[1]
	if first["allowSnapshots"] != true || first["title"] != "title" {
		t.Fatalf("first update must only turn on allowSnapshots, got allow %v title %v", first["allowSnapshots"], first["title"])
	}
	if second["title"] != "changed" || second["allowSnapshots"] != true {
		t.Fatalf("second update must carry the new config, got %v", second)
	}
}

func TestApplicationInstanceSnapshots_PlanChecks(t *testing.T) {
	fc := newFakeController()
	meta := newFakeMeta(t, fc, true)
	state := seedInstance(t, fc, meta, nil)

	cases := map[string]struct {
		state   *terraform.InstanceState
		cfg     map[string]interface{}
		wantErr string
	}{
		"snapshot_on_update without allow_snapshots":     {state, withCfg(map[string]interface{}{"snapshot_on_update": true}), "requires allow_snapshots"},
		"new instance, snapshot_on_update without allow": {nil, withCfg(map[string]interface{}{"snapshot_on_update": true}), "requires allow_snapshots"},
		"docker compose": {nil, withCfg(map[string]interface{}{"app_type": "APP_TYPE_DOCKER_COMPOSE", "allow_snapshots": true}), "docker-compose"},
		"cluster allow": {nil, withCfg(map[string]interface{}{
			"device_id": "", "edge_node_cluster": []interface{}{map[string]interface{}{"id": "cluster-1"}}, "allow_snapshots": true,
		}), "edge node cluster cluster-1"},
		"cluster max": {nil, withCfg(map[string]interface{}{
			"device_id": "", "edge_node_cluster": []interface{}{map[string]interface{}{"id": "cluster-1"}}, "max_snapshots": 2,
		}), "edge node cluster"},
		"cluster without settings": {nil, withCfg(map[string]interface{}{
			"device_id": "", "edge_node_cluster": []interface{}{map[string]interface{}{"id": "cluster-1"}},
		}), ""},
		"device with settings": {nil, withCfg(map[string]interface{}{"allow_snapshots": true, "max_snapshots": 2, "snapshot_on_update": true}), ""},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			_, err := plan(t, meta, tc.state, tc.cfg)
			switch {
			case tc.wantErr == "" && err != nil:
				t.Fatalf("unexpected error: %v", err)
			case tc.wantErr != "" && (err == nil || !strings.Contains(err.Error(), tc.wantErr)):
				t.Fatalf("err = %v, want %q", err, tc.wantErr)
			}
		})
	}
}

func TestApplicationInstanceSnapshots_UnsetMaxSnapshotsIsNotSent(t *testing.T) {
	// The controller reports a stored 0 as 1. Sending the 1 back would change
	// what it stores, so an unset max_snapshots of 1 is sent as 0 (omitted).
	fc := newFakeController()
	meta := newFakeMeta(t, fc, true)
	state := seedInstance(t, fc, meta, nil)

	apply(t, meta, state, withCfg(map[string]interface{}{"title": "changed"}))
	if len(fc.puts) != 1 {
		t.Fatalf("puts = %v", fc.puts)
	}
	if v, ok := fc.puts[0]["maxSnapshots"]; ok {
		t.Fatalf("maxSnapshots must not be sent when unset and reported as 1, sent %v", v)
	}

	// A value other than 1 is unambiguous and is kept.
	fc2 := newFakeController()
	meta2 := newFakeMeta(t, fc2, true)
	state2 := seedInstance(t, fc2, meta2, map[string]interface{}{"maxSnapshots": 3})
	apply(t, meta2, state2, withCfg(map[string]interface{}{"title": "changed"}))
	if len(fc2.puts) != 1 || fc2.puts[0]["maxSnapshots"] != float64(3) {
		t.Fatalf("maxSnapshots 3 must be kept: %v", fc2.puts)
	}
}

func TestApplicationInstanceSnapshots_MaxSnapshotsZeroRejected(t *testing.T) {
	fc := newFakeController()
	meta := newFakeMeta(t, fc, true)
	r := ApplicationInstanceResource()
	cfg := terraform.NewResourceConfigRaw(withCfg(map[string]interface{}{"max_snapshots": 0}))
	if diags := r.Validate(cfg); !diags.HasError() {
		t.Fatal("max_snapshots = 0 must be rejected: the controller reads it back as 1")
	}
	_ = meta
}

func TestApplicationInstanceSnapshots_SnapshotUpdateJobFailed(t *testing.T) {
	fc := newFakeController()
	meta := newFakeMeta(t, fc, true)
	state := seedInstance(t, fc, meta, map[string]interface{}{"allowSnapshots": true})
	state = apply(t, meta, state, withCfg(map[string]interface{}{"allow_snapshots": true, "snapshot_on_update": true}))
	fc.failUpdate = "failed to create snapshot"

	cfg := withCfg(map[string]interface{}{"allow_snapshots": true, "snapshot_on_update": true, "title": "changed"})
	diff, err := plan(t, meta, state, cfg)
	if err != nil {
		t.Fatal(err)
	}
	_, ds := ApplicationInstanceResource().Apply(context.Background(), state, diff, meta)
	txt := ""
	for _, d := range ds {
		txt += d.Summary + "\n"
	}
	if !ds.HasError() || !strings.Contains(txt, "failed to create snapshot") || !strings.Contains(txt, "was not applied") {
		t.Fatalf("diags = %s", txt)
	}
	if fc.instances["inst-s"]["title"] != "title" {
		t.Fatal("the fake must not have applied the update")
	}
}

func TestApplicationInstanceSnapshots_SyncNonSnapshotUpdateUsesV1(t *testing.T) {
	fc := newFakeController()
	meta := newFakeMeta(t, fc, false)
	state := seedInstance(t, fc, meta, nil)

	apply(t, meta, state, withCfg(map[string]interface{}{"title": "changed"}))
	log := fc.requestLog()
	if countPrefix(log, "PUT /v1/apps/instances/id/"+snapInstID) != 1 || strings.Contains(log, "/v2/") {
		t.Fatalf("async_writes=false without snapshot_on_update must keep using v1:\n%s", log)
	}
}

func TestApplicationInstanceSnapshots_CreateIgnoresSnapshotOnUpdate(t *testing.T) {
	fc := newFakeController()
	meta := newFakeMeta(t, fc, true)
	apply(t, meta, nil, withCfg(map[string]interface{}{"allow_snapshots": true, "snapshot_on_update": true}))
	log := fc.requestLog()
	if strings.Contains(log, "makeSnapshot") || countPrefix(log, "POST /v2/apps/instances") != 1 {
		t.Fatalf("create must not ask for a snapshot:\n%s", log)
	}
}
