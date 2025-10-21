package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ConfigurationLockPolicyModel(d *schema.ResourceData) *models.ConfigurationLockPolicy {
	var configLock *models.ConfigurationLock // ConfigurationLock
	configLockInterface, configLockIsSet := d.GetOk("config_lock")
	if configLockIsSet {
		configLockModel := configLockInterface.(string)
		configLock = models.NewConfigurationLock(models.ConfigurationLock(configLockModel))
	}
	id, _ := d.Get("id").(string)
	return &models.ConfigurationLockPolicy{
		ConfigLock: configLock,
		ID:         id,
	}
}

func ConfigurationLockPolicyModelFromMap(m map[string]interface{}) *models.ConfigurationLockPolicy {
	var configLock *models.ConfigurationLock // ConfigurationLock
	configLockInterface, configLockIsSet := m["config_lock"]
	if configLockIsSet {
		configLockModel := configLockInterface.(string)
		configLock = models.NewConfigurationLock(models.ConfigurationLock(configLockModel))
	}
	id := m["id"].(string)
	return &models.ConfigurationLockPolicy{
		ConfigLock: configLock,
		ID:         id,
	}
}

func SetConfigurationLockPolicyResourceData(d *schema.ResourceData, m *models.ConfigurationLockPolicy) {
	d.Set("config_lock", m.ConfigLock)
	d.Set("id", m.ID)
}

func SetConfigurationLockPolicySubResourceData(m []*models.ConfigurationLockPolicy) (d []*map[string]interface{}) {
	for _, ConfigurationLockPolicyModel := range m {
		if ConfigurationLockPolicyModel != nil {
			properties := make(map[string]interface{})
			properties["config_lock"] = ConfigurationLockPolicyModel.ConfigLock
			properties["id"] = ConfigurationLockPolicyModel.ID
			d = append(d, &properties)
		}
	}
	return
}

func ConfigurationLockPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"config_lock": {
			Description: `configuration lock setting`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"id": {
			Description: `unique policy id`,
			Type:        schema.TypeString,
			Computed:    true,
		},
	}
}

func GetConfigurationLockPolicyPropertyFields() (t []string) {
	return []string{
		"config_lock",
		"id",
	}
}
