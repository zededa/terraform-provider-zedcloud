package resources

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	snapshots "github.com/zededa/terraform-provider-zedcloud/v2/client/app_instance_snapshot"
	zschema "github.com/zededa/terraform-provider-zedcloud/v2/schemas"
)

// AppInstanceSnapshotResource manages an immediate snapshot of an edge
// application instance, as zcli's `edge-app-instance snapshot-create`.
// See docs/design/ple-576-app-instance-immediate-snapshots.md §4.2.
//
// The controller keeps one snapshot per app instance: a new snapshot (from
// this resource, another one, or an update with makeSnapshot) deletes the
// previous one, and Read then drops this resource from state.
func AppInstanceSnapshotResource() *schema.Resource {
	return &schema.Resource{
		Description: "Takes an immediate snapshot of an edge application instance. The controller keeps one " +
			"snapshot per app instance, so a new snapshot replaces the previous one. Not supported for app " +
			"instances deployed to an edge node cluster.",
		CreateContext: CreateAppInstanceSnapshot,
		ReadContext:   ReadAppInstanceSnapshot,
		UpdateContext: UpdateAppInstanceSnapshot,
		DeleteContext: DeleteAppInstanceSnapshot,
		CustomizeDiff: snapshotClusterGuardDiff,
		Schema:        zschema.AppInstanceSnapshotSchema(),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
		},
	}
}

func CreateAppInstanceSnapshot(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*api_client.ZedcloudAPI)
	appInstanceID := d.Get("app_instance_id").(string)

	inst, err := readAppInstanceForSnapshot(ctx, client, appInstanceID)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := checkSnapshotNotOnCluster(inst); err != nil {
		return diag.FromErr(err)
	}
	if !d.Get("skip_capability_check").(bool) {
		if err := checkImmediateSnapshotCapability(ctx, client, inst); err != nil {
			return diag.FromErr(err)
		}
	}

	created, err := client.AppInstanceSnapshot.Create(ctx, appInstanceID)
	if err != nil {
		return snapshotAPIErrorDiags(fmt.Sprintf("creating a snapshot of edge application instance %q",
			appInstanceDisplayName(inst)), err)
	}
	id, name := created.ID, created.Name
	if id == "" {
		// Controllers older than ENG-1676 don't return the snapshot identity.
		// Create leaves exactly one snapshot, so list and take it.
		s, err := findCreatedSnapshot(ctx, client.AppInstanceSnapshot, appInstanceID, name)
		if err != nil {
			return diag.Errorf("snapshot of edge application instance %q was requested, but finding it failed: %s",
				appInstanceDisplayName(inst), err)
		}
		id, name = s.ID, s.Name
	}

	// Set the id before waiting, so a failed or timed-out wait leaves a tainted
	// resource instead of an orphaned snapshot.
	d.SetId(id)
	_ = d.Set("name", name)

	if d.Get("wait_for_ready").(bool) {
		st, err := waitForSnapshotState(ctx, client.AppInstanceSnapshot, id, d.Timeout(schema.TimeoutCreate),
			func(st *snapshots.State) (bool, error) {
				switch st.Name() {
				case snapshots.StateCreated, snapshots.StateActive:
					return true, nil
				case snapshots.StateCreateFailed:
					return false, errors.New("the edge node failed to create the snapshot (CREATE_FAILED)")
				case snapshots.StateDeleting, snapshots.StateDeleted, snapshots.StateDeleteFailed:
					return false, fmt.Errorf("the snapshot was deleted while being created (%s); "+
						"another snapshot of the same app instance replaces it", st.Name())
				}
				return false, nil
			})
		if st != nil {
			_ = d.Set("state", st.Name())
		}
		if err != nil {
			return diag.Errorf("snapshot %s of edge application instance %q: %s", name, appInstanceDisplayName(inst), err)
		}
	}

	return ReadAppInstanceSnapshot(ctx, d, m)
}

// findCreatedSnapshot finds the snapshot a Create made when the controller did
// not return its id.
func findCreatedSnapshot(ctx context.Context, c snapshots.ClientService, appInstanceID, name string) (*snapshots.Snapshot, error) {
	list, err := c.List(ctx, appInstanceID)
	if err != nil {
		return nil, err
	}
	for _, s := range list {
		if s != nil && name != "" && s.Name == name {
			return s, nil
		}
	}
	if len(list) == 1 && list[0] != nil && list[0].ID != "" {
		return list[0], nil
	}
	return nil, fmt.Errorf("expected one snapshot, found %d", len(list))
}

func ReadAppInstanceSnapshot(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*api_client.ZedcloudAPI)
	id := d.Id()

	s, err := client.AppInstanceSnapshot.Get(ctx, id)
	if err != nil {
		if errors.Is(err, snapshots.ErrNotFound) {
			log.Printf("[WARN] app instance snapshot %s not found (a newer snapshot of the same app instance "+
				"replaces it), removing it from state", id)
			d.SetId("")
			return nil
		}
		return snapshotAPIErrorDiags(fmt.Sprintf("reading app instance snapshot %s", id), err)
	}

	// Import: only the id is known. Recover the rest and the defaults.
	if d.Get("app_instance_id").(string) == "" {
		_ = d.Set("app_instance_id", s.AppInstanceID())
		_ = d.Set("wait_for_ready", true)
		_ = d.Set("skip_capability_check", false)
	}
	_ = d.Set("name", s.Name)
	cfg, err := s.ConfigJSON()
	if err != nil {
		log.Printf("[WARN] app instance snapshot %s: %s", id, err)
	}
	_ = d.Set("config_json", cfg)

	st, err := client.AppInstanceSnapshot.State(ctx, id)
	switch {
	case err == nil:
		_ = d.Set("state", st.Name())
	case errors.Is(err, snapshots.ErrNotFound):
		// ganges hasn't seen it yet; keep the last known state.
	default:
		log.Printf("[WARN] app instance snapshot %s: reading state: %s", id, err)
	}
	return nil
}

// UpdateAppInstanceSnapshot only handles wait_for_ready and
// skip_capability_check, which live in state alone. Everything else is
// ForceNew.
func UpdateAppInstanceSnapshot(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return ReadAppInstanceSnapshot(ctx, d, m)
}

func DeleteAppInstanceSnapshot(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*api_client.ZedcloudAPI)
	appInstanceID := d.Get("app_instance_id").(string)
	name := d.Get("name").(string)
	if appInstanceID == "" || name == "" {
		d.SetId("")
		return nil
	}
	// seine removes the snapshot at once (a missing one is not an error), so
	// there is nothing to wait for.
	if err := client.AppInstanceSnapshot.Delete(ctx, appInstanceID, name); err != nil {
		return snapshotAPIErrorDiags(fmt.Sprintf("deleting app instance snapshot %s", name), err)
	}
	d.SetId("")
	return nil
}

// snapshotAPIErrorDiags reports an API error with the controller's details.
func snapshotAPIErrorDiags(what string, err error) diag.Diagnostics {
	if errors.Is(err, snapshots.ErrNotFound) {
		return diag.Errorf("%s: not found", what)
	}
	var apiErr *snapshots.APIError
	if errors.As(err, &apiErr) {
		return diag.Errorf("%s: %s", what, apiErr.Error())
	}
	return diag.Errorf("%s: %s", what, err)
}
