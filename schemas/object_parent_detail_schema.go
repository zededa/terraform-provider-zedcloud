package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate ObjectParentDetail resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ObjectParentDetailModel(d *schema.ResourceData) *models.ObjectParentDetail {
	referenceExists := d.Get("reference_exists").(bool)
	updateAvailable := d.Get("update_available").(bool)
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

// Update the underlying ObjectParentDetail resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetObjectParentDetailResourceData(d *schema.ResourceData, m *models.ObjectParentDetail) {
	d.Set("id_of_parent_object", m.IDOfParentObject)
	d.Set("reference_exists", m.ReferenceExists)
	d.Set("update_available", m.UpdateAvailable)
	d.Set("version_of_parent_object", m.VersionOfParentObject)
}

// Iterate throught and update the ObjectParentDetail resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
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

// Schema mapping representing the ObjectParentDetail resource defined in the Terraform configuration
func ObjectParentDetailSchema() map[string]*schema.Schema {
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

// Retrieve property field names for updating the ObjectParentDetail resource
func GetObjectParentDetailPropertyFields() (t []string) {
	return []string{
		"reference_exists",
		"update_available",
	}
}
