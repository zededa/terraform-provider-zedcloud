package resources

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"sync"
	"testing"
	"time"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	"github.com/zededa/terraform-provider-zedcloud/v2/client/jobs"
	zschema "github.com/zededa/terraform-provider-zedcloud/v2/schemas"
)

// fakeController is a minimal in-memory zedcloud for the app-instance routes
// the provider uses (UE-177). No controller or credentials are needed.
type fakeController struct {
	mu        sync.Mutex
	instances map[string]map[string]interface{} // by name
	jobs      map[string][]jobs.Status          // job id -> remaining statuses
	jobErr    map[string]string
	requests  []string
	nextJob   int

	// failCreate makes the next v2 create job fail without creating anything.
	failCreate string
	// gatewayHTML makes every request answer 504 with an HTML page.
	gatewayHTML bool
}

func newFakeController() *fakeController {
	return &fakeController{
		instances: map[string]map[string]interface{}{},
		jobs:      map[string][]jobs.Status{},
		jobErr:    map[string]string{},
	}
}

func (f *fakeController) newJob(statuses ...jobs.Status) string {
	f.nextJob++
	id := fmt.Sprintf("00000000-0000-0000-0000-%012d", f.nextJob)
	f.jobs[id] = statuses
	return id
}

func (f *fakeController) byID(id string) map[string]interface{} {
	for _, inst := range f.instances {
		if inst["id"] == id {
			return inst
		}
	}
	return nil
}

func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func notFound(w http.ResponseWriter) {
	writeJSON(w, http.StatusNotFound, map[string]interface{}{
		"httpStatusCode": 404, "httpStatusMsg": "Not Found",
		"error": []map[string]string{{"details": "record not found"}},
	})
}

func (f *fakeController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f.mu.Lock()
	defer f.mu.Unlock()
	path := strings.TrimPrefix(r.URL.Path, "/api")
	f.requests = append(f.requests, r.Method+" "+path)

	if f.gatewayHTML {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusGatewayTimeout)
		_, _ = io.WriteString(w, "<html><body><h1>504 Gateway Time-out</h1></body></html>")
		return
	}

	var body map[string]interface{}
	if r.Body != nil {
		raw, _ := io.ReadAll(r.Body)
		if len(raw) > 0 {
			_ = json.Unmarshal(raw, &body)
		}
	}

	switch {
	case r.Method == http.MethodPost && path == "/v2/apps/instances":
		name, _ := body["name"].(string)
		if pid, _ := body["projectId"].(string); pid == "" {
			// seine appInstCreateV2 rejects this up front; v1 derives it.
			writeJSON(w, http.StatusBadRequest, map[string]interface{}{
				"httpStatusCode": 400, "error": []map[string]string{{"details": "project id is empty"}},
			})
			return
		}
		if f.failCreate == name {
			job := f.newJob(jobs.StatusInProgress, jobs.StatusFailed)
			f.jobErr[job] = "app instance " + name + " already exists"
			writeJSON(w, http.StatusOK, map[string]interface{}{"objectId": job})
			return
		}
		f.instances[name] = body
		writeJSON(w, http.StatusOK, map[string]interface{}{"objectId": f.newJob(jobs.StatusInProgress, jobs.StatusCompleted)})

	case r.Method == http.MethodPost && path == "/v1/apps/instances":
		name, _ := body["name"].(string)
		f.instances[name] = body
		writeJSON(w, http.StatusOK, map[string]interface{}{"objectId": body["id"]})

	case r.Method == http.MethodGet && strings.HasPrefix(path, "/v1/devices/id/"):
		writeJSON(w, http.StatusOK, map[string]interface{}{
			"id": strings.TrimPrefix(path, "/v1/devices/id/"), "name": "dev", "title": "dev",
			"modelId": "m", "projectId": "proj-of-device",
		})

	case strings.HasPrefix(path, "/v1/jobs/id/"):
		parts := strings.Split(strings.TrimPrefix(path, "/v1/jobs/id/"), "/")
		statuses, ok := f.jobs[parts[0]]
		if !ok || len(parts) != 3 || parts[2] != jobs.ObjectTypeEdgeAppInstance {
			writeJSON(w, http.StatusInternalServerError, map[string]interface{}{"httpStatusCode": 500})
			return
		}
		st := statuses[0]
		if len(statuses) > 1 {
			f.jobs[parts[0]] = statuses[1:]
		}
		writeJSON(w, http.StatusOK, map[string]interface{}{
			"id": parts[0], "status": st, "errorDetailOfJob": f.jobErr[parts[0]],
		})

	case r.Method == http.MethodGet && strings.HasPrefix(path, "/v1/apps/instances/name/"):
		inst, ok := f.instances[strings.TrimPrefix(path, "/v1/apps/instances/name/")]
		if !ok {
			notFound(w)
			return
		}
		writeJSON(w, http.StatusOK, inst)

	case r.Method == http.MethodGet && strings.HasPrefix(path, "/v1/apps/instances/id/"):
		inst := f.byID(strings.TrimPrefix(path, "/v1/apps/instances/id/"))
		if inst == nil {
			notFound(w)
			return
		}
		writeJSON(w, http.StatusOK, inst)

	case r.Method == http.MethodPut && strings.HasPrefix(path, "/v2/apps/instances/id/"):
		id := strings.TrimPrefix(path, "/v2/apps/instances/id/")
		if body["id"] != id {
			writeJSON(w, http.StatusBadRequest, map[string]interface{}{"httpStatusCode": 400})
			return
		}
		inst := f.byID(id)
		if inst == nil {
			notFound(w)
			return
		}
		for k, v := range body {
			inst[k] = v
		}
		writeJSON(w, http.StatusOK, map[string]interface{}{"objectId": f.newJob(jobs.StatusCompleted)})

	case r.Method == http.MethodDelete && strings.HasPrefix(path, "/v2/apps/instances/id/"):
		inst := f.byID(strings.TrimPrefix(path, "/v2/apps/instances/id/"))
		if inst == nil {
			notFound(w)
			return
		}
		delete(f.instances, inst["name"].(string))
		writeJSON(w, http.StatusOK, map[string]interface{}{"objectId": f.newJob(jobs.StatusInProgress, jobs.StatusCompleted)})

	case r.Method == http.MethodDelete && strings.HasPrefix(path, "/v1/apps/instances/id/"):
		inst := f.byID(strings.TrimPrefix(path, "/v1/apps/instances/id/"))
		if inst == nil {
			notFound(w)
			return
		}
		delete(f.instances, inst["name"].(string))
		writeJSON(w, http.StatusOK, map[string]interface{}{"objectId": inst["id"]})

	default:
		writeJSON(w, http.StatusNotImplemented, map[string]interface{}{"httpStatusCode": 501})
	}
}

func (f *fakeController) requestLog() string {
	f.mu.Lock()
	defer f.mu.Unlock()
	return strings.Join(f.requests, "\n")
}

// newFakeMeta wires a ZedcloudAPI to fc the way ProviderConfigure does: the
// retrying client with checkRetry, the transport wrapper that rewrites
// non-JSON errors, https and the /api base path.
func newFakeMeta(t *testing.T, fc *fakeController, async bool) *api_client.ZedcloudAPI {
	t.Helper()
	t.Setenv("TF_INSECURE_SKIP_TLS_VERIFY", "true")
	srv := httptest.NewTLSServer(fc)
	t.Cleanup(srv.Close)

	opts := httpOptions{requestTimeout: 5 * time.Second, maxRetries: 0, retryWaitMax: 10 * time.Millisecond}
	std := getRetryClient(opts).StandardClient()
	std.Transport = NewHttpTransportWrapper(std.Transport)
	transport := httptransport.NewWithClient(strings.TrimPrefix(srv.URL, "https://"), "/api", []string{"https"}, std)

	c := api_client.New(transport, strfmt.Default)
	c.Settings = &api_client.Settings{
		AsyncWrites:        async,
		JobPollIntervalMin: time.Millisecond,
		JobPollIntervalMax: 5 * time.Millisecond,
	}
	return c
}

func newInstanceData(t *testing.T, name string) *schema.ResourceData {
	t.Helper()
	return schema.TestResourceDataRaw(t, zschema.ApplicationInstance(), map[string]interface{}{
		"name":       name,
		"title":      "title",
		"app_id":     "app-1",
		"device_id":  "dev-1",
		"project_id": "proj-1",
	})
}

func TestApplicationInstance_AsyncCreateUpdateDelete(t *testing.T) {
	fc := newFakeController()
	meta := newFakeMeta(t, fc, true)
	ctx := context.Background()
	d := newInstanceData(t, "inst-a")

	if diags := CreateApplicationInstance(ctx, d, meta); diags.HasError() {
		t.Fatalf("create: %v\n%s", diags, fc.requestLog())
	}
	sent := fc.instances["inst-a"]
	if sent == nil || sent["id"] == "" || sent["id"] != d.Id() {
		t.Fatalf("the provider must send the id it picked: sent %v, state id %q", sent["id"], d.Id())
	}
	if !strings.Contains(fc.requestLog(), "POST /v2/apps/instances") {
		t.Fatalf("expected the async route:\n%s", fc.requestLog())
	}

	if err := d.Set("title", "new title"); err != nil {
		t.Fatal(err)
	}
	if diags := UpdateApplicationInstance(ctx, d, meta); diags.HasError() {
		t.Fatalf("update: %v\n%s", diags, fc.requestLog())
	}
	if fc.instances["inst-a"]["title"] != "new title" {
		t.Fatalf("update not applied: %v", fc.instances["inst-a"]["title"])
	}

	if diags := DeleteApplicationInstance(ctx, d, meta); diags.HasError() {
		t.Fatalf("delete: %v\n%s", diags, fc.requestLog())
	}
	if d.Id() != "" || len(fc.instances) != 0 {
		t.Fatalf("delete: id %q, instances %v", d.Id(), fc.instances)
	}
}

func TestApplicationInstance_AsyncCreateResolvesProjectFromDevice(t *testing.T) {
	fc := newFakeController()
	meta := newFakeMeta(t, fc, true)
	d := schema.TestResourceDataRaw(t, zschema.ApplicationInstance(), map[string]interface{}{
		"name": "no-project", "title": "t", "app_id": "app-1", "device_id": "dev-1",
	})

	if diags := CreateApplicationInstance(context.Background(), d, meta); diags.HasError() {
		t.Fatalf("create: %v\n%s", diags, fc.requestLog())
	}
	if got := fc.instances["no-project"]["projectId"]; got != "proj-of-device" {
		t.Fatalf("projectId sent = %v, want the device's project", got)
	}
}

func TestApplicationInstance_AsyncCreateJobFailed(t *testing.T) {
	fc := newFakeController()
	fc.failCreate = "inst-dup"
	meta := newFakeMeta(t, fc, true)
	d := newInstanceData(t, "inst-dup")

	diags := CreateApplicationInstance(context.Background(), d, meta)
	if !diags.HasError() || !strings.Contains(fmt.Sprint(diags), "already exists") {
		t.Fatalf("diags = %v, want the job's error", diags)
	}
	if d.Id() != "" {
		t.Fatalf("nothing was created, so the id must be cleared, got %q", d.Id())
	}
}

func TestApplicationInstance_SyncCreate(t *testing.T) {
	fc := newFakeController()
	meta := newFakeMeta(t, fc, false)
	d := newInstanceData(t, "inst-sync")

	if diags := CreateApplicationInstance(context.Background(), d, meta); diags.HasError() {
		t.Fatalf("create: %v\n%s", diags, fc.requestLog())
	}
	log := fc.requestLog()
	if !strings.Contains(log, "POST /v1/apps/instances") || strings.Contains(log, "/v2/") || strings.Contains(log, "/jobs/") {
		t.Fatalf("expected only the sync route:\n%s", log)
	}
	if d.Id() == "" || fc.instances["inst-sync"]["id"] != d.Id() {
		t.Fatalf("state id %q, sent %v", d.Id(), fc.instances["inst-sync"]["id"])
	}
}

func TestApplicationInstance_RefreshRemovesMissing(t *testing.T) {
	fc := newFakeController()
	meta := newFakeMeta(t, fc, true)
	d := newInstanceData(t, "gone")
	d.SetId("11111111-1111-1111-1111-111111111111")

	if diags := resourceReadApplicationInstance(context.Background(), d, meta); diags.HasError() {
		t.Fatalf("refresh of a missing instance must not fail: %v", diags)
	}
	if d.Id() != "" {
		t.Fatalf("id = %q, want it removed from state", d.Id())
	}
}

func TestApplicationInstance_DataSourceErrorsOnMissing(t *testing.T) {
	fc := newFakeController()
	meta := newFakeMeta(t, fc, true)
	d := newInstanceData(t, "gone")

	if diags := ReadApplicationInstance(context.Background(), d, meta); !diags.HasError() {
		t.Fatal("the data source must fail for a missing instance")
	}
}

func TestApplicationInstance_DeleteMissingIsSuccess(t *testing.T) {
	fc := newFakeController()
	meta := newFakeMeta(t, fc, true)
	d := newInstanceData(t, "gone")
	d.SetId("11111111-1111-1111-1111-111111111111")

	if diags := DeleteApplicationInstance(context.Background(), d, meta); diags.HasError() {
		t.Fatalf("delete: %v", diags)
	}
	if d.Id() != "" {
		t.Fatalf("id = %q", d.Id())
	}
}

func TestApplicationInstance_ReadAfterWriteRetriesNotFound(t *testing.T) {
	old := readAfterWriteTimeout
	readAfterWriteTimeout = 3 * time.Second
	t.Cleanup(func() { readAfterWriteTimeout = old })

	fc := newFakeController()
	meta := newFakeMeta(t, fc, true)
	d := newInstanceData(t, "late")
	d.SetId("22222222-2222-2222-2222-222222222222")

	// The instance becomes readable a moment after the write.
	go func() {
		time.Sleep(300 * time.Millisecond)
		fc.mu.Lock()
		fc.instances["late"] = map[string]interface{}{"id": d.Id(), "name": "late", "title": "t", "appId": "a", "deviceId": "d", "projectId": "p"}
		fc.mu.Unlock()
	}()

	if diags := readApplicationInstanceAfterWrite(context.Background(), d, meta); diags.HasError() {
		t.Fatalf("read after write: %v", diags)
	}
	if d.Id() != "22222222-2222-2222-2222-222222222222" {
		t.Fatalf("id = %q", d.Id())
	}
}

func TestApplicationInstance_ReadAfterWriteRejectsOtherID(t *testing.T) {
	fc := newFakeController()
	fc.instances["taken"] = map[string]interface{}{"id": "33333333-3333-3333-3333-333333333333", "name": "taken"}
	meta := newFakeMeta(t, fc, true)
	d := newInstanceData(t, "taken")
	d.SetId("44444444-4444-4444-4444-444444444444")

	diags := readApplicationInstanceAfterWrite(context.Background(), d, meta)
	if !diags.HasError() || !strings.Contains(fmt.Sprint(diags), "terraform import") {
		t.Fatalf("diags = %v", diags)
	}
	if d.Id() != "44444444-4444-4444-4444-444444444444" {
		t.Fatalf("read after write must never change or drop the id, got %q", d.Id())
	}
}

func TestApplicationInstance_GatewayHTMLErrorIsReported(t *testing.T) {
	fc := newFakeController()
	fc.gatewayHTML = true
	meta := newFakeMeta(t, fc, false)
	d := newInstanceData(t, "x")
	d.SetId("55555555-5555-5555-5555-555555555555")

	diags := resourceReadApplicationInstance(context.Background(), d, meta)
	msg := fmt.Sprint(diags)
	if !diags.HasError() || strings.Contains(msg, "invalid character") || !strings.Contains(msg, "504") {
		t.Fatalf("diags = %s, want the HTTP 504 status, not a JSON decode error", msg)
	}
}

func TestNewAppInstanceID_IsUUIDv4(t *testing.T) {
	re := regexp.MustCompile("^[a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89ab][a-f0-9]{3}-[a-f0-9]{12}$")
	seen := map[string]bool{}
	for i := 0; i < 100; i++ {
		id, err := newAppInstanceID()
		if err != nil || !re.MatchString(id) || seen[id] {
			t.Fatalf("id %q (err %v) is not a fresh UUID v4", id, err)
		}
		seen[id] = true
	}
}

func TestApplicationInstanceSchema_ServerFilledBlocksAreComputed(t *testing.T) {
	s := zschema.ApplicationInstance()
	for _, k := range []string{"logs", "manifest_info", "vminfo", "tags"} {
		if !s[k].Optional || !s[k].Computed {
			t.Errorf("%s: Optional=%v Computed=%v, want both (UE-177 perpetual diff)", k, s[k].Optional, s[k].Computed)
		}
	}
}

func TestApplicationInstanceResource_Timeouts(t *testing.T) {
	r := ApplicationInstanceResource()
	if r.Timeouts == nil || r.Timeouts.Create == nil || *r.Timeouts.Create != appInstanceWriteTimeout {
		t.Fatalf("timeouts = %+v", r.Timeouts)
	}
	if ApplicationResource().Timeouts == nil {
		t.Fatal("zedcloud_application has no timeouts")
	}
}
