package schemas

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

// The SDK hands nested blocks to *FromMap as []interface{} of maps. This used
// to panic with "interface {} is []interface {}, not []*models.PCRValue".
func TestPCRTemplateModelFromMap(t *testing.T) {
	m := map[string]interface{}{
		"eve_version":      "17.0.0-lts-kvm-amd64",
		"firmware_version": "fw-1",
		"id":               "",
		"name":             "tmpl",
		"p_c_r_values": []interface{}{
			map[string]interface{}{"index": 0, "type": "PCR_TYPE_HASH", "value": "aa"},
			map[string]interface{}{"index": 14, "type": "PCR_TYPE_HASH", "value": "bb"},
		},
	}

	got := PCRTemplateModelFromMap(m)
	if *got.EveVersion != "17.0.0-lts-kvm-amd64" || got.FirmwareVersion != "fw-1" || got.Name != "tmpl" {
		t.Fatalf("unexpected template fields: %+v", got)
	}
	if len(got.PCRValues) != 2 {
		t.Fatalf("expected 2 PCR values, got %d", len(got.PCRValues))
	}
	v := got.PCRValues[1]
	if *v.Index != 14 || *v.Type != models.PCRTypePCRTYPEHASH || *v.Value != "bb" {
		t.Fatalf("unexpected PCR value: index=%d type=%v value=%s", *v.Index, *v.Type, *v.Value)
	}
}

// Round trip through the zedcloud_model schema, as in create and read: the SDK
// builds the config, SysModelModel converts it, and the setter writes it back.
func TestSysModelPCRTemplatesRoundTrip(t *testing.T) {
	tmpl := map[string]interface{}{
		"eve_version":      "17.0.0-lts-kvm-amd64",
		"firmware_version": "*",
		"name":             "17.0.0-lts-kvm-amd64",
		"p_c_r_values": []interface{}{
			map[string]interface{}{"index": 0, "type": "PCR_TYPE_HASH", "value": "aa"},
			map[string]interface{}{"index": 14, "type": "PCR_TYPE_HASH", "value": "bb"},
		},
	}
	d := schema.TestResourceDataRaw(t, SysModelSchema(), map[string]interface{}{
		"name":            "m",
		"p_c_r_templates": []interface{}{tmpl},
	})

	model := SysModelModel(d)
	if len(model.PCRTemplates) != 1 || len(model.PCRTemplates[0].PCRValues) != 2 {
		t.Fatalf("unexpected templates: %+v", model.PCRTemplates)
	}
	if v := model.PCRTemplates[0].PCRValues[1]; *v.Index != 14 || *v.Type != models.PCRTypePCRTYPEHASH {
		t.Fatalf("unexpected PCR value: index=%d type=%v", *v.Index, *v.Type)
	}

	read := schema.TestResourceDataRaw(t, SysModelSchema(), map[string]interface{}{})
	SetSysModelResourceData(read, model)
	if got := read.Get("p_c_r_templates.0.p_c_r_values.1.type"); got != "PCR_TYPE_HASH" {
		t.Fatalf("type in state is %#v", got)
	}
	if got := read.Get("p_c_r_templates.0.p_c_r_values.1.index"); got != 14 {
		t.Fatalf("index in state is %#v", got)
	}
	if got := read.Get("p_c_r_templates.0.eve_version"); got != "17.0.0-lts-kvm-amd64" {
		t.Fatalf("eve_version in state is %#v", got)
	}
}
