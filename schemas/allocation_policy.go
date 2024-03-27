package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

// Function to perform the following actions:
// (1) Translate AllocationPolicy resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func AllocationPolicyModel(d *schema.ResourceData) *models.AllocationPolicy {
	allocationPolicy, _ := d.Get("allocation_policy").(models.AllocationPolicy)
	return &allocationPolicy
}

func AllocationPolicyModelFromMap(m map[string]interface{}) *models.AllocationPolicy {
	allocationPolicy := m["allocation_policy"].(models.AllocationPolicy)
	return &allocationPolicy
}

// Update the underlying AllocationPolicy resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetAllocationPolicyResourceData(d *schema.ResourceData, m *models.AllocationPolicy) {
}

// Iterate through and update the AllocationPolicy resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetAllocationPolicySubResourceData(m []*models.AllocationPolicy) (d []*map[string]interface{}) {
	for _, AllocationPolicyModel := range m {
		if AllocationPolicyModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the AllocationPolicy resource defined in the Terraform configuration
func AllocationPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the AllocationPolicy resource
func GetAllocationPolicyPropertyFields() (t []string) {
	return []string{}
}
