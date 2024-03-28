package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func AppsModel(d *schema.ResourceData) *models.Apps {
	var list []*models.AppSummary // []*AppSummary
	listInterface, listIsSet := d.GetOk("list")
	if listIsSet {
		var items []interface{}
		if listItems, isList := listInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = listInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := AppSummaryModelFromMap(v.(map[string]interface{}))
			list = append(list, m)
		}
	}
	var next *models.Cursor // Cursor
	nextInterface, nextIsSet := d.GetOk("next")
	if nextIsSet && nextInterface != nil {
		nextMap := nextInterface.([]interface{})
		if len(nextMap) > 0 {
			next = CursorModelFromMap(nextMap[0].(map[string]interface{}))
		}
	}
	var summaryByAppInstanceDistribution *models.Summary // Summary
	summaryByAppInstanceDistributionInterface, summaryByAppInstanceDistributionIsSet := d.GetOk("summary_by_app_instance_distribution")
	if summaryByAppInstanceDistributionIsSet && summaryByAppInstanceDistributionInterface != nil {
		summaryByAppInstanceDistributionMap := summaryByAppInstanceDistributionInterface.([]interface{})
		if len(summaryByAppInstanceDistributionMap) > 0 {
			summaryByAppInstanceDistribution = SummaryModelFromMap(summaryByAppInstanceDistributionMap[0].(map[string]interface{}))
		}
	}
	var summaryByAppType *models.Summary // Summary
	summaryByAppTypeInterface, summaryByAppTypeIsSet := d.GetOk("summary_by_app_type")
	if summaryByAppTypeIsSet && summaryByAppTypeInterface != nil {
		summaryByAppTypeMap := summaryByAppTypeInterface.([]interface{})
		if len(summaryByAppTypeMap) > 0 {
			summaryByAppType = SummaryModelFromMap(summaryByAppTypeMap[0].(map[string]interface{}))
		}
	}
	var summaryByCategory *models.Summary // Summary
	summaryByCategoryInterface, summaryByCategoryIsSet := d.GetOk("summary_by_category")
	if summaryByCategoryIsSet && summaryByCategoryInterface != nil {
		summaryByCategoryMap := summaryByCategoryInterface.([]interface{})
		if len(summaryByCategoryMap) > 0 {
			summaryByCategory = SummaryModelFromMap(summaryByCategoryMap[0].(map[string]interface{}))
		}
	}
	var summaryByOrigin *models.Summary // Summary
	summaryByOriginInterface, summaryByOriginIsSet := d.GetOk("summary_by_origin")
	if summaryByOriginIsSet && summaryByOriginInterface != nil {
		summaryByOriginMap := summaryByOriginInterface.([]interface{})
		if len(summaryByOriginMap) > 0 {
			summaryByOrigin = SummaryModelFromMap(summaryByOriginMap[0].(map[string]interface{}))
		}
	}
	return &models.Apps{
		List:                             list,
		Next:                             next,
		SummaryByAppInstanceDistribution: summaryByAppInstanceDistribution,
		SummaryByAppType:                 summaryByAppType,
		SummaryByCategory:                summaryByCategory,
		SummaryByOrigin:                  summaryByOrigin,
	}
}

func AppsModelFromMap(m map[string]interface{}) *models.Apps {
	var list []*models.AppSummary // []*AppSummary
	listInterface, listIsSet := m["list"]
	if listIsSet {
		var items []interface{}
		if listItems, isList := listInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = listInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := AppSummaryModelFromMap(v.(map[string]interface{}))
			list = append(list, m)
		}
	}
	var next *models.Cursor // Cursor
	nextInterface, nextIsSet := m["next"]
	if nextIsSet && nextInterface != nil {
		nextMap := nextInterface.([]interface{})
		if len(nextMap) > 0 {
			next = CursorModelFromMap(nextMap[0].(map[string]interface{}))
		}
	}
	//
	var summaryByAppInstanceDistribution *models.Summary // Summary
	summaryByAppInstanceDistributionInterface, summaryByAppInstanceDistributionIsSet := m["summary_by_app_instance_distribution"]
	if summaryByAppInstanceDistributionIsSet && summaryByAppInstanceDistributionInterface != nil {
		summaryByAppInstanceDistributionMap := summaryByAppInstanceDistributionInterface.([]interface{})
		if len(summaryByAppInstanceDistributionMap) > 0 {
			summaryByAppInstanceDistribution = SummaryModelFromMap(summaryByAppInstanceDistributionMap[0].(map[string]interface{}))
		}
	}
	//
	var summaryByAppType *models.Summary // Summary
	summaryByAppTypeInterface, summaryByAppTypeIsSet := m["summary_by_app_type"]
	if summaryByAppTypeIsSet && summaryByAppTypeInterface != nil {
		summaryByAppTypeMap := summaryByAppTypeInterface.([]interface{})
		if len(summaryByAppTypeMap) > 0 {
			summaryByAppType = SummaryModelFromMap(summaryByAppTypeMap[0].(map[string]interface{}))
		}
	}
	//
	var summaryByCategory *models.Summary // Summary
	summaryByCategoryInterface, summaryByCategoryIsSet := m["summary_by_category"]
	if summaryByCategoryIsSet && summaryByCategoryInterface != nil {
		summaryByCategoryMap := summaryByCategoryInterface.([]interface{})
		if len(summaryByCategoryMap) > 0 {
			summaryByCategory = SummaryModelFromMap(summaryByCategoryMap[0].(map[string]interface{}))
		}
	}
	//
	var summaryByOrigin *models.Summary // Summary
	summaryByOriginInterface, summaryByOriginIsSet := m["summary_by_origin"]
	if summaryByOriginIsSet && summaryByOriginInterface != nil {
		summaryByOriginMap := summaryByOriginInterface.([]interface{})
		if len(summaryByOriginMap) > 0 {
			summaryByOrigin = SummaryModelFromMap(summaryByOriginMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.Apps{
		List:                             list,
		Next:                             next,
		SummaryByAppInstanceDistribution: summaryByAppInstanceDistribution,
		SummaryByAppType:                 summaryByAppType,
		SummaryByCategory:                summaryByCategory,
		SummaryByOrigin:                  summaryByOrigin,
	}
}

// Update the underlying Apps resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetAppsResourceData(d *schema.ResourceData, m *models.Apps) {
	d.Set("list", SetAppSummarySubResourceData(m.List))
	d.Set("next", SetCursorSubResourceData([]*models.Cursor{m.Next}))
	d.Set("summary_by_app_instance_distribution", SetSummarySubResourceData([]*models.Summary{m.SummaryByAppInstanceDistribution}))
	d.Set("summary_by_app_type", SetSummarySubResourceData([]*models.Summary{m.SummaryByAppType}))
	d.Set("summary_by_category", SetSummarySubResourceData([]*models.Summary{m.SummaryByCategory}))
	d.Set("summary_by_origin", SetSummarySubResourceData([]*models.Summary{m.SummaryByOrigin}))
}

// Iterate through and update the Apps resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetAppsSubResourceData(m []*models.Apps) (d []*map[string]interface{}) {
	for _, AppsModel := range m {
		if AppsModel != nil {
			properties := make(map[string]interface{})
			properties["list"] = SetAppSummarySubResourceData(AppsModel.List)
			properties["next"] = SetCursorSubResourceData([]*models.Cursor{AppsModel.Next})
			properties["summary_by_app_instance_distribution"] = SetSummarySubResourceData([]*models.Summary{AppsModel.SummaryByAppInstanceDistribution})
			properties["summary_by_app_type"] = SetSummarySubResourceData([]*models.Summary{AppsModel.SummaryByAppType})
			properties["summary_by_category"] = SetSummarySubResourceData([]*models.Summary{AppsModel.SummaryByCategory})
			properties["summary_by_origin"] = SetSummarySubResourceData([]*models.Summary{AppsModel.SummaryByOrigin})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the Apps resource defined in the Terraform configuration
func AppsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"list": {
			Description: `edge application summary list`,
			Type:        schema.TypeList, //GoType: []*AppSummary
			Elem: &schema.Resource{
				Schema: AppSummarySchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"next": {
			Description: `cursor next`,
			Type:        schema.TypeList, //GoType: Cursor
			Elem: &schema.Resource{
				Schema: CursorSchema(),
			},
			Optional: true,
		},

		"summary_by_app_instance_distribution": {
			Description: `app distribution summary`,
			Type:        schema.TypeList, //GoType: Summary
			Elem: &schema.Resource{
				Schema: SummarySchema(),
			},
			Optional: true,
		},

		"summary_by_app_type": {
			Description: `edge applications by appType`,
			Type:        schema.TypeList, //GoType: Summary
			Elem: &schema.Resource{
				Schema: SummarySchema(),
			},
			Optional: true,
		},

		"summary_by_category": {
			Description: `edge applications by category`,
			Type:        schema.TypeList, //GoType: Summary
			Elem: &schema.Resource{
				Schema: SummarySchema(),
			},
			Optional: true,
		},

		"summary_by_origin": {
			Description: `edge applications by origin`,
			Type:        schema.TypeList, //GoType: Summary
			Elem: &schema.Resource{
				Schema: SummarySchema(),
			},
			Optional: true,
		},
	}
}

// Retrieve property field names for updating the Apps resource
func GetAppsPropertyFields() (t []string) {
	return []string{
		"list",
		"next",
		"summary_by_app_instance_distribution",
		"summary_by_app_type",
		"summary_by_category",
		"summary_by_origin",
	}
}
