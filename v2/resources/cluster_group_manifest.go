package resources

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	cluster_groups "github.com/zededa/terraform-provider-zedcloud/v2/client/cluster_group"
)

func ClusterGroupManifestResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: schema.NoopContext,
		ReadContext:   GetClusterGroupManifest,
		UpdateContext: schema.NoopContext,
		DeleteContext: schema.NoopContext,
		Schema:        ClusterGroupManifestSchema(),
	}
}

func ClusterGroupManifestDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: GetClusterGroupManifest,
		Schema:      ClusterGroupManifestSchema(),
	}
}

func ClusterGroupManifestSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"project_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"tags": {
			Type:     schema.TypeMap,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"yaml_content": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

func GetClusterGroupManifest(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*api_client.ZedcloudAPI)

	params := cluster_groups.NewGetClusterGroupManifestParams()

	if v, ok := d.GetOk("name"); ok {
		name := v.(string)
		params.SetName(&name)
	}

	if v, ok := d.GetOk("project_name"); ok {
		projectName := v.(string)
		params.SetProjectName(&projectName)
	}

	if v, ok := d.GetOk("tags"); ok {
		tagsMap := v.(map[string]interface{})
		params.SetTags(tagsMap)
	}

	resp, err := client.ClusterGroup.GetClusterGroupManifest(params, nil)
	if err != nil {
		log.Printf("[TRACE] GetClusterGroupManifest error: %s", spew.Sdump(err))
		return append(diags, diag.Errorf("GetClusterGroupManifest error: %s", err)...)
	}

	if resp.Payload != nil && resp.Payload.Manifest != nil {
		d.Set("yaml_content", resp.Payload.Manifest.YamlContent)
	}

	d.SetId(generateClusterGroupManifestID(d))

	return diags
}

func generateClusterGroupManifestID(d *schema.ResourceData) string {
	h := sha256.New()
	if v, ok := d.GetOk("name"); ok {
		fmt.Fprintf(h, "%s", v.(string))
	}
	if v, ok := d.GetOk("project_name"); ok {
		fmt.Fprintf(h, "%s", v.(string))
	}
	return hex.EncodeToString(h.Sum(nil))
}
