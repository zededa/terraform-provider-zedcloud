package schemas

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"regexp"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ArtifactMaxSize caps what the provider will upload. The controller has no
// limit today (ENG-3009 item 4) and reads every chunk into memory; the UI caps
// app logos at 5 MB and enterprise logos at 1 MB.
const ArtifactMaxSize = 10 * 1024 * 1024

// artifactNameRegexp is the controller's object-name rule (zutils.ValidateName,
// swagger Artifact.name): 3-256 characters.
var artifactNameRegexp = regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_.-]{2,255}$`)

// ArtifactContentHash is what state stores for content_base64, so a large
// inline payload is not kept in state verbatim.
func ArtifactContentHash(v interface{}) string {
	s, _ := v.(string)
	if s == "" {
		return ""
	}
	sum := sha256.Sum256([]byte(s))
	return "sha256:" + hex.EncodeToString(sum[:])
}

func ArtifactSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Description: "Artifact name. The controller generates the id as `<uuid>_<name>`. " +
				"3-256 characters matching `[a-zA-Z0-9][a-zA-Z0-9_.-]+`. The name is not the key used in an " +
				"Edge app's `manifest.desc.logo` map: that key is always `logo`.",
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
			ValidateDiagFunc: func(v interface{}, p cty.Path) diag.Diagnostics {
				s, _ := v.(string)
				if !artifactNameRegexp.MatchString(s) {
					return diag.Diagnostics{{
						Severity:      diag.Error,
						Summary:       "Invalid artifact name",
						Detail:        fmt.Sprintf("%q must be 3-256 characters matching [a-zA-Z0-9][a-zA-Z0-9_.-]+", s),
						AttributePath: p,
					}}
				}
				return nil
			},
		},

		"source": {
			Description: "Path to a local file to upload. Exactly one of `source` and `content_base64` must be set. " +
				"Terraform does not notice when the file's content changes unless `source_hash` is also set " +
				"(for example to `filesha256(...)`).",
			Type:         schema.TypeString,
			Optional:     true,
			ForceNew:     true,
			ExactlyOneOf: []string{"source", "content_base64"},
		},

		"content_base64": {
			Description: "Base64-encoded content to upload, for example `filebase64(\"logo.png\")`. " +
				"Exactly one of `source` and `content_base64` must be set. Only a SHA-256 of the value is stored in state.",
			Type:         schema.TypeString,
			Optional:     true,
			ForceNew:     true,
			ExactlyOneOf: []string{"source", "content_base64"},
			StateFunc:    ArtifactContentHash,
			ValidateDiagFunc: func(v interface{}, p cty.Path) diag.Diagnostics {
				s, _ := v.(string)
				if _, err := base64.StdEncoding.DecodeString(s); err != nil {
					return diag.Diagnostics{{
						Severity:      diag.Error,
						Summary:       "content_base64 is not valid base64",
						Detail:        err.Error(),
						AttributePath: p,
					}}
				}
				return nil
			},
		},

		"source_hash": {
			Description: "Any string that changes when the content of `source` changes, typically " +
				"`filesha256(...)`. Changing it uploads a new artifact (and so a new `id`).",
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"retain_on_destroy": {
			Description: "When true, destroying or replacing this resource only removes it from Terraform state and " +
				"leaves the artifact on the controller. Use it when the artifact may be referenced from outside " +
				"Terraform, for example by an app duplicated in the UI, which shares the original's logo artifact. " +
				"Default false.",
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},

		"size": {
			Description: "Number of bytes uploaded. The controller may store fewer or more: it strips metadata from " +
				"PNG and JPEG images, which re-encodes them.",
			Type:     schema.TypeInt,
			Computed: true,
		},
	}
}
