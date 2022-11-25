package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate ECDSA resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ECDSAModel(d *schema.ResourceData) *models.ECDSA {
	curve := d.Get("curve").(string)
	return &models.ECDSA{
		Curve: curve,
	}
}

func ECDSAModelFromMap(m map[string]interface{}) *models.ECDSA {
	curve := m["curve"].(string)
	return &models.ECDSA{
		Curve: curve,
	}
}

// Update the underlying ECDSA resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetECDSAResourceData(d *schema.ResourceData, m *models.ECDSA) {
	d.Set("curve", m.Curve)
}

// Iterate throught and update the ECDSA resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetECDSASubResourceData(m []*models.ECDSA) (d []*map[string]interface{}) {
	for _, ECDSAModel := range m {
		if ECDSAModel != nil {
			properties := make(map[string]interface{})
			properties["curve"] = ECDSAModel.Curve
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the ECDSA resource defined in the Terraform configuration
func ECDSASchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"curve": {
			Description: `ECDSA curve to be used while signing the certificate.`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the ECDSA resource
func GetECDSAPropertyFields() (t []string) {
	return []string{
		"curve",
	}
}
