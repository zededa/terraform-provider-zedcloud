package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

func ObjectDetailModel(d *schema.ResourceData) *models.ObjectDetail {
	id, _ := d.Get("id").(string)
	name, _ := d.Get("name").(string)
	return &models.ObjectDetail{
		ID:   id,
		Name: name,
	}
}

func ObjectDetailModelFromMap(m map[string]interface{}) *models.ObjectDetail {
	id := m["id"].(string)
	name := m["name"].(string)
	return &models.ObjectDetail{
		ID:   id,
		Name: name,
	}
}

// Update the underlying ObjectDetail resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetObjectDetailResourceData(d *schema.ResourceData, m *models.ObjectDetail) {
	d.Set("id", m.ID)
	d.Set("name", m.Name)
}

// Iterate through and update the ObjectDetail resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetObjectDetailSubResourceData(m []*models.ObjectDetail) (d []*map[string]interface{}) {
	for _, ObjectDetailModel := range m {
		if ObjectDetailModel != nil {
			properties := make(map[string]interface{})
			properties["id"] = ObjectDetailModel.ID
			properties["name"] = ObjectDetailModel.Name
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the ObjectDetail resource defined in the Terraform configuration
func ObjectDetailSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Description: ``,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"name": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the ObjectDetail resource
func GetObjectDetailPropertyFields() (t []string) {
	return []string{
		"id",
		"name",
	}
}
