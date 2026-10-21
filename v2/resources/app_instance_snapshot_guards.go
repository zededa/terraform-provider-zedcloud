package resources

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	snapshots "github.com/zededa/terraform-provider-zedcloud/v2/client/app_instance_snapshot"
	config "github.com/zededa/terraform-provider-zedcloud/v2/client/application_instance"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

// Shared helpers for the app instance snapshot resources (PLE-576). See
// docs/design/ple-576-app-instance-immediate-snapshots.md §4.4 and §4.4a.

// snapshotPollInterval is how often the snapshot resources poll the snapshot
// state. State changes need the edge node to report back, which takes tens of
// seconds at best.
var snapshotPollInterval = 5 * time.Second

// readAppInstanceForSnapshot reads the app instance a snapshot operation
// targets.
func readAppInstanceForSnapshot(ctx context.Context, client *api_client.ZedcloudAPI, id string) (*models.AppInstance, error) {
	params := config.GetByIDParams().WithContext(ctx)
	params.ID = id
	resp, err := client.ApplicationInstance.GetByID(params, nil)
	if err != nil {
		if isStatusNotFound(err) {
			return nil, fmt.Errorf("edge application instance %s not found", id)
		}
		if ds, ok := ZsrvResponderToDiags(err); ok && len(ds) > 0 {
			return nil, fmt.Errorf("reading edge application instance %s: %s", id, ds[len(ds)-1].Summary)
		}
		return nil, fmt.Errorf("reading edge application instance %s: %w", id, err)
	}
	inst := resp.GetPayload()
	if inst == nil || inst.ID == "" {
		return nil, fmt.Errorf("edge application instance %s: the controller returned no instance", id)
	}
	return inst, nil
}

// appInstanceEdgeNodeClusterID returns the edge node cluster an app instance
// is deployed to, or "". This is edgeNodeCluster.id, which seine stores as
// EdgeNodeClusterID. clusterID is a different thing (the cluster instance).
func appInstanceEdgeNodeClusterID(inst *models.AppInstance) string {
	if inst == nil || inst.EdgeNodeCluster == nil {
		return ""
	}
	return inst.EdgeNodeCluster.ID
}

func appInstanceDisplayName(inst *models.AppInstance) string {
	if inst != nil && inst.Name != nil && *inst.Name != "" {
		return *inst.Name
	}
	if inst != nil {
		return inst.ID
	}
	return ""
}

// errSnapshotsOnCluster is returned for any snapshot operation on an app
// instance deployed to an edge node cluster (decisions D3 and D7).
func errSnapshotsOnCluster(appInstance, clusterID, what string) error {
	return fmt.Errorf("edge application instance %q is deployed to edge node cluster %s; "+
		"%s are not supported for edge node clusters", appInstance, clusterID, what)
}

// checkSnapshotNotOnCluster rejects app instances deployed to an edge node
// cluster (§4.4a). It is not affected by skip_capability_check.
func checkSnapshotNotOnCluster(inst *models.AppInstance) error {
	if cluster := appInstanceEdgeNodeClusterID(inst); cluster != "" {
		return errSnapshotsOnCluster(appInstanceDisplayName(inst), cluster, "app instance snapshots")
	}
	return nil
}

// checkImmediateSnapshotCapability checks that the edge node hosting inst
// reports an API capability level of at least
// API_CAPABILITY_VOLUME_SNAPSHOTS_IMMEDIATE (§4.4). zcli makes this check
// before snapshot-create, but tests for that exact name; EVE reports the
// highest level it implements, so the comparison must be >=. zcli only prints
// a message; Terraform fails.
func checkImmediateSnapshotCapability(ctx context.Context, client *api_client.ZedcloudAPI, inst *models.AppInstance) error {
	deviceID := ""
	if inst.DeviceID != nil {
		deviceID = *inst.DeviceID
	}
	if deviceID == "" {
		return fmt.Errorf("edge application instance %q has no edge node; cannot check that it supports "+
			"immediate snapshots. Set skip_capability_check = true to skip the check", appInstanceDisplayName(inst))
	}
	raw, err := client.AppInstanceSnapshot.DeviceRawStatus(ctx, deviceID)
	if err != nil {
		if errors.Is(err, snapshots.ErrNotFound) {
			raw = ""
		} else {
			return fmt.Errorf("reading the status of edge node %s: %w", deviceID, err)
		}
	}
	if raw == "" {
		return fmt.Errorf("edge node %s has not reported its status, so support for immediate snapshots "+
			"(%s) is unknown. Bring the node online, or set skip_capability_check = true",
			deviceID, snapshots.ImmediateSnapshotCapability)
	}
	ok, reported, err := snapshots.SupportsImmediateSnapshots(raw)
	if err != nil {
		return fmt.Errorf("edge node %s: %s. Set skip_capability_check = true to skip the check", deviceID, err)
	}
	if !ok {
		return fmt.Errorf("edge node %s reports API capability %s, below %s, so it cannot take immediate "+
			"snapshots; upgrade EVE-OS on the node, or set skip_capability_check = true",
			deviceID, reported, snapshots.ImmediateSnapshotCapability)
	}
	return nil
}

// snapshotClusterGuardDiff fails the plan for a new snapshot or activation on
// an app instance deployed to an edge node cluster, when the app instance id
// is already known. An unknown id (app instance created in the same apply) is
// checked again in Create. Read errors are left for Create to report, so a
// transient failure does not block planning.
func snapshotClusterGuardDiff(ctx context.Context, d *schema.ResourceDiff, m interface{}) error {
	if d.Id() != "" && !d.HasChange("app_instance_id") {
		return nil
	}
	if !d.NewValueKnown("app_instance_id") {
		return nil
	}
	id, _ := d.Get("app_instance_id").(string)
	client, ok := m.(*api_client.ZedcloudAPI)
	if id == "" || !ok || client == nil {
		return nil
	}
	inst, err := readAppInstanceForSnapshot(ctx, client, id)
	if err != nil {
		log.Printf("[DEBUG] app instance snapshot plan check skipped: %s", err)
		return nil
	}
	return checkSnapshotNotOnCluster(inst)
}

// waitForSnapshotState polls the snapshot state until check reports done or
// an error, or timeout passes. A missing state is polled again: ganges learns
// about a new snapshot a moment after seine creates it.
func waitForSnapshotState(
	ctx context.Context,
	c snapshots.ClientService,
	snapshotID string,
	timeout time.Duration,
	check func(st *snapshots.State) (done bool, err error),
) (*snapshots.State, error) {
	deadline := time.Now().Add(timeout)
	var last *snapshots.State
	var lastErr error
	for {
		st, err := c.State(ctx, snapshotID)
		switch {
		case err == nil:
			last, lastErr = st, nil
			done, cerr := check(st)
			if cerr != nil {
				return st, cerr
			}
			if done {
				return st, nil
			}
		case errors.Is(err, snapshots.ErrNotFound):
			// not visible yet
		default:
			// A 4xx won't go away; anything else (5xx, transport) may.
			var apiErr *snapshots.APIError
			if errors.As(err, &apiErr) && apiErr.Status >= 400 && apiErr.Status < 500 {
				return last, fmt.Errorf("reading snapshot state: %w", err)
			}
			lastErr = err
		}
		if time.Now().Add(snapshotPollInterval).After(deadline) {
			if lastErr != nil {
				return last, fmt.Errorf("timed out after %s; last state %s; last error reading it: %s",
					timeout, last.Name(), lastErr)
			}
			if last == nil {
				// Seen on a real controller: ganges dropped its state row for a
				// snapshot the device had not reported yet, and never recreated it.
				return nil, fmt.Errorf("timed out after %s; the controller never reported a state for the "+
					"snapshot. Check the edge node is online, and the snapshot with "+
					"data.zedcloud_app_instance_snapshots", timeout)
			}
			return last, fmt.Errorf("timed out after %s; last state %s. State changes need the edge node to "+
				"report back: check that it is online", timeout, last.Name())
		}
		select {
		case <-ctx.Done():
			return last, ctx.Err()
		case <-time.After(snapshotPollInterval):
		}
	}
}
