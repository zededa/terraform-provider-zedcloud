// Copyright (c) Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package testing

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

// Environment contract for tests that need a REAL, onboarded edge node.
//
// Most acceptance tests in this package fabricate a zedcloud_edgenode against a
// synthetic brand/model with an invented serial. Nothing ever checks in, so
// those tests cannot observe anything the device actually does -- an app
// instance is "created" but never runs.
//
// These variables point a test at a node that genuinely exists and is onboarded
// (see test/e2e/ for the Terraform config that creates one). When they are
// absent the dependent tests skip, so the default suite is unaffected.
const (
	// EnvRealNode gates the real-node tests. Set to "1"/"true" to enable.
	EnvRealNode = "ZEDCLOUD_ACC_REAL_NODE"

	// EnvRealNodeName is the name of the onboarded edge node to target. The
	// fixture resolves it through the zedcloud_edgenode data source rather
	// than taking an ID, so node lifecycle and test run stay decoupled.
	EnvRealNodeName = "ZEDCLOUD_TEST_NODE_NAME"

	// EnvTestSuffix is appended to every object name the fixture creates, so
	// concurrent runs and leftovers cannot collide. Required because the suite
	// otherwise uses fixed literal names and alpha is a shared enterprise.
	EnvTestSuffix = "ZEDCLOUD_TEST_SUFFIX"
)

// RealNodeInfo identifies the live edge node a test should run against.
type RealNodeInfo struct {
	// Name of the onboarded edge node.
	Name string
	// ID is the node's UUID, resolved from Name against the API.
	//
	// Why resolved here instead of via the zedcloud_edgenode data source: that
	// data source is built from zschema.Node(), the *resource* schema, so it
	// inherits every write-side Required field. `data "zedcloud_edgenode"` with
	// only a name fails with "The argument \"model_id\" is required", likewise
	// project_id, title and at least one interfaces block. It cannot currently
	// be used to look a node up. Resolving by name here keeps the ergonomics
	// the data source should have provided (callers pass a name, not a UUID)
	// without depending on that defect being fixed first.
	ID string
	// Suffix appended to created object names for run isolation.
	Suffix string
}

// RealNode returns the live-node parameters, skipping the test when real-node
// testing is not enabled. Call it first in any test that needs a device.
//
// It skips rather than fails on purpose: `go test ./...` with no lab access must
// stay green.
func RealNode(t *testing.T) RealNodeInfo {
	t.Helper()

	if !envEnabled(EnvRealNode) {
		t.Skipf("%s is not set; skipping test that requires a real onboarded edge node", EnvRealNode)
	}

	name := os.Getenv(EnvRealNodeName)
	if name == "" {
		// Enabled but unconfigured is a mistake, not a reason to skip
		// silently -- that is how a "passing" CI job ends up testing nothing.
		t.Fatalf("%s is set but %s is empty; it must name an onboarded edge node", EnvRealNode, EnvRealNodeName)
	}

	suffix := os.Getenv(EnvTestSuffix)
	if suffix == "" {
		suffix = fmt.Sprintf("_%d", time.Now().Unix())
		t.Logf("%s not set, generated run suffix %q", EnvTestSuffix, suffix)
	}

	// A node that is not online cannot run a workload, and the resulting
	// failure would look like an unrelated timeout deep in the test, so this
	// is checked up front where the message is actionable.
	//
	// But it is POLLED rather than sampled once, because "not online" is
	// routinely transient. An edge node reboots shortly after onboarding, so
	// the sequence in CI is:
	//
	//	tofu apply -> zedamigo_wait_until sees RUN_STATE_ONLINE -> apply ends
	//	  -> device reboots -> this test starts -> RUN_STATE_BOOTING
	//
	// A single reading therefore fails or passes depending on where in that
	// window the test happens to start. Observed both ways on consecutive CI
	// runs against an otherwise healthy node: run 34103767333 sailed past this
	// check, run 34106625131 died on it in 0.43s. Treating one non-ONLINE
	// sample as fatal is the bug; the node being briefly unavailable is not.
	id, runState, err := waitNodeOnline(t, name, nodeOnlineTimeout)
	if err != nil {
		t.Fatalf("%s", err)
	}

	t.Logf("real node %q resolved to %s (%s), run suffix %q", name, id, runState, suffix)
	return RealNodeInfo{Name: name, ID: id, Suffix: suffix}
}

// nodeOnlineTimeout bounds the wait for a node to report RUN_STATE_ONLINE.
//
// Sized against the measured cold path: a freshly installed node reaches
// ONLINE about 1m20s after boot, so a post-onboard reboot should clear well
// inside this. Long enough to absorb that, short enough that a genuinely dead
// node fails the job in minutes rather than at the 40m go test timeout.
const nodeOnlineTimeout = 4 * time.Minute

// waitNodeOnline resolves an edge node by name and polls until it reports
// RUN_STATE_ONLINE, returning its id and final run state.
//
// Lookup errors are retried alongside the state check: the same reboot that
// flips runState can also make the status endpoint briefly return nothing
// useful, and failing on the first such blip would defeat the point.
func waitNodeOnline(t *testing.T, name string, timeout time.Duration) (string, string, error) {
	t.Helper()

	const interval = 10 * time.Second

	deadline := time.Now().Add(timeout)
	var (
		lastID    string
		lastState string
		lastErr   error
	)

	for attempt := 1; ; attempt++ {
		id, runState, err := resolveNodeByName(name)
		if id != "" {
			lastID = id
		}
		switch {
		case err != nil:
			lastErr = err
			t.Logf("attempt %d: resolving edge node %q failed: %s", attempt, name, err)
		case runState == "RUN_STATE_ONLINE":
			if attempt > 1 {
				t.Logf("edge node %q reached RUN_STATE_ONLINE after %d attempt(s)", name, attempt)
			}
			return id, runState, nil
		default:
			lastState = runState
			lastErr = nil
			t.Logf("attempt %d: edge node %q (%s) is %s, waiting for RUN_STATE_ONLINE", attempt, name, id, runState)
		}

		if time.Now().After(deadline) {
			if lastErr != nil {
				return lastID, lastState, fmt.Errorf(
					"gave up after %s resolving edge node %q: %w", timeout, name, lastErr)
			}
			return lastID, lastState, fmt.Errorf(
				"edge node %q (%s) still %s after %s, never reached RUN_STATE_ONLINE; bring it online before running real-node tests",
				name, lastID, lastState, timeout)
		}
		time.Sleep(interval)
	}
}

// resolveNodeByName looks up an edge node's UUID and current run state.
func resolveNodeByName(name string) (id string, runState string, err error) {
	var cfg struct {
		ID string `json:"id"`
	}
	if err := apiGet(fmt.Sprintf("devices/name/%s", name), &cfg); err != nil {
		return "", "", err
	}
	if cfg.ID == "" {
		return "", "", fmt.Errorf("no id returned for node %q", name)
	}

	var st struct {
		RunState string `json:"runState"`
	}
	if err := apiGet(fmt.Sprintf("devices/id/%s/status", cfg.ID), &st); err != nil {
		return cfg.ID, "", err
	}
	return cfg.ID, st.RunState, nil
}

func envEnabled(key string) bool {
	switch strings.ToLower(os.Getenv(key)) {
	case "1", "true", "yes":
		return true
	}
	return false
}

// MustGetTestInputWithVars reads ./testdata/<path> and expands __TOKEN__
// placeholders from vars.
//
// Use it for fixtures needing tokens MustGetTestInput does not know about --
// __NODE_ID__ for the real-node fixture. vars is the complete set: nothing is
// filled in from the environment, so a caller wanting the suite-wide suffix
// passes testhelper.Suffix() explicitly (RealNode supplies its own, which is
// generated when ZEDCLOUD_TEST_SUFFIX is unset).
//
// The placeholder form is __TOKEN__ rather than ${TOKEN} because ${...} is HCL
// interpolation syntax and would break `terraform fmt` and `validate` on the
// fixture files.
func MustGetTestInputWithVars(t *testing.T, path string, vars map[string]string) string {
	t.Helper()

	return expandTokens(t, path, mustReadTestdata(t, path), vars)
}

func keysOf(m map[string]string) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	return out
}

// WaitForAppInstanceSwState polls an app instance's status until swState reaches
// want, and fails the test otherwise.
//
// This is the assertion that only a real node makes possible. The progression is
//
//	SW_STATE_INITIAL -> SW_STATE_RESOLVING_TAG -> SW_STATE_DOWNLOAD_IN_PROGRESS
//	  -> SW_STATE_INSTALLED -> SW_STATE_RUNNING
//
// so reaching SW_STATE_RUNNING means the controller pushed the config, the
// device pulled the image and actually started the workload.
//
// The generated client has no operation for the status endpoint (only GetByID /
// GetByName, which return config), so this talks to the REST API directly using
// the same environment the provider reads.
// It returns an error rather than calling t.Fatal so it composes with
// resource.TestCheckFunc; t is used only for progress logging (visible with
// `go test -v`), which matters because this can legitimately run for minutes.
func WaitForAppInstanceSwState(t *testing.T, appInstID string, want models.SWState, timeout time.Duration) error {
	t.Helper()

	const interval = 15 * time.Second

	deadline := time.Now().Add(timeout)
	var last string

	for attempt := 1; ; attempt++ {
		st, err := getAppInstanceStatus(appInstID)
		if err != nil {
			// Transient API errors are expected right after create, while the
			// object propagates. Keep going until the deadline.
			t.Logf("attempt %d: status fetch failed: %s", attempt, err)
		} else {
			sw := derefSwState(st.SwState)
			run := derefRunState(st.RunState)
			last = fmt.Sprintf("swState=%s runState=%s", sw, run)
			t.Logf("attempt %d: %s", attempt, last)

			if sw == string(want) {
				t.Logf("app instance %s reached %s after %d attempt(s)", appInstID, want, attempt)
				return nil
			}

			// Surface device-reported errors immediately rather than burning
			// the whole timeout on a workload that has already failed.
			if msgs := errDescriptions(st.ErrInfo); len(msgs) > 0 {
				return fmt.Errorf("app instance %s reported error(s) while waiting for %s: %s (last: %s)",
					appInstID, want, strings.Join(msgs, "; "), last)
			}
		}

		if time.Now().After(deadline) {
			return fmt.Errorf("timed out after %s waiting for app instance %s to reach %s (last: %s)",
				timeout, appInstID, want, last)
		}
		time.Sleep(interval)
	}
}

func errDescriptions(errs []*models.DeviceError) []string {
	var msgs []string
	for _, e := range errs {
		if e == nil || e.Description == nil || *e.Description == "" {
			continue
		}
		msgs = append(msgs, *e.Description)
	}
	return msgs
}

func getAppInstanceStatus(id string) (*models.AppInstStatusMsg, error) {
	var st models.AppInstStatusMsg
	if err := apiGet(fmt.Sprintf("apps/instances/id/%s/status", id), &st); err != nil {
		return nil, err
	}
	return &st, nil
}

// apiGet performs a GET against /api/v1/<path> and decodes the JSON body.
//
// The generated client covers only the config endpoints (GetByID / GetByName);
// there is no generated operation for the /status endpoints these helpers need,
// so they talk to the REST API directly, reading the same environment the
// provider itself reads.
func apiGet(path string, out interface{}) error {
	host := os.Getenv("TF_VAR_zedcloud_url")
	if host == "" {
		return fmt.Errorf("TF_VAR_zedcloud_url is not set")
	}
	token := os.Getenv("TF_VAR_zedcloud_token")
	if token == "" {
		return fmt.Errorf("TF_VAR_zedcloud_token is not set")
	}
	host = strings.TrimPrefix(strings.TrimPrefix(host, "https://"), "http://")
	host = strings.TrimSuffix(host, "/")

	url := fmt.Sprintf("https://%s/api/v1/%s", host, path)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{Timeout: 30 * time.Second}
	if envEnabled("TF_INSECURE_SKIP_TLS_VERIFY") {
		client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //nolint:gosec // opt-in, dev clusters only
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("GET %s: unexpected HTTP status %s: %s",
			path, resp.Status, truncate(string(body), 200))
	}

	if err := json.Unmarshal(body, out); err != nil {
		return fmt.Errorf("decoding response from %s: %w", path, err)
	}
	return nil
}

func derefSwState(s *models.SWState) string {
	if s == nil {
		return "<nil>"
	}
	return string(*s)
}

func derefRunState(s *models.RunState) string {
	if s == nil {
		return "<nil>"
	}
	return string(*s)
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "..."
}
