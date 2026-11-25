package resources

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// ENG-3118: an in-place update of a deployment creates a new version with a
// new id and revision. id cannot be planned as unknown (the plugin SDK then
// forces replacement, which destroys every version), so version_id and
// revision are planned as unknown, for outputs and dependants.

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

func TestDeploymentDiff_UpdatePlansNewVersionID(t *testing.T) {
	diff := depPlan(t, "v2", "changed")
	if diff == nil {
		t.Fatal("expected a diff")
	}
	if a := diff.Attributes["version_id"]; a == nil || !a.NewComputed {
		t.Errorf("version_id must be planned as unknown on an update, got %#v", a)
	}
	if a := diff.Attributes["revision.#"]; a == nil || !a.NewComputed {
		t.Errorf("revision must be planned as unknown on an update, got %#v", a)
	}
	assertNoReplace(t, diff)
}

func TestDeploymentDiff_TagOnlyChangePlansNewVersionID(t *testing.T) {
	diff := depPlan(t, "v2", "dep")
	if a := diff.Attributes["version_id"]; a == nil || !a.NewComputed {
		t.Errorf("a deployment_tag change creates a version too, got %#v", a)
	}
	assertNoReplace(t, diff)
}

func TestDeploymentDiff_NoChangeLeavesVersionID(t *testing.T) {
	diff := depPlan(t, "v1", "dep")
	if diff != nil {
		if a := diff.Attributes["version_id"]; a != nil {
			t.Errorf("version_id must not be touched without a change (it would be a perpetual diff), got %#v", a)
		}
	}
}

// An unknown id makes the plugin SDK force replacement, and destroy removes
// every deployment version. The plan must never touch id.
func assertNoReplace(t *testing.T, diff *terraform.InstanceDiff) {
	t.Helper()
	if diff.RequiresNew() {
		t.Error("an update must not force replacement")
	}
	if a := diff.Attributes["id"]; a != nil {
		t.Errorf("id must not be part of the diff, got %#v", a)
	}
}
