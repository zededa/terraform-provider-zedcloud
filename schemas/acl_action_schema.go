package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate ACLAction resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ACLActionModel(d *schema.ResourceData) *models.ACLAction {
	drop, _ := d.Get("drop").(bool)
	limit, _ := d.Get("limit").(bool)
	var limitValue *models.LimitParams // LimitParams
	limitValueInterface, limitValueIsSet := d.GetOk("limit_value")
	if limitValueIsSet {
		limitValueMap := limitValueInterface.([]interface{})[0].(map[string]interface{})
		limitValue = LimitParamsModelFromMap(limitValueMap)
	}
	limitburstInt, _ := d.Get("limitburst").(int)
	limitburst := int64(limitburstInt)
	limitrateInt, _ := d.Get("limitrate").(int)
	limitrate := int64(limitrateInt)
	limitunit, _ := d.Get("limitunit").(string)
	portmap, _ := d.Get("portmap").(bool)
	var portmapto *models.MapParams // MapParams
	portmaptoInterface, portmaptoIsSet := d.GetOk("portmapto")
	if portmaptoIsSet {
		portmaptoMap := portmaptoInterface.([]interface{})[0].(map[string]interface{})
		portmapto = MapParamsModelFromMap(portmaptoMap)
	}
	return &models.ACLAction{
		Drop:       drop,
		Limit:      limit,
		LimitValue: limitValue,
		Limitburst: limitburst,
		Limitrate:  limitrate,
		Limitunit:  limitunit,
		Portmap:    portmap,
		Portmapto:  portmapto,
	}
}

func ACLActionModelFromMap(m map[string]interface{}) *models.ACLAction {
	drop := m["drop"].(bool)
	limit := m["limit"].(bool)
	var limitValue *models.LimitParams // LimitParams
	limitValueInterface, limitValueIsSet := m["limit_value"]
	if limitValueIsSet {
		limitValueMap := limitValueInterface.([]interface{})[0].(map[string]interface{})
		limitValue = LimitParamsModelFromMap(limitValueMap)
	}
	//
	limitburst := int64(m["limitburst"].(int)) // int64 false false false
	limitrate := int64(m["limitrate"].(int))   // int64 false false false
	limitunit := m["limitunit"].(string)
	portmap := m["portmap"].(bool)
	var portmapto *models.MapParams // MapParams
	portmaptoInterface, portmaptoIsSet := m["portmapto"]
	if portmaptoIsSet {
		portmaptoMap := portmaptoInterface.([]interface{})[0].(map[string]interface{})
		portmapto = MapParamsModelFromMap(portmaptoMap)
	}
	//
	return &models.ACLAction{
		Drop:       drop,
		Limit:      limit,
		LimitValue: limitValue,
		Limitburst: limitburst,
		Limitrate:  limitrate,
		Limitunit:  limitunit,
		Portmap:    portmap,
		Portmapto:  portmapto,
	}
}

// Update the underlying ACLAction resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetACLActionResourceData(d *schema.ResourceData, m *models.ACLAction) {
	d.Set("drop", m.Drop)
	d.Set("limit", m.Limit)
	d.Set("limit_value", SetLimitParamsSubResourceData([]*models.LimitParams{m.LimitValue}))
	d.Set("limitburst", m.Limitburst)
	d.Set("limitrate", m.Limitrate)
	d.Set("limitunit", m.Limitunit)
	d.Set("portmap", m.Portmap)
	d.Set("portmapto", SetMapParamsSubResourceData([]*models.MapParams{m.Portmapto}))
}

// Iterate through and update the ACLAction resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetACLActionSubResourceData(m []*models.ACLAction) (d []*map[string]interface{}) {
	for _, ACLActionModel := range m {
		if ACLActionModel != nil {
			properties := make(map[string]interface{})
			properties["drop"] = ACLActionModel.Drop
			properties["limit"] = ACLActionModel.Limit
			properties["limit_value"] = SetLimitParamsSubResourceData([]*models.LimitParams{ACLActionModel.LimitValue})
			properties["limitburst"] = ACLActionModel.Limitburst
			properties["limitrate"] = ACLActionModel.Limitrate
			properties["limitunit"] = ACLActionModel.Limitunit
			properties["portmap"] = ACLActionModel.Portmap
			properties["portmapto"] = SetMapParamsSubResourceData([]*models.MapParams{ACLActionModel.Portmapto})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the ACLAction resource defined in the Terraform configuration
func ACLActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"drop": {
			Description: `Drop the packet`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"limit": {
			Description: `UI map: AppDetailsPage:EnvironmentsPane, AppDetailsPage:EnvironmentsPane`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"limit_value": {
			Description: `Value to be used for limit action (Required if limit is true)`,
			Type:        schema.TypeList, //GoType: LimitParams
			Elem: &schema.Resource{
				Schema: LimitParamsSchema(),
			},
			Optional: true,
		},

		"limitburst": {
			Description: `UI map: AppDetailsPage:EnvironmentsPane, AppDetailsPage:EnvironmentsPane`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"limitrate": {
			Description: `UI map: AppDetailsPage:EnvironmentsPane, AppDetailsPage:EnvironmentsPane`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"limitunit": {
			Description: `UI map: AppDetailsPage:EnvironmentsPane, AppDetailsPage:EnvironmentsPane`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"portmap": {
			Description: `UI map: AppDetailsPage:EnvironmentsPane, AppDetailsPage:EnvironmentsPane`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"portmapto": {
			Description: `UI map: AppDetailsPage:EnvironmentsPane, AppDetailsPage:EnvironmentsPane`,
			Type:        schema.TypeList, //GoType: MapParams
			Elem: &schema.Resource{
				Schema: MapParamsSchema(),
			},
			Optional: true,
		},
	}
}

// Retrieve property field names for updating the ACLAction resource
func GetACLActionPropertyFields() (t []string) {
	return []string{
		"drop",
		"limit",
		"limit_value",
		"limitburst",
		"limitrate",
		"limitunit",
		"portmap",
		"portmapto",
	}
}
