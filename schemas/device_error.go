package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

func DeviceErrorModel(d *schema.ResourceData) *models.DeviceError {
	description, _ := d.Get("description").(string)
	var entities []*models.DeviceEntity // []*DeviceEntity
	entitiesInterface, entitiesIsSet := d.GetOk("entities")
	if entitiesIsSet {
		var items []interface{}
		if listItems, isList := entitiesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = entitiesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := DeviceEntityModelFromMap(v.(map[string]interface{}))
			entities = append(entities, m)
		}
	}
	retryCondition, _ := d.Get("retry_condition").(string)
	var severity *models.Severity // Severity
	severityInterface, severityIsSet := d.GetOk("severity")
	if severityIsSet {
		severityModel := severityInterface.(string)
		severity = models.NewSeverity(models.Severity(severityModel))
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
	var entities []*models.DeviceEntity // []*DeviceEntity
	entitiesInterface, entitiesIsSet := m["entities"]
	if entitiesIsSet {
		var items []interface{}
		if listItems, isList := entitiesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = entitiesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := DeviceEntityModelFromMap(v.(map[string]interface{}))
			entities = append(entities, m)
		}
	}
	retryCondition := m["retry_condition"].(string)
	var severity *models.Severity // Severity
	severityInterface, severityIsSet := m["severity"]
	if severityIsSet {
		severityModel := severityInterface.(string)
		severity = models.NewSeverity(models.Severity(severityModel))
	}
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
