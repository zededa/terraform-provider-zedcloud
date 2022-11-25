package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate RSA resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func RSAModel(d *schema.ResourceData) *models.RSA {
	rsaBits := d.Get("rsa_bits").(string)
	return &models.RSA{
		RsaBits: rsaBits,
	}
}

func RSAModelFromMap(m map[string]interface{}) *models.RSA {
	rsaBits := m["rsa_bits"].(string)
	return &models.RSA{
		RsaBits: rsaBits,
	}
}

// Update the underlying RSA resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetRSAResourceData(d *schema.ResourceData, m *models.RSA) {
	d.Set("rsa_bits", m.RsaBits)
}

// Iterate throught and update the RSA resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetRSASubResourceData(m []*models.RSA) (d []*map[string]interface{}) {
	for _, RSAModel := range m {
		if RSAModel != nil {
			properties := make(map[string]interface{})
			properties["rsa_bits"] = RSAModel.RsaBits
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the RSA resource defined in the Terraform configuration
func RSASchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"rsa_bits": {
			Description: `RSA Encryption Key bit size.`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the RSA resource
func GetRSAPropertyFields() (t []string) {
	return []string{
		"rsa_bits",
	}
}
