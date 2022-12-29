package schemas

import (
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate DeviceCerts resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func DeviceCertsModel(d *schema.ResourceData) *models.DeviceCerts {
	pemCert, _ := d.Get("pem_cert").(strfmt.Base64)
	pemKey, _ := d.Get("pem_key").(strfmt.Base64)
	return &models.DeviceCerts{
		PemCert: pemCert,
		PemKey:  pemKey,
	}
}

func DeviceCertsModelFromMap(m map[string]interface{}) *models.DeviceCerts {
	pemCert := m["pem_cert"].(strfmt.Base64)
	pemKey := m["pem_key"].(strfmt.Base64)
	return &models.DeviceCerts{
		PemCert: pemCert,
		PemKey:  pemKey,
	}
}

// Update the underlying DeviceCerts resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDeviceCertsResourceData(d *schema.ResourceData, m *models.DeviceCerts) {
	d.Set("pem_cert", m.PemCert)
	d.Set("pem_key", m.PemKey)
}

// Iterate through and update the DeviceCerts resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDeviceCertsSubResourceData(m []*models.DeviceCerts) (d []*map[string]interface{}) {
	for _, DeviceCertsModel := range m {
		if DeviceCertsModel != nil {
			properties := make(map[string]interface{})
			properties["pem_cert"] = DeviceCertsModel.PemCert
			properties["pem_key"] = DeviceCertsModel.PemKey
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the DeviceCerts resource defined in the Terraform configuration
func DeviceCertsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"pem_cert": {
			Description:  `pem certificate`,
			Type:         schema.TypeString,
			ValidateFunc: validation.IsRFC3339Time,
			Optional:     true,
		},

		"pem_key": {
			Description:  `pem key`,
			Type:         schema.TypeString,
			ValidateFunc: validation.IsRFC3339Time,
			Optional:     true,
		},
	}
}

// Retrieve property field names for updating the DeviceCerts resource
func GetDeviceCertsPropertyFields() (t []string) {
	return []string{
		"pem_cert",
		"pem_key",
	}
}
