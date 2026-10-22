package resources

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	snapshots "github.com/zededa/terraform-provider-zedcloud/v2/client/app_instance_snapshot"
	zschema "github.com/zededa/terraform-provider-zedcloud/v2/schemas"
)

// fakeSnapshotController is an in-memory zedcloud for the app instance
// snapshot routes (PLE-576). States are served from a per-snapshot queue; the
// last entry repeats.
type fakeSnapshotController struct {
	mu        sync.Mutex
	instances map[string]map[string]interface{} // by id
	rawStatus map[string]string                 // device id -> rawStatus
	snaps     map[string]map[string]interface{} // by snapshot id
	states    map[string][]snapshots.State      // by snapshot id
	requests  []string
	next      int

	// noObjectID makes create answer like a controller older than ENG-1676.
	noObjectID bool
	// createStates is the state queue given to each new snapshot.
	createStates []snapshots.State
	// activateStates replaces a snapshot's state queue on activate.
	activateStates []snapshots.State
}

func newFakeSnapshotController() *fakeSnapshotController {
	return &fakeSnapshotController{
		instances: map[string]map[string]interface{}{
			"app-1": {"id": "app-1", "name": "inst-1", "title": "t", "appId": "a", "deviceId": "dev-1", "projectId": "p"},
			"app-cluster": {"id": "app-cluster", "name": "inst-c", "title": "t", "appId": "a", "deviceId": "",
				"projectId": "p", "edgeNodeCluster": map[string]interface{}{"id": "cluster-1"}},
		},
		// EVE reports the highest level it implements (here 21 >= 13).
		rawStatus: map[string]string{"dev-1": `{"apiCapability":"API_CAPABILITY_REPORT_TPM_EVENTLOG"}`},
		snaps:     map[string]map[string]interface{}{},
		states:    map[string][]snapshots.State{},
		createStates: []snapshots.State{
			{State: snapshots.StateCreating, Timestamp: "t1"},
			{State: snapshots.StateCreated, Timestamp: "t2"},
		},
		activateStates: []snapshots.State{
			{State: snapshots.StateRollingBack, Timestamp: "t3"},
			{State: snapshots.StateActive, Timestamp: "t4"},
		},
	}
}

func (f *fakeSnapshotController) addSnapshot(appID, name string, states ...snapshots.State) string {
	f.next++
	id := fmt.Sprintf("snap-%d", f.next)
	f.snaps[id] = map[string]interface{}{"id": id, "name": name, "config": f.instances[appID]}
	f.states[id] = states
	return id
}

func (f *fakeSnapshotController) snapshotByName(appID, name string) (string, map[string]interface{}) {
	for id, s := range f.snaps {
		if s["name"] == name && s["config"].(map[string]interface{})["id"] == appID {
			return id, s
		}
	}
	return "", nil
}

func (f *fakeSnapshotController) log() string {
	f.mu.Lock()
	defer f.mu.Unlock()
	return strings.Join(f.requests, "\n")
}

func (f *fakeSnapshotController) count(prefix string) int {
	n := 0
	for _, r := range strings.Split(f.log(), "\n") {
		if strings.HasPrefix(r, prefix) {
			n++
		}
	}
	return n
}

func (f *fakeSnapshotController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f.mu.Lock()
	defer f.mu.Unlock()
	path := strings.TrimPrefix(r.URL.Path, "/api")
	line := r.Method + " " + path
	if r.URL.RawQuery != "" {
		line += "?" + r.URL.RawQuery
	}
	f.requests = append(f.requests, line)
	_, _ = io.Copy(io.Discard, r.Body)
	parts := strings.Split(strings.Trim(path, "/"), "/")

	switch {
	// GET /v1/devices/id/{id}/status
	case r.Method == http.MethodGet && len(parts) == 5 && parts[1] == "devices" && parts[4] == "status":
		writeJSON(w, http.StatusOK, map[string]interface{}{"id": parts[3], "rawStatus": f.rawStatus[parts[3]]})

	// GET /v1/apps/instances/snapshot/id/{id}[/state]
	case r.Method == http.MethodGet && len(parts) >= 6 && parts[3] == "snapshot" && parts[4] == "id":
		id := parts[5]
		s, ok := f.snaps[id]
		if !ok {
			notFound(w)
			return
		}
		if len(parts) == 7 && parts[6] == "state" {
			q := f.states[id]
			if len(q) == 0 {
				notFound(w)
				return
			}
			st := q[0]
			if len(q) > 1 {
				f.states[id] = q[1:]
			}
			writeJSON(w, http.StatusOK, st)
			return
		}
		writeJSON(w, http.StatusOK, s)

	// /v1/apps/instances/id/{app}/...
	case len(parts) >= 5 && parts[2] == "instances" && parts[3] == "id":
		appID := parts[4]
		inst, ok := f.instances[appID]
		if !ok {
			notFound(w)
			return
		}
		rest := parts[5:]
		switch {
		case r.Method == http.MethodGet && len(rest) == 0:
			writeJSON(w, http.StatusOK, inst)
		case r.Method == http.MethodPost && len(rest) == 1 && rest[0] == "snapshot":
			for id, s := range f.snaps {
				if s["config"].(map[string]interface{})["id"] == appID {
					delete(f.snaps, id)
				}
			}
			name := fmt.Sprintf("snapshot-2026102112000%d-abcdef12", f.next+1)
			id := f.addSnapshot(appID, name, f.createStates...)
			if f.noObjectID {
				writeJSON(w, http.StatusOK, map[string]interface{}{"httpStatusCode": 200})
				return
			}
			writeJSON(w, http.StatusOK, map[string]interface{}{"httpStatusCode": 200, "objectId": id, "objectName": name})
		case r.Method == http.MethodGet && len(rest) == 1 && rest[0] == "snapshots":
			var l []interface{}
			for _, s := range f.snaps {
				if s["config"].(map[string]interface{})["id"] == appID {
					l = append(l, s)
				}
			}
			writeJSON(w, http.StatusOK, map[string]interface{}{"snapshots": l})
		case len(rest) >= 3 && rest[0] == "snapshot" && rest[1] == "name":
			id, s := f.snapshotByName(appID, rest[2])
			if s == nil {
				notFound(w)
				return
			}
			switch {
			case r.Method == http.MethodGet && len(rest) == 3:
				writeJSON(w, http.StatusOK, s)
			case r.Method == http.MethodDelete && len(rest) == 3:
				delete(f.snaps, id)
				writeJSON(w, http.StatusOK, map[string]interface{}{"httpStatusCode": 200})
			case r.Method == http.MethodPut && len(rest) == 4 && rest[3] == "activate":
				f.states[id] = append([]snapshots.State(nil), f.activateStates...)
				writeJSON(w, http.StatusOK, map[string]interface{}{"httpStatusCode": 200, "objectId": id})
			default:
				writeJSON(w, http.StatusNotImplemented, map[string]interface{}{"httpStatusCode": 501})
			}
		default:
			writeJSON(w, http.StatusNotImplemented, map[string]interface{}{"httpStatusCode": 501})
		}

	default:
		writeJSON(w, http.StatusNotImplemented, map[string]interface{}{"httpStatusCode": 501})
	}
}

func newSnapshotFakeMeta(t *testing.T, h http.Handler) *api_client.ZedcloudAPI {
	t.Helper()
	old := snapshotPollInterval
	snapshotPollInterval = time.Millisecond
	t.Cleanup(func() { snapshotPollInterval = old })

	t.Setenv("TF_INSECURE_SKIP_TLS_VERIFY", "true")
	srv := httptest.NewTLSServer(h)
	t.Cleanup(srv.Close)
	opts := httpOptions{requestTimeout: 5 * time.Second, maxRetries: 0, retryWaitMax: 10 * time.Millisecond}
	std := getRetryClient(opts).StandardClient()
	std.Transport = NewHttpTransportWrapper(std.Transport)
	transport := httptransport.NewWithClient(strings.TrimPrefix(srv.URL, "https://"), "/api", []string{"https"}, std)
	c := api_client.New(transport, strfmt.Default)
	c.Settings = &api_client.Settings{AsyncWrites: true, JobPollIntervalMin: time.Millisecond, JobPollIntervalMax: 5 * time.Millisecond}
	return c
}

func newSnapshotData(t *testing.T, raw map[string]interface{}) *schema.ResourceData {
	t.Helper()
	return schema.TestResourceDataRaw(t, zschema.AppInstanceSnapshotSchema(), raw)
}

func newActivationData(t *testing.T, raw map[string]interface{}) *schema.ResourceData {
	t.Helper()
	return schema.TestResourceDataRaw(t, zschema.AppInstanceSnapshotActivationSchema(), raw)
}

func diagText(ds diag.Diagnostics) string {
	var b strings.Builder
	for _, d := range ds {
		b.WriteString(d.Summary + " " + d.Detail + "\n")
	}
	return b.String()
}

// --- zedcloud_app_instance_snapshot ---

func TestAppInstanceSnapshot_CreateReadDelete(t *testing.T) {
	fc := newFakeSnapshotController()
	meta := newSnapshotFakeMeta(t, fc)
	ctx := context.Background()
	d := newSnapshotData(t, map[string]interface{}{"app_instance_id": "app-1"})

	if ds := CreateAppInstanceSnapshot(ctx, d, meta); ds.HasError() {
		t.Fatalf("create: %s\n%s", diagText(ds), fc.log())
	}
	if d.Id() != "snap-1" || !strings.HasPrefix(d.Get("name").(string), "snapshot-") {
		t.Fatalf("id %q, name %q", d.Id(), d.Get("name"))
	}
	if got := d.Get("state").(string); got != snapshots.StateCreated {
		t.Fatalf("state = %q, want CREATED", got)
	}
	if !strings.Contains(d.Get("config_json").(string), `"name":"inst-1"`) {
		t.Fatalf("config_json = %s", d.Get("config_json"))
	}
	if fc.count("GET /v1/devices/id/dev-1/status") != 1 {
		t.Fatalf("expected one capability check:\n%s", fc.log())
	}

	// A newer snapshot of the same app instance replaces this one.
	fc.mu.Lock()
	delete(fc.snaps, "snap-1")
	fc.mu.Unlock()
	if ds := ReadAppInstanceSnapshot(ctx, d, meta); ds.HasError() || d.Id() != "" {
		t.Fatalf("read of a replaced snapshot must remove it: id %q, %s", d.Id(), diagText(ds))
	}

	d2 := newSnapshotData(t, map[string]interface{}{"app_instance_id": "app-1"})
	if ds := CreateAppInstanceSnapshot(ctx, d2, meta); ds.HasError() {
		t.Fatalf("create: %s", diagText(ds))
	}
	if ds := DeleteAppInstanceSnapshot(ctx, d2, meta); ds.HasError() || d2.Id() != "" {
		t.Fatalf("delete: id %q, %s", d2.Id(), diagText(ds))
	}
	if !strings.Contains(fc.log(), "DELETE /v1/apps/instances/id/app-1/snapshot/name/") || len(fc.snaps) != 0 {
		t.Fatalf("delete not sent:\n%s", fc.log())
	}
	// Deleting again (already gone) is success.
	d2.SetId("snap-x")
	if ds := DeleteAppInstanceSnapshot(ctx, d2, meta); ds.HasError() {
		t.Fatalf("delete of a missing snapshot: %s", diagText(ds))
	}
}

func TestAppInstanceSnapshot_CapabilityMissing(t *testing.T) {
	for name, raw := range map[string]string{
		"level too low": `{"apiCapability":"API_CAPABILITY_WIN_LIC_PASSTHROUGH"}`,
		"no level":      `{"swList":[]}`,
		"no status":     "",
	} {
		t.Run(name, func(t *testing.T) {
			fc := newFakeSnapshotController()
			fc.rawStatus["dev-1"] = raw
			meta := newSnapshotFakeMeta(t, fc)
			d := newSnapshotData(t, map[string]interface{}{"app_instance_id": "app-1"})

			ds := CreateAppInstanceSnapshot(context.Background(), d, meta)
			if !ds.HasError() || !strings.Contains(diagText(ds), "skip_capability_check") {
				t.Fatalf("diags = %s", diagText(ds))
			}
			if fc.count("POST") != 0 || d.Id() != "" {
				t.Fatalf("no snapshot may be requested:\n%s", fc.log())
			}
		})
	}
}

func TestAppInstanceSnapshot_SkipCapabilityCheck(t *testing.T) {
	fc := newFakeSnapshotController()
	fc.rawStatus["dev-1"] = ""
	meta := newSnapshotFakeMeta(t, fc)
	d := newSnapshotData(t, map[string]interface{}{"app_instance_id": "app-1", "skip_capability_check": true})

	if ds := CreateAppInstanceSnapshot(context.Background(), d, meta); ds.HasError() {
		t.Fatalf("create: %s", diagText(ds))
	}
	if fc.count("GET /v1/devices/") != 0 || fc.count("POST") != 1 {
		t.Fatalf("expected no status read and one create:\n%s", fc.log())
	}
}

func TestAppInstanceSnapshot_RejectsClusterAppInstance(t *testing.T) {
	fc := newFakeSnapshotController()
	meta := newSnapshotFakeMeta(t, fc)
	d := newSnapshotData(t, map[string]interface{}{"app_instance_id": "app-cluster", "skip_capability_check": true})

	ds := CreateAppInstanceSnapshot(context.Background(), d, meta)
	if !ds.HasError() || !strings.Contains(diagText(ds), "edge node cluster cluster-1") {
		t.Fatalf("diags = %s", diagText(ds))
	}
	if fc.count("POST") != 0 {
		t.Fatalf("no snapshot may be requested:\n%s", fc.log())
	}
}

func TestAppInstanceSnapshot_CreateFailedLeavesID(t *testing.T) {
	fc := newFakeSnapshotController()
	fc.createStates = []snapshots.State{{State: snapshots.StateCreating}, {State: snapshots.StateCreateFailed}}
	meta := newSnapshotFakeMeta(t, fc)
	d := newSnapshotData(t, map[string]interface{}{"app_instance_id": "app-1"})

	ds := CreateAppInstanceSnapshot(context.Background(), d, meta)
	if !ds.HasError() || !strings.Contains(diagText(ds), "CREATE_FAILED") {
		t.Fatalf("diags = %s", diagText(ds))
	}
	if d.Id() == "" {
		t.Fatal("the id must stay set so Terraform taints the resource")
	}
}

func TestAppInstanceSnapshot_NoWait(t *testing.T) {
	fc := newFakeSnapshotController()
	fc.createStates = []snapshots.State{{State: snapshots.StateCreating}}
	meta := newSnapshotFakeMeta(t, fc)
	d := newSnapshotData(t, map[string]interface{}{"app_instance_id": "app-1", "wait_for_ready": false})

	if ds := CreateAppInstanceSnapshot(context.Background(), d, meta); ds.HasError() {
		t.Fatalf("create: %s", diagText(ds))
	}
	if d.Get("state").(string) != snapshots.StateCreating {
		t.Fatalf("state = %q", d.Get("state"))
	}
}

func TestAppInstanceSnapshot_OldControllerListsForID(t *testing.T) {
	fc := newFakeSnapshotController()
	fc.noObjectID = true
	meta := newSnapshotFakeMeta(t, fc)
	d := newSnapshotData(t, map[string]interface{}{"app_instance_id": "app-1"})

	if ds := CreateAppInstanceSnapshot(context.Background(), d, meta); ds.HasError() {
		t.Fatalf("create: %s\n%s", diagText(ds), fc.log())
	}
	if d.Id() != "snap-1" || fc.count("GET /v1/apps/instances/id/app-1/snapshots") != 1 {
		t.Fatalf("id %q\n%s", d.Id(), fc.log())
	}
}

func TestAppInstanceSnapshot_Import(t *testing.T) {
	fc := newFakeSnapshotController()
	id := fc.addSnapshot("app-1", "imported", snapshots.State{State: snapshots.StateCreated})
	meta := newSnapshotFakeMeta(t, fc)
	d := newSnapshotData(t, map[string]interface{}{})
	d.SetId(id)

	if ds := ReadAppInstanceSnapshot(context.Background(), d, meta); ds.HasError() {
		t.Fatalf("read: %s", diagText(ds))
	}
	if d.Get("app_instance_id") != "app-1" || d.Get("name") != "imported" || d.Get("state") != snapshots.StateCreated ||
		d.Get("wait_for_ready") != true {
		t.Fatalf("state after import: app %v, name %v, state %v, wait %v",
			d.Get("app_instance_id"), d.Get("name"), d.Get("state"), d.Get("wait_for_ready"))
	}
}

func TestAppInstanceSnapshot_PlanRejectsClusterAppInstance(t *testing.T) {
	fc := newFakeSnapshotController()
	meta := newSnapshotFakeMeta(t, fc)
	for _, r := range []*schema.Resource{AppInstanceSnapshotResource(), AppInstanceSnapshotActivationResource()} {
		cfg := map[string]interface{}{"app_instance_id": "app-cluster", "snapshot_name": "s"}
		if _, ok := r.Schema["snapshot_name"]; !ok {
			delete(cfg, "snapshot_name")
		}
		_, err := r.Diff(context.Background(), nil, terraform.NewResourceConfigRaw(cfg), meta)
		if err == nil || !strings.Contains(err.Error(), "edge node cluster") {
			t.Fatalf("plan must fail for a cluster app instance, err = %v", err)
		}
		cfg["app_instance_id"] = "app-1"
		if _, err := r.Diff(context.Background(), nil, terraform.NewResourceConfigRaw(cfg), meta); err != nil {
			t.Fatalf("plan must pass for a device app instance: %v", err)
		}
	}
}

// --- zedcloud_app_instance_snapshot_activation ---

func TestAppInstanceSnapshotActivation_Create(t *testing.T) {
	fc := newFakeSnapshotController()
	fc.addSnapshot("app-1", "s1", snapshots.State{State: snapshots.StateCreated, Timestamp: "t2"})
	meta := newSnapshotFakeMeta(t, fc)
	d := newActivationData(t, map[string]interface{}{"app_instance_id": "app-1", "snapshot_name": "s1"})

	if ds := CreateAppInstanceSnapshotActivation(context.Background(), d, meta); ds.HasError() {
		t.Fatalf("create: %s\n%s", diagText(ds), fc.log())
	}
	if d.Id() != "app-1/s1" || d.Get("snapshot_id") != "snap-1" || d.Get("state") != snapshots.StateActive ||
		d.Get("activated_at") != "t4" {
		t.Fatalf("id %q, snapshot_id %v, state %v, activated_at %v",
			d.Id(), d.Get("snapshot_id"), d.Get("state"), d.Get("activated_at"))
	}
	if !strings.Contains(fc.log(), "PUT /v1/apps/instances/id/app-1/snapshot/name/s1/activate?force=true") {
		t.Fatalf("expected a forced activate:\n%s", fc.log())
	}
}

func TestAppInstanceSnapshotActivation_RefusesSnapshotRolledBackBefore(t *testing.T) {
	// The controller sends EVE RollbackCmd{Counter: 1} for every activation, so
	// EVE rolls back to a snapshot only once. Refuse before sending anything.
	for _, st := range []string{snapshots.StateActive, snapshots.StateRollingBack, snapshots.StateRollbackFailed} {
		fc := newFakeSnapshotController()
		fc.addSnapshot("app-1", "s1", snapshots.State{State: st, Timestamp: "t4"})
		meta := newSnapshotFakeMeta(t, fc)
		d := newActivationData(t, map[string]interface{}{"app_instance_id": "app-1", "snapshot_name": "s1"})

		ds := CreateAppInstanceSnapshotActivation(context.Background(), d, meta)
		if !ds.HasError() || !strings.Contains(diagText(ds), "Take a new snapshot") {
			t.Fatalf("%s: diags = %s", st, diagText(ds))
		}
		if fc.count("PUT") != 0 || d.Id() != "" {
			t.Fatalf("%s: nothing may be activated:\n%s", st, fc.log())
		}
	}
}

func TestAppInstanceSnapshotActivation_RollbackFailed(t *testing.T) {
	fc := newFakeSnapshotController()
	fc.addSnapshot("app-1", "s1", snapshots.State{State: snapshots.StateCreated, Timestamp: "t2"})
	fc.activateStates = []snapshots.State{
		{State: snapshots.StateRollingBack, Timestamp: "t3"},
		{State: snapshots.StateRollbackFailed, Timestamp: "t4"},
	}
	meta := newSnapshotFakeMeta(t, fc)
	d := newActivationData(t, map[string]interface{}{"app_instance_id": "app-1", "snapshot_name": "s1"})

	ds := CreateAppInstanceSnapshotActivation(context.Background(), d, meta)
	if !ds.HasError() || !strings.Contains(diagText(ds), "ROLLBACK_FAILED") {
		t.Fatalf("diags = %s", diagText(ds))
	}
	if d.Id() == "" {
		t.Fatal("the id must stay set so Terraform taints the resource and retries")
	}
}

func TestAppInstanceSnapshotActivation_ForcesAndIgnoresStaleState(t *testing.T) {
	// Always forced (a stale activate flag would make a plain activate a
	// no-op), and the state from before the request must not count.
	fc := newFakeSnapshotController()
	fc.addSnapshot("app-1", "s1", snapshots.State{State: snapshots.StateCreated, Timestamp: "t2"})
	fc.activateStates = []snapshots.State{
		{State: snapshots.StateCreated, Timestamp: "t2"}, // stale
		{State: snapshots.StateRollingBack, Timestamp: "t5"},
		{State: snapshots.StateActive, Timestamp: "t6"},
	}
	meta := newSnapshotFakeMeta(t, fc)
	d := newActivationData(t, map[string]interface{}{"app_instance_id": "app-1", "snapshot_name": "s1"})

	if ds := CreateAppInstanceSnapshotActivation(context.Background(), d, meta); ds.HasError() {
		t.Fatalf("create: %s", diagText(ds))
	}
	if !strings.Contains(fc.log(), "/activate?force=true") {
		t.Fatalf("a rollback must always force:\n%s", fc.log())
	}
	if d.Get("state") != snapshots.StateActive || d.Get("activated_at") != "t6" {
		t.Fatalf("state %v, activated_at %v", d.Get("state"), d.Get("activated_at"))
	}
}

func TestAppInstanceSnapshotActivation_RefusesSnapshotNotCreated(t *testing.T) {
	for _, st := range []string{snapshots.StateCreating, snapshots.StateCreateFailed, snapshots.StateDeleting, ""} {
		fc := newFakeSnapshotController()
		if st == "" {
			fc.addSnapshot("app-1", "s1") // ganges has no state yet
		} else {
			fc.addSnapshot("app-1", "s1", snapshots.State{State: st, Timestamp: "t1"})
		}
		meta := newSnapshotFakeMeta(t, fc)
		d := newActivationData(t, map[string]interface{}{"app_instance_id": "app-1", "snapshot_name": "s1"})

		ds := CreateAppInstanceSnapshotActivation(context.Background(), d, meta)
		if !ds.HasError() || !strings.Contains(diagText(ds), "not CREATED") {
			t.Fatalf("%q: diags = %s", st, diagText(ds))
		}
		if fc.count("PUT") != 0 {
			t.Fatalf("%q: nothing may be activated:\n%s", st, fc.log())
		}
	}
}

func TestAppInstanceSnapshotActivation_Timeout(t *testing.T) {
	fc := newFakeSnapshotController()
	fc.addSnapshot("app-1", "s1", snapshots.State{State: snapshots.StateCreated, Timestamp: "t2"})
	fc.activateStates = []snapshots.State{{State: snapshots.StateRollingBack, Timestamp: "t3"}}
	meta := newSnapshotFakeMeta(t, fc)
	d := newActivationData(t, map[string]interface{}{"app_instance_id": "app-1", "snapshot_name": "s1"})
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	ds := CreateAppInstanceSnapshotActivation(ctx, d, meta)
	txt := diagText(ds)
	if !ds.HasError() || !strings.Contains(txt, "not confirmed") || d.Get("state") != snapshots.StateRollingBack {
		t.Fatalf("diags = %s", txt)
	}
	if d.Id() == "" {
		t.Fatal("the id must stay set: the controller config has already changed")
	}
}

func TestAppInstanceSnapshotActivation_MissingSnapshot(t *testing.T) {
	fc := newFakeSnapshotController()
	meta := newSnapshotFakeMeta(t, fc)
	d := newActivationData(t, map[string]interface{}{"app_instance_id": "app-1", "snapshot_name": "nope"})

	ds := CreateAppInstanceSnapshotActivation(context.Background(), d, meta)
	if !ds.HasError() || !strings.Contains(diagText(ds), "no snapshot named") {
		t.Fatalf("diags = %s", diagText(ds))
	}
	if fc.count("PUT") != 0 || d.Id() != "" {
		t.Fatalf("nothing may be activated:\n%s", fc.log())
	}
}

func TestAppInstanceSnapshotActivation_RejectsClusterAppInstance(t *testing.T) {
	fc := newFakeSnapshotController()
	meta := newSnapshotFakeMeta(t, fc)
	d := newActivationData(t, map[string]interface{}{"app_instance_id": "app-cluster", "snapshot_name": "s1"})

	ds := CreateAppInstanceSnapshotActivation(context.Background(), d, meta)
	if !ds.HasError() || !strings.Contains(diagText(ds), "edge node cluster") || fc.count("PUT") != 0 {
		t.Fatalf("diags = %s\n%s", diagText(ds), fc.log())
	}
}

func TestAppInstanceSnapshotActivation_ReadKeepsResourceWhenSnapshotGone(t *testing.T) {
	fc := newFakeSnapshotController()
	meta := newSnapshotFakeMeta(t, fc)
	d := newActivationData(t, map[string]interface{}{"app_instance_id": "app-1", "snapshot_name": "s1"})
	d.SetId("app-1/s1")
	_ = d.Set("snapshot_id", "snap-gone")
	_ = d.Set("state", snapshots.StateActive)

	if ds := ReadAppInstanceSnapshotActivation(context.Background(), d, meta); ds.HasError() {
		t.Fatalf("read: %s", diagText(ds))
	}
	if d.Id() != "app-1/s1" || d.Get("state") != "" {
		t.Fatalf("id %q, state %v", d.Id(), d.Get("state"))
	}
}

func TestAppInstanceSnapshotActivation_DeleteWarns(t *testing.T) {
	d := newActivationData(t, map[string]interface{}{"app_instance_id": "app-1", "snapshot_name": "s1"})
	d.SetId("app-1/s1")
	ds := DeleteAppInstanceSnapshotActivation(context.Background(), d, nil)
	if ds.HasError() || len(ds) != 1 || ds[0].Severity != diag.Warning || d.Id() != "" {
		t.Fatalf("diags %v, id %q", ds, d.Id())
	}
}

// --- data.zedcloud_app_instance_snapshots ---

func TestAppInstanceSnapshotsDataSource(t *testing.T) {
	fc := newFakeSnapshotController()
	fc.addSnapshot("app-1", "s1", snapshots.State{State: snapshots.StateCreated})
	fc.addSnapshot("app-1", "s2") // no state yet
	meta := newSnapshotFakeMeta(t, fc)

	d := schema.TestResourceDataRaw(t, zschema.AppInstanceSnapshotsDataSourceSchema(), map[string]interface{}{"app_instance_id": "app-1"})
	if ds := ReadAppInstanceSnapshots(context.Background(), d, meta); ds.HasError() {
		t.Fatalf("read: %s", diagText(ds))
	}
	got := map[string]string{}
	for _, s := range d.Get("snapshots").([]interface{}) {
		m := s.(map[string]interface{})
		got[m["name"].(string)] = m["state"].(string)
	}
	if len(got) != 2 || got["s1"] != snapshots.StateCreated || got["s2"] != "" || d.Id() != "app-1" {
		t.Fatalf("snapshots = %v, id %q", got, d.Id())
	}

	d = schema.TestResourceDataRaw(t, zschema.AppInstanceSnapshotsDataSourceSchema(),
		map[string]interface{}{"app_instance_id": "app-1", "name": "s2"})
	if ds := ReadAppInstanceSnapshots(context.Background(), d, meta); ds.HasError() {
		t.Fatalf("read: %s", diagText(ds))
	}
	if l := d.Get("snapshots").([]interface{}); len(l) != 1 || l[0].(map[string]interface{})["name"] != "s2" {
		t.Fatalf("filtered snapshots = %v", l)
	}
}

func TestSnapshotFake_RoundTrip(t *testing.T) {
	// Guards the fake itself: state queues advance and repeat the last entry.
	fc := newFakeSnapshotController()
	id := fc.addSnapshot("app-1", "s", snapshots.State{State: "A"}, snapshots.State{State: "B"})
	meta := newSnapshotFakeMeta(t, fc)
	var seen []string
	for i := 0; i < 3; i++ {
		st, err := meta.AppInstanceSnapshot.State(context.Background(), id)
		if err != nil {
			t.Fatal(err)
		}
		seen = append(seen, st.State)
	}
	b, _ := json.Marshal(seen)
	if string(b) != `["A","B","B"]` {
		t.Fatalf("states = %s", b)
	}
}

// stateSeq is a ClientService whose State answers from a list of results.
type stateSeq struct {
	snapshots.ClientService
	results []func() (*snapshots.State, error)
	calls   int
}

func (s *stateSeq) State(context.Context, string) (*snapshots.State, error) {
	f := s.results[len(s.results)-1]
	if s.calls < len(s.results) {
		f = s.results[s.calls]
	}
	s.calls++
	return f()
}

func TestWaitForSnapshotState_RidesOutTransientErrors(t *testing.T) {
	old := snapshotPollInterval
	snapshotPollInterval = time.Millisecond
	t.Cleanup(func() { snapshotPollInterval = old })

	c := &stateSeq{results: []func() (*snapshots.State, error){
		func() (*snapshots.State, error) { return nil, snapshots.ErrNotFound },
		func() (*snapshots.State, error) { return nil, &snapshots.APIError{Op: "state", Status: 502} },
		func() (*snapshots.State, error) { return &snapshots.State{State: snapshots.StateCreated}, nil },
	}}
	st, err := waitForSnapshotState(context.Background(), c, "s", time.Second, func(st *snapshots.State) (bool, error) {
		return st.Name() == snapshots.StateCreated, nil
	})
	if err != nil || st.Name() != snapshots.StateCreated {
		t.Fatalf("state %v, err %v", st, err)
	}
}

func TestWaitForSnapshotState_StopsOnClientError(t *testing.T) {
	c := &stateSeq{results: []func() (*snapshots.State, error){
		func() (*snapshots.State, error) { return nil, &snapshots.APIError{Op: "state", Status: 403} },
	}}
	_, err := waitForSnapshotState(context.Background(), c, "s", time.Minute, func(*snapshots.State) (bool, error) {
		return false, nil
	})
	if err == nil || c.calls != 1 {
		t.Fatalf("a 403 must end the wait at once: err %v, calls %d", err, c.calls)
	}
}

func TestWaitForSnapshotState_Timeout(t *testing.T) {
	old := snapshotPollInterval
	snapshotPollInterval = time.Millisecond
	t.Cleanup(func() { snapshotPollInterval = old })

	c := &stateSeq{results: []func() (*snapshots.State, error){
		func() (*snapshots.State, error) { return &snapshots.State{State: snapshots.StateCreating}, nil },
	}}
	_, err := waitForSnapshotState(context.Background(), c, "s", 20*time.Millisecond, func(*snapshots.State) (bool, error) {
		return false, nil
	})
	if err == nil || !strings.Contains(err.Error(), "timed out") || !strings.Contains(err.Error(), "CREATING") ||
		!strings.Contains(err.Error(), "online") {
		t.Fatalf("err = %v", err)
	}
}
