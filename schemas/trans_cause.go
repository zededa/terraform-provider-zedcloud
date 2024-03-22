package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

func TransCauseModel(d *schema.ResourceData) *models.TransCause {
	transCause, _ := d.Get("trans_cause").(models.TransCause)
	return &transCause
}

func TransCauseModelFromMap(m map[string]interface{}) *models.TransCause {
	transCause := m["trans_cause"].(models.TransCause)
	return &transCause
}

// Update the underlying TransCause resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetTransCauseResourceData(d *schema.ResourceData, m *models.TransCause) {
}

// Iterate through and update the TransCause resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetTransCauseSubResourceData(m []*models.TransCause) (d []*map[string]interface{}) {
	for _, TransCauseModel := range m {
		if TransCauseModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the TransCause resource defined in the Terraform configuration
func TransCauseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the TransCause resource
func GetTransCausePropertyFields() (t []string) {
	return []string{}
}
