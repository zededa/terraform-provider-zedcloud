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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	zschema "github.com/zededa/terraform-provider-zedcloud/v2/schemas"
)

// fakeController is a minimal in-memory zedcloud for the app-instance routes
// the provider uses (UE-177). No controller or credentials are needed.
type fakeController struct {
	mu        sync.Mutex
	instances map[string]map[string]interface{} // by name
	jobs      map[string][]string               // job id -> remaining statuses
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
		jobs:      map[string][]string{},
		jobErr:    map[string]string{},
	}
}

func (f *fakeController) newJob(statuses ...string) string {
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
		if f.failCreate == name {
			job := f.newJob("JOB_STATUS_INPROGRESS", "JOB_STATUS_FAILED")
			f.jobErr[job] = "app instance " + name + " already exists"
			writeJSON(w, http.StatusOK, map[string]interface{}{"objectId": job})
			return
		}
		f.instances[name] = body
		writeJSON(w, http.StatusOK, map[string]interface{}{"objectId": f.newJob("JOB_STATUS_INPROGRESS", "JOB_STATUS_COMPLETED")})

	case r.Method == http.MethodPost && path == "/v1/apps/instances":
		name, _ := body["name"].(string)
		f.instances[name] = body
		writeJSON(w, http.StatusOK, map[string]interface{}{"objectId": body["id"]})

	case strings.HasPrefix(path, "/v1/jobs/id/"):
		parts := strings.Split(strings.TrimPrefix(path, "/v1/jobs/id/"), "/")
		statuses, ok := f.jobs[parts[0]]
		if !ok || len(parts) != 3 || parts[2] != "OBJECT_TYPE_EDGE_APP_INSTANCE" {
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
		writeJSON(w, http.StatusOK, map[string]interface{}{"objectId": f.newJob("JOB_STATUS_COMPLETED")})

	case r.Method == http.MethodDelete && strings.HasPrefix(path, "/v2/apps/instances/id/"):
		inst := f.byID(strings.TrimPrefix(path, "/v2/apps/instances/id/"))
		if inst == nil {
			notFound(w)
			return
		}
		delete(f.instances, inst["name"].(string))
		writeJSON(w, http.StatusOK, map[string]interface{}{"objectId": f.newJob("JOB_STATUS_INPROGRESS", "JOB_STATUS_COMPLETED")})

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
func newFakeMeta(t *testing.T, fc *fakeController) *api_client.ZedcloudAPI {
	t.Helper()
	t.Setenv("TF_INSECURE_SKIP_TLS_VERIFY", "true")
	srv := httptest.NewTLSServer(fc)
	t.Cleanup(srv.Close)

	opts := httpOptions{requestTimeout: 5 * time.Second, maxRetries: 0, retryWaitMax: 10 * time.Millisecond}
	std := getRetryClient(opts).StandardClient()
	std.Transport = NewHttpTransportWrapper(std.Transport)
	transport := httptransport.NewWithClient(strings.TrimPrefix(srv.URL, "https://"), "/api", []string{"https"}, std)

	return api_client.New(transport, strfmt.Default)
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

func TestApplicationInstance_CreateSendsProviderID(t *testing.T) {
	fc := newFakeController()
	meta := newFakeMeta(t, fc)
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
	meta := newFakeMeta(t, fc)
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
	meta := newFakeMeta(t, fc)
	d := newInstanceData(t, "gone")

	if diags := ReadApplicationInstance(context.Background(), d, meta); !diags.HasError() {
		t.Fatal("the data source must fail for a missing instance")
	}
}

func TestApplicationInstance_DeleteMissingIsSuccess(t *testing.T) {
	fc := newFakeController()
	meta := newFakeMeta(t, fc)
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
	meta := newFakeMeta(t, fc)
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
	meta := newFakeMeta(t, fc)
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
	meta := newFakeMeta(t, fc)
	d := newInstanceData(t, "x")
	d.SetId("55555555-5555-5555-5555-555555555555")

	diags := resourceReadApplicationInstance(context.Background(), d, meta)
	msg := fmt.Sprint(diags)
	if !diags.HasError() || strings.Contains(msg, "invalid character") || !strings.Contains(msg, "504") {
		t.Fatalf("diags = %s, want the HTTP 504 status, not a JSON decode error", msg)
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
