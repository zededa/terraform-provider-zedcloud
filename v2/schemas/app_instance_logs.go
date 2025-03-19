package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func AppInstanceLogsModel(d *schema.ResourceData) *models.AppInstanceLogs {
	access, _ := d.Get("access").(bool)
	return &models.AppInstanceLogs{
		Access: &access, // bool
	}
}

func AppInstanceLogsModelFromMap(m map[string]interface{}) *models.AppInstanceLogs {
	access := m["access"].(bool)
	return &models.AppInstanceLogs{
		Access: &access,
	}
}

func SetAppInstanceLogsResourceData(d *schema.ResourceData, m *models.AppInstanceLogs) {
	d.Set("access", m.Access)
}

func SetAppInstanceLogsSubResourceData(m []*models.AppInstanceLogs) (d []*map[string]interface{}) {
	for _, AppInstanceLogsModel := range m {
		if AppInstanceLogsModel != nil {
			properties := make(map[string]interface{})
			properties["access"] = AppInstanceLogsModel.Access
			d = append(d, &properties)
		}
	}
	return
}

func AppInstanceLogsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"access": {
			Description: `Flags to enable / disable sending of logs generated by app instance to zedcloud`,
			Type:        schema.TypeBool,
			Required:    true,
		},
	}
}

func GetAppInstanceLogsPropertyFields() (t []string) {
	return []string{
		"access",
	}
}
