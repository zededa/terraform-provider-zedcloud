package resources

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	config "github.com/zededa/terraform-provider-zedcloud/v2/client/application"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
	zschema "github.com/zededa/terraform-provider-zedcloud/v2/schemas"
)

// Default timeouts for zedcloud_application. Bundle deletes used to be
// processed one at a time by seine; at Terraform parallelism 160 the last of
// them took ~2 minutes (ENG-3015). There is no async route for bundles, so
// the provider waits on the synchronous call (UE-177 §4.7).
const (
	applicationWriteTimeout = 20 * time.Minute
	applicationReadTimeout  = 5 * time.Minute
)

func ApplicationResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: CreateApplication,
		ReadContext:   resourceReadApplication,
		UpdateContext: UpdateApplication,
		DeleteContext: DeleteApplication,
		Schema:        zschema.Application(),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Timeouts: &schema.ResourceTimeout{
			Create:  schema.DefaultTimeout(applicationWriteTimeout),
			Read:    schema.DefaultTimeout(applicationReadTimeout),
			Update:  schema.DefaultTimeout(applicationWriteTimeout),
			Delete:  schema.DefaultTimeout(applicationWriteTimeout),
			Default: schema.DefaultTimeout(applicationWriteTimeout),
		},
	}
}

func ApplicationDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: ReadApplication,
		Schema:      zschema.Application(),
	}
}

func applicationAPIErrorDiags(op string, err error) diag.Diagnostics {
	log.Printf("[TRACE] edge application %s error: %s", op, spew.Sdump(err))
	if ds, ok := ZsrvResponderToDiags(err); ok {
		return ds
	}
	return diag.Errorf("edge application %s error: %s", op, err)
}

func CreateApplication(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	model := zschema.ApplicationModel(d)

	// we support reading the manifest from a JSON file for legacy reasons
	if path, ok := d.GetOk("manifest_file"); ok {
		manifest, err := readManifestFromFile(path.(string))
		if err != nil {
			diags = append(diags, diag.FromErr(err)...)
			return diags
		}
		model.Manifest = manifest
		log.Println("[WARN] note, diffs for manifest files are suppressed and cannot be considered by the terraform plan command!")
	}

	// the ac_kind depends on the app_type set in the manifest.
	// this should be read-only and be computed in the backend but is not, so we have to compute it in the client.
	if model.Manifest != nil {
		appType := getAppType(model.Manifest.AppType)
		model.Manifest.AppType = &appType

		deploymentType := getDeploymentType(model.Manifest.DeploymentType)
		model.Manifest.DeploymentType = &deploymentType

		acKind, err := getAcKind(model.Manifest.AppType)
		if err != nil {
			diags = append(diags, diag.FromErr(err)...)
			return diags
		}
		model.Manifest.AcKind = acKind
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: manifest or manifest_file")...)
		return diags
	}

	params := config.CreateParams().WithContext(ctx)
	params.SetBody(model)

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.Application.Create(params, nil)
	if err != nil {
		return applicationAPIErrorDiags("create", err)
	}

	responseData := resp.GetPayload()
	if ds := zsrvResponseDiags(responseData); ds.HasError() {
		return ds
	}

	// note, we need to set the ID in any case, even GetByName endpoint seems to require the ID
	// but doesn't return any error if it's not set.
	d.SetId(responseData.ObjectID)

	// the zedcloud API does not return the partially updated object but a custom response.
	// thus, we need to fetch the object and populate the state.
	return readApplicationAfterWrite(ctx, d, m)
}

// fetchApplication reads the edge app by name when name is set, else by id.
func fetchApplication(ctx context.Context, d *schema.ResourceData, m interface{}) (*models.Application, error) {
	client := m.(*api_client.ZedcloudAPI)

	if name, ok := d.GetOk("name"); ok {
		params := config.GetByNameParams().WithContext(ctx)
		params.Name = name.(string)
		resp, err := client.Application.GetByName(params, nil)
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
	resp, err := client.Application.GetByID(params, nil)
	if err != nil {
		return nil, err
	}
	return resp.GetPayload(), nil
}

// ReadApplication is the data source read: a missing edge app is an error.
func ReadApplication(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	app, err := fetchApplication(ctx, d, m)
	if err != nil {
		return applicationAPIErrorDiags("read", err)
	}
	if app == nil || app.ID == "" {
		return diag.Errorf("edge application read error: the controller returned no edge application")
	}
	zschema.SetApplicationResourceData(d, app)
	d.SetId(app.ID)
	return nil
}

// resourceReadApplication is the resource refresh: an edge app that no longer
// exists is removed from state (UE-177 §4.2).
//
// Tags are only tracked in state once Terraform manages them (they are in
// state before the read, i.e. they came from config; see managed.go). Tags set
// outside Terraform (UI, API) on an app whose config has no tags argument stay
// out of state, so they cause no diff, and UpdateApplication carries them over.
func resourceReadApplication(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	tagsManaged := isManagedMap(d, "tags")
	app, err := fetchApplication(ctx, d, m)
	if err != nil {
		if isStatusNotFound(err) {
			log.Printf("[WARN] edge application %q (%s) not found, removing it from state", d.Get("name"), d.Id())
			d.SetId("")
			return nil
		}
		return applicationAPIErrorDiags("read", err)
	}
	if app == nil || app.ID == "" {
		log.Printf("[WARN] edge application %q (%s): the controller returned no edge application, removing it from state", d.Get("name"), d.Id())
		d.SetId("")
		return nil
	}
	zschema.SetApplicationResourceData(d, app)
	d.SetId(app.ID)
	return unsetUnmanagedAppTags(d, tagsManaged)
}

// unsetUnmanagedAppTags drops tags from state when Terraform doesn't manage
// them. An empty map, not nil: the SDK's map writer only clears previously
// written keys for a non-nil value.
func unsetUnmanagedAppTags(d *schema.ResourceData, managed bool) diag.Diagnostics {
	if managed {
		return nil
	}
	if err := d.Set("tags", map[string]interface{}{}); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

// readApplicationAfterWrite reads back the edge app Create or Update just
// wrote, retrying not-found for a while. It never removes the resource from
// state (UE-177 §4.3).
func readApplicationAfterWrite(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	wantID := d.Id()
	tagsManaged := isManagedMap(d, "tags") // the planned value: tags are in config
	err := resource.RetryContext(ctx, readAfterWriteTimeout, func() *resource.RetryError {
		app, err := fetchApplication(ctx, d, m)
		if err != nil {
			if isStatusNotFound(err) {
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		if app == nil || app.ID == "" {
			return resource.RetryableError(errors.New("the controller returned no edge application yet"))
		}
		if wantID != "" && app.ID != wantID {
			return resource.NonRetryableError(fmt.Errorf(
				"name %q belongs to another edge application (id %s, expected %s)", d.Get("name"), app.ID, wantID))
		}
		zschema.SetApplicationResourceData(d, app)
		d.SetId(app.ID)
		return nil
	})
	if err != nil {
		return applicationAPIErrorDiags("read after write", err)
	}
	return unsetUnmanagedAppTags(d, tagsManaged)
}

func UpdateApplication(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	model := zschema.ApplicationModel(d)

	// we support reading the manifest from a JSON file for legacy reasons
	if path, ok := d.GetOk("manifest_file"); ok {
		manifest, err := readManifestFromFile(path.(string))
		if err != nil {
			diags = append(diags, diag.FromErr(err)...)
			return diags
		}
		model.Manifest = manifest
		log.Println("[WARN] note, diffs for manifest files are suppressed and cannot be considered by the terraform plan command!")
	}
	if model.Manifest == nil {
		return diag.Errorf("missing client parameter: manifest or manifest_file")
	}

	// the ac_kind depends on the app_type set in the manifest.
	// this should be read-only and be computed in the backend but is not, so we have to compute it in the client.
	acKind, err := getAcKind(model.Manifest.AppType)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}
	model.Manifest.AcKind = acKind

	id := d.Id()
	if id == "" {
		return diag.Errorf("missing client parameter: id")
	}
	params := config.UpdateParams().WithContext(ctx)
	params.ID = id
	client := m.(*api_client.ZedcloudAPI)

	// The update replaces the whole app and the controller treats a missing
	// tags field as "clear". When Terraform doesn't manage tags (not in config
	// before or after this change), send the tags the app has now so that tags
	// set in the UI or through the API survive the apply.
	oldTags, newTags := d.GetChange("tags")
	if len(asMap(oldTags)) == 0 && len(asMap(newTags)) == 0 {
		current, err := getApplicationByID(ctx, client, params.ID)
		if err != nil {
			return applicationAPIErrorDiags("read before update", err)
		}
		model.Tags = current.Tags
	}
	params.SetBody(model)

	// makes a bulk update for all properties that were changed
	resp, err := client.Application.Update(params, nil)
	if err != nil {
		return applicationAPIErrorDiags("update", err)
	}

	if ds := zsrvResponseDiags(resp.GetPayload()); ds.HasError() {
		return ds
	}

	// the zedcloud API does not return the partially updated object but a custom response.
	// thus, we need to fetch the object and populate the state.
	return readApplicationAfterWrite(ctx, d, m)
}

func getApplicationByID(ctx context.Context, client *api_client.ZedcloudAPI, id string) (*models.Application, error) {
	params := config.GetByIDParams().WithContext(ctx)
	params.ID = id
	resp, err := client.Application.GetByID(params, nil)
	if err != nil {
		return nil, err
	}
	return resp.GetPayload(), nil
}

func DeleteApplication(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	id := d.Id()
	if id == "" {
		return diag.Errorf("missing client parameter: id")
	}

	params := config.DeleteParams().WithContext(ctx)
	params.ID = id

	client := m.(*api_client.ZedcloudAPI)

	// A delete that timed out here may still complete on the controller. It
	// is not retried (UE-177 §4.1); the next run gets 404, which is success,
	// or refresh drops the resource.
	if _, err := client.Application.Delete(params, nil); err != nil {
		if isStatusNotFound(err) {
			d.SetId("")
			return nil
		}
		return applicationAPIErrorDiags("delete", err)
	}

	d.SetId("")
	return nil
}

func readManifestFromFile(path string) (*models.VMManifest, error) {
	jsonData, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("could not read manifest file (path=%s): %w", path, err)
	}
	var manifest models.VMManifest
	err = json.Unmarshal([]byte(jsonData), &manifest)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal manifest file (path=%s): %w", path, err)
	}
	return &manifest, nil
}

func getAcKind(appType *models.AppType) (*string, error) {
	if appType == nil {
		return nil, nil
	}
	switch *appType {
	case models.AppTypeAPPTYPEVM:
		acKind := "VMManifest"
		return &acKind, nil
	case models.AppTypeAPPTYPECONTAINER:
		acKind := "PodManifest"
		return &acKind, nil
	case models.AppTypeAPPTYPEMODULE:
		acKind := "ModuleManifest"
		return &acKind, nil
	case models.AppTypeAPPTYPEDOCKERCOMPOSE:
		acKind := "ComposeManifest"
		return &acKind, nil
	default:
		return nil, fmt.Errorf("could not determine ac_kind, unsupported app_type (%s)", *appType)
	}
}

func getAppType(appType *models.AppType) models.AppType {
	if appType == nil || *appType == "" {
		return models.AppTypeAPPTYPEUNSPECIFIED
	}
	return *appType
}

func getDeploymentType(deploymentType *models.DeploymentType) models.DeploymentType {
	if deploymentType == nil || *deploymentType == "" {
		return models.DeploymentTypeDEPLOYMENTTYPEUNSPECIFIED
	}
	return *deploymentType
}
