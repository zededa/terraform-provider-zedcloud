// Copyright (c) Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package resources

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/zededa/terraform-provider-zedcloud/v2/schemas"
	testhelper "github.com/zededa/terraform-provider-zedcloud/v2/testing"
)

// ===========================================================================
// CI-834 regression coverage.
//
// When an edge node joins an edge-node cluster the controller populates its
// `cluster_interface` and `edge_node_cluster` attributes by itself. No
// Terraform configuration declares them, so with both fields marked
// `Optional` the provider read the values back, found nothing in the
// configuration to match, and planned to REMOVE them:
//
//	~ resource "zedcloud_edgenode" "ENODE_001" {
//	    - cluster_interface = "eth1" -> null
//	    - edge_node_cluster {
//	        - cluster_prefix = "10.244.244.2/28" -> null
//	        ...
//
// Applying that plan tears down the cluster membership. Commit c932b75d
// ("fix: diffs for cluster_interface and edge_node_cluster fields(CI-834)",
// PR #199) fixed it by changing exactly two lines in v2/schemas/node.go,
// Optional -> Computed.
//
// That fix shipped in February 2026 and stayed UNTESTED for seven months:
// nothing anywhere under v2/ referenced either attribute from a test. This
// file is that missing coverage. It has three layers, deliberately:
//
//	TestEdgeNodeSchema_ClusterFieldsAreComputed
//	  pins the two schema lines. Cheap, runs everywhere, and is what actually
//	  fails first if someone regenerates the schema from swagger.
//
//	TestEdgeNode_ClusterFieldsProduceNoDiff
//	  exercises the SDK's diff machinery, which is where the bug lived. It
//	  carries a NEGATIVE CONTROL that reverts the two fields to Optional and
//	  asserts the diff comes back, so the test cannot pass vacuously.
//
//	TestEdgeNodeCluster_RealNode
//	  the same diff, but against cluster values read off a live clustered node
//	  rather than values this file invented. Gated on the lab.
//
// The end-to-end layer lives in test/e2e-cluster, which forms a real cluster
// and asserts `tofu plan` reports no changes.
// ===========================================================================

// clusterAttrPrefixes are the state keys CI-834 is about.
//
// Matched by prefix because `edge_node_cluster` is a TypeList of a nested
// resource, so it flattens to edge_node_cluster.#, edge_node_cluster.0.id,
// edge_node_cluster.0.seed_node_ip and so on -- a plain equality check on the
// block name would miss every one of them.
var clusterAttrPrefixes = []string{"cluster_interface", "edge_node_cluster"}

func isClusterAttr(key string) bool {
	for _, p := range clusterAttrPrefixes {
		if key == p || strings.HasPrefix(key, p+".") {
			return true
		}
	}
	return false
}

// TestEdgeNodeSchema_ClusterFieldsAreComputed pins the two lines c932b75d
// changed.
//
// Optional and Computed are not interchangeable here and the distinction is
// the whole bug: Optional says "the configuration may set this", which makes
// its absence from the configuration meaningful, so a value present only in
// state reads as "the user removed it". Computed says "the server owns this",
// and absence from the configuration means nothing at all.
//
// Optional+Computed would also suppress the diff, so it is accepted -- that is
// what the CI-709 sibling fix used for network_instance.device_id, where the
// user legitimately MAY set the field. What must never come back is Optional
// WITHOUT Computed.
func TestEdgeNodeSchema_ClusterFieldsAreComputed(t *testing.T) {
	node := schemas.Node()

	for _, name := range clusterAttrPrefixes {
		s, ok := node[name]
		if !ok {
			t.Fatalf("schemas.Node() has no %q attribute; CI-834 concerns exactly this field", name)
		}
		if !s.Computed {
			t.Errorf("%s: Computed is false. The controller populates this field when the node "+
				"joins a cluster; without Computed, Terraform plans to remove it and tearing "+
				"down cluster membership. See CI-834 / commit c932b75d.", name)
		}
		if s.Optional && !s.Computed {
			t.Errorf("%s: Optional without Computed -- this is precisely the CI-834 regression", name)
		}
		if s.Required {
			t.Errorf("%s: Required is true; a server-populated field cannot be required", name)
		}
	}
}

// TestEdgeNode_ClusterFieldsProduceNoDiff drives the SDK's own diff for the
// node schema with cluster values in state and nothing in the configuration,
// which is the exact shape of a post-cluster-create plan.
func TestEdgeNode_ClusterFieldsProduceNoDiff(t *testing.T) {
	state := clusteredNodeState(exampleClusterFields())

	// The configuration a user actually writes: no cluster_interface, no
	// edge_node_cluster. Everything else matches state so the only thing the
	// assertion can see is the fields under test.
	cfg := terraform.NewResourceConfigRaw(map[string]interface{}{
		"name":       "ENODE_001",
		"title":      "ENODE 001",
		"model_id":   "eeee7b73-9063-4c03-9b26-dabb81e84cca",
		"project_id": "b91b2c7d-bc40-4422-81b6-304886b712ba",
	})

	t.Run("fixed schema produces no cluster diff", func(t *testing.T) {
		got := clusterAttrsInDiff(t, schemas.Node(), state, cfg)
		if len(got) != 0 {
			t.Errorf("CI-834 regression: plan would change %v on a clustered node.\n"+
				"cluster_interface and edge_node_cluster must be Computed in "+
				"v2/schemas/node.go; see commit c932b75d.", got)
		}
	})

	// NEGATIVE CONTROL.
	//
	// Without this the test above could pass for the wrong reason -- a typo in
	// clusterAttrPrefixes, a diff that came back nil because the schema failed
	// to build, a future SDK that stops reporting removals here. Reverting the
	// two fields to Optional must make the diff reappear; if it does not, the
	// assertion above is not measuring anything.
	t.Run("negative control: Optional brings the diff back", func(t *testing.T) {
		broken := schemas.Node()
		for _, name := range clusterAttrPrefixes {
			s := *broken[name] // copy, so the shared schema map is untouched
			s.Computed = false
			s.Optional = true
			broken[name] = &s
		}

		got := clusterAttrsInDiff(t, broken, state, cfg)
		if len(got) == 0 {
			t.Fatal("negative control did not reproduce CI-834: reverting cluster_interface " +
				"and edge_node_cluster to Optional produced no diff, so the positive " +
				"assertion above proves nothing. Fix this test before trusting it.")
		}
		t.Logf("negative control reproduced the diff on %v", got)
	})
}

// TestEdgeNodeCluster_RealNode repeats the diff assertion against a node that
// is genuinely in a cluster, using the values the controller wrote.
//
// Why this is not redundant with the test above: that one asserts the schema
// behaves correctly for values THIS FILE invented. Only a live node shows what
// the controller actually populates -- field names, which of them are set, and
// the shape of edge_node_cluster on the wire. A provider that read the cluster
// block into the wrong attribute, or missed it entirely, would still pass the
// synthetic test while producing an empty state and a broken plan.
//
// Requires (otherwise skipped):
//
//	ZEDCLOUD_ACC_REAL_NODE=1
//	ZEDCLOUD_TEST_NODE_NAME=<name of an onboarded node THAT IS IN A CLUSTER>
//	TF_VAR_zedcloud_url / TF_VAR_zedcloud_token
//
// test/e2e-cluster creates a suitable node; its `device_names` output names
// them.
func TestEdgeNodeCluster_RealNode(t *testing.T) {
	node := testhelper.RealNode(t)

	live, err := testhelper.GetNodeClusterFields(node.ID)
	if err != nil {
		t.Fatalf("reading cluster fields from the controller: %s", err)
	}

	// Enabled-but-unclustered is an operator mistake, not a reason to skip.
	// Skipping here is how this test would silently stop covering anything the
	// moment someone points it at a plain node.
	if !live.Clustered() {
		t.Fatalf("node %q (%s) is not in an edge-node cluster: clusterInterface=%q edgeNodeCluster=%v.\n"+
			"This test needs a CLUSTERED node -- run test/e2e-cluster first and point "+
			"ZEDCLOUD_TEST_NODE_NAME at one of its device_names.",
			node.Name, node.ID, live.ClusterInterface, live.EdgeNodeCluster != nil)
	}

	t.Logf("live node %q: cluster_interface=%q cluster=%q prefix=%s seed_ip=%s is_master=%t",
		node.Name, live.ClusterInterface, live.EdgeNodeCluster.Name,
		live.EdgeNodeCluster.ClusterPrefix, live.EdgeNodeCluster.SeedNodeIP,
		live.EdgeNodeCluster.IsMaster)

	state := clusteredNodeState(clusterFields{
		ID:               node.ID,
		ClusterInterface: live.ClusterInterface,
		ClusterID:        live.EdgeNodeCluster.ID,
		ClusterName:      live.EdgeNodeCluster.Name,
		ClusterPrefix:    live.EdgeNodeCluster.ClusterPrefix,
		SeedNodeID:       live.EdgeNodeCluster.SeedNodeID,
		SeedNodeIP:       live.EdgeNodeCluster.SeedNodeIP,
		IsMaster:         live.EdgeNodeCluster.IsMaster,
		ProjectID:        live.EdgeNodeCluster.ProjectID,
	})

	cfg := terraform.NewResourceConfigRaw(map[string]interface{}{
		"name":       node.Name,
		"project_id": live.EdgeNodeCluster.ProjectID,
	})

	if got := clusterAttrsInDiff(t, schemas.Node(), state, cfg); len(got) != 0 {
		t.Errorf("CI-834 is live on node %q: a plan would change %v.\n"+
			"Applying it would remove the node from cluster %q.",
			node.Name, got, live.EdgeNodeCluster.Name)
	}
}

// ---------------------------------------------------------------------------
// helpers
// ---------------------------------------------------------------------------

// clusterFields is the flat form of what the controller populates.
type clusterFields struct {
	ID               string
	ClusterInterface string
	ClusterID        string
	ClusterName      string
	ClusterPrefix    string
	SeedNodeID       string
	SeedNodeIP       string
	IsMaster         bool
	ProjectID        string
}

// exampleClusterFields mirrors the values in the CI-834 report.
func exampleClusterFields() clusterFields {
	return clusterFields{
		ID:               "6e1233fc-8e8e-43bf-b048-87b2a309bcaa",
		ClusterInterface: "eth1",
		ClusterID:        "f73f1bdb-b229-4900-b564-99c2adf69f18",
		ClusterName:      "TEST_CLUSTER_x7",
		ClusterPrefix:    "10.244.244.2/28",
		SeedNodeID:       "264e393d-54bb-49de-a9b4-f6ca62b56f29",
		SeedNodeIP:       "10.244.244.1",
		IsMaster:         true,
		ProjectID:        "b91b2c7d-bc40-4422-81b6-304886b712ba",
	}
}

// clusteredNodeState builds the flatmap state Terraform would hold after
// refreshing a node that the controller has put into a cluster.
func clusteredNodeState(f clusterFields) *terraform.InstanceState {
	isMaster := "false"
	if f.IsMaster {
		isMaster = "true"
	}

	return &terraform.InstanceState{
		ID: f.ID,
		Attributes: map[string]string{
			"id":         f.ID,
			"name":       "ENODE_001",
			"project_id": f.ProjectID,

			// The fields under test, as the provider's SetNodeResourceData
			// writes them (v2/schemas/node.go).
			"cluster_interface":                  f.ClusterInterface,
			"edge_node_cluster.#":                "1",
			"edge_node_cluster.0.id":             f.ClusterID,
			"edge_node_cluster.0.name":           f.ClusterName,
			"edge_node_cluster.0.cluster_prefix": f.ClusterPrefix,
			"edge_node_cluster.0.seed_node_id":   f.SeedNodeID,
			"edge_node_cluster.0.seed_node_ip":   f.SeedNodeIP,
			"edge_node_cluster.0.is_master":      isMaster,
			"edge_node_cluster.0.project_id":     f.ProjectID,
		},
	}
}

// clusterAttrsInDiff returns the cluster attribute keys the SDK would change,
// sorted for a stable failure message.
//
// The resource is built from the schema map alone rather than via
// NodeResource(): NodeResource carries a CustomizeDiff
// (validateBondMemberInterfaces) that reads the raw cty config, which a
// hand-built ResourceConfig does not carry. Including it would test the
// CustomizeDiff's tolerance for synthetic input, not the schema behaviour
// CI-834 is about.
//
// Only cluster attributes are inspected, not the whole diff: the configuration
// here deliberately omits most of the node's other attributes, so an
// "is the diff empty" assertion would fail for reasons unrelated to CI-834.
func clusterAttrsInDiff(
	t *testing.T,
	nodeSchema map[string]*schema.Schema,
	state *terraform.InstanceState,
	cfg *terraform.ResourceConfig,
) []string {
	t.Helper()

	r := &schema.Resource{Schema: nodeSchema}

	diff, err := r.Diff(context.Background(), state, cfg, nil)
	if err != nil {
		t.Fatalf("computing diff: %s", err)
	}
	if diff == nil {
		return nil
	}

	var got []string
	for key, attr := range diff.Attributes {
		if attr == nil || !isClusterAttr(key) || !isDrop(attr) {
			continue
		}
		got = append(got, describeAttrDiff(key, attr))
	}
	sort.Strings(got)
	return got
}

// describeAttrDiff renders one diff entry for a failure message. Shared with
// the CI-709 tests in cluster_instances_test.go -- the old/new/flag triple is
// what makes a failure diagnosable rather than just red.
func describeAttrDiff(key string, attr *terraform.ResourceAttrDiff) string {
	return fmt.Sprintf("%s: %q -> %q (removed=%t computed=%t)",
		key, attr.Old, attr.New, attr.NewRemoved, attr.NewComputed)
}

// isDrop reports whether a diff entry would DISCARD a value the controller
// populated. That, and only that, is CI-834.
//
// *** NewComputed IS NOT THE BUG, AND MISTAKING IT FOR THE BUG IS EASY. ***
//
// The first version of this file reported any non-identity diff entry, and the
// real-node test then "failed" on all three nodes of a healthy cluster with:
//
//	cluster_interface:   "eth1" -> "" (removed=false computed=true)
//	edge_node_cluster.#: "1"    -> "" (removed=false computed=true)
//
// NewComputed means "the provider will supply this during apply" -- Terraform
// renders it as "(known after apply)", the Read then returns the same value
// that is already in state, and the plan is empty. It is the ordinary and
// correct treatment of a Computed attribute, and whether Resource.Diff emits
// it here depends on how complete the hand-built ResourceConfig is, which has
// nothing whatever to do with CI-834.
//
// The regression looks different: the value is DROPPED, either as an explicit
// removal or as a concrete empty New with no promise to recompute it:
//
//	cluster_interface: "eth1" -> "" (removed=true  computed=false)
//
// The negative control in TestEdgeNode_ClusterFieldsProduceNoDiff exists to
// keep this distinction honest -- it must still fire under this predicate.
func isDrop(attr *terraform.ResourceAttrDiff) bool {
	if attr.NewComputed {
		return false
	}
	if attr.NewRemoved {
		return true
	}
	return attr.Old != "" && attr.New == ""
}
