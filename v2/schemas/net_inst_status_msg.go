package schemas

import (
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func NetInstStatusMsgModel(d *schema.ResourceData) *models.NetInstStatusMsg {
	var assignedAdapters []*models.IoBundleStatus // []*IoBundleStatus
	assignedAdaptersInterface, assignedAdaptersIsSet := d.GetOk("assigned_adapters")
	if assignedAdaptersIsSet {
		var items []interface{}
		if listItems, isList := assignedAdaptersInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = assignedAdaptersInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := IoBundleStatusModelFromMap(v.(map[string]interface{}))
			assignedAdapters = append(assignedAdapters, m)
		}
	}
	bridgeIPAddr, _ := d.Get("bridge_ip_addr").(string)
	bridgeName, _ := d.Get("bridge_name").(string)
	bridgeNumInt, _ := d.Get("bridge_num").(int)
	bridgeNum := int64(bridgeNumInt)
	clusterID, _ := d.Get("cluster_id").(string)
	deviceID, _ := d.Get("device_id").(string)
	var errInfo []*models.DeviceError // []*DeviceError
	errInfoInterface, errInfoIsSet := d.GetOk("err_info")
	if errInfoIsSet {
		var items []interface{}
		if listItems, isList := errInfoInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = errInfoInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := DeviceErrorModelFromMap(v.(map[string]interface{}))
			errInfo = append(errInfo, m)
		}
	}
	id, _ := d.Get("id").(string)
	var iPMappings []*models.IPAssignment // []*IPAssignment
	ipMappingsInterface, ipMappingsIsSet := d.GetOk("ip_mappings")
	if ipMappingsIsSet {
		var items []interface{}
		if listItems, isList := ipMappingsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = ipMappingsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := IPAssignmentModelFromMap(v.(map[string]interface{}))
			iPMappings = append(iPMappings, m)
		}
	}
	iPV4Eid, _ := d.Get("ipv4_eid").(bool)
	var kind *models.NetworkInstanceKind // NetworkInstanceKind
	kindInterface, kindIsSet := d.GetOk("kind")
	if kindIsSet {
		kindModel := kindInterface.(string)
		kind = models.NewNetworkInstanceKind(models.NetworkInstanceKind(kindModel))
	}
	name, _ := d.Get("name").(string)
	projectID, _ := d.Get("project_id").(string)
	rawStatus, _ := d.Get("raw_status").(string)
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

	upTimeStamp, _ := d.Get("up_time_stamp").(strfmt.DateTime)
	uplinkIntf, _ := d.Get("uplink_intf").(string)
	var vifs []*models.VifInfo // []*VifInfo
	vifsInterface, vifsIsSet := d.GetOk("vifs")
	if vifsIsSet {
		var items []interface{}
		if listItems, isList := vifsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = vifsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := VifInfoModelFromMap(v.(map[string]interface{}))
			vifs = append(vifs, m)
		}
	}
	return &models.NetInstStatusMsg{
		AssignedAdapters: assignedAdapters,
		BridgeIPAddr:     bridgeIPAddr,
		BridgeName:       bridgeName,
		BridgeNum:        bridgeNum,
		ClusterID:        clusterID,
		DeviceID:         deviceID,
		ErrInfo:          errInfo,
		ID:               id,
		IPMappings:       iPMappings,
		IPV4Eid:          iPV4Eid,
		Kind:             kind,
		Name:             name,
		ProjectID:        projectID,
		RawStatus:        rawStatus,
		RunState:         runState,
		Tags:             tags,
		UpTimeStamp:      upTimeStamp,
		UplinkIntf:       uplinkIntf,
		Vifs:             vifs,
	}
}

func NetInstStatusMsgModelFromMap(m map[string]interface{}) *models.NetInstStatusMsg {
	var assignedAdapters []*models.IoBundleStatus // []*IoBundleStatus
	assignedAdaptersInterface, assignedAdaptersIsSet := m["assigned_adapters"]
	if assignedAdaptersIsSet {
		var items []interface{}
		if listItems, isList := assignedAdaptersInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = assignedAdaptersInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := IoBundleStatusModelFromMap(v.(map[string]interface{}))
			assignedAdapters = append(assignedAdapters, m)
		}
	}
	bridgeIPAddr := m["bridge_ip_addr"].(string)
	bridgeName := m["bridge_name"].(string)
	bridgeNum := int64(m["bridge_num"].(int)) // int64
	clusterID := m["cluster_id"].(string)
	deviceID := m["device_id"].(string)
	var errInfo []*models.DeviceError // []*DeviceError
	errInfoInterface, errInfoIsSet := m["err_info"]
	if errInfoIsSet {
		var items []interface{}
		if listItems, isList := errInfoInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = errInfoInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := DeviceErrorModelFromMap(v.(map[string]interface{}))
			errInfo = append(errInfo, m)
		}
	}
	id := m["id"].(string)
	var iPMappings []*models.IPAssignment // []*IPAssignment
	ipMappingsInterface, ipMappingsIsSet := m["ip_mappings"]
	if ipMappingsIsSet {
		var items []interface{}
		if listItems, isList := ipMappingsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = ipMappingsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := IPAssignmentModelFromMap(v.(map[string]interface{}))
			iPMappings = append(iPMappings, m)
		}
	}
	iPV4Eid := m["ipv4_eid"].(bool)
	var kind *models.NetworkInstanceKind // NetworkInstanceKind
	kindInterface, kindIsSet := m["kind"]
	if kindIsSet {
		kindModel := kindInterface.(string)
		kind = models.NewNetworkInstanceKind(models.NetworkInstanceKind(kindModel))
	}
	name := m["name"].(string)
	projectID := m["project_id"].(string)
	rawStatus := m["raw_status"].(string)
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

	upTimeStamp := m["up_time_stamp"].(strfmt.DateTime)
	uplinkIntf := m["uplink_intf"].(string)
	var vifs []*models.VifInfo // []*VifInfo
	vifsInterface, vifsIsSet := m["vifs"]
	if vifsIsSet {
		var items []interface{}
		if listItems, isList := vifsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = vifsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := VifInfoModelFromMap(v.(map[string]interface{}))
			vifs = append(vifs, m)
		}
	}
	return &models.NetInstStatusMsg{
		AssignedAdapters: assignedAdapters,
		BridgeIPAddr:     bridgeIPAddr,
		BridgeName:       bridgeName,
		BridgeNum:        bridgeNum,
		ClusterID:        clusterID,
		DeviceID:         deviceID,
		ErrInfo:          errInfo,
		ID:               id,
		IPMappings:       iPMappings,
		IPV4Eid:          iPV4Eid,
		Kind:             kind,
		Name:             name,
		ProjectID:        projectID,
		RawStatus:        rawStatus,
		RunState:         runState,
		Tags:             tags,
		UpTimeStamp:      upTimeStamp,
		UplinkIntf:       uplinkIntf,
		Vifs:             vifs,
	}
}

// Update the underlying NetInstStatusMsg resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetNetInstStatusMsgResourceData(d *schema.ResourceData, m *models.NetInstStatusMsg) {
	d.Set("assigned_adapters", SetIoBundleStatusSubResourceData(m.AssignedAdapters))
	d.Set("bridge_ip_addr", m.BridgeIPAddr)
	d.Set("bridge_name", m.BridgeName)
	d.Set("bridge_num", m.BridgeNum)
	d.Set("cluster_id", m.ClusterID)
	d.Set("device_id", m.DeviceID)
	d.Set("err_info", SetDeviceErrorSubResourceData(m.ErrInfo))
	d.Set("id", m.ID)
	d.Set("ip_mappings", SetIPAssignmentSubResourceData(m.IPMappings))
	d.Set("ipv4_eid", m.IPV4Eid)
	d.Set("kind", m.Kind)
	d.Set("name", m.Name)
	d.Set("project_id", m.ProjectID)
	d.Set("raw_status", m.RawStatus)
	d.Set("run_state", m.RunState)
	d.Set("tags", m.Tags)
	d.Set("up_time_stamp", m.UpTimeStamp)
	d.Set("uplink_intf", m.UplinkIntf)
	d.Set("vifs", SetVifInfoSubResourceData(m.Vifs))
}

// Iterate through and update the NetInstStatusMsg resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetNetInstStatusMsgSubResourceData(m []*models.NetInstStatusMsg) (d []*map[string]interface{}) {
	for _, NetInstStatusMsgModel := range m {
		if NetInstStatusMsgModel != nil {
			properties := make(map[string]interface{})
			properties["assigned_adapters"] = SetIoBundleStatusSubResourceData(NetInstStatusMsgModel.AssignedAdapters)
			properties["bridge_ip_addr"] = NetInstStatusMsgModel.BridgeIPAddr
			properties["bridge_name"] = NetInstStatusMsgModel.BridgeName
			properties["bridge_num"] = NetInstStatusMsgModel.BridgeNum
			properties["cluster_id"] = NetInstStatusMsgModel.ClusterID
			properties["device_id"] = NetInstStatusMsgModel.DeviceID
			properties["err_info"] = SetDeviceErrorSubResourceData(NetInstStatusMsgModel.ErrInfo)
			properties["id"] = NetInstStatusMsgModel.ID
			properties["ip_mappings"] = SetIPAssignmentSubResourceData(NetInstStatusMsgModel.IPMappings)
			properties["ipv4_eid"] = NetInstStatusMsgModel.IPV4Eid
			properties["kind"] = NetInstStatusMsgModel.Kind
			properties["name"] = NetInstStatusMsgModel.Name
			properties["project_id"] = NetInstStatusMsgModel.ProjectID
			properties["raw_status"] = NetInstStatusMsgModel.RawStatus
			properties["run_state"] = NetInstStatusMsgModel.RunState
			properties["tags"] = NetInstStatusMsgModel.Tags
			properties["up_time_stamp"] = NetInstStatusMsgModel.UpTimeStamp.String()
			properties["uplink_intf"] = NetInstStatusMsgModel.UplinkIntf
			properties["vifs"] = SetVifInfoSubResourceData(NetInstStatusMsgModel.Vifs)
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the NetInstStatusMsg resource defined in the Terraform configuration
func NetInstStatusMsgSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"assigned_adapters": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []*IoBundleStatus
			Elem: &schema.Resource{
				Schema: IoBundleStatusSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"bridge_ip_addr": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"bridge_name": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"bridge_num": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"cluster_id": {
			Description: `System defined universally unique clusterInstance ID, unique across the enterprise.`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"device_id": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"err_info": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []*DeviceError
			Elem: &schema.Resource{
				Schema: DeviceErrorSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"id": {
			Description: `System defined universally unique Id of the network instance`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"ip_mappings": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []*IPAssignment
			Elem: &schema.Resource{
				Schema: IPAssignmentSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"ipv4_eid": {
			Description: ``,
			Type:        schema.TypeBool,
			Optional:    true,
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

		"project_id": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"raw_status": {
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

		"vifs": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []*VifInfo
			Elem: &schema.Resource{
				Schema: VifInfoSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
	}
}

// Retrieve property field names for updating the NetInstStatusMsg resource
func GetNetInstStatusMsgPropertyFields() (t []string) {
	return []string{
		"assigned_adapters",
		"bridge_ip_addr",
		"bridge_name",
		"bridge_num",
		"cluster_id",
		"device_id",
		"err_info",
		"id",
		"ip_mappings",
		"ipv4_eid",
		"kind",
		"name",
		"project_id",
		"raw_status",
		"run_state",
		"tags",
		"up_time_stamp",
		"uplink_intf",
		"vifs",
	}
}
