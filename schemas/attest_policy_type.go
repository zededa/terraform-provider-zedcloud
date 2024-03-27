package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

// Function to perform the following actions:
// (1) Translate AttestPolicyType resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func AttestPolicyTypeModel(d *schema.ResourceData) *models.AttestPolicyType {
	attestPolicyType, _ := d.Get("attest_policy_type").(models.AttestPolicyType)
	return &attestPolicyType
}

func AttestPolicyTypeModelFromMap(m map[string]interface{}) *models.AttestPolicyType {
	attestPolicyType := m["attest_policy_type"].(models.AttestPolicyType)
	return &attestPolicyType
}

// Update the underlying AttestPolicyType resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetAttestPolicyTypeResourceData(d *schema.ResourceData, m *models.AttestPolicyType) {
}

// Iterate through and update the AttestPolicyType resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetAttestPolicyTypeSubResourceData(m []*models.AttestPolicyType) (d []*map[string]interface{}) {
	for _, AttestPolicyTypeModel := range m {
		if AttestPolicyTypeModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the AttestPolicyType resource defined in the Terraform configuration
func AttestPolicyTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the AttestPolicyType resource
func GetAttestPolicyTypePropertyFields() (t []string) {
	return []string{}
}
