// Package artifact is a hand-written client for the zedcloud ArtifactManager
// API (/v1/artifacts): small files such as Edge app logos, custom licenses and
// brand/model logos.
//
// It is hand-written rather than generated because the published swagger is
// wrong for the upload operation: it declares a JSON string/byte body, while
// gilas only accepts a raw application/octet-stream body (ENG-3009 item 5).
//
// Behaviour this client encodes, verified against a controller (see
// docs/design/nfr-19-edge-app-logo-artifacts.md §4.7):
//
//   - Upload is sent as ONE chunk. Multi-chunk uploads are accepted with 202 but
//     never materialise (ENG-3009 item 1).
//   - Content-Range uses an exclusive end, "bytes 0-<N>/<N>", the form the UI
//     sends. gilas ignores the end value today.
//   - "Not found" is 307 Temporary Redirect, documented in the proto as "not
//     available at the requested time". It is returned for ids that never
//     existed, were never uploaded, were deleted, or whose upload failed.
//     GetSignedURL maps 307 (and 404) to ErrNotFound.
//   - DELETE returns 200 whether or not the artifact exists.
//   - The list endpoint (GET /v1/artifacts) times out, so it is not exposed.
package artifact

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

// ErrNotFound is returned by GetSignedURL when the controller reports the
// artifact as unavailable (HTTP 307 or 404).
var ErrNotFound = errors.New("artifact not found")

// Artifact is the subset of the controller's Artifact message this client uses.
type Artifact struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	SignedURL string `json:"signedUrl,omitempty"`
	TTL       string `json:"ttl,omitempty"`
}

// APIError is a non-success HTTP response from the artifact API.
type APIError struct {
	Op     string
	Status int
	Body   string
}

func (e *APIError) Error() string {
	if e.Body == "" {
		return fmt.Sprintf("artifact %s: unexpected HTTP status %d", e.Op, e.Status)
	}
	return fmt.Sprintf("artifact %s: unexpected HTTP status %d: %s", e.Op, e.Status, e.Body)
}

// Code returns the HTTP status, so callers that test for a Code() int (such as
// resources.isStatusNotFound) recognise this error.
func (e *APIError) Code() int { return e.Status }

// ClientService is the artifact API.
type ClientService interface {
	// Create registers a new artifact and returns it with its generated id
	// ("<uuid>_<name>"). Nothing is stored until Upload is called.
	Create(ctx context.Context, name string) (*Artifact, error)
	// Upload sends data as a single chunk. The controller answers 202 before
	// the object is stored; poll GetSignedURL to confirm it landed.
	Upload(ctx context.Context, id string, data []byte) error
	// GetSignedURL returns a short-lived signed download URL, or ErrNotFound.
	GetSignedURL(ctx context.Context, id string) (*Artifact, error)
	// Delete removes the artifact. The controller succeeds for missing ids too.
	Delete(ctx context.Context, id string) error

	SetTransport(transport runtime.ClientTransport)
}

// New creates a new artifact API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// Client is the artifact API client.
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// SetTransport changes the transport on the client.
func (c *Client) SetTransport(transport runtime.ClientTransport) {
	c.transport = transport
}

// ContentRange returns the Content-Range header value for a single-chunk upload
// of size bytes, using the exclusive-end form the zedcloud UI sends.
func ContentRange(size int) string {
	return fmt.Sprintf("bytes 0-%d/%d", size, size)
}

func (c *Client) Create(ctx context.Context, name string) (*Artifact, error) {
	if name == "" {
		return nil, errors.New("artifact create: name must not be empty")
	}
	res, err := c.submit(ctx, "create", http.MethodPost, "/v1/artifacts", "application/json",
		func(r runtime.ClientRequest) error {
			return r.SetBodyParam(&Artifact{Name: name})
		},
		func(resp runtime.ClientResponse) (interface{}, error) {
			if resp.Code() != http.StatusOK {
				return nil, apiError("create", resp)
			}
			var a Artifact
			if err := json.NewDecoder(resp.Body()).Decode(&a); err != nil {
				return nil, fmt.Errorf("artifact create: decoding response: %w", err)
			}
			if a.ID == "" {
				return nil, errors.New("artifact create: response carried no id")
			}
			return &a, nil
		})
	if err != nil {
		return nil, err
	}
	return res.(*Artifact), nil
}

func (c *Client) Upload(ctx context.Context, id string, data []byte) error {
	if id == "" {
		return errors.New("artifact upload: id must not be empty")
	}
	if len(data) == 0 {
		return errors.New("artifact upload: content must not be empty")
	}
	_, err := c.submit(ctx, "upload", http.MethodPut, "/v1/artifacts/id/{id}/upload/chunked", runtime.DefaultMime,
		func(r runtime.ClientRequest) error {
			if err := r.SetPathParam("id", id); err != nil {
				return err
			}
			if err := r.SetHeaderParam("Content-Range", ContentRange(len(data))); err != nil {
				return err
			}
			return r.SetBodyParam(data)
		},
		func(resp runtime.ClientResponse) (interface{}, error) {
			if resp.Code() < 200 || resp.Code() > 299 {
				return nil, apiError("upload", resp)
			}
			return nil, nil
		})
	return err
}

func (c *Client) GetSignedURL(ctx context.Context, id string) (*Artifact, error) {
	if id == "" {
		return nil, errors.New("artifact read: id must not be empty")
	}
	res, err := c.submit(ctx, "read", http.MethodGet, "/v1/artifacts/id/{id}/url", "application/json",
		func(r runtime.ClientRequest) error {
			return r.SetPathParam("id", id)
		},
		func(resp runtime.ClientResponse) (interface{}, error) {
			switch resp.Code() {
			case http.StatusOK:
				var a Artifact
				if err := json.NewDecoder(resp.Body()).Decode(&a); err != nil {
					return nil, fmt.Errorf("artifact read: decoding response: %w", err)
				}
				if a.SignedURL == "" {
					// Success without a URL: treat as not available yet.
					return nil, ErrNotFound
				}
				return &a, nil
			case http.StatusTemporaryRedirect, http.StatusNotFound:
				return nil, ErrNotFound
			default:
				return nil, apiError("read", resp)
			}
		})
	if err != nil {
		// The provider's retrying HTTP client retries GET 404s and, once it
		// gives up, returns a transport error instead of the response.
		if !errors.Is(err, ErrNotFound) && strings.Contains(err.Error(), "404 Not Found") {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return res.(*Artifact), nil
}

func (c *Client) Delete(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("artifact delete: id must not be empty")
	}
	_, err := c.submit(ctx, "delete", http.MethodDelete, "/v1/artifacts/id/{id}", "application/json",
		func(r runtime.ClientRequest) error {
			return r.SetPathParam("id", id)
		},
		func(resp runtime.ClientResponse) (interface{}, error) {
			if resp.Code() == http.StatusNotFound || (resp.Code() >= 200 && resp.Code() <= 299) {
				return nil, nil
			}
			return nil, apiError("delete", resp)
		})
	return err
}

func (c *Client) submit(
	ctx context.Context,
	op, method, path, consumes string,
	write func(runtime.ClientRequest) error,
	read func(runtime.ClientResponse) (interface{}, error),
) (interface{}, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	return c.transport.Submit(&runtime.ClientOperation{
		ID:                 "Artifact_" + op,
		Method:             method,
		PathPattern:        path,
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{consumes},
		Schemes:            []string{"https"},
		Params: runtime.ClientRequestWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
			return write(r)
		}),
		Reader: runtime.ClientResponseReaderFunc(func(resp runtime.ClientResponse, _ runtime.Consumer) (interface{}, error) {
			return read(resp)
		}),
		Context: ctx,
	})
}

func apiError(op string, resp runtime.ClientResponse) error {
	var body string
	if b := resp.Body(); b != nil {
		raw, _ := io.ReadAll(io.LimitReader(b, 2048))
		body = strings.TrimSpace(string(raw))
	}
	return &APIError{Op: op, Status: resp.Code(), Body: body}
}
