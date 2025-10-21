package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func VMModel(d *schema.ResourceData) *models.VM {
	cPUPinningEnabled, _ := d.Get("cpu_pinning_enabled").(bool)
	cpusInt, _ := d.Get("cpus").(int)
	cpus := int64(cpusInt)
	disableVTPM, _ := d.Get("disable_v_t_p_m").(bool)
	memoryInt, _ := d.Get("memory").(int)
	memory := int64(memoryInt)
	var mode *models.HvMode // HvMode
	modeInterface, modeIsSet := d.GetOk("mode")
	if modeIsSet {
		modeModel := modeInterface.(string)
		mode = models.NewHvMode(models.HvMode(modeModel))
	}
	vnc, _ := d.Get("vnc").(bool)
	enableOemWinLicenseKey, _ := d.Get("enable_oem_win_license_key").(bool)
	return &models.VM{
		CPUPinningEnabled:      cPUPinningEnabled,
		Cpus:                   &cpus, // int64
		DisableVTPM:            disableVTPM,
		EnableOemWinLicenseKey: enableOemWinLicenseKey,
		Memory:                 &memory, // int64
		Mode:                   mode,
		Vnc:                    &vnc, // bool
	}
}

func VMModelFromMap(m map[string]interface{}) *models.VM {
	cPUPinningEnabled := m["cpu_pinning_enabled"].(bool)
	cpus := int64(m["cpus"].(int)) // int64
	disableVTPM := m["disable_v_t_p_m"].(bool)
	memory := int64(m["memory"].(int)) // int64
	var mode *models.HvMode            // HvMode
	modeInterface, modeIsSet := m["mode"]
	if modeIsSet {
		modeModel := modeInterface.(string)
		mode = models.NewHvMode(models.HvMode(modeModel))
	}
	vnc := m["vnc"].(bool)
	enableOemWinLicenseKey := m["enable_oem_win_license_key"].(bool)
	return &models.VM{
		CPUPinningEnabled:      cPUPinningEnabled,
		Cpus:                   &cpus,
		DisableVTPM:            disableVTPM,
		EnableOemWinLicenseKey: enableOemWinLicenseKey,
		Memory:                 &memory,
		Mode:                   mode,
		Vnc:                    &vnc,
	}
}

func SetVMResourceData(d *schema.ResourceData, m *models.VM) {
	d.Set("cpu_pinning_enabled", m.CPUPinningEnabled)
	d.Set("cpus", m.Cpus)
	d.Set("disable_v_t_p_m", m.DisableVTPM)
	d.Set("memory", m.Memory)
	d.Set("mode", m.Mode)
	d.Set("vnc", m.Vnc)
	d.Set("vnc_display", m.VncDisplay)
	d.Set("enable_oem_win_license_key", m.EnableOemWinLicenseKey)
}

func SetVMSubResourceData(m []*models.VM) (d []*map[string]interface{}) {
	for _, VMModel := range m {
		if VMModel != nil {
			properties := make(map[string]interface{})
			properties["cpu_pinning_enabled"] = VMModel.CPUPinningEnabled
			properties["cpus"] = VMModel.Cpus
			properties["disable_v_t_p_m"] = VMModel.DisableVTPM
			properties["memory"] = VMModel.Memory
			properties["mode"] = VMModel.Mode
			properties["vnc"] = VMModel.Vnc
			properties["vnc_display"] = VMModel.VncDisplay
			properties["enable_oem_win_license_key"] = VMModel.EnableOemWinLicenseKey
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
			Optional:    true,
		},

		"cpus": {
			Description: `CPUs`,
			Type:        schema.TypeInt,
			Required:    true,
		},

		"disable_v_t_p_m": {
			Description: `Disable vTPM for virtual machines (VM)`,
			Type:        schema.TypeBool,
			Optional:    true,
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

		"enable_oem_win_license_key": {
			Description: `Enable device which has VM to receive the Windows license embedded in the ACPI tables`,
			Type:        schema.TypeBool,
			Optional:    true,
		},
	}
}

func GetVMPropertyFields() (t []string) {
	return []string{
		"cpu_pinning_enabled",
		"cpus",
		"disable_v_t_p_m",
		"memory",
		"mode",
		"vnc",
		"enable_oem_win_license_key",
	}
}
