package schemas

import (
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

func ObjectRevisionModel(d *schema.ResourceData) *models.ObjectRevision {
	createdAt, _ := d.Get("created_at").(strfmt.DateTime)
	createdBy, _ := d.Get("created_by").(string)
	curr, _ := d.Get("curr").(string)
	prev, _ := d.Get("prev").(string)
	updatedAt, _ := d.Get("updated_at").(strfmt.DateTime)
	updatedBy, _ := d.Get("updated_by").(string)
	return &models.ObjectRevision{
		CreatedAt: &createdAt,
		CreatedBy: &createdBy,
		Curr:      &curr,
		Prev:      prev,
		UpdatedAt: &updatedAt,
		UpdatedBy: &updatedBy,
	}
}

func ObjectRevisionModelFromMap(m map[string]interface{}) *models.ObjectRevision {
	t, err := time.Parse(time.RFC3339, m["created_at"].(string))
	if err != nil {
		panic(err)
	}
	createdAt := strfmt.DateTime(t)

	createdBy := m["created_by"].(string)
	curr := m["curr"].(string)
	prev := m["prev"].(string)

	t, err = time.Parse(time.RFC3339, m["updated_at"].(string))
	if err != nil {
		panic(err)
	}
	updatedAt := strfmt.DateTime(t)
	updatedBy := m["updated_by"].(string)
	return &models.ObjectRevision{
		CreatedAt: &createdAt,
		CreatedBy: &createdBy,
		Curr:      &curr,
		Prev:      prev,
		UpdatedAt: &updatedAt,
		UpdatedBy: &updatedBy,
	}
}

func SetObjectRevisionResourceData(d *schema.ResourceData, m *models.ObjectRevision) {
	d.Set("created_at", m.CreatedAt)
	d.Set("created_by", m.CreatedBy)
	d.Set("curr", m.Curr)
	d.Set("prev", m.Prev)
	d.Set("updated_at", m.UpdatedAt)
	d.Set("updated_by", m.UpdatedBy)
}

func SetObjectRevisionSubResourceData(m []*models.ObjectRevision) (d []*map[string]interface{}) {
	for _, ObjectRevisionModel := range m {
		if ObjectRevisionModel != nil {
			properties := make(map[string]interface{})
			properties["created_at"] = ObjectRevisionModel.CreatedAt.String()
			properties["created_by"] = ObjectRevisionModel.CreatedBy
			properties["curr"] = ObjectRevisionModel.Curr
			properties["prev"] = ObjectRevisionModel.Prev
			properties["updated_at"] = ObjectRevisionModel.UpdatedAt.String()
			properties["updated_by"] = ObjectRevisionModel.UpdatedBy
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the ObjectRevision resource defined in the Terraform configuration
func ObjectRevision() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"created_at": {
			Description: `The time, in milliseconds since the epoch, when the record was created.`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"created_by": {
			Description: `User data: Created By`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"curr": {
			Description: `Current Database version of the record`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"prev": {
			Description: `Previous`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"updated_at": {
			Description: `The time, in milliseconds since the epoch, when the record was last updated.`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"updated_by": {
			Description: `User data: Updated By`,
			Type:        schema.TypeString,
			Computed:    true,
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
