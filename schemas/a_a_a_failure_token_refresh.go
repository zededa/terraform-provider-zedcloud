package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AAAFailureTokenRefreshModel(d *schema.ResourceData) *models.AAAFailureTokenRefresh {
	var cause *models.AAAFailureTokenRefreshCause // AAAFailureTokenRefreshCause
	causeInterface, causeIsSet := d.GetOk("cause")
	if causeIsSet {
		causeModel := causeInterface.(string)
		cause = models.NewAAAFailureTokenRefreshCause(models.AAAFailureTokenRefreshCause(causeModel))
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
	return &models.AAAFailureTokenRefresh{
		Cause:    cause,
		Error:    error,
		Original: original,
	}
}

func AAAFailureTokenRefreshModelFromMap(m map[string]interface{}) *models.AAAFailureTokenRefresh {
	var cause *models.AAAFailureTokenRefreshCause // AAAFailureTokenRefreshCause
	causeInterface, causeIsSet := m["cause"]
	if causeIsSet {
		causeModel := causeInterface.(string)
		cause = models.NewAAAFailureTokenRefreshCause(models.AAAFailureTokenRefreshCause(causeModel))
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
	return &models.AAAFailureTokenRefresh{
		Cause:    cause,
		Error:    error,
		Original: original,
	}
}

func SetAAAFailureTokenRefreshResourceData(d *schema.ResourceData, m *models.AAAFailureTokenRefresh) {
	d.Set("cause", m.Cause)
	d.Set("error", m.Error)
	d.Set("original", SetOpaqueToken64SubResourceData([]*models.OpaqueToken64{m.Original}))
}

func SetAAAFailureTokenRefreshSubResourceData(m []*models.AAAFailureTokenRefresh) (d []*map[string]interface{}) {
	for _, AAAFailureTokenRefreshModel := range m {
		if AAAFailureTokenRefreshModel != nil {
			properties := make(map[string]interface{})
			properties["cause"] = AAAFailureTokenRefreshModel.Cause
			properties["error"] = AAAFailureTokenRefreshModel.Error
			properties["original"] = SetOpaqueToken64SubResourceData([]*models.OpaqueToken64{AAAFailureTokenRefreshModel.Original})
			d = append(d, &properties)
		}
	}
	return
}

func AAAFailureTokenRefreshSchema() map[string]*schema.Schema {
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

func GetAAAFailureTokenRefreshPropertyFields() (t []string) {
	return []string{
		"cause",
		"error",
		"original",
	}
}
