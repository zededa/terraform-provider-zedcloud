package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate ZManufacturerInfo resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ZManufacturerInfoModel(d *schema.ResourceData) *models.ZManufacturerInfo {
	biosReleaseDate := d.Get("bios_release_date").(string)
	biosVendor := d.Get("bios_vendor").(string)
	biosVersion := d.Get("bios_version").(string)
	compatible := d.Get("compatible").(string)
	endorsementKey := d.Get("endorsement_key").(string)
	hSMInfo := d.Get("h_s_m_info").(string)
	hSMStatus := d.Get("h_s_m_status").(*models.DeviceHWSecurityModuleStatus) // DeviceHWSecurityModuleStatus
	manufacturer := d.Get("manufacturer").(string)
	productName := d.Get("product_name").(string)
	serialNumber := d.Get("serial_number").(string)
	uuid := d.Get("uuid").(string)
	version := d.Get("version").(string)
	return &models.ZManufacturerInfo{
		BiosReleaseDate: biosReleaseDate,
		BiosVendor:      biosVendor,
		BiosVersion:     biosVersion,
		Compatible:      compatible,
		EndorsementKey:  endorsementKey,
		HSMInfo:         hSMInfo,
		HSMStatus:       hSMStatus,
		Manufacturer:    manufacturer,
		ProductName:     productName,
		SerialNumber:    serialNumber,
		UUID:            uuid,
		Version:         version,
	}
}

func ZManufacturerInfoModelFromMap(m map[string]interface{}) *models.ZManufacturerInfo {
	biosReleaseDate := m["bios_release_date"].(string)
	biosVendor := m["bios_vendor"].(string)
	biosVersion := m["bios_version"].(string)
	compatible := m["compatible"].(string)
	endorsementKey := m["endorsement_key"].(string)
	hSMInfo := m["h_s_m_info"].(string)
	hSMStatus := m["h_s_m_status"].(*models.DeviceHWSecurityModuleStatus) // DeviceHWSecurityModuleStatus
	manufacturer := m["manufacturer"].(string)
	productName := m["product_name"].(string)
	serialNumber := m["serial_number"].(string)
	uuid := m["uuid"].(string)
	version := m["version"].(string)
	return &models.ZManufacturerInfo{
		BiosReleaseDate: biosReleaseDate,
		BiosVendor:      biosVendor,
		BiosVersion:     biosVersion,
		Compatible:      compatible,
		EndorsementKey:  endorsementKey,
		HSMInfo:         hSMInfo,
		HSMStatus:       hSMStatus,
		Manufacturer:    manufacturer,
		ProductName:     productName,
		SerialNumber:    serialNumber,
		UUID:            uuid,
		Version:         version,
	}
}

// Update the underlying ZManufacturerInfo resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetZManufacturerInfoResourceData(d *schema.ResourceData, m *models.ZManufacturerInfo) {
	d.Set("bios_release_date", m.BiosReleaseDate)
	d.Set("bios_vendor", m.BiosVendor)
	d.Set("bios_version", m.BiosVersion)
	d.Set("compatible", m.Compatible)
	d.Set("endorsement_key", m.EndorsementKey)
	d.Set("h_s_m_info", m.HSMInfo)
	d.Set("h_s_m_status", m.HSMStatus)
	d.Set("manufacturer", m.Manufacturer)
	d.Set("product_name", m.ProductName)
	d.Set("serial_number", m.SerialNumber)
	d.Set("uuid", m.UUID)
	d.Set("version", m.Version)
}

// Iterate throught and update the ZManufacturerInfo resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetZManufacturerInfoSubResourceData(m []*models.ZManufacturerInfo) (d []*map[string]interface{}) {
	for _, ZManufacturerInfoModel := range m {
		if ZManufacturerInfoModel != nil {
			properties := make(map[string]interface{})
			properties["bios_release_date"] = ZManufacturerInfoModel.BiosReleaseDate
			properties["bios_vendor"] = ZManufacturerInfoModel.BiosVendor
			properties["bios_version"] = ZManufacturerInfoModel.BiosVersion
			properties["compatible"] = ZManufacturerInfoModel.Compatible
			properties["endorsement_key"] = ZManufacturerInfoModel.EndorsementKey
			properties["h_s_m_info"] = ZManufacturerInfoModel.HSMInfo
			properties["h_s_m_status"] = ZManufacturerInfoModel.HSMStatus
			properties["manufacturer"] = ZManufacturerInfoModel.Manufacturer
			properties["product_name"] = ZManufacturerInfoModel.ProductName
			properties["serial_number"] = ZManufacturerInfoModel.SerialNumber
			properties["uuid"] = ZManufacturerInfoModel.UUID
			properties["version"] = ZManufacturerInfoModel.Version
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the ZManufacturerInfo resource defined in the Terraform configuration
func ZManufacturerInfoSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"bios_release_date": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"bios_vendor": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"bios_version": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"compatible": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"endorsement_key": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"h_s_m_info": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"h_s_m_status": {
			Description: ``,
			Type:        schema.TypeList, //GoType: DeviceHWSecurityModuleStatus
			Elem: &schema.Resource{
				Schema: DeviceHWSecurityModuleStatusSchema(),
			},
			Optional: true,
		},

		"manufacturer": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"product_name": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"serial_number": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"uuid": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"version": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the ZManufacturerInfo resource
func GetZManufacturerInfoPropertyFields() (t []string) {
	return []string{
		"bios_release_date",
		"bios_vendor",
		"bios_version",
		"compatible",
		"endorsement_key",
		"h_s_m_info",
		"h_s_m_status",
		"manufacturer",
		"product_name",
		"serial_number",
		"uuid",
		"version",
	}
}
