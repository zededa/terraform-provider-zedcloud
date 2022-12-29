package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate AppACEAction resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func AppACEActionModel(d *schema.ResourceData) *models.AppACEAction {
	drop, _ := d.Get("drop").(bool)
	limit, _ := d.Get("limit").(bool)
	limitburstInt, _ := d.Get("limitburst").(int)
	limitburst := int64(limitburstInt)
	limitrateInt, _ := d.Get("limitrate").(int)
	limitrate := int64(limitrateInt)
	limitunit, _ := d.Get("limitunit").(string)
	var mapparams *models.AppMapParams // AppMapParams
	mapparamsInterface, mapparamsIsSet := d.GetOk("mapparams")
	if mapparamsIsSet {
		mapparamsMap := mapparamsInterface.([]interface{})[0].(map[string]interface{})
		mapparams = AppMapParamsModelFromMap(mapparamsMap)
	}
	portmap, _ := d.Get("portmap").(bool)
	return &models.AppACEAction{
		Drop:       &drop,       // bool true false false
		Limit:      &limit,      // bool true false false
		Limitburst: &limitburst, // int64 true false false
		Limitrate:  &limitrate,  // int64 true false false
		Limitunit:  &limitunit,  // string true false false
		Mapparams:  mapparams,
		Portmap:    &portmap, // bool true false false
	}
}

func AppACEActionModelFromMap(m map[string]interface{}) *models.AppACEAction {
	drop := m["drop"].(bool)
	limit := m["limit"].(bool)
	limitburst := int64(m["limitburst"].(int)) // int64 true false false
	limitrate := int64(m["limitrate"].(int))   // int64 true false false
	limitunit := m["limitunit"].(string)
	var mapparams *models.AppMapParams // AppMapParams
	mapparamsInterface, mapparamsIsSet := m["mapparams"]
	if mapparamsIsSet {
		mapparamsMap := mapparamsInterface.([]interface{})[0].(map[string]interface{})
		mapparams = AppMapParamsModelFromMap(mapparamsMap)
	}
	//
	portmap := m["portmap"].(bool)
	return &models.AppACEAction{
		Drop:       &drop,
		Limit:      &limit,
		Limitburst: &limitburst,
		Limitrate:  &limitrate,
		Limitunit:  &limitunit,
		Mapparams:  mapparams,
		Portmap:    &portmap,
	}
}

// Update the underlying AppACEAction resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetAppACEActionResourceData(d *schema.ResourceData, m *models.AppACEAction) {
	d.Set("drop", m.Drop)
	d.Set("limit", m.Limit)
	d.Set("limitburst", m.Limitburst)
	d.Set("limitrate", m.Limitrate)
	d.Set("limitunit", m.Limitunit)
	d.Set("mapparams", SetAppMapParamsSubResourceData([]*models.AppMapParams{m.Mapparams}))
	d.Set("portmap", m.Portmap)
}

// Iterate through and update the AppACEAction resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetAppACEActionSubResourceData(m []*models.AppACEAction) (d []*map[string]interface{}) {
	for _, AppACEActionModel := range m {
		if AppACEActionModel != nil {
			properties := make(map[string]interface{})
			properties["drop"] = AppACEActionModel.Drop
			properties["limit"] = AppACEActionModel.Limit
			properties["limitburst"] = AppACEActionModel.Limitburst
			properties["limitrate"] = AppACEActionModel.Limitrate
			properties["limitunit"] = AppACEActionModel.Limitunit
			properties["mapparams"] = SetAppMapParamsSubResourceData([]*models.AppMapParams{AppACEActionModel.Mapparams})
			properties["portmap"] = AppACEActionModel.Portmap
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the AppACEAction resource defined in the Terraform configuration
func AppACEActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"drop": {
			Description: `ACE drop flag`,
			Type:        schema.TypeBool,
			Required:    true,
		},

		"limit": {
			Description: `ACE limit flag`,
			Type:        schema.TypeBool,
			Required:    true,
		},

		"limitburst": {
			Description: `ACE limit burst`,
			Type:        schema.TypeInt,
			Required:    true,
		},

		"limitrate": {
			Description: `ACE limit rate`,
			Type:        schema.TypeInt,
			Required:    true,
		},

		"limitunit": {
			Description: `ACE limit unit`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"mapparams": {
			Description: `Application map params`,
			Type:        schema.TypeList, //GoType: AppMapParams
			Elem: &schema.Resource{
				Schema: AppMapParamsSchema(),
			},
			Required: true,
		},

		"portmap": {
			Description: `application port map flag`,
			Type:        schema.TypeBool,
			Required:    true,
		},
	}
}

// Retrieve property field names for updating the AppACEAction resource
func GetAppACEActionPropertyFields() (t []string) {
	return []string{
		"drop",
		"limit",
		"limitburst",
		"limitrate",
		"limitunit",
		"mapparams",
		"portmap",
	}
}
