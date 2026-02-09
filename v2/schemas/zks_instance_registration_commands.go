package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ZksInstanceRegistrationCommandsModel(d *schema.ResourceData) *models.ZksInstanceRegistrationCommands {
	command, _ := d.Get("command").(string)
	insecureCommand, _ := d.Get("insecure_command").(string)
	nodeCommand, _ := d.Get("node_command").(string)
	windowsNodeCommand, _ := d.Get("windows_node_command").(string)
	return &models.ZksInstanceRegistrationCommands{
		Command:            command,
		InsecureCommand:    insecureCommand,
		NodeCommand:        nodeCommand,
		WindowsNodeCommand: windowsNodeCommand,
	}
}

func ZksInstanceRegistrationCommandsModelFromMap(m map[string]interface{}) *models.ZksInstanceRegistrationCommands {
	command := m["command"].(string)
	insecureCommand := m["insecure_command"].(string)
	nodeCommand := m["node_command"].(string)
	windowsNodeCommand := m["windows_node_command"].(string)
	return &models.ZksInstanceRegistrationCommands{
		Command:            command,
		InsecureCommand:    insecureCommand,
		NodeCommand:        nodeCommand,
		WindowsNodeCommand: windowsNodeCommand,
	}
}

func SetZksInstanceRegistrationCommandsResourceData(d *schema.ResourceData, m *models.ZksInstanceRegistrationCommands) {
	d.Set("command", m.Command)
	d.Set("insecure_command", m.InsecureCommand)
	d.Set("node_command", m.NodeCommand)
	d.Set("windows_node_command", m.WindowsNodeCommand)
}

func SetZksInstanceRegistrationCommandsSubResourceData(m []*models.ZksInstanceRegistrationCommands) (d []*map[string]interface{}) {
	for _, ZksInstanceRegistrationCommandsModel := range m {
		if ZksInstanceRegistrationCommandsModel != nil {
			properties := make(map[string]interface{})
			properties["command"] = ZksInstanceRegistrationCommandsModel.Command
			properties["insecure_command"] = ZksInstanceRegistrationCommandsModel.InsecureCommand
			properties["node_command"] = ZksInstanceRegistrationCommandsModel.NodeCommand
			properties["windows_node_command"] = ZksInstanceRegistrationCommandsModel.WindowsNodeCommand
			d = append(d, &properties)
		}
	}
	return
}

func ZksInstanceRegistrationCommandsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"command": {
			Description: `Command to register a node with the ZKS instance`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"insecure_command": {
			Description: `Insecure command to register a node with the ZKS instance`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"node_command": {
			Description: `Node command to register a node with the ZKS instance`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"windows_node_command": {
			Description: `Windows node command to register a node with the ZKS instance`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetZksInstanceRegistrationCommandsPropertyFields() (t []string) {
	return []string{
		"command",
		"insecure_command",
		"node_command",
		"windows_node_command",
	}
}
