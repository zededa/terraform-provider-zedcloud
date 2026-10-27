package resources

import (
	"context"
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	"github.com/zededa/terraform-provider-zedcloud/v2/client/application_instance_v2"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

// Snapshot settings on zedcloud_application_instance: allow_snapshots,
// max_snapshots and snapshot_on_update (PLE-576). See
// docs/design/ple-576-app-instance-immediate-snapshots.md §4.6 and §4.7.

// snapshotSettingsReader is what ResourceData and ResourceDiff have in common.
type snapshotSettingsReader interface {
	Get(key string) interface{}
	HasChange(key string) bool
}

func configuredEdgeNodeClusterID(d snapshotSettingsReader) string {
	l, _ := d.Get("edge_node_cluster").([]interface{})
	if len(l) == 0 || l[0] == nil {
		return ""
	}
	m, _ := l[0].(map[string]interface{})
	id, _ := m["id"].(string)
	return id
}

// checkAppInstanceSnapshotSettings validates the snapshot settings of an app
// instance. Only settings being turned on count: max_snapshots is Computed and
// the controller reports 0 as 1, so an untouched value must not trip the
// docker-compose or cluster rules. allowKnown is false when allow_snapshots is
// not known yet (plan of a new instance that doesn't set it).
func checkAppInstanceSnapshotSettings(d snapshotSettingsReader, allowKnown bool) error {
	allow, _ := d.Get("allow_snapshots").(bool)
	maxSnapshots, _ := d.Get("max_snapshots").(int)
	onUpdate, _ := d.Get("snapshot_on_update").(bool)

	var turnedOn []string
	if allow && d.HasChange("allow_snapshots") {
		turnedOn = append(turnedOn, "allow_snapshots")
	}
	if maxSnapshots > 0 && d.HasChange("max_snapshots") {
		turnedOn = append(turnedOn, "max_snapshots")
	}
	if onUpdate {
		turnedOn = append(turnedOn, "snapshot_on_update")
	}
	if len(turnedOn) == 0 {
		return nil
	}

	name, _ := d.Get("name").(string)
	if appType, _ := d.Get("app_type").(string); appType == string(models.AppTypeAPPTYPEDOCKERCOMPOSE) {
		return fmt.Errorf("edge application instance %q is a docker-compose app; snapshots (%v) are not supported "+
			"for docker-compose apps", name, turnedOn)
	}
	if cluster := configuredEdgeNodeClusterID(d); cluster != "" {
		return errSnapshotsOnCluster(name, cluster, fmt.Sprintf("snapshots (%v)", turnedOn))
	}
	if onUpdate && (!allowKnown || !allow) {
		return errors.New("snapshot_on_update requires allow_snapshots = true")
	}
	return nil
}

// normalizeMaxSnapshots keeps the controller's stored maxSnapshots when the
// configuration doesn't set it. The controller reads a stored 0 back as 1, so
// the 1 in state may stand for 0; sending it would change what the controller
// stores (and make it reject app-update snapshots while allowSnapshots is
// false). Values other than 1 are unambiguous and are sent as they are.
func normalizeMaxSnapshots(d *schema.ResourceData, model *models.AppInstance) {
	if model.MaxSnapshots != 1 || d.HasChange("max_snapshots") {
		return
	}
	raw := d.GetRawConfig()
	if !raw.IsNull() && raw.IsKnown() && raw.Type().IsObjectType() && raw.Type().HasAttribute("max_snapshots") &&
		!raw.GetAttr("max_snapshots").IsNull() {
		return
	}
	model.MaxSnapshots = 0
}

// applicationInstanceCustomizeDiff fails the plan for snapshot settings the
// controller would reject, and drops interface ACL changes the controller
// would ignore (UE-203).
func applicationInstanceCustomizeDiff(ctx context.Context, d *schema.ResourceDiff, m interface{}) error {
	if err := checkAppInstanceSnapshotSettings(d, d.NewValueKnown("allow_snapshots")); err != nil {
		return err
	}
	return gateInterfaceACLDiffs(ctx, d, m)
}

// updateApplicationInstanceWithSnapshot sends the update with
// ?makeSnapshot=true on the v2 route, whatever async_writes says: the v1 route
// ignores the parameter (§4.7).
//
// The controller takes the snapshot before it stores the new allowSnapshots,
// so when this same update turns allow_snapshots on, that is sent first on its
// own, and the snapshot still captures the pre-update config.
func updateApplicationInstanceWithSnapshot(
	ctx context.Context, d *schema.ResourceData, client *api_client.ZedcloudAPI, model *models.AppInstance,
) diag.Diagnostics {
	id := d.Id()
	name := d.Get("name").(string)
	timeout := d.Timeout(schema.TimeoutUpdate)

	if oldAllow, newAllow := d.GetChange("allow_snapshots"); !oldAllow.(bool) && newAllow.(bool) {
		current, err := readAppInstanceForSnapshot(ctx, client, id)
		if err != nil {
			return diag.Errorf("edge application instance %q: turning on allow_snapshots: %s", name, err)
		}
		current.AllowSnapshots = true
		current.MaxSnapshots = model.MaxSnapshots // already normalized by the caller
		jobID, err := client.ApplicationInstanceV2.Update(ctx, id, current)
		if err != nil {
			return appInstanceAPIErrorDiags("update (allow_snapshots)", err)
		}
		if _, err := appInstancePoller(client).wait(ctx, jobID, timeout); err != nil {
			return diag.Errorf("edge application instance %q: turning on allow_snapshots: %s", name, err)
		}
		if rev := currentApplicationInstanceRevision(ctx, client, id); rev != nil {
			model.Revision = rev
		}
	}

	jobID, err := client.ApplicationInstanceV2.UpdateWithOptions(ctx, id, model,
		application_instance_v2.UpdateOptions{MakeSnapshot: true})
	if err != nil {
		return appInstanceAPIErrorDiags("update", err)
	}
	if _, err := appInstancePoller(client).wait(ctx, jobID, timeout); err != nil {
		return diag.Errorf("edge application instance %q update with snapshot_on_update: %s. "+
			"If taking the snapshot failed, the update was not applied", name, err)
	}
	return nil
}
