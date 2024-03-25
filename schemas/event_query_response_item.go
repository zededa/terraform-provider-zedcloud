package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

// Function to perform the following actions:
// (1) Translate EventQueryResponseItem resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func EventQueryResponseItemModel(d *schema.ResourceData) *models.EventQueryResponseItem {
	clusterInstance, _ := d.Get("cluster_instance").(string)
	description, _ := d.Get("description").(string)
	device, _ := d.Get("device").(string)
	eventType, _ := d.Get("event_type").(string)
	instance, _ := d.Get("instance").(string)
	jSONData, _ := d.Get("json_data").(interface{}) // interface{}
	project, _ := d.Get("project").(string)
	resource, _ := d.Get("resource").(string)
	resourceName, _ := d.Get("resource_name").(string)
	severity, _ := d.Get("severity").(string)
	var source *models.EventSource // EventSource
	sourceInterface, sourceIsSet := d.GetOk("source")
	if sourceIsSet {
		sourceModel := sourceInterface.(string)
		source = models.NewEventSource(models.EventSource(sourceModel))
	}
	tags, _ := d.Get("tags").([]string)
	timestamp, _ := d.Get("timestamp").(interface{}) // interface{}
	user, _ := d.Get("user").(string)
	return &models.EventQueryResponseItem{
		ClusterInstance: clusterInstance,
		Description:     description,
		Device:          device,
		EventType:       eventType,
		Instance:        instance,
		JSONData:        jSONData,
		Project:         project,
		Resource:        resource,
		ResourceName:    resourceName,
		Severity:        severity,
		Source:          source,
		Tags:            tags,
		Timestamp:       timestamp,
		User:            user,
	}
}

func EventQueryResponseItemModelFromMap(m map[string]interface{}) *models.EventQueryResponseItem {
	clusterInstance := m["cluster_instance"].(string)
	description := m["description"].(string)
	device := m["device"].(string)
	eventType := m["event_type"].(string)
	instance := m["instance"].(string)
	jSONData := m["json_data"].(interface{})
	project := m["project"].(string)
	resource := m["resource"].(string)
	resourceName := m["resource_name"].(string)
	severity := m["severity"].(string)
	source := m["source"].(*models.EventSource) // EventSource
	tags := m["tags"].([]string)
	timestamp := m["timestamp"].(interface{})
	user := m["user"].(string)
	return &models.EventQueryResponseItem{
		ClusterInstance: clusterInstance,
		Description:     description,
		Device:          device,
		EventType:       eventType,
		Instance:        instance,
		JSONData:        jSONData,
		Project:         project,
		Resource:        resource,
		ResourceName:    resourceName,
		Severity:        severity,
		Source:          source,
		Tags:            tags,
		Timestamp:       timestamp,
		User:            user,
	}
}

// Update the underlying EventQueryResponseItem resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetEventQueryResponseItemResourceData(d *schema.ResourceData, m *models.EventQueryResponseItem) {
	d.Set("cluster_instance", m.ClusterInstance)
	d.Set("description", m.Description)
	d.Set("device", m.Device)
	d.Set("event_type", m.EventType)
	d.Set("instance", m.Instance)
	d.Set("json_data", m.JSONData)
	d.Set("project", m.Project)
	d.Set("resource", m.Resource)
	d.Set("resource_name", m.ResourceName)
	d.Set("severity", m.Severity)
	d.Set("source", m.Source)
	d.Set("tags", m.Tags)
	d.Set("timestamp", m.Timestamp)
	d.Set("user", m.User)
}

// Iterate through and update the EventQueryResponseItem resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetEventQueryResponseItemSubResourceData(m []*models.EventQueryResponseItem) (d []*map[string]interface{}) {
	for _, EventQueryResponseItemModel := range m {
		if EventQueryResponseItemModel != nil {
			properties := make(map[string]interface{})
			properties["cluster_instance"] = EventQueryResponseItemModel.ClusterInstance
			properties["description"] = EventQueryResponseItemModel.Description
			properties["device"] = EventQueryResponseItemModel.Device
			properties["event_type"] = EventQueryResponseItemModel.EventType
			properties["instance"] = EventQueryResponseItemModel.Instance
			properties["json_data"] = EventQueryResponseItemModel.JSONData
			properties["project"] = EventQueryResponseItemModel.Project
			properties["resource"] = EventQueryResponseItemModel.Resource
			properties["resource_name"] = EventQueryResponseItemModel.ResourceName
			properties["severity"] = EventQueryResponseItemModel.Severity
			properties["source"] = EventQueryResponseItemModel.Source
			properties["tags"] = EventQueryResponseItemModel.Tags
			properties["timestamp"] = EventQueryResponseItemModel.Timestamp
			properties["user"] = EventQueryResponseItemModel.User
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the EventQueryResponseItem resource defined in the Terraform configuration
func EventQueryResponseItemSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cluster_instance": {
			Description: `clusterInstance name`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"description": {
			Description: `Event description`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"device": {
			Description: `device name`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"event_type": {
			Description: `Event type`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"instance": {
			Description: `instance name`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"json_data": {
			Description: `Event resources`,
			Type:        schema.TypeMap, //GoType: interface{}
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"project": {
			Description: `project name`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"resource": {
			Description: `Event resources`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"resource_name": {
			Description: `Event resources`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"severity": {
			Description: `severity: FIXME: should be ENUM`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"source": {
			Description: `source`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"tags": {
			Description: `Event tags`,
			Type:        schema.TypeList, //GoType: []string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"timestamp": {
			Description: `event timestamp`,
			Type:        schema.TypeMap, //GoType: interface{}
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"user": {
			Description: `User name`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the EventQueryResponseItem resource
func GetEventQueryResponseItemPropertyFields() (t []string) {
	return []string{
		"cluster_instance",
		"description",
		"device",
		"event_type",
		"instance",
		"json_data",
		"project",
		"resource",
		"resource_name",
		"severity",
		"source",
		"tags",
		"timestamp",
		"user",
	}
}
