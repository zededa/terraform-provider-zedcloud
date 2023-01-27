package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func ToStaticDNSListModel(d *schema.ResourceData) *models.StaticDNSList {
	var addrs []string
	addrsInterface, addrsIsSet := d.GetOk("addrs")
	if addrsIsSet {
		var items []interface{}
		if listItems, isList := addrsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = addrsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			addrs = append(addrs, v.(string))
		}
	}
	hostname, _ := d.Get("hostname").(string)
	return &models.StaticDNSList{
		Addrs:    addrs,
		Hostname: hostname,
	}
}

func StaticDNSListModelFromMap(m map[string]interface{}) *models.StaticDNSList {
	var addrs []string
	addrsInterface, addrsIsSet := m["addrs"]
	if addrsIsSet {
		var items []interface{}
		if listItems, isList := addrsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = addrsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			addrs = append(addrs, v.(string))
		}
	}
	hostname := m["hostname"].(string)
	return &models.StaticDNSList{
		Addrs:    addrs,
		Hostname: hostname,
	}
}

func SetStaticDNSListResourceData(d *schema.ResourceData, m *models.StaticDNSList) {
	d.Set("addrs", m.Addrs)
	d.Set("hostname", m.Hostname)
}

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
func StaticDNSList() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"addrs": {
			Description: "Set of IP addresses for the specified hostname",
			Type:        schema.TypeSet, //GoType: []string
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

func StaticDNSListPropertyFields() (t []string) {
	return []string{
		"addrs",
		"hostname",
	}
}
