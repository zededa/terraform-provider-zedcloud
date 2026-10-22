package resources

import (
	"context"
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	snapshots "github.com/zededa/terraform-provider-zedcloud/v2/client/app_instance_snapshot"
	zschema "github.com/zededa/terraform-provider-zedcloud/v2/schemas"
)

// AppInstanceSnapshotsDataSource lists an edge application instance's
// snapshots, as zcli's `edge-app-instance snapshot-show`. See the design doc
// §4.5.
func AppInstanceSnapshotsDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Lists the snapshots of an edge application instance, immediate and app-update ones.",
		ReadContext: ReadAppInstanceSnapshots,
		Schema:      zschema.AppInstanceSnapshotsDataSourceSchema(),
	}
}

func ReadAppInstanceSnapshots(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*api_client.ZedcloudAPI)
	appInstanceID := d.Get("app_instance_id").(string)
	nameFilter := d.Get("name").(string)

	list, err := client.AppInstanceSnapshot.List(ctx, appInstanceID)
	if err != nil {
		return snapshotAPIErrorDiags(fmt.Sprintf("listing snapshots of edge application instance %s", appInstanceID), err)
	}

	out := make([]interface{}, 0, len(list))
	for _, s := range list {
		if s == nil || (nameFilter != "" && s.Name != nameFilter) {
			continue
		}
		// Get and List never fill the state; ask ganges.
		state := ""
		st, err := client.AppInstanceSnapshot.State(ctx, s.ID)
		switch {
		case err == nil:
			state = st.Name()
		case errors.Is(err, snapshots.ErrNotFound):
		default:
			return snapshotAPIErrorDiags(fmt.Sprintf("reading the state of snapshot %q", s.Name), err)
		}
		out = append(out, map[string]interface{}{"id": s.ID, "name": s.Name, "state": state})
	}

	if err := d.Set("snapshots", out); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(appInstanceID)
	return nil
}
