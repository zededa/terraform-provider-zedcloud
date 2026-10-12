package resources

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

// httpOptions are the provider-level HTTP knobs (UE-177 §4.1).
type httpOptions struct {
	// requestTimeout bounds a single HTTP attempt.
	requestTimeout time.Duration
	// maxRetries is the number of retries after the first attempt.
	maxRetries int
	// retryWaitMax caps the backoff between attempts.
	retryWaitMax time.Duration
}

func defaultHTTPOptions() httpOptions {
	return httpOptions{
		requestTimeout: 60 * time.Second,
		maxRetries:     4,
		retryWaitMax:   30 * time.Second,
	}
}

// callBudget is the deadline the provider gives an API call that does not
// carry a Terraform context: every attempt plus every backoff. Resources that
// pass their ctx (params.WithContext) get Terraform's timeouts instead.
func (o httpOptions) callBudget() time.Duration {
	return time.Duration(o.maxRetries+1) * (o.requestTimeout + o.retryWaitMax)
}

func getRetryClient(opts httpOptions) *retryablehttp.Client {
	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = opts.maxRetries
	retryClient.RetryWaitMax = opts.retryWaitMax
	if retryClient.RetryWaitMin > retryClient.RetryWaitMax {
		retryClient.RetryWaitMin = retryClient.RetryWaitMax
	}
	skipTlsVerification := envVarIsEnabled("TF_INSECURE_SKIP_TLS_VERIFY")
	if skipTlsVerification {
		insecureTransport := retryClient.HTTPClient.Transport.(*http.Transport).Clone()
		insecureTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		retryClient.HTTPClient = &http.Client{
			Transport: insecureTransport,
		}
	}
	// Bound each attempt, so one hung request cannot eat a resource's whole
	// Terraform timeout.
	retryClient.HTTPClient.Timeout = opts.requestTimeout
	retryClient.CheckRetry = checkRetry
	retryClient.ErrorHandler = giveUpWithLastResponse
	return retryClient
}

// giveUpWithLastResponse runs when retries are exhausted. If there was an
// HTTP response (a 404 or 5xx that kept coming back), it is returned as is, so
// the generated readers decode the controller's error and its status code
// instead of a bare "giving up after N attempt(s)" (UE-177 §4.5). Transport
// errors are still returned as errors.
func giveUpWithLastResponse(resp *http.Response, err error, attempts int) (*http.Response, error) {
	if resp != nil {
		log.Printf("[DEBUG] HTTP %s: giving up after %d attempt(s)", resp.Status, attempts)
		return resp, nil
	}
	if err == nil {
		return nil, fmt.Errorf("giving up after %d attempt(s)", attempts)
	}
	return nil, fmt.Errorf("giving up after %d attempt(s): %w", attempts, err)
}

// checkRetry decides whether retryablehttp retries a request.
//
//   - GET: the default policy (connection errors, 429, 5xx), plus 404 to absorb
//     read-after-write eventual consistency (CI-723).
//   - Any other method: our API is not idempotent, so a request that may have
//     reached the server is never sent again. That includes a client timeout:
//     the controller often completes the work after the client gave up
//     (ENG-3015), and a second POST would create a duplicate. Only errors that
//     prove the request never left (DNS, connection refused, TLS handshake)
//     are retried (UE-177 §4.1.4).
func checkRetry(ctx context.Context, resp *http.Response, err error) (bool, error) {
	if ctx.Err() != nil {
		return false, ctx.Err()
	}

	if resp == nil {
		if method := methodOf(err); method != "" && method != http.MethodGet {
			if requestNeverSent(err) {
				return retryablehttp.DefaultRetryPolicy(ctx, resp, err)
			}
			return false, nil
		}
		return retryablehttp.DefaultRetryPolicy(ctx, resp, err)
	}

	if resp.Request == nil {
		// Be conservative: without the request we cannot tell the method.
		if err != nil {
			return false, nil
		}
		return retryablehttp.DefaultRetryPolicy(ctx, resp, err)
	}

	if resp.Request.Method != http.MethodGet {
		return false, nil
	}

	// On GET, retry 404 to absorb read-after-write eventual consistency.
	// For non-GET methods we never retry 404: on DELETE it is the success
	// case, and on POST/PUT it indicates a missing parent that retries
	// won't resolve. See CI-723.
	if resp.StatusCode == http.StatusNotFound {
		return true, fmt.Errorf("unexpected HTTP status %s", resp.Status)
	}

	return retryablehttp.DefaultRetryPolicy(ctx, resp, err)
}

// methodOf returns the HTTP method of a failed request. net/http reports
// transport errors as *url.Error whose Op is the method with only its first
// letter upper-cased ("Post", "Get").
func methodOf(err error) string {
	var ue *url.Error
	if errors.As(err, &ue) {
		return strings.ToUpper(ue.Op)
	}
	return ""
}

// requestNeverSent reports whether err proves that the request did not reach
// the server, so sending it again cannot duplicate work.
func requestNeverSent(err error) bool {
	if err == nil {
		return false
	}
	var dnsErr *net.DNSError
	if errors.As(err, &dnsErr) {
		return true
	}
	var opErr *net.OpError
	if errors.As(err, &opErr) && opErr.Op == "dial" {
		return true
	}
	var recErr tls.RecordHeaderError
	if errors.As(err, &recErr) {
		return true
	}
	var unknownAuth x509.UnknownAuthorityError
	if errors.As(err, &unknownAuth) {
		return false // retrying will not fix a certificate problem
	}
	return strings.Contains(err.Error(), "tls: handshake")
}

// maxErrorDetails limits how much of a non-JSON error body is reported.
const maxErrorDetails = 512

var (
	htmlTagRE    = regexp.MustCompile(`(?s)<(script|style)[^>]*>.*?</(script|style)>|<[^>]+>`)
	whitespaceRE = regexp.MustCompile(`\s+`)
)

// rewriteNonJSONError replaces a non-JSON error body, such as a load
// balancer's or Cloudflare's HTML 502/504 page, with a ZsrvResponse JSON body
// that carries the HTTP status and the page text. Without it, the generated
// readers fail with "invalid character '<' looking for beginning of value"
// and the status is lost (UE-177 §4.5).
//
// It runs outside retryablehttp, so the retry policy still sees the original
// response.
func rewriteNonJSONError(resp *http.Response) {
	if resp == nil || resp.StatusCode < 400 || resp.Body == nil {
		return
	}
	raw, err := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
	_ = resp.Body.Close()
	if err != nil {
		raw = nil
	}

	ct := strings.ToLower(resp.Header.Get("Content-Type"))
	trimmed := bytes.TrimSpace(raw)
	if (ct == "" || strings.Contains(ct, "json")) && len(trimmed) > 0 && trimmed[0] == '{' && json.Valid(trimmed) {
		resp.Body = io.NopCloser(bytes.NewReader(raw))
		return
	}

	details := whitespaceRE.ReplaceAllString(htmlTagRE.ReplaceAllString(string(raw), " "), " ")
	details = strings.TrimSpace(details)
	if len(details) > maxErrorDetails {
		details = details[:maxErrorDetails] + "…"
	}
	if details == "" {
		details = "the response body was empty or not JSON"
	}

	synth := map[string]interface{}{
		"httpStatusCode": resp.StatusCode,
		"httpStatusMsg":  resp.Status,
		"error":          []map[string]string{{"details": details}},
	}
	body, _ := json.Marshal(synth)
	log.Printf("[DEBUG] HTTP %s: replaced non-JSON error body (%s) with a ZsrvResponse", resp.Status, ct)

	resp.Body = io.NopCloser(bytes.NewReader(body))
	resp.ContentLength = int64(len(body))
	resp.Header.Set("Content-Type", "application/json")
	resp.Header.Del("Content-Length")
}
