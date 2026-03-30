package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ArpMonitorModel(d *schema.ResourceData) *models.ArpMonitor {
	intervalInt, _ := d.Get("interval").(int)
	interval := int64(intervalInt)
	var ipTargets []string
	ipTargetsInterface, ipTargetsIsSet := d.GetOk("ip_targets")
	if ipTargetsIsSet {
		var items []interface{}
		if listItems, isList := ipTargetsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = ipTargetsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			ipTargets = append(ipTargets, v.(string))
		}
	}
	return &models.ArpMonitor{
		Interval:  interval,
		IPTargets: ipTargets,
	}
}

func ArpMonitorModelFromMap(m map[string]interface{}) *models.ArpMonitor {
	interval := int64(m["interval"].(int)) // int64
	var ipTargets []string
	ipTargetsInterface, ipTargetsIsSet := m["ip_targets"]
	if ipTargetsIsSet {
		var items []interface{}
		if listItems, isList := ipTargetsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = ipTargetsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			ipTargets = append(ipTargets, v.(string))
		}
	}
	return &models.ArpMonitor{
		Interval:  interval,
		IPTargets: ipTargets,
	}
}

func SetArpMonitorResourceData(d *schema.ResourceData, m *models.ArpMonitor) {
	d.Set("interval", m.Interval)
	d.Set("ip_targets", m.IPTargets)
}

func SetArpMonitorSubResourceData(m []*models.ArpMonitor) (d []*map[string]interface{}) {
	for _, ArpMonitorModel := range m {
		if ArpMonitorModel != nil {
			properties := make(map[string]interface{})
			properties["interval"] = ArpMonitorModel.Interval
			properties["ip_targets"] = ArpMonitorModel.IPTargets
			d = append(d, &properties)
		}
	}
	return
}

func ArpMonitorSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"interval": {
			Description: `Interval specifies the ARP link monitoring frequency in milliseconds.`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"ip_targets": {
			Description: `IpTargets specifies the IPv4 addresses to use as ARP monitoring peers.
These are the targets of ARP requests sent to determine the health of links.`,
			Type: schema.TypeList, //GoType: []string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},
	}
}

func GetArpMonitorPropertyFields() (t []string) {
	return []string{
		"interval",
		"ip_targets",
	}
}
