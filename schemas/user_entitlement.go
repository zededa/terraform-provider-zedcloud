package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func UserEntitlementModel(d *schema.ResourceData) *models.UserEntitlement {
	limitInt, _ := d.Get("limit").(int)
	limit := int32(limitInt)
	usageInt, _ := d.Get("usage").(int)
	usage := int64(usageInt)
	return &models.UserEntitlement{
		Limit: limit,
		Usage: usage,
	}
}

func UserEntitlementModelFromMap(m map[string]interface{}) *models.UserEntitlement {
	limit := int32(m["limit"].(int)) // int32
	usage := int64(m["usage"].(int)) // int64
	return &models.UserEntitlement{
		Limit: limit,
		Usage: usage,
	}
}

func SetUserEntitlementResourceData(d *schema.ResourceData, m *models.UserEntitlement) {
	d.Set("limit", m.Limit)
	d.Set("usage", m.Usage)
}

func SetUserEntitlementSubResourceData(m []*models.UserEntitlement) (d []*map[string]interface{}) {
	for _, UserEntitlementModel := range m {
		if UserEntitlementModel != nil {
			properties := make(map[string]interface{})
			properties["limit"] = UserEntitlementModel.Limit
			properties["usage"] = UserEntitlementModel.Usage
			d = append(d, &properties)
		}
	}
	return
}

func UserEntitlementSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"limit": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"usage": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},
	}
}

func GetUserEntitlementPropertyFields() (t []string) {
	return []string{
		"limit",
		"usage",
	}
}
