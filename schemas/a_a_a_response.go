package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AAAResponseModel(d *schema.ResourceData) *models.AAAResponse {
	var failure *models.AAAFailureResponse // AAAFailureResponse
	failureInterface, failureIsSet := d.GetOk("failure")
	if failureIsSet && failureInterface != nil {
		failureMap := failureInterface.([]interface{})
		if len(failureMap) > 0 {
			failure = AAAFailureResponseModelFromMap(failureMap[0].(map[string]interface{}))
		}
	}
	var mode *models.AAALoginModeResponse // AAALoginModeResponse
	modeInterface, modeIsSet := d.GetOk("mode")
	if modeIsSet && modeInterface != nil {
		modeMap := modeInterface.([]interface{})
		if len(modeMap) > 0 {
			mode = AAALoginModeResponseModelFromMap(modeMap[0].(map[string]interface{}))
		}
	}
	var notify *models.AAANotifyResponse // AAANotifyResponse
	notifyInterface, notifyIsSet := d.GetOk("notify")
	if notifyIsSet && notifyInterface != nil {
		notifyMap := notifyInterface.([]interface{})
		if len(notifyMap) > 0 {
			notify = AAANotifyResponseModelFromMap(notifyMap[0].(map[string]interface{}))
		}
	}
	var redirect *models.AAARedirectResponse // AAARedirectResponse
	redirectInterface, redirectIsSet := d.GetOk("redirect")
	if redirectIsSet && redirectInterface != nil {
		redirectMap := redirectInterface.([]interface{})
		if len(redirectMap) > 0 {
			redirect = AAARedirectResponseModelFromMap(redirectMap[0].(map[string]interface{}))
		}
	}
	var result *models.ZsrvResponse // ZsrvResponse
	resultInterface, resultIsSet := d.GetOk("result")
	if resultIsSet && resultInterface != nil {
		resultMap := resultInterface.([]interface{})
		if len(resultMap) > 0 {
			result = ZsrvResponseModelFromMap(resultMap[0].(map[string]interface{}))
		}
	}
	var success *models.AAASuccessResponse // AAASuccessResponse
	successInterface, successIsSet := d.GetOk("success")
	if successIsSet && successInterface != nil {
		successMap := successInterface.([]interface{})
		if len(successMap) > 0 {
			success = AAASuccessResponseModelFromMap(successMap[0].(map[string]interface{}))
		}
	}
	var totpURL *models.TOTPEnrolmentURL // TOTPEnrolmentURL
	totpURLInterface, totpURLIsSet := d.GetOk("totp_url")
	if totpURLIsSet && totpURLInterface != nil {
		totpURLMap := totpURLInterface.([]interface{})
		if len(totpURLMap) > 0 {
			totpURL = TOTPEnrolmentURLModelFromMap(totpURLMap[0].(map[string]interface{}))
		}
	}
	var typeVar *models.AAAResponseType // AAAResponseType
	typeInterface, typeIsSet := d.GetOk("type")
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewAAAResponseType(models.AAAResponseType(typeModel))
	}
	return &models.AAAResponse{
		Failure:  failure,
		Mode:     mode,
		Notify:   notify,
		Redirect: redirect,
		Result:   result,
		Success:  success,
		TotpURL:  totpURL,
		Type:     typeVar,
	}
}

func AAAResponseModelFromMap(m map[string]interface{}) *models.AAAResponse {
	var failure *models.AAAFailureResponse // AAAFailureResponse
	failureInterface, failureIsSet := m["failure"]
	if failureIsSet && failureInterface != nil {
		failureMap := failureInterface.([]interface{})
		if len(failureMap) > 0 {
			failure = AAAFailureResponseModelFromMap(failureMap[0].(map[string]interface{}))
		}
	}
	//
	var mode *models.AAALoginModeResponse // AAALoginModeResponse
	modeInterface, modeIsSet := m["mode"]
	if modeIsSet && modeInterface != nil {
		modeMap := modeInterface.([]interface{})
		if len(modeMap) > 0 {
			mode = AAALoginModeResponseModelFromMap(modeMap[0].(map[string]interface{}))
		}
	}
	//
	var notify *models.AAANotifyResponse // AAANotifyResponse
	notifyInterface, notifyIsSet := m["notify"]
	if notifyIsSet && notifyInterface != nil {
		notifyMap := notifyInterface.([]interface{})
		if len(notifyMap) > 0 {
			notify = AAANotifyResponseModelFromMap(notifyMap[0].(map[string]interface{}))
		}
	}
	//
	var redirect *models.AAARedirectResponse // AAARedirectResponse
	redirectInterface, redirectIsSet := m["redirect"]
	if redirectIsSet && redirectInterface != nil {
		redirectMap := redirectInterface.([]interface{})
		if len(redirectMap) > 0 {
			redirect = AAARedirectResponseModelFromMap(redirectMap[0].(map[string]interface{}))
		}
	}
	//
	var result *models.ZsrvResponse // ZsrvResponse
	resultInterface, resultIsSet := m["result"]
	if resultIsSet && resultInterface != nil {
		resultMap := resultInterface.([]interface{})
		if len(resultMap) > 0 {
			result = ZsrvResponseModelFromMap(resultMap[0].(map[string]interface{}))
		}
	}
	//
	var success *models.AAASuccessResponse // AAASuccessResponse
	successInterface, successIsSet := m["success"]
	if successIsSet && successInterface != nil {
		successMap := successInterface.([]interface{})
		if len(successMap) > 0 {
			success = AAASuccessResponseModelFromMap(successMap[0].(map[string]interface{}))
		}
	}
	//
	var totpURL *models.TOTPEnrolmentURL // TOTPEnrolmentURL
	totpURLInterface, totpURLIsSet := m["totp_url"]
	if totpURLIsSet && totpURLInterface != nil {
		totpURLMap := totpURLInterface.([]interface{})
		if len(totpURLMap) > 0 {
			totpURL = TOTPEnrolmentURLModelFromMap(totpURLMap[0].(map[string]interface{}))
		}
	}
	//
	var typeVar *models.AAAResponseType // AAAResponseType
	typeInterface, typeIsSet := m["type"]
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewAAAResponseType(models.AAAResponseType(typeModel))
	}
	return &models.AAAResponse{
		Failure:  failure,
		Mode:     mode,
		Notify:   notify,
		Redirect: redirect,
		Result:   result,
		Success:  success,
		TotpURL:  totpURL,
		Type:     typeVar,
	}
}

func SetAAAResponseResourceData(d *schema.ResourceData, m *models.AAAResponse) {
	d.Set("failure", SetAAAFailureResponseSubResourceData([]*models.AAAFailureResponse{m.Failure}))
	d.Set("mode", SetAAALoginModeResponseSubResourceData([]*models.AAALoginModeResponse{m.Mode}))
	d.Set("notify", SetAAANotifyResponseSubResourceData([]*models.AAANotifyResponse{m.Notify}))
	d.Set("redirect", SetAAARedirectResponseSubResourceData([]*models.AAARedirectResponse{m.Redirect}))
	d.Set("result", SetZsrvResponseSubResourceData([]*models.ZsrvResponse{m.Result}))
	d.Set("success", SetAAASuccessResponseSubResourceData([]*models.AAASuccessResponse{m.Success}))
	d.Set("totp_url", SetTOTPEnrolmentURLSubResourceData([]*models.TOTPEnrolmentURL{m.TotpURL}))
	d.Set("type", m.Type)
}

func SetAAAResponseSubResourceData(m []*models.AAAResponse) (d []*map[string]interface{}) {
	for _, AAAResponseModel := range m {
		if AAAResponseModel != nil {
			properties := make(map[string]interface{})
			properties["failure"] = SetAAAFailureResponseSubResourceData([]*models.AAAFailureResponse{AAAResponseModel.Failure})
			properties["mode"] = SetAAALoginModeResponseSubResourceData([]*models.AAALoginModeResponse{AAAResponseModel.Mode})
			properties["notify"] = SetAAANotifyResponseSubResourceData([]*models.AAANotifyResponse{AAAResponseModel.Notify})
			properties["redirect"] = SetAAARedirectResponseSubResourceData([]*models.AAARedirectResponse{AAAResponseModel.Redirect})
			properties["result"] = SetZsrvResponseSubResourceData([]*models.ZsrvResponse{AAAResponseModel.Result})
			properties["success"] = SetAAASuccessResponseSubResourceData([]*models.AAASuccessResponse{AAAResponseModel.Success})
			properties["totp_url"] = SetTOTPEnrolmentURLSubResourceData([]*models.TOTPEnrolmentURL{AAAResponseModel.TotpURL})
			properties["type"] = AAAResponseModel.Type
			d = append(d, &properties)
		}
	}
	return
}

func AAAResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"failure": {
			Description: ``,
			Type:        schema.TypeList, //GoType: AAAFailureResponse
			Elem: &schema.Resource{
				Schema: AAAFailureResponseSchema(),
			},
			Optional: true,
		},

		"mode": {
			Description: ``,
			Type:        schema.TypeList, //GoType: AAALoginModeResponse
			Elem: &schema.Resource{
				Schema: AAALoginModeResponseSchema(),
			},
			Optional: true,
		},

		"notify": {
			Description: ``,
			Type:        schema.TypeList, //GoType: AAANotifyResponse
			Elem: &schema.Resource{
				Schema: AAANotifyResponseSchema(),
			},
			Optional: true,
		},

		"redirect": {
			Description: ``,
			Type:        schema.TypeList, //GoType: AAARedirectResponse
			Elem: &schema.Resource{
				Schema: AAARedirectResponseSchema(),
			},
			Optional: true,
		},

		"result": {
			Description: ``,
			Type:        schema.TypeList, //GoType: ZsrvResponse
			Elem: &schema.Resource{
				Schema: ZsrvResponseSchema(),
			},
			Optional: true,
		},

		"success": {
			Description: ``,
			Type:        schema.TypeList, //GoType: AAASuccessResponse
			Elem: &schema.Resource{
				Schema: AAASuccessResponseSchema(),
			},
			Optional: true,
		},

		"totp_url": {
			Description: ``,
			Type:        schema.TypeList, //GoType: TOTPEnrolmentURL
			Elem: &schema.Resource{
				Schema: TOTPEnrolmentURLSchema(),
			},
			Optional: true,
		},

		"type": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetAAAResponsePropertyFields() (t []string) {
	return []string{
		"failure",
		"mode",
		"notify",
		"redirect",
		"result",
		"success",
		"totp_url",
		"type",
	}
}
