package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate PhyAdapter resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func PhyAdapterModel(d *schema.ResourceData) *models.PhyAdapter {
	name, _ := d.Get("name").(string)
	tags, _ := d.Get("tags").(map[string]string) // map[string]string
	var typeVar *models.IoType                   // IoType
	typeInterface, typeIsSet := d.GetOk("type")
	if typeIsSet {
		typeModel := typeInterface.(models.IoType)
		typeVar = &typeModel
	}
	return &models.PhyAdapter{
		Name: name,
		Tags: tags,
		Type: typeVar,
	}
}

func PhyAdapterModelFromMap(m map[string]interface{}) *models.PhyAdapter {
	name := m["name"].(string)
	tags := m["tags"].(map[string]string)
	typeVar := m["type"].(*models.IoType) // IoType
	return &models.PhyAdapter{
		Name: name,
		Tags: tags,
		Type: typeVar,
	}
}

// Update the underlying PhyAdapter resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetPhyAdapterResourceData(d *schema.ResourceData, m *models.PhyAdapter) {
	d.Set("name", m.Name)
	d.Set("tags", m.Tags)
	d.Set("type", m.Type)
}

// Iterate through and update the PhyAdapter resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetPhyAdapterSubResourceData(m []*models.PhyAdapter) (d []*map[string]interface{}) {
	for _, PhyAdapterModel := range m {
		if PhyAdapterModel != nil {
			properties := make(map[string]interface{})
			properties["name"] = PhyAdapterModel.Name
			properties["tags"] = PhyAdapterModel.Tags
			properties["type"] = PhyAdapterModel.Type
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the PhyAdapter resource defined in the Terraform configuration
func PhyAdapterSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Description: `Physical Adapter name`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"tags": {
			Description: `Tags are name/value pairs that enable you to categorize resources. Tag names are case insensitive with max_length 512 and min_length 3. Tag values are case sensitive with max_length 256 and min_length 3.`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"type": {
			Description: `IoType specifies the type of the Input output of the device`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the PhyAdapter resource
func GetPhyAdapterPropertyFields() (t []string) {
	return []string{
		"name",
		"tags",
		"type",
	}
}
