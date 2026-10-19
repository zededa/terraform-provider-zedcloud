package resources

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	project "github.com/zededa/terraform-provider-zedcloud/v2/client/projects"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
	zschema "github.com/zededa/terraform-provider-zedcloud/v2/schemas"
)

func ProjectResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: CreateProject,
		ReadContext:   GetProject,
		UpdateContext: UpdateProject,
		DeleteContext: DeleteProject,
		CustomizeDiff: customizeProjectDiff,
		Schema:        zschema.Project(),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func ProjectDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: GetProject,
		Schema:      zschema.Project(),
	}
}

func CreateProject(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	model := zschema.TagModel(d)
	setProjectPCRPolicyVersion(d, model, true)
	params := project.CreateParams()
	params.SetBody(model)

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.Project.Create(params, nil)
	if err != nil {
		log.Printf("[TRACE] project create error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("project create error: %s", err)...)
		return diags
	}

	responseData := resp.GetPayload()
	if responseData != nil && len(responseData.Error) > 0 {
		for _, err := range responseData.Error {
			// FIXME: zedcloud api returns a response that contains and error even in case of success.
			// remove this code once it is fixed on API side.
			if err.ErrorCode != nil && *err.ErrorCode == models.ErrorCodeSuccess {
				continue
			}
			diags = append(diags, diag.FromErr(errors.New(err.Details))...)
		}
		if diags.HasError() {
			return diags
		}
	}

	// note, we need to set the ID before fetching the newly created project.
	d.SetId(responseData.ObjectID)

	// the zedcloud API does not return the partially updated object but a custom response.
	// thus, we need to fetch the object and populate the state.
	if errs := GetProject(ctx, d, m); errs != nil {
		return append(diags, errs...)
	}

	return diags
}

func GetProject(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var project *models.Tag
	var diags diag.Diagnostics

	if _, isSet := d.GetOk("name"); isSet {
		project, diags = readProjectByName(ctx, d, m)
	} else if _, isSet := d.GetOk("id"); isSet {
		project, diags = readProjectByID(ctx, d, m)
	}

	if diags.HasError() {
		return diags
	}

	zschema.SetTagResourceData(d, project)
	d.SetId(project.ID)

	return diags
}

func readProjectByID(ctx context.Context, d *schema.ResourceData, m interface{}) (*models.Tag, diag.Diagnostics) {
	var diags diag.Diagnostics

	params := project.GetByIDParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	id, isSet := d.GetOk("id")
	if isSet {
		params.ID = id.(string)
	} else {
		return nil, append(diags, diag.Errorf("missing client parameter: id")...)
	}

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.Project.GetByID(params, nil)
	if err != nil {
		log.Printf("[TRACE] project read error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return nil, diags
		}

		diags = append(diags, diag.Errorf("project read error: %s", err)...)
		return nil, diags
	}

	return resp.GetPayload(), diags
}

func readProjectByName(ctx context.Context, d *schema.ResourceData, m interface{}) (*models.Tag, diag.Diagnostics) {
	var diags diag.Diagnostics

	params := project.GetByNameParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	name, isSet := d.GetOk("name")
	if isSet {
		params.Name = name.(string)
	} else {
		return nil, append(diags, diag.Errorf("missing client parameter: name")...)
	}

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.Project.GetByName(params, nil)
	if err != nil {
		log.Printf("[TRACE] project read error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return nil, diags
		}

		diags = append(diags, diag.Errorf("project read error: %s", err)...)
		return nil, diags
	}

	return resp.GetPayload(), diags
}

func UpdateProject(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := project.UpdateParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	model := zschema.TagModel(d)
	setProjectPCRPolicyVersion(d, model, false)
	params.SetBody(model)

	idVal, idIsSet := d.GetOk("id")
	if idIsSet {
		id, _ := idVal.(string)
		params.ID = id
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: id")...)
		return diags
	}

	client := m.(*api_client.ZedcloudAPI)
	resp, err := client.Project.Update(params, nil)
	if err != nil {
		log.Printf("[TRACE] project update error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("project update error: %s", err)...)
		return diags
	}

	responseData := resp.GetPayload()
	if responseData != nil && len(responseData.Error) > 0 {
		for _, err := range responseData.Error {
			// FIXME: zedcloud api returns a response that contains and error even in case of success.
			// remove this code once it is fixed on API side.
			if err.ErrorCode != nil && *err.ErrorCode == models.ErrorCodeSuccess {
				continue
			}
			diags = append(diags, diag.FromErr(errors.New(err.Details))...)
		}
		if diags.HasError() {
			return diags
		}
	}

	// the zedcloud API does not return the partially updated object but a custom response.
	// thus, we need to fetch the object and populate the state.
	if errs := GetProject(ctx, d, m); errs != nil {
		return append(diags, errs...)
	}

	return diags
}

func DeleteProject(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := project.DeleteParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	idVal, idIsSet := d.GetOk("id")
	if idIsSet {
		id, _ := idVal.(string)
		params.ID = id
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: id")...)
		return diags
	}

	client := m.(*api_client.ZedcloudAPI)

	_, err := client.Project.Delete(params, nil)
	if err != nil {
		if isStatusNotFound(err) {
			d.SetId("")
			return diags
		}
		log.Printf("[TRACE] project delete error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("project delete error: %s", err)...)
		return diags
	}

	d.SetId("")
	return diags
}

// customizeProjectDiff validates pcr_policy at plan time: the policy type must
// be POLICY_TYPE_PCR, the nested pcr_policy block must be present, and the
// indices must satisfy the controller's rules (PCR 0 required, no PCR 5, 0-15).
func customizeProjectDiff(ctx context.Context, d *schema.ResourceDiff, m interface{}) error {
	pol := asList(d.Get("pcr_policy"))
	if len(pol) == 0 || pol[0] == nil {
		return nil
	}
	pm := pol[0].(map[string]interface{})
	if t, _ := pm["type"].(string); t != "" && t != string(models.PolicyTypePOLICYTYPEPCR) {
		return fmt.Errorf("pcr_policy.type must be %q, got %q", models.PolicyTypePOLICYTYPEPCR, t)
	}
	inner := asList(pm["pcr_policy"])
	if len(inner) == 0 || inner[0] == nil {
		return fmt.Errorf("pcr_policy requires a nested pcr_policy { pcr_indices = [...] } block")
	}
	if !d.NewValueKnown("pcr_policy") {
		return nil
	}
	indices := zschema.PCRIndicesFromInterface(inner[0].(map[string]interface{})["pcr_indices"])
	if err := zschema.ValidatePCRIndices(int64sToInts(indices)); err != nil {
		return fmt.Errorf("pcr_policy.pcr_policy.pcr_indices: %w", err)
	}
	return nil
}

// setProjectPCRPolicyVersion fills in the type and version of the project PCR
// policy before it is sent. The controller doesn't manage the version for
// project policies (zcli sends whatever is in its file), so the provider does:
// 1 on create, +1 when the indices change, unchanged otherwise.
func setProjectPCRPolicyVersion(d *schema.ResourceData, model *models.Tag, create bool) {
	if model.PcrPolicy == nil || model.PcrPolicy.PcrPolicy == nil {
		return
	}
	if model.PcrPolicy.Type == nil || *model.PcrPolicy.Type == "" {
		model.PcrPolicy.Type = models.NewPolicyType(models.PolicyTypePOLICYTYPEPCR)
	}
	newIndices := model.PcrPolicy.PcrPolicy.PcrIndices
	if create {
		model.PcrPolicy.PcrPolicy.PolicyVersion = 1
		return
	}
	oldVal, _ := d.GetChange("pcr_policy")
	var oldIndices []int64
	var oldVersion int64
	if old := asList(oldVal); len(old) > 0 && old[0] != nil {
		if inner := asList(old[0].(map[string]interface{})["pcr_policy"]); len(inner) > 0 && inner[0] != nil {
			im := inner[0].(map[string]interface{})
			oldIndices = zschema.PCRIndicesFromInterface(im["pcr_indices"])
			if v, ok := im["policy_version"].(int); ok {
				oldVersion = int64(v)
			}
		}
	}
	model.PcrPolicy.PcrPolicy.PolicyVersion = zschema.NextPCRPolicyVersion(oldIndices, oldVersion, newIndices)
}

func int64sToInts(in []int64) []int {
	out := make([]int, len(in))
	for i, v := range in {
		out[i] = int(v)
	}
	return out
}
