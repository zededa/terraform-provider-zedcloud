package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate AttestationPolicy resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func AttestationPolicyModel(d *schema.ResourceData) *models.AttestationPolicy {
	id := d.Get("id").(string)
	typeVar := d.Get("type").(*models.AttestPolicyType) // AttestPolicyType
	return &models.AttestationPolicy{
		ID:   id,
		Type: typeVar,
	}
}

func AttestationPolicyModelFromMap(m map[string]interface{}) *models.AttestationPolicy {
	id := m["id"].(string)
	typeVar := m["type"].(*models.AttestPolicyType) // AttestPolicyType
	return &models.AttestationPolicy{
		ID:   id,
		Type: typeVar,
	}
}

// Update the underlying AttestationPolicy resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetAttestationPolicyResourceData(d *schema.ResourceData, m *models.AttestationPolicy) {
	d.Set("id", m.ID)
	d.Set("type", m.Type)
}

// Iterate throught and update the AttestationPolicy resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetAttestationPolicySubResourceData(m []*models.AttestationPolicy) (d []*map[string]interface{}) {
	for _, AttestationPolicyModel := range m {
		if AttestationPolicyModel != nil {
			properties := make(map[string]interface{})
			properties["id"] = AttestationPolicyModel.ID
			properties["type"] = AttestationPolicyModel.Type
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the AttestationPolicy resource defined in the Terraform configuration
func AttestationPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Description: `unique policy id`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"type": {
			Description: `Attestation policy type`,
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

// Retrieve property field names for updating the AttestationPolicy resource
func GetAttestationPolicyPropertyFields() (t []string) {
	return []string{
		"id",
		"type",
	}
}
