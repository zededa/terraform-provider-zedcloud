package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ObjectParentDetailModel(d *schema.ResourceData) *models.ObjectParentDetail {
	referenceExists, _ := d.Get("reference_exists").(bool)
	updateAvailable, _ := d.Get("update_available").(bool)
	return &models.ObjectParentDetail{
		ReferenceExists: referenceExists,
		UpdateAvailable: updateAvailable,
	}
}

func ObjectParentDetailModelFromMap(m map[string]interface{}) *models.ObjectParentDetail {
	referenceExists := m["reference_exists"].(bool)
	updateAvailable := m["update_available"].(bool)
	return &models.ObjectParentDetail{
		ReferenceExists: referenceExists,
		UpdateAvailable: updateAvailable,
	}
}

func SetObjectParentDetailResourceData(d *schema.ResourceData, m *models.ObjectParentDetail) {
	d.Set("id_of_parent_object", m.IDOfParentObject)
	d.Set("reference_exists", m.ReferenceExists)
	d.Set("update_available", m.UpdateAvailable)
	d.Set("version_of_parent_object", m.VersionOfParentObject)
}

func SetObjectParentDetailSubResourceData(m []*models.ObjectParentDetail) (d []*map[string]interface{}) {
	for _, ObjectParentDetailModel := range m {
		if ObjectParentDetailModel != nil {
			properties := make(map[string]interface{})
			properties["id_of_parent_object"] = ObjectParentDetailModel.IDOfParentObject
			properties["reference_exists"] = ObjectParentDetailModel.ReferenceExists
			properties["update_available"] = ObjectParentDetailModel.UpdateAvailable
			properties["version_of_parent_object"] = ObjectParentDetailModel.VersionOfParentObject
			d = append(d, &properties)
		}
	}
	return
}

func ObjectParentDetail() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id_of_parent_object": {
			Description: `system defined unique id of parent object`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"reference_exists": {
			Description: `Relation with child and parent object exists or not`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"update_available": {
			Description: `Update required flag`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"version_of_parent_object": {
			Description: `version of object present in parent`,
			Type:        schema.TypeInt,
			Computed:    true,
		},
	}
}

func GetObjectParentDetailPropertyFields() (t []string) {
	return []string{
		"reference_exists",
		"update_available",
	}
}
