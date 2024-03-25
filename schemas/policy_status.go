package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

// Function to perform the following actions:
// (1) Translate PolicyStatus resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func PolicyStatusModel(d *schema.ResourceData) *models.PolicyStatus {
	policyStatus, _ := d.Get("policy_status").(models.PolicyStatus)
	return &policyStatus
}

func PolicyStatusModelFromMap(m map[string]interface{}) *models.PolicyStatus {
	policyStatus := m["policy_status"].(models.PolicyStatus)
	return &policyStatus
}

// Update the underlying PolicyStatus resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetPolicyStatusResourceData(d *schema.ResourceData, m *models.PolicyStatus) {
}

// Iterate through and update the PolicyStatus resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetPolicyStatusSubResourceData(m []*models.PolicyStatus) (d []*map[string]interface{}) {
	for _, PolicyStatusModel := range m {
		if PolicyStatusModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the PolicyStatus resource defined in the Terraform configuration
func PolicyStatusSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the PolicyStatus resource
func GetPolicyStatusPropertyFields() (t []string) {
	return []string{}
}
