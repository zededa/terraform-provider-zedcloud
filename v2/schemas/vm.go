package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func VMModel(d *schema.ResourceData) *models.VM {
	cPUPinningEnabled, _ := d.Get("cpu_pinning_enabled").(bool)
	cpusInt, _ := d.Get("cpus").(int)
	cpus := int64(cpusInt)
	memoryInt, _ := d.Get("memory").(int)
	memory := int64(memoryInt)
	var mode *models.HvMode // HvMode
	modeInterface, modeIsSet := d.GetOk("mode")
	if modeIsSet {
		modeModel := modeInterface.(string)
		mode = models.NewHvMode(models.HvMode(modeModel))
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
	var mode *models.HvMode            // HvMode
	modeInterface, modeIsSet := m["mode"]
	if modeIsSet {
		modeModel := modeInterface.(string)
		mode = models.NewHvMode(models.HvMode(modeModel))
	}
	vnc := m["vnc"].(bool)
	return &models.VM{
		CPUPinningEnabled: cPUPinningEnabled,
		Cpus:              &cpus,
		Memory:            &memory,
		Mode:              mode,
		Vnc:               &vnc,
	}
}

func SetVMResourceData(d *schema.ResourceData, m *models.VM) {
	d.Set("cpu_pinning_enabled", m.CPUPinningEnabled)
	d.Set("cpus", m.Cpus)
	d.Set("memory", m.Memory)
	d.Set("mode", m.Mode)
	d.Set("vnc", m.Vnc)
	d.Set("vnc_display", m.VncDisplay)
}

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

func VMSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cpu_pinning_enabled": {
			Description: `Enable CpuPinning`,
			Type:        schema.TypeBool,
			Computed:    true,
		},

		"cpus": {
			Description: `CPUs`,
			Type:        schema.TypeInt,
			Required:    true,
		},

		"memory": {
			Description: `Memory`,
			Type:        schema.TypeInt,
			Computed:    true,
		},

		"mode": {
			Description: `Hardware Virtualization`,
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
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
