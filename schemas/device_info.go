package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

// Function to perform the following actions:
// (1) Translate DeviceInfo resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func DeviceInfoModel(d *schema.ResourceData) *models.DeviceInfo {
	cPUArch, _ := d.Get("cpu_arch").(string)
	machineArch, _ := d.Get("machine_arch").(string)
	memMB, _ := d.Get("mem_m_b").(string)
	nCPUInt, _ := d.Get("n_cpu").(int)
	nCPU := int64(nCPUInt)
	platform, _ := d.Get("platform").(string)
	storageMB, _ := d.Get("storage_m_b").(string)
	return &models.DeviceInfo{
		CPUArch:     cPUArch,
		MachineArch: machineArch,
		MemMB:       memMB,
		NCPU:        nCPU,
		Platform:    platform,
		StorageMB:   storageMB,
	}
}

func DeviceInfoModelFromMap(m map[string]interface{}) *models.DeviceInfo {
	cPUArch := m["cpu_arch"].(string)
	machineArch := m["machine_arch"].(string)
	memMB := m["mem_m_b"].(string)
	nCPU := int64(m["n_cpu"].(int)) // int64 false false false
	platform := m["platform"].(string)
	storageMB := m["storage_m_b"].(string)
	return &models.DeviceInfo{
		CPUArch:     cPUArch,
		MachineArch: machineArch,
		MemMB:       memMB,
		NCPU:        nCPU,
		Platform:    platform,
		StorageMB:   storageMB,
	}
}

// Update the underlying DeviceInfo resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDeviceInfoResourceData(d *schema.ResourceData, m *models.DeviceInfo) {
	d.Set("cpu_arch", m.CPUArch)
	d.Set("machine_arch", m.MachineArch)
	d.Set("mem_m_b", m.MemMB)
	d.Set("n_cpu", m.NCPU)
	d.Set("platform", m.Platform)
	d.Set("storage_m_b", m.StorageMB)
}

// Iterate through and update the DeviceInfo resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDeviceInfoSubResourceData(m []*models.DeviceInfo) (d []*map[string]interface{}) {
	for _, DeviceInfoModel := range m {
		if DeviceInfoModel != nil {
			properties := make(map[string]interface{})
			properties["cpu_arch"] = DeviceInfoModel.CPUArch
			properties["machine_arch"] = DeviceInfoModel.MachineArch
			properties["mem_m_b"] = DeviceInfoModel.MemMB
			properties["n_cpu"] = DeviceInfoModel.NCPU
			properties["platform"] = DeviceInfoModel.Platform
			properties["storage_m_b"] = DeviceInfoModel.StorageMB
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the DeviceInfo resource defined in the Terraform configuration
func DeviceInfoSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cpu_arch": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"machine_arch": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"mem_m_b": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"n_cpu": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"platform": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"storage_m_b": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the DeviceInfo resource
func GetDeviceInfoPropertyFields() (t []string) {
	return []string{
		"cpu_arch",
		"machine_arch",
		"mem_m_b",
		"n_cpu",
		"platform",
		"storage_m_b",
	}
}
