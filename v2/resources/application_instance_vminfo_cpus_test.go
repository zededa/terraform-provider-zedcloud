package resources

import (
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// ENG-3123: the controller takes the CPU count of an application instance from
// the application bundle and ignores vminfo.cpus, so a configured value that
// differs from the bundle's must not produce a permanent diff.

func vminfoCfg(cpus int) map[string]interface{} {
	return withCfg(map[string]interface{}{
		"vminfo": []interface{}{map[string]interface{}{"cpus": cpus, "vnc": false}},
	})
}

func seedWithBundleCPUs(t *testing.T, cpus int) (*fakeController, *terraform.InstanceState) {
	t.Helper()
	fc := newFakeController()
	meta := newFakeMeta(t, fc, false)
	state := seedInstance(t, fc, meta, map[string]interface{}{
		"vminfo": map[string]interface{}{"cpus": cpus, "vnc": false},
	})
	if got := state.Attributes["vminfo.0.cpus"]; got != "1" {
		t.Fatalf("seeded state vminfo.0.cpus = %q, want the controller's 1", got)
	}
	return fc, state
}

func TestApplicationInstance_VminfoCpusChangeOnExistingInstanceIsNotPlanned(t *testing.T) {
	fc, state := seedWithBundleCPUs(t, 1)
	meta := newFakeMeta(t, fc, false)

	diff, err := plan(t, meta, state, vminfoCfg(3))
	if err != nil {
		t.Fatalf("plan: %v", err)
	}
	if diff != nil {
		if a, ok := diff.Attributes["vminfo.0.cpus"]; ok {
			t.Fatalf("vminfo.0.cpus must not be planned for an existing instance, got %+v", a)
		}
	}
}

func TestApplicationInstance_VminfoCpusOnNewInstanceIsStillSent(t *testing.T) {
	meta := newFakeMeta(t, newFakeController(), false)

	diff, err := plan(t, meta, nil, vminfoCfg(2))
	if err != nil {
		t.Fatalf("plan: %v", err)
	}
	a, ok := diff.Attributes["vminfo.0.cpus"]
	if !ok || a.New != "2" {
		t.Fatalf("a new instance must keep the configured cpus in the plan, got %+v (present %v)", a, ok)
	}
}

func TestApplicationInstance_VminfoCpusIsDeprecated(t *testing.T) {
	diags := ApplicationInstanceResource().Validate(terraform.NewResourceConfigRaw(vminfoCfg(2)))
	for _, d := range diags {
		if strings.Contains(d.Summary+d.Detail, "application bundle manifest") {
			return
		}
	}
	t.Fatalf("setting vminfo.cpus must warn that the bundle decides, got %v", diags)
}
