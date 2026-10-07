package schemas

import (
	"fmt"
	"regexp"
	"sort"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

// AppLogoKey is the only key the zedcloud UI writes into manifest.desc.logo
// (zededa-services zedui-dev features/marketplace/services/edge-apps.ts).
const AppLogoKey = "logo"

// unknownValue is the SDK's placeholder for a value not known until apply,
// such as zedcloud_artifact.x.id on first create (hcl2shim.UnknownVariableValue,
// which lives in an internal package).
const unknownValue = "74D93920-ED26-11E3-AC10-0800200C9A66"

var logoURLRegexp = regexp.MustCompile(`(?i)^(https?:|data:)`)

// ValidateAppLogo checks manifest.desc.logo against what the UI can display:
//
//   - error: a value that is a URL. The UI treats every value as an artifact id
//     and asks the controller for /artifacts/id/<value>/url, so a URL logo
//     never renders.
//   - warning: more than one entry, or a single entry whose key is not "logo".
//     The UI shows one logo: the "logo" key if present, else the first entry.
func ValidateAppLogo(v interface{}, path cty.Path) diag.Diagnostics {
	m, ok := v.(map[string]interface{})
	if !ok || len(m) == 0 {
		return nil
	}

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var diags diag.Diagnostics
	for _, k := range keys {
		s, _ := m[k].(string)
		if s == "" || s == unknownValue {
			continue
		}
		if logoURLRegexp.MatchString(s) {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "App logo must be an artifact id, not a URL",
				Detail: fmt.Sprintf("logo[%q] = %q. The zedcloud UI only displays logos uploaded as artifacts. "+
					"Upload the image with a zedcloud_artifact resource and set logo = { logo = zedcloud_artifact.<name>.id }.", k, s),
				AttributePath: path.IndexString(k),
			})
		}
	}

	if len(m) > 1 {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "App logo has more than one entry",
			Detail: fmt.Sprintf("The zedcloud UI displays a single logo: the %q entry if present, otherwise the "+
				"first one. Entries: %v.", AppLogoKey, keys),
			AttributePath: path,
		})
	} else if keys[0] != AppLogoKey {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  fmt.Sprintf("App logo key should be %q", AppLogoKey),
			Detail: fmt.Sprintf("The zedcloud UI writes and prefers the key %q; got %q. "+
				"Use logo = { logo = <artifact id> }.", AppLogoKey, keys[0]),
			AttributePath: path.IndexString(keys[0]),
		})
	}
	return diags
}
