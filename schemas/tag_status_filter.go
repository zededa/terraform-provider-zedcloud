package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate TagStatusFilter resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func TagStatusFilterModel(d *schema.ResourceData) *models.TagStatusFilter {
	namePattern, _ := d.Get("name_pattern").(string)
	var status *models.TagStatus // TagStatus
	statusInterface, statusIsSet := d.GetOk("status")
	if statusIsSet {
		statusModel := statusInterface.(string)
		status = models.NewTagStatus(models.TagStatus(statusModel))
	}
	var typeVar *models.TagType // TagType
	typeInterface, typeIsSet := d.GetOk("type")
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewTagType(models.TagType(typeModel))
	}
	return &models.TagStatusFilter{
		NamePattern: namePattern,
		Status:      status,
		Type:        typeVar,
	}
}

func TagStatusFilterModelFromMap(m map[string]interface{}) *models.TagStatusFilter {
	namePattern := m["name_pattern"].(string)
	status := m["status"].(*models.TagStatus) // TagStatus
	typeVar := m["type"].(*models.TagType)    // TagType
	return &models.TagStatusFilter{
		NamePattern: namePattern,
		Status:      status,
		Type:        typeVar,
	}
}

// Update the underlying TagStatusFilter resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetTagStatusFilterResourceData(d *schema.ResourceData, m *models.TagStatusFilter) {
	d.Set("name_pattern", m.NamePattern)
	d.Set("status", m.Status)
	d.Set("type", m.Type)
}

// Iterate through and update the TagStatusFilter resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetTagStatusFilterSubResourceData(m []*models.TagStatusFilter) (d []*map[string]interface{}) {
	for _, TagStatusFilterModel := range m {
		if TagStatusFilterModel != nil {
			properties := make(map[string]interface{})
			properties["name_pattern"] = TagStatusFilterModel.NamePattern
			properties["status"] = TagStatusFilterModel.Status
			properties["type"] = TagStatusFilterModel.Type
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the TagStatusFilter resource defined in the Terraform configuration
func TagStatusFilterSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name_pattern": {
			Description: `Resource group name pattern to be matched.`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"status": {
			Description: `Resource group status to be matched.`,
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

// Retrieve property field names for updating the TagStatusFilter resource
func GetTagStatusFilterPropertyFields() (t []string) {
	return []string{
		"name_pattern",
		"status",
		"type",
	}
}
