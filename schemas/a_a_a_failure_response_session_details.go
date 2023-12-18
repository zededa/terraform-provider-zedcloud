package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AAAFailureResponseSessionDetailsModel(d *schema.ResourceData) *models.AAAFailureResponseSessionDetails {
	var cause *models.AAAFailureResponseSessionDetailsCause // AAAFailureResponseSessionDetailsCause
	causeInterface, causeIsSet := d.GetOk("cause")
	if causeIsSet {
		causeModel := causeInterface.(string)
		cause = models.NewAAAFailureResponseSessionDetailsCause(models.AAAFailureResponseSessionDetailsCause(causeModel))
	}
	var original *models.OpaqueToken64 // OpaqueToken64
	originalInterface, originalIsSet := d.GetOk("original")
	if originalIsSet && originalInterface != nil {
		originalMap := originalInterface.([]interface{})
		if len(originalMap) > 0 {
			original = OpaqueToken64ModelFromMap(originalMap[0].(map[string]interface{}))
		}
	}
	return &models.AAAFailureResponseSessionDetails{
		Cause:    cause,
		Original: original,
	}
}

func AAAFailureResponseSessionDetailsModelFromMap(m map[string]interface{}) *models.AAAFailureResponseSessionDetails {
	var cause *models.AAAFailureResponseSessionDetailsCause // AAAFailureResponseSessionDetailsCause
	causeInterface, causeIsSet := m["cause"]
	if causeIsSet {
		causeModel := causeInterface.(string)
		cause = models.NewAAAFailureResponseSessionDetailsCause(models.AAAFailureResponseSessionDetailsCause(causeModel))
	}
	var original *models.OpaqueToken64 // OpaqueToken64
	originalInterface, originalIsSet := m["original"]
	if originalIsSet && originalInterface != nil {
		originalMap := originalInterface.([]interface{})
		if len(originalMap) > 0 {
			original = OpaqueToken64ModelFromMap(originalMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.AAAFailureResponseSessionDetails{
		Cause:    cause,
		Original: original,
	}
}

func SetAAAFailureResponseSessionDetailsResourceData(d *schema.ResourceData, m *models.AAAFailureResponseSessionDetails) {
	d.Set("cause", m.Cause)
	d.Set("original", SetOpaqueToken64SubResourceData([]*models.OpaqueToken64{m.Original}))
}

func SetAAAFailureResponseSessionDetailsSubResourceData(m []*models.AAAFailureResponseSessionDetails) (d []*map[string]interface{}) {
	for _, AAAFailureResponseSessionDetailsModel := range m {
		if AAAFailureResponseSessionDetailsModel != nil {
			properties := make(map[string]interface{})
			properties["cause"] = AAAFailureResponseSessionDetailsModel.Cause
			properties["original"] = SetOpaqueToken64SubResourceData([]*models.OpaqueToken64{AAAFailureResponseSessionDetailsModel.Original})
			d = append(d, &properties)
		}
	}
	return
}

func AAAFailureResponseSessionDetailsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cause": {
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

func GetAAAFailureResponseSessionDetailsPropertyFields() (t []string) {
	return []string{
		"cause",
		"original",
	}
}
