package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func DhcpServerConfigModel(d *schema.ResourceData) *models.DhcpServerConfig {
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
		var items []interface{}
		if listItems, isList := dnsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = dnsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			dns = append(dns, v.(string))
		}
	}
	domain, _ := d.Get("domain").(string)
	gateway, _ := d.Get("gateway").(string)
	mask, _ := d.Get("mask").(string)
	ntp, _ := d.Get("ntp").(string)
	subnet, _ := d.Get("subnet").(string)
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
		var items []interface{}
		if listItems, isList := dnsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = dnsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			dns = append(dns, v.(string))
		}
	}
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

func SetDhcpServerConfigResourceData(d *schema.ResourceData, m *models.DhcpServerConfig) {
	d.Set("dhcp_range", SetDhcpIPRangeSubResourceData([]*models.DhcpIPRange{m.DhcpRange}))
	d.Set("dns", m.DNS)
	d.Set("domain", m.Domain)
	d.Set("gateway", m.Gateway)
	d.Set("mask", m.Mask)
	d.Set("ntp", m.Ntp)
	d.Set("subnet", m.Subnet)
}

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

func DhcpServerConfig() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"dhcp_range": {
			Description: `Range of IP addresses to be used for DHCP`,
			Type:        schema.TypeList, //GoType: DhcpIPRange
			Elem: &schema.Resource{
				Schema: DhcpIPRange(),
			},
			Optional: true,
		},

		"dns": {
			Description: `IP Addresses of DNS servers`,
			Type:        schema.TypeList, //GoType: []string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"domain": {
			Description: `Network domain`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"gateway": {
			Description: `IP Address of Network Gateway`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"mask": {
			Description: `Subnet Mask`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"ntp": {
			Description: `IP Address of NTP Server`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"subnet": {
			Description: `Subnet address`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

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
