package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func NtpServerModel(d *schema.ResourceData) *models.NtpServer {
	server, _ := d.Get("server").(string)
	return &models.NtpServer{
		Server: server,
	}
}

func NtpServerModelFromMap(m map[string]interface{}) *models.NtpServer {
	server := m["server"].(string)
	return &models.NtpServer{
		Server: server,
	}
}

func SetNtpServerResourceData(d *schema.ResourceData, m *models.NtpServer) {
	d.Set("server", m.Server)
}

func SetNtpServerSubResourceData(m []*models.NtpServer) (d []*map[string]interface{}) {
	for _, NtpServerModel := range m {
		if NtpServerModel != nil {
			properties := make(map[string]interface{})
			properties["server"] = NtpServerModel.Server
			d = append(d, &properties)
		}
	}
	return
}

func NtpServerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"server": {
			Description: `NTP Server IP or FQDN`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetNtpServerPropertyFields() (t []string) {
	return []string{
		"server",
	}
}
