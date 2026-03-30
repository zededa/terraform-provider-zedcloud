package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func MIIMonitorModel(d *schema.ResourceData) *models.MIIMonitor {
	downdelayInt, _ := d.Get("downdelay").(int)
	downdelay := int64(downdelayInt)
	intervalInt, _ := d.Get("interval").(int)
	interval := int64(intervalInt)
	updelayInt, _ := d.Get("updelay").(int)
	updelay := int64(updelayInt)
	return &models.MIIMonitor{
		Downdelay: downdelay,
		Interval:  interval,
		Updelay:   updelay,
	}
}

func MIIMonitorModelFromMap(m map[string]interface{}) *models.MIIMonitor {
	downdelay := int64(m["downdelay"].(int)) // int64
	interval := int64(m["interval"].(int))   // int64
	updelay := int64(m["updelay"].(int))     // int64
	return &models.MIIMonitor{
		Downdelay: downdelay,
		Interval:  interval,
		Updelay:   updelay,
	}
}

func SetMIIMonitorResourceData(d *schema.ResourceData, m *models.MIIMonitor) {
	d.Set("downdelay", m.Downdelay)
	d.Set("interval", m.Interval)
	d.Set("updelay", m.Updelay)
}

func SetMIIMonitorSubResourceData(m []*models.MIIMonitor) (d []*map[string]interface{}) {
	for _, MIIMonitorModel := range m {
		if MIIMonitorModel != nil {
			properties := make(map[string]interface{})
			properties["downdelay"] = MIIMonitorModel.Downdelay
			properties["interval"] = MIIMonitorModel.Interval
			properties["updelay"] = MIIMonitorModel.Updelay
			d = append(d, &properties)
		}
	}
	return
}

func MIIMonitorSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"downdelay": {
			Description: `Specifies the time, in milliseconds, to wait before disabling a bond
slave after a link failure has been detected.
The downdelay value should be a multiple of the monitoring interval; if not,
it will be rounded down to the nearest multiple.
The default value is 0.`,
			Type:     schema.TypeInt,
			Optional: true,
		},

		"interval": {
			Description: `Specifies the MII link monitoring frequency in milliseconds.
This determines how often the link state of each bond slave is inspected
for link failures.`,
			Type:     schema.TypeInt,
			Optional: true,
		},

		"updelay": {
			Description: `Updelay specifies the time, in milliseconds, to wait before enabling
a bond slave after a link recovery has been detected.
The updelay value should be a multiple of the monitoring interval; if not,
it will be rounded down to the nearest multiple.
The default value is 0.`,
			Type:     schema.TypeInt,
			Optional: true,
		},
	}
}

func GetMIIMonitorPropertyFields() (t []string) {
	return []string{
		"downdelay",
		"interval",
		"updelay",
	}
}
