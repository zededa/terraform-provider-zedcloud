package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

func LispServerModel(d *schema.ResourceData) *models.LispServer {
	credential, _ := d.Get("credential").(string)
	nameOrIP, _ := d.Get("name_or_ip").(string)
	return &models.LispServer{
		Credential: &credential, // string true false false
		NameOrIP:   &nameOrIP,   // string true false false
	}
}

func LispServerModelFromMap(m map[string]interface{}) *models.LispServer {
	credential := m["credential"].(string)
	nameOrIP := m["name_or_ip"].(string)
	return &models.LispServer{
		Credential: &credential,
		NameOrIP:   &nameOrIP,
	}
}

func SetLispServerResourceData(d *schema.ResourceData, m *models.LispServer) {
	d.Set("credential", m.Credential)
	d.Set("name_or_ip", m.NameOrIP)
}

func SetLispServerSubResourceData(m []*models.LispServer) (d []*map[string]interface{}) {
	for _, LispServerModel := range m {
		if LispServerModel != nil {
			properties := make(map[string]interface{})
			properties["credential"] = LispServerModel.Credential
			properties["name_or_ip"] = LispServerModel.NameOrIP
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the LispServer resource defined in the Terraform configuration
func LispServer() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"credential": {
			Description: `lisp credential`,
			Type:        schema.TypeString,
			Required:    true,
			Sensitive:   true,
		},

		"name_or_ip": {
			Description: `name/IP`,
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

// Retrieve property field names for updating the LispServer resource
func GetLispServerPropertyFields() (t []string) {
	return []string{
		"credential",
		"name_or_ip",
	}
}
