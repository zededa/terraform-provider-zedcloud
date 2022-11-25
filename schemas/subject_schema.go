package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate Subject resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func SubjectModel(d *schema.ResourceData) *models.Subject {
	commonName := d.Get("common_name").(string)
	country := d.Get("country").([]string)
	locality := d.Get("locality").([]string)
	organization := d.Get("organization").([]string)
	organizationalUnit := d.Get("organizational_unit").([]string)
	postalCode := d.Get("postal_code").([]string)
	province := d.Get("province").([]string)
	serialNumber := d.Get("serial_number").(string)
	return &models.Subject{
		CommonName:         commonName,
		Country:            country,
		Locality:           locality,
		Organization:       organization,
		OrganizationalUnit: organizationalUnit,
		PostalCode:         postalCode,
		Province:           province,
		SerialNumber:       serialNumber,
	}
}

func SubjectModelFromMap(m map[string]interface{}) *models.Subject {
	commonName := m["common_name"].(string)
	country := m["country"].([]string)
	locality := m["locality"].([]string)
	organization := m["organization"].([]string)
	organizationalUnit := m["organizational_unit"].([]string)
	postalCode := m["postal_code"].([]string)
	province := m["province"].([]string)
	serialNumber := m["serial_number"].(string)
	return &models.Subject{
		CommonName:         commonName,
		Country:            country,
		Locality:           locality,
		Organization:       organization,
		OrganizationalUnit: organizationalUnit,
		PostalCode:         postalCode,
		Province:           province,
		SerialNumber:       serialNumber,
	}
}

// Update the underlying Subject resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetSubjectResourceData(d *schema.ResourceData, m *models.Subject) {
	d.Set("common_name", m.CommonName)
	d.Set("country", m.Country)
	d.Set("locality", m.Locality)
	d.Set("organization", m.Organization)
	d.Set("organizational_unit", m.OrganizationalUnit)
	d.Set("postal_code", m.PostalCode)
	d.Set("province", m.Province)
	d.Set("serial_number", m.SerialNumber)
}

// Iterate throught and update the Subject resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetSubjectSubResourceData(m []*models.Subject) (d []*map[string]interface{}) {
	for _, SubjectModel := range m {
		if SubjectModel != nil {
			properties := make(map[string]interface{})
			properties["common_name"] = SubjectModel.CommonName
			properties["country"] = SubjectModel.Country
			properties["locality"] = SubjectModel.Locality
			properties["organization"] = SubjectModel.Organization
			properties["organizational_unit"] = SubjectModel.OrganizationalUnit
			properties["postal_code"] = SubjectModel.PostalCode
			properties["province"] = SubjectModel.Province
			properties["serial_number"] = SubjectModel.SerialNumber
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the Subject resource defined in the Terraform configuration
func SubjectSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"common_name": {
			Description: `Certificate common name.`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"country": {
			Description: `List of countries.`,
			Type:        schema.TypeList, //GoType: []string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"locality": {
			Description: `List of locallity.`,
			Type:        schema.TypeList, //GoType: []string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"organization": {
			Description: `List of organization.`,
			Type:        schema.TypeList, //GoType: []string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"organizational_unit": {
			Description: `List of Organizational Unit.`,
			Type:        schema.TypeList, //GoType: []string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"postal_code": {
			Description: `List of Postal codes.`,
			Type:        schema.TypeList, //GoType: []string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"province": {
			Description: `List of List of Prvince.`,
			Type:        schema.TypeList, //GoType: []string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"serial_number": {
			Description: `Subject cerial number`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the Subject resource
func GetSubjectPropertyFields() (t []string) {
	return []string{
		"common_name",
		"country",
		"locality",
		"organization",
		"organizational_unit",
		"postal_code",
		"province",
		"serial_number",
	}
}
