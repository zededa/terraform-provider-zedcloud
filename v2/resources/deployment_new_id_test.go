package resources

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// ENG-3118: an in-place update of a deployment creates a new version with a
// new id and revision. The plan must show them as unknown, or outputs and
// dependants keep the old id until a second apply.

const depNewIDProject = "0b6f4b1e-3f6c-4c36-9d9f-4a4a3b3c1f11"

func depPlan(t *testing.T, tag, title string) *terraform.InstanceDiff {
	t.Helper()
	state := &terraform.InstanceState{
		ID: "8e9c48c7-0514-458a-b414-7ab2de1e9199",
		Attributes: map[string]string{
			"id": "8e9c48c7-0514-458a-b414-7ab2de1e9199", "name": "dep", "title": "dep",
			"project_id": depNewIDProject, "deployment_tag": "v1",
		},
	}
	cfg := terraform.NewResourceConfigRaw(map[string]interface{}{
		"name": "dep", "title": title, "project_id": depNewIDProject, "deployment_tag": tag,
	})
	diff, err := DeploymentResource().Diff(context.Background(), state, cfg, nil)
	if err != nil {
		t.Fatalf("plan failed: %s", err)
	}
	return diff
}

func TestDeploymentDiff_UpdatePlansNewID(t *testing.T) {
	diff := depPlan(t, "v2", "changed")
	if diff == nil {
		t.Fatal("expected a diff")
	}
	if a := diff.Attributes["id"]; a == nil || !a.NewComputed {
		t.Errorf("id must be planned as unknown on an update, got %#v", a)
	}
	if a := diff.Attributes["revision.#"]; a == nil || !a.NewComputed {
		t.Errorf("revision must be planned as unknown on an update, got %#v", a)
	}
}

func TestDeploymentDiff_TagOnlyChangePlansNewID(t *testing.T) {
	diff := depPlan(t, "v2", "dep")
	if a := diff.Attributes["id"]; a == nil || !a.NewComputed {
		t.Errorf("a deployment_tag change creates a version too, got %#v", a)
	}
}

func TestDeploymentDiff_NoChangeKeepsID(t *testing.T) {
	diff := depPlan(t, "v1", "dep")
	if diff != nil {
		if a := diff.Attributes["id"]; a != nil {
			t.Errorf("id must not be touched without a change (it would be a perpetual diff), got %#v", a)
		}
	}
}
