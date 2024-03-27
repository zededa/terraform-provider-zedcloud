package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

func ClassificationResponseModel(d *schema.ResourceData) *models.ClassificationResponse {
	var list []*models.ClassificationItem // []*ClassificationItem
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
			m := ClassificationItemModelFromMap(v.(map[string]interface{}))
			list = append(list, m)
		}
	}
	return &models.ClassificationResponse{
		List: list,
	}
}

func ClassificationResponseModelFromMap(m map[string]interface{}) *models.ClassificationResponse {
	var list []*models.ClassificationItem // []*ClassificationItem
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
			m := ClassificationItemModelFromMap(v.(map[string]interface{}))
			list = append(list, m)
		}
	}
	return &models.ClassificationResponse{
		List: list,
	}
}

// Update the underlying ClassificationResponse resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetClassificationResponseResourceData(d *schema.ResourceData, m *models.ClassificationResponse) {
	d.Set("list", SetClassificationItemSubResourceData(m.List))
}

// Iterate through and update the ClassificationResponse resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetClassificationResponseSubResourceData(m []*models.ClassificationResponse) (d []*map[string]interface{}) {
	for _, ClassificationResponseModel := range m {
		if ClassificationResponseModel != nil {
			properties := make(map[string]interface{})
			properties["list"] = SetClassificationItemSubResourceData(ClassificationResponseModel.List)
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the ClassificationResponse resource defined in the Terraform configuration
func ClassificationResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"list": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []*ClassificationItem
			Elem: &schema.Resource{
				Schema: ClassificationItemSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
	}
}

// Retrieve property field names for updating the ClassificationResponse resource
func GetClassificationResponsePropertyFields() (t []string) {
	return []string{
		"list",
	}
}
