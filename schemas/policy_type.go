package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

// Function to perform the following actions:
// (1) Translate PolicyType resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func PolicyTypeModel(d *schema.ResourceData) *models.PolicyType {
	policyType, _ := d.Get("policy_type").(models.PolicyType)
	return &policyType
}

func PolicyTypeModelFromMap(m map[string]interface{}) *models.PolicyType {
	policyType := m["policy_type"].(models.PolicyType)
	return &policyType
}

// Update the underlying PolicyType resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetPolicyTypeResourceData(d *schema.ResourceData, m *models.PolicyType) {
}

// Iterate through and update the PolicyType resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetPolicyTypeSubResourceData(m []*models.PolicyType) (d []*map[string]interface{}) {
	for _, PolicyTypeModel := range m {
		if PolicyTypeModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the PolicyType resource defined in the Terraform configuration
func PolicyTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the PolicyType resource
func GetPolicyTypePropertyFields() (t []string) {
	return []string{}
}
