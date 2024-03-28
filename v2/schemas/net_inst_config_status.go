package schemas

import (
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func NetInstConfigStatusModel(d *schema.ResourceData) *models.NetInstConfigStatus {
	clusterID, _ := d.Get("cluster_id").(string)
	deviceDefault, _ := d.Get("device_default").(bool)
	deviceID, _ := d.Get("device_id").(string)
	deviceName, _ := d.Get("device_name").(string)
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
	projectName, _ := d.Get("project_name").(string)
	var runState *models.RunState // RunState
	runStateInterface, runStateIsSet := d.GetOk("run_state")
	if runStateIsSet {
		runStateModel := runStateInterface.(string)
		runState = models.NewRunState(models.RunState(runStateModel))
	}
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
	upTimeStamp, _ := d.Get("up_time_stamp").(strfmt.DateTime)
	uplinkIntf, _ := d.Get("uplink_intf").(string)
	return &models.NetInstConfigStatus{
		ClusterID:       clusterID,
		DeviceDefault:   deviceDefault,
		DeviceID:        deviceID,
		DeviceName:      deviceName,
		ID:              id,
		Kind:            kind,
		Name:            name,
		NetworkPolicyID: networkPolicyID,
		ProjectID:       projectID,
		ProjectName:     projectName,
		RunState:        runState,
		Tags:            tags,
		Type:            typeVar,
		UpTimeStamp:     upTimeStamp,
		UplinkIntf:      uplinkIntf,
	}
}

func NetInstConfigStatusModelFromMap(m map[string]interface{}) *models.NetInstConfigStatus {
	clusterID := m["cluster_id"].(string)
	deviceDefault := m["device_default"].(bool)
	deviceID := m["device_id"].(string)
	deviceName := m["device_name"].(string)
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
	projectName := m["project_name"].(string)
	var runState *models.RunState // RunState
	runStateInterface, runStateIsSet := m["run_state"]
	if runStateIsSet {
		runStateModel := runStateInterface.(string)
		runState = models.NewRunState(models.RunState(runStateModel))
	}
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
	upTimeStamp := m["up_time_stamp"].(strfmt.DateTime)
	uplinkIntf := m["uplink_intf"].(string)
	return &models.NetInstConfigStatus{
		ClusterID:       clusterID,
		DeviceDefault:   deviceDefault,
		DeviceID:        deviceID,
		DeviceName:      deviceName,
		ID:              id,
		Kind:            kind,
		Name:            name,
		NetworkPolicyID: networkPolicyID,
		ProjectID:       projectID,
		ProjectName:     projectName,
		RunState:        runState,
		Tags:            tags,
		Type:            typeVar,
		UpTimeStamp:     upTimeStamp,
		UplinkIntf:      uplinkIntf,
	}
}

// Update the underlying NetInstConfigStatus resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetNetInstConfigStatusResourceData(d *schema.ResourceData, m *models.NetInstConfigStatus) {
	d.Set("cluster_id", m.ClusterID)
	d.Set("device_default", m.DeviceDefault)
	d.Set("device_id", m.DeviceID)
	d.Set("device_name", m.DeviceName)
	d.Set("id", m.ID)
	d.Set("kind", m.Kind)
	d.Set("name", m.Name)
	d.Set("network_policy_id", m.NetworkPolicyID)
	d.Set("project_id", m.ProjectID)
	d.Set("project_name", m.ProjectName)
	d.Set("run_state", m.RunState)
	d.Set("tags", m.Tags)
	d.Set("type", m.Type)
	d.Set("up_time_stamp", m.UpTimeStamp)
	d.Set("uplink_intf", m.UplinkIntf)
}

// Iterate through and update the NetInstConfigStatus resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetNetInstConfigStatusSubResourceData(m []*models.NetInstConfigStatus) (d []*map[string]interface{}) {
	for _, NetInstConfigStatusModel := range m {
		if NetInstConfigStatusModel != nil {
			properties := make(map[string]interface{})
			properties["cluster_id"] = NetInstConfigStatusModel.ClusterID
			properties["device_default"] = NetInstConfigStatusModel.DeviceDefault
			properties["device_id"] = NetInstConfigStatusModel.DeviceID
			properties["device_name"] = NetInstConfigStatusModel.DeviceName
			properties["id"] = NetInstConfigStatusModel.ID
			properties["kind"] = NetInstConfigStatusModel.Kind
			properties["name"] = NetInstConfigStatusModel.Name
			properties["network_policy_id"] = NetInstConfigStatusModel.NetworkPolicyID
			properties["project_id"] = NetInstConfigStatusModel.ProjectID
			properties["project_name"] = NetInstConfigStatusModel.ProjectName
			properties["run_state"] = NetInstConfigStatusModel.RunState
			properties["tags"] = NetInstConfigStatusModel.Tags
			properties["type"] = NetInstConfigStatusModel.Type
			properties["up_time_stamp"] = NetInstConfigStatusModel.UpTimeStamp.String()
			properties["uplink_intf"] = NetInstConfigStatusModel.UplinkIntf
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the NetInstConfigStatus resource defined in the Terraform configuration
func NetInstConfigStatusSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cluster_id": {
			Description: `User defined name of the clusterInstance, unique across the enterprise.`,
			Type:        schema.TypeString,
			Optional:    true,
		},

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

		"device_name": {
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
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"project_id": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"project_name": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"run_state": {
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

		"up_time_stamp": {
			Description:  ``,
			Type:         schema.TypeString,
			ValidateFunc: validation.IsRFC3339Time,
			Optional:     true,
		},

		"uplink_intf": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the NetInstConfigStatus resource
func GetNetInstConfigStatusPropertyFields() (t []string) {
	return []string{
		"cluster_id",
		"device_default",
		"device_id",
		"device_name",
		"id",
		"kind",
		"name",
		"network_policy_id",
		"project_id",
		"project_name",
		"run_state",
		"tags",
		"type",
		"up_time_stamp",
		"uplink_intf",
	}
}
