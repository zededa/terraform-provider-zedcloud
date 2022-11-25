package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate DhcpServerConfig resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func DhcpServerConfigModel(d *schema.ResourceData) *models.DhcpServerConfig {
	var dhcpRange *models.DhcpIPRange // DhcpIPRange
	dhcpRangeInterface, dhcpRangeIsSet := d.GetOk("dhcp_range")
	if dhcpRangeIsSet {
		dhcpRangeMap := dhcpRangeInterface.([]interface{})[0].(map[string]interface{})
		dhcpRange = DhcpIPRangeModelFromMap(dhcpRangeMap)
	}
	dns := d.Get("dns").([]string)
	domain := d.Get("domain").(string)
	gateway := d.Get("gateway").(string)
	mask := d.Get("mask").(string)
	ntp := d.Get("ntp").(string)
	subnet := d.Get("subnet").(string)
	return &models.DhcpServerConfig{
		DhcpRange: dhcpRange,
		DNS:       dns,
		Domain:    domain,
		Gateway:   gateway,
		Mask:      mask,
		Ntp:       ntp,
		Subnet:    subnet,
	}
}

func DhcpServerConfigModelFromMap(m map[string]interface{}) *models.DhcpServerConfig {
	var dhcpRange *models.DhcpIPRange // DhcpIPRange
	dhcpRangeInterface, dhcpRangeIsSet := m["dhcp_range"]
	if dhcpRangeIsSet {
		dhcpRangeMap := dhcpRangeInterface.([]interface{})[0].(map[string]interface{})
		dhcpRange = DhcpIPRangeModelFromMap(dhcpRangeMap)
	}
	//
	dns := m["dns"].([]string)
	domain := m["domain"].(string)
	gateway := m["gateway"].(string)
	mask := m["mask"].(string)
	ntp := m["ntp"].(string)
	subnet := m["subnet"].(string)
	return &models.DhcpServerConfig{
		DhcpRange: dhcpRange,
		DNS:       dns,
		Domain:    domain,
		Gateway:   gateway,
		Mask:      mask,
		Ntp:       ntp,
		Subnet:    subnet,
	}
}

// Update the underlying DhcpServerConfig resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDhcpServerConfigResourceData(d *schema.ResourceData, m *models.DhcpServerConfig) {
	d.Set("dhcp_range", SetDhcpIPRangeSubResourceData([]*models.DhcpIPRange{m.DhcpRange}))
	d.Set("dns", m.DNS)
	d.Set("domain", m.Domain)
	d.Set("gateway", m.Gateway)
	d.Set("mask", m.Mask)
	d.Set("ntp", m.Ntp)
	d.Set("subnet", m.Subnet)
}

// Iterate throught and update the DhcpServerConfig resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDhcpServerConfigSubResourceData(m []*models.DhcpServerConfig) (d []*map[string]interface{}) {
	for _, DhcpServerConfigModel := range m {
		if DhcpServerConfigModel != nil {
			properties := make(map[string]interface{})
			properties["dhcp_range"] = SetDhcpIPRangeSubResourceData([]*models.DhcpIPRange{DhcpServerConfigModel.DhcpRange})
			properties["dns"] = DhcpServerConfigModel.DNS
			properties["domain"] = DhcpServerConfigModel.Domain
			properties["gateway"] = DhcpServerConfigModel.Gateway
			properties["mask"] = DhcpServerConfigModel.Mask
			properties["ntp"] = DhcpServerConfigModel.Ntp
			properties["subnet"] = DhcpServerConfigModel.Subnet
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the DhcpServerConfig resource defined in the Terraform configuration
func DhcpServerConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"dhcp_range": {
			Description: `for IPAM management when dhcp is turned on.If none provided, system will default pool.`,
			Type:        schema.TypeList, //GoType: DhcpIPRange
			Elem: &schema.Resource{
				Schema: DhcpIPRangeSchema(),
			},
			Optional: true,
		},

		"dns": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"domain": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"gateway": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"mask": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"ntp": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"subnet": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the DhcpServerConfig resource
func GetDhcpServerConfigPropertyFields() (t []string) {
	return []string{
		"dhcp_range",
		"dns",
		"domain",
		"gateway",
		"mask",
		"ntp",
		"subnet",
	}
}
