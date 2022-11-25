package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate PCRTemplate resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func PCRTemplateModel(d *schema.ResourceData) *models.PCRTemplate {
	pCRValues := d.Get("p_c_r_values").([]*models.PCRValue) // []*PCRValue
	eveVersion := d.Get("eve_version").(string)
	firmwareVersion := d.Get("firmware_version").(string)
	return &models.PCRTemplate{
		PCRValues:       pCRValues,
		EveVersion:      eveVersion,
		FirmwareVersion: firmwareVersion,
	}
}

func PCRTemplateModelFromMap(m map[string]interface{}) *models.PCRTemplate {
	pCRValues := m["p_c_r_values"].([]*models.PCRValue) // []*PCRValue
	eveVersion := m["eve_version"].(string)
	firmwareVersion := m["firmware_version"].(string)
	return &models.PCRTemplate{
		PCRValues:       pCRValues,
		EveVersion:      eveVersion,
		FirmwareVersion: firmwareVersion,
	}
}

// Update the underlying PCRTemplate resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetPCRTemplateResourceData(d *schema.ResourceData, m *models.PCRTemplate) {
	d.Set("p_c_r_values", SetPCRValueSubResourceData(m.PCRValues))
	d.Set("eve_version", m.EveVersion)
	d.Set("firmware_version", m.FirmwareVersion)
}

// Iterate throught and update the PCRTemplate resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetPCRTemplateSubResourceData(m []*models.PCRTemplate) (d []*map[string]interface{}) {
	for _, PCRTemplateModel := range m {
		if PCRTemplateModel != nil {
			properties := make(map[string]interface{})
			properties["p_c_r_values"] = SetPCRValueSubResourceData(PCRTemplateModel.PCRValues)
			properties["eve_version"] = PCRTemplateModel.EveVersion
			properties["firmware_version"] = PCRTemplateModel.FirmwareVersion
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the PCRTemplate resource defined in the Terraform configuration
func PCRTemplateSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"p_c_r_values": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []*PCRValue
			Elem: &schema.Resource{
				Schema: PCRValueSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"eve_version": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"firmware_version": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the PCRTemplate resource
func GetPCRTemplatePropertyFields() (t []string) {
	return []string{
		"p_c_r_values",
		"eve_version",
		"firmware_version",
	}
}
