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

// AppInstanceSnapshotActivationResource rolls an edge application instance
// back to one of its snapshots, as zcli's `edge-app-instance
// snapshot-activate`. It is one-shot: Create performs the rollback and waits
// for the edge node to confirm it; destroy only removes it from state, since
// a rollback can't be undone. See the design doc §4.3 and §5.2.
//
// After a rollback the controller holds the snapshot's config, so the
// zedcloud_application_instance resource shows it as drift on its next
// refresh, like any other out-of-band change (decision D4).
func AppInstanceSnapshotActivationResource() *schema.Resource {
	return &schema.Resource{
		Description: "Rolls an edge application instance back to one of its snapshots and waits until the edge " +
			"node reports the snapshot as `ACTIVE`. Destroying this resource does not undo the rollback. After a " +
			"rollback, `zedcloud_application_instance` shows the rolled-back configuration as drift. Not " +
			"supported for app instances deployed to an edge node cluster.",
		CreateContext: CreateAppInstanceSnapshotActivation,
		ReadContext:   ReadAppInstanceSnapshotActivation,
		DeleteContext: DeleteAppInstanceSnapshotActivation,
		CustomizeDiff: snapshotClusterGuardDiff,
		Schema:        zschema.AppInstanceSnapshotActivationSchema(),
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
		},
	}
}

func CreateAppInstanceSnapshotActivation(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*api_client.ZedcloudAPI)
	appInstanceID := d.Get("app_instance_id").(string)
	name := d.Get("snapshot_name").(string)

	inst, err := readAppInstanceForSnapshot(ctx, client, appInstanceID)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := checkSnapshotNotOnCluster(inst); err != nil {
		return diag.FromErr(err)
	}
	instName := appInstanceDisplayName(inst)

	snap, err := client.AppInstanceSnapshot.GetByName(ctx, appInstanceID, name)
	if err != nil {
		if errors.Is(err, snapshots.ErrNotFound) {
			return diag.Errorf("edge application instance %q has no snapshot named %q. The controller keeps one "+
				"snapshot per app instance, so a newer snapshot may have replaced it", instName, name)
		}
		return snapshotAPIErrorDiags(fmt.Sprintf("reading snapshot %q of edge application instance %q", name, instName), err)
	}

	prev, err := client.AppInstanceSnapshot.State(ctx, snap.ID)
	if err != nil {
		if !errors.Is(err, snapshots.ErrNotFound) {
			return snapshotAPIErrorDiags(fmt.Sprintf("reading the state of snapshot %q", name), err)
		}
		prev = nil
	}

	// A snapshot can be rolled back to only once. The controller always sends
	// EVE RollbackCmd{Counter: 1} for the active snapshot (purus appinst.go),
	// and EVE rolls back only when that counter changes. A second activate
	// would make seine rewrite the app instance config to the snapshot's, while
	// the device keeps its volumes and the state stays ROLLING_BACK forever.
	// Refuse before anything is sent (verified against a device, PLE-576 §9).
	//
	// Only a CREATED snapshot can be rolled back to. An app-update snapshot
	// (snapshot_on_update) stays CREATING until an update restarts the app;
	// activating it then makes EVE reject the whole app config ("Error
	// restoring config from snapshot ... not available"), and the app stops
	// receiving config until the snapshot is deleted.
	switch prev.Name() {
	case snapshots.StateCreated:
	case snapshots.StateActive, snapshots.StateRollingBack, snapshots.StateRollbackFailed:
		return diag.Errorf("snapshot %q of edge application instance %q was already rolled back to (state %s). "+
			"The controller can roll back to a snapshot only once: a second rollback would restore the "+
			"configuration but not the volumes, and never complete. Take a new snapshot and roll back to that",
			name, instName, prev.Name())
	default:
		return diag.Errorf("snapshot %q of edge application instance %q is %s, not CREATED, so there is nothing "+
			"to roll back to yet. A snapshot taken by snapshot_on_update is only created on the edge node when an "+
			"update restarts the app; wait for CREATED (data.zedcloud_app_instance_snapshots shows the state)",
			name, instName, prev.Name())
	}

	// force=true: the controller never clears a snapshot's activate flag, and
	// answers a non-forced activate of a flagged snapshot with a silent no-op
	// (for example after an earlier attempt failed before reaching the
	// device). For a snapshot never activated, force changes nothing.
	if err := client.AppInstanceSnapshot.Activate(ctx, appInstanceID, name, true); err != nil {
		return snapshotAPIErrorDiags(fmt.Sprintf("activating snapshot %q of edge application instance %q", name, instName), err)
	}

	// From here on the controller config has changed. Set the id first, so a
	// failure below leaves a tainted resource and the next apply checks again.
	d.SetId(appInstanceID + "/" + name)
	_ = d.Set("snapshot_id", snap.ID)

	final, err := waitForSnapshotState(ctx, client.AppInstanceSnapshot, snap.ID, d.Timeout(schema.TimeoutCreate),
		func(st *snapshots.State) (bool, error) {
			// Ignore the state from before the request: until ganges
			// records the rollback it still reports the old one.
			if prev != nil && st.Name() == prev.Name() && st.Timestamp == prev.Timestamp {
				return false, nil
			}
			switch st.Name() {
			case snapshots.StateActive:
				return true, nil
			case snapshots.StateRollbackFailed:
				return false, errors.New("the edge node failed to roll back (ROLLBACK_FAILED)")
			case snapshots.StateDeleting, snapshots.StateDeleted, snapshots.StateDeleteFailed:
				return false, fmt.Errorf("the snapshot was deleted during the rollback (%s)", st.Name())
			}
			return false, nil
		})
	if final != nil {
		_ = d.Set("state", final.Name())
	}
	if err != nil {
		return diag.Errorf("rollback of edge application instance %q to snapshot %q was requested and the "+
			"controller configuration has changed, but the rollback is not confirmed: %s", instName, name, err)
	}
	_ = d.Set("activated_at", final.Timestamp)
	return nil
}

// ReadAppInstanceSnapshotActivation refreshes the state. If the snapshot is
// gone the resource stays: the rollback already happened, and removing it
// would make the next apply roll back again.
func ReadAppInstanceSnapshotActivation(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*api_client.ZedcloudAPI)
	snapID := d.Get("snapshot_id").(string)
	if snapID == "" {
		return nil
	}
	st, err := client.AppInstanceSnapshot.State(ctx, snapID)
	switch {
	case err == nil:
		_ = d.Set("state", st.Name())
	case errors.Is(err, snapshots.ErrNotFound):
		_ = d.Set("state", "")
	default:
		log.Printf("[WARN] app instance snapshot activation %s: reading state: %s", d.Id(), err)
	}
	return nil
}

func DeleteAppInstanceSnapshotActivation(_ context.Context, d *schema.ResourceData, _ interface{}) diag.Diagnostics {
	d.SetId("")
	return diag.Diagnostics{{
		Severity: diag.Warning,
		Summary:  "Rollback not undone",
		Detail: "A snapshot rollback cannot be undone. zedcloud_app_instance_snapshot_activation was only " +
			"removed from the Terraform state; the edge application instance keeps the rolled-back configuration.",
	}}
}
