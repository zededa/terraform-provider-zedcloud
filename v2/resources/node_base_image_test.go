package resources

import (
	"context"
	"strings"
	"testing"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
	zschema "github.com/zededa/terraform-provider-zedcloud/v2/schemas"
)

// The controller writes base_image onto a node's device config on its own: an
// edge-node cluster EVE-OS upgrade calls devBaseImagePublish + devBaseImageApply
// for each member node in turn, and the UI can do the same for a standalone
// node. The provider refreshes that into state.
//
// A configuration that never mentions base_image must not be dragged into a
// diff by it -- and it must especially not be dragged into an apply, because
// PUT /v1/devices/id/{id} silently refuses to clear baseImage
// (devproc.go: dev.BaseImage = exists.BaseImage), so the diff would come back
// on the next plan, forever.
//
// Reported twice: CI-736, then CI-836 after a cluster upgrade.
func TestNodeDiff_BaseImageAbsentFromConfig_NoDiff(t *testing.T) {
	r := &schema.Resource{Schema: zschema.Node()}

	state := &terraform.InstanceState{
		ID: "51fa1d42-8a05-42bb-91fc-45c3fa81415d",
		Attributes: map[string]string{
			"id":                      "51fa1d42-8a05-42bb-91fc-45c3fa81415d",
			"name":                    "ENODE_001",
			"base_image.#":            "1",
			"base_image.0.image_name": "16.5.0-k-amd64",
			"base_image.0.activate":   "true",
			"base_image.0.imvol_id":   "e0203864-2848-4ef9-879b-bae2aad353bf",
			"base_image.0.uuid":       "",
			"base_image.0.version":    "",
		},
		// The user's HCL: no base_image block anywhere in it.
		RawConfig: cty.ObjectVal(map[string]cty.Value{
			"name": cty.StringVal("ENODE_001"),
		}),
	}

	cfg := terraform.NewResourceConfigRaw(map[string]interface{}{
		"name": "ENODE_001",
	})

	diff, err := r.Diff(context.Background(), state, cfg, nil)
	if err != nil {
		t.Fatalf("Diff returned an error: %s", err)
	}

	if diff == nil {
		return // no diff at all is the ideal outcome
	}
	for k := range diff.Attributes {
		if strings.HasPrefix(k, "base_image") {
			t.Errorf("base_image must not appear in the diff when it is absent from the config, got %q: %#v",
				k, diff.Attributes[k])
		}
	}
}

// The suppression above must not become a blanket "never diff base_image":
// a standalone node that does declare base_image still has to be able to
// change it.
func TestNodeDiff_BaseImageInConfig_StillDiffs(t *testing.T) {
	r := &schema.Resource{Schema: zschema.Node()}

	state := &terraform.InstanceState{
		ID: "51fa1d42-8a05-42bb-91fc-45c3fa81415d",
		Attributes: map[string]string{
			"id":                      "51fa1d42-8a05-42bb-91fc-45c3fa81415d",
			"name":                    "ENODE_001",
			"base_image.#":            "1",
			"base_image.0.image_name": "16.5.0-k-amd64",
			"base_image.0.activate":   "true",
		},
		RawConfig: cty.ObjectVal(map[string]cty.Value{
			"name": cty.StringVal("ENODE_001"),
			"base_image": cty.ListVal([]cty.Value{
				cty.ObjectVal(map[string]cty.Value{
					"image_name": cty.StringVal("17.0.0-lts-k-amd64"),
					"activate":   cty.True,
				}),
			}),
		}),
	}

	cfg := terraform.NewResourceConfigRaw(map[string]interface{}{
		"name": "ENODE_001",
		"base_image": []interface{}{
			map[string]interface{}{
				"image_name": "17.0.0-lts-k-amd64",
				"activate":   true,
			},
		},
	})

	diff, err := r.Diff(context.Background(), state, cfg, nil)
	if err != nil {
		t.Fatalf("Diff returned an error: %s", err)
	}
	if diff == nil {
		t.Fatal("expected a diff when the config changes base_image, got none")
	}

	found := false
	for k := range diff.Attributes {
		if strings.HasPrefix(k, "base_image") {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("expected base_image in the diff when the config changes it, got: %#v", diff.Attributes)
	}
}

// A node in an edge-node cluster must be told at plan time that base_image
// belongs on the cluster, rather than discovering it mid-apply as a 400 from
// devBaseImageApply (CI-836).
func TestNodeDiff_BaseImageOnClusterMember_FailsAtPlanTime(t *testing.T) {
	r := NodeResource()

	state := &terraform.InstanceState{
		ID: "51fa1d42-8a05-42bb-91fc-45c3fa81415d",
		Attributes: map[string]string{
			"id":                       "51fa1d42-8a05-42bb-91fc-45c3fa81415d",
			"name":                     "ENODE_001",
			"edge_node_cluster.#":      "1",
			"edge_node_cluster.0.id":   "bfda031f-a79d-4724-b6b8-29e3b38dc673",
			"edge_node_cluster.0.name": "TEST_CLUSTER",
		},
		RawConfig: cty.ObjectVal(map[string]cty.Value{
			"name": cty.StringVal("ENODE_001"),
			"base_image": cty.ListVal([]cty.Value{
				cty.ObjectVal(map[string]cty.Value{
					"image_name": cty.StringVal("16.5.0-k-amd64"),
					"activate":   cty.True,
				}),
			}),
		}),
	}

	cfg := terraform.NewResourceConfigRaw(map[string]interface{}{
		"name": "ENODE_001",
		"base_image": []interface{}{
			map[string]interface{}{
				"image_name": "16.5.0-k-amd64",
				"activate":   true,
			},
		},
	})

	_, err := r.Diff(context.Background(), state, cfg, nil)
	if err == nil {
		t.Fatal("expected the plan to fail for base_image on a cluster member, got no error")
	}
	for _, want := range []string{"bfda031f-a79d-4724-b6b8-29e3b38dc673", "zedcloud_edgenode_cluster"} {
		if !strings.Contains(err.Error(), want) {
			t.Errorf("error should mention %q so the user knows where base_image belongs, got: %s", want, err)
		}
	}
}

// The same configuration on a node that is NOT in a cluster is perfectly legal.
func TestNodeDiff_BaseImageOnStandaloneNode_Allowed(t *testing.T) {
	r := NodeResource()

	state := &terraform.InstanceState{
		ID: "51fa1d42-8a05-42bb-91fc-45c3fa81415d",
		Attributes: map[string]string{
			"id":   "51fa1d42-8a05-42bb-91fc-45c3fa81415d",
			"name": "ENODE_001",
		},
		RawConfig: cty.ObjectVal(map[string]cty.Value{
			"name": cty.StringVal("ENODE_001"),
			"base_image": cty.ListVal([]cty.Value{
				cty.ObjectVal(map[string]cty.Value{
					"image_name": cty.StringVal("16.5.0-k-amd64"),
					"activate":   cty.True,
				}),
			}),
		}),
	}

	cfg := terraform.NewResourceConfigRaw(map[string]interface{}{
		"name": "ENODE_001",
		"base_image": []interface{}{
			map[string]interface{}{
				"image_name": "16.5.0-k-amd64",
				"activate":   true,
			},
		},
	})

	if _, err := r.Diff(context.Background(), state, cfg, nil); err != nil {
		t.Fatalf("base_image on a standalone node must be allowed, got: %s", err)
	}
}

func TestRejectBaseImageOnClusterMember(t *testing.T) {
	name := "ENODE_001"

	tests := []struct {
		name    string
		node    *models.Node
		wantErr bool
	}{
		{
			name:    "nil node",
			node:    nil,
			wantErr: false,
		},
		{
			name:    "standalone node",
			node:    &models.Node{Name: &name},
			wantErr: false,
		},
		{
			name: "cluster block present but empty",
			node: &models.Node{
				Name:            &name,
				EdgeNodeCluster: &models.EdgeNodeClusterConfig{},
			},
			wantErr: false,
		},
		{
			name: "cluster member",
			node: &models.Node{
				Name: &name,
				EdgeNodeCluster: &models.EdgeNodeClusterConfig{
					ID: "bfda031f-a79d-4724-b6b8-29e3b38dc673",
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diags := rejectBaseImageOnClusterMember(tt.node)
			if got := diags.HasError(); got != tt.wantErr {
				t.Fatalf("HasError() = %v, want %v (diags: %v)", got, tt.wantErr, diags)
			}
		})
	}
}
