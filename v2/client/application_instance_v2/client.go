// Package application_instance_v2 is a hand-written client for the async
// app-instance write routes:
//
//	POST   /v2/apps/instances
//	PUT    /v2/apps/instances/id/{id}
//	DELETE /v2/apps/instances/id/{id}
//
// Each answers 200 as soon as seine has recorded a job, and does the work in
// the background. The caller polls the job (see client/jobs) until it reaches
// a final state. See docs/design/ue-177-high-parallelism-robustness.md §3.
//
// It is hand-written rather than generated because the response does not
// match what the swagger implies: the job id is returned in
// ZsrvResponse.objectId, not in jobId, for all three routes
// (srvs/seine/appinstproc.go, zedcloud main cdedb70283). gilas returns an
// AppInstanceServiceResp instead when seine fills its jobID, so both shapes
// are accepted.
//
// Other behaviour this client encodes:
//
//   - Create keeps a caller-supplied id (it must be a UUID). The provider uses
//     that to know the instance id before the job finishes.
//   - Update requires the body id to equal the path id (400 otherwise).
//   - Delete of an id that does not exist answers 404 synchronously.
//   - A create whose name or id already exists still answers 200; the
//     conflict is reported by the job failing.
package application_instance_v2

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

// APIError is a non-success HTTP response. It implements Code() so
// resources.isStatusNotFound recognises a 404, and GetPayload()/String() so
// resources.ZsrvResponderToDiags can report the controller's error details.
type APIError struct {
	Op      string
	Status  int
	Body    string
	Payload *models.ZsrvResponse
}

func (e *APIError) Error() string {
	msg := e.Body
	if e.Payload != nil {
		var details []string
		for _, ze := range e.Payload.Error {
			if ze != nil && ze.Details != "" {
				details = append(details, ze.Details)
			}
		}
		if len(details) > 0 {
			msg = strings.Join(details, "; ")
		} else if e.Payload.HTTPStatusMsg != "" {
			msg = e.Payload.HTTPStatusMsg
		}
	}
	if msg == "" {
		return fmt.Sprintf("app instance %s: unexpected HTTP status %d", e.Op, e.Status)
	}
	return fmt.Sprintf("app instance %s: unexpected HTTP status %d: %s", e.Op, e.Status, msg)
}

// Code returns the HTTP status.
func (e *APIError) Code() int { return e.Status }

// GetPayload returns the decoded ZsrvResponse, if the body was one.
func (e *APIError) GetPayload() *models.ZsrvResponse { return e.Payload }

// String returns the error message.
func (e *APIError) String() string { return e.Error() }

// ClientService is the async app-instance write API. Every method returns the
// id of the job that carries out the request.
type ClientService interface {
	Create(ctx context.Context, body *models.AppInstance) (string, error)
	Update(ctx context.Context, id string, body *models.AppInstance) (string, error)
	Delete(ctx context.Context, id string) (string, error)

	SetTransport(transport runtime.ClientTransport)
}

// New creates a new client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// Client is the async app-instance write client.
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// SetTransport changes the transport on the client.
func (c *Client) SetTransport(transport runtime.ClientTransport) {
	c.transport = transport
}

func (c *Client) Create(ctx context.Context, body *models.AppInstance) (string, error) {
	if body == nil {
		return "", errors.New("app instance create: body must not be nil")
	}
	return c.submit(ctx, "create", http.MethodPost, "/v2/apps/instances", "", body)
}

func (c *Client) Update(ctx context.Context, id string, body *models.AppInstance) (string, error) {
	if id == "" {
		return "", errors.New("app instance update: id must not be empty")
	}
	if body == nil {
		return "", errors.New("app instance update: body must not be nil")
	}
	if body.ID != id {
		return "", fmt.Errorf("app instance update: body id %q does not match id %q", body.ID, id)
	}
	return c.submit(ctx, "update", http.MethodPut, "/v2/apps/instances/id/{id}", id, body)
}

func (c *Client) Delete(ctx context.Context, id string) (string, error) {
	if id == "" {
		return "", errors.New("app instance delete: id must not be empty")
	}
	return c.submit(ctx, "delete", http.MethodDelete, "/v2/apps/instances/id/{id}", id, nil)
}

// jobResponse covers both success shapes: ZsrvResponse (job id in objectId)
// and AppInstanceServiceResp (job id in jobID).
type jobResponse struct {
	ObjectID string `json:"objectId"`
	JobID    string `json:"jobID"`
	JobId    string `json:"jobId"` //nolint:revive,stylecheck // ZsrvResponse spelling
}

func (r *jobResponse) jobID() string {
	switch {
	case r.JobID != "":
		return r.JobID
	case r.JobId != "":
		return r.JobId
	default:
		return r.ObjectID
	}
}

func (c *Client) submit(ctx context.Context, op, method, path, id string, body *models.AppInstance) (string, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	res, err := c.transport.Submit(&runtime.ClientOperation{
		ID:                 "AppInstanceV2_" + op,
		Method:             method,
		PathPattern:        path,
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params: runtime.ClientRequestWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
			if id != "" {
				if err := r.SetPathParam("id", id); err != nil {
					return err
				}
			}
			if body != nil {
				return r.SetBodyParam(body)
			}
			return nil
		}),
		Reader: runtime.ClientResponseReaderFunc(func(resp runtime.ClientResponse, _ runtime.Consumer) (interface{}, error) {
			raw := readAll(resp)
			if resp.Code() < 200 || resp.Code() > 299 {
				return nil, newAPIError(op, resp.Code(), raw)
			}
			var jr jobResponse
			if err := json.Unmarshal(raw, &jr); err != nil {
				return nil, fmt.Errorf("app instance %s: decoding response: %w", op, err)
			}
			if jr.jobID() == "" {
				return nil, fmt.Errorf("app instance %s: response carried no job id: %s", op, truncate(string(raw)))
			}
			return jr.jobID(), nil
		}),
		Context: ctx,
	})
	if err != nil {
		return "", err
	}
	return res.(string), nil
}

func newAPIError(op string, status int, raw []byte) *APIError {
	e := &APIError{Op: op, Status: status, Body: truncate(string(bytes.TrimSpace(raw)))}
	var zr models.ZsrvResponse
	if len(raw) > 0 && json.Unmarshal(raw, &zr) == nil {
		e.Payload = &zr
	}
	return e
}

func readAll(resp runtime.ClientResponse) []byte {
	b := resp.Body()
	if b == nil {
		return nil
	}
	raw, _ := io.ReadAll(io.LimitReader(b, 1<<20))
	return raw
}

func truncate(s string) string {
	const max = 1024
	if len(s) > max {
		return s[:max] + "…"
	}
	return s
}
