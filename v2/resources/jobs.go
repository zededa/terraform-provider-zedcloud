package resources

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/zededa/terraform-provider-zedcloud/v2/client/jobs"
)

// maxJobPollErrors is how many job reads in a row may fail before waitForJob
// gives up. Job reads go through seine's Kafka consumer, which can be slow or
// briefly unavailable under the very load async writes are meant to absorb.
const maxJobPollErrors = 5

// jobFailedError is returned by waitForJob when the job finished in FAILED or
// ABORTED.
type jobFailedError struct {
	Job *jobs.Job
}

func (e *jobFailedError) Error() string {
	return fmt.Sprintf("job %s finished with status %s: %s", e.Job.ID, e.Job.Status, e.Job.ErrorText())
}

// jobTimeoutError is returned by waitForJob when the job did not finish in
// time. The job keeps running on the controller.
type jobTimeoutError struct {
	JobID      string
	LastStatus jobs.Status
	Timeout    time.Duration
}

func (e *jobTimeoutError) Error() string {
	status := string(e.LastStatus)
	if status == "" {
		status = "unknown"
	}
	return fmt.Sprintf("job %s did not finish within %s (last status: %s). It is still running on ZEDCloud; "+
		"re-run terraform apply to pick up the result, or raise the resource's timeouts", e.JobID, e.Timeout, status)
}

// jobPoller polls one job until it reaches a final state.
type jobPoller struct {
	client      jobs.ClientService
	objectType  string
	intervalMin time.Duration
	intervalMax time.Duration
	// sleep is replaceable in tests.
	sleep func(ctx context.Context, d time.Duration) error
}

func sleepCtx(ctx context.Context, d time.Duration) error {
	t := time.NewTimer(d)
	defer t.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-t.C:
		return nil
	}
}

// jitter spreads polls of parallel resources apart by ±20%.
func jitter(d time.Duration) time.Duration {
	if d <= 0 {
		return d
	}
	delta := time.Duration(rand.Int63n(int64(d)/5*2+1)) - d/5 //nolint:gosec // not security relevant
	return d + delta
}

// wait polls jobID until it is COMPLETED (nil), FAILED or ABORTED
// (*jobFailedError), or timeout passes (*jobTimeoutError). Transient read
// errors are tolerated up to maxJobPollErrors in a row; 401 and 403 fail at
// once.
func (p *jobPoller) wait(ctx context.Context, jobID string, timeout time.Duration) (*jobs.Job, error) {
	if p.sleep == nil {
		p.sleep = sleepCtx
	}
	deadline := time.Now().Add(timeout)
	interval := p.intervalMin
	var (
		last       jobs.Status
		readErrors int
	)
	for {
		job, err := p.client.Get(ctx, jobID, p.objectType)
		switch {
		case err != nil:
			if ctx.Err() != nil {
				return nil, &jobTimeoutError{JobID: jobID, LastStatus: last, Timeout: timeout}
			}
			var apiErr *jobs.APIError
			if errors.As(err, &apiErr) && (apiErr.Status == http.StatusUnauthorized || apiErr.Status == http.StatusForbidden) {
				return nil, fmt.Errorf("reading job %s: %w (the token needs read access to %s jobs)", jobID, err, p.objectType)
			}
			readErrors++
			log.Printf("[WARN] reading job %s (attempt %d of %d): %s", jobID, readErrors, maxJobPollErrors, err)
			if readErrors >= maxJobPollErrors {
				return nil, fmt.Errorf("reading job %s: %w", jobID, err)
			}
		case job.Status == jobs.StatusCompleted:
			return job, nil
		case job.Status == jobs.StatusFailed || job.Status == jobs.StatusAborted:
			return job, &jobFailedError{Job: job}
		default:
			readErrors = 0
			if job.Status != last {
				log.Printf("[DEBUG] job %s: %s", jobID, job.Status)
			}
			last = job.Status
		}

		wait := jitter(interval)
		if time.Now().Add(wait).After(deadline) {
			// One last read right at the deadline is not worth it; the
			// controller keeps going regardless.
			return nil, &jobTimeoutError{JobID: jobID, LastStatus: last, Timeout: timeout}
		}
		if err := p.sleep(ctx, wait); err != nil {
			return nil, &jobTimeoutError{JobID: jobID, LastStatus: last, Timeout: timeout}
		}
		interval *= 2
		if interval > p.intervalMax {
			interval = p.intervalMax
		}
	}
}
