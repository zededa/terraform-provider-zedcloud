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
	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	identity_access_management "github.com/zededa/terraform-provider-zedcloud/v2/client/identity_access_management"
	models "github.com/zededa/terraform-provider-zedcloud/v2/models"
	zschema "github.com/zededa/terraform-provider-zedcloud/v2/schemas"
)

/*
User management API
*/

func UserResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: IdentityAccessManagement_CreateUser,
		DeleteContext: IdentityAccessManagement_DeleteUser,
		ReadContext:   IdentityAccessManagement_GetUser,
		UpdateContext: IdentityAccessManagement_UpdateUser2,
		Schema:        zschema.DetailedUserSchema(),
		Importer: &schema.ResourceImporter{
			StateContext: importUserStateContext,
		},
	}
}

// importUserStateContext accepts either the 28-character user ID or the
// username as the import ID (NFR-165 §4.1):
//
//	terraform import zedcloud_user.alice alice@corp.com
//	terraform import zedcloud_user.alice AAGFABAEqnH4je5PHZTXSmHOs-XC
//
// A name is resolved to its ID here so that the Read which follows takes the
// by-ID path and state ends up keyed on the immutable identifier.
func importUserStateContext(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	raw := strings.TrimSpace(d.Id())
	if raw == "" {
		return nil, errors.New("import ID is empty: pass either the user ID or the username")
	}

	if looksLikeObjectID(raw) {
		d.SetId(raw)
		return []*schema.ResourceData{d}, nil
	}

	params := identity_access_management.NewIdentityAccessManagementGetUserByNameParams()
	params.Name = raw

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.IdentityAccessManagement.IdentityAccessManagementGetUserByName(params, nil)
	if err != nil {
		if isStatusNotFound(err) {
			return nil, fmt.Errorf("no user found with username %q", raw)
		}
		return nil, fmt.Errorf("could not resolve username %q: %w", raw, err)
	}

	user := resp.GetPayload()
	if user == nil || user.ID == "" {
		return nil, fmt.Errorf("user %q resolved to an empty ID", raw)
	}

	d.SetId(user.ID)
	return []*schema.ResourceData{d}, nil
}

func UserDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: IdentityAccessManagement_GetUser,
		Schema:      zschema.DetailedUserSchema(),
	}
}

func IdentityAccessManagement_GetUser(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	var user *models.DetailedUser

	// NFR-165 §3.2: prefer the immutable system ID whenever state carries one.
	// Reading by username first meant that after a rename in configuration the
	// provider queried the *new* name, reported a lookup failure, and never
	// detected the drift on the object it was actually managing.
	id := d.Id()
	if id == "" {
		if idVal, isSet := d.GetOk("id"); isSet {
			id, _ = idVal.(string)
		}
	}

	switch {
	case id != "":
		var found bool
		user, found, diags = getUserById(ctx, d, m, id)
		if diags.HasError() {
			return diags
		}
		// NFR-165 §3.3: deleted out of band. Drop it from state so the next
		// plan proposes a re-create instead of failing the refresh.
		if !found {
			log.Printf("[INFO] user %s no longer exists, removing from state", id)
			d.SetId("")
			return diags
		}
	case hasUsername(d):
		user, diags = getUserByName(ctx, d, m)
		if diags.HasError() {
			return diags
		}
	default:
		// NFR-165 §3.1: without this the nil dereference below panicked the
		// provider instead of reporting a usable error.
		return append(diags, diag.Errorf("cannot read user: set either id or username")...)
	}

	if user == nil {
		return append(diags, diag.Errorf("user read returned an empty response")...)
	}

	zschema.SetDetailedUserResourceData(d, user)
	d.SetId(user.ID)

	return diags
}

func hasUsername(d *schema.ResourceData) bool {
	_, isSet := d.GetOk("username")
	return isSet
}

// getUserById fetches a user by its system ID. The boolean result is false
// when the API answered 404, which callers translate into "gone from state"
// rather than an error.
func getUserById(ctx context.Context, d *schema.ResourceData, m interface{}, id string) (*models.DetailedUser, bool, diag.Diagnostics) {
	var diags diag.Diagnostics

	params := identity_access_management.NewIdentityAccessManagementGetUserParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	params.ID = id

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.IdentityAccessManagement.IdentityAccessManagementGetUser(params, nil)
	if err != nil {
		if isStatusNotFound(err) {
			return nil, false, diags
		}

		log.Printf("[TRACE] user read error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return nil, false, diags
		}

		diags = append(diags, diag.Errorf("user read error: %s", err)...)
		return nil, false, diags
	}

	return resp.GetPayload(), true, diags
}

func getUserByName(ctx context.Context, d *schema.ResourceData, m interface{}) (*models.DetailedUser, diag.Diagnostics) {
	var diags diag.Diagnostics

	params := identity_access_management.NewIdentityAccessManagementGetUserByNameParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	nameVal, nameIsSet := d.GetOk("username")
	if nameIsSet {
		params.Name = nameVal.(string)
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: username")...)
		return nil, diags
	}

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.IdentityAccessManagement.IdentityAccessManagementGetUserByName(params, nil)
	if err != nil {
		log.Printf("[TRACE] user read error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return nil, diags
		}

		diags = append(diags, diag.Errorf("user read error: %s", err)...)
		return nil, diags
	}

	return resp.GetPayload(), diags
}

func IdentityAccessManagement_CreateUser(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	model := zschema.DetailedUserModel(d)
	params := identity_access_management.NewIdentityAccessManagementCreateUserParams()
	params.SetBody(model)

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.IdentityAccessManagement.IdentityAccessManagementCreateUser(params, nil)
	if err != nil {
		log.Printf("[TRACE] user create error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("user create error: %s", err)...)
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
	if diags := IdentityAccessManagement_GetUser(ctx, d, m); diags.HasError() {
		return diags
	}

	return diags
}

func IdentityAccessManagement_UpdateUser2(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	d.Partial(true)

	params := identity_access_management.NewIdentityAccessManagementUpdateUser2Params()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	params.SetBody(zschema.DetailedUserModel(d))
	// IdentityAccessManagementUpdateUser2Body

	// d.Id() is authoritative: it is set for a managed resource whether it was
	// created here or imported, whereas the "id" attribute is only populated
	// once a Read has run.
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
	resp, err := client.IdentityAccessManagement.IdentityAccessManagementUpdateUser2(params, nil)
	if err != nil {
		log.Printf("[TRACE] user update error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("user update error: %s", err)...)
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
	//
	// This used to read `if errs := ...; err != nil`, testing the update error
	// that is necessarily nil by this point, so the refresh result was
	// discarded and state was never repopulated after an update.
	if errs := IdentityAccessManagement_GetUser(ctx, d, m); errs.HasError() {
		return append(diags, errs...)
	}

	d.Partial(false)

	return diags
}

func IdentityAccessManagement_DeleteUser(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := identity_access_management.NewIdentityAccessManagementDeleteUserParams()

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

	_, err := client.IdentityAccessManagement.IdentityAccessManagementDeleteUser(params, nil)
	if err != nil {
		// Already gone: a delete of a missing object is a success.
		if isStatusNotFound(err) {
			d.SetId("")
			return diags
		}

		log.Printf("[TRACE] user delete error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("user delete error: %s", err)...)
		return diags
	}

	d.SetId("")
	return diags
}
