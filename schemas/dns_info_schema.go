package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate DNSInfo resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func DNSInfoModel(d *schema.ResourceData) *models.DNSInfo {
	domain, _ := d.Get("domain").(string)
	search, _ := d.Get("search").([]string)
	servers, _ := d.Get("servers").([]string)
	return &models.DNSInfo{
		Domain:  &domain, // string true false false
		Search:  search,
		Servers: servers,
	}
}

func DNSInfoModelFromMap(m map[string]interface{}) *models.DNSInfo {
	domain := m["domain"].(string)
	search := m["search"].([]string)
	servers := m["servers"].([]string)
	return &models.DNSInfo{
		Domain:  &domain,
		Search:  search,
		Servers: servers,
	}
}

// Update the underlying DNSInfo resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDNSInfoResourceData(d *schema.ResourceData, m *models.DNSInfo) {
	d.Set("domain", m.Domain)
	d.Set("search", m.Search)
	d.Set("servers", m.Servers)
}

// Iterate through and update the DNSInfo resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDNSInfoSubResourceData(m []*models.DNSInfo) (d []*map[string]interface{}) {
	for _, DNSInfoModel := range m {
		if DNSInfoModel != nil {
			properties := make(map[string]interface{})
			properties["domain"] = DNSInfoModel.Domain
			properties["search"] = DNSInfoModel.Search
			properties["servers"] = DNSInfoModel.Servers
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the DNSInfo resource defined in the Terraform configuration
func DNSInfoSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"domain": {
			Description: `domain name`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"search": {
			Description: `Array of search strings`,
			Type:        schema.TypeList, //GoType: []string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Required: true,
		},

		"servers": {
			Description: `Array of dns server`,
			Type:        schema.TypeList, //GoType: []string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Required: true,
		},
	}
}

// Retrieve property field names for updating the DNSInfo resource
func GetDNSInfoPropertyFields() (t []string) {
	return []string{
		"domain",
		"search",
		"servers",
	}
}
