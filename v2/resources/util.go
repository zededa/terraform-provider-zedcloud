package resources

import (
	"log"
	"os"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

// HTLogger is a logger which satisfies the go-openapi/runtime/logger.Logger
// interface using go stdlib "log". It appends a prefix to each message to make
// them compatible with what TF expects for legacy logging 
// (see https://developer.hashicorp.com/terraform/plugin/log/writing#legacy-logging).
type HTLogger struct{}

// Printf just prints a log message. The TF "logging level" is hardcoded to
// "DEBUG" since it's expected that this will only be used in the
// go-openapi/runtime to log HTTP transport debugging info.
func (l HTLogger) Printf(format string, args ...interface{}) {
	log.Printf("[DEBUG] HTTP transport: "+format, args...)
}

// Debugf prints a log message the same way as Printf does. The TF "logging
// level" is hardcoded to "DEBUG" since it's expected that this will only
// be used in the go-openapi/runtime to log HTTP transport debugging info.
func (l HTLogger) Debugf(format string, args ...interface{}) {
	log.Printf("[DEBUG] HTTP transport: "+format, args...)
}

// envVarIsEnabled checks if an environment variable is set to a value of
// `yes` or `true` (case insensitive).
func envVarIsEnabled(key string) bool {
	val := strings.TrimSpace(os.Getenv(key))

	return len(val) > 0 && (strings.EqualFold(val, "true") || strings.EqualFold(val, "yes"))
}

// ZsrvResponder interface can be satistied by any API reponse struct that
// contains a models.ZsrvResponse and thus we can use that to extract additional
// information about the response (for example any error reasons send by the
// Zedcloud API).
type ZsrvResponder interface {
	GetPayload() *models.ZsrvResponse
	String() string
}

// ZsrvResponderToDiags checks if x satisfies the `ZsrvResponder` interface
// and maps any error reasons from the response to TF `Diagnostic`s.
func ZsrvResponderToDiags(intf interface{}) (diag.Diagnostics, bool) {
	diags := diag.Diagnostics{}

	responder, ok := intf.(ZsrvResponder)
	if !ok {
		return diags, false
	}

	payload := responder.GetPayload()
	if payload != nil {
		if payload.HTTPStatusCode != 0 {
			diags = append(diags, diag.Errorf("Zedcloud API call returned HTTP status code: %d", payload.HTTPStatusCode)...)
		}

		if len(payload.HTTPStatusMsg) > 0 {
			diags = append(diags, diag.Errorf("%s", payload.HTTPStatusMsg)...)
		}

		for _, w := range payload.Error {
			if w != nil && len(w.Details) > 0 {
				diags = append(diags, diag.Errorf("%s", w.Details)...)
			}
		}
	}

	if len(diags) == 0 {
		diags = append(diags, diag.Errorf("%s", responder.String())...)
	}

	return diags, true
}
