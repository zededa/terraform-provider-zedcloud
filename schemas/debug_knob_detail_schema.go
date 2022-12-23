package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate DebugKnobDetail resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func DebugKnobDetailModel(d *schema.ResourceData) *models.DebugKnobDetail {
	debugKnob := d.Get("debug_knob").(bool)
	expired := d.Get("expired").(bool)
	expiry := d.Get("expiry").(string)
	id := d.Get("id").(string)
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

// Update the underlying DebugKnobDetail resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDebugKnobDetailResourceData(d *schema.ResourceData, m *models.DebugKnobDetail) {
	d.Set("debug_knob", m.DebugKnob)
	d.Set("expired", m.Expired)
	d.Set("expiry", m.Expiry)
	d.Set("id", m.ID)
}

// Iterate throught and update the DebugKnobDetail resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
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
func DebugKnobDetailSchema() map[string]*schema.Schema {
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
