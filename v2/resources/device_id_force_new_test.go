package resources

import (
	"context"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// ENG-3121: device_id of a network instance or application instance cannot be
// changed in place. The controller rejects it for a network instance and
// ignores it for an application instance, so a change must plan a replacement.

const (
	devIDOld = "129e9cc9-a4db-4f0b-824b-1eb4114a5a91"
	devIDNew = "f2429fac-cc2d-4895-981e-767418ab9649"
)

type differ interface {
	Diff(context.Context, *terraform.InstanceState, *terraform.ResourceConfig, interface{}) (*terraform.InstanceDiff, error)
}

func TestDeviceID_ChangeForcesReplacement(t *testing.T) {
	resources := map[string]differ{
		"zedcloud_network_instance":     NetworkInstanceResource(),
		"zedcloud_application_instance": ApplicationInstanceResource(),
	}
	cases := []struct {
		name     string
		cfg      map[string]interface{}
		replaces bool
	}{
		{"device_id changed", map[string]interface{}{"device_id": devIDNew, "title": "t"}, true},
		{"device_id unchanged, title changed", map[string]interface{}{"device_id": devIDOld, "title": "t2"}, false},
		{"device_id not configured", map[string]interface{}{"title": "t"}, false},
	}

	for resName, r := range resources {
		for _, c := range cases {
			t.Run(resName+"/"+c.name, func(t *testing.T) {
				state := &terraform.InstanceState{ID: "inst-1", Attributes: map[string]string{
					"id": "inst-1", "name": "inst", "title": "t", "device_id": devIDOld,
				}}
				cfg := map[string]interface{}{"name": "inst"}
				for k, v := range c.cfg {
					cfg[k] = v
				}
				diff, err := r.Diff(context.Background(), state, terraform.NewResourceConfigRaw(cfg), nil)
				if err != nil {
					t.Fatalf("diff: %v", err)
				}
				if got := diff != nil && diff.RequiresNew(); got != c.replaces {
					t.Errorf("requires replacement = %v, want %v (diff: %v)", got, c.replaces, diff)
				}
			})
		}
	}
}

// The controller refuses to move a network instance and says so in
// error[].details with an empty httpStatusMsg. The generated update client had
// no case for 400, decoded the body into the wrong type, and the user saw
// "&{Code:0 Details:[] Message:}".
func TestNetworkInstanceUpdate_BadRequestReasonReachesUser(t *testing.T) {
	const reason = "Network instance edge node change is not supported"
	m := newNodeMeta(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == http.MethodGet {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"operationType":"OPS_TYPE_UPDATE","objectKind":"NetInst","httpStatusCode":400,"httpStatusMsg":"","error":[{"ec":"FunctionUnsupported","details":"` + reason + `"}]}`))
	}))

	r := NetworkInstanceResource()
	d := r.TestResourceData()
	d.SetId("ni-1")
	_ = d.Set("id", "ni-1")
	_ = d.Set("name", "ni")
	_ = d.Set("title", "ni")

	got := joinDiags(UpdateNetworkInstance(context.Background(), d, m))
	if !strings.Contains(got, reason) || strings.Contains(got, "Message:") {
		t.Errorf("the controller's reason must reach the user, got: %s", got)
	}
}
