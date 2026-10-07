package resources

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	"github.com/zededa/terraform-provider-zedcloud/v2/client/artifact"
	zschema "github.com/zededa/terraform-provider-zedcloud/v2/schemas"
)

// artifactPollInterval is how often Create checks that an uploaded artifact
// has landed in object storage. On a controller it took about 2s.
var artifactPollInterval = time.Second

// ArtifactResource manages a file uploaded to the zedcloud ArtifactManager API,
// such as an Edge app logo. See docs/design/nfr-19-edge-app-logo-artifacts.md.
//
// Artifacts are immutable, so every input is ForceNew; Update exists only so
// retain_on_destroy can change without a replacement.
func ArtifactResource() *schema.Resource {
	return &schema.Resource{
		Description: "Uploads a file (an Edge app logo, a custom license, a brand or model logo) to zedcloud and " +
			"exposes its artifact id. Reference the id from, for example, " +
			"`zedcloud_application.manifest.desc.logo = { logo = zedcloud_artifact.x.id }`.",
		CreateContext: CreateArtifact,
		ReadContext:   ReadArtifact,
		UpdateContext: UpdateArtifact,
		DeleteContext: DeleteArtifact,
		Schema:        zschema.ArtifactSchema(),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(time.Minute),
		},
	}
}

// artifactContent returns the bytes to upload from source or content_base64.
func artifactContent(d *schema.ResourceData) ([]byte, error) {
	var data []byte
	if p, ok := d.GetOk("source"); ok && p.(string) != "" {
		b, err := os.ReadFile(p.(string))
		if err != nil {
			return nil, fmt.Errorf("reading source: %w", err)
		}
		data = b
	} else if c, ok := d.GetOk("content_base64"); ok && c.(string) != "" {
		b, err := base64.StdEncoding.DecodeString(c.(string))
		if err != nil {
			return nil, fmt.Errorf("decoding content_base64: %w", err)
		}
		data = b
	} else {
		return nil, errors.New("one of source or content_base64 must be set")
	}

	if len(data) == 0 {
		return nil, errors.New("artifact content is empty")
	}
	if len(data) > zschema.ArtifactMaxSize {
		return nil, fmt.Errorf("artifact content is %d bytes; the provider uploads at most %d bytes", len(data), zschema.ArtifactMaxSize)
	}
	return data, nil
}

func CreateArtifact(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*api_client.ZedcloudAPI)
	name := d.Get("name").(string)

	data, err := artifactContent(d)
	if err != nil {
		return diag.FromErr(err)
	}

	created, err := client.Artifact.Create(ctx, name)
	if err != nil {
		return diag.Errorf("creating artifact %q: %s", name, err)
	}
	id := created.ID

	// cleanup removes the half-created artifact when a later step fails, so a
	// failed apply does not leave an orphan behind. Best effort only.
	cleanup := func() {
		if derr := client.Artifact.Delete(context.Background(), id); derr != nil {
			log.Printf("[WARN] artifact %s: cleanup after failed create: %s", id, derr)
		}
	}

	if err := client.Artifact.Upload(ctx, id, data); err != nil {
		cleanup()
		return diag.Errorf("uploading artifact %q (%s): %s", name, id, err)
	}

	// The controller answers the upload with 202 before the object is stored,
	// and reports a failed store (for example, image metadata stripping
	// rejecting the file) only by the artifact never becoming available.
	if err := waitForArtifact(ctx, client.Artifact, id, d.Timeout(schema.TimeoutCreate)); err != nil {
		cleanup()
		return diag.Errorf("artifact %q (%s): %s", name, id, err)
	}

	d.SetId(id)
	if err := d.Set("size", len(data)); err != nil {
		return diag.FromErr(err)
	}
	return ReadArtifact(ctx, d, m)
}

func waitForArtifact(ctx context.Context, c artifact.ClientService, id string, timeout time.Duration) error {
	deadline := time.Now().Add(timeout)
	for {
		_, err := c.GetSignedURL(ctx, id)
		if err == nil {
			return nil
		}
		if !errors.Is(err, artifact.ErrNotFound) {
			return fmt.Errorf("checking upload: %w", err)
		}
		if time.Now().Add(artifactPollInterval).After(deadline) {
			return fmt.Errorf("upload was accepted but the artifact did not become available within %s. "+
				"The controller may have rejected the file (it strips metadata from PNG and JPEG images and "+
				"rejects files it cannot process)", timeout)
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(artifactPollInterval):
		}
	}
}

func ReadArtifact(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*api_client.ZedcloudAPI)
	id := d.Id()

	if _, err := client.Artifact.GetSignedURL(ctx, id); err != nil {
		if errors.Is(err, artifact.ErrNotFound) {
			log.Printf("[WARN] artifact %s not found on the controller, removing it from state", id)
			d.SetId("")
			return nil
		}
		return diag.Errorf("reading artifact %s: %s", id, err)
	}

	// The API returns no metadata for an artifact, only a download URL. After
	// an import, recover the name from the id ("<uuid>_<name>").
	if _, ok := d.GetOk("name"); !ok {
		if name := artifactNameFromID(id); name != "" {
			if err := d.Set("name", name); err != nil {
				return diag.FromErr(err)
			}
		}
	}
	if _, ok := d.GetOk("retain_on_destroy"); !ok {
		_ = d.Set("retain_on_destroy", false)
	}
	return nil
}

// artifactNameFromID returns the <name> part of an artifact id "<uuid>_<name>".
func artifactNameFromID(id string) string {
	if i := strings.Index(id, "_"); i >= 0 && i+1 < len(id) {
		return id[i+1:]
	}
	return ""
}

func UpdateArtifact(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Only retain_on_destroy can change in place; it lives in state alone.
	return ReadArtifact(ctx, d, m)
}

func DeleteArtifact(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if d.Get("retain_on_destroy").(bool) {
		log.Printf("[INFO] artifact %s: retain_on_destroy is set, leaving it on the controller", d.Id())
		d.SetId("")
		return nil
	}

	client := m.(*api_client.ZedcloudAPI)
	if err := client.Artifact.Delete(ctx, d.Id()); err != nil {
		return diag.Errorf("deleting artifact %s: %s", d.Id(), err)
	}
	d.SetId("")
	return nil
}
