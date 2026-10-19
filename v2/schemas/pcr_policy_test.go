package schemas

import (
	"reflect"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func TestValidatePCRIndices(t *testing.T) {
	cases := []struct {
		name    string
		in      []int
		wantErr string
	}{
		{"recommended default", PCRDefaultIndices, ""},
		{"pcr 0 only", []int{0}, ""},
		{"empty", nil, "cannot be empty"},
		{"missing pcr 0", []int{1, 2}, "PCR 0"},
		{"volatile pcr 5", []int{0, 5}, "volatile"},
		{"out of range", []int{0, 16}, "between 0 and 15"},
		{"negative", []int{0, -1}, "between 0 and 15"},
		{"duplicate", []int{0, 1, 1}, "duplicate"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidatePCRIndices(tc.in)
			if tc.wantErr == "" {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				return
			}
			if err == nil || !strings.Contains(err.Error(), tc.wantErr) {
				t.Fatalf("error = %v, want it to contain %q", err, tc.wantErr)
			}
		})
	}
}

func TestValidatePCRIndexElem(t *testing.T) {
	for _, ok := range []int{0, 1, 4, 6, 15} {
		if _, es := validatePCRIndexElem(ok, "k"); len(es) != 0 {
			t.Errorf("%d: unexpected errors %v", ok, es)
		}
	}
	for _, bad := range []int{5, 16, -1} {
		if _, es := validatePCRIndexElem(bad, "k"); len(es) == 0 {
			t.Errorf("%d: expected an error", bad)
		}
	}
}

func TestNextPCRPolicyVersion(t *testing.T) {
	cases := []struct {
		name       string
		cur        []int64
		curVersion int64
		next       []int64
		want       int64
	}{
		{"new policy", nil, 0, []int64{0, 1}, 1},
		{"unchanged keeps version", []int64{0, 1, 2}, 3, []int64{0, 1, 2}, 3},
		{"reordered is unchanged", []int64{0, 2, 1}, 3, []int64{1, 2, 0}, 3},
		{"changed bumps", []int64{0, 1, 2}, 3, []int64{0, 1}, 4},
		{"changed after zcli bump", []int64{0, 7}, 9, []int64{0, 1, 7}, 10},
		{"unchanged but no version yet", []int64{0, 1}, 0, []int64{0, 1}, 1},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := NextPCRPolicyVersion(tc.cur, tc.curVersion, tc.next); got != tc.want {
				t.Fatalf("got %d, want %d", got, tc.want)
			}
		})
	}
}

func TestPCRIndicesFromInterfaceSorts(t *testing.T) {
	set := schema.NewSet(schema.HashInt, []interface{}{14, 0, 7, 3})
	got := PCRIndicesFromInterface(set)
	want := []int64{0, 3, 7, 14}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	if got := PCRIndicesFromInterface([]interface{}{2, 0}); !reflect.DeepEqual(got, []int64{0, 2}) {
		t.Fatalf("list: got %v", got)
	}
}

func TestNodePCRPolicyRoundTrip(t *testing.T) {
	if NodePCRPolicyModel(nil) != nil || NodePCRPolicyModel([]interface{}{}) != nil {
		t.Fatal("absent block must produce no policy")
	}
	block := []interface{}{map[string]interface{}{
		"pcr_indices":    schema.NewSet(schema.HashInt, []interface{}{7, 0, 4}),
		"policy_version": 2,
	}}
	p := NodePCRPolicyModel(block)
	if p == nil || p.Type == nil || *p.Type != models.PolicyTypePOLICYTYPEPCR {
		t.Fatalf("type not set: %+v", p)
	}
	if !reflect.DeepEqual(p.PcrPolicy.PcrIndices, []int64{0, 4, 7}) || p.PcrPolicy.PolicyVersion != 2 {
		t.Fatalf("unexpected policy %+v", p.PcrPolicy)
	}

	flat := FlattenNodePCRPolicy(p)
	m := flat[0].(map[string]interface{})
	if !reflect.DeepEqual(m["pcr_indices"], []interface{}{0, 4, 7}) || m["policy_version"] != 2 {
		t.Fatalf("unexpected flatten %v", m)
	}
	if FlattenNodePCRPolicy(nil) != nil || FlattenNodePCRPolicy(&models.Policy{}) != nil {
		t.Fatal("no policy must flatten to nil")
	}
}

func TestNodeSchemaPCRPolicyRoundTrip(t *testing.T) {
	d := schema.TestResourceDataRaw(t, Node(), map[string]interface{}{
		"pcr_policy": []interface{}{map[string]interface{}{
			"pcr_indices": []interface{}{0, 1, 2},
		}},
	})
	n := NodeModel(d)
	if n.PcrPolicy == nil || !reflect.DeepEqual(n.PcrPolicy.PcrPolicy.PcrIndices, []int64{0, 1, 2}) {
		t.Fatalf("NodeModel pcr policy = %+v", n.PcrPolicy)
	}
	n.PcrPolicy.PcrPolicy.PolicyVersion = 5
	SetNodeResourceData(d, n)
	if v := d.Get("pcr_policy.0.policy_version"); v != 5 {
		t.Fatalf("policy_version = %v", v)
	}
}

func TestProjectSchemaPCRPolicyRoundTrip(t *testing.T) {
	d := schema.TestResourceDataRaw(t, Project(), map[string]interface{}{
		"name":  "p1",
		"title": "p1",
		"type":  "TAG_TYPE_PROJECT",
		"pcr_policy": []interface{}{map[string]interface{}{
			"title": "pcr",
			"type":  "POLICY_TYPE_PCR",
			"pcr_policy": []interface{}{map[string]interface{}{
				"pcr_indices": []interface{}{13, 0, 14},
			}},
		}},
	})
	tag := TagModel(d)
	if tag.PcrPolicy == nil || tag.PcrPolicy.PcrPolicy == nil {
		t.Fatalf("TagModel pcr policy = %+v", tag.PcrPolicy)
	}
	if got := tag.PcrPolicy.PcrPolicy.PcrIndices; !reflect.DeepEqual(got, []int64{0, 13, 14}) {
		t.Fatalf("indices = %v", got)
	}
	if *tag.PcrPolicy.Type != models.PolicyTypePOLICYTYPEPCR {
		t.Fatalf("type = %v", *tag.PcrPolicy.Type)
	}

	tag.PcrPolicy.ID = "policy-id"
	tag.PcrPolicy.PcrPolicy.PolicyVersion = 3
	SetTagResourceData(d, tag)
	if v := d.Get("pcr_policy.0.pcr_policy.0.policy_version"); v != 3 {
		t.Fatalf("policy_version = %v", v)
	}
	if v := d.Get("pcr_policy.0.id"); v != "policy-id" {
		t.Fatalf("id = %v", v)
	}
}

func TestDevicePolicyPCRRoundTrip(t *testing.T) {
	m := map[string]interface{}{
		"policy_sub_type": "DEVICE_POLICY_TYPE_PCR",
		"pcr_policy": []interface{}{map[string]interface{}{
			"pcr_indices":    schema.NewSet(schema.HashInt, []interface{}{4, 0}),
			"policy_version": 0,
		}},
	}
	p := DevicePolicyModelFromMap(m)
	if p.PcrPolicy == nil || !reflect.DeepEqual(p.PcrPolicy.PcrIndices, []int64{0, 4}) {
		t.Fatalf("pcr policy = %+v", p.PcrPolicy)
	}
	p.PcrPolicy.PolicyVersion = 2
	flat := SetDevicePolicySubResourceData([]*models.DevicePolicy{p})
	pcr := (*flat[0])["pcr_policy"].([]*map[string]interface{})
	if (*pcr[0])["policy_version"] != 2 {
		t.Fatalf("flatten = %v", *pcr[0])
	}
}
