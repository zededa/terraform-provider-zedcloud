package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func DatastoresModel(d *schema.ResourceData) *models.Datastores {
	var list []*models.Datastore // []*DatastoreInfo
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
			m := DatastoreModelFromMap(v.(map[string]interface{}))
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
	var summaryByCategory *models.Summary // Summary
	summaryByCategoryInterface, summaryByCategoryIsSet := d.GetOk("summary_by_category")
	if summaryByCategoryIsSet && summaryByCategoryInterface != nil {
		summaryByCategoryMap := summaryByCategoryInterface.([]interface{})
		if len(summaryByCategoryMap) > 0 {
			summaryByCategory = SummaryModelFromMap(summaryByCategoryMap[0].(map[string]interface{}))
		}
	}
	var summaryByType *models.Summary // Summary
	summaryByTypeInterface, summaryByTypeIsSet := d.GetOk("summary_by_type")
	if summaryByTypeIsSet && summaryByTypeInterface != nil {
		summaryByTypeMap := summaryByTypeInterface.([]interface{})
		if len(summaryByTypeMap) > 0 {
			summaryByType = SummaryModelFromMap(summaryByTypeMap[0].(map[string]interface{}))
		}
	}
	return &models.Datastores{
		List:              list,
		Next:              next,
		SummaryByCategory: summaryByCategory,
		SummaryByType:     summaryByType,
	}
}

func DatastoresModelFromMap(m map[string]interface{}) *models.Datastores {
	var list []*models.Datastore // []*DatastoreInfo
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
			m := DatastoreModelFromMap(v.(map[string]interface{}))
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
	var summaryByCategory *models.Summary // Summary
	summaryByCategoryInterface, summaryByCategoryIsSet := m["summary_by_category"]
	if summaryByCategoryIsSet && summaryByCategoryInterface != nil {
		summaryByCategoryMap := summaryByCategoryInterface.([]interface{})
		if len(summaryByCategoryMap) > 0 {
			summaryByCategory = SummaryModelFromMap(summaryByCategoryMap[0].(map[string]interface{}))
		}
	}
	//
	var summaryByType *models.Summary // Summary
	summaryByTypeInterface, summaryByTypeIsSet := m["summary_by_type"]
	if summaryByTypeIsSet && summaryByTypeInterface != nil {
		summaryByTypeMap := summaryByTypeInterface.([]interface{})
		if len(summaryByTypeMap) > 0 {
			summaryByType = SummaryModelFromMap(summaryByTypeMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.Datastores{
		List:              list,
		Next:              next,
		SummaryByCategory: summaryByCategory,
		SummaryByType:     summaryByType,
	}
}

// Update the underlying Datastores resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDatastoresResourceData(d *schema.ResourceData, m *models.Datastores) {
	d.Set("list", SetDatastoreInfoSubResourceData(m.List))
	d.Set("next", SetCursorSubResourceData([]*models.Cursor{m.Next}))
	d.Set("summary_by_category", SetSummarySubResourceData([]*models.Summary{m.SummaryByCategory}))
	d.Set("summary_by_type", SetSummarySubResourceData([]*models.Summary{m.SummaryByType}))
}

// Iterate through and update the Datastores resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDatastoresSubResourceData(m []*models.Datastores) (d []*map[string]interface{}) {
	for _, DatastoresModel := range m {
		if DatastoresModel != nil {
			properties := make(map[string]interface{})
			properties["list"] = SetDatastoreInfoSubResourceData(DatastoresModel.List)
			properties["next"] = SetCursorSubResourceData([]*models.Cursor{DatastoresModel.Next})
			properties["summary_by_category"] = SetSummarySubResourceData([]*models.Summary{DatastoresModel.SummaryByCategory})
			properties["summary_by_type"] = SetSummarySubResourceData([]*models.Summary{DatastoresModel.SummaryByType})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the Datastores resource defined in the Terraform configuration
func Datastores() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"list": {
			Description: `List of filtered Datastore records`,
			Type:        schema.TypeList, //GoType: []*DatastoreInfo
			Elem: &schema.Resource{
				Schema: Datastore(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"next": {
			Description: `Responded page details of filtered records`,
			Type:        schema.TypeList, //GoType: Cursor
			Elem: &schema.Resource{
				Schema: CursorSchema(),
			},
			Optional: true,
		},

		"summary_by_category": {
			Description: `Category distribution summary of filtered Datastore records`,
			Type:        schema.TypeList, //GoType: Summary
			Elem: &schema.Resource{
				Schema: SummarySchema(),
			},
			Optional: true,
		},

		"summary_by_type": {
			Description: `Type distribution summary of filtered Datastore records`,
			Type:        schema.TypeList, //GoType: Summary
			Elem: &schema.Resource{
				Schema: SummarySchema(),
			},
			Optional: true,
		},
	}
}

// Retrieve property field names for updating the Datastores resource
func GetDatastoresPropertyFields() (t []string) {
	return []string{
		"list",
		"next",
		"summary_by_category",
		"summary_by_type",
	}
}
