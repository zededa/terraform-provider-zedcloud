package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func AppInstanceLogsQueryResponseModel(d *schema.ResourceData) *models.AppInstanceLogsQueryResponse {
	var list []*models.AppInstanceLogsQueryResponseItem // []*AppInstanceLogsQueryResponseItem
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
			m := AppInstanceLogsQueryResponseItemModelFromMap(v.(map[string]interface{}))
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
	totalRecordsInt, _ := d.Get("total_records").(int)
	totalRecords := int64(totalRecordsInt)
	return &models.AppInstanceLogsQueryResponse{
		List:         list,
		Next:         next,
		TotalRecords: totalRecords,
	}
}

func AppInstanceLogsQueryResponseModelFromMap(m map[string]interface{}) *models.AppInstanceLogsQueryResponse {
	var list []*models.AppInstanceLogsQueryResponseItem // []*AppInstanceLogsQueryResponseItem
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
			m := AppInstanceLogsQueryResponseItemModelFromMap(v.(map[string]interface{}))
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
	totalRecords := int64(m["total_records"].(int)) // int64
	return &models.AppInstanceLogsQueryResponse{
		List:         list,
		Next:         next,
		TotalRecords: totalRecords,
	}
}

// Update the underlying AppInstanceLogsQueryResponse resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetAppInstanceLogsQueryResponseResourceData(d *schema.ResourceData, m *models.AppInstanceLogsQueryResponse) {
	d.Set("list", SetAppInstanceLogsQueryResponseItemSubResourceData(m.List))
	d.Set("next", SetCursorSubResourceData([]*models.Cursor{m.Next}))
	d.Set("total_records", m.TotalRecords)
}

// Iterate through and update the AppInstanceLogsQueryResponse resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetAppInstanceLogsQueryResponseSubResourceData(m []*models.AppInstanceLogsQueryResponse) (d []*map[string]interface{}) {
	for _, AppInstanceLogsQueryResponseModel := range m {
		if AppInstanceLogsQueryResponseModel != nil {
			properties := make(map[string]interface{})
			properties["list"] = SetAppInstanceLogsQueryResponseItemSubResourceData(AppInstanceLogsQueryResponseModel.List)
			properties["next"] = SetCursorSubResourceData([]*models.Cursor{AppInstanceLogsQueryResponseModel.Next})
			properties["total_records"] = AppInstanceLogsQueryResponseModel.TotalRecords
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the AppInstanceLogsQueryResponse resource defined in the Terraform configuration
func AppInstanceLogsQueryResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"list": {
			Description: `app instance logs query response`,
			Type:        schema.TypeList, //GoType: []*AppInstanceLogsQueryResponseItem
			Elem: &schema.Resource{
				Schema: AppInstanceLogsQueryResponseItemSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"next": {
			Description: `next cursor`,
			Type:        schema.TypeList, //GoType: Cursor
			Elem: &schema.Resource{
				Schema: CursorSchema(),
			},
			Optional: true,
		},

		"total_records": {
			Description: `app instance logs query response`,
			Type:        schema.TypeInt,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the AppInstanceLogsQueryResponse resource
func GetAppInstanceLogsQueryResponsePropertyFields() (t []string) {
	return []string{
		"list",
		"next",
		"total_records",
	}
}
