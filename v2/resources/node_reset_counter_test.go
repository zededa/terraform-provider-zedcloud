package resources

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	zschema "github.com/zededa/terraform-provider-zedcloud/v2/schemas"
)

// ENG-3115: the controller owns reset_counter. A counter in an update body is
// ignored; only PUT /v1/devices/id/{id}/reboot reboots the node and adds 1.
// The provider used to send the counter in the body, so a raised reset_counter
// applied "successfully", never rebooted and kept planning the same diff.

const resetNodeID = "7da54ab0-dcf2-403d-8f0f-115c2695ed23"

type fakeNodeController struct {
	mu          sync.Mutex
	node        map[string]interface{}
	reboots     int
	failReboot  bool
	updateBodys []map[string]interface{}
}

func newFakeNodeController(counter int) *fakeNodeController {
	return &fakeNodeController{node: map[string]interface{}{
		"id": resetNodeID, "name": "eng3115", "title": "eng3115",
		"modelId": testModelA, "projectId": "proj-1", "serialno": "SN-1",
		"adminState": "ADMIN_STATE_ACTIVE", "resetCounter": counter,
		"revision": map[string]interface{}{
			"curr": "1", "createdAt": "2026-11-25T00:00:00Z", "createdBy": "tf",
			"updatedAt": "2026-11-25T00:00:00Z", "updatedBy": "tf",
		},
	}}
}

func (f *fakeNodeController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f.mu.Lock()
	defer f.mu.Unlock()
	w.Header().Set("Content-Type", "application/json")
	path := strings.TrimPrefix(r.URL.Path, "/api")
	switch {
	case r.Method == http.MethodPut && path == "/v1/devices/id/"+resetNodeID+"/reboot":
		if f.failReboot {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(`{"error":[{"details":"reboot failed"}]}`))
			return
		}
		f.reboots++
		f.node["resetCounter"] = f.node["resetCounter"].(int) + 1
		_, _ = w.Write([]byte(`{}`))
	case r.Method == http.MethodPut && path == "/v1/devices/id/"+resetNodeID:
		var body map[string]interface{}
		_ = json.NewDecoder(r.Body).Decode(&body)
		f.updateBodys = append(f.updateBodys, body)
		for k, v := range body {
			if k != "resetCounter" && k != "id" {
				f.node[k] = v
			}
		}
		_, _ = w.Write([]byte(`{}`))
	case r.Method == http.MethodGet && (path == "/v1/devices/id/"+resetNodeID || path == "/v1/devices/name/eng3115"):
		_ = json.NewEncoder(w).Encode(f.node)
	default:
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{}`))
	}
}

func (f *fakeNodeController) counter() int {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.node["resetCounter"].(int)
}

func newNodeMeta(t *testing.T, h http.Handler) *api_client.ZedcloudAPI {
	t.Helper()
	t.Setenv("TF_INSECURE_SKIP_TLS_VERIFY", "true")
	srv := httptest.NewTLSServer(h)
	t.Cleanup(srv.Close)
	opts := httpOptions{requestTimeout: 5 * time.Second, maxRetries: 0, retryWaitMax: 10 * time.Millisecond}
	std := getRetryClient(opts).StandardClient()
	std.Transport = NewHttpTransportWrapper(std.Transport)
	transport := httptransport.NewWithClient(strings.TrimPrefix(srv.URL, "https://"), "/api", []string{"https"}, std)
	return api_client.New(transport, strfmt.Default)
}

// resetUpdateData is the ResourceData for an update of reset_counter from
// state oldCounter to the configured newCounter.
func resetUpdateData(t *testing.T, oldCounter, newCounter int) *schema.ResourceData {
	t.Helper()
	r := &schema.Resource{Schema: zschema.Node()}
	state := &terraform.InstanceState{ID: resetNodeID, Attributes: map[string]string{
		"id": resetNodeID, "name": "eng3115", "title": "eng3115", "model_id": testModelA,
		"project_id": "proj-1", "serialno": "SN-1", "admin_state": "ADMIN_STATE_ACTIVE",
		"reset_counter": itoa(oldCounter),
	}}
	cfg := terraform.NewResourceConfigRaw(map[string]interface{}{
		"name": "eng3115", "title": "eng3115", "model_id": testModelA, "project_id": "proj-1",
		"serialno": "SN-1", "admin_state": "ADMIN_STATE_ACTIVE", "reset_counter": newCounter,
	})
	diff, err := r.Diff(context.Background(), state, cfg, nil)
	if err != nil {
		t.Fatal(err)
	}
	d, err := schema.InternalMap(r.Schema).Data(state, diff)
	if err != nil {
		t.Fatal(err)
	}
	return d
}

func itoa(i int) string { b, _ := json.Marshal(i); return string(b) }

func TestNodeResetCounter_IncreaseRebootsOnce(t *testing.T) {
	fc := newFakeNodeController(1000)
	m := newNodeMeta(t, fc)
	// A jump of 5 still reboots once; the controller adds 1.
	d := resetUpdateData(t, 1000, 1005)
	if diags := UpdateNode(context.Background(), d, m); diags.HasError() {
		t.Fatalf("update failed: %v", diags)
	}
	if fc.reboots != 1 {
		t.Fatalf("reboots = %d, want 1", fc.reboots)
	}
	if got := d.Get("reset_counter").(int); got != 1005 {
		t.Errorf("state reset_counter = %d, want the configured 1005", got)
	}
	// A refresh must not bring the diff back.
	if diags := readNode(context.Background(), d, m); diags.HasError() {
		t.Fatal(diags)
	}
	if got := d.Get("reset_counter").(int); got != 1005 {
		t.Errorf("after refresh reset_counter = %d, want 1005", got)
	}
}

func TestNodeResetCounter_NoRebootWithoutIncrease(t *testing.T) {
	for _, c := range []struct {
		name     string
		old, new int
	}{
		{"unchanged", 1000, 1000},
		{"decreased", 1000, 999},
	} {
		t.Run(c.name, func(t *testing.T) {
			fc := newFakeNodeController(1000)
			m := newNodeMeta(t, fc)
			if diags := UpdateNode(context.Background(), resetUpdateData(t, c.old, c.new), m); diags.HasError() {
				t.Fatal(diags)
			}
			if fc.reboots != 0 {
				t.Errorf("reboots = %d, want 0", fc.reboots)
			}
		})
	}
}

func TestNodeResetCounter_FailedRebootIsRetried(t *testing.T) {
	fc := newFakeNodeController(1000)
	fc.failReboot = true
	m := newNodeMeta(t, fc)
	d := resetUpdateData(t, 1000, 1001)
	if diags := UpdateNode(context.Background(), d, m); !diags.HasError() {
		t.Fatal("a failed reboot must fail the apply")
	}
	if got := d.Get("reset_counter").(int); got != 1000 {
		t.Fatalf("state reset_counter = %d, want the old 1000 so the next apply retries", got)
	}
	fc.failReboot = false
	if diags := UpdateNode(context.Background(), resetUpdateData(t, 1000, 1001), m); diags.HasError() {
		t.Fatal(diags)
	}
	if fc.reboots != 1 || fc.counter() != 1001 {
		t.Errorf("after retry: reboots=%d counter=%d, want 1 and 1001", fc.reboots, fc.counter())
	}
}

func TestNodeResetCounter_ControllerAheadOfStateIsTaken(t *testing.T) {
	// Someone rebooted the node out of band: refresh shows the controller's value.
	fc := newFakeNodeController(1003)
	m := newNodeMeta(t, fc)
	d := resetUpdateData(t, 1000, 1000)
	if diags := readNode(context.Background(), d, m); diags.HasError() {
		t.Fatal(diags)
	}
	if got := d.Get("reset_counter").(int); got != 1003 {
		t.Errorf("reset_counter = %d, want 1003", got)
	}
}
