package resources

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	identity_access_management "github.com/zededa/terraform-provider-zedcloud/v2/client/identity_access_management"
	zschema "github.com/zededa/terraform-provider-zedcloud/v2/schemas"
)

/*
IdentityAccessManagement identity access management API
*/

func RoleResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: IdentityAccessManagement_CreateRole,
		DeleteContext: IdentityAccessManagement_DeleteRole,
		ReadContext:   IdentityAccessManagement_GetRole,
		UpdateContext: IdentityAccessManagement_UpdateRole,
		Schema:        zschema.RoleSchema(),
		Importer: &schema.ResourceImporter{
			StateContext: importRoleStateContext,
		},
	}
}

// importRoleStateContext accepts either the 28-character role ID or the role
// name as the import ID (NFR-165 §4.1):
//
//	terraform import zedcloud_role.readonly readonly-operators
//	terraform import zedcloud_role.readonly CCGFABAEqnH4je5PHZTXSmHOs-ZE
func importRoleStateContext(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	raw := strings.TrimSpace(d.Id())
	if raw == "" {
		return nil, errors.New("import ID is empty: pass either the role ID or the role name")
	}

	if looksLikeObjectID(raw) {
		d.SetId(raw)
		return []*schema.ResourceData{d}, nil
	}

	params := identity_access_management.NewIdentityAccessManagementGetRoleByNameParams()
	params.Name = raw

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.IdentityAccessManagement.IdentityAccessManagementGetRoleByName(params, nil)
	if err != nil {
		if isStatusNotFound(err) {
			return nil, fmt.Errorf("no role found with name %q", raw)
		}
		return nil, fmt.Errorf("could not resolve role name %q: %w", raw, err)
	}

	role := resp.GetPayload()
	if role == nil || role.ID == "" {
		return nil, fmt.Errorf("role %q resolved to an empty ID", raw)
	}

	d.SetId(role.ID)
	return []*schema.ResourceData{d}, nil
}

func RoleDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext:   IdentityAccessManagement_GetRole,
		Schema: zschema.RoleSchema(),
	}
}

func IdentityAccessManagement_GetRole(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	var role *models.Role

	// NFR-165 §3.2: the immutable system ID wins over the name whenever state
	// has one, so a renamed role is still found rather than reported missing.
	id := d.Id()
	if id == "" {
		if idVal, isSet := d.GetOk("id"); isSet {
			id, _ = idVal.(string)
		}
	}

	switch {
	case id != "":
		var found bool
		role, found, diags = getRoleById(ctx, d, m, id)
		if diags.HasError() {
			return diags
		}
		// NFR-165 §3.3: deleted out of band — plan a re-create.
		if !found {
			log.Printf("[INFO] role %s no longer exists, removing from state", id)
			d.SetId("")
			return diags
		}
	case hasName(d):
		role, diags = getRoleByName(ctx, d, m)
		if diags.HasError() {
			return diags
		}
	default:
		// NFR-165 §3.1.
		return append(diags, diag.Errorf("cannot read role: set either id or name")...)
	}

	if role == nil {
		return append(diags, diag.Errorf("role read returned an empty response")...)
	}

	zschema.SetRoleResourceData(d, role)
	d.SetId(role.ID)

	return diags
}

func hasName(d *schema.ResourceData) bool {
	_, isSet := d.GetOk("name")
	return isSet
}

// getRoleById fetches a role by its system ID. The boolean result is false
// when the API answered 404.
func getRoleById(ctx context.Context, d *schema.ResourceData, m interface{}, id string) (*models.Role, bool, diag.Diagnostics) {
	var diags diag.Diagnostics

	params := identity_access_management.NewIdentityAccessManagementGetRoleParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	params.ID = id

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.IdentityAccessManagement.IdentityAccessManagementGetRole(params, nil)
	if err != nil {
		if isStatusNotFound(err) {
			return nil, false, diags
		}

		log.Printf("[TRACE] role read error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return nil, false, diags
		}

		diags = append(diags, diag.Errorf("role read error: %s", err)...)
		return nil, false, diags
	}

	return resp.GetPayload(), true, diags
}

func getRoleByName(ctx context.Context, d *schema.ResourceData, m interface{}) (*models.Role, diag.Diagnostics) {
	var diags diag.Diagnostics

	params := identity_access_management.NewIdentityAccessManagementGetRoleByNameParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	nameVal, nameIsSet := d.GetOk("name")
	if nameIsSet {
		params.Name = nameVal.(string)
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: name")...)
		return nil, diags
	}

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.IdentityAccessManagement.IdentityAccessManagementGetRoleByName(params, nil)
	if err != nil {
		log.Printf("[TRACE] role read error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return nil, diags
		}

		diags = append(diags, diag.Errorf("role read error: %s", err)...)
		return nil, diags
	}

	return resp.GetPayload(), diags
}

func IdentityAccessManagement_CreateRole(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	model := zschema.RoleModel(d)
	params := identity_access_management.NewIdentityAccessManagementCreateRoleParams()
	params.SetBody(model)

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.IdentityAccessManagement.IdentityAccessManagementCreateRole(params, nil)
	if err != nil {
		log.Printf("[TRACE] role create error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("role create error: %s", err)...)
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
	if diags := IdentityAccessManagement_GetRole(ctx, d, m); diags.HasError() {
		return diags
	}

	return diags
}

func IdentityAccessManagement_UpdateRole(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	d.Partial(true)

	params := identity_access_management.NewIdentityAccessManagementUpdateRoleParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	params.SetBody(zschema.RoleModel(d))
	// IdentityAccessManagementUpdateRoleBody

	// d.Id() is authoritative for a managed resource, created or imported.
	if id := d.Id(); id != "" {
		params.ID = id
	} else if idVal, isSet := d.GetOk("id"); isSet {
		params.ID, _ = idVal.(string)
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: id")...)
		return diags
	}

	// makes a bulk update for all properties that were changed
	client := m.(*api_client.ZedcloudAPI)
	resp, err := client.IdentityAccessManagement.IdentityAccessManagementUpdateRole(params, nil)
	if err != nil {
		log.Printf("[TRACE] role update error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("role update error: %s", err)...)
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
	if diags := IdentityAccessManagement_GetRole(ctx, d, m); diags.HasError() {
		return diags
	}

	d.Partial(false)

	return diags
}

func IdentityAccessManagement_DeleteRole(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := identity_access_management.NewIdentityAccessManagementDeleteRoleParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	if id := d.Id(); id != "" {
		params.ID = id
	} else if idVal, isSet := d.GetOk("id"); isSet {
		params.ID, _ = idVal.(string)
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: id")...)
		return diags
	}

	client := m.(*api_client.ZedcloudAPI)

	_, err := client.IdentityAccessManagement.IdentityAccessManagementDeleteRole(params, nil)
	if err != nil {
		// Already gone: a delete of a missing object is a success.
		if isStatusNotFound(err) {
			d.SetId("")
			return diags
		}

		log.Printf("[TRACE] role delete error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("role delete error: %s", err)...)
		return diags
	}

	d.SetId("")
	return diags
}
