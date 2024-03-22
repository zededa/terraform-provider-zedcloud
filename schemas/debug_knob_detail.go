package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

func DebugKnobDetailModel(d *schema.ResourceData) *models.DebugKnobDetail {
	debugKnob, _ := d.Get("debug_knob").(bool)
	expired, _ := d.Get("expired").(bool)
	expiry, _ := d.Get("expiry").(string)
	id, _ := d.Get("id").(string)
	return &models.DebugKnobDetail{
		DebugKnob: debugKnob,
		Expired:   expired,
		Expiry:    expiry,
		ID:        id,
	}
}

func DebugKnobDetailModelFromMap(m map[string]interface{}) *models.DebugKnobDetail {
	debugKnob := m["debug_knob"].(bool)
	expired := m["expired"].(bool)
	expiry := m["expiry"].(string)
	id := m["id"].(string)
	return &models.DebugKnobDetail{
		DebugKnob: debugKnob,
		Expired:   expired,
		Expiry:    expiry,
		ID:        id,
	}
}

func SetDebugKnobDetailResourceData(d *schema.ResourceData, m *models.DebugKnobDetail) {
	d.Set("debug_knob", m.DebugKnob)
	d.Set("expired", m.Expired)
	d.Set("expiry", m.Expiry)
	d.Set("id", m.ID)
}

func SetDebugKnobDetailSubResourceData(m []*models.DebugKnobDetail) (d []*map[string]interface{}) {
	for _, DebugKnobDetailModel := range m {
		if DebugKnobDetailModel != nil {
			properties := make(map[string]interface{})
			properties["debug_knob"] = DebugKnobDetailModel.DebugKnob
			properties["expired"] = DebugKnobDetailModel.Expired
			properties["expiry"] = DebugKnobDetailModel.Expiry
			properties["id"] = DebugKnobDetailModel.ID
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the DebugKnobDetail resource defined in the Terraform configuration
func DebugKnobDetail() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"debug_knob": {
			Description: `debug knob flag`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"expired": {
			Description: `debug knob expiry status flag`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"expiry": {
			Description: `debug expiry time in minutes`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"id": {
			Description: `system generated unique id for a device`,
			Type:        schema.TypeString,
			Computed:    true,
		},
	}
}

// Retrieve property field names for updating the DebugKnobDetail resource
func GetDebugKnobDetailPropertyFields() (t []string) {
	return []string{
		"debug_knob",
		"expired",
		"expiry",
		"id",
	}
}
