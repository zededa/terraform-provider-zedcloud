package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func StaticDNSListModel(d *schema.ResourceData) *models.StaticDNSList {
	var addrs []string
	addrsInterface, addrsIsSet := d.GetOk("addrs")
	if addrsIsSet {
		addrsSlice := addrsInterface.([]interface{})
		for _, i := range addrsSlice {
			addrsSlice = append(addrsSlice, i.(string))
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
		addrsSlice := addrsInterface.([]interface{})
		for _, i := range addrsSlice {
			addrsSlice = append(addrsSlice, i.(string))
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
func StaticDNSListSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"addrs": {
			Description: "Set of IP addresses for the specified hostname",
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
