package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate TagStatusMsg resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func TagStatusMsgModel(d *schema.ResourceData) *models.TagStatusMsg {
	edgeviewSessionCount := int64(d.Get("edgeview_session_count").(int))
	id := d.Get("id").(string)
	name := d.Get("name").(string)
	typeVar := d.Get("type").(*models.TagType) // TagType
	return &models.TagStatusMsg{
		EdgeviewSessionCount: edgeviewSessionCount,
		ID:                   id,
		Name:                 name,
		Type:                 typeVar,
	}
}

func TagStatusMsgModelFromMap(m map[string]interface{}) *models.TagStatusMsg {
	edgeviewSessionCount := int64(m["edgeview_session_count"].(int)) // int64 false false false
	id := m["id"].(string)
	name := m["name"].(string)
	typeVar := m["type"].(*models.TagType) // TagType
	return &models.TagStatusMsg{
		EdgeviewSessionCount: edgeviewSessionCount,
		ID:                   id,
		Name:                 name,
		Type:                 typeVar,
	}
}

// Update the underlying TagStatusMsg resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetTagStatusMsgResourceData(d *schema.ResourceData, m *models.TagStatusMsg) {
	d.Set("edgeview_session_count", m.EdgeviewSessionCount)
	d.Set("id", m.ID)
	d.Set("name", m.Name)
	d.Set("status", m.Status)
	d.Set("type", m.Type)
}

// Iterate throught and update the TagStatusMsg resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetTagStatusMsgSubResourceData(m []*models.TagStatusMsg) (d []*map[string]interface{}) {
	for _, TagStatusMsgModel := range m {
		if TagStatusMsgModel != nil {
			properties := make(map[string]interface{})
			properties["edgeview_session_count"] = TagStatusMsgModel.EdgeviewSessionCount
			properties["id"] = TagStatusMsgModel.ID
			properties["name"] = TagStatusMsgModel.Name
			properties["status"] = TagStatusMsgModel.Status
			properties["type"] = TagStatusMsgModel.Type
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the TagStatusMsg resource defined in the Terraform configuration
func TagStatusMsgSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"edgeview_session_count": {
			Description: `total count of devices enabled with edgeview session`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"id": {
			Description: `System defined universally unique Id of the resource group.`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"name": {
			Description: `User defined name of the resource group, unique across the enterprise. Once resource group is created, name can’t be changed.`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"status": {
			Description: `Resource group status`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"type": {
			Description: `Resource group type`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the TagStatusMsg resource
func GetTagStatusMsgPropertyFields() (t []string) {
	return []string{
		"edgeview_session_count",
		"id",
		"name",
		"type",
	}
}
