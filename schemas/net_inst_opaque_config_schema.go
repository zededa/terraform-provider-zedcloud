package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate NetInstOpaqueConfig resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func NetInstOpaqueConfigModel(d *schema.ResourceData) *models.NetInstOpaqueConfig {
	var lisp *models.LispConfig // LispConfig
	lispInterface, lispIsSet := d.GetOk("lisp")
	if lispIsSet {
		lispMap := lispInterface.([]interface{})[0].(map[string]interface{})
		lisp = LispConfigModelFromMap(lispMap)
	}
	oconfig := d.Get("oconfig").(string)
	typeVar := d.Get("type").(*models.OpaqueConfigType) // OpaqueConfigType
	return &models.NetInstOpaqueConfig{
		Lisp:    lisp,
		Oconfig: oconfig,
		Type:    typeVar,
	}
}

func NetInstOpaqueConfigModelFromMap(m map[string]interface{}) *models.NetInstOpaqueConfig {
	var lisp *models.LispConfig // LispConfig
	lispInterface, lispIsSet := m["lisp"]
	if lispIsSet {
		lispMap := lispInterface.([]interface{})[0].(map[string]interface{})
		lisp = LispConfigModelFromMap(lispMap)
	}
	//
	oconfig := m["oconfig"].(string)
	typeVar := m["type"].(*models.OpaqueConfigType) // OpaqueConfigType
	return &models.NetInstOpaqueConfig{
		Lisp:    lisp,
		Oconfig: oconfig,
		Type:    typeVar,
	}
}

// Update the underlying NetInstOpaqueConfig resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetNetInstOpaqueConfigResourceData(d *schema.ResourceData, m *models.NetInstOpaqueConfig) {
	d.Set("lisp", SetLispConfigSubResourceData([]*models.LispConfig{m.Lisp}))
	d.Set("oconfig", m.Oconfig)
	d.Set("type", m.Type)
}

// Iterate throught and update the NetInstOpaqueConfig resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
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

// Schema mapping representing the NetInstOpaqueConfig resource defined in the Terraform configuration
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

// Retrieve property field names for updating the NetInstOpaqueConfig resource
func GetNetInstOpaqueConfigPropertyFields() (t []string) {
	return []string{
		"lisp",
		"oconfig",
		"type",
	}
}
