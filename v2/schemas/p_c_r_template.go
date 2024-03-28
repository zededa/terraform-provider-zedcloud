package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

// Function to perform the following actions:
// (1) Translate PCRTemplate resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func PCRTemplateModel(d *schema.ResourceData) *models.PCRTemplate {
	pCRValues, _ := d.Get("p_c_r_values").([]*models.PCRValue) // []*PCRValue
	eveVersion, _ := d.Get("eve_version").(string)
	firmwareVersion, _ := d.Get("firmware_version").(string)
	id, _ := d.Get("id").(string)
	name, _ := d.Get("name").(string)
	return &models.PCRTemplate{
		PCRValues:       pCRValues,
		EveVersion:      &eveVersion, // string true false false
		FirmwareVersion: firmwareVersion,
		ID:              id,
		Name:            name,
	}
}

func PCRTemplateModelFromMap(m map[string]interface{}) *models.PCRTemplate {
	pCRValues := m["p_c_r_values"].([]*models.PCRValue) // []*PCRValue
	eveVersion := m["eve_version"].(string)
	firmwareVersion := m["firmware_version"].(string)
	id := m["id"].(string)
	name := m["name"].(string)
	return &models.PCRTemplate{
		PCRValues:       pCRValues,
		EveVersion:      &eveVersion,
		FirmwareVersion: firmwareVersion,
		ID:              id,
		Name:            name,
	}
}

// Update the underlying PCRTemplate resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetPCRTemplateResourceData(d *schema.ResourceData, m *models.PCRTemplate) {
	d.Set("p_c_r_values", SetPCRValueSubResourceData(m.PCRValues))
	d.Set("eve_version", m.EveVersion)
	d.Set("firmware_version", m.FirmwareVersion)
	d.Set("id", m.ID)
	d.Set("name", m.Name)
}

// Iterate through and update the PCRTemplate resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetPCRTemplateSubResourceData(m []*models.PCRTemplate) (d []*map[string]interface{}) {
	for _, PCRTemplateModel := range m {
		if PCRTemplateModel != nil {
			properties := make(map[string]interface{})
			properties["p_c_r_values"] = SetPCRValueSubResourceData(PCRTemplateModel.PCRValues)
			properties["eve_version"] = PCRTemplateModel.EveVersion
			properties["firmware_version"] = PCRTemplateModel.FirmwareVersion
			properties["id"] = PCRTemplateModel.ID
			properties["name"] = PCRTemplateModel.Name
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the PCRTemplate resource defined in the Terraform configuration
func PCRTemplateSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"p_c_r_values": {
			Description: `List of PCR values for the PCR template`,
			Type:        schema.TypeList, //GoType: []*PCRValue
			Elem: &schema.Resource{
				Schema: PCRValueSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Required: true,
		},

		"eve_version": {
			Description: `EVE version related to the PCR template`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"firmware_version": {
			Description: `Firmware version related to the PCR template. If user doesn't set it, it will be set to '*'`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"id": {
			Description: `System defined universally unique Id of the PCR template. If not set in POST / PUT API calls, this will be treated as a new entry and a unique System Generated ID assigned.`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"name": {
			Description: `Name of the PCR template. The name is Unique among PCR templates for that System Model. If it is not specified, a system-generated name will be assigned.`,
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
		"id",
		"name",
	}
}
