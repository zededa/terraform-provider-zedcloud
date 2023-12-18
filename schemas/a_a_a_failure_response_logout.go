package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AAAFailureResponseLogoutModel(d *schema.ResourceData) *models.AAAFailureResponseLogout {
	var cause *models.AAAFailureResponseLogoutCause // AAAFailureResponseLogoutCause
	causeInterface, causeIsSet := d.GetOk("cause")
	if causeIsSet {
		causeModel := causeInterface.(string)
		cause = models.NewAAAFailureResponseLogoutCause(models.AAAFailureResponseLogoutCause(causeModel))
	}
	error, _ := d.Get("error").(string)
	var original *models.OpaqueToken64 // OpaqueToken64
	originalInterface, originalIsSet := d.GetOk("original")
	if originalIsSet && originalInterface != nil {
		originalMap := originalInterface.([]interface{})
		if len(originalMap) > 0 {
			original = OpaqueToken64ModelFromMap(originalMap[0].(map[string]interface{}))
		}
	}
	return &models.AAAFailureResponseLogout{
		Cause:    cause,
		Error:    error,
		Original: original,
	}
}

func AAAFailureResponseLogoutModelFromMap(m map[string]interface{}) *models.AAAFailureResponseLogout {
	var cause *models.AAAFailureResponseLogoutCause // AAAFailureResponseLogoutCause
	causeInterface, causeIsSet := m["cause"]
	if causeIsSet {
		causeModel := causeInterface.(string)
		cause = models.NewAAAFailureResponseLogoutCause(models.AAAFailureResponseLogoutCause(causeModel))
	}
	error := m["error"].(string)
	var original *models.OpaqueToken64 // OpaqueToken64
	originalInterface, originalIsSet := m["original"]
	if originalIsSet && originalInterface != nil {
		originalMap := originalInterface.([]interface{})
		if len(originalMap) > 0 {
			original = OpaqueToken64ModelFromMap(originalMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.AAAFailureResponseLogout{
		Cause:    cause,
		Error:    error,
		Original: original,
	}
}

func SetAAAFailureResponseLogoutResourceData(d *schema.ResourceData, m *models.AAAFailureResponseLogout) {
	d.Set("cause", m.Cause)
	d.Set("error", m.Error)
	d.Set("original", SetOpaqueToken64SubResourceData([]*models.OpaqueToken64{m.Original}))
}

func SetAAAFailureResponseLogoutSubResourceData(m []*models.AAAFailureResponseLogout) (d []*map[string]interface{}) {
	for _, AAAFailureResponseLogoutModel := range m {
		if AAAFailureResponseLogoutModel != nil {
			properties := make(map[string]interface{})
			properties["cause"] = AAAFailureResponseLogoutModel.Cause
			properties["error"] = AAAFailureResponseLogoutModel.Error
			properties["original"] = SetOpaqueToken64SubResourceData([]*models.OpaqueToken64{AAAFailureResponseLogoutModel.Original})
			d = append(d, &properties)
		}
	}
	return
}

func AAAFailureResponseLogoutSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cause": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"error": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"original": {
			Description: ``,
			Type:        schema.TypeList, //GoType: OpaqueToken64
			Elem: &schema.Resource{
				Schema: OpaqueToken64Schema(),
			},
			Optional: true,
		},
	}
}

func GetAAAFailureResponseLogoutPropertyFields() (t []string) {
	return []string{
		"cause",
		"error",
		"original",
	}
}
