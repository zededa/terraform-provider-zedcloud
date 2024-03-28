package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func TopTalkersResponseModel(d *schema.ResourceData) *models.TopTalkersResponse {
	var list []*models.TopTalkersResponseItem // []*TopTalkersResponseItem
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
			m := TopTalkersResponseItemModelFromMap(v.(map[string]interface{}))
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
	return &models.TopTalkersResponse{
		List: list,
		Next: next,
	}
}

func TopTalkersResponseModelFromMap(m map[string]interface{}) *models.TopTalkersResponse {
	var list []*models.TopTalkersResponseItem // []*TopTalkersResponseItem
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
			m := TopTalkersResponseItemModelFromMap(v.(map[string]interface{}))
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
	return &models.TopTalkersResponse{
		List: list,
		Next: next,
	}
}

// Update the underlying TopTalkersResponse resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetTopTalkersResponseResourceData(d *schema.ResourceData, m *models.TopTalkersResponse) {
	d.Set("list", SetTopTalkersResponseItemSubResourceData(m.List))
	d.Set("next", SetCursorSubResourceData([]*models.Cursor{m.Next}))
}

// Iterate through and update the TopTalkersResponse resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetTopTalkersResponseSubResourceData(m []*models.TopTalkersResponse) (d []*map[string]interface{}) {
	for _, TopTalkersResponseModel := range m {
		if TopTalkersResponseModel != nil {
			properties := make(map[string]interface{})
			properties["list"] = SetTopTalkersResponseItemSubResourceData(TopTalkersResponseModel.List)
			properties["next"] = SetCursorSubResourceData([]*models.Cursor{TopTalkersResponseModel.Next})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the TopTalkersResponse resource defined in the Terraform configuration
func TopTalkersResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"list": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []*TopTalkersResponseItem
			Elem: &schema.Resource{
				Schema: TopTalkersResponseItemSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"next": {
			Description: ``,
			Type:        schema.TypeList, //GoType: Cursor
			Elem: &schema.Resource{
				Schema: CursorSchema(),
			},
			Optional: true,
		},
	}
}

// Retrieve property field names for updating the TopTalkersResponse resource
func GetTopTalkersResponsePropertyFields() (t []string) {
	return []string{
		"list",
		"next",
	}
}
