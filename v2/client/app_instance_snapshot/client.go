// Package app_instance_snapshot is a hand-written client for the zedcloud app
// instance snapshot API, plus the one device-status read the provider needs to
// check that an edge node supports immediate snapshots.
//
// It is hand-written rather than generated because the published swagger is
// incomplete (docs/design/ple-576-app-instance-immediate-snapshots.md §1.2):
//
//   - Create (POST /v1/apps/instances/id/{id}/snapshot) and activate
//     (PUT .../snapshot/name/{name}/activate) are gilas-only routes and are not
//     in the swagger.
//   - The state route (GET /v1/apps/instances/snapshot/id/{id}/state) is served
//     by ganges, not gilas, and answers protojson.
//
// Behaviour this client encodes:
//
//   - Create takes no body fields: the controller always makes an immediate
//     snapshot and generates its name. The snapshot id and name come back in
//     ZsrvResponse.objectId / objectName (ENG-1676). Creating a snapshot deletes
//     every other snapshot of the app instance.
//   - Get and List fill only id, name and config. state, activate and immediate
//     are always empty there; use State.
//   - State omits "state" when it is INVALID (protojson drops zero values).
//   - A missing snapshot is a 404. The provider's retrying HTTP client retries
//     GET 404s and then returns a transport error instead of the response; both
//     map to ErrNotFound.
//   - Delete answers 404 for a missing snapshot; Delete treats that as success.
package app_instance_snapshot

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

// ErrNotFound is returned when the controller reports the snapshot (or device
// status) as missing.
var ErrNotFound = errors.New("not found")

// Snapshot states, as the controller names them (app.SnapshotState).
const (
	StateInvalid        = "INVALID"
	StateIdle           = "IDLE"
	StateCreating       = "CREATING"
	StateCreated        = "CREATED"
	StateCreateFailed   = "CREATE_FAILED"
	StateRollingBack    = "ROLLING_BACK"
	StateActive         = "ACTIVE"
	StateRollbackFailed = "ROLLBACK_FAILED"
	StateDeleting       = "DELETING"
	StateDeleteFailed   = "DELETE_FAILED"
	StateDeleted        = "DELETED"
)

// ImmediateSnapshotCapability is the lowest EVE API capability level that
// supports immediate snapshots.
const ImmediateSnapshotCapability = "API_CAPABILITY_VOLUME_SNAPSHOTS_IMMEDIATE"

// immediateSnapshotLevel is ImmediateSnapshotCapability's value.
const immediateSnapshotLevel = 13

// apiCapabilityLevels maps EVE's APICapability enum names to their values
// (eveproto info.proto). The capability is a LEVEL: a device reports the
// highest one it implements, and that asserts every lower one. A controller
// must compare with >=, never test for one name. zcli's snapshot-create tests
// for the exact name, which refuses every EVE newer than level 13.
var apiCapabilityLevels = map[string]int{
	"API_CAPABILITY_UNSPECIFIED":                       0,
	"API_CAPABILITY_RETRY_UPDATE":                      1,
	"API_CAPABILITY_SHUTDOWN":                          2,
	"API_CAPABILITY_START_DELAY_IN_SECONDS":            3,
	"API_CAPABILITY_EDGEVIEW":                          4,
	"API_CAPABILITY_VOLUME_SNAPSHOTS":                  5,
	"API_CAPABILITY_NETWORK_INSTANCE_ROUTING":          6,
	"API_CAPABILITY_BOOT_MODE":                         7,
	"API_CAPABILITY_MTU":                               8,
	"API_CAPABILITY_ADAPTER_USER_LABELS":               9,
	"API_CAPABILITY_ENFORCED_NET_INTERFACE_ORDER":      10,
	"API_CAPABILITY_NTPS_FQDN":                         11,
	"API_CAPABILITY_WIN_LIC_PASSTHROUGH":               12,
	"API_CAPABILITY_VOLUME_SNAPSHOTS_IMMEDIATE":        13,
	"API_CAPABILITY_ENCRYPTED_PATCH_ENVELOPE":          14,
	"API_CAPABILITY_SINGLE_STACK_IP_NETWORK":           15,
	"API_CAPABILITY_CELLULAR_ATTACH_CONFIG":            16,
	"API_CAPABILITY_EDGEVIEW_AUTHENTICATION":           17,
	"API_CAPABILITY_DISABLE_VTPM":                      18,
	"API_CAPABILITY_LOC_REBOOT_COLLECT_INFO":           19,
	"API_CAPABILITY_SMART_REPORT":                      20,
	"API_CAPABILITY_REPORT_TPM_EVENTLOG":               21,
	"API_CAPABILITY_APP_INSTANCE_NET_INTERFACE_CHANGE": 22,
	"API_CAPABILITY_DEFERRED_QUEUE_METRICS":            23,
}

// SupportsImmediateSnapshots reads the API capability level from a device's
// rawStatus (EVE's ZInfoDevice as JSON) and reports whether it is at least
// API_CAPABILITY_VOLUME_SNAPSHOTS_IMMEDIATE. reported is the value as found,
// for messages. A name this table does not know is newer than every entry in
// it, so it counts as supported.
func SupportsImmediateSnapshots(rawStatus string) (ok bool, reported string, err error) {
	var info struct {
		APICapability      json.RawMessage `json:"apiCapability"`
		APICapabilitySnake json.RawMessage `json:"api_capability"`
	}
	if err := json.Unmarshal([]byte(rawStatus), &info); err != nil {
		return false, "", fmt.Errorf("decoding the edge node's raw status: %w", err)
	}
	v := info.APICapability
	if len(v) == 0 {
		v = info.APICapabilitySnake
	}
	if len(v) == 0 || string(v) == "null" {
		return false, "API_CAPABILITY_UNSPECIFIED", nil
	}
	var name string
	if json.Unmarshal(v, &name) == nil {
		level, known := apiCapabilityLevels[name]
		if !known {
			return true, name, nil
		}
		return level >= immediateSnapshotLevel, name, nil
	}
	var level int
	if err := json.Unmarshal(v, &level); err != nil {
		return false, string(v), fmt.Errorf("unexpected apiCapability value %s", v)
	}
	return level >= immediateSnapshotLevel, strconv.Itoa(level), nil
}

// Snapshot is the subset of the controller's AppInstanceSnapshot message this
// client uses.
type Snapshot struct {
	ID        string          `json:"id,omitempty"`
	Name      string          `json:"name,omitempty"`
	State     string          `json:"state,omitempty"`
	Activate  bool            `json:"activate,omitempty"`
	Immediate bool            `json:"immediate,omitempty"`
	Config    json.RawMessage `json:"config,omitempty"`
}

// AppInstanceID returns the id of the app instance the snapshot belongs to,
// taken from the stored config.
func (s *Snapshot) AppInstanceID() string {
	var c struct {
		ID string `json:"id"`
	}
	if len(s.Config) == 0 || json.Unmarshal(s.Config, &c) != nil {
		return ""
	}
	return c.ID
}

// ConfigJSON returns the snapshot's app instance config as compact JSON,
// without encryptedSecrets (zcli strips them too). Empty if there is no config.
func (s *Snapshot) ConfigJSON() (string, error) {
	if len(s.Config) == 0 {
		return "", nil
	}
	var m map[string]json.RawMessage
	if err := json.Unmarshal(s.Config, &m); err != nil {
		return "", fmt.Errorf("decoding snapshot config: %w", err)
	}
	delete(m, "encryptedSecrets")
	b, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// State is the controller's AppInstanceSnapshotState message.
type State struct {
	State     string `json:"state,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
}

// Name returns the state, with the omitted zero value spelled out.
func (s *State) Name() string {
	if s == nil || s.State == "" {
		return StateInvalid
	}
	return s.State
}

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
		return fmt.Sprintf("app instance snapshot %s: unexpected HTTP status %d", e.Op, e.Status)
	}
	return fmt.Sprintf("app instance snapshot %s: unexpected HTTP status %d: %s", e.Op, e.Status, msg)
}

// Code returns the HTTP status.
func (e *APIError) Code() int { return e.Status }

// GetPayload returns the decoded ZsrvResponse, if the body was one.
func (e *APIError) GetPayload() *models.ZsrvResponse { return e.Payload }

// String returns the error message.
func (e *APIError) String() string { return e.Error() }

// CreateResult identifies the snapshot a Create made.
type CreateResult struct {
	ID   string
	Name string
}

// ClientService is the app instance snapshot API.
type ClientService interface {
	// Create takes an immediate snapshot of the app instance. The result may be
	// empty on controllers older than ENG-1676; callers then List.
	Create(ctx context.Context, appInstanceID string) (*CreateResult, error)
	// Get reads a snapshot by its id, or returns ErrNotFound.
	Get(ctx context.Context, snapshotID string) (*Snapshot, error)
	// GetByName reads a snapshot by app instance id and name, or returns ErrNotFound.
	GetByName(ctx context.Context, appInstanceID, name string) (*Snapshot, error)
	// List returns the app instance's snapshots.
	List(ctx context.Context, appInstanceID string) ([]*Snapshot, error)
	// State returns the snapshot's state, or ErrNotFound.
	State(ctx context.Context, snapshotID string) (*State, error)
	// Activate rolls the app instance back to the named snapshot.
	Activate(ctx context.Context, appInstanceID, name string, force bool) error
	// Delete removes the named snapshot. A missing snapshot is not an error.
	Delete(ctx context.Context, appInstanceID, name string) error
	// DeviceRawStatus returns the rawStatus field of an edge node's status
	// (EVE's info message as JSON), or ErrNotFound.
	DeviceRawStatus(ctx context.Context, deviceID string) (string, error)

	SetTransport(transport runtime.ClientTransport)
}

// New creates a new client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// Client is the app instance snapshot client.
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// SetTransport changes the transport on the client.
func (c *Client) SetTransport(transport runtime.ClientTransport) {
	c.transport = transport
}

func (c *Client) Create(ctx context.Context, appInstanceID string) (*CreateResult, error) {
	if appInstanceID == "" {
		return nil, errors.New("app instance snapshot create: app instance id must not be empty")
	}
	raw, err := c.do(ctx, "create", http.MethodPost, "/v1/apps/instances/id/{id}/snapshot",
		map[string]string{"id": appInstanceID}, nil, struct{}{})
	if err != nil {
		return nil, err
	}
	var zr models.ZsrvResponse
	if len(raw) > 0 {
		if err := json.Unmarshal(raw, &zr); err != nil {
			return nil, fmt.Errorf("app instance snapshot create: decoding response: %w", err)
		}
	}
	return &CreateResult{ID: zr.ObjectID, Name: zr.ObjectName}, nil
}

func (c *Client) Get(ctx context.Context, snapshotID string) (*Snapshot, error) {
	if snapshotID == "" {
		return nil, errors.New("app instance snapshot read: id must not be empty")
	}
	raw, err := c.do(ctx, "read", http.MethodGet, "/v1/apps/instances/snapshot/id/{id}",
		map[string]string{"id": snapshotID}, nil, nil)
	if err != nil {
		return nil, err
	}
	return decodeSnapshot("read", raw)
}

func (c *Client) GetByName(ctx context.Context, appInstanceID, name string) (*Snapshot, error) {
	if appInstanceID == "" || name == "" {
		return nil, errors.New("app instance snapshot read: app instance id and name must not be empty")
	}
	raw, err := c.do(ctx, "read", http.MethodGet, "/v1/apps/instances/id/{id}/snapshot/name/{name}",
		map[string]string{"id": appInstanceID, "name": name}, nil, nil)
	if err != nil {
		return nil, err
	}
	return decodeSnapshot("read", raw)
}

func (c *Client) List(ctx context.Context, appInstanceID string) ([]*Snapshot, error) {
	if appInstanceID == "" {
		return nil, errors.New("app instance snapshot list: app instance id must not be empty")
	}
	raw, err := c.do(ctx, "list", http.MethodGet, "/v1/apps/instances/id/{id}/snapshots",
		map[string]string{"id": appInstanceID}, nil, nil)
	if err != nil {
		return nil, err
	}
	var l struct {
		Snapshots []*Snapshot `json:"snapshots"`
	}
	if len(raw) > 0 {
		if err := json.Unmarshal(raw, &l); err != nil {
			return nil, fmt.Errorf("app instance snapshot list: decoding response: %w", err)
		}
	}
	return l.Snapshots, nil
}

func (c *Client) State(ctx context.Context, snapshotID string) (*State, error) {
	if snapshotID == "" {
		return nil, errors.New("app instance snapshot state: id must not be empty")
	}
	raw, err := c.do(ctx, "state", http.MethodGet, "/v1/apps/instances/snapshot/id/{id}/state",
		map[string]string{"id": snapshotID}, nil, nil)
	if err != nil {
		return nil, err
	}
	var s State
	if len(raw) > 0 {
		if err := json.Unmarshal(raw, &s); err != nil {
			return nil, fmt.Errorf("app instance snapshot state: decoding response: %w", err)
		}
	}
	return &s, nil
}

func (c *Client) Activate(ctx context.Context, appInstanceID, name string, force bool) error {
	if appInstanceID == "" || name == "" {
		return errors.New("app instance snapshot activate: app instance id and name must not be empty")
	}
	var query map[string]string
	if force {
		query = map[string]string{"force": strconv.FormatBool(force)}
	}
	_, err := c.do(ctx, "activate", http.MethodPut, "/v1/apps/instances/id/{id}/snapshot/name/{name}/activate",
		map[string]string{"id": appInstanceID, "name": name}, query, nil)
	return err
}

func (c *Client) Delete(ctx context.Context, appInstanceID, name string) error {
	if appInstanceID == "" || name == "" {
		return errors.New("app instance snapshot delete: app instance id and name must not be empty")
	}
	_, err := c.do(ctx, "delete", http.MethodDelete, "/v1/apps/instances/id/{id}/snapshot/name/{name}",
		map[string]string{"id": appInstanceID, "name": name}, nil, nil)
	if errors.Is(err, ErrNotFound) {
		return nil
	}
	return err
}

func (c *Client) DeviceRawStatus(ctx context.Context, deviceID string) (string, error) {
	if deviceID == "" {
		return "", errors.New("edge node status: device id must not be empty")
	}
	raw, err := c.do(ctx, "device status", http.MethodGet, "/v1/devices/id/{id}/status",
		map[string]string{"id": deviceID}, nil, nil)
	if err != nil {
		return "", err
	}
	var s struct {
		RawStatus string `json:"rawStatus"`
	}
	if len(raw) > 0 {
		if err := json.Unmarshal(raw, &s); err != nil {
			return "", fmt.Errorf("edge node status: decoding response: %w", err)
		}
	}
	return s.RawStatus, nil
}

func decodeSnapshot(op string, raw []byte) (*Snapshot, error) {
	var s Snapshot
	if err := json.Unmarshal(raw, &s); err != nil {
		return nil, fmt.Errorf("app instance snapshot %s: decoding response: %w", op, err)
	}
	if s.ID == "" && s.Name == "" {
		return nil, ErrNotFound
	}
	return &s, nil
}

// do sends one request and returns the raw body of a 2xx response. A 404 is
// ErrNotFound; any other non-2xx status is an *APIError.
func (c *Client) do(
	ctx context.Context,
	op, method, path string,
	pathParams, query map[string]string,
	body interface{},
) ([]byte, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	res, err := c.transport.Submit(&runtime.ClientOperation{
		ID:                 "AppInstanceSnapshot_" + op,
		Method:             method,
		PathPattern:        path,
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params: runtime.ClientRequestWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
			for k, v := range pathParams {
				if err := r.SetPathParam(k, v); err != nil {
					return err
				}
			}
			for k, v := range query {
				if err := r.SetQueryParam(k, v); err != nil {
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
			if resp.Code() == http.StatusNotFound {
				return nil, ErrNotFound
			}
			if resp.Code() < 200 || resp.Code() > 299 {
				return nil, newAPIError(op, resp.Code(), raw)
			}
			return raw, nil
		}),
		Context: ctx,
	})
	if err != nil {
		// The provider's retrying HTTP client retries GET 404s and, once it
		// gives up, returns a transport error instead of the response.
		if !errors.Is(err, ErrNotFound) && strings.Contains(err.Error(), "404 Not Found") {
			return nil, ErrNotFound
		}
		return nil, err
	}
	raw, _ := res.([]byte)
	return raw, nil
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
	raw, _ := io.ReadAll(io.LimitReader(b, 4<<20))
	return raw
}

func truncate(s string) string {
	const max = 1024
	if len(s) > max {
		return s[:max] + "…"
	}
	return s
}
