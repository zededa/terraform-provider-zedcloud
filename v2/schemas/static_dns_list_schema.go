package schemas

import (
	"slices"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func StaticDNSListModel(d *schema.ResourceData) *models.StaticDNSList {
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

func SetStaticDNSListSubResourceData(m []*models.StaticDNSList) []*map[string]interface{} {
	result := []*map[string]interface{}{}
	for _, StaticDNSListModel := range m {
		if StaticDNSListModel != nil {
			properties := make(map[string]interface{})
			properties["addrs"] = StaticDNSListModel.Addrs
			properties["hostname"] = StaticDNSListModel.Hostname
			result = append(result, &properties)
		}
	}
	return result
}

func StaticDNSList() map[string]*schema.Schema {
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

func GetStaticDNSListPropertyFields() (t []string) {
	return []string{
		"addrs",
		"hostname",
	}
}

func CompareDNSLists(x, y []*models.StaticDNSList) bool {
	if len(x) != len(y) {
		return false
	}

	// Sort both slices by Name (which is required)
	slices.SortFunc(x, func(i, j *models.StaticDNSList) int {
		return strings.Compare(strings.ToLower(i.Hostname), strings.ToLower(j.Hostname))
	})

	slices.SortFunc(y, func(i, j *models.StaticDNSList) int {
		return strings.Compare(strings.ToLower(i.Hostname), strings.ToLower(j.Hostname))
	})

	// Compare each StaticDNSList
	equal := slices.EqualFunc(x, y, func(a, b *models.StaticDNSList) bool {
		if a.Hostname != b.Hostname {
			return false
		}
		slices.Sort(a.Addrs)
		slices.Sort(b.Addrs)
		if !Equal(a.Addrs, b.Addrs) {
			return false
		}
		return true
	})
	return equal
}
