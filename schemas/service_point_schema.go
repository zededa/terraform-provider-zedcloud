package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func ServicePointModel(d *schema.ResourceData) *models.ServicePoint {
	credential, _ := d.Get("credential").(string)
	nameOrIP, _ := d.Get("name_or_ip").(string)
	var typeVar *models.SpType // SpType
	typeInterface, typeIsSet := d.GetOk("type")
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewSpType(models.SpType(typeModel))
	}
	return &models.ServicePoint{
		Credential: credential,
		NameOrIP:   nameOrIP,
		Type:       typeVar,
	}
}

func ServicePointModelFromMap(m map[string]interface{}) *models.ServicePoint {
	credential := m["credential"].(string)
	nameOrIP := m["name_or_ip"].(string)
	var typeVar *models.SpType // SpType
	typeInterface, typeIsSet := m["type"]
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewSpType(models.SpType(typeModel))
	}
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

// Iterate through and update the ServicePoint resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
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
			Type:        schema.TypeString,
			Optional:    true,
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
