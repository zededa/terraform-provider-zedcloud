package resources

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/go-openapi/runtime"
	"github.com/zededa/terraform-provider-zedcloud/v2/client/jobs"
)

// fakeJobs answers Get from a script of results, repeating the last one.
type fakeJobs struct {
	script []fakeJobResult
	calls  int
}

type fakeJobResult struct {
	status jobs.Status
	err    error
	detail string
}

func (f *fakeJobs) Get(_ context.Context, id, objectType string) (*jobs.Job, error) {
	r := f.script[len(f.script)-1]
	if f.calls < len(f.script) {
		r = f.script[f.calls]
	}
	f.calls++
	if r.err != nil {
		return nil, r.err
	}
	return &jobs.Job{ID: id, Status: r.status, ErrorDetailOfJob: r.detail}, nil
}

func (f *fakeJobs) SetTransport(runtime.ClientTransport) {}

func testPoller(f *fakeJobs) *jobPoller {
	return &jobPoller{
		client:      f,
		objectType:  jobs.ObjectTypeEdgeAppInstance,
		intervalMin: time.Millisecond,
		intervalMax: 4 * time.Millisecond,
		sleep:       func(context.Context, time.Duration) error { return nil },
	}
}

func TestJobPoller_Completed(t *testing.T) {
	f := &fakeJobs{script: []fakeJobResult{
		{status: jobs.StatusInitialized}, {status: jobs.StatusInProgress}, {status: jobs.StatusCompleted},
	}}
	job, err := testPoller(f).wait(context.Background(), "job-1", time.Minute)
	if err != nil || job == nil || job.Status != jobs.StatusCompleted {
		t.Fatalf("job = %+v, err = %v", job, err)
	}
	if f.calls != 3 {
		t.Errorf("calls = %d, want 3", f.calls)
	}
}

func TestJobPoller_FailedCarriesErrorText(t *testing.T) {
	f := &fakeJobs{script: []fakeJobResult{
		{status: jobs.StatusInProgress}, {status: jobs.StatusFailed, detail: "app instance x already exists"},
	}}
	_, err := testPoller(f).wait(context.Background(), "job-1", time.Minute)
	var failed *jobFailedError
	if !errors.As(err, &failed) || !strings.Contains(err.Error(), "already exists") {
		t.Fatalf("err = %v, want a jobFailedError with the job's error text", err)
	}
}

func TestJobPoller_Aborted(t *testing.T) {
	f := &fakeJobs{script: []fakeJobResult{{status: jobs.StatusAborted}}}
	_, err := testPoller(f).wait(context.Background(), "job-1", time.Minute)
	var failed *jobFailedError
	if !errors.As(err, &failed) {
		t.Fatalf("err = %v, want a jobFailedError", err)
	}
}

func TestJobPoller_ToleratesTransientErrors(t *testing.T) {
	transient := &jobs.APIError{Status: http.StatusInternalServerError}
	f := &fakeJobs{script: []fakeJobResult{
		{err: transient}, {err: transient}, {status: jobs.StatusInProgress}, {err: transient}, {status: jobs.StatusCompleted},
	}}
	if _, err := testPoller(f).wait(context.Background(), "job-1", time.Minute); err != nil {
		t.Fatalf("err = %v", err)
	}
}

func TestJobPoller_GivesUpAfterRepeatedErrors(t *testing.T) {
	f := &fakeJobs{script: []fakeJobResult{{err: &jobs.APIError{Status: http.StatusInternalServerError}}}}
	_, err := testPoller(f).wait(context.Background(), "job-1", time.Minute)
	if err == nil || f.calls != maxJobPollErrors {
		t.Fatalf("err = %v after %d calls, want an error after %d", err, f.calls, maxJobPollErrors)
	}
}

func TestJobPoller_ForbiddenFailsAtOnce(t *testing.T) {
	f := &fakeJobs{script: []fakeJobResult{{err: &jobs.APIError{Status: http.StatusForbidden}}}}
	_, err := testPoller(f).wait(context.Background(), "job-1", time.Minute)
	if err == nil || f.calls != 1 || !strings.Contains(err.Error(), "read access") {
		t.Fatalf("err = %v after %d calls", err, f.calls)
	}
}

func TestJobPoller_Timeout(t *testing.T) {
	f := &fakeJobs{script: []fakeJobResult{{status: jobs.StatusInProgress}}}
	p := testPoller(f)
	p.intervalMin, p.intervalMax = time.Hour, time.Hour // the next poll would pass the deadline
	_, err := p.wait(context.Background(), "job-1", time.Minute)
	var te *jobTimeoutError
	if !errors.As(err, &te) || te.LastStatus != jobs.StatusInProgress {
		t.Fatalf("err = %v, want a jobTimeoutError with the last status", err)
	}
}

func TestJobPoller_ContextCancelled(t *testing.T) {
	f := &fakeJobs{script: []fakeJobResult{{status: jobs.StatusInProgress}}}
	p := testPoller(f)
	p.sleep = sleepCtx
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := p.wait(ctx, "job-1", time.Minute)
	var te *jobTimeoutError
	if !errors.As(err, &te) {
		t.Fatalf("err = %v, want a jobTimeoutError", err)
	}
}

func TestJitter(t *testing.T) {
	for i := 0; i < 1000; i++ {
		got := jitter(10 * time.Second)
		if got < 8*time.Second || got > 12*time.Second {
			t.Fatalf("jitter(10s) = %s, want within ±20%%", got)
		}
	}
}
