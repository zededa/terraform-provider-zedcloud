package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate ObjectRevision resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ObjectRevisionModel(d *schema.ResourceData) *models.ObjectRevision {
	createdAt, _ := d.Get("created_at").(interface{}) // interface{}
	createdBy, _ := d.Get("created_by").(string)
	curr, _ := d.Get("curr").(string)
	prev, _ := d.Get("prev").(string)
	updatedAt, _ := d.Get("updated_at").(interface{}) // interface{}
	updatedBy, _ := d.Get("updated_by").(string)
	return &models.ObjectRevision{
		CreatedAt: createdAt,
		CreatedBy: &createdBy, // string true false false
		Curr:      &curr,      // string true false false
		Prev:      prev,
		UpdatedAt: updatedAt,
		UpdatedBy: &updatedBy, // string true false false
	}
}

func ObjectRevisionModelFromMap(m map[string]interface{}) *models.ObjectRevision {
	createdAt := m["created_at"].(interface{})
	createdBy := m["created_by"].(string)
	curr := m["curr"].(string)
	prev := m["prev"].(string)
	updatedAt := m["updated_at"].(interface{})
	updatedBy := m["updated_by"].(string)
	return &models.ObjectRevision{
		CreatedAt: createdAt,
		CreatedBy: &createdBy,
		Curr:      &curr,
		Prev:      prev,
		UpdatedAt: updatedAt,
		UpdatedBy: &updatedBy,
	}
}

// Update the underlying ObjectRevision resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetObjectRevisionResourceData(d *schema.ResourceData, m *models.ObjectRevision) {
	d.Set("created_at", m.CreatedAt)
	d.Set("created_by", m.CreatedBy)
	d.Set("curr", m.Curr)
	d.Set("prev", m.Prev)
	d.Set("updated_at", m.UpdatedAt)
	d.Set("updated_by", m.UpdatedBy)
}

// Iterate through and update the ObjectRevision resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetObjectRevisionSubResourceData(m []*models.ObjectRevision) (d []*map[string]interface{}) {
	for _, ObjectRevisionModel := range m {
		if ObjectRevisionModel != nil {
			properties := make(map[string]interface{})
			properties["created_at"] = ObjectRevisionModel.CreatedAt
			properties["created_by"] = ObjectRevisionModel.CreatedBy
			properties["curr"] = ObjectRevisionModel.Curr
			properties["prev"] = ObjectRevisionModel.Prev
			properties["updated_at"] = ObjectRevisionModel.UpdatedAt
			properties["updated_by"] = ObjectRevisionModel.UpdatedBy
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the ObjectRevision resource defined in the Terraform configuration
func ObjectRevisionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"created_at": {
			Description: `The time, in milliseconds since the epoch, when the record was created.`,
			Type:        schema.TypeMap, //GoType: interface{}
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Required: true,
		},

		"created_by": {
			Description: `User data: Created By`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"curr": {
			Description: `Current Database version of the record`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"prev": {
			Description: `Previous`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"updated_at": {
			Description: `The time, in milliseconds since the epoch, when the record was last updated.`,
			Type:        schema.TypeMap, //GoType: interface{}
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Required: true,
		},

		"updated_by": {
			Description: `User data: Updated By`,
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

// Retrieve property field names for updating the ObjectRevision resource
func GetObjectRevisionPropertyFields() (t []string) {
	return []string{
		"created_at",
		"created_by",
		"curr",
		"prev",
		"updated_at",
		"updated_by",
	}
}
