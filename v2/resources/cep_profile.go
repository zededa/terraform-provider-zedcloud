package resources

import (
	"context"
	"errors"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	cep_profile "github.com/zededa/terraform-provider-zedcloud/v2/client/certificate_enrollment_profile"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
	zschema "github.com/zededa/terraform-provider-zedcloud/v2/schemas"
)

func CEPProfileResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: CreateCEPProfile,
		ReadContext:   GetCEPProfile,
		UpdateContext: UpdateCEPProfile,
		DeleteContext: DeleteCEPProfile,
		Schema:        zschema.CEPProfileSchema(),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func CEPProfileDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: GetCEPProfile,
		Schema:      zschema.CEPProfileDataSourceSchema(),
	}
}

func CreateCEPProfile(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	model := zschema.CEPProfileModel(d)
	params := cep_profile.CreateParams()
	params.SetBody(model)

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.CertificateEnrollmentProfile.Create(params, nil)
	if err != nil {
		log.Printf("[TRACE] CEP profile create error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("CEP profile create error: %s", err)...)
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

	// note, we need to set the ID before fetching the newly created CEP profile.
	d.SetId(responseData.ObjectID)

	// the zedcloud API does not return the partially updated object but a custom response.
	// thus, we need to fetch the object and populate the state.
	if errs := GetCEPProfile(ctx, d, m); errs != nil {
		return append(diags, errs...)
	}

	return diags
}

func GetCEPProfile(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var cepProfile *models.CEPCommonSCEPProfile
	var diags diag.Diagnostics

	if _, isSet := d.GetOk("name"); isSet {
		cepProfile, diags = readCEPProfileByName(ctx, d, m)
	} else if _, isSet := d.GetOk("id"); isSet {
		cepProfile, diags = readCEPProfileByID(ctx, d, m)
	}

	if diags.HasError() {
		return diags
	}

	zschema.SetCEPProfileResourceData(d, cepProfile)
	d.SetId(cepProfile.ID)

	return diags
}

func readCEPProfileByID(ctx context.Context, d *schema.ResourceData, m interface{}) (*models.CEPCommonSCEPProfile, diag.Diagnostics) {
	var diags diag.Diagnostics

	params := cep_profile.GetByIDParams()

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

	resp, err := client.CertificateEnrollmentProfile.GetByID(params, nil)
	if err != nil {
		log.Printf("[TRACE] CEP profile read error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return nil, diags
		}

		diags = append(diags, diag.Errorf("CEP profile read error: %s", err)...)
		return nil, diags
	}

	return resp.GetPayload(), diags
}

func readCEPProfileByName(ctx context.Context, d *schema.ResourceData, m interface{}) (*models.CEPCommonSCEPProfile, diag.Diagnostics) {
	var diags diag.Diagnostics

	params := cep_profile.GetByNameParams()

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

	resp, err := client.CertificateEnrollmentProfile.GetByName(params, nil)
	if err != nil {
		log.Printf("[TRACE] CEP profile read error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return nil, diags
		}

		diags = append(diags, diag.Errorf("CEP profile read error: %s", err)...)
		return nil, diags
	}

	return resp.GetPayload(), diags
}

func UpdateCEPProfile(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := cep_profile.UpdateParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	params.SetBody(zschema.CEPProfileModel(d))

	idVal, idIsSet := d.GetOk("id")
	if idIsSet {
		id, _ := idVal.(string)
		params.ID = id
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: id")...)
		return diags
	}

	client := m.(*api_client.ZedcloudAPI)
	resp, err := client.CertificateEnrollmentProfile.Update(params, nil)
	if err != nil {
		log.Printf("[TRACE] CEP profile update error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("CEP profile update error: %s", err)...)
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
	if errs := GetCEPProfile(ctx, d, m); errs != nil {
		return append(diags, errs...)
	}

	return diags
}

func DeleteCEPProfile(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := cep_profile.DeleteParams()

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

	_, err := client.CertificateEnrollmentProfile.Delete(params, nil)
	if err != nil {
		log.Printf("[TRACE] CEP profile delete error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("CEP profile delete error: %s", err)...)
		return diags
	}

	d.SetId("")
	return diags
}
