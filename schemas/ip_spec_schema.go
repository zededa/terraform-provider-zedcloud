package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func IPSpecModel(d *schema.ResourceData) *models.IPSpec {
	deprecatedDhcp, _ := d.Get("deprecated_dhcp").(bool)
	var dhcp *models.NetworkDHCPType // NetworkDHCPType
	dhcpInterface, dhcpIsSet := d.GetOk("dhcp")
	if dhcpIsSet {
		dhcpModel := dhcpInterface.(string)
		dhcp = models.NewNetworkDHCPType(models.NetworkDHCPType(dhcpModel))
	}
	var dhcpRange *models.DhcpIPRange // DhcpIPRange
	dhcpRangeInterface, dhcpRangeIsSet := d.GetOk("dhcp_range")
	if dhcpRangeIsSet && dhcpRangeInterface != nil {
		dhcpRangeMap := dhcpRangeInterface.([]interface{})
		if len(dhcpRangeMap) > 0 {
			dhcpRange = DhcpIPRangeModelFromMap(dhcpRangeMap[0].(map[string]interface{}))
		}
	}
	var dns []string
	dnsInterface, dnsIsSet := d.GetOk("dns")
	if dnsIsSet {
		dnsSlice := dnsInterface.([]interface{})
		for _, i := range dnsSlice {
			dnsSlice = append(dnsSlice, i.(string))
		}
	}
	domain, _ := d.Get("domain").(string)
	gateway, _ := d.Get("gateway").(string)
	mask, _ := d.Get("mask").(string)
	ntp, _ := d.Get("ntp").(string)
	subnet, _ := d.Get("subnet").(string)
	return &models.IPSpec{
		DeprecatedDhcp: deprecatedDhcp,
		Dhcp:           dhcp,
		DhcpRange:      dhcpRange,
		DNS:            dns,
		Domain:         domain,
		Gateway:        gateway,
		Mask:           mask,
		Ntp:            ntp,
		Subnet:         subnet,
	}
}

func IPSpecModelFromMap(m map[string]interface{}) *models.IPSpec {
	deprecatedDhcp := m["deprecated_dhcp"].(bool)
	var dhcp *models.NetworkDHCPType // NetworkDHCPType
	dhcpInterface, dhcpIsSet := m["dhcp"]
	if dhcpIsSet {
		dhcpModel := dhcpInterface.(string)
		dhcp = models.NewNetworkDHCPType(models.NetworkDHCPType(dhcpModel))
	}
	var dhcpRange *models.DhcpIPRange // DhcpIPRange
	dhcpRangeInterface, dhcpRangeIsSet := m["dhcp_range"]
	if dhcpRangeIsSet && dhcpRangeInterface != nil {
		dhcpRangeMap := dhcpRangeInterface.([]interface{})
		if len(dhcpRangeMap) > 0 {
			dhcpRange = DhcpIPRangeModelFromMap(dhcpRangeMap[0].(map[string]interface{}))
		}
	}
	//
	var dns []string
	dnsInterface, dnsIsSet := m["dns"]
	if dnsIsSet {
		dnsSlice := dnsInterface.([]interface{})
		for _, i := range dnsSlice {
			dnsSlice = append(dnsSlice, i.(string))
		}
	}
	domain := m["domain"].(string)
	gateway := m["gateway"].(string)
	mask := m["mask"].(string)
	ntp := m["ntp"].(string)
	subnet := m["subnet"].(string)
	return &models.IPSpec{
		DeprecatedDhcp: deprecatedDhcp,
		Dhcp:           dhcp,
		DhcpRange:      dhcpRange,
		DNS:            dns,
		Domain:         domain,
		Gateway:        gateway,
		Mask:           mask,
		Ntp:            ntp,
		Subnet:         subnet,
	}
}

// Update the underlying IPSpec resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetIPSpecResourceData(d *schema.ResourceData, m *models.IPSpec) {
	d.Set("deprecated_dhcp", m.DeprecatedDhcp)
	d.Set("dhcp", m.Dhcp)
	d.Set("dhcp_range", SetDhcpIPRangeSubResourceData([]*models.DhcpIPRange{m.DhcpRange}))
	d.Set("dns", m.DNS)
	d.Set("domain", m.Domain)
	d.Set("gateway", m.Gateway)
	d.Set("mask", m.Mask)
	d.Set("ntp", m.Ntp)
	d.Set("subnet", m.Subnet)
}

// Iterate through and update the IPSpec resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetIPSpecSubResourceData(m []*models.IPSpec) (d []*map[string]interface{}) {
	for _, IPSpecModel := range m {
		if IPSpecModel != nil {
			properties := make(map[string]interface{})
			properties["deprecated_dhcp"] = IPSpecModel.DeprecatedDhcp
			properties["dhcp"] = IPSpecModel.Dhcp
			properties["dhcp_range"] = SetDhcpIPRangeSubResourceData([]*models.DhcpIPRange{IPSpecModel.DhcpRange})
			properties["dns"] = IPSpecModel.DNS
			properties["domain"] = IPSpecModel.Domain
			properties["gateway"] = IPSpecModel.Gateway
			properties["mask"] = IPSpecModel.Mask
			properties["ntp"] = IPSpecModel.Ntp
			properties["subnet"] = IPSpecModel.Subnet
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the IPSpec resource defined in the Terraform configuration
func IPSpecSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"deprecated_dhcp": {
			Description: ``,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"dhcp": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"dhcp_range": {
			Description: `for IPAM management when dhcp is turned on.
If none provided, system will default pool.`,
			Type: schema.TypeList, //GoType: DhcpIPRange
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

// Retrieve property field names for updating the IPSpec resource
func GetIPSpecPropertyFields() (t []string) {
	return []string{
		"deprecated_dhcp",
		"dhcp",
		"dhcp_range",
		"dns",
		"domain",
		"gateway",
		"mask",
		"ntp",
		"subnet",
	}
}
