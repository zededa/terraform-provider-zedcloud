package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate ServicePoint resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ServicePointModel(d *schema.ResourceData) *models.ServicePoint {
	credential := d.Get("credential").(string)
	nameOrIP := d.Get("name_or_ip").(string)
	typeVar := d.Get("type").(*models.SpType) // SpType
	return &models.ServicePoint{
		Credential: credential,
		NameOrIP:   nameOrIP,
		Type:       typeVar,
	}
}

func ServicePointModelFromMap(m map[string]interface{}) *models.ServicePoint {
	credential := m["credential"].(string)
	nameOrIP := m["name_or_ip"].(string)
	typeVar := m["type"].(*models.SpType) // SpType
	return &models.ServicePoint{
		Credential: credential,
		NameOrIP:   nameOrIP,
		Type:       typeVar,
	}
}

// Update the underlying ServicePoint resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetServicePointResourceData(d *schema.ResourceData, m *models.ServicePoint) {
	d.Set("credential", m.Credential)
	d.Set("name_or_ip", m.NameOrIP)
	d.Set("type", m.Type)
}

// Iterate throught and update the ServicePoint resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetServicePointSubResourceData(m []*models.ServicePoint) (d []*map[string]interface{}) {
	for _, ServicePointModel := range m {
		if ServicePointModel != nil {
			properties := make(map[string]interface{})
			properties["credential"] = ServicePointModel.Credential
			properties["name_or_ip"] = ServicePointModel.NameOrIP
			properties["type"] = ServicePointModel.Type
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the ServicePoint resource defined in the Terraform configuration
func ServicePointSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"credential": {
			Description: `Service credentials`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"name_or_ip": {
			Description: `Service name/ service name`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"type": {
			Description: `Service Point Type`,
			Type:        schema.TypeList, //GoType: SpType
			Elem: &schema.Resource{
				Schema: SpTypeSchema(),
			},
			Optional: true,
		},
	}
}

// Retrieve property field names for updating the ServicePoint resource
func GetServicePointPropertyFields() (t []string) {
	return []string{
		"credential",
		"name_or_ip",
		"type",
	}
}
