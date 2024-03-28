package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func DevZedServerModel(d *schema.ResourceData) *models.DevZedServer {
	var eID []string
	for _, i := range d.Get("e_id").([]interface{}) {
		eID = append(eID, i.(string))
	}
	hostName, _ := d.Get("host_name").(string)
	return &models.DevZedServer{
		EID:      eID,
		HostName: &hostName, // string true false false
	}
}

func DevZedServerModelFromMap(m map[string]interface{}) *models.DevZedServer {
	var eID []string
	for _, i := range m["e_id"].([]interface{}) {
		eID = append(eID, i.(string))
	}
	hostName := m["host_name"].(string)
	return &models.DevZedServer{
		EID:      eID,
		HostName: &hostName,
	}
}

func SetDevZedServerResourceData(d *schema.ResourceData, m *models.DevZedServer) {
	d.Set("e_id", m.EID)
	d.Set("host_name", m.HostName)
}

func SetDevZedServerSubResourceData(m []*models.DevZedServer) (d []*map[string]interface{}) {
	for _, DevZedServerModel := range m {
		if DevZedServerModel != nil {
			properties := make(map[string]interface{})
			properties["e_id"] = DevZedServerModel.EID
			properties["host_name"] = DevZedServerModel.HostName
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the DevZedServer resource defined in the Terraform configuration
func DevZedServer() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"e_id": {
			Description: `EID`,
			Type:        schema.TypeList, //GoType: []string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Required: true,
		},

		"host_name": {
			Description: `Hostname for dev zed server`,
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

// Retrieve property field names for updating the DevZedServer resource
func GetDevZedServerPropertyFields() (t []string) {
	return []string{
		"e_id",
		"host_name",
	}
}
