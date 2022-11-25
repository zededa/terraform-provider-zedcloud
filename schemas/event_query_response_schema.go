package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate EventQueryResponse resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func EventQueryResponseModel(d *schema.ResourceData) *models.EventQueryResponse {
	list := d.Get("list").([]*models.EventQueryResponseItem) // []*EventQueryResponseItem
	var next *models.Cursor                                  // Cursor
	nextInterface, nextIsSet := d.GetOk("next")
	if nextIsSet {
		nextMap := nextInterface.([]interface{})[0].(map[string]interface{})
		next = CursorModelFromMap(nextMap)
	}
	return &models.EventQueryResponse{
		List: list,
		Next: next,
	}
}

func EventQueryResponseModelFromMap(m map[string]interface{}) *models.EventQueryResponse {
	list := m["list"].([]*models.EventQueryResponseItem) // []*EventQueryResponseItem
	var next *models.Cursor                              // Cursor
	nextInterface, nextIsSet := m["next"]
	if nextIsSet {
		nextMap := nextInterface.([]interface{})[0].(map[string]interface{})
		next = CursorModelFromMap(nextMap)
	}
	//
	return &models.EventQueryResponse{
		List: list,
		Next: next,
	}
}

// Update the underlying EventQueryResponse resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetEventQueryResponseResourceData(d *schema.ResourceData, m *models.EventQueryResponse) {
	d.Set("list", SetEventQueryResponseItemSubResourceData(m.List))
	d.Set("next", SetCursorSubResourceData([]*models.Cursor{m.Next}))
}

// Iterate throught and update the EventQueryResponse resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetEventQueryResponseSubResourceData(m []*models.EventQueryResponse) (d []*map[string]interface{}) {
	for _, EventQueryResponseModel := range m {
		if EventQueryResponseModel != nil {
			properties := make(map[string]interface{})
			properties["list"] = SetEventQueryResponseItemSubResourceData(EventQueryResponseModel.List)
			properties["next"] = SetCursorSubResourceData([]*models.Cursor{EventQueryResponseModel.Next})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the EventQueryResponse resource defined in the Terraform configuration
func EventQueryResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"list": {
			Description: `Event Query response list`,
			Type:        schema.TypeList, //GoType: []*EventQueryResponseItem
			Elem: &schema.Resource{
				Schema: EventQueryResponseItemSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Required:   true,
		},

		"next": {
			Description: `Cursor filter`,
			Type:        schema.TypeList, //GoType: Cursor
			Elem: &schema.Resource{
				Schema: CursorSchema(),
			},
			Required: true,
		},
	}
}

// Retrieve property field names for updating the EventQueryResponse resource
func GetEventQueryResponsePropertyFields() (t []string) {
	return []string{
		"list",
		"next",
	}
}
