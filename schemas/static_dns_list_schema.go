package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate StaticDNSList resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func StaticDNSListModel(d *schema.ResourceData) *models.StaticDNSList {
	addrs := d.Get("addrs").([]string)
	hostname := d.Get("hostname").(string)
	return &models.StaticDNSList{
		Addrs:    addrs,
		Hostname: hostname,
	}
}

func StaticDNSListModelFromMap(m map[string]interface{}) *models.StaticDNSList {
	addrs := m["addrs"].([]string)
	hostname := m["hostname"].(string)
	return &models.StaticDNSList{
		Addrs:    addrs,
		Hostname: hostname,
	}
}

// Update the underlying StaticDNSList resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetStaticDNSListResourceData(d *schema.ResourceData, m *models.StaticDNSList) {
	d.Set("addrs", m.Addrs)
	d.Set("hostname", m.Hostname)
}

// Iterate throught and update the StaticDNSList resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetStaticDNSListSubResourceData(m []*models.StaticDNSList) (d []*map[string]interface{}) {
	for _, StaticDNSListModel := range m {
		if StaticDNSListModel != nil {
			properties := make(map[string]interface{})
			properties["addrs"] = StaticDNSListModel.Addrs
			properties["hostname"] = StaticDNSListModel.Hostname
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the StaticDNSList resource defined in the Terraform configuration
func StaticDNSListSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"addrs": {
			Description: `Addresses`,
			Type:        schema.TypeList, //GoType: []string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"hostname": {
			Description: `Host name`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the StaticDNSList resource
func GetStaticDNSListPropertyFields() (t []string) {
	return []string{
		"addrs",
		"hostname",
	}
}
