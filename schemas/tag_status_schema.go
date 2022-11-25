package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate TagStatus resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func TagStatusModel(d *schema.ResourceData) *models.TagStatus {
	tagStatus := d.Get("tag_status").(models.TagStatus)
	return &tagStatus
}

func TagStatusModelFromMap(m map[string]interface{}) *models.TagStatus {
	tagStatus := m["tag_status"].(models.TagStatus)
	return &tagStatus
}

// Update the underlying TagStatus resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetTagStatusResourceData(d *schema.ResourceData, m *models.TagStatus) {
}

// Iterate throught and update the TagStatus resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetTagStatusSubResourceData(m []*models.TagStatus) (d []*map[string]interface{}) {
	for _, TagStatusModel := range m {
		if TagStatusModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the TagStatus resource defined in the Terraform configuration
func TagStatusSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the TagStatus resource
func GetTagStatusPropertyFields() (t []string) {
	return []string{}
}
