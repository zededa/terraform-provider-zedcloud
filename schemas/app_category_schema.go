package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate AppCategory resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func AppCategoryModel(d *schema.ResourceData) *models.AppCategory {
	appCategory, _ := d.Get("app_category").(models.AppCategory)
	return &appCategory
}

func AppCategoryModelFromMap(m map[string]interface{}) *models.AppCategory {
	appCategory := m["app_category"].(models.AppCategory)
	return &appCategory
}

// Update the underlying AppCategory resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetAppCategoryResourceData(d *schema.ResourceData, m *models.AppCategory) {
}

// Iterate through and update the AppCategory resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetAppCategorySubResourceData(m []*models.AppCategory) (d []*map[string]interface{}) {
	for _, AppCategoryModel := range m {
		if AppCategoryModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the AppCategory resource defined in the Terraform configuration
func AppCategorySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the AppCategory resource
func GetAppCategoryPropertyFields() (t []string) {
	return []string{}
}
