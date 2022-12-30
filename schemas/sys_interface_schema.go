package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate SysInterface resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func SysInterfaceModel(d *schema.ResourceData) *models.SysInterface {
	costInt, _ := d.Get("cost").(int)
	cost := int64(costInt)
	var intfUsage *models.AdapterUsage // AdapterUsage
	intfUsageInterface, intfUsageIsSet := d.GetOk("intf_usage")
	if intfUsageIsSet {
		intfUsageModel := intfUsageInterface.(models.AdapterUsage)
		intfUsage = &intfUsageModel
	}
	intfname, _ := d.Get("intfname").(string)
	ipaddr, _ := d.Get("ipaddr").(string)
	macaddr, _ := d.Get("macaddr").(string)
	netname, _ := d.Get("netname").(string)
	tags, _ := d.Get("tags").(map[string]string) // map[string]string
	return &models.SysInterface{
		Cost:      cost,
		IntfUsage: intfUsage,
		Intfname:  intfname,
		Ipaddr:    ipaddr,
		Macaddr:   macaddr,
		Netname:   netname,
		Tags:      tags,
	}
}

func SysInterfaceModelFromMap(m map[string]interface{}) *models.SysInterface {
	cost := int64(m["cost"].(int))                      // int64 false false false
	intfUsage := m["intf_usage"].(*models.AdapterUsage) // AdapterUsage
	intfname := m["intfname"].(string)
	ipaddr := m["ipaddr"].(string)
	macaddr := m["macaddr"].(string)
	netname := m["netname"].(string)
	tags := m["tags"].(map[string]string)
	return &models.SysInterface{
		Cost:      cost,
		IntfUsage: intfUsage,
		Intfname:  intfname,
		Ipaddr:    ipaddr,
		Macaddr:   macaddr,
		Netname:   netname,
		Tags:      tags,
	}
}

// Update the underlying SysInterface resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetSysInterfaceResourceData(d *schema.ResourceData, m *models.SysInterface) {
	d.Set("cost", m.Cost)
	d.Set("intf_usage", m.IntfUsage)
	d.Set("intfname", m.Intfname)
	d.Set("ipaddr", m.Ipaddr)
	d.Set("macaddr", m.Macaddr)
	d.Set("netname", m.Netname)
	d.Set("tags", m.Tags)
}

// Iterate through and update the SysInterface resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetSysInterfaceSubResourceData(m []*models.SysInterface) (d []*map[string]interface{}) {
	for _, SysInterfaceModel := range m {
		if SysInterfaceModel != nil {
			properties := make(map[string]interface{})
			properties["cost"] = SysInterfaceModel.Cost
			properties["intf_usage"] = SysInterfaceModel.IntfUsage
			properties["intfname"] = SysInterfaceModel.Intfname
			properties["ipaddr"] = SysInterfaceModel.Ipaddr
			properties["macaddr"] = SysInterfaceModel.Macaddr
			properties["netname"] = SysInterfaceModel.Netname
			properties["tags"] = SysInterfaceModel.Tags
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the SysInterface resource defined in the Terraform configuration
func SysInterfaceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cost": {
			Description: `cost of using this interface. Default is 0.`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"intf_usage": {
			Description: `Adapter Udage`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"intfname": {
			Description: `name of interface in the manifest to which this network or adapter maps to`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"ipaddr": {
			Description: `IP address: we will be needing this in cae of static network`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"macaddr": {
			Description: `mac address needs to be over-written in some cases`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"netname": {
			Description: `network name: if attaching a network use netname`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"tags": {
			Description: `Tags are name/value pairs that enable you to categorize resources. Tag names are case insensitive with max_length 512 and min_length 3. Tag values are case sensitive with max_length 256 and min_length 3.`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},
	}
}

// Retrieve property field names for updating the SysInterface resource
func GetSysInterfacePropertyFields() (t []string) {
	return []string{
		"cost",
		"intf_usage",
		"intfname",
		"ipaddr",
		"macaddr",
		"netname",
		"tags",
	}
}
