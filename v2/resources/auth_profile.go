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
IdentityAccessManagement identity access management API
*/

func AuthProfileResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: IdentityAccessManagement_CreateAuthProfile,
		DeleteContext: IdentityAccessManagement_DeleteAuthProfile,
		ReadContext:   IdentityAccessManagement_GetAuthProfile,
		UpdateContext: IdentityAccessManagement_UpdateAuthProfile,
		Schema:        zschema.AuthorizationProfileSchema(),
		Importer: &schema.ResourceImporter{
			StateContext: importAuthProfileStateContext,
		},
	}
}

func AuthProfileDataSource() *schema.Resource {
	return &schema.Resource{
		// UE-168: this data source had no ReadContext at all, so it silently
		// produced an empty object.
		ReadContext: IdentityAccessManagement_GetAuthProfile,
		Schema:      zschema.AuthorizationProfileSchema(),
	}
}

// importAuthProfileStateContext accepts either the 28-character profile ID or
// the profile name as the import ID (UE-166 §4.1, extended to auth profiles by
// UE-168):
//
//	terraform import zedcloud_auth_profile.okta corp-okta
//	terraform import zedcloud_auth_profile.okta AAGFABAEqnH4je5PHZTXSmHOs-XC
//
// Note on secrets: the controller blanks ClientSecret, CryptoKey and
// EncryptedSecrets on every read (fillAuthProfile in zedcloud
// srvs/indusv2/authprofileproc.go), so an imported profile carries no secret
// material. The operator has to supply client_secret in configuration after
// importing. See UE-168 and the acceptance test that pins this behaviour.
func importAuthProfileStateContext(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	raw := strings.TrimSpace(d.Id())
	if raw == "" {
		return nil, errors.New("import ID is empty: pass either the auth profile ID or its name")
	}

	if looksLikeObjectID(raw) {
		d.SetId(raw)
		return []*schema.ResourceData{d}, nil
	}

	params := identity_access_management.NewIdentityAccessManagementGetAuthProfileByNameParams()
	params.Name = raw

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.IdentityAccessManagement.IdentityAccessManagementGetAuthProfileByName(params, nil)
	if err != nil {
		if isStatusNotFound(err) {
			return nil, fmt.Errorf("no auth profile found with name %q", raw)
		}
		return nil, fmt.Errorf("could not resolve auth profile name %q: %w", raw, err)
	}

	profile := resp.GetPayload()
	if profile == nil || profile.ID == "" {
		return nil, fmt.Errorf("auth profile %q resolved to an empty ID", raw)
	}

	d.SetId(profile.ID)
	return []*schema.ResourceData{d}, nil
}

func IdentityAccessManagement_GetAuthProfile(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	var profile *models.AuthorizationProfile

	// UE-166 §3.2: the immutable system ID wins over the name whenever state
	// carries one, so a renamed profile is still found rather than reported
	// missing.
	id := d.Id()
	if id == "" {
		if idVal, isSet := d.GetOk("id"); isSet {
			id, _ = idVal.(string)
		}
	}

	switch {
	case id != "":
		var found bool
		profile, found, diags = GetAuthProfileById(ctx, d, m, id)
		if diags.HasError() {
			return diags
		}
		// UE-166 §3.3: deleted out of band — plan a re-create.
		if !found {
			log.Printf("[INFO] auth profile %s no longer exists, removing from state", id)
			d.SetId("")
			return diags
		}
	case hasName(d):
		profile, diags = GetAuthProfileByName(ctx, d, m)
		if diags.HasError() {
			return diags
		}
	default:
		// UE-166 §3.1: this dereferenced nil and panicked the provider.
		return append(diags, diag.Errorf("cannot read auth profile: set either id or name")...)
	}

	if profile == nil {
		return append(diags, diag.Errorf("auth profile read returned an empty response")...)
	}

	preserveOAuthWriteOnlyFields(d, profile)

	zschema.SetAuthorizationProfileResourceData(d, profile)
	d.SetId(profile.ID)

	return diags
}

// preserveOAuthWriteOnlyFields carries the write-only OAuth secrets from the
// prior state into the model just read from the API.
//
// UE-168: the controller blanks client_secret, crypto_key and
// encrypted_secrets on every read. Writing those blanks straight into state
// would erase what the operator configured, and the next plan would propose
// setting them again, for ever. Terraform's own post-apply "plan must be
// empty" check catches this, which is how it was found.
//
// On import there is no prior state, so the fields stay empty and the operator
// supplies client_secret in configuration afterwards -- that first plan showing
// client_secret as the only change is expected, and is documented in the
// resource's import example.
func preserveOAuthWriteOnlyFields(d *schema.ResourceData, profile *models.AuthorizationProfile) {
	if profile.OauthProfile == nil {
		return
	}

	if v, ok := d.GetOk("oauth_profile.0.client_secret"); ok && profile.OauthProfile.ClientSecret == "" {
		profile.OauthProfile.ClientSecret, _ = v.(string)
	}
	if v, ok := d.GetOk("oauth_profile.0.crypto_key"); ok && profile.OauthProfile.CryptoKey == "" {
		profile.OauthProfile.CryptoKey, _ = v.(string)
	}
	if v, ok := d.GetOk("oauth_profile.0.encrypted_secrets"); ok && len(profile.OauthProfile.EncryptedSecrets) == 0 {
		if m, isMap := v.(map[string]interface{}); isMap && len(m) > 0 {
			preserved := make(map[string]string, len(m))
			for k, raw := range m {
				s, isStr := raw.(string)
				if !isStr {
					continue
				}
				preserved[k] = s
			}
			profile.OauthProfile.EncryptedSecrets = preserved
		}
	}
}

// GetAuthProfileById fetches an auth profile by its system ID. The boolean
// result is false when the API answered 404.
func GetAuthProfileById(ctx context.Context, d *schema.ResourceData, m interface{}, id string) (*models.AuthorizationProfile, bool, diag.Diagnostics) {
	var diags diag.Diagnostics

	params := identity_access_management.NewIdentityAccessManagementGetAuthProfileParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	params.ID = id

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.IdentityAccessManagement.IdentityAccessManagementGetAuthProfile(params, nil)
	if err != nil {
		if isStatusNotFound(err) {
			return nil, false, diags
		}

		log.Printf("[TRACE] auth profile read error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return nil, false, diags
		}

		diags = append(diags, diag.Errorf("auth profile read error: %s", err)...)
		return nil, false, diags
	}

	return resp.GetPayload(), true, diags
}

func GetAuthProfileByName(ctx context.Context, d *schema.ResourceData, m interface{}) (*models.AuthorizationProfile, diag.Diagnostics) {
	var diags diag.Diagnostics

	params := identity_access_management.NewIdentityAccessManagementGetAuthProfileByNameParams()

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

	resp, err := client.IdentityAccessManagement.IdentityAccessManagementGetAuthProfileByName(params, nil)
	if err != nil {
		log.Printf("[TRACE] auth profile read error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return nil, diags
		}

		diags = append(diags, diag.Errorf("auth profile read error: %s", err)...)
		return nil, diags
	}

	return resp.GetPayload(), diags
}

func IdentityAccessManagement_CreateAuthProfile(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	model := zschema.AuthorizationProfileModel(d)
	params := identity_access_management.NewIdentityAccessManagementCreateAuthProfileParams()
	params.SetBody(model)

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.IdentityAccessManagement.IdentityAccessManagementCreateAuthProfile(params, nil)
	if err != nil {
		log.Printf("[TRACE] auth profile create error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("auth profile create error: %s", err)...)
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
	if diags := IdentityAccessManagement_GetAuthProfile(ctx, d, m); diags.HasError() {
		return diags
	}

	return diags
}

func IdentityAccessManagement_UpdateAuthProfile(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	d.Partial(true)

	params := identity_access_management.NewIdentityAccessManagementUpdateAuthProfileParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	params.SetBody(zschema.AuthorizationProfileModel(d))
	// IdentityAccessManagementUpdateAuthProfileBody

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
	resp, err := client.IdentityAccessManagement.IdentityAccessManagementUpdateAuthProfile(params, nil)
	if err != nil {
		log.Printf("[TRACE] auth profile update error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("auth profile update error: %s", err)...)
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
	// UE-168: this called IdentityAccessManagement_Get*User* — a copy-paste
	// slip that refreshed the wrong object entirely. It went unnoticed because
	// the guard tested `err`, the update error that is necessarily nil here,
	// rather than `errs`. Once the user read learned to treat a 404 as "gone"
	// (UE-166 §3.3), looking up an auth profile ID in the user API would have
	// cleared this resource's ID from state on every update.
	if errs := IdentityAccessManagement_GetAuthProfile(ctx, d, m); errs.HasError() {
		return append(diags, errs...)
	}

	d.Partial(false)

	return diags
}

func IdentityAccessManagement_DeleteAuthProfile(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := identity_access_management.NewIdentityAccessManagementDeleteAuthProfileParams()

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

	_, err := client.IdentityAccessManagement.IdentityAccessManagementDeleteAuthProfile(params, nil)
	if err != nil {
		// Already gone: a delete of a missing object is a success.
		if isStatusNotFound(err) {
			d.SetId("")
			return diags
		}

		log.Printf("[TRACE] auth profile delete error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("auth profile delete error: %s", err)...)
		return diags
	}

	d.SetId("")
	return diags
}
