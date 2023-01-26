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
			lisp = LispModelFromMap(lispMap[0].(map[string]interface{}))
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
			lisp = LispModelFromMap(lispMap[0].(map[string]interface{}))
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

// Update the underlying NetInstOpaqueConfig resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetNetInstOpaqueConfigResourceData(d *schema.ResourceData, m *models.NetInstOpaqueConfig) {
	d.Set("lisp", SetLispSubResourceData([]*models.LispConfig{m.Lisp}))
	d.Set("oconfig", m.Oconfig)
	d.Set("type", m.Type)
}

// Iterate through and update the NetInstOpaqueConfig resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetNetInstOpaqueConfigSubResourceData(m []*models.NetInstOpaqueConfig) (d []*map[string]interface{}) {
	for _, NetInstOpaqueConfigModel := range m {
		if NetInstOpaqueConfigModel != nil {
			properties := make(map[string]interface{})
			properties["lisp"] = SetLispSubResourceData([]*models.LispConfig{NetInstOpaqueConfigModel.Lisp})
			properties["oconfig"] = NetInstOpaqueConfigModel.Oconfig
			properties["type"] = NetInstOpaqueConfigModel.Type
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the NetInstOpaqueConfig resource defined in the Terraform configuration
func NetInstOpaqueConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"lisp": {
			Description: `Deprecated - Lisp config`,
			Type:        schema.TypeList, //GoType: LispConfig
			Elem: &schema.Resource{
				Schema: Lisp(),
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

// Retrieve property field names for updating the NetInstOpaqueConfig resource
func GetNetInstOpaqueConfigPropertyFields() (t []string) {
	return []string{
		"lisp",
		"oconfig",
		"type",
	}
}
