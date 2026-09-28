package schemas

import (
	"testing"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

// clusterData builds a ResourceData the way the SDK does at runtime, from an
// InstanceState. TestResourceDataRaw does not hydrate nested blocks, so a
// base_image written through it is invisible to d.GetOk.
func clusterData(attrs map[string]string) *schema.ResourceData {
	r := &schema.Resource{Schema: ClusterSchema()}
	return r.Data(&terraform.InstanceState{
		ID:         "bfda031f-a79d-4724-b6b8-29e3b38dc673",
		Attributes: attrs,
	})
}

func clusterWithBaseImage(imageName, activate string) map[string]string {
	return map[string]string{
		"id":                      "bfda031f-a79d-4724-b6b8-29e3b38dc673",
		"name":                    "CLUSTER",
		"base_image.#":            "1",
		"base_image.0.image_name": imageName,
		"base_image.0.activate":   activate,
	}
}

func TestClusterBaseImageModel(t *testing.T) {
	t.Run("absent from config", func(t *testing.T) {
		d := clusterData(map[string]string{"name": "CLUSTER"})
		if got := ClusterBaseImageModel(d); got != nil {
			t.Fatalf("expected nil when base_image is not configured, got %+v", got)
		}
	})

	t.Run("image name and activate", func(t *testing.T) {
		d := clusterData(clusterWithBaseImage("16.5.0-k-amd64", "true"))

		got := ClusterBaseImageModel(d)
		if got == nil {
			t.Fatal("expected a body, got nil")
		}
		if got.ImageName == nil || *got.ImageName != "16.5.0-k-amd64" {
			t.Errorf("ImageName = %v, want 16.5.0-k-amd64", got.ImageName)
		}
		if got.Activate == nil || !*got.Activate {
			t.Errorf("Activate = %v, want true", got.Activate)
		}
		// uuid and version are required by the swagger definition but carry no
		// meaning for a cluster upgrade; they must go out empty rather than
		// nil, which would marshal as null.
		if got.UUID == nil || *got.UUID != "" {
			t.Errorf("UUID = %v, want a pointer to the empty string", got.UUID)
		}
		if got.Version == nil || *got.Version != "" {
			t.Errorf("Version = %v, want a pointer to the empty string", got.Version)
		}
	})

	t.Run("empty image name is not a request", func(t *testing.T) {
		d := clusterData(clusterWithBaseImage("", "true"))
		if got := ClusterBaseImageModel(d); got != nil {
			t.Fatalf("expected nil for an empty image_name, got %+v", got)
		}
	})
}

func TestSetClusterUpgradeResourceData(t *testing.T) {
	completed := models.EdgeNodeClusterUpgradeStatusSTATUSCOMPLETED
	unspecified := models.EdgeNodeClusterUpgradeStatusSTATUSUNSPECIFIED

	older := strfmt.DateTime(time.Date(2026, 2, 26, 16, 36, 39, 0, time.UTC))
	newer := strfmt.DateTime(time.Date(2026, 2, 26, 17, 10, 0, 0, time.UTC))

	t.Run("derives base_image from the most recently updated node", func(t *testing.T) {
		d := clusterData(map[string]string{"name": "CLUSTER"})

		SetClusterUpgradeResourceData(d, &models.EdgeNodeClusterUpgradeStatusResp{
			Nodes: []*models.EdgeNodeClusterUpgradeStatusRespNode{
				{NodeID: "a", UpgradeableEveOs: "16.4.0-k-amd64", Status: &completed, UpdatedAt: older},
				{NodeID: "b", UpgradeableEveOs: "16.5.0-k-amd64", Status: &completed, UpdatedAt: newer},
				{NodeID: "c", UpgradeableEveOs: "", Status: &unspecified},
			},
		})

		raw, isSet := d.GetOk("base_image")
		if !isSet {
			t.Fatal("base_image was not set")
		}
		items := raw.([]interface{})
		if len(items) != 1 {
			t.Fatalf("expected exactly one base_image block, got %d", len(items))
		}
		got := items[0].(map[string]interface{})
		if got["image_name"] != "16.5.0-k-amd64" {
			t.Errorf("image_name = %v, want 16.5.0-k-amd64", got["image_name"])
		}

		status := d.Get("upgrade_status").([]interface{})
		if len(status) != 3 {
			t.Fatalf("expected a status row per node, got %d", len(status))
		}
		first := status[0].(map[string]interface{})
		if first["node_id"] != "a" {
			t.Errorf("node_id = %v, want a", first["node_id"])
		}
		if first["status"] != string(models.EdgeNodeClusterUpgradeStatusSTATUSCOMPLETED) {
			t.Errorf("status = %v, want STATUS_COMPLETED", first["status"])
		}
	})

	// A cluster that has never been upgraded reports rows with no image. Writing
	// an empty base_image block into state would invent a diff against a config
	// that says nothing about base images.
	t.Run("no rollout leaves base_image alone", func(t *testing.T) {
		d := clusterData(map[string]string{"name": "CLUSTER"})

		SetClusterUpgradeResourceData(d, &models.EdgeNodeClusterUpgradeStatusResp{
			Nodes: []*models.EdgeNodeClusterUpgradeStatusRespNode{
				{NodeID: "a", UpgradeableEveOs: "", Status: &unspecified},
				{NodeID: "b", UpgradeableEveOs: "", Status: &unspecified},
			},
		})

		if raw, isSet := d.GetOk("base_image"); isSet {
			t.Fatalf("base_image should stay unset when nothing has been rolled out, got %v", raw)
		}
		if got := len(d.Get("upgrade_status").([]interface{})); got != 2 {
			t.Errorf("upgrade_status rows = %d, want 2", got)
		}
	})

	t.Run("nil response is a no-op", func(t *testing.T) {
		d := clusterData(map[string]string{"name": "CLUSTER"})
		SetClusterUpgradeResourceData(d, nil)
		if raw, isSet := d.GetOk("base_image"); isSet {
			t.Fatalf("expected base_image to stay unset, got %v", raw)
		}
	})

	// activate is not reported by the API, so a configured value has to survive
	// the refresh or every plan would show it flipping.
	t.Run("keeps the configured activate", func(t *testing.T) {
		d := clusterData(clusterWithBaseImage("16.5.0-k-amd64", "false"))

		SetClusterUpgradeResourceData(d, &models.EdgeNodeClusterUpgradeStatusResp{
			Nodes: []*models.EdgeNodeClusterUpgradeStatusRespNode{
				{NodeID: "a", UpgradeableEveOs: "16.5.0-k-amd64", Status: &completed, UpdatedAt: newer},
			},
		})

		items := d.Get("base_image").([]interface{})
		got := items[0].(map[string]interface{})
		if got["activate"] != false {
			t.Errorf("activate = %v, want the configured false", got["activate"])
		}
	})
}
