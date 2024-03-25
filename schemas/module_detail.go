package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ModuleDetailModel(d *schema.ResourceData) *models.ModuleDetail {
	environment := map[string]string{}
	environmentInterface, environmentIsSet := d.GetOk("environment")
	if environmentIsSet {
		environmentMap := environmentInterface.(map[string]interface{})
		for k, v := range environmentMap {
			if v == nil {
				continue
			}
			environment[k] = v.(string)
		}
	}

	var moduleType *models.ModuleType // ModuleType
	moduleTypeInterface, moduleTypeIsSet := d.GetOk("module_type")
	if moduleTypeIsSet {
		moduleTypeModel := moduleTypeInterface.(string)
		moduleType = models.NewModuleType(models.ModuleType(moduleTypeModel))
	}
	routes := map[string]string{}
	routesInterface, routesIsSet := d.GetOk("routes")
	if routesIsSet {
		routesMap := routesInterface.(map[string]interface{})
		for k, v := range routesMap {
			if v == nil {
				continue
			}
			routes[k] = v.(string)
		}
	}

	twinDetail, _ := d.Get("twin_detail").(string)
	return &models.ModuleDetail{
		Environment: environment,
		ModuleType:  moduleType,
		Routes:      routes,
		TwinDetail:  twinDetail,
	}
}

func ModuleDetailModelFromMap(m map[string]interface{}) *models.ModuleDetail {
	environment := map[string]string{}
	environmentInterface, environmentIsSet := m["environment"]
	if environmentIsSet {
		environmentMap := environmentInterface.(map[string]interface{})
		for k, v := range environmentMap {
			if v == nil {
				continue
			}
			environment[k] = v.(string)
		}
	}

	var moduleType *models.ModuleType // ModuleType
	moduleTypeInterface, moduleTypeIsSet := m["module_type"]
	if moduleTypeIsSet {
		moduleTypeModel := moduleTypeInterface.(string)
		moduleType = models.NewModuleType(models.ModuleType(moduleTypeModel))
	}
	routes := map[string]string{}
	routesInterface, routesIsSet := m["routes"]
	if routesIsSet {
		routesMap := routesInterface.(map[string]interface{})
		for k, v := range routesMap {
			if v == nil {
				continue
			}
			routes[k] = v.(string)
		}
	}

	twinDetail := m["twin_detail"].(string)
	return &models.ModuleDetail{
		Environment: environment,
		ModuleType:  moduleType,
		Routes:      routes,
		TwinDetail:  twinDetail,
	}
}

func SetModuleDetailResourceData(d *schema.ResourceData, m *models.ModuleDetail) {
	d.Set("environment", m.Environment)
	d.Set("module_type", m.ModuleType)
	d.Set("routes", m.Routes)
	d.Set("twin_detail", m.TwinDetail)
}

func SetModuleDetailSubResourceData(m []*models.ModuleDetail) (d []*map[string]interface{}) {
	for _, ModuleDetailModel := range m {
		if ModuleDetailModel != nil {
			properties := make(map[string]interface{})
			properties["environment"] = ModuleDetailModel.Environment
			properties["module_type"] = ModuleDetailModel.ModuleType
			properties["routes"] = ModuleDetailModel.Routes
			properties["twin_detail"] = ModuleDetailModel.TwinDetail
			d = append(d, &properties)
		}
	}
	return
}

func ModuleDetail() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"environment": {
			Description: `Extra information to module to make configuration easier`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"module_type": {
			Description: `Type of modules`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"routes": {
			Description: `Send messages between modules or send messages from modules to iot hub`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"twin_detail": {
			Description: `Base64 encoded module twin details, desired properties of the module will be updated to reflect these values`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetModuleDetailPropertyFields() (t []string) {
	return []string{
		"environment",
		"module_type",
		"routes",
		"twin_detail",
	}
}
