// Copyright (c) Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package resources

import (
	"context"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/zededa/terraform-provider-zedcloud/v2/schemas"
	testhelper "github.com/zededa/terraform-provider-zedcloud/v2/testing"
)

// ===========================================================================
// CI-709 regression coverage.
//
// Create a network, volume or application instance with only an
// `edge_node_cluster` and no `device_id`, and the controller resolves a
// designated node and persists the choice
// (the controller's network-instance create path, the app-instance create path). The provider
// declared `device_id` as `Optional` and NOT `Computed`, and PUTs the whole
// model on update with no per-field change detection -- so refresh wrote the
// server's value into state, the config had none, Terraform planned to remove
// it, and the follow-up apply 409'd.
//
// `app_type` was the same shape and worse: it carried
// `Default: "APP_TYPE_UNSPECIFIED"`, so the config always asserted a value and
// overwrote whatever the cluster flow had set -- the
// `~ app_type = "APP_TYPE_VM" -> "APP_TYPE_UNSPECIFIED"` line in the ticket.
//
// PR #234 (44f74c4a, merged) fixed six fields across three resources, and
// removed that Default because SDKv2 rejects Default together with Computed.
//
// *** WHY THESE TESTS EXIST EVEN THOUGH #234's CI WAS GREEN. ***
//
// CI could not have caught a regression in any of this:
//
//   - no fixture anywhere creates a cluster-scoped instance --
//     `grep -l edge_node_cluster v2/resources/testdata/*/*.tf` finds none, and
//     application_instance/create.tf pins `cluster_id = ""`
//   - test/e2e, which the e2e workflow applies, forms no cluster at all
//   - EVERY app-instance fixture sets `app_type` explicitly, so the removed
//     Default is invisible to the suite
//   - TestApplicationInstance_CreateCompose, the one app-instance test that
//     might have noticed, begins `if os.Getenv("CI") != "" { t.Skip(...) }`
//
// The end-to-end layer is test/e2e-cluster/ci709_repro.tf
// (TF_VAR_repro_ci709=true), which creates all three on a real cluster and
// asserts the plan stays clean.
// ===========================================================================

// serverAssignedFields are the fields PR #234 made Optional + Computed,
// grouped by the schema that carries them.
var serverAssignedFields = []struct {
	resource string
	schema   func() map[string]*schema.Schema
	fields   []string
}{
	{"zedcloud_network_instance", schemas.NetworkInstance, []string{"device_id", "cluster_id"}},
	{"zedcloud_volume_instance", schemas.VolumeInstance, []string{"device_id"}},
	{"zedcloud_application_instance", schemas.ApplicationInstance, []string{"device_id", "cluster_id", "app_type"}},
}

// TestClusterScopedSchemas_ServerFieldsAreComputed pins all six fields #234
// changed.
//
// `Optional + Computed` is the right shape here, unlike the CI-834 fields on
// zedcloud_edgenode which are Computed-only: a user legitimately MAY pin
// `device_id` to place an instance on a specific node, and may set `app_type`.
// The regression to guard against is losing `Computed`, which makes the
// absence of the field from the configuration mean "remove it".
func TestClusterScopedSchemas_ServerFieldsAreComputed(t *testing.T) {
	for _, res := range serverAssignedFields {
		s := res.schema()
		for _, name := range res.fields {
			f, ok := s[name]
			if !ok {
				t.Errorf("%s: no %q attribute", res.resource, name)
				continue
			}
			if !f.Computed {
				t.Errorf("%s.%s: Computed is false. The controller assigns this field for a "+
					"cluster-scoped instance; without Computed, Terraform plans to remove it "+
					"and the follow-up apply 409s. See CI-709 / PR #234.", res.resource, name)
			}
			if !f.Optional {
				t.Errorf("%s.%s: Optional is false; a user may legitimately set this field",
					res.resource, name)
			}
			// Default and Computed are mutually exclusive in SDKv2 -- the
			// provider would fail InternalValidate. Asserting it here names the
			// constraint instead of leaving a future reader to rediscover why
			// app_type's Default had to go.
			if f.Default != nil {
				t.Errorf("%s.%s: has Default %#v alongside Computed; SDKv2 rejects that "+
					"combination. #234 removed exactly this from app_type.",
					res.resource, name, f.Default)
			}
		}
	}
}

// TestApplicationInstance_ClusterFieldsProduceNoDiff is the CI-709 equivalent
// of the CI-834 diff test: cluster-assigned values in state, a configuration
// that omits them, and no entry that would discard one.
//
// The negative control reproduces the ticket's exact line by restoring the old
// schema shape -- `Optional` with `Default: "APP_TYPE_UNSPECIFIED"` -- which
// is why this test can be trusted to fail if the fix is reverted.
func TestApplicationInstance_ClusterFieldsProduceNoDiff(t *testing.T) {
	const (
		deviceID  = "08d09c84-1f0e-4a0f-8b5e-2c5b6d9f1a33"
		clusterID = "f73f1bdb-b229-4900-b564-99c2adf69f18"
		appType   = "APP_TYPE_VM"
	)

	state := &terraform.InstanceState{
		ID: "3f2a9c11-7d4e-4b8a-9c21-5e6f7a8b9c01",
		Attributes: map[string]string{
			"id":     "3f2a9c11-7d4e-4b8a-9c21-5e6f7a8b9c01",
			"name":   "demo_app_0",
			"app_id": "a1b2c3d4-0000-4000-8000-000000000001",

			// What the controller assigned, as the provider's
			// SetAppInstResourceData writes it.
			"device_id":              deviceID,
			"cluster_id":             clusterID,
			"app_type":               appType,
			"edge_node_cluster.#":    "1",
			"edge_node_cluster.0.id": clusterID,
		},
	}

	// The configuration a user writes for a cluster-scoped instance: the
	// cluster, and none of the fields the controller owns.
	cfg := terraform.NewResourceConfigRaw(map[string]interface{}{
		"name":   "demo_app_0",
		"title":  "demo app 0",
		"app_id": "a1b2c3d4-0000-4000-8000-000000000001",
		"edge_node_cluster": []interface{}{
			map[string]interface{}{"id": clusterID},
		},
	})

	appFields := []string{"device_id", "cluster_id", "app_type"}

	t.Run("fixed schema discards nothing", func(t *testing.T) {
		got := overwrittenAttrs(t, schemas.ApplicationInstance(), state, cfg, appFields)
		if len(got) != 0 {
			t.Errorf("CI-709 regression: plan would discard %v on a cluster-scoped app "+
				"instance.\nThe follow-up apply then 409s. device_id, cluster_id and "+
				"app_type must be Optional + Computed; see PR #234.", got)
		}
	})

	// NEGATIVE CONTROL: the pre-#234 schema, field by field.
	//
	// app_type is restored WITH its Default, because that is what produced the
	// ticket's distinctive line -- not a removal to null but a rewrite to the
	// default:
	//
	//	~ app_type = "APP_TYPE_VM" -> "APP_TYPE_UNSPECIFIED"
	t.Run("negative control: the pre-#234 shape reintroduces the diff", func(t *testing.T) {
		broken := schemas.ApplicationInstance()
		for _, name := range appFields {
			s := *broken[name] // copy; leave the shared schema alone
			s.Computed = false
			s.Optional = true
			if name == "app_type" {
				s.Default = "APP_TYPE_UNSPECIFIED"
			}
			broken[name] = &s
		}

		got := overwrittenAttrs(t, broken, state, cfg, appFields)
		if len(got) == 0 {
			t.Fatal("negative control did not reproduce CI-709: restoring the pre-#234 " +
				"schema produced no discarding diff, so the assertion above proves " +
				"nothing. Fix this test before trusting it.")
		}
		t.Logf("negative control reproduced the diff on %v", got)
	})
}

// TestVolumeInstance_ClusterFieldsProduceNoDiff is CI-752's own shape.
//
// CI-709 was reported against a network instance and an application instance;
// CI-752 (https://zededa.atlassian.net/browse/CI-752, "Errors after creating
// persistent storage to cluster") is the SAME defect reported against a volume
// instance, with the ticket's plan output reading
//
//	~ resource "zedcloud_volume_instance" "ignition_vol1_persist" {
//	    - device_id = "19a0f34f-..." -> null
//
// The volume path is worth its own case rather than a row in
// serverAssignedFields: the server side differs from the network-instance and
// app-instance paths, and only `device_id` is at stake, with no `cluster_id`
// and no `app_type` alongside it to mask a regression.
//
// Server side, for the record. `bindVolInstData`
// (srvs/seine/volinstproc.go:383-397) resolves a node for a volume created with
// an `edgeNodeCluster` and no `deviceId` by calling
// `clusterproc.GetEdgeNodeForClusterObject`, which picks at RANDOM among the
// eligible non-tie-breaker nodes (clusterproc.go:2843), and volinstproc.go:435
// persists it. Note the difference from the network-instance path, which keeps
// the existing device on update (netinstproc.go:102-110): `volinstUpdate`
// (volinstproc.go:1726-1734) copies only title and description, so a PUT that
// carries an empty `deviceId` changes nothing server-side. That is why the
// pre-#234 diff never converged -- each apply "removed" a value the server
// then still reported on the next GET.
func TestVolumeInstance_ClusterFieldsProduceNoDiff(t *testing.T) {
	const (
		deviceID  = "19a0f34f-8d2e-4a1b-9c33-92052fb2c14c"
		clusterID = "0370bb66-5d53-426f-96da-1575cf5efba0"
	)

	state := &terraform.InstanceState{
		ID: "7d721b72-54ab-4729-aca4-3ed7f790735c",
		Attributes: map[string]string{
			"id":   "7d721b72-54ab-4729-aca4-3ed7f790735c",
			"name": "ignition-vol1-sjc-persist",

			// What the controller assigned, as SetVolumeInstanceResourceData
			// writes it back after refresh.
			"device_id":              deviceID,
			"edge_node_cluster.#":    "1",
			"edge_node_cluster.0.id": clusterID,
		},
	}

	// The configuration from the ticket, minus the fields the reporter did not
	// write: no device_id anywhere in it.
	cfg := terraform.NewResourceConfigRaw(map[string]interface{}{
		"name":       "ignition-vol1-sjc-persist",
		"title":      "ignition-vol1-portainer-persist",
		"accessmode": "VOLUME_INSTANCE_ACCESS_MODE_READWRITE",
		"cleartext":  false,
		"size_bytes": "200000000",
		"type":       "VOLUME_INSTANCE_TYPE_BLOCKSTORAGE",
		"edge_node_cluster": []interface{}{
			map[string]interface{}{"id": clusterID},
		},
	})

	volFields := []string{"device_id"}

	t.Run("fixed schema discards nothing", func(t *testing.T) {
		got := overwrittenAttrs(t, schemas.VolumeInstance(), state, cfg, volFields)
		if len(got) != 0 {
			t.Errorf("CI-752 regression: plan would discard %v on a cluster-scoped volume "+
				"instance. device_id must be Optional + Computed; see PR #234.", got)
		}
	})

	// NEGATIVE CONTROL. Without this the assertion above could pass for the
	// wrong reason -- e.g. a state/config pair that never produces a diff at
	// all -- and volume_instance has no Default anywhere, so a silent no-diff
	// is entirely plausible.
	t.Run("negative control: the pre-#234 shape reintroduces the diff", func(t *testing.T) {
		broken := schemas.VolumeInstance()
		s := *broken["device_id"] // copy; leave the shared schema alone
		s.Computed = false
		s.Optional = true
		broken["device_id"] = &s

		got := overwrittenAttrs(t, broken, state, cfg, volFields)
		if len(got) == 0 {
			t.Fatal("negative control did not reproduce CI-752: restoring the pre-#234 " +
				"schema produced no discarding diff, so the assertion above proves " +
				"nothing. Fix this test before trusting it.")
		}
		t.Logf("negative control reproduced the diff on %v", got)
	})
}

// TestClusterScopedInstances_RealNode asserts the same thing against instances
// that genuinely live on a cluster, using the values the controller assigned.
//
// Only a live object shows what the controller actually does: which fields it
// populates, and -- for the app instance created with NO `app_type` -- what it
// defaults that to now that the provider no longer sends
// "APP_TYPE_UNSPECIFIED". A synthetic test cannot answer either question.
//
// Requires (otherwise skipped):
//
//	ZEDCLOUD_ACC_REAL_NODE=1
//	ZEDCLOUD_TEST_CLUSTER_NI_NAME      / _VOL_NAME / _APPINST_NAME
//	TF_VAR_zedcloud_url / TF_VAR_zedcloud_token
//
// Names come from test/e2e-cluster applied with TF_VAR_repro_ci709=true:
//
//	tofu output -json ci709_application_instance | jq -r .name
func TestClusterScopedInstances_RealNode(t *testing.T) {
	if !realNodeEnabled() {
		t.Skipf("%s is not set; skipping test that requires a real cluster",
			testhelper.EnvRealNode)
	}

	cases := []struct {
		env    string
		kind   testhelper.InstanceKind
		schema func() map[string]*schema.Schema
		fields []string
	}{
		{testhelper.EnvClusterNIName, testhelper.KindNetworkInstance,
			schemas.NetworkInstance, []string{"device_id", "cluster_id"}},
		{testhelper.EnvClusterVolName, testhelper.KindVolumeInstance,
			schemas.VolumeInstance, []string{"device_id"}},
		{testhelper.EnvClusterAppInstName, testhelper.KindApplicationInstance,
			schemas.ApplicationInstance, []string{"device_id", "cluster_id", "app_type"}},
	}

	ran := 0
	for _, c := range cases {
		name := os.Getenv(c.env)
		if name == "" {
			t.Logf("%s is empty; skipping %s", c.env, c.kind)
			continue
		}
		ran++

		t.Run(string(c.kind), func(t *testing.T) {
			live, err := testhelper.GetClusterScopedInstance(c.kind, name)
			if err != nil {
				t.Fatalf("%s", err)
			}

			// Enabled-but-not-cluster-scoped is an operator mistake. Skipping
			// would leave a green test that asserts nothing.
			if !live.ClusterScoped() {
				t.Fatalf("%s %q has no controller-assigned device_id (deviceId=%q "+
					"edgeNodeCluster set=%t).\nThis test needs an instance created with "+
					"edge_node_cluster and NO device_id -- apply test/e2e-cluster with "+
					"TF_VAR_repro_ci709=true.",
					c.kind, name, live.DeviceID, live.EdgeNodeCluster != nil)
			}

			t.Logf("%s %q: device_id=%s cluster_id=%s app_type=%q",
				c.kind, name, live.DeviceID, live.ClusterID, live.AppType)

			attrs := map[string]string{
				"id":                     live.ID,
				"name":                   live.Name,
				"device_id":              live.DeviceID,
				"cluster_id":             live.ClusterID,
				"edge_node_cluster.#":    "1",
				"edge_node_cluster.0.id": live.EdgeNodeCluster.ID,
			}
			if live.AppType != "" {
				attrs["app_type"] = live.AppType
			}

			state := &terraform.InstanceState{ID: live.ID, Attributes: attrs}
			cfg := terraform.NewResourceConfigRaw(map[string]interface{}{
				"name": live.Name,
				"edge_node_cluster": []interface{}{
					map[string]interface{}{"id": live.EdgeNodeCluster.ID},
				},
			})

			if got := overwrittenAttrs(t, c.schema(), state, cfg, c.fields); len(got) != 0 {
				t.Errorf("CI-709 is live on %s %q: a plan would discard %v, and the "+
					"follow-up apply would 409.", c.kind, name, got)
			}
		})
	}

	if ran == 0 {
		t.Skipf("none of %s / %s / %s is set; nothing to check",
			testhelper.EnvClusterNIName, testhelper.EnvClusterVolName,
			testhelper.EnvClusterAppInstName)
	}
}

// realNodeEnabled mirrors the gate RealNode applies, without its node-name and
// online requirements -- CI-709 is about instances on a cluster, so there is no
// single node to resolve or wait for.
func realNodeEnabled() bool {
	switch os.Getenv(testhelper.EnvRealNode) {
	case "1", "true", "yes", "TRUE", "True", "YES", "Yes":
		return true
	}
	return false
}

// isOverwrite reports whether a diff entry would REPLACE a value the
// controller assigned, with anything at all.
//
// *** DELIBERATELY WIDER THAN isDrop, AND THE DIFFERENCE MATTERS. ***
//
// The CI-834 fields on zedcloud_edgenode have no Default, so the only way to
// lose a controller-assigned value there is a drop to empty -- isDrop is
// sufficient. CI-709's `app_type` carried
// `Default: "APP_TYPE_UNSPECIFIED"`, and a Default does not drop a value, it
// REWRITES it. That is the ticket's own line:
//
//	~ app_type = "APP_TYPE_VM" -> "APP_TYPE_UNSPECIFIED"
//
// The first version of this test reused isDrop, and its negative control duly
// reported only device_id and cluster_id -- silently missing app_type, the one
// field whose regression looks like a rewrite. So: any non-NewComputed change
// away from a non-empty prior value counts.
func isOverwrite(attr *terraform.ResourceAttrDiff) bool {
	if attr.NewComputed {
		return false
	}
	if attr.NewRemoved {
		return true
	}
	return attr.Old != "" && attr.New != attr.Old
}

// overwrittenAttrs returns the named attributes whose controller-assigned
// value the SDK's diff would replace.
//
// NewComputed is excluded: it means "(known after apply)", the Read returns
// the same value already in state, and the plan is empty. See isDrop in
// edgenode_cluster_test.go for the false failure that taught us to exclude it.
//
// The resource is built from the schema map alone rather than via the real
// *Resource, to avoid dragging CustomizeDiff hooks into a diff computed from a
// hand-built ResourceConfig.
func overwrittenAttrs(
	t *testing.T,
	resSchema map[string]*schema.Schema,
	state *terraform.InstanceState,
	cfg *terraform.ResourceConfig,
	fields []string,
) []string {
	t.Helper()

	r := &schema.Resource{Schema: resSchema}
	diff, err := r.Diff(context.Background(), state, cfg, nil)
	if err != nil {
		t.Fatalf("computing diff: %s", err)
	}
	if diff == nil {
		return nil
	}

	var got []string
	for _, f := range fields {
		attr, ok := diff.Attributes[f]
		if !ok || attr == nil || !isOverwrite(attr) {
			continue
		}
		got = append(got, describeAttrDiff(f, attr))
	}
	return got
}
