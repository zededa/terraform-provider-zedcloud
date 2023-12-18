package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AAASuccessResponseLogoutModel(d *schema.ResourceData) *models.AAASuccessResponseLogout {
	var original *models.OpaqueToken64 // OpaqueToken64
	originalInterface, originalIsSet := d.GetOk("original")
	if originalIsSet && originalInterface != nil {
		originalMap := originalInterface.([]interface{})
		if len(originalMap) > 0 {
			original = OpaqueToken64ModelFromMap(originalMap[0].(map[string]interface{}))
		}
	}
	return &models.AAASuccessResponseLogout{
		Original: original,
	}
}

func AAASuccessResponseLogoutModelFromMap(m map[string]interface{}) *models.AAASuccessResponseLogout {
	var original *models.OpaqueToken64 // OpaqueToken64
	originalInterface, originalIsSet := m["original"]
	if originalIsSet && originalInterface != nil {
		originalMap := originalInterface.([]interface{})
		if len(originalMap) > 0 {
			original = OpaqueToken64ModelFromMap(originalMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.AAASuccessResponseLogout{
		Original: original,
	}
}

func SetAAASuccessResponseLogoutResourceData(d *schema.ResourceData, m *models.AAASuccessResponseLogout) {
	d.Set("original", SetOpaqueToken64SubResourceData([]*models.OpaqueToken64{m.Original}))
}

func SetAAASuccessResponseLogoutSubResourceData(m []*models.AAASuccessResponseLogout) (d []*map[string]interface{}) {
	for _, AAASuccessResponseLogoutModel := range m {
		if AAASuccessResponseLogoutModel != nil {
			properties := make(map[string]interface{})
			properties["original"] = SetOpaqueToken64SubResourceData([]*models.OpaqueToken64{AAASuccessResponseLogoutModel.Original})
			d = append(d, &properties)
		}
	}
	return
}

func AAASuccessResponseLogoutSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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

func GetAAASuccessResponseLogoutPropertyFields() (t []string) {
	return []string{
		"original",
	}
}
