package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// Enterprise attribute keys the console reads to white-label the UI. These are the
// only keys the enterprise API accepts: it validates every attribute key against an
// allowlist and rejects the whole create/update request on any other key.
const (
	whiteLabelPrimaryColorKey   = "$ztag.entp.zui.ux.color.primary"
	whiteLabelSecondaryColorKey = "$ztag.entp.zui.ux.color.secondary"
	whiteLabelLogoKey           = "$ztag.entp.zui.ux.logo"
	whiteLabelProductNameKey    = "$ztag.entp.zui.ux.product.name"
)

// whiteLabelAttributeKeys maps each white_labeling block field to its attribute key.
var whiteLabelAttributeKeys = map[string]string{
	"primary_color":   whiteLabelPrimaryColorKey,
	"secondary_color": whiteLabelSecondaryColorKey,
	"logo_url":        whiteLabelLogoKey,
	"product_name":    whiteLabelProductNameKey,
}

// EnterpriseWhiteLabelingAttributes converts a white_labeling block into the
// enterprise attribute entries the console consumes. Empty fields are left out of
// the map: an absent key is how a white-label value gets cleared, since the API
// replaces the whole attribute map on every update.
func EnterpriseWhiteLabelingAttributes(v interface{}) map[string]string {
	attributes := map[string]string{}

	blocks, isList := v.([]interface{})
	if !isList || len(blocks) == 0 || blocks[0] == nil {
		return attributes
	}
	block, isMap := blocks[0].(map[string]interface{})
	if !isMap {
		return attributes
	}

	for field, key := range whiteLabelAttributeKeys {
		if value, _ := block[field].(string); value != "" {
			attributes[key] = value
		}
	}
	return attributes
}

// SetEnterpriseWhiteLabelingSubResourceData projects the white-label entries of an
// enterprise attribute map back into a white_labeling block. It returns nil when the
// enterprise carries none of them, so the block stays absent from state instead of
// showing up as an all-empty block.
func SetEnterpriseWhiteLabelingSubResourceData(attributes map[string]string) (d []*map[string]interface{}) {
	properties := make(map[string]interface{})
	isWhiteLabeled := false
	for field, key := range whiteLabelAttributeKeys {
		value := attributes[key]
		properties[field] = value
		if value != "" {
			isWhiteLabeled = true
		}
	}
	if !isWhiteLabeled {
		return nil
	}
	return append(d, &properties)
}

// isWhiteLabelAttributeKey reports whether an attribute key is owned by the
// white_labeling block.
func isWhiteLabelAttributeKey(key string) bool {
	for _, whiteLabelKey := range whiteLabelAttributeKeys {
		if key == whiteLabelKey {
			return true
		}
	}
	return false
}

// hasWhiteLabelAttributeKey reports whether an attribute map carries any white-label
// entry.
func hasWhiteLabelAttributeKey(attributes map[string]string) bool {
	for key := range attributes {
		if isWhiteLabelAttributeKey(key) {
			return true
		}
	}
	return false
}

// nonWhiteLabelAttributes returns the attribute entries that are not owned by the
// white_labeling block, so they keep round-tripping through the raw attributes map.
func nonWhiteLabelAttributes(attributes map[string]string) map[string]string {
	rest := map[string]string{}
	for key, value := range attributes {
		if !isWhiteLabelAttributeKey(key) {
			rest[key] = value
		}
	}
	return rest
}

func EnterpriseWhiteLabelingSchema() map[string]*schema.Schema {
	// The API stores these as enterprise tag values, which it requires to be between
	// 3 and 256 characters. The cap in particular is easy to trip over: it rules out
	// inlining an image, so the logo has to be a URL.
	return map[string]*schema.Schema{
		"primary_color": {
			Description:  `Primary theme color of the console, as a CSS color value, e.g. "#0A2540"`,
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringLenBetween(3, 256),
		},

		"secondary_color": {
			Description:  `Secondary theme color of the console, as a CSS color value, e.g. "#00B3A4"`,
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringLenBetween(3, 256),
		},

		"logo_url": {
			Description:  `URL of the logo shown in the console. Must be at most 256 characters, so the image cannot be inlined`,
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringLenBetween(3, 256),
		},

		"product_name": {
			Description:  `Product name shown in the console in place of the default one`,
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringLenBetween(3, 256),
		},
	}
}

func GetEnterpriseWhiteLabelingPropertyFields() (t []string) {
	return []string{
		"primary_color",
		"secondary_color",
		"logo_url",
		"product_name",
	}
}
