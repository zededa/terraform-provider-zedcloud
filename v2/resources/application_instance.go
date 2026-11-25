package resources

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	config "github.com/zededa/terraform-provider-zedcloud/v2/client/application_instance"
	clustercfg "github.com/zededa/terraform-provider-zedcloud/v2/client/cluster"
	"github.com/zededa/terraform-provider-zedcloud/v2/client/jobs"
	nodecfg "github.com/zededa/terraform-provider-zedcloud/v2/client/node"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
	zschema "github.com/zededa/terraform-provider-zedcloud/v2/schemas"
)

// Default timeouts for zedcloud_application_instance. At high parallelism a
// single device's instance writes queue on the controller, so creates and
// updates can take minutes (ENG-3015, UE-177).
const (
	appInstanceWriteTimeout = 20 * time.Minute
	appInstanceReadTimeout  = 5 * time.Minute
)

// readAfterWriteTimeout bounds how long Create and Update wait for the
// instance they just wrote to become readable.
var readAfterWriteTimeout = 2 * time.Minute

func ApplicationInstanceResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: CreateApplicationInstance,
		ReadContext:   resourceReadApplicationInstance,
		UpdateContext: UpdateApplicationInstance,
		DeleteContext: DeleteApplicationInstance,
		CustomizeDiff: applicationInstanceCustomizeDiff,
		Schema:        zschema.ApplicationInstanceResourceSchema(),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Timeouts: &schema.ResourceTimeout{
			Create:  schema.DefaultTimeout(appInstanceWriteTimeout),
			Read:    schema.DefaultTimeout(appInstanceReadTimeout),
			Update:  schema.DefaultTimeout(appInstanceWriteTimeout),
			Delete:  schema.DefaultTimeout(appInstanceWriteTimeout),
			Default: schema.DefaultTimeout(appInstanceWriteTimeout),
		},
	}
}

func ApplicationInstanceDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: ReadApplicationInstance,
		Schema:      zschema.ApplicationInstance(),
	}
}

func appInstancePoller(c *api_client.ZedcloudAPI) *jobPoller {
	s := c.Settings
	if s == nil {
		s = api_client.DefaultSettings()
	}
	return &jobPoller{
		client:      c.Jobs,
		objectType:  jobs.ObjectTypeEdgeAppInstance,
		intervalMin: s.JobPollIntervalMin,
		intervalMax: s.JobPollIntervalMax,
	}
}

func asyncWrites(c *api_client.ZedcloudAPI) bool {
	return c.Settings == nil || c.Settings.AsyncWrites
}

// appInstanceAPIErrorDiags turns an API error into diagnostics.
func appInstanceAPIErrorDiags(op string, err error) diag.Diagnostics {
	log.Printf("[TRACE] edge application instance %s error: %s", op, spew.Sdump(err))
	if ds, ok := ZsrvResponderToDiags(err); ok {
		return ds
	}
	return diag.Errorf("edge application instance %s error: %s", op, err)
}

// zsrvResponseDiags returns the errors carried in a 200 ZsrvResponse.
func zsrvResponseDiags(resp *models.ZsrvResponse) diag.Diagnostics {
	var diags diag.Diagnostics
	if resp == nil {
		return diags
	}
	for _, e := range resp.Error {
		// FIXME: zedcloud api returns a response that contains and error even in case of success.
		// remove this code once it is fixed on API side.
		if e == nil || (e.ErrorCode != nil && *e.ErrorCode == models.ErrorCodeSuccess) {
			continue
		}
		diags = append(diags, diag.FromErr(errors.New(e.Details))...)
	}
	return diags
}

// isHTTPResponseError reports whether err came from an HTTP response, as
// opposed to a transport failure (timeout, connection reset) after which the
// controller may still have done the work.
func isHTTPResponseError(err error) bool {
	if _, ok := err.(ZsrvResponder); ok {
		return true
	}
	type statusCoder interface{ Code() int }
	var sc statusCoder
	return errors.As(err, &sc)
}

// newAppInstanceID returns a random RFC 4122 version 4 UUID, the form the
// controller generates itself (uuid.New in seine).
func newAppInstanceID() (string, error) {
	var b [16]byte
	if _, err := rand.Read(b[:]); err != nil {
		return "", err
	}
	b[6] = (b[6] & 0x0f) | 0x40 // version 4
	b[8] = (b[8] & 0x3f) | 0x80 // RFC 4122 variant
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:16]), nil
}

func CreateApplicationInstance(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	_, isDevIDSet := d.GetOk("device_id")
	_, isEdgeNodeClusterSet := d.GetOk("edge_node_cluster")
	if !isDevIDSet && !isEdgeNodeClusterSet {
		return diag.Errorf("either device_id or edge_node_cluster has to be specified")
	}

	if err := checkAppInstanceSnapshotSettings(d, true); err != nil {
		return diag.FromErr(err)
	}

	client := m.(*api_client.ZedcloudAPI)
	name := d.Get("name").(string)

	// The provider picks the id, so it knows the instance's id even if the
	// call times out after the controller accepted it. Then the instance is
	// tracked (tainted) in state instead of orphaned (ENG-3016, UE-177 §4.6).
	// seine keeps a caller-supplied UUID (appInstCreateCommon).
	id, err := newAppInstanceID()
	if err != nil {
		return diag.Errorf("generating an id for edge application instance %q: %s", name, err)
	}
	model := zschema.ApplicationInstanceModel(d)
	model.ID = id
	sentInterfaces, _ := d.Get("interfaces").([]interface{})

	if asyncWrites(client) {
		// v1 derives the project from the device; v2 rejects a create
		// without projectId up front (seine appInstCreateV2).
		if err := resolveAppInstanceProjectID(ctx, client, model); err != nil {
			return diag.Errorf("edge application instance %q create: %s", name, err)
		}
		jobID, err := client.ApplicationInstanceV2.Create(ctx, model)
		if err != nil {
			if !isHTTPResponseError(err) {
				// The request may have reached the controller.
				d.SetId(id)
			}
			return appInstanceAPIErrorDiags("create", err)
		}
		d.SetId(id)

		if _, err := appInstancePoller(client).wait(ctx, jobID, d.Timeout(schema.TimeoutCreate)); err != nil {
			var failed *jobFailedError
			if errors.As(err, &failed) && !applicationInstanceExists(ctx, client, id) {
				// Nothing was created; do not leave a tainted resource behind.
				d.SetId("")
			}
			return diag.Errorf("edge application instance %q create: %s", name, err)
		}
	} else {
		params := config.CreateParams().WithContext(ctx)
		params.SetBody(model)

		resp, err := client.ApplicationInstance.Create(params, nil)
		if err != nil {
			if !isHTTPResponseError(err) {
				d.SetId(id)
			}
			return appInstanceAPIErrorDiags("create", err)
		}
		if diags := zsrvResponseDiags(resp.GetPayload()); diags.HasError() {
			return diags
		}
		if objID := resp.GetPayload().ObjectID; objID != "" {
			id = objID
		}
		d.SetId(id)
	}

	// The zedcloud API does not return the object, so read it back.
	diags := readApplicationInstanceAfterWrite(ctx, d, m)
	if diags.HasError() {
		return diags
	}
	return append(diags, aclReadBackDiags(ctx, d, m, sentInterfaces)...)
}

// projectLookups caches, per provider process, the project of each edge node
// or cluster that resolveAppInstanceProjectID looked up. At high
// parallelism many instances share a few devices (CI-1004: 120 instances on
// two), and a read per instance is load the controller does not need.
var projectLookups = &projectCache{entries: map[string]*projectEntry{}}

type projectCache struct {
	mu      sync.Mutex
	entries map[string]*projectEntry
}

type projectEntry struct {
	done      chan struct{}
	projectID string
	err       error
}

// get returns the cached project for key, or calls fetch once for all
// concurrent callers. Failures are not cached.
func (c *projectCache) get(ctx context.Context, key string, fetch func() (string, error)) (string, error) {
	c.mu.Lock()
	if e, ok := c.entries[key]; ok {
		c.mu.Unlock()
		select {
		case <-e.done:
			if e.err == nil {
				return e.projectID, nil
			}
			// The lookup that failed has already been forgotten; try again.
			return c.get(ctx, key, fetch)
		case <-ctx.Done():
			return "", ctx.Err()
		}
	}
	e := &projectEntry{done: make(chan struct{})}
	c.entries[key] = e
	c.mu.Unlock()

	e.projectID, e.err = fetch()
	if e.err != nil || e.projectID == "" {
		if e.err == nil {
			e.err = errors.New("no project")
		}
		c.mu.Lock()
		delete(c.entries, key)
		c.mu.Unlock()
	}
	close(e.done)
	return e.projectID, e.err
}

// resolveAppInstanceProjectID fills model.ProjectID from the edge node, or
// from the edge node cluster, when the configuration does not set
// project_id. The synchronous v1 create does this on the controller; the
// async v2 create requires it in the request.
func resolveAppInstanceProjectID(ctx context.Context, client *api_client.ZedcloudAPI, model *models.AppInstance) error {
	if model.ProjectID != nil && *model.ProjectID != "" {
		return nil
	}
	var (
		projectID string
		err       error
	)
	switch {
	case model.DeviceID != nil && *model.DeviceID != "":
		devID := *model.DeviceID
		projectID, err = projectLookups.get(ctx, "device/"+devID, func() (string, error) {
			params := nodecfg.GetByIDParams().WithContext(ctx)
			params.ID = devID
			resp, err := client.Node.GetByID(params, nil)
			if err != nil {
				return "", fmt.Errorf("reading edge node %s to find its project: %w", devID, err)
			}
			if p := resp.GetPayload(); p != nil && p.ProjectID != nil {
				return *p.ProjectID, nil
			}
			return "", nil
		})
	case model.EdgeNodeCluster != nil && model.EdgeNodeCluster.ID != "":
		clusterID := model.EdgeNodeCluster.ID
		projectID, err = projectLookups.get(ctx, "cluster/"+clusterID, func() (string, error) {
			params := clustercfg.NewGetParams().WithContext(ctx)
			params.ID = clusterID
			resp, err := client.Cluster.GetCluster(params, nil)
			if err != nil {
				return "", fmt.Errorf("reading edge node cluster %s to find its project: %w", clusterID, err)
			}
			if p := resp.GetPayload(); p != nil {
				return p.ProjectID, nil
			}
			return "", nil
		})
	}
	if err != nil && err.Error() != "no project" {
		return err
	}
	if projectID == "" {
		return errors.New("project_id is not set and could not be read from the edge node or cluster")
	}
	model.ProjectID = &projectID
	return nil
}

// fetchApplicationInstance reads the instance by name when name is set, else
// by id.
func fetchApplicationInstance(ctx context.Context, d *schema.ResourceData, m interface{}) (*models.AppInstance, error) {
	client := m.(*api_client.ZedcloudAPI)

	if name, ok := d.GetOk("name"); ok {
		params := config.GetByNameParams().WithContext(ctx)
		params.Name = name.(string)
		resp, err := client.ApplicationInstance.GetByName(params, nil)
		if err != nil {
			return nil, err
		}
		return resp.GetPayload(), nil
	}

	id, ok := d.GetOk("id")
	if !ok || id.(string) == "" {
		return nil, errors.New("missing client parameter: name or id")
	}
	params := config.GetByIDParams().WithContext(ctx)
	params.ID = id.(string)
	resp, err := client.ApplicationInstance.GetByID(params, nil)
	if err != nil {
		return nil, err
	}
	return resp.GetPayload(), nil
}

// ReadApplicationInstance is the data source read: a missing instance is an
// error.
func ReadApplicationInstance(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	appInstance, err := fetchApplicationInstance(ctx, d, m)
	if err != nil {
		return appInstanceAPIErrorDiags("read", err)
	}
	if appInstance == nil || appInstance.ID == "" {
		return diag.Errorf("edge application instance read error: the controller returned no instance")
	}
	zschema.SetApplicationInstanceResourceData(d, appInstance)
	d.SetId(appInstance.ID)
	return nil
}

// resourceReadApplicationInstance is the resource refresh: an instance that no
// longer exists is removed from state, so Terraform plans to create it again
// instead of failing every run until `terraform state rm` (UE-177 §4.2).
func resourceReadApplicationInstance(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	appInstance, err := fetchApplicationInstance(ctx, d, m)
	if err != nil {
		if isStatusNotFound(err) {
			log.Printf("[WARN] edge application instance %q (%s) not found, removing it from state", d.Get("name"), d.Id())
			d.SetId("")
			return nil
		}
		return appInstanceAPIErrorDiags("read", err)
	}
	if appInstance == nil || appInstance.ID == "" {
		log.Printf("[WARN] edge application instance %q (%s): the controller returned no instance, removing it from state", d.Get("name"), d.Id())
		d.SetId("")
		return nil
	}
	zschema.SetApplicationInstanceResourceData(d, appInstance)
	// snapshot_on_update lives in state only; after an import it is unset.
	_ = d.Set("snapshot_on_update", d.Get("snapshot_on_update"))
	d.SetId(appInstance.ID)
	return nil
}

// readApplicationInstanceAfterWrite reads back the instance Create or Update
// just wrote. Under load the write can become visible a moment later, so
// not-found is retried for a while. It never removes the resource from state
// (UE-177 §4.3).
func readApplicationInstanceAfterWrite(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	wantID := d.Id()
	err := resource.RetryContext(ctx, readAfterWriteTimeout, func() *resource.RetryError {
		appInstance, err := fetchApplicationInstance(ctx, d, m)
		if err != nil {
			if isStatusNotFound(err) {
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		if appInstance == nil || appInstance.ID == "" {
			return resource.RetryableError(errors.New("the controller returned no instance yet"))
		}
		if wantID != "" && appInstance.ID != wantID {
			return resource.NonRetryableError(fmt.Errorf(
				"name %q belongs to another edge application instance (id %s, expected %s). "+
					"Import it with `terraform import` or choose another name",
				d.Get("name"), appInstance.ID, wantID))
		}
		zschema.SetApplicationInstanceResourceData(d, appInstance)
		d.SetId(appInstance.ID)
		return nil
	})
	if err != nil {
		return appInstanceAPIErrorDiags("read after write", err)
	}
	return nil
}

// applicationInstanceExists reports whether an instance with id exists. It
// answers true when it cannot tell, so callers keep tracking the resource.
func applicationInstanceExists(ctx context.Context, client *api_client.ZedcloudAPI, id string) bool {
	params := config.GetByIDParams().WithContext(ctx)
	params.ID = id
	_, err := client.ApplicationInstance.GetByID(params, nil)
	if err == nil {
		return true
	}
	return !isStatusNotFound(err)
}

func UpdateApplicationInstance(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*api_client.ZedcloudAPI)
	id := d.Id()
	if id == "" {
		return diag.Errorf("missing client parameter: id")
	}

	// snapshot_on_update lives in state only: changing just that must not
	// send an update (which would itself take a snapshot).
	if !d.HasChangesExcept("snapshot_on_update") {
		return nil
	}
	if err := checkAppInstanceSnapshotSettings(d, true); err != nil {
		return diag.FromErr(err)
	}

	model := zschema.ApplicationInstanceModel(d)
	model.ID = id
	normalizeMaxSnapshots(d, model)
	sentInterfaces, _ := d.Get("interfaces").([]interface{})

	// CI-752 / CI-790: lock on the server's current revision, not the state's.
	// See the comment on currentVolumeInstanceRevision.
	if rev := currentApplicationInstanceRevision(ctx, client, id); rev != nil {
		model.Revision = rev
	}

	if snapshotOnUpdate, _ := d.Get("snapshot_on_update").(bool); snapshotOnUpdate {
		if diags := updateApplicationInstanceWithSnapshot(ctx, d, client, model); diags.HasError() {
			return diags
		}
	} else if asyncWrites(client) {
		jobID, err := client.ApplicationInstanceV2.Update(ctx, id, model)
		if err != nil {
			return appInstanceAPIErrorDiags("update", err)
		}
		if _, err := appInstancePoller(client).wait(ctx, jobID, d.Timeout(schema.TimeoutUpdate)); err != nil {
			return diag.Errorf("edge application instance %q update: %s", d.Get("name"), err)
		}
	} else {
		params := config.UpdateParams().WithContext(ctx)
		params.ID = id
		params.SetBody(model)

		resp, err := client.ApplicationInstance.Update(params, nil)
		if err != nil {
			return appInstanceAPIErrorDiags("update", err)
		}
		if diags := zsrvResponseDiags(resp.GetPayload()); diags.HasError() {
			return diags
		}
	}

	diags := readApplicationInstanceAfterWrite(ctx, d, m)
	if diags.HasError() {
		return diags
	}
	return append(diags, aclReadBackDiags(ctx, d, m, sentInterfaces)...)
}

// currentApplicationInstanceRevision fetches the app instance's revision as the
// server holds it now. nil means "could not determine"; see
// currentVolumeInstanceRevision.
func currentApplicationInstanceRevision(ctx context.Context, client *api_client.ZedcloudAPI, id string) *models.ObjectRevision {
	params := config.GetByIDParams().WithContext(ctx)
	params.ID = id

	resp, err := client.ApplicationInstance.GetByID(params, nil)
	if err != nil {
		log.Printf("[TRACE] application instance revision read error: %s", spew.Sdump(err))
		return nil
	}
	if payload := resp.GetPayload(); payload != nil {
		return payload.Revision
	}
	return nil
}

func DeleteApplicationInstance(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*api_client.ZedcloudAPI)
	id := d.Id()
	if id == "" {
		return diag.Errorf("missing client parameter: id")
	}

	if asyncWrites(client) {
		jobID, err := client.ApplicationInstanceV2.Delete(ctx, id)
		if err != nil {
			if isStatusNotFound(err) {
				d.SetId("")
				return nil
			}
			return appInstanceAPIErrorDiags("delete", err)
		}
		// A COMPLETED delete job is authoritative; no read-back needed.
		if _, err := appInstancePoller(client).wait(ctx, jobID, d.Timeout(schema.TimeoutDelete)); err != nil {
			return diag.Errorf("edge application instance %q delete: %s", d.Get("name"), err)
		}
		d.SetId("")
		return nil
	}

	params := config.DeleteParams().WithContext(ctx)
	params.ID = id
	if _, err := client.ApplicationInstance.Delete(params, nil); err != nil {
		if isStatusNotFound(err) {
			d.SetId("")
			return nil
		}
		return appInstanceAPIErrorDiags("delete", err)
	}

	d.SetId("")
	return nil
}
