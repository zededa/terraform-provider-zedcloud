package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func NetInstOpaqueConfigModel(d *schema.ResourceData) *models.NetInstOpaqueConfig {
	var lisp *models.LispConfig // LispConfig
	lispInterface, lispIsSet := d.GetOk("lisp")
	if lispIsSet && lispInterface != nil {
		lispMap := lispInterface.([]interface{})
		if len(lispMap) > 0 {
			lisp = LispConfigModelFromMap(lispMap[0].(map[string]interface{}))
		}
	}
	oconfig, _ := d.Get("oconfig").(string)
	var typeVar *models.OpaqueConfigType // OpaqueConfigType
	typeInterface, typeIsSet := d.GetOk("type")
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewOpaqueConfigType(models.OpaqueConfigType(typeModel))
	}
	return &models.NetInstOpaqueConfig{
		Lisp:    lisp,
		Oconfig: oconfig,
		Type:    typeVar,
	}
}

func NetInstOpaqueConfigModelFromMap(m map[string]interface{}) *models.NetInstOpaqueConfig {
	var lisp *models.LispConfig // LispConfig
	lispInterface, lispIsSet := m["lisp"]
	if lispIsSet && lispInterface != nil {
		lispMap := lispInterface.([]interface{})
		if len(lispMap) > 0 {
			lisp = LispConfigModelFromMap(lispMap[0].(map[string]interface{}))
		}
	}
	//
	oconfig := m["oconfig"].(string)
	var typeVar *models.OpaqueConfigType // OpaqueConfigType
	typeInterface, typeIsSet := m["type"]
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewOpaqueConfigType(models.OpaqueConfigType(typeModel))
	}
	return &models.NetInstOpaqueConfig{
		Lisp:    lisp,
		Oconfig: oconfig,
		Type:    typeVar,
	}
}

func SetNetInstOpaqueConfigResourceData(d *schema.ResourceData, m *models.NetInstOpaqueConfig) {
	d.Set("lisp", SetLispConfigSubResourceData([]*models.LispConfig{m.Lisp}))
	d.Set("oconfig", m.Oconfig)
	d.Set("type", m.Type)
}

func SetNetInstOpaqueConfigSubResourceData(m []*models.NetInstOpaqueConfig) (d []*map[string]interface{}) {
	for _, NetInstOpaqueConfigModel := range m {
		if NetInstOpaqueConfigModel != nil {
			properties := make(map[string]interface{})
			properties["lisp"] = SetLispConfigSubResourceData([]*models.LispConfig{NetInstOpaqueConfigModel.Lisp})
			properties["oconfig"] = NetInstOpaqueConfigModel.Oconfig
			properties["type"] = NetInstOpaqueConfigModel.Type
			d = append(d, &properties)
		}
	}
	return
}

func NetInstOpaqueConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"lisp": {
			Description: `Deprecated - Lisp config`,
			Type:        schema.TypeList, //GoType: LispConfig
			Elem: &schema.Resource{
				Schema: LispConfigSchema(),
			},
			Optional: true,
		},

		"oconfig": {
			Description: `base64 encoded string of opaque config`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"type": {
			Description: `type of Opaque config`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetNetInstOpaqueConfigPropertyFields() (t []string) {
	return []string{
		"lisp",
		"oconfig",
		"type",
	}
}
