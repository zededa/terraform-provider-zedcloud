package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ContainerPortMapModel(d *schema.ResourceData) *models.ContainerPortMap {
	ip, _ := d.Get("ip").(string)
	privatePortInt, _ := d.Get("private_port").(int)
	privatePort := int64(privatePortInt)
	publicPortInt, _ := d.Get("public_port").(int)
	publicPort := int64(publicPortInt)
	runtimeIP, _ := d.Get("runtime_ip").(string)
	typeVar, _ := d.Get("type").(string)
	return &models.ContainerPortMap{
		IP:          ip,
		PrivatePort: privatePort,
		PublicPort:  publicPort,
		RuntimeIP:   runtimeIP,
		Type:        typeVar,
	}
}

func ContainerPortMapModelFromMap(m map[string]interface{}) *models.ContainerPortMap {
	ip := m["ip"].(string)
	privatePort := int64(m["private_port"].(int)) // int64
	publicPort := int64(m["public_port"].(int))   // int64
	runtimeIP := m["runtime_ip"].(string)
	typeVar := m["type"].(string)
	return &models.ContainerPortMap{
		IP:          ip,
		PrivatePort: privatePort,
		PublicPort:  publicPort,
		RuntimeIP:   runtimeIP,
		Type:        typeVar,
	}
}

func SetContainerPortMapResourceData(d *schema.ResourceData, m *models.ContainerPortMap) {
	d.Set("ip", m.IP)
	d.Set("private_port", m.PrivatePort)
	d.Set("public_port", m.PublicPort)
	d.Set("runtime_ip", m.RuntimeIP)
	d.Set("type", m.Type)
}

func SetContainerPortMapSubResourceData(m []*models.ContainerPortMap) (d []*map[string]interface{}) {
	for _, ContainerPortMapModel := range m {
		if ContainerPortMapModel != nil {
			properties := make(map[string]interface{})
			properties["ip"] = ContainerPortMapModel.IP
			properties["private_port"] = ContainerPortMapModel.PrivatePort
			properties["public_port"] = ContainerPortMapModel.PublicPort
			properties["runtime_ip"] = ContainerPortMapModel.RuntimeIP
			properties["type"] = ContainerPortMapModel.Type
			d = append(d, &properties)
		}
	}
	return
}

func ContainerPortMapSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"ip": {
			Description: `container ip`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"private_port": {
			Description: `container private port`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"public_port": {
			Description: `container public port`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"runtime_ip": {
			Description: `IP address that can be used for connecting to port mapped containers inside runtime`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"type": {
			Description: `port type (tcp, udp)`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetContainerPortMapPropertyFields() (t []string) {
	return []string{
		"ip",
		"private_port",
		"public_port",
		"runtime_ip",
		"type",
	}
}
