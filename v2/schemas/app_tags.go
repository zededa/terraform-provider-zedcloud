package schemas

import (
	"fmt"
	"regexp"
	"sort"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

// Edge app tag limits, mirrored from zedcloud libs/zutils/validate.go
// (ValidateTag). The controller enforces them again; checking at plan time
// turns an apply-time 400 into a plan error that points at the bad entry.
const (
	AppTagLimit       = 32
	AppTagKeyMinLen   = 3
	AppTagKeyMaxLen   = 512
	AppTagValueMinLen = 3
	AppTagValueMaxLen = 256
)

var (
	// Keys may not contain < > % & ? / \ or quotes/backtick.
	appTagKeyRegexp = regexp.MustCompile("^[^<%&?/\\\\>'\"`]+$")
	// Values may not contain quotes, backtick or backslash.
	appTagValueInvalidRegexp = regexp.MustCompile("['\"`\\\\]")
)

// ValidateAppTags validates zedcloud_application.tags against the controller's
// rules for Edge app tags (NFR-96).
func ValidateAppTags(v interface{}, path cty.Path) diag.Diagnostics {
	m, ok := v.(map[string]interface{})
	if !ok || len(m) == 0 {
		return nil
	}

	var diags diag.Diagnostics
	if len(m) > AppTagLimit {
		diags = append(diags, diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Too many tags",
			Detail:        fmt.Sprintf("%d tags set, the controller allows at most %d.", len(m), AppTagLimit),
			AttributePath: path,
		})
	}

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		s, _ := m[k].(string)
		// A value not known until apply (e.g. from another resource): only the key can be checked.
		unknown := s == unknownValue
		if err := validateAppTag(k, s, unknown); err != "" {
			diags = append(diags, diag.Diagnostic{
				Severity:      diag.Error,
				Summary:       "Invalid tag",
				Detail:        fmt.Sprintf("tags[%q]: %s", k, err),
				AttributePath: path.IndexString(k),
			})
		}
	}
	return diags
}

// mapStringFromInterface converts a TypeMap value (map[string]interface{}) into
// map[string]string. It returns nil for a nil or empty map, so the field is
// omitted from the request body.
func mapStringFromInterface(v interface{}) map[string]string {
	m, ok := v.(map[string]interface{})
	if !ok || len(m) == 0 {
		return nil
	}
	out := make(map[string]string, len(m))
	for k, val := range m {
		s, _ := val.(string)
		out[k] = s
	}
	return out
}

// validateAppTag returns a description of what is wrong with one entry, or "".
// skipValue is set when the value is not known at plan time.
func validateAppTag(key, value string, skipValue bool) string {
	if len(key) < AppTagKeyMinLen || len(key) > AppTagKeyMaxLen {
		return fmt.Sprintf("key must be %d-%d characters long", AppTagKeyMinLen, AppTagKeyMaxLen)
	}
	if !appTagKeyRegexp.MatchString(key) {
		return "key must not contain any of < > % & ? / \\ ' \" `"
	}
	if skipValue {
		return ""
	}
	if len(value) < AppTagValueMinLen || len(value) > AppTagValueMaxLen {
		return fmt.Sprintf("value must be %d-%d characters long", AppTagValueMinLen, AppTagValueMaxLen)
	}
	if appTagValueInvalidRegexp.MatchString(value) {
		return "value must not contain any of ' \" ` \\"
	}
	return ""
}
