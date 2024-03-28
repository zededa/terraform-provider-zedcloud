package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ObjectTagsListModel(d *schema.ResourceData) *models.ObjectTagsList {
	var next *models.Cursor // Cursor
	nextInterface, nextIsSet := d.GetOk("next")
	if nextIsSet && nextInterface != nil {
		nextMap := nextInterface.([]interface{})
		if len(nextMap) > 0 {
			next = CursorModelFromMap(nextMap[0].(map[string]interface{}))
		}
	}
	var objectTags []*models.ObjectTag // []*ObjectTag
	objectTagsInterface, objectTagsIsSet := d.GetOk("object_tags")
	if objectTagsIsSet {
		var items []interface{}
		if listItems, isList := objectTagsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = objectTagsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := ObjectTagModelFromMap(v.(map[string]interface{}))
			objectTags = append(objectTags, m)
		}
	}
	return &models.ObjectTagsList{
		Next:       next,
		ObjectTags: objectTags,
	}
}

func ObjectTagsListModelFromMap(m map[string]interface{}) *models.ObjectTagsList {
	var next *models.Cursor // Cursor
	nextInterface, nextIsSet := m["next"]
	if nextIsSet && nextInterface != nil {
		nextMap := nextInterface.([]interface{})
		if len(nextMap) > 0 {
			next = CursorModelFromMap(nextMap[0].(map[string]interface{}))
		}
	}
	//
	var objectTags []*models.ObjectTag // []*ObjectTag
	objectTagsInterface, objectTagsIsSet := m["object_tags"]
	if objectTagsIsSet {
		var items []interface{}
		if listItems, isList := objectTagsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = objectTagsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := ObjectTagModelFromMap(v.(map[string]interface{}))
			objectTags = append(objectTags, m)
		}
	}
	return &models.ObjectTagsList{
		Next:       next,
		ObjectTags: objectTags,
	}
}

// Update the underlying ObjectTagsList resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetObjectTagsListResourceData(d *schema.ResourceData, m *models.ObjectTagsList) {
	d.Set("next", SetCursorSubResourceData([]*models.Cursor{m.Next}))
	d.Set("object_tags", SetObjectTagSubResourceData(m.ObjectTags))
}

// Iterate through and update the ObjectTagsList resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetObjectTagsListSubResourceData(m []*models.ObjectTagsList) (d []*map[string]interface{}) {
	for _, ObjectTagsListModel := range m {
		if ObjectTagsListModel != nil {
			properties := make(map[string]interface{})
			properties["next"] = SetCursorSubResourceData([]*models.Cursor{ObjectTagsListModel.Next})
			properties["object_tags"] = SetObjectTagSubResourceData(ObjectTagsListModel.ObjectTags)
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the ObjectTagsList resource defined in the Terraform configuration
func ObjectTagsListSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"next": {
			Description: `Returned record page.`,
			Type:        schema.TypeList, //GoType: Cursor
			Elem: &schema.Resource{
				Schema: CursorSchema(),
			},
			Optional: true,
		},

		"object_tags": {
			Description: `Details of Object tags list records.`,
			Type:        schema.TypeList, //GoType: []*ObjectTag
			Elem: &schema.Resource{
				Schema: ObjectTagSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
	}
}

// Retrieve property field names for updating the ObjectTagsList resource
func GetObjectTagsListPropertyFields() (t []string) {
	return []string{
		"next",
		"object_tags",
	}
}
