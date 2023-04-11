package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func SANValuesModel(d *schema.ResourceData) *models.SANValues {
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
	var emaildIds []string
	emaildIdsInterface, emaildIdsIsSet := d.GetOk("emaildIds")
	if emaildIdsIsSet {
		var items []interface{}
		if listItems, isList := emaildIdsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = emaildIdsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			emaildIds = append(emaildIds, v.(string))
		}
	}
	var hosts []string
	hostsInterface, hostsIsSet := d.GetOk("hosts")
	if hostsIsSet {
		var items []interface{}
		if listItems, isList := hostsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = hostsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			hosts = append(hosts, v.(string))
		}
	}
	var ips []string
	ipsInterface, ipsIsSet := d.GetOk("ips")
	if ipsIsSet {
		var items []interface{}
		if listItems, isList := ipsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = ipsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			ips = append(ips, v.(string))
		}
	}
	var upns []string
	upnsInterface, upnsIsSet := d.GetOk("upns")
	if upnsIsSet {
		var items []interface{}
		if listItems, isList := upnsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = upnsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			upns = append(upns, v.(string))
		}
	}
	var uris []string
	urisInterface, urisIsSet := d.GetOk("uris")
	if urisIsSet {
		var items []interface{}
		if listItems, isList := urisInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = urisInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			uris = append(uris, v.(string))
		}
	}
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
	var emaildIds []string
	emaildIdsInterface, emaildIdsIsSet := m["emaildIds"]
	if emaildIdsIsSet {
		var items []interface{}
		if listItems, isList := emaildIdsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = emaildIdsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			emaildIds = append(emaildIds, v.(string))
		}
	}
	var hosts []string
	hostsInterface, hostsIsSet := m["hosts"]
	if hostsIsSet {
		var items []interface{}
		if listItems, isList := hostsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = hostsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			hosts = append(hosts, v.(string))
		}
	}
	var ips []string
	ipsInterface, ipsIsSet := m["ips"]
	if ipsIsSet {
		var items []interface{}
		if listItems, isList := ipsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = ipsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			ips = append(ips, v.(string))
		}
	}
	var upns []string
	upnsInterface, upnsIsSet := m["upns"]
	if upnsIsSet {
		var items []interface{}
		if listItems, isList := upnsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = upnsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			upns = append(upns, v.(string))
		}
	}
	var uris []string
	urisInterface, urisIsSet := m["uris"]
	if urisIsSet {
		var items []interface{}
		if listItems, isList := urisInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = urisInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			uris = append(uris, v.(string))
		}
	}
	return &models.SANValues{
		DNS:       dns,
		EmaildIds: emaildIds,
		Hosts:     hosts,
		Ips:       ips,
		Upns:      upns,
		Uris:      uris,
	}
}

func SetSANValuesResourceData(d *schema.ResourceData, m *models.SANValues) {
	d.Set("dns", m.DNS)
	d.Set("emaild_ids", m.EmaildIds)
	d.Set("hosts", m.Hosts)
	d.Set("ips", m.Ips)
	d.Set("upns", m.Upns)
	d.Set("uris", m.Uris)
}

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

func SANValues() map[string]*schema.Schema {
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
