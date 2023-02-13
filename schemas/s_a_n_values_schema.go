package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate SANValues resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func SANValuesModel(d *schema.ResourceData) *models.SANValues {
	dns, _ := d.Get("dns").([]string)
	emaildIds, _ := d.Get("emaild_ids").([]string)
	hosts, _ := d.Get("hosts").([]string)
	ips, _ := d.Get("ips").([]string)
	upns, _ := d.Get("upns").([]string)
	uris, _ := d.Get("uris").([]string)
	return &models.SANValues{
		DNS:       dns,
		EmaildIds: emaildIds,
		Hosts:     hosts,
		Ips:       ips,
		Upns:      upns,
		Uris:      uris,
	}
}

func SANValuesModelFromMap(m map[string]interface{}) *models.SANValues {
	dns := m["dns"].([]string)
	emaildIds := m["emaild_ids"].([]string)
	hosts := m["hosts"].([]string)
	ips := m["ips"].([]string)
	upns := m["upns"].([]string)
	uris := m["uris"].([]string)
	return &models.SANValues{
		DNS:       dns,
		EmaildIds: emaildIds,
		Hosts:     hosts,
		Ips:       ips,
		Upns:      upns,
		Uris:      uris,
	}
}

// Update the underlying SANValues resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetSANValuesResourceData(d *schema.ResourceData, m *models.SANValues) {
	d.Set("dns", m.DNS)
	d.Set("emaild_ids", m.EmaildIds)
	d.Set("hosts", m.Hosts)
	d.Set("ips", m.Ips)
	d.Set("upns", m.Upns)
	d.Set("uris", m.Uris)
}

// Iterate through and update the SANValues resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetSANValuesSubResourceData(m []*models.SANValues) (d []*map[string]interface{}) {
	for _, SANValuesModel := range m {
		if SANValuesModel != nil {
			properties := make(map[string]interface{})
			properties["dns"] = SANValuesModel.DNS
			properties["emaild_ids"] = SANValuesModel.EmaildIds
			properties["hosts"] = SANValuesModel.Hosts
			properties["ips"] = SANValuesModel.Ips
			properties["upns"] = SANValuesModel.Upns
			properties["uris"] = SANValuesModel.Uris
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the SANValues resource defined in the Terraform configuration
func SANValuesSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"dns": {
			Description: `List of permitted DNS names.`,
			Type:        schema.TypeList, //GoType: []string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"emaild_ids": {
			Description: `List of permitted email addresses.`,
			Type:        schema.TypeList, //GoType: []string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"hosts": {
			Description: `List of permitted hosts.`,
			Type:        schema.TypeList, //GoType: []string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"ips": {
			Description: `List of permitted IP addresses.`,
			Type:        schema.TypeList, //GoType: []string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"upns": {
			Description: `List of permitted User principal names.`,
			Type:        schema.TypeList, //GoType: []string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"uris": {
			Description: `List of permitted URIs.`,
			Type:        schema.TypeList, //GoType: []string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},
	}
}

// Retrieve property field names for updating the SANValues resource
func GetSANValuesPropertyFields() (t []string) {
	return []string{
		"dns",
		"emaild_ids",
		"hosts",
		"ips",
		"upns",
		"uris",
	}
}
