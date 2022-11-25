package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate LispServer resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func LispServerModel(d *schema.ResourceData) *models.LispServer {
	credential := d.Get("credential").(string)
	nameOrIP := d.Get("name_or_ip").(string)
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

// Update the underlying LispServer resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetLispServerResourceData(d *schema.ResourceData, m *models.LispServer) {
	d.Set("credential", m.Credential)
	d.Set("name_or_ip", m.NameOrIP)
}

// Iterate throught and update the LispServer resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
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
func LispServerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"credential": {
			Description: `lisp credential`,
			Type:        schema.TypeString,
			Required:    true,
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
