package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func NetInstShortConfigModel(d *schema.ResourceData) *models.NetInstShortConfig {
	deviceDefault, _ := d.Get("device_default").(bool)
	deviceID, _ := d.Get("device_id").(string)
	id, _ := d.Get("id").(string)
	var kind *models.NetworkInstanceKind // NetworkInstanceKind
	kindInterface, kindIsSet := d.GetOk("kind")
	if kindIsSet {
		kindModel := kindInterface.(string)
		kind = models.NewNetworkInstanceKind(models.NetworkInstanceKind(kindModel))
	}
	name, _ := d.Get("name").(string)
	networkPolicyID, _ := d.Get("network_policy_id").(string)
	projectID, _ := d.Get("project_id").(string)
	tags := map[string]string{}
	tagsInterface, tagsIsSet := d.GetOk("tags")
	if tagsIsSet {
		tagsMap := tagsInterface.(map[string]interface{})
		for k, v := range tagsMap {
			if v == nil {
				continue
			}
			tags[k] = v.(string)
		}
	}

	var typeVar *models.NetworkInstanceDhcpType // NetworkInstanceDhcpType
	typeInterface, typeIsSet := d.GetOk("type")
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewNetworkInstanceDhcpType(models.NetworkInstanceDhcpType(typeModel))
	}
	return &models.NetInstShortConfig{
		DeviceDefault:   deviceDefault,
		DeviceID:        deviceID,
		ID:              id,
		Kind:            kind,
		Name:            name,
		NetworkPolicyID: networkPolicyID,
		ProjectID:       projectID,
		Tags:            tags,
		Type:            typeVar,
	}
}

func NetInstShortConfigModelFromMap(m map[string]interface{}) *models.NetInstShortConfig {
	deviceDefault := m["device_default"].(bool)
	deviceID := m["device_id"].(string)
	id := m["id"].(string)
	var kind *models.NetworkInstanceKind // NetworkInstanceKind
	kindInterface, kindIsSet := m["kind"]
	if kindIsSet {
		kindModel := kindInterface.(string)
		kind = models.NewNetworkInstanceKind(models.NetworkInstanceKind(kindModel))
	}
	name := m["name"].(string)
	networkPolicyID := m["network_policy_id"].(string)
	projectID := m["project_id"].(string)
	tags := map[string]string{}
	tagsInterface, tagsIsSet := m["tags"]
	if tagsIsSet {
		tagsMap := tagsInterface.(map[string]interface{})
		for k, v := range tagsMap {
			if v == nil {
				continue
			}
			tags[k] = v.(string)
		}
	}

	var typeVar *models.NetworkInstanceDhcpType // NetworkInstanceDhcpType
	typeInterface, typeIsSet := m["type"]
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewNetworkInstanceDhcpType(models.NetworkInstanceDhcpType(typeModel))
	}
	return &models.NetInstShortConfig{
		DeviceDefault:   deviceDefault,
		DeviceID:        deviceID,
		ID:              id,
		Kind:            kind,
		Name:            name,
		NetworkPolicyID: networkPolicyID,
		ProjectID:       projectID,
		Tags:            tags,
		Type:            typeVar,
	}
}

// Update the underlying NetInstShortConfig resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetNetInstShortConfigResourceData(d *schema.ResourceData, m *models.NetInstShortConfig) {
	d.Set("device_default", m.DeviceDefault)
	d.Set("device_id", m.DeviceID)
	d.Set("id", m.ID)
	d.Set("kind", m.Kind)
	d.Set("name", m.Name)
	d.Set("network_policy_id", m.NetworkPolicyID)
	d.Set("project_id", m.ProjectID)
	d.Set("tags", m.Tags)
	d.Set("type", m.Type)
}

// Iterate through and update the NetInstShortConfig resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetNetInstShortConfigSubResourceData(m []*models.NetInstShortConfig) (d []*map[string]interface{}) {
	for _, NetInstShortConfigModel := range m {
		if NetInstShortConfigModel != nil {
			properties := make(map[string]interface{})
			properties["device_default"] = NetInstShortConfigModel.DeviceDefault
			properties["device_id"] = NetInstShortConfigModel.DeviceID
			properties["id"] = NetInstShortConfigModel.ID
			properties["kind"] = NetInstShortConfigModel.Kind
			properties["name"] = NetInstShortConfigModel.Name
			properties["network_policy_id"] = NetInstShortConfigModel.NetworkPolicyID
			properties["project_id"] = NetInstShortConfigModel.ProjectID
			properties["tags"] = NetInstShortConfigModel.Tags
			properties["type"] = NetInstShortConfigModel.Type
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the NetInstShortConfig resource defined in the Terraform configuration
func NetInstShortConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"device_default": {
			Description: ``,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"device_id": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"id": {
			Description: `System defined universally unique Id of the network instance`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"kind": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"name": {
			Description: `User defined name of the network instance, unique across the enterprise. Once object is created, name can’t be changed`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"network_policy_id": {
			Description: `network policy id`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"project_id": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"tags": {
			Description: `Tags are name/value pairs that enable you to categorize resources. Tag names are case insensitive with max_length 512 and min_length 3. Tag values are case sensitive with max_length 256 and min_length 3.`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"type": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the NetInstShortConfig resource
func GetNetInstShortConfigPropertyFields() (t []string) {
	return []string{
		"device_default",
		"device_id",
		"id",
		"kind",
		"name",
		"network_policy_id",
		"project_id",
		"tags",
		"type",
	}
}
