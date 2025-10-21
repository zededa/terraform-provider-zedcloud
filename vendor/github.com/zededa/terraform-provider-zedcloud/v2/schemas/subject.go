package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func SubjectModel(d *schema.ResourceData) *models.Subject {
	commonName, _ := d.Get("common_name").(string)
	var country []string
	countryInterface, countryIsSet := d.GetOk("country")
	if countryIsSet {
		var items []interface{}
		if listItems, isList := countryInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = countryInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			country = append(country, v.(string))
		}
	}
	var locality []string
	localityInterface, localityIsSet := d.GetOk("locality")
	if localityIsSet {
		var items []interface{}
		if listItems, isList := localityInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = localityInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			locality = append(locality, v.(string))
		}
	}
	var organization []string
	organizationInterface, organizationIsSet := d.GetOk("organization")
	if organizationIsSet {
		var items []interface{}
		if listItems, isList := organizationInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = organizationInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			organization = append(organization, v.(string))
		}
	}
	var organizationalUnit []string
	organizationalUnitInterface, organizationalUnitIsSet := d.GetOk("organizationalUnit")
	if organizationalUnitIsSet {
		var items []interface{}
		if listItems, isList := organizationalUnitInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = organizationalUnitInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			organizationalUnit = append(organizationalUnit, v.(string))
		}
	}
	var postalCode []string
	postalCodeInterface, postalCodeIsSet := d.GetOk("postalCode")
	if postalCodeIsSet {
		var items []interface{}
		if listItems, isList := postalCodeInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = postalCodeInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			postalCode = append(postalCode, v.(string))
		}
	}
	var province []string
	provinceInterface, provinceIsSet := d.GetOk("province")
	if provinceIsSet {
		var items []interface{}
		if listItems, isList := provinceInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = provinceInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			province = append(province, v.(string))
		}
	}
	serialNumber, _ := d.Get("serial_number").(string)
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
	var country []string
	countryInterface, countryIsSet := m["country"]
	if countryIsSet {
		var items []interface{}
		if listItems, isList := countryInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = countryInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			country = append(country, v.(string))
		}
	}
	var locality []string
	localityInterface, localityIsSet := m["locality"]
	if localityIsSet {
		var items []interface{}
		if listItems, isList := localityInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = localityInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			locality = append(locality, v.(string))
		}
	}
	var organization []string
	organizationInterface, organizationIsSet := m["organization"]
	if organizationIsSet {
		var items []interface{}
		if listItems, isList := organizationInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = organizationInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			organization = append(organization, v.(string))
		}
	}
	var organizationalUnit []string
	organizationalUnitInterface, organizationalUnitIsSet := m["organizationalUnit"]
	if organizationalUnitIsSet {
		var items []interface{}
		if listItems, isList := organizationalUnitInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = organizationalUnitInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			organizationalUnit = append(organizationalUnit, v.(string))
		}
	}
	var postalCode []string
	postalCodeInterface, postalCodeIsSet := m["postalCode"]
	if postalCodeIsSet {
		var items []interface{}
		if listItems, isList := postalCodeInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = postalCodeInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			postalCode = append(postalCode, v.(string))
		}
	}
	var province []string
	provinceInterface, provinceIsSet := m["province"]
	if provinceIsSet {
		var items []interface{}
		if listItems, isList := provinceInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = provinceInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			province = append(province, v.(string))
		}
	}
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

func Subject() map[string]*schema.Schema {
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
