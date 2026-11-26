package resources

import "testing"

// ENG-3122: the controller keeps userData from create and ignores it on update,
// so changing user_data must replace the instance instead of planning an
// in-place update that never converges.

func TestApplicationInstance_UserDataChangeForcesReplacement(t *testing.T) {
	fc := newFakeController()
	meta := newFakeMeta(t, fc, false)
	state := seedInstance(t, fc, meta, map[string]interface{}{"userData": "QQ=="})
	if got := state.Attributes["user_data"]; got != "QQ==" {
		t.Fatalf("seeded user_data = %q, want QQ==", got)
	}

	same, err := plan(t, meta, state, withCfg(map[string]interface{}{"user_data": "QQ=="}))
	if err != nil {
		t.Fatalf("plan: %v", err)
	}
	if same != nil && !same.Empty() {
		t.Fatalf("unchanged user_data must plan no change, got %v", same)
	}

	diff, err := plan(t, meta, state, withCfg(map[string]interface{}{"user_data": "Qg=="}))
	if err != nil {
		t.Fatalf("plan: %v", err)
	}
	if diff == nil || !diff.RequiresNew() {
		t.Fatalf("changing user_data must require replacement, got %v", diff)
	}
}
