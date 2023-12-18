package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AllowedEnterpriseListModel(d *schema.ResourceData) *models.AllowedEnterpriseList {
	var enterpriseEntitlements []*models.Entitlement // []*Entitlement
	enterpriseEntitlementsInterface, enterpriseEntitlementsIsSet := d.GetOk("enterprise_entitlements")
	if enterpriseEntitlementsIsSet {
		var items []interface{}
		if listItems, isList := enterpriseEntitlementsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = enterpriseEntitlementsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := EntitlementModelFromMap(v.(map[string]interface{}))
			enterpriseEntitlements = append(enterpriseEntitlements, m)
		}
	}
	var enterprises []*models.AllowedEnterprise // []*AllowedEnterprise
	enterprisesInterface, enterprisesIsSet := d.GetOk("enterprises")
	if enterprisesIsSet {
		var items []interface{}
		if listItems, isList := enterprisesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = enterprisesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := AllowedEnterpriseModelFromMap(v.(map[string]interface{}))
			enterprises = append(enterprises, m)
		}
	}
	error, _ := d.Get("error").(string)
	var next *models.Cursor // Cursor
	nextInterface, nextIsSet := d.GetOk("next")
	if nextIsSet && nextInterface != nil {
		nextMap := nextInterface.([]interface{})
		if len(nextMap) > 0 {
			next = CursorModelFromMap(nextMap[0].(map[string]interface{}))
		}
	}
	total, _ := d.Get("total").(string)
	return &models.AllowedEnterpriseList{
		EnterpriseEntitlements: enterpriseEntitlements,
		Enterprises:            enterprises,
		Error:                  error,
		Next:                   next,
		Total:                  total,
	}
}

func AllowedEnterpriseListModelFromMap(m map[string]interface{}) *models.AllowedEnterpriseList {
	var enterpriseEntitlements []*models.Entitlement // []*Entitlement
	enterpriseEntitlementsInterface, enterpriseEntitlementsIsSet := m["enterprise_entitlements"]
	if enterpriseEntitlementsIsSet {
		var items []interface{}
		if listItems, isList := enterpriseEntitlementsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = enterpriseEntitlementsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := EntitlementModelFromMap(v.(map[string]interface{}))
			enterpriseEntitlements = append(enterpriseEntitlements, m)
		}
	}
	var enterprises []*models.AllowedEnterprise // []*AllowedEnterprise
	enterprisesInterface, enterprisesIsSet := m["enterprises"]
	if enterprisesIsSet {
		var items []interface{}
		if listItems, isList := enterprisesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = enterprisesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := AllowedEnterpriseModelFromMap(v.(map[string]interface{}))
			enterprises = append(enterprises, m)
		}
	}
	error := m["error"].(string)
	var next *models.Cursor // Cursor
	nextInterface, nextIsSet := m["next"]
	if nextIsSet && nextInterface != nil {
		nextMap := nextInterface.([]interface{})
		if len(nextMap) > 0 {
			next = CursorModelFromMap(nextMap[0].(map[string]interface{}))
		}
	}
	//
	total := m["total"].(string)
	return &models.AllowedEnterpriseList{
		EnterpriseEntitlements: enterpriseEntitlements,
		Enterprises:            enterprises,
		Error:                  error,
		Next:                   next,
		Total:                  total,
	}
}

func SetAllowedEnterpriseListResourceData(d *schema.ResourceData, m *models.AllowedEnterpriseList) {
	d.Set("enterprise_entitlements", SetEntitlementSubResourceData(m.EnterpriseEntitlements))
	d.Set("enterprises", SetAllowedEnterpriseSubResourceData(m.Enterprises))
	d.Set("error", m.Error)
	d.Set("next", SetCursorSubResourceData([]*models.Cursor{m.Next}))
	d.Set("total", m.Total)
}

func SetAllowedEnterpriseListSubResourceData(m []*models.AllowedEnterpriseList) (d []*map[string]interface{}) {
	for _, AllowedEnterpriseListModel := range m {
		if AllowedEnterpriseListModel != nil {
			properties := make(map[string]interface{})
			properties["enterprise_entitlements"] = SetEntitlementSubResourceData(AllowedEnterpriseListModel.EnterpriseEntitlements)
			properties["enterprises"] = SetAllowedEnterpriseSubResourceData(AllowedEnterpriseListModel.Enterprises)
			properties["error"] = AllowedEnterpriseListModel.Error
			properties["next"] = SetCursorSubResourceData([]*models.Cursor{AllowedEnterpriseListModel.Next})
			properties["total"] = AllowedEnterpriseListModel.Total
			d = append(d, &properties)
		}
	}
	return
}

func AllowedEnterpriseListSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"enterprise_entitlements": {
			Description: `Enterprises entitlement data`,
			Type:        schema.TypeList, //GoType: []*Entitlement
			Elem: &schema.Resource{
				Schema: EntitlementSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"enterprises": {
			Description: `List of child enterprises`,
			Type:        schema.TypeList, //GoType: []*AllowedEnterprise
			Elem: &schema.Resource{
				Schema: AllowedEnterpriseSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"error": {
			Description: `Error while fetching list of child enterprises, if any`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"next": {
			Description: `Next page information`,
			Type:        schema.TypeList, //GoType: Cursor
			Elem: &schema.Resource{
				Schema: CursorSchema(),
			},
			Optional: true,
		},

		"total": {
			Description: `Total number of allowed enterprises`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetAllowedEnterpriseListPropertyFields() (t []string) {
	return []string{
		"enterprise_entitlements",
		"enterprises",
		"error",
		"next",
		"total",
	}
}
