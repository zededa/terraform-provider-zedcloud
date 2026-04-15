package schemas

import (
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// TestEdgeNodeClusterConfigModelFromMap_ExampleJSON verifies that
// EdgeNodeClusterConfigModelFromMap does not panic when the manifest
// field is provided as a plain string (as returned from Terraform state).
func TestEdgeNodeClusterConfigModelFromMap_ExampleJSON(t *testing.T) {
	m := map[string]interface{}{
		"cluster_prefix":      "10.244.244.0/28",
		"id":                  "afec1006-4a0d-41e6-bccf-ffdb065deb7d",
		"is_master":           true,
		"manifest":            "", // comes in as string from Terraform state
		"name":                "test-cluster",
		"project_id":          "4fdc30c8-0bc8-4f78-b649-94053514f106",
		"seed_node_id":        "e523cb50-7e82-4193-8879-0d687efef70c",
		"seed_node_ip":        "10.244.244.2",
		"tags":                map[string]interface{}{},
		"tie_breaker_node_id": "14ba9bc9-31f6-4986-bcd8-4d545a5317e9",
		"token":               "WAi6N8OI6D1wqYjWzNMd7N5tNGnM9zsrKh0eZIKg6EE=",
	}

	result := EdgeNodeClusterConfigModelFromMap(m)
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if result.ID != "afec1006-4a0d-41e6-bccf-ffdb065deb7d" {
		t.Errorf("expected id 'afec1006-4a0d-41e6-bccf-ffdb065deb7d', got %q", result.ID)
	}
	if result.Name != "test-cluster" {
		t.Errorf("expected name 'test-cluster', got %q", result.Name)
	}
}

// TestEdgeNodeClusterConfigModelFromMap_StringManifest verifies that
// a non-empty base64 string manifest is parsed correctly without panicking.
func TestEdgeNodeClusterConfigModelFromMap_StringManifest(t *testing.T) {
	// "hello world" base64 encoded
	b64manifest := "aGVsbG8gd29ybGQ="
	m := map[string]interface{}{
		"cluster_prefix":      "10.244.244.0/28",
		"id":                  "afec1006-4a0d-41e6-bccf-ffdb065deb7d",
		"is_master":           false,
		"manifest":            b64manifest,
		"name":                "test-cluster",
		"project_id":          "4fdc30c8-0bc8-4f78-b649-94053514f106",
		"seed_node_id":        "e523cb50-7e82-4193-8879-0d687efef70c",
		"seed_node_ip":        "10.244.244.2",
		"tags":                map[string]interface{}{},
		"tie_breaker_node_id": "14ba9bc9-31f6-4986-bcd8-4d545a5317e9",
		"token":               "WAi6N8OI6D1wqYjWzNMd7N5tNGnM9zsrKh0eZIKg6EE=",
	}

	result := EdgeNodeClusterConfigModelFromMap(m)
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if string(result.Manifest) != "hello world" {
		t.Errorf("expected manifest 'hello world', got %q", string(result.Manifest))
	}
}

// TestEdgeNodeClusterConfigModelFromMap_Base64Manifest verifies that
// a strfmt.Base64 manifest value is handled correctly.
func TestEdgeNodeClusterConfigModelFromMap_Base64Manifest(t *testing.T) {
	manifest := strfmt.Base64("hello world")
	m := map[string]interface{}{
		"cluster_prefix":      "10.244.244.0/28",
		"id":                  "afec1006-4a0d-41e6-bccf-ffdb065deb7d",
		"is_master":           false,
		"manifest":            manifest,
		"name":                "test-cluster",
		"project_id":          "4fdc30c8-0bc8-4f78-b649-94053514f106",
		"seed_node_id":        "e523cb50-7e82-4193-8879-0d687efef70c",
		"seed_node_ip":        "10.244.244.2",
		"tags":                map[string]interface{}{},
		"tie_breaker_node_id": "14ba9bc9-31f6-4986-bcd8-4d545a5317e9",
		"token":               "WAi6N8OI6D1wqYjWzNMd7N5tNGnM9zsrKh0eZIKg6EE=",
	}

	result := EdgeNodeClusterConfigModelFromMap(m)
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if string(result.Manifest) != "hello world" {
		t.Errorf("expected manifest 'hello world', got %q", string(result.Manifest))
	}
}

// TestEdgeNodeClusterConfigModel_ExampleJSON verifies that
// EdgeNodeClusterConfigModel reads from ResourceData without panicking,
// including when manifest is stored as a string in Terraform state.
func TestEdgeNodeClusterConfigModel_ExampleJSON(t *testing.T) {
	raw := map[string]interface{}{
		"cluster_prefix":      "10.244.244.0/28",
		"id":                  "afec1006-4a0d-41e6-bccf-ffdb065deb7d",
		"is_master":           true,
		"manifest":            "aGVsbG8gd29ybGQ=",
		"name":                "test-cluster",
		"project_id":          "4fdc30c8-0bc8-4f78-b649-94053514f106",
		"seed_node_id":        "e523cb50-7e82-4193-8879-0d687efef70c",
		"seed_node_ip":        "10.244.244.2",
		"tags":                map[string]interface{}{},
		"tie_breaker_node_id": "14ba9bc9-31f6-4986-bcd8-4d545a5317e9",
		"token":               "WAi6N8OI6D1wqYjWzNMd7N5tNGnM9zsrKh0eZIKg6EE=",
	}

	d := schema.TestResourceDataRaw(t, EdgeNodeClusterConfigSchema(), raw)

	result := EdgeNodeClusterConfigModel(d)
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if result.Name != "test-cluster" {
		t.Errorf("expected name 'test-cluster', got %q", result.Name)
	}
	if string(result.Manifest) != "hello world" {
		t.Errorf("expected manifest 'hello world', got %q", string(result.Manifest))
	}
}
