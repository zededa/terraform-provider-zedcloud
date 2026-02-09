package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ZksInstanceRegistrationInfoModel(d *schema.ResourceData) *models.ZksInstanceRegistrationInfo {
	var commands *models.ZksInstanceRegistrationCommands // ZksInstanceRegistrationCommands
	commandsInterface, commandsIsSet := d.GetOk("commands")
	if commandsIsSet && commandsInterface != nil {
		commandsMap := commandsInterface.([]interface{})
		if len(commandsMap) > 0 {
			commands = ZksInstanceRegistrationCommandsModelFromMap(commandsMap[0].(map[string]interface{}))
		}
	}
	created, _ := d.Get("created").(string)
	manifestURL, _ := d.Get("manifest_url").(string)
	return &models.ZksInstanceRegistrationInfo{
		Commands:    commands,
		Created:     created,
		ManifestURL: manifestURL,
	}
}

func ZksInstanceRegistrationInfoModelFromMap(m map[string]interface{}) *models.ZksInstanceRegistrationInfo {
	var commands *models.ZksInstanceRegistrationCommands // ZksInstanceRegistrationCommands
	commandsInterface, commandsIsSet := m["commands"]
	if commandsIsSet && commandsInterface != nil {
		commandsMap := commandsInterface.([]interface{})
		if len(commandsMap) > 0 {
			commands = ZksInstanceRegistrationCommandsModelFromMap(commandsMap[0].(map[string]interface{}))
		}
	}
	//
	created := m["created"].(string)
	manifestURL := m["manifest_url"].(string)
	return &models.ZksInstanceRegistrationInfo{
		Commands:    commands,
		Created:     created,
		ManifestURL: manifestURL,
	}
}

func SetZksInstanceRegistrationInfoResourceData(d *schema.ResourceData, m *models.ZksInstanceRegistrationInfo) {
	d.Set("commands", SetZksInstanceRegistrationCommandsSubResourceData([]*models.ZksInstanceRegistrationCommands{m.Commands}))
	d.Set("created", m.Created)
	d.Set("manifest_url", m.ManifestURL)
}

func SetZksInstanceRegistrationInfoSubResourceData(m []*models.ZksInstanceRegistrationInfo) (d []*map[string]interface{}) {
	for _, ZksInstanceRegistrationInfoModel := range m {
		if ZksInstanceRegistrationInfoModel != nil {
			properties := make(map[string]interface{})
			properties["commands"] = SetZksInstanceRegistrationCommandsSubResourceData([]*models.ZksInstanceRegistrationCommands{ZksInstanceRegistrationInfoModel.Commands})
			properties["created"] = ZksInstanceRegistrationInfoModel.Created
			properties["manifest_url"] = ZksInstanceRegistrationInfoModel.ManifestURL
			d = append(d, &properties)
		}
	}
	return
}

func ZksInstanceRegistrationInfoSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"commands": {
			Description: `Registration commands for the ZKS instance`,
			Type:        schema.TypeList, //GoType: ZksInstanceRegistrationCommands
			Elem: &schema.Resource{
				Schema: ZksInstanceRegistrationCommandsSchema(),
			},
			Optional: true,
		},

		"created": {
			Description: `Timestamp when the registration commands were created`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"manifest_url": {
			Description: `URL to fetch the registration manifest`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetZksInstanceRegistrationInfoPropertyFields() (t []string) {
	return []string{
		"commands",
		"created",
		"manifest_url",
	}
}
