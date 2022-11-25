package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate TagStatusMsg resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func TagStatusMsgModel(d *schema.ResourceData) *models.TagStatusMsg {
	id := d.Get("id").(string)
	name := d.Get("name").(string)
	return &models.TagStatusMsg{
		ID:   id,
		Name: name,
	}
}

func TagStatusMsgModelFromMap(m map[string]interface{}) *models.TagStatusMsg {
	id := m["id"].(string)
	name := m["name"].(string)
	return &models.TagStatusMsg{
		ID:   id,
		Name: name,
	}
}

// Update the underlying TagStatusMsg resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetTagStatusMsgResourceData(d *schema.ResourceData, m *models.TagStatusMsg) {
	d.Set("id", m.ID)
	d.Set("name", m.Name)
	d.Set("status", m.Status)
}

// Iterate throught and update the TagStatusMsg resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetTagStatusMsgSubResourceData(m []*models.TagStatusMsg) (d []*map[string]interface{}) {
	for _, TagStatusMsgModel := range m {
		if TagStatusMsgModel != nil {
			properties := make(map[string]interface{})
			properties["id"] = TagStatusMsgModel.ID
			properties["name"] = TagStatusMsgModel.Name
			properties["status"] = TagStatusMsgModel.Status
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the TagStatusMsg resource defined in the Terraform configuration
func TagStatusMsgSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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
			Type:        schema.TypeList, //GoType: TagStatus
			Elem: &schema.Resource{
				Schema: TagStatusSchema(),
			},
			Computed: true,
		},
	}
}

// Retrieve property field names for updating the TagStatusMsg resource
func GetTagStatusMsgPropertyFields() (t []string) {
	return []string{
		"id",
		"name",
	}
}
