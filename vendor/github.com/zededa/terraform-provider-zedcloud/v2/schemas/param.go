package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ParamModel(d *schema.ResourceData) *models.Param {
	name, _ := d.Get("name").(string)
	value, _ := d.Get("value").(string)
	return &models.Param{
		Name:  name,
		Value: value,
	}
}

func ParamModelFromMap(m map[string]interface{}) *models.Param {
	name := m["name"].(string)
	value := m["value"].(string)
	return &models.Param{
		Name:  name,
		Value: value,
	}
}

func SetParamResourceData(d *schema.ResourceData, m *models.Param) {
	d.Set("name", m.Name)
	d.Set("value", m.Value)
}

func SetParamSubResourceData(m []*models.Param) (d []*map[string]interface{}) {
	for _, ParamModel := range m {
		if ParamModel != nil {
			properties := make(map[string]interface{})
			properties["name"] = ParamModel.Name
			properties["value"] = ParamModel.Value
			d = append(d, &properties)
		}
	}
	return
}

func Param() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Description: `Name of the Parameter (Required)`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"value": {
			Description: `Value of the parameter (Required)`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetParamPropertyFields() (t []string) {
	return []string{
		"name",
		"value",
	}
}
