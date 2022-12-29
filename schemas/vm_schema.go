package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate VM resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func VMModel(d *schema.ResourceData) *models.VM {
	cPUPinningEnabled, _ := d.Get("cpu_pinning_enabled").(bool)
	cpusInt, _ := d.Get("cpus").(int)
	cpus := int64(cpusInt)
	memoryInt, _ := d.Get("memory").(int)
	memory := int64(memoryInt)
	modeModel, ok := d.Get("mode").(models.HvMode) // HvMode
	mode := &modeModel
	if !ok {
		mode = nil
	}
	vnc, _ := d.Get("vnc").(bool)
	return &models.VM{
		CPUPinningEnabled: cPUPinningEnabled,
		Cpus:              &cpus,   // int64 true false false
		Memory:            &memory, // int64 true false false
		Mode:              mode,
		Vnc:               &vnc, // bool true false false
	}
}

func VMModelFromMap(m map[string]interface{}) *models.VM {
	cPUPinningEnabled := m["cpu_pinning_enabled"].(bool)
	cpus := int64(m["cpus"].(int))     // int64 true false false
	memory := int64(m["memory"].(int)) // int64 true false false
	mode := m["mode"].(*models.HvMode) // HvMode
	vnc := m["vnc"].(bool)
	return &models.VM{
		CPUPinningEnabled: cPUPinningEnabled,
		Cpus:              &cpus,
		Memory:            &memory,
		Mode:              mode,
		Vnc:               &vnc,
	}
}

// Update the underlying VM resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetVMResourceData(d *schema.ResourceData, m *models.VM) {
	d.Set("cpu_pinning_enabled", m.CPUPinningEnabled)
	d.Set("cpus", m.Cpus)
	d.Set("memory", m.Memory)
	d.Set("mode", m.Mode)
	d.Set("vnc", m.Vnc)
	d.Set("vnc_display", m.VncDisplay)
}

// Iterate through and update the VM resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetVMSubResourceData(m []*models.VM) (d []*map[string]interface{}) {
	for _, VMModel := range m {
		if VMModel != nil {
			properties := make(map[string]interface{})
			properties["cpu_pinning_enabled"] = VMModel.CPUPinningEnabled
			properties["cpus"] = VMModel.Cpus
			properties["memory"] = VMModel.Memory
			properties["mode"] = VMModel.Mode
			properties["vnc"] = VMModel.Vnc
			properties["vnc_display"] = VMModel.VncDisplay
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the VM resource defined in the Terraform configuration
func VMSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cpu_pinning_enabled": {
			Description: `Enable CpuPinning`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"cpus": {
			Description: `CPUs`,
			Type:        schema.TypeInt,
			Required:    true,
		},

		"memory": {
			Description: `Memory`,
			Type:        schema.TypeInt,
			Required:    true,
		},

		"mode": {
			Description: `Hardware Virtualization`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"vnc": {
			Description: `VNC`,
			Type:        schema.TypeBool,
			Required:    true,
		},

		"vnc_display": {
			Description: `VNC display`,
			Type:        schema.TypeInt,
			Computed:    true,
		},
	}
}

// Retrieve property field names for updating the VM resource
func GetVMPropertyFields() (t []string) {
	return []string{
		"cpu_pinning_enabled",
		"cpus",
		"memory",
		"mode",
		"vnc",
	}
}
