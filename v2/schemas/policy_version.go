package schemas

import (
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func PolicyVersionModel(d *schema.ResourceData) *models.PolicyVersion {
	acceptedAt, _ := d.Get("accepted_at").(strfmt.DateTime)
	acceptedBy, _ := d.Get("accepted_by").(string)
	effectiveFrom, _ := d.Get("effective_from").(strfmt.DateTime)
	policy, _ := d.Get("policy").(string)
	version, _ := d.Get("version").(string)
	return &models.PolicyVersion{
		AcceptedAt:    acceptedAt,
		AcceptedBy:    acceptedBy,
		EffectiveFrom: effectiveFrom,
		Policy:        policy,
		Version:       version,
	}
}

func PolicyVersionModelFromMap(m map[string]interface{}) *models.PolicyVersion {
	acceptedAt := m["accepted_at"].(strfmt.DateTime)
	acceptedBy := m["accepted_by"].(string)
	effectiveFrom := m["effective_from"].(strfmt.DateTime)
	policy := m["policy"].(string)
	version := m["version"].(string)
	return &models.PolicyVersion{
		AcceptedAt:    acceptedAt,
		AcceptedBy:    acceptedBy,
		EffectiveFrom: effectiveFrom,
		Policy:        policy,
		Version:       version,
	}
}

func SetPolicyVersionResourceData(d *schema.ResourceData, m *models.PolicyVersion) {
	d.Set("accepted_at", m.AcceptedAt)
	d.Set("accepted_by", m.AcceptedBy)
	d.Set("effective_from", m.EffectiveFrom)
	d.Set("policy", m.Policy)
	d.Set("version", m.Version)
}

func SetPolicyVersionSubResourceData(m []*models.PolicyVersion) (d []*map[string]interface{}) {
	for _, PolicyVersionModel := range m {
		if PolicyVersionModel != nil {
			properties := make(map[string]interface{})
			properties["accepted_at"] = PolicyVersionModel.AcceptedAt.String()
			properties["accepted_by"] = PolicyVersionModel.AcceptedBy
			properties["effective_from"] = PolicyVersionModel.EffectiveFrom.String()
			properties["policy"] = PolicyVersionModel.Policy
			properties["version"] = PolicyVersionModel.Version
			d = append(d, &properties)
		}
	}
	return
}

func PolicyVersionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"accepted_at": {
			Description:  ``,
			Type:         schema.TypeString,
			ValidateFunc: validation.IsRFC3339Time,
			Optional:     true,
		},

		"accepted_by": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"effective_from": {
			Description:  ``,
			Type:         schema.TypeString,
			ValidateFunc: validation.IsRFC3339Time,
			Optional:     true,
		},

		"policy": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"version": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetPolicyVersionPropertyFields() (t []string) {
	return []string{
		"accepted_at",
		"accepted_by",
		"effective_from",
		"policy",
		"version",
	}
}
