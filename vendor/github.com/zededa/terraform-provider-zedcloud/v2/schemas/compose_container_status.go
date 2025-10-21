package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ComposeContainerStatusModel(d *schema.ResourceData) *models.ComposeContainerStatus {
	containerEntryPoint, _ := d.Get("container_entry_point").(string)
	containerImage, _ := d.Get("container_image").(string)
	containerName, _ := d.Get("container_name").(string)
	containerState, _ := d.Get("container_state").(string)
	createTimeUTC, _ := d.Get("create_time_u_t_c").(string)
	var errInfo *models.DeviceError // DeviceError
	errInfoInterface, errInfoIsSet := d.GetOk("err_info")
	if errInfoIsSet && errInfoInterface != nil {
		errInfoMap := errInfoInterface.([]interface{})
		if len(errInfoMap) > 0 {
			errInfo = DeviceErrorModelFromMap(errInfoMap[0].(map[string]interface{}))
		}
	}
	var portMaps []*models.ContainerPortMap // []*ContainerPortMap
	portMapsInterface, portMapsIsSet := d.GetOk("port_maps")
	if portMapsIsSet {
		var items []interface{}
		if listItems, isList := portMapsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = portMapsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := ContainerPortMapModelFromMap(v.(map[string]interface{}))
			portMaps = append(portMaps, m)
		}
	}
	uptime, _ := d.Get("uptime").(uint64)
	return &models.ComposeContainerStatus{
		ContainerEntryPoint: containerEntryPoint,
		ContainerImage:      containerImage,
		ContainerName:       containerName,
		ContainerState:      containerState,
		CreateTimeUTC:       createTimeUTC,
		ErrInfo:             errInfo,
		PortMaps:            portMaps,
		Uptime:              uptime,
	}
}

func ComposeContainerStatusModelFromMap(m map[string]interface{}) *models.ComposeContainerStatus {
	containerEntryPoint := m["container_entry_point"].(string)
	containerImage := m["container_image"].(string)
	containerName := m["container_name"].(string)
	containerState := m["container_state"].(string)
	createTimeUTC := m["create_time_u_t_c"].(string)
	var errInfo *models.DeviceError // DeviceError
	errInfoInterface, errInfoIsSet := m["err_info"]
	if errInfoIsSet && errInfoInterface != nil {
		errInfoMap := errInfoInterface.([]interface{})
		if len(errInfoMap) > 0 {
			errInfo = DeviceErrorModelFromMap(errInfoMap[0].(map[string]interface{}))
		}
	}
	//
	var portMaps []*models.ContainerPortMap // []*ContainerPortMap
	portMapsInterface, portMapsIsSet := m["port_maps"]
	if portMapsIsSet {
		var items []interface{}
		if listItems, isList := portMapsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = portMapsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := ContainerPortMapModelFromMap(v.(map[string]interface{}))
			portMaps = append(portMaps, m)
		}
	}
	uptime := m["uptime"].(uint64)
	return &models.ComposeContainerStatus{
		ContainerEntryPoint: containerEntryPoint,
		ContainerImage:      containerImage,
		ContainerName:       containerName,
		ContainerState:      containerState,
		CreateTimeUTC:       createTimeUTC,
		ErrInfo:             errInfo,
		PortMaps:            portMaps,
		Uptime:              uptime,
	}
}

func SetComposeContainerStatusResourceData(d *schema.ResourceData, m *models.ComposeContainerStatus) {
	d.Set("container_entry_point", m.ContainerEntryPoint)
	d.Set("container_image", m.ContainerImage)
	d.Set("container_name", m.ContainerName)
	d.Set("container_state", m.ContainerState)
	d.Set("create_time_u_t_c", m.CreateTimeUTC)
	d.Set("err_info", SetDeviceErrorSubResourceData([]*models.DeviceError{m.ErrInfo}))
	d.Set("port_maps", SetContainerPortMapSubResourceData(m.PortMaps))
	d.Set("uptime", m.Uptime)
}

func SetComposeContainerStatusSubResourceData(m []*models.ComposeContainerStatus) (d []*map[string]interface{}) {
	for _, ComposeContainerStatusModel := range m {
		if ComposeContainerStatusModel != nil {
			properties := make(map[string]interface{})
			properties["container_entry_point"] = ComposeContainerStatusModel.ContainerEntryPoint
			properties["container_image"] = ComposeContainerStatusModel.ContainerImage
			properties["container_name"] = ComposeContainerStatusModel.ContainerName
			properties["container_state"] = ComposeContainerStatusModel.ContainerState
			properties["create_time_u_t_c"] = ComposeContainerStatusModel.CreateTimeUTC
			properties["err_info"] = SetDeviceErrorSubResourceData([]*models.DeviceError{ComposeContainerStatusModel.ErrInfo})
			properties["port_maps"] = SetContainerPortMapSubResourceData(ComposeContainerStatusModel.PortMaps)
			properties["uptime"] = ComposeContainerStatusModel.Uptime
			d = append(d, &properties)
		}
	}
	return
}

func ComposeContainerStatusSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"container_entry_point": {
			Description: `container entry point`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"container_image": {
			Description: `container image`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"container_name": {
			Description: `container name`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"container_state": {
			Description: `running status of container (running, stopped, exited)`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"create_time_u_t_c": {
			Description: `container creation time - UTC`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"err_info": {
			Description: `container error information`,
			Type:        schema.TypeList, //GoType: DeviceError
			Elem: &schema.Resource{
				Schema: DeviceErrorSchema(),
			},
			Optional: true,
		},

		"port_maps": {
			Description: `ports mapped to container`,
			Type:        schema.TypeList, //GoType: []*ContainerPortMap
			Elem: &schema.Resource{
				Schema: ContainerPortMapSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"uptime": {
			Description: `container up-time in seconds from start`,
			Type:        schema.TypeInt,
			Optional:    true,
		},
	}
}

func GetComposeContainerStatusPropertyFields() (t []string) {
	return []string{
		"container_entry_point",
		"container_image",
		"container_name",
		"container_state",
		"create_time_u_t_c",
		"err_info",
		"port_maps",
		"uptime",
	}
}
