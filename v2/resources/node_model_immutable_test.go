package resources

import (
	"context"
	"strings"
	"testing"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// CI-1026: the controller accepts a different modelId in a node update, never
// applies it and reports no error, so a changed model_id used to "apply" with
// no error and come back as the same diff on every later plan. The plan must
// reject the change instead.

const (
	testNodeID    = "7da54ab0-dcf2-403d-8f0f-115c2695ed23"
	testModelA    = "b4e7f3cf-e720-4f33-83ba-7cb8a28b2a6d"
	testModelB    = "afcbc2c5-ff51-49e5-8956-cd274ef566ae"
	unknownMarker = "74D93920-ED26-11E3-AC10-0800200C9A66" // the SDK's "known after apply" value
)

func nodeStateWithModel(model string) *terraform.InstanceState {
	return &terraform.InstanceState{
		ID: testNodeID,
		Attributes: map[string]string{
			"id":       testNodeID,
			"name":     "ci1026-node",
			"model_id": model,
		},
		RawConfig: cty.ObjectVal(map[string]cty.Value{
			"name": cty.StringVal("ci1026-node"),
		}),
	}
}

func diffNodeModel(t *testing.T, state *terraform.InstanceState, model string) error {
	t.Helper()
	cfg := terraform.NewResourceConfigRaw(map[string]interface{}{
		"name":     "ci1026-node",
		"model_id": model,
	})
	_, err := NodeResource().Diff(context.Background(), state, cfg, nil)
	return err
}

func TestNodeDiff_ModelIDChange_Rejected(t *testing.T) {
	err := diffNodeModel(t, nodeStateWithModel(testModelA), testModelB)
	if err == nil {
		t.Fatal("changing model_id on an existing node must fail at plan time, got no error")
	}
	for _, want := range []string{testModelA, testModelB, "ci1026-node", "offboard"} {
		if !strings.Contains(err.Error(), want) {
			t.Errorf("error should mention %q, got: %s", want, err)
		}
	}
}

func TestNodeDiff_ModelIDUnchanged_NoError(t *testing.T) {
	if err := diffNodeModel(t, nodeStateWithModel(testModelA), testModelA); err != nil {
		t.Fatalf("an unchanged model_id must not fail: %s", err)
	}
}

func TestNodeDiff_ModelIDOnCreate_NoError(t *testing.T) {
	// No prior state: a new node may be created with any model.
	if err := diffNodeModel(t, nil, testModelB); err != nil {
		t.Fatalf("creating a node with any model_id must not fail: %s", err)
	}
}

func TestNodeDiff_ModelIDEmptyInState_NoError(t *testing.T) {
	// A node with no recorded model (model-less) has nothing to protect.
	if err := diffNodeModel(t, nodeStateWithModel(""), testModelB); err != nil {
		t.Fatalf("a node with no model in state must not fail: %s", err)
	}
}

func TestNodeDiff_ModelIDUnknownAtPlan_PassesPlan(t *testing.T) {
	// model_id refers to something created in the same apply, so it is not
	// known at plan time; UpdateNode rejects it before sending (nodeModelDiffers).
	if err := diffNodeModel(t, nodeStateWithModel(testModelA), unknownMarker); err != nil {
		t.Fatalf("an unknown model_id cannot be judged at plan time and must pass: %s", err)
	}
}

func TestNodeDiff_OtherAttributeChange_NoModelError(t *testing.T) {
	cfg := terraform.NewResourceConfigRaw(map[string]interface{}{
		"name":     "ci1026-node",
		"model_id": testModelA,
		"title":    "renamed",
	})
	if _, err := NodeResource().Diff(context.Background(), nodeStateWithModel(testModelA), cfg, nil); err != nil {
		t.Fatalf("changing another attribute must not fail: %s", err)
	}
}

func TestNodeModelDiffers(t *testing.T) {
	a, b, empty := testModelA, testModelB, ""
	cases := []struct {
		name string
		x, y *string
		want bool
	}{
		{"different ids", &a, &b, true},
		{"same id", &a, &a, false},
		{"first nil", nil, &a, false},
		{"second nil", &a, nil, false},
		{"first empty", &empty, &a, false},
		{"second empty", &a, &empty, false},
		{"both nil", nil, nil, false},
	}
	for _, c := range cases {
		if got := nodeModelDiffers(c.x, c.y); got != c.want {
			t.Errorf("%s: nodeModelDiffers = %v, want %v", c.name, got, c.want)
		}
	}
}
