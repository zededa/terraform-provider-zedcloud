package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func CertificateChainModel(d *schema.ResourceData) *models.CertificateChain {
	var certificates []*models.Certificate // []*Certificate
	certificatesInterface, certificatesIsSet := d.GetOk("certificates")
	if certificatesIsSet {
		var items []interface{}
		if listItems, isList := certificatesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = certificatesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := CertificateModelFromMap(v.(map[string]interface{}))
			certificates = append(certificates, m)
		}
	}
	return &models.CertificateChain{
		Certificates: certificates,
	}
}

func CertificateChainModelFromMap(m map[string]interface{}) *models.CertificateChain {
	var certificates []*models.Certificate // []*Certificate
	certificatesInterface, certificatesIsSet := m["certificates"]
	if certificatesIsSet {
		var items []interface{}
		if listItems, isList := certificatesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = certificatesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := CertificateModelFromMap(v.(map[string]interface{}))
			certificates = append(certificates, m)
		}
	}
	return &models.CertificateChain{
		Certificates: certificates,
	}
}

// Update the underlying CertificateChain resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetCertificateChainResourceData(d *schema.ResourceData, m *models.CertificateChain) {
	d.Set("certificates", SetCertificateSubResourceData(m.Certificates))
}

// Iterate through and update the CertificateChain resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetCertificateChainSubResourceData(m []*models.CertificateChain) (d []*map[string]interface{}) {
	for _, CertificateChainModel := range m {
		if CertificateChainModel != nil {
			properties := make(map[string]interface{})
			properties["certificates"] = SetCertificateSubResourceData(CertificateChainModel.Certificates)
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the CertificateChain resource defined in the Terraform configuration
func CertificateChainSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"certificates": {
			Description: `List of Certificate object holds the details of certificate like cert block, encryption type, validity, subject etc`,
			Type:        schema.TypeList, //GoType: []*Certificate
			Elem: &schema.Resource{
				Schema: Certificate(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
	}
}

// Retrieve property field names for updating the CertificateChain resource
func GetCertificateChainPropertyFields() (t []string) {
	return []string{
		"certificates",
	}
}
