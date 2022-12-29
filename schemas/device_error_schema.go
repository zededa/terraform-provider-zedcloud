package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate DeviceError resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func DeviceErrorModel(d *schema.ResourceData) *models.DeviceError {
	description, _ := d.Get("description").(string)
	entities, _ := d.Get("entities").([]*models.DeviceEntity) // []*DeviceEntity
	retryCondition, _ := d.Get("retry_condition").(string)
	severityModel, ok := d.Get("severity").(models.Severity) // Severity
	severity := &severityModel
	if !ok {
		severity = nil
	}
	timestamp, _ := d.Get("timestamp").(interface{}) // interface{}
	return &models.DeviceError{
		Description:    &description, // string true false false
		Entities:       entities,
		RetryCondition: retryCondition,
		Severity:       severity,
		Timestamp:      timestamp,
	}
}

func DeviceErrorModelFromMap(m map[string]interface{}) *models.DeviceError {
	description := m["description"].(string)
	entities := m["entities"].([]*models.DeviceEntity) // []*DeviceEntity
	retryCondition := m["retry_condition"].(string)
	severity := m["severity"].(*models.Severity) // Severity
	timestamp := m["timestamp"].(interface{})
	return &models.DeviceError{
		Description:    &description,
		Entities:       entities,
		RetryCondition: retryCondition,
		Severity:       severity,
		Timestamp:      timestamp,
	}
}

// Update the underlying DeviceError resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDeviceErrorResourceData(d *schema.ResourceData, m *models.DeviceError) {
	d.Set("description", m.Description)
	d.Set("entities", SetDeviceEntitySubResourceData(m.Entities))
	d.Set("retry_condition", m.RetryCondition)
	d.Set("severity", m.Severity)
	d.Set("timestamp", m.Timestamp)
}

// Iterate through and update the DeviceError resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDeviceErrorSubResourceData(m []*models.DeviceError) (d []*map[string]interface{}) {
	for _, DeviceErrorModel := range m {
		if DeviceErrorModel != nil {
			properties := make(map[string]interface{})
			properties["description"] = DeviceErrorModel.Description
			properties["entities"] = SetDeviceEntitySubResourceData(DeviceErrorModel.Entities)
			properties["retry_condition"] = DeviceErrorModel.RetryCondition
			properties["severity"] = DeviceErrorModel.Severity
			properties["timestamp"] = DeviceErrorModel.Timestamp
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the DeviceError resource defined in the Terraform configuration
func DeviceErrorSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"description": {
			Description: `Description of the error`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"entities": {
			Description: `objects referenced by the description or retry_condition`,
			Type:        schema.TypeList, //GoType: []*DeviceEntity
			Elem: &schema.Resource{
				Schema: DeviceEntitySchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Required: true,
		},

		"retry_condition": {
			Description: `condition for retry`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"severity": {
			Description: `Severity of the error`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"timestamp": {
			Description: `Timestamp at which error had occurred`,
			Type:        schema.TypeMap, //GoType: interface{}
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Required: true,
		},
	}
}

// Retrieve property field names for updating the DeviceError resource
func GetDeviceErrorPropertyFields() (t []string) {
	return []string{
		"description",
		"entities",
		"retry_condition",
		"severity",
		"timestamp",
	}
}
