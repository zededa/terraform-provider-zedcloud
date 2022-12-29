package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate TagFilter resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func TagFilterModel(d *schema.ResourceData) *models.TagFilter {
	namePattern, _ := d.Get("name_pattern").(string)
	typeVarModel, _ := d.Get("type").(models.TagType) // TagType
	typeVar := &typeVarModel
	if !ok {
		typeVar = nil
	}
	return &models.TagFilter{
		NamePattern: namePattern,
		Type:        typeVar,
	}
}

func TagFilterModelFromMap(m map[string]interface{}) *models.TagFilter {
	namePattern := m["name_pattern"].(string)
	typeVar := m["type"].(*models.TagType) // TagType
	return &models.TagFilter{
		NamePattern: namePattern,
		Type:        typeVar,
	}
}

// Update the underlying TagFilter resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetTagFilterResourceData(d *schema.ResourceData, m *models.TagFilter) {
	d.Set("name_pattern", m.NamePattern)
	d.Set("type", m.Type)
}

// Iterate through and update the TagFilter resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetTagFilterSubResourceData(m []*models.TagFilter) (d []*map[string]interface{}) {
	for _, TagFilterModel := range m {
		if TagFilterModel != nil {
			properties := make(map[string]interface{})
			properties["name_pattern"] = TagFilterModel.NamePattern
			properties["type"] = TagFilterModel.Type
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the TagFilter resource defined in the Terraform configuration
func TagFilterSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name_pattern": {
			Description: `Resource group name pattern to be matched.`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"type": {
			Description: `Resource group type to ne matched.`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the TagFilter resource
func GetTagFilterPropertyFields() (t []string) {
	return []string{
		"name_pattern",
		"type",
	}
}
