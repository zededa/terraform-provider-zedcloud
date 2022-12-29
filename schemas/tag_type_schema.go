package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate TagType resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func TagTypeModel(d *schema.ResourceData) *models.TagType {
	tagType, _ := d.Get("tag_type").(models.TagType)
	return &tagType
}

func TagTypeModelFromMap(m map[string]interface{}) *models.TagType {
	tagType := m["tag_type"].(models.TagType)
	return &tagType
}

// Update the underlying TagType resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetTagTypeResourceData(d *schema.ResourceData, m *models.TagType) {
}

// Iterate through and update the TagType resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetTagTypeSubResourceData(m []*models.TagType) (d []*map[string]interface{}) {
	for _, TagTypeModel := range m {
		if TagTypeModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the TagType resource defined in the Terraform configuration
func TagTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the TagType resource
func GetTagTypePropertyFields() (t []string) {
	return []string{}
}
