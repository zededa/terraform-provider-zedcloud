package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ConfigurationLockModel(d *schema.ResourceData) *models.ConfigurationLock {
	configurationLock, _ := d.Get("configuration_lock").(models.ConfigurationLock)
	return &configurationLock
}

func ConfigurationLockModelFromMap(m map[string]interface{}) *models.ConfigurationLock {
	configurationLock := m["configuration_lock"].(models.ConfigurationLock)
	return &configurationLock
}

func SetConfigurationLockResourceData(d *schema.ResourceData, m *models.ConfigurationLock) {
}

func SetConfigurationLockSubResourceData(m []*models.ConfigurationLock) (d []*map[string]interface{}) {
	for _, ConfigurationLockModel := range m {
		if ConfigurationLockModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func ConfigurationLockSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func GetConfigurationLockPropertyFields() (t []string) {
	return []string{}
}
