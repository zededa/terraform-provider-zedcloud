// Package jobs is a hand-written client for reading zedcloud jobs
// (GET /v1/jobs/id/{id}/objectType/{objectType}).
//
// Async write routes, such as /v2/apps/instances, answer 200 with a job id and
// do the work in the background. The caller polls the job until it reaches a
// final state. See docs/design/ue-177-high-parallelism-robustness.md §3, §4.6.
//
// It is hand-written because the job service is not part of the generated
// client, and the provider only needs this one read.
//
// Behaviour this client encodes (zedcloud main, cdedb70283):
//
//   - The objectType path segment selects the RBAC check. Without it gilas
//     checks device access, so users who only have app-instance permissions
//     would be refused (srvs/gilas/job_handlers.go getJobAccessType).
//   - A job the controller cannot find is reported as a generic error, not
//     404 (srvs/seine/jobproc.go jobRead), so callers must not treat a failed
//     read as "job gone".
//   - Responses are jsonpb: zero values, such as an UNSPECIFIED status, are
//     omitted.
package jobs

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// Status is a zedcloud JobStatus.
type Status string

const (
	StatusUnspecified Status = "JOB_STATUS_UNSPECIFIED"
	StatusInitialized Status = "JOB_STATUS_INITIALIZED"
	StatusReady       Status = "JOB_STATUS_READY"
	StatusInProgress  Status = "JOB_STATUS_INPROGRESS"
	StatusCompleted   Status = "JOB_STATUS_COMPLETED"
	StatusFailed      Status = "JOB_STATUS_FAILED"
	StatusAborted     Status = "JOB_STATUS_ABORTED"
)

// Final reports whether s is a terminal job status.
func (s Status) Final() bool {
	return s == StatusCompleted || s == StatusFailed || s == StatusAborted
}

// ObjectTypeEdgeAppInstance is the objectType for app-instance jobs.
const ObjectTypeEdgeAppInstance = "OBJECT_TYPE_EDGE_APP_INSTANCE"

// Job is the subset of the controller's JobConfig message this client uses.
type Job struct {
	ID                 string            `json:"id,omitempty"`
	Name               string            `json:"name,omitempty"`
	Status             Status            `json:"status,omitempty"`
	ObjectList         []string          `json:"objectList,omitempty"`
	IndividualOpStatus map[string]string `json:"individualOpStatus,omitempty"`
	ErrorDetailOfJob   string            `json:"errorDetailOfJob,omitempty"`
}

// ErrorText returns the job's error details: errorDetailOfJob plus every
// object whose individual status is not SUCCESS.
func (j *Job) ErrorText() string {
	var parts []string
	if s := strings.TrimSpace(j.ErrorDetailOfJob); s != "" {
		parts = append(parts, s)
	}
	for obj, st := range j.IndividualOpStatus {
		if st == "" || strings.HasSuffix(st, "SUCCESS") {
			continue
		}
		parts = append(parts, fmt.Sprintf("%s: %s", obj, st))
	}
	if len(parts) == 0 {
		return "no error details reported"
	}
	return strings.Join(parts, "; ")
}

// APIError is a non-success HTTP response from the jobs API.
type APIError struct {
	Status int
	Body   string
}

func (e *APIError) Error() string {
	if e.Body == "" {
		return fmt.Sprintf("job read: unexpected HTTP status %d", e.Status)
	}
	return fmt.Sprintf("job read: unexpected HTTP status %d: %s", e.Status, e.Body)
}

// Code returns the HTTP status.
func (e *APIError) Code() int { return e.Status }

// ClientService is the jobs API.
type ClientService interface {
	// Get reads one job. objectType is a zedcloud ObjectType name, such as
	// ObjectTypeEdgeAppInstance.
	Get(ctx context.Context, id, objectType string) (*Job, error)

	SetTransport(transport runtime.ClientTransport)
}

// New creates a new jobs API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// Client is the jobs API client.
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// SetTransport changes the transport on the client.
func (c *Client) SetTransport(transport runtime.ClientTransport) {
	c.transport = transport
}

func (c *Client) Get(ctx context.Context, id, objectType string) (*Job, error) {
	if id == "" {
		return nil, errors.New("job read: id must not be empty")
	}
	if objectType == "" {
		return nil, errors.New("job read: objectType must not be empty")
	}
	if ctx == nil {
		ctx = context.Background()
	}
	res, err := c.transport.Submit(&runtime.ClientOperation{
		ID:                 "Jobs_get",
		Method:             http.MethodGet,
		PathPattern:        "/v1/jobs/id/{id}/objectType/{objectType}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params: runtime.ClientRequestWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
			if err := r.SetPathParam("id", id); err != nil {
				return err
			}
			return r.SetPathParam("objectType", objectType)
		}),
		Reader: runtime.ClientResponseReaderFunc(func(resp runtime.ClientResponse, _ runtime.Consumer) (interface{}, error) {
			if resp.Code() != http.StatusOK {
				return nil, &APIError{Status: resp.Code(), Body: readBody(resp)}
			}
			var j Job
			if err := json.NewDecoder(resp.Body()).Decode(&j); err != nil {
				return nil, fmt.Errorf("job read: decoding response: %w", err)
			}
			if j.ID == "" {
				// gilas answers with a ZsrvResponse instead of the job when
				// the read did not succeed but still returned 200.
				return nil, errors.New("job read: response carried no job")
			}
			return &j, nil
		}),
		Context: ctx,
	})
	if err != nil {
		return nil, err
	}
	return res.(*Job), nil
}

func readBody(resp runtime.ClientResponse) string {
	b := resp.Body()
	if b == nil {
		return ""
	}
	raw, _ := io.ReadAll(io.LimitReader(b, 2048))
	return strings.TrimSpace(string(raw))
}
