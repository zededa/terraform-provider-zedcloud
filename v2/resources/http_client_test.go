package resources

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"
)

func urlErr(method string, err error) error {
	// net/http reports transport errors with the method as Op, e.g. "Post".
	op := method[:1] + strings.ToLower(method[1:])
	return &url.Error{Op: op, URL: "https://zedcontrol.example/api/v1/apps", Err: err}
}

type timeoutErr struct{}

func (timeoutErr) Error() string   { return "Client.Timeout exceeded while awaiting headers" }
func (timeoutErr) Timeout() bool   { return true }
func (timeoutErr) Temporary() bool { return true }

func TestCheckRetry(t *testing.T) {
	dialErr := &net.OpError{Op: "dial", Net: "tcp", Err: errors.New("connection refused")}
	readErr := &net.OpError{Op: "read", Net: "tcp", Err: errors.New("connection reset by peer")}

	resp := func(method string, status int) *http.Response {
		return &http.Response{
			StatusCode: status,
			Status:     http.StatusText(status),
			Request:    &http.Request{Method: method},
		}
	}

	cases := []struct {
		name  string
		resp  *http.Response
		err   error
		retry bool
	}{
		{"POST client timeout is not retried", nil, urlErr(http.MethodPost, timeoutErr{}), false},
		{"PUT connection reset is not retried", nil, urlErr(http.MethodPut, readErr), false},
		{"DELETE client timeout is not retried", nil, urlErr(http.MethodDelete, timeoutErr{}), false},
		{"POST dial error is retried", nil, urlErr(http.MethodPost, dialErr), true},
		{"POST DNS error is retried", nil, urlErr(http.MethodPost, &net.DNSError{Err: "no such host", Name: "x"}), true},
		{"GET client timeout is retried", nil, urlErr(http.MethodGet, timeoutErr{}), true},
		{"GET 404 is retried", resp(http.MethodGet, http.StatusNotFound), nil, true},
		{"GET 503 is retried", resp(http.MethodGet, http.StatusServiceUnavailable), nil, true},
		{"GET 200 is not retried", resp(http.MethodGet, http.StatusOK), nil, false},
		{"POST 503 is not retried", resp(http.MethodPost, http.StatusServiceUnavailable), nil, false},
		{"DELETE 404 is not retried", resp(http.MethodDelete, http.StatusNotFound), nil, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, _ := checkRetry(context.Background(), tc.resp, tc.err)
			if got != tc.retry {
				t.Fatalf("retry = %v, want %v", got, tc.retry)
			}
		})
	}
}

func TestCheckRetry_ContextDone(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	retry, err := checkRetry(ctx, nil, urlErr(http.MethodGet, timeoutErr{}))
	if retry || err == nil {
		t.Fatalf("retry = %v, err = %v; want no retry and the context error", retry, err)
	}
}

func TestHTTPOptions_CallBudget(t *testing.T) {
	o := httpOptions{requestTimeout: 60 * time.Second, maxRetries: 4, retryWaitMax: 30 * time.Second}
	if got, want := o.callBudget(), 450*time.Second; got != want {
		t.Fatalf("callBudget = %s, want %s", got, want)
	}
	rc := getRetryClient(o)
	if rc.RetryMax != 4 || rc.RetryWaitMax != 30*time.Second || rc.HTTPClient.Timeout != 60*time.Second {
		t.Fatalf("retry client = max %d, wait %s, timeout %s", rc.RetryMax, rc.RetryWaitMax, rc.HTTPClient.Timeout)
	}
}

func newResp(status int, contentType, body string) *http.Response {
	h := http.Header{}
	if contentType != "" {
		h.Set("Content-Type", contentType)
	}
	return &http.Response{
		StatusCode: status,
		Status:     http.StatusText(status),
		Header:     h,
		Body:       io.NopCloser(strings.NewReader(body)),
	}
}

func TestRewriteNonJSONError_HTMLGatewayPage(t *testing.T) {
	r := newResp(http.StatusGatewayTimeout, "text/html",
		`<html><head><style>body{}</style><title>504</title></head><body><h1>504 Gateway Time-out</h1>
		<p>cloudflare</p></body></html>`)
	r.Status = "504 Gateway Time-out"

	rewriteNonJSONError(r)

	if ct := r.Header.Get("Content-Type"); ct != "application/json" {
		t.Fatalf("content-type = %q", ct)
	}
	var got struct {
		HTTPStatusCode int    `json:"httpStatusCode"`
		HTTPStatusMsg  string `json:"httpStatusMsg"`
		Error          []struct {
			Details string `json:"details"`
		} `json:"error"`
	}
	raw, _ := io.ReadAll(r.Body)
	if err := json.Unmarshal(raw, &got); err != nil {
		t.Fatalf("body is not JSON: %v: %s", err, raw)
	}
	if got.HTTPStatusCode != 504 || got.HTTPStatusMsg != "504 Gateway Time-out" {
		t.Errorf("status = %d %q", got.HTTPStatusCode, got.HTTPStatusMsg)
	}
	if len(got.Error) != 1 || !strings.Contains(got.Error[0].Details, "504 Gateway Time-out cloudflare") ||
		strings.Contains(got.Error[0].Details, "<") || strings.Contains(got.Error[0].Details, "body{}") {
		t.Errorf("details = %+v", got.Error)
	}
}

func TestRewriteNonJSONError_KeepsJSON(t *testing.T) {
	body := `{"httpStatusCode":409,"error":[{"details":"exists"}]}`
	r := newResp(http.StatusConflict, "application/json", body)
	rewriteNonJSONError(r)
	raw, _ := io.ReadAll(r.Body)
	if string(raw) != body {
		t.Fatalf("body changed: %s", raw)
	}
}

func TestRewriteNonJSONError_JSONWithoutContentType(t *testing.T) {
	body := `{"error":[{"details":"bad"}]}`
	r := newResp(http.StatusBadRequest, "", body)
	rewriteNonJSONError(r)
	raw, _ := io.ReadAll(r.Body)
	if string(raw) != body {
		t.Fatalf("body changed: %s", raw)
	}
}

func TestRewriteNonJSONError_IgnoresSuccess(t *testing.T) {
	r := newResp(http.StatusOK, "text/plain", "ok")
	rewriteNonJSONError(r)
	raw, _ := io.ReadAll(r.Body)
	if string(raw) != "ok" {
		t.Fatalf("body changed: %s", raw)
	}
}

func TestRewriteNonJSONError_EmptyBody(t *testing.T) {
	r := newResp(http.StatusBadGateway, "", "")
	rewriteNonJSONError(r)
	raw, _ := io.ReadAll(r.Body)
	if !json.Valid(raw) || !strings.Contains(string(raw), "502") {
		t.Fatalf("body = %s", raw)
	}
}
