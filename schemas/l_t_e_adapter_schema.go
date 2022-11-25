package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate LTEAdapter resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func LTEAdapterModel(d *schema.ResourceData) *models.LTEAdapter {
	cellModuleName := d.Get("cell_module_name").(string)
	firmwareVersion := d.Get("firmware_version").(string)
	iccid := d.Get("iccid").(string)
	imei := d.Get("imei").(string)
	imsi := d.Get("imsi").(string)
	simName := d.Get("sim_name").(string)
	simcardState := d.Get("simcard_state").(*models.SimcardState) // SimcardState
	return &models.LTEAdapter{
		CellModuleName:  cellModuleName,
		FirmwareVersion: firmwareVersion,
		Iccid:           iccid,
		Imei:            imei,
		Imsi:            imsi,
		SimName:         simName,
		SimcardState:    simcardState,
	}
}

func LTEAdapterModelFromMap(m map[string]interface{}) *models.LTEAdapter {
	cellModuleName := m["cell_module_name"].(string)
	firmwareVersion := m["firmware_version"].(string)
	iccid := m["iccid"].(string)
	imei := m["imei"].(string)
	imsi := m["imsi"].(string)
	simName := m["sim_name"].(string)
	simcardState := m["simcard_state"].(*models.SimcardState) // SimcardState
	return &models.LTEAdapter{
		CellModuleName:  cellModuleName,
		FirmwareVersion: firmwareVersion,
		Iccid:           iccid,
		Imei:            imei,
		Imsi:            imsi,
		SimName:         simName,
		SimcardState:    simcardState,
	}
}

// Update the underlying LTEAdapter resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetLTEAdapterResourceData(d *schema.ResourceData, m *models.LTEAdapter) {
	d.Set("cell_module_name", m.CellModuleName)
	d.Set("firmware_version", m.FirmwareVersion)
	d.Set("iccid", m.Iccid)
	d.Set("imei", m.Imei)
	d.Set("imsi", m.Imsi)
	d.Set("sim_name", m.SimName)
	d.Set("simcard_state", m.SimcardState)
}

// Iterate throught and update the LTEAdapter resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetLTEAdapterSubResourceData(m []*models.LTEAdapter) (d []*map[string]interface{}) {
	for _, LTEAdapterModel := range m {
		if LTEAdapterModel != nil {
			properties := make(map[string]interface{})
			properties["cell_module_name"] = LTEAdapterModel.CellModuleName
			properties["firmware_version"] = LTEAdapterModel.FirmwareVersion
			properties["iccid"] = LTEAdapterModel.Iccid
			properties["imei"] = LTEAdapterModel.Imei
			properties["imsi"] = LTEAdapterModel.Imsi
			properties["sim_name"] = LTEAdapterModel.SimName
			properties["simcard_state"] = LTEAdapterModel.SimcardState
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the LTEAdapter resource defined in the Terraform configuration
func LTEAdapterSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cell_module_name": {
			Description: `Name of Cell Module`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"firmware_version": {
			Description: `Firmware Version of Cell Radio.`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"iccid": {
			Description: `iccid of the SIM`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"imei": {
			Description: `IMEI of Cell Radio.`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"imsi": {
			Description: `imsi of the SIM`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"sim_name": {
			Description: `Name of SIM card.`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"simcard_state": {
			Description: `State of SimCard`,
			Type:        schema.TypeList, //GoType: SimcardState
			Elem: &schema.Resource{
				Schema: SimcardStateSchema(),
			},
			Optional: true,
		},
	}
}

// Retrieve property field names for updating the LTEAdapter resource
func GetLTEAdapterPropertyFields() (t []string) {
	return []string{
		"cell_module_name",
		"firmware_version",
		"iccid",
		"imei",
		"imsi",
		"sim_name",
		"simcard_state",
	}
}
