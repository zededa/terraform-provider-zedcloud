package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func SeverityModel(d *schema.ResourceData) *models.Severity {
	severity, _ := d.Get("severity").(models.Severity)
	return &severity
}

func SeverityModelFromMap(m map[string]interface{}) *models.Severity {
	severity := m["severity"].(models.Severity)
	return &severity
}

// Update the underlying Severity resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetSeverityResourceData(d *schema.ResourceData, m *models.Severity) {
}

// Iterate through and update the Severity resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetSeveritySubResourceData(m []*models.Severity) (d []*map[string]interface{}) {
	for _, SeverityModel := range m {
		if SeverityModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the Severity resource defined in the Terraform configuration
func SeveritySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the Severity resource
func GetSeverityPropertyFields() (t []string) {
	return []string{}
}
