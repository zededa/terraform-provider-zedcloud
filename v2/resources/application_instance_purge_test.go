package resources

import (
	"context"
	"strings"
	"testing"
)

// Unit tests for purge { counter } on zedcloud_application_instance (ENG-3114).
// The controller ignores purge.counter in an update and keeps its own counter,
// which each PUT .../refresh/purge raises by one. The fake controller behaves
// the same way.

func purgeCfg(counter int) map[string]interface{} {
	return withCfg(map[string]interface{}{
		"purge": []interface{}{map[string]interface{}{"counter": counter}},
	})
}

func purgeCalls(fc *fakeController) int {
	return countPrefix(fc.requestLog(), "PUT /v1/apps/instances/id/"+snapInstID+"/refresh/purge")
}

func controllerPurgeCounter(fc *fakeController) float64 {
	fc.mu.Lock()
	defer fc.mu.Unlock()
	if p, ok := fc.instances["inst-s"]["purge"].(map[string]interface{}); ok {
		c, _ := p["counter"].(float64)
		return c
	}
	return 0
}

func TestApplicationInstancePurge_CounterIncreasePurgesOnce(t *testing.T) {
	for _, async := range []bool{false, true} {
		fc := newFakeController()
		meta := newFakeMeta(t, fc, async)
		state := seedInstance(t, fc, meta, nil)

		state = apply(t, meta, state, purgeCfg(1))
		if n := purgeCalls(fc); n != 1 {
			t.Fatalf("async=%v: counter 0 -> 1 must purge once, got %d calls:\n%s", async, n, fc.requestLog())
		}
		if got := controllerPurgeCounter(fc); got != 1 {
			t.Fatalf("async=%v: controller counter %v, want 1", async, got)
		}
		if got := state.Attributes["purge.0.counter"]; got != "1" {
			t.Fatalf("async=%v: state purge.0.counter = %q, want the configured 1", async, got)
		}
		for _, p := range fc.puts {
			if _, ok := p["purge"]; ok {
				t.Fatalf("async=%v: the update body must not carry a purge counter the controller keeps, got %v", async, p["purge"])
			}
		}

		// The plan converges although the controller counter is its own.
		if diff, err := plan(t, meta, state, purgeCfg(1)); err != nil || (diff != nil && !diff.Empty()) {
			t.Fatalf("async=%v: plan after the purge must be clean, got %v (err %v)", async, diff, err)
		}

		// A jump by more than one still purges once.
		state = apply(t, meta, state, purgeCfg(5))
		if n := purgeCalls(fc); n != 2 {
			t.Fatalf("async=%v: counter 1 -> 5 must purge once more, got %d calls in total", async, n)
		}
		if got := state.Attributes["purge.0.counter"]; got != "5" {
			t.Fatalf("async=%v: state purge.0.counter = %q, want 5", async, got)
		}
		if diff, err := plan(t, meta, state, purgeCfg(5)); err != nil || (diff != nil && !diff.Empty()) {
			t.Fatalf("async=%v: plan after the second purge must be clean, got %v (err %v)", async, diff, err)
		}
	}
}

func TestApplicationInstancePurge_NoPurgeWithoutAnIncrease(t *testing.T) {
	fc := newFakeController()
	meta := newFakeMeta(t, fc, false)
	state := seedInstance(t, fc, meta, nil)
	state = apply(t, meta, state, purgeCfg(3))
	if n := purgeCalls(fc); n != 1 {
		t.Fatalf("setup: want 1 purge call, got %d", n)
	}

	// Another change with the same counter does not purge.
	cfg := purgeCfg(3)
	cfg["title"] = "other title"
	state = apply(t, meta, state, cfg)
	if n := purgeCalls(fc); n != 1 {
		t.Fatalf("an unrelated update must not purge, got %d calls:\n%s", n, fc.requestLog())
	}

	// Lowering the counter does not purge either.
	state = apply(t, meta, state, purgeCfg(1))
	if n := purgeCalls(fc); n != 1 {
		t.Fatalf("lowering the counter must not purge, got %d calls", n)
	}
	if got := state.Attributes["purge.0.counter"]; got != "1" {
		t.Fatalf("state purge.0.counter = %q, want 1", got)
	}
}

func TestApplicationInstancePurge_FailedPurgeIsRetried(t *testing.T) {
	fc := newFakeController()
	meta := newFakeMeta(t, fc, false)
	state := seedInstance(t, fc, meta, nil)

	fc.mu.Lock()
	fc.failPurge = true
	fc.mu.Unlock()

	diff, err := plan(t, meta, state, purgeCfg(1))
	if err != nil || diff == nil || diff.Empty() {
		t.Fatalf("plan: %v (err %v)", diff, err)
	}
	newState, ds := ApplicationInstanceResource().Apply(context.Background(), state, diff, meta)
	if !ds.HasError() {
		t.Fatalf("a failed purge must fail the apply:\n%s", fc.requestLog())
	}
	reported := false
	for _, d := range ds {
		if strings.Contains(d.Summary, "purge failed") {
			reported = true
		}
	}
	if !reported {
		t.Fatalf("the controller's error must be reported, got %v", ds)
	}
	if newState != nil {
		if got := newState.Attributes["purge.0.counter"]; got == "1" {
			t.Fatalf("state must not claim the purge happened, got purge.0.counter = %q", got)
		}
	}

	// The next apply asks for the purge again.
	fc.mu.Lock()
	fc.failPurge = false
	fc.mu.Unlock()
	if newState == nil {
		newState = state
	}
	newState = apply(t, meta, newState, purgeCfg(1))
	if got := controllerPurgeCounter(fc); got != 1 {
		t.Fatalf("the retry must purge, controller counter %v:\n%s", got, fc.requestLog())
	}
	if got := newState.Attributes["purge.0.counter"]; got != "1" {
		t.Fatalf("state purge.0.counter = %q, want 1", got)
	}
}
